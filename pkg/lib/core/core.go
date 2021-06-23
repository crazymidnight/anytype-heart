package core

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/libp2p/go-libp2p-core/peer"
	pstore "github.com/libp2p/go-libp2p-core/peerstore"
	"github.com/libp2p/go-libp2p/p2p/discovery"

	"github.com/anytypeio/go-anytype-middleware/app"
	"github.com/anytypeio/go-anytype-middleware/core/anytype/config"
	"github.com/anytypeio/go-anytype-middleware/core/wallet"
	"github.com/anytypeio/go-anytype-middleware/metrics"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/cafe"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core/smartblock"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/datastore"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/files"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/ipfs"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/filestore"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/objectstore"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/logging"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pin"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/threads"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/util"
)

var log = logging.Logger("anytype-core")

const (
	CName  = "anytype"
	tmpDir = "tmp"
)

type PredefinedBlockIds struct {
	Account string
	Profile string
	Home    string
	Archive string

	SetPages string
}

type Service interface {
	Account() string // deprecated, use wallet component
	Device() string  // deprecated, use wallet component
	Start() error
	Stop() error
	IsStarted() bool
	BecameOnline(ch chan<- error)

	// InitNewSmartblocksChan allows to init the chan to inform when there is a new smartblock becomes available
	// Can be called only once. Returns error if called more than once
	InitNewSmartblocksChan(ch chan<- string) error
	InitPredefinedBlocks(ctx context.Context, mustSyncFromRemote bool) error
	PredefinedBlocks() threads.DerivedSmartblockIds
	GetBlock(blockId string) (SmartBlock, error)
	DeleteBlock(blockId string) error
	CreateBlock(t smartblock.SmartBlockType) (SmartBlock, error)

	FileByHash(ctx context.Context, hash string) (File, error)
	FileAdd(ctx context.Context, opts ...files.AddOption) (File, error)
	FileAddWithBytes(ctx context.Context, content []byte, filename string) (File, error)         // deprecated
	FileAddWithReader(ctx context.Context, content io.ReadSeeker, filename string) (File, error) // deprecated
	FileGetKeys(hash string) (*files.FileKeys, error)
	FileStoreKeys(fileKeys ...files.FileKeys) error

	ImageByHash(ctx context.Context, hash string) (Image, error)
	ImageAdd(ctx context.Context, opts ...files.AddOption) (Image, error)
	ImageAddWithBytes(ctx context.Context, content []byte, filename string) (Image, error)         // deprecated
	ImageAddWithReader(ctx context.Context, content io.ReadSeeker, filename string) (Image, error) // deprecated

	ObjectStore() objectstore.ObjectStore // deprecated
	FileStore() filestore.FileStore       // deprecated
	ThreadsIds() ([]string, error)        // deprecated

	ObjectInfoWithLinks(id string) (*model.ObjectInfoWithLinks, error)
	ObjectList() ([]*model.ObjectInfo, error)

	ProfileInfo

	app.ComponentRunnable
	TempDir() string
}

var _ app.Component = (*Anytype)(nil)

var _ Service = (*Anytype)(nil)

type Anytype struct {
	files       *files.Service
	cafe        cafe.Client
	mdns        discovery.Service
	objectStore objectstore.ObjectStore
	fileStore   filestore.FileStore

	ds datastore.Datastore

	predefinedBlockIds threads.DerivedSmartblockIds
	threadService      threads.Service
	pinService         pin.FilePinService
	ipfs               ipfs.Node
	logLevels          map[string]string

	opts ServiceOptions

	replicationWG    sync.WaitGroup
	migrationOnce    sync.Once
	lock             sync.Mutex
	isStarted        bool // use under the lock
	shutdownStartsCh chan struct {
	} // closed when node shutdown starts
	onlineCh chan struct {
	} // closed when became online

	recordsbatch        batchAdder
	subscribeOnce       sync.Once
	config              *config.Config
	wallet              wallet.Wallet
	tmpFolderAutocreate sync.Once
	tempDir             string
}

func (a *Anytype) ThreadsIds() ([]string, error) {
	tids, err := a.ThreadService().Logstore().Threads()
	if err != nil {
		return nil, err
	}
	return util.ThreadIdsToStings(tids), nil
}

type batchAdder interface {
	Add(msgs ...SmartblockRecordWithThreadID) error
	Close() error
}

func New() *Anytype {
	return &Anytype{
		shutdownStartsCh: make(chan struct{}),
		onlineCh:         make(chan struct{}),
	}
}

func (a *Anytype) Init(ap *app.App) (err error) {
	a.wallet = ap.MustComponent(wallet.CName).(wallet.Wallet)
	a.config = ap.MustComponent(config.CName).(*config.Config)
	a.recordsbatch = ap.MustComponent("recordsbatcher").(batchAdder)
	a.objectStore = ap.MustComponent(objectstore.CName).(objectstore.ObjectStore)
	a.fileStore = ap.MustComponent(filestore.CName).(filestore.FileStore)
	a.ds = ap.MustComponent(datastore.CName).(datastore.Datastore)
	a.threadService = ap.MustComponent(threads.CName).(threads.Service)
	a.cafe = ap.MustComponent(cafe.CName).(cafe.Client)
	a.files = ap.MustComponent(files.CName).(*files.Service)
	a.pinService = ap.MustComponent(pin.CName).(pin.FilePinService)
	a.ipfs = ap.MustComponent(ipfs.CName).(ipfs.Node)

	return
}

func (a *Anytype) Name() string {
	return CName
}

// Deprecated, use wallet component directly
func (a *Anytype) Account() string {
	pk, _ := a.wallet.GetAccountPrivkey()
	if pk == nil {
		return ""
	}
	return pk.Address()
}

