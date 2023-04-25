package config

import (
	"fmt"
	"github.com/anytypeio/any-sync/commonspace"
	"github.com/anytypeio/any-sync/metric"
	commonnet "github.com/anytypeio/any-sync/net"
	"github.com/anytypeio/any-sync/nodeconf"
	"github.com/anytypeio/go-anytype-middleware/util/files"
	"github.com/anytypeio/go-anytype-middleware/util/netutil"
	"gopkg.in/yaml.v2"
	"net"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kelseyhightower/envconfig"

	"github.com/anytypeio/any-sync/app"
	"github.com/anytypeio/go-anytype-middleware/core/wallet"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/datastore/clientds"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/ipfs"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/logging"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/threads"
)

var log = logging.Logger("anytype-config")

const (
	CName = "config"
)

type FileConfig interface {
	GetFileConfig() (ConfigRequired, error)
	WriteFileConfig(cfg ConfigRequired) (ConfigRequired, error)
}

type ConfigRequired struct {
	HostAddr        string `json:",omitempty"`
	IPFSStorageAddr string `json:",omitempty"`
	TimeZone        string `json:",omitempty"`
}

type Config struct {
	ConfigRequired           `json:",inline"`
	NewAccount               bool `ignored:"true"` // set to true if a new account is creating. This option controls whether mw should wait for the existing data to arrive before creating the new log
	Offline                  bool
	DisableThreadsSyncEvents bool

	RepoPath string

	PrivateNetworkSecret string

	SwarmLowWater  int
	SwarmHighWater int
	BootstrapNodes []string
	RelayNodes     []string

	CafeAddr        string
	CafeGrpcPort    int
	CafeP2PPort     int
	CafePeerId      string
	CafeAPIInsecure bool

	DebugAddr       string
	LocalServerAddr string

	Threads              threads.Config
	DS                   clientds.Config
	FS                   clientds.FSConfig
	DisableFileConfig    bool `ignored:"true"` // set in order to skip reading/writing config from/to file
	CreateBuiltinObjects bool
}

type DebugAPIConfig struct {
	commonnet.Config
	IsEnabled bool
}

const (
	ConfigFileName = "config.json"
)

var DefaultConfig = Config{
	Offline:              false,
	SwarmLowWater:        10,
	SwarmHighWater:       50,
	LocalServerAddr:      "0.0.0.0:9091",
	PrivateNetworkSecret: ipfs.IpfsPrivateNetworkKey,
	BootstrapNodes: []string{
		"/ip4/54.93.109.23/tcp/4001/p2p/QmZ4P1Q8HhtKpMshHorM2HDg4iVGZdhZ7YN7WeWDWFH3Hi",           // fra1
		"/dns4/bootstrap2.anytype.io/tcp/4001/p2p/QmSxuiczQTjgj5agSoNtp4esSsj64RisDyKt2MCZQsKZUx", // sfo1
		"/dns4/bootstrap3.anytype.io/tcp/4001/p2p/QmUdDTWzgdcf4cM4aHeihoYSUfQJJbLVLTZFZvm1b46NNT", // sgp1
	},
	RelayNodes: []string{
		"/dns4/relay2.anytype.io/tcp/4101/p2p/12D3KooWMLuW43JqNzUHbXMJH2Ted5Nf26sxv1VMcZAxXV3d3YtB",
		"/dns4/relay1.anytype.io/tcp/4101/p2p/12D3KooWNPqCu4BC5WMBuHmqdiNWwAHGTNKbNy6JP5W1DML2psg1",
	},
	CafeAPIInsecure: false,
	CafeAddr:        "cafe1.anytype.io",
	CafeP2PPort:     4001,
	CafeGrpcPort:    3006,
	CafePeerId:      "12D3KooWKwPC165PptjnzYzGrEs7NSjsF5vvMmxmuqpA2VfaBbLw",

	DS:      clientds.DefaultConfig,
	Threads: threads.DefaultConfig,
}

func WithNewAccount(isNewAccount bool) func(*Config) {
	return func(c *Config) {
		c.NewAccount = isNewAccount
	}
}

func WithStagingCafe(isStaging bool) func(*Config) {
	return func(c *Config) {
		if isStaging {
			c.CafeAddr = "cafe-staging.anytype.io"
			c.CafePeerId = "12D3KooWPGR6LQyTEtBzFnJ7fGEMe6hKiQKeNof29zLH4bGq2djR"
		}
	}
}

func WithDebugAddr(addr string) func(*Config) {
	return func(c *Config) {
		c.DebugAddr = addr
	}
}

func WithLocalServer(addr string) func(*Config) {
	return func(c *Config) {
		c.LocalServerAddr = addr
	}
}

func WithCreateBuiltinObjects(createBuiltinObjects bool) func(*Config) {
	return func(c *Config) {
		c.CreateBuiltinObjects = createBuiltinObjects
	}
}

func DisableFileConfig(disable bool) func(*Config) {
	return func(c *Config) {
		c.DisableFileConfig = disable
	}
}

