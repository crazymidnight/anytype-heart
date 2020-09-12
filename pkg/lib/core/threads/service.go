package threads

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/anytypeio/go-anytype-library/core/smartblock"
	"github.com/anytypeio/go-anytype-library/logging"
	net2 "github.com/anytypeio/go-anytype-library/net"
	util2 "github.com/anytypeio/go-anytype-library/util"
	"github.com/anytypeio/go-anytype-library/wallet"
	ma "github.com/multiformats/go-multiaddr"
	db2 "github.com/textileio/go-threads/core/db"
	"github.com/textileio/go-threads/core/net"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/go-threads/crypto/symmetric"
	"github.com/textileio/go-threads/db"
	"github.com/textileio/go-threads/util"
)

var log = logging.Logger("anytype-threads")

type service struct {
	t                 net2.NetBoostrapper
	db                *db.DB
	threadsCollection *db.Collection
	threadsGetter     ThreadsGetter
	device            wallet.Keypair
	account           wallet.Keypair
	repoRootPath      string
	newHeadProcessor  func(id thread.ID) error
	replicatorAddr    ma.Multiaddr
	closeCh           chan struct{}
	sync.Mutex
}

func New(threadsAPI net2.NetBoostrapper, threadsGetter ThreadsGetter, repoRootPath string, deviceKeypair wallet.Keypair, accountKeypair wallet.Keypair, newHeadProcessor func(id thread.ID) error, replicatorAddr ma.Multiaddr) Service {
	return &service{
		t:                threadsAPI,
		threadsGetter:    threadsGetter,
		device:           deviceKeypair,
		repoRootPath:     repoRootPath,
		account:          accountKeypair,
		newHeadProcessor: newHeadProcessor,
		replicatorAddr:   replicatorAddr,
		closeCh:          make(chan struct{}),
	}
}

type Service interface {
	ThreadsCollection() (*db.Collection, error)
	CreateThread(blockType smartblock.SmartBlockType) (thread.Info, error)
	DeleteThread(id string) error

	EnsurePredefinedThreads(ctx context.Context, newAccount bool) (DerivedSmartblockIds, error)
	Close() error
}

type ThreadsGetter interface {
	Threads() (thread.IDSlice, error)
}

func (s *service) ThreadsCollection() (*db.Collection, error) {
	if s.threadsCollection == nil {
		return nil, fmt.Errorf("thread collection not initialized: need to call EnsurePredefinedThreads first")
	}

	return s.threadsCollection, nil
}

func (s *service) Close() error {
	close(s.closeCh)
	return nil
}

func (s *service) CreateThread(blockType smartblock.SmartBlockType) (thread.Info, error) {
	if s.threadsCollection == nil {
		return thread.Info{}, fmt.Errorf("thread collection not initialized: need to call EnsurePredefinedThreads first")
	}

	thrdId, err := threadCreateID(thread.AccessControlled, blockType)
	if err != nil {
		return thread.Info{}, err
	}
	followKey, err := symmetric.NewRandom()
	if err != nil {
		return thread.Info{}, err
	}

	readKey, err := symmetric.NewRandom()
	if err != nil {
		return thread.Info{}, err
	}

	thrd, err := s.t.CreateThread(context.TODO(), thrdId, net.WithThreadKey(thread.NewKey(followKey, readKey)), net.WithLogKey(s.device))
	if err != nil {
		return thread.Info{}, err
	}

	hasReplAddress := false
	var replAddrWithThread ma.Multiaddr
	if s.replicatorAddr != nil {
		replAddrWithThread, err = util2.MultiAddressAddThread(s.replicatorAddr, thrdId)
		if err != nil {
			return thread.Info{}, err
		}
	}

	var multiAddrs []ma.Multiaddr
	hasReplAddress = util2.MultiAddressHasReplicator(thrd.Addrs, s.replicatorAddr)

	if !hasReplAddress && replAddrWithThread != nil {
		multiAddrs = append(multiAddrs, replAddrWithThread)
	}

	threadInfo := threadInfo{
		ID:    db2.InstanceID(thrd.ID.String()),
		Key:   thrd.Key.String(),
		Addrs: util2.MultiAddressesToStrings(multiAddrs),
	}

	// todo: wait for threadsCollection to push?
	_, err = s.threadsCollection.Create(util.JSONFromInstance(threadInfo))
	if err != nil {
		log.With("thread", thrd.ID.String()).Errorf("failed to create thread at collection: %s: ", err.Error())
	}

	if replAddrWithThread != nil {
		go func() {
			attempt := 0
			// todo: rewrite to job queue in badger
			for {
				attempt++
				p, err := s.t.AddReplicator(context.TODO(), thrd.ID, replAddrWithThread)
				if err != nil {
					log.Errorf("failed to add log replicator after %d attempt: %s", attempt, err.Error())
					select {
					case <-time.After(time.Second * 3 * time.Duration(attempt)):
					case <-s.closeCh:
						return
					}
					continue
				}

				log.With("thread", thrd.ID.String()).Infof("added log replicator after %d attempt: %s", attempt, p.String())
				return
			}
		}()
	}

	return thrd, nil
}

func (s *service) DeleteThread(id string) error {
	if s.threadsCollection == nil {
		return fmt.Errorf("thread collection not initialized: need to call EnsurePredefinedThreads first")
	}

	tid, err := thread.Decode(id)
	if err != nil {
		return fmt.Errorf("incorrect block id: %w", err)
	}

	err = s.t.DeleteThread(context.Background(), tid)
	if err != nil {
		return err
	}

	err = s.threadsCollection.Delete(db2.InstanceID(id))
	if err != nil {
		// todo: here we can get an error if we didn't yet added thead keys into DB
		log.With("thread", id).Error("DeleteThread failed to remove thread from collection: %s", err.Error())
	}
	return nil
}