// Deprecated, use wallet component directly
func (a *Anytype) Device() string {
	pk, _ := a.wallet.GetDevicePrivkey()
	if pk == nil {
		return ""
	}
	return pk.Address()
}

func (a *Anytype) Run() (err error) {
	if err = a.Start(); err != nil {
		return
	}

	metrics.SharedClient.SetUserId(a.Account())
	metrics.SharedClient.StartAggregating()

	return a.InitPredefinedBlocks(context.TODO(), a.config.NewAccount)
}

func (a *Anytype) IsStarted() bool {
	a.lock.Lock()
	defer a.lock.Unlock()

	return a.isStarted
}

func (a *Anytype) BecameOnline(ch chan<- error) {
	for {
		select {
		case <-a.onlineCh:
			ch <- nil
			close(ch)
		case <-a.shutdownStartsCh:
			ch <- fmt.Errorf("node was shutdown")
			close(ch)
		}
	}
}

func (a *Anytype) CreateBlock(t smartblock.SmartBlockType) (SmartBlock, error) {
	thrd, err := a.threadService.CreateThread(t)
	if err != nil {
		return nil, err
	}
	sb := &smartBlock{thread: thrd, node: a}
	return sb, nil
}

// PredefinedBlocks returns default blocks like home and archive
// ⚠️ Will return empty struct in case it runs before Anytype.Start()
func (a *Anytype) PredefinedBlocks() threads.DerivedSmartblockIds {
	return a.predefinedBlockIds
}

func (a *Anytype) HandlePeerFound(p peer.AddrInfo) {
	a.ThreadService().Threads().Host().Peerstore().AddAddrs(p.ID, p.Addrs, pstore.ConnectedAddrTTL)
}

func (a *Anytype) Start() error {
	err := a.RunMigrations()
	if err != nil {
		return err
	}

	return a.start()
}

func (a *Anytype) start() error {
	a.lock.Lock()
	defer a.lock.Unlock()

	if a.isStarted {
		return nil
	}

	if err := a.subscribeForNewRecords(); err != nil {
		return err
	}

	a.isStarted = true
	return nil
}

func (a *Anytype) InitPredefinedBlocks(ctx context.Context, newAccount bool) error {
	cctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		select {
		case <-cctx.Done():
			return
		case <-a.shutdownStartsCh:
			cancel()
		}
	}()

	ids, err := a.threadService.EnsurePredefinedThreads(cctx, newAccount)
	if err != nil {
		return err
	}

	a.predefinedBlockIds = ids

	return nil
}

func (a *Anytype) Close() (err error) {
	metrics.SharedClient.StopAggregating()
	return a.Stop()
}

func (a *Anytype) Stop() error {
	fmt.Printf("stopping the library...\n")
	defer fmt.Println("library has been successfully stopped")
	a.lock.Lock()
	defer a.lock.Unlock()
	a.isStarted = false

	if a.shutdownStartsCh != nil {
		close(a.shutdownStartsCh)
		a.shutdownStartsCh = nil
	}

	// fixme useless!
	a.replicationWG.Wait()

	return nil
}

func (a *Anytype) ThreadService() threads.Service {
	return a.threadService
}

func (a *Anytype) InitNewSmartblocksChan(ch chan<- string) error {
	if a.threadService == nil {
		return fmt.Errorf("thread service not ready yet")
	}

	return a.threadService.InitNewThreadsChan(ch)
}

func (a *Anytype) TempDir() string {
	// it shouldn't be a case when it is called before wallet init, but just in case lets add the check here
	if a.wallet == nil || a.wallet.RootPath() == "" {
		return os.TempDir()
	}

	var err error
	// simultaneous calls to TempDir will wait for the once func to finish, so it will be fine
	a.tmpFolderAutocreate.Do(func() {
		path := filepath.Join(a.wallet.RootPath(), tmpDir)
		err = os.MkdirAll(path, 0655)
		if err != nil {
			log.Errorf("failed to make temp dir, use the default system one: %s", err.Error())
			a.tempDir = os.TempDir()
		} else {
			a.tempDir = path
		}
	})

	return a.tempDir
}

// subscribeForNewRecords should be called only once as early as possible.
// Subscribes to new records for all threads and add them to the batcher
func (a *Anytype) subscribeForNewRecords() (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	// do not defer cancel, cancel only on shutdown
	threadsCh, err := a.threadService.PresubscribedNewRecords()
	if err != nil {
		return err
	}

	go func() {
		smartBlocksCache := make(map[string]*smartBlock)
		defer a.recordsbatch.Close()
		for {
			select {
			case val, ok := <-threadsCh:
				if !ok {
					return
				}
				var block *smartBlock
				id := val.ThreadID().String()
				if id == a.predefinedBlockIds.Account {
					continue
				}
				if block, ok = smartBlocksCache[id]; !ok {
					if block, err = a.GetSmartBlock(id); err != nil {
						log.Errorf("failed to open smartblock %s: %v", id, err)
						continue
					} else {
						smartBlocksCache[id] = block
					}
				}
				rec, err := block.decodeRecord(ctx, val.Value(), true)
				if err != nil {
					log.Errorf("failed to decode thread record: %s", err.Error())
					continue
				}
				err = a.recordsbatch.Add(SmartblockRecordWithThreadID{
					SmartblockRecordEnvelope: *rec,
					ThreadID:                 id,
				})
				if err != nil {
					log.Errorf("failed to add thread record to batcher: %s", err.Error())
					continue
				}
			case <-ctx.Done():
				return
			case <-a.shutdownStartsCh:
				cancel()
			}
		}
	}()

	return nil
}

func init() {

	/* logs */

	// apply log levels in go-threads and go-ipfs deps
	logging.ApplyLevelsFromEnv()
}