func New(options ...func(*Config)) *Config {
	cfg := DefaultConfig
	if randomPort, err := netutil.GetRandomPort(); err == nil {
		cfg.LocalServerAddr = fmt.Sprintf("0.0.0.0:%d", randomPort)
	}
	for _, opt := range options {
		opt(&cfg)
	}
	cfg.Threads.CafeP2PAddr = cfg.CafeP2PFullAddr()
	cfg.Threads.CafePID = cfg.CafePeerId

	return &cfg
}

func (c *Config) CafeNodeGrpcAddr() string {
	return c.CafeAddr + ":" + strconv.Itoa(c.CafeGrpcPort)
}

func (c *Config) CafeUrl() string {
	if net.ParseIP(c.CafeAddr) != nil {
		return c.CafeAddr
	}
	prefix := "https://"
	if c.CafeAPIInsecure {
		prefix = "http://"
	}
	return prefix + c.CafeAddr
}

func (c *Config) CafeP2PFullAddr() string {
	prefix := "dns4"
	if net.ParseIP(c.CafeAddr) != nil {
		prefix = "ip4"
	}
	return fmt.Sprintf("/%s/%s/tcp/%d/p2p/%s", prefix, c.CafeAddr, c.CafeP2PPort, c.CafePeerId)
}

func (c *Config) Init(a *app.App) (err error) {
	repoPath := a.MustComponent(wallet.CName).(wallet.Wallet).RepoPath()
	if err = c.initFromFileAndEnv(repoPath); err != nil {
		return
	}

	return
}

func (c *Config) initFromFileAndEnv(repoPath string) error {
	c.RepoPath = repoPath

	if !c.DisableFileConfig {
		err := files.GetFileConfig(c.GetConfigPath(), &c.ConfigRequired)
		if err != nil {
			return fmt.Errorf("failed to get config from file: %s", err.Error())
		}

		saveRandomHostAddr := func() error {
			port, err := getRandomPort()
			if err != nil {
				port = 4006
				log.Errorf("failed to get random port for gateway, go with the default %d: %s", port, err.Error())
			}

			c.HostAddr = fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port)

			err = files.WriteJsonConfig(c.GetConfigPath(), c.ConfigRequired)
			if err != nil {
				return fmt.Errorf("failed to save port to the cfg file: %s", err.Error())
			}
			return nil
		}

		if c.HostAddr == "" {
			err = saveRandomHostAddr()
			if err != nil {
				return err
			}
		} else {
			parts := strings.Split(c.HostAddr, "/")
			if len(parts) == 0 {
				log.Errorf("failed to parse cfg.HostAddr: %s", c.HostAddr)
			} else {
				// lets test the existing port in config
				addr, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:"+parts[len(parts)-1])
				if err == nil {
					l, err := net.ListenTCP("tcp4", addr)
					if err != nil {
						// the port from config is no longer available. It may be used by other app or blocked by the OS(e.g. port exclusion range on windows)
						// lets find another available port and save it to config
						err = saveRandomHostAddr()
						if err != nil {
							return err
						}
					} else {
						_ = l.Close()
					}
				}
			}
		}

	}

	err := envconfig.Process("ANYTYPE", c)
	if err != nil {
		log.Errorf("failed to read config from env: %v", err)
	}

	return nil
}

func (c *Config) Name() (name string) {
	return CName
}

func (c *Config) DSConfig() clientds.Config {
	return c.DS
}

func (c *Config) FSConfig() (clientds.FSConfig, error) {
	res := ConfigRequired{}
	err := files.GetFileConfig(c.GetConfigPath(), &res)
	if err != nil {
		return clientds.FSConfig{}, err
	}

	return clientds.FSConfig{IPFSStorageAddr: res.IPFSStorageAddr}, nil
}

func (c *Config) ThreadsConfig() threads.Config {
	return c.Threads
}

func (c *Config) GetConfigPath() string {
	return filepath.Join(c.RepoPath, ConfigFileName)
}

func getRandomPort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}

	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

func (c *Config) GetSpace() commonspace.Config {
	return commonspace.Config{
		GCTTL:      60,
		SyncPeriod: 20,
	}
}

func (c *Config) GetMetric() metric.Config {
	return metric.Config{}
}

func (c *Config) GetNet() commonnet.Config {
	return commonnet.Config{
		Server: commonnet.ServerConfig{
			ListenAddrs: []string{c.LocalServerAddr},
		},
		Stream: commonnet.StreamConfig{
			TimeoutMilliseconds: 1000,
			MaxMsgSizeMb:        256,
		},
	}
}

func (c *Config) GetDebugAPIConfig() DebugAPIConfig {
	return DebugAPIConfig{
		Config: commonnet.Config{
			Server: commonnet.ServerConfig{
				ListenAddrs: []string{c.DebugAddr},
			},
			Stream: commonnet.StreamConfig{
				TimeoutMilliseconds: 1000,
				MaxMsgSizeMb:        256,
			},
		},
		IsEnabled: len(c.DebugAddr) != 0,
	}
}

func (c *Config) GetNodes() []nodeconf.NodeConfig {
	var nodes struct {
		Nodes []nodeconf.NodeConfig `yaml:"nodes"`
	}
	if err := yaml.Unmarshal(nodesConfYmlBytes, &nodes); err != nil {
		panic(fmt.Errorf("unable to parse node config: %v", err))
	}
	return nodes.Nodes
}
