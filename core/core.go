package core

import (
	"io/ioutil"
	"path/filepath"
	"time"

	ipfsCore "github.com/ipfs/go-ipfs/core"
	"github.com/libp2p/go-libp2p-core/crypto"
	tcore "github.com/textileio/go-textile/core"
	tmobile "github.com/textileio/go-textile/mobile"
)

const privateKey = `/key/swarm/psk/1.0.0/
/base16/
fee6e180af8fc354d321fde5c84cab22138f9c62fec0d1bc0e99f4439968b02c`

var BootstrapNodes = []string{
	"/ip4/68.183.2.167/tcp/4001/ipfs/12D3KooWB2Ya2GkLLRSR322Z13ZDZ9LP4fDJxauscYwUMKLFCqaD",
	"/ip4/157.230.124.182/tcp/4001/ipfs/12D3KooWKLLf9Qc6SHaLWNPvx7Tk4AMc9i71CLdnbZuRiFMFMnEf",
}

type Anytype struct {
	Textile *tmobile.Mobile
}

func (a *Anytype) ipfs() *ipfsCore.IpfsNode {
	return a.Textile.Node().Ipfs()
}

func (a *Anytype) textile() *tcore.Textile {
	return a.Textile.Node()
}

func New(repoPath string, account string) (*Anytype, error) {
	// todo: remove this temp workaround after release of go-ipfs v0.4.23
	crypto.MinRsaKeyBits = 1024

	msg := messenger{}
	tm, err := tmobile.NewTextile(&tmobile.RunConfig{filepath.Join(repoPath, account), true, nil}, &msg)
	if err != nil {
		return nil, err
	}

	return &Anytype{Textile: tm}, nil
}

func (a *Anytype) Run() error {
	swarmKeyFilePath := filepath.Join(a.textile().RepoPath(), "swarm.key")
	err := ioutil.WriteFile(swarmKeyFilePath, []byte(privateKey), 0644)
	if err != nil {
		return err
	}

	err = a.Textile.Start()
	if err != nil {
		return err
	}

	err = a.ipfs().Repo.SetConfigKey("Addresses.Bootstrap", BootstrapNodes)
	if err != nil {
		return err
	}

	go func() {
		for {
			if !a.textile().Started() {
				break
			}

			if !a.ipfs().IsOnline {
				time.Sleep(time.Second)
				continue
			}

			_, err = a.textile().RegisterCafe("12D3KooWB2Ya2GkLLRSR322Z13ZDZ9LP4fDJxauscYwUMKLFCqaD", "2MsR9h7mfq53oNt8vh7RfdPr57qPsn28X3dwbviZWs3E8kEu6kpdcDHyMx7Qo")
			if err != nil {
				log.Errorf("failed to register cafe: %s", err.Error())
				time.Sleep(time.Second * 5)
				continue
			}
			break
		}
	}()

	err = a.createPredefinedThreads()
	if err != nil {
		return err
	}

	return nil
}

func (a *Anytype) Stop() error {
	return a.textile().Stop()
}
