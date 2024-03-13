package filesync

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/cheggaaa/mb/v3"
	"github.com/dgraph-io/badger/v4"
	"go.uber.org/zap"

	"github.com/anyproto/anytype-heart/core/domain"
	"github.com/anyproto/anytype-heart/util/badgerhelper"
)

const (
	keyPrefix = "/filesyncindex/"
)

var (
	errQueueIsEmpty = errors.New("queue is empty")

	sepByte = []byte("/")[0]

	uploadKeyPrefix       = []byte(keyPrefix + "queue/upload/")
	removeKeyPrefix       = []byte(keyPrefix + "queue/remove/")
	discardedKeyPrefix    = []byte(keyPrefix + "queue/discarded/")
	queueSchemaVersionKey = []byte(keyPrefix + "queue/schema_version")
)

type fileSyncStore struct {
	db *badger.DB

	inboxQueue     *mb.MB[*QueueItem]
	discardedQueue *mb.MB[*QueueItem]
}

func newFileSyncStore(db *badger.DB) (*fileSyncStore, error) {
	s := &fileSyncStore{
		db:             db,
		inboxQueue:     mb.New[*QueueItem](0),
		discardedQueue: mb.New[*QueueItem](0),
	}
	err := s.migrateQueue()
	if err != nil {
		return nil, fmt.Errorf("migrate queue: %w", err)
	}

	{
		items, err := s.listItemsByPrefix(uploadKeyPrefix)
		if err != nil {
			return nil, fmt.Errorf("get saved queue items: %w", err)
		}
		err = s.inboxQueue.Add(context.Background(), items...)
		if err != nil {
			return nil, fmt.Errorf("add to inbox queue: %w", err)
		}
	}

	{
		items, err := s.listItemsByPrefix(discardedKeyPrefix)
		if err != nil {
			return nil, fmt.Errorf("get saved discarded items: %w", err)
		}
		err = s.discardedQueue.Add(context.Background(), items...)
		if err != nil {
			return nil, fmt.Errorf("add to discarded queue: %w", err)
		}
	}

	return s, nil
}

type QueueItem struct {
	SpaceId     string
	FileId      domain.FileId
	Timestamp   int64
	AddedByUser bool
	Imported    bool
}

func (it *QueueItem) FullFileId() domain.FullFileId {
	return domain.FullFileId{
		SpaceId: it.SpaceId,
		FileId:  it.FileId,
	}
}

func (it *QueueItem) less(other *QueueItem) bool {
	return it.Timestamp < other.Timestamp
}

const queueSchemaVersion = 1

func (s *fileSyncStore) updateTxn(f func(txn *badger.Txn) error) error {
	return badgerhelper.RetryOnConflict(func() error {
		return s.db.Update(f)
	})
}

func (s *fileSyncStore) migrateQueue() error {
	return s.updateTxn(func(txn *badger.Txn) error {
		raw, err := txn.Get(queueSchemaVersionKey)
		if err != nil && err != badger.ErrKeyNotFound {
			return fmt.Errorf("get schema version: %w", err)
		}
		version, err := versionFromItem(raw)
		if err != nil {
			return fmt.Errorf("get schema version from item: %w", err)
		}

		if version < queueSchemaVersion {
			err = runMigrationFromVersion0(txn)
			if err != nil {
				return fmt.Errorf("run migration from version 0: %w", err)
			}
		}

		return txn.Set(queueSchemaVersionKey, []byte(strconv.Itoa(queueSchemaVersion)))
	})
}

func runMigrationFromVersion0(txn *badger.Txn) error {
	for _, prefix := range [][]byte{
		uploadKeyPrefix,
		discardedKeyPrefix,
		removeKeyPrefix,
	} {
		err := migrateByPrefix(txn, prefix)
		if err != nil {
			return fmt.Errorf("migrate by prefix %s: %w", string(prefix), err)
		}
	}
	return nil
}

func migrateByPrefix(txn *badger.Txn, prefix []byte) error {
	it := txn.NewIterator(badger.IteratorOptions{
		PrefetchSize:   100,
		PrefetchValues: true,
		Prefix:         prefix,
	})
	defer it.Close()

	for it.Rewind(); it.Valid(); it.Next() {
		item := it.Item()
		err := migrateItem(txn, item)
		if err != nil {
			return fmt.Errorf("migrate item %s: %w", item.Key(), err)
		}
		log.Warn("migrate queue item", zap.String("key", string(item.Key())))
	}
	return nil
}

func migrateItem(txn *badger.Txn, item *badger.Item) error {
	timestamp, err := getTimestamp(item)
	if err != nil {
		return fmt.Errorf("get timestamp: %w", err)
	}
	it := QueueItem{
		Timestamp: int64(timestamp),
	}
	raw, err := json.Marshal(it)
	if err != nil {
		return fmt.Errorf("marshal queue item: %w", err)
	}
	return txn.Set(item.Key(), raw)
}

func versionFromItem(it *badger.Item) (int, error) {
	if it == nil {
		return 0, nil
	}
	var res int
	err := it.Value(func(raw []byte) error {
		res, _ = strconv.Atoi(string(raw))
		return nil
	})
	return res, err
}

func (s *fileSyncStore) QueueUpload(fileId domain.FullFileId, addedByUser bool, isImported bool) (err error) {
	err = s.inboxQueue.Add(context.Background(), &QueueItem{
		SpaceId:     fileId.SpaceId,
		FileId:      fileId.FileId,
		AddedByUser: addedByUser,
		Imported:    isImported,
	})
	if err != nil {
		return fmt.Errorf("add to inbox queue: %w", err)
	}

	return s.storeQueueItem(fileId, addedByUser, isImported)
}

func (s *fileSyncStore) storeQueueItem(fileId domain.FullFileId, addedByUser bool, isImported bool) error {
	return s.updateTxn(func(txn *badger.Txn) error {
		logger := log.With(zap.String("spaceId", fileId.SpaceId), zap.String("fileId", fileId.FileId.String()), zap.Bool("addedByUser", addedByUser))
		ok, err := isKeyExists(txn, discardedKey(fileId))
		if err != nil {
			return fmt.Errorf("check discarded key: %w", err)
		}
		if ok {
			logger.Info("add file to upload queue: file is in discarded queue")
			return nil
		}
		ok, err = isKeyExists(txn, uploadKey(fileId))
		if err != nil {
			return fmt.Errorf("check upload key: %w", err)
		}
		if ok {
			logger.Info("add file to upload queue: file is already in queue, update timestamp")
		} else {
			logger.Info("add file to upload queue")
		}
		raw, err := createQueueItem(addedByUser, isImported)
		if err != nil {
			return fmt.Errorf("create queue item: %w", err)
		}
		return txn.Set(uploadKey(fileId), raw)
	})
}

func createQueueItem(addedByUser bool, imported bool) ([]byte, error) {
	return json.Marshal(QueueItem{
		Timestamp:   time.Now().UnixMilli(),
		AddedByUser: addedByUser,
		Imported:    imported,
	})
}

func (s *fileSyncStore) QueueDiscarded(fileId domain.FullFileId) (err error) {
	err = s.discardedQueue.Add(context.Background(), &QueueItem{
		SpaceId:     fileId.SpaceId,
		FileId:      fileId.FileId,
		AddedByUser: false,
		Imported:    false,
	})
	if err != nil {
		return fmt.Errorf("add to discarded queue: %w", err)
	}
	return s.updateTxn(func(txn *badger.Txn) error {
		if err = txn.Delete(uploadKey(fileId)); err != nil {
			return err
		}
		raw, err := createQueueItem(false, false)
		if err != nil {
			return fmt.Errorf("create queue item: %w", err)
		}
		return txn.Set(discardedKey(fileId), raw)
	})
}

func (s *fileSyncStore) QueueRemove(fileId domain.FullFileId) (err error) {
	return s.updateTxn(func(txn *badger.Txn) error {
		if err = removeFromUploadingQueue(txn, fileId); err != nil {
			return err
		}
		raw, err := createQueueItem(false, false)
		if err != nil {
			return fmt.Errorf("create queue item: %w", err)
		}
		return txn.Set(removeKey(fileId), raw)
	})
}

func (s *fileSyncStore) DoneUpload(fileId domain.FullFileId) (err error) {
	return s.updateTxn(func(txn *badger.Txn) error {
		if err = removeFromUploadingQueue(txn, fileId); err != nil {
			return err
		}
		return txn.Set(doneUploadKey(fileId), binTime(time.Now().UnixMilli()))
	})
}

func removeFromUploadingQueue(txn *badger.Txn, fileId domain.FullFileId) error {
	if err := txn.Delete(uploadKey(fileId)); err != nil {
		return fmt.Errorf("remove from uploading queue: %w", err)
	}
	if err := txn.Delete(discardedKey(fileId)); err != nil {
		return fmt.Errorf("remove from discarded uploading queue: %w", err)
	}
	return nil
}

func (s *fileSyncStore) DoneRemove(fileId domain.FullFileId) (err error) {
	return s.updateTxn(func(txn *badger.Txn) error {
		if err = txn.Delete(removeKey(fileId)); err != nil {
			return err
		}
		if err = txn.Delete(doneUploadKey(fileId)); err != nil {
			return err
		}
		return txn.Set(doneRemoveKey(fileId), binTime(time.Now().UnixMilli()))
	})
}

func (s *fileSyncStore) GetUpload(ctx context.Context) (it *QueueItem, err error) {
	return s.inboxQueue.WaitOne(ctx)
}

func (s *fileSyncStore) GetDiscardedUpload(ctx context.Context) (it *QueueItem, err error) {
	return s.discardedQueue.WaitOne(ctx)
}

func isKeyExists(txn *badger.Txn, key []byte) (bool, error) {
	_, err := txn.Get(key)
	if err == badger.ErrKeyNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *fileSyncStore) HasUpload(fileId domain.FullFileId) (ok bool, err error) {
	err = s.db.View(func(txn *badger.Txn) error {
		ok, err = isKeyExists(txn, uploadKey(fileId))
		return err
	})
	return
}

func (s *fileSyncStore) IsFileUploadLimited(fileId domain.FullFileId) (ok bool, err error) {
	err = s.db.View(func(txn *badger.Txn) error {
		ok, err = isKeyExists(txn, discardedKey(fileId))
		return err
	})
	return
}

func (s *fileSyncStore) GetRemove() (it *QueueItem, err error) {
	// TODO FIX
	// return s.getOne(removeKeyPrefix)
	return nil, fmt.Errorf("fix")
}

func (s *fileSyncStore) listItemsByPrefix(prefix []byte) ([]*QueueItem, error) {
	var items []*QueueItem
	err := s.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.IteratorOptions{
			PrefetchSize:   100,
			PrefetchValues: true,
			Prefix:         prefix,
		})
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			qItem, err := getQueueItem(item)
			if err != nil {
				return fmt.Errorf("get queue item %s: %w", item.Key(), err)
			}
			fileId, spaceId := extractFileAndSpaceID(item)
			qItem.FileId = domain.FileId(fileId)
			qItem.SpaceId = spaceId
			items = append(items, qItem)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return items, nil
}

func extractFileAndSpaceID(item *badger.Item) (string, string) {
	k := item.Key()
	idx := bytes.LastIndexByte(k, sepByte)
	fileId := string(k[idx+1:])
	k = k[:idx]
	idx = bytes.LastIndexByte(k, sepByte)
	spaceId := string(k[idx+1:])
	return fileId, spaceId
}

func (s *fileSyncStore) QueueLen() (l int, err error) {
	err = s.db.View(func(txn *badger.Txn) error {
		for _, prefix := range [][]byte{uploadKeyPrefix, removeKeyPrefix} {
			it := txn.NewIterator(badger.IteratorOptions{
				PrefetchSize:   100,
				PrefetchValues: false,
				Prefix:         prefix,
			})
			for it.Rewind(); it.Valid(); it.Next() {
				l++
			}
			it.Close()
		}
		return nil
	})
	return
}

func (s *fileSyncStore) IsAlreadyUploaded(fileId domain.FullFileId) (done bool, err error) {
	err = s.db.View(func(txn *badger.Txn) error {
		_, e := txn.Get(doneUploadKey(fileId))
		if e != nil && e != badger.ErrKeyNotFound {
			return e
		}
		if e != badger.ErrKeyNotFound {
			done = true
		}
		return nil
	})
	return
}

func (s *fileSyncStore) setNodeUsage(usage NodeUsage) error {
	data, err := json.Marshal(usage)
	if err != nil {
		return err
	}
	return badgerhelper.SetValue(s.db, nodeUsageKey(), data)
}

func (s *fileSyncStore) getNodeUsage() (NodeUsage, error) {
	return badgerhelper.GetValue(s.db, nodeUsageKey(), func(raw []byte) (NodeUsage, error) {
		var usage NodeUsage
		err := json.Unmarshal(raw, &usage)
		return usage, err
	})
}

func nodeUsageKey() []byte {
	return []byte(keyPrefix + "node_usage/")
}

func uploadKey(fileId domain.FullFileId) (key []byte) {
	return []byte(keyPrefix + "queue/upload/" + fileId.SpaceId + "/" + fileId.FileId.String())
}

func discardedKey(fileId domain.FullFileId) (key []byte) {
	return []byte(keyPrefix + "queue/discarded/" + fileId.SpaceId + "/" + fileId.FileId.String())
}

func removeKey(fileId domain.FullFileId) (key []byte) {
	return []byte(keyPrefix + "queue/remove/" + fileId.SpaceId + "/" + fileId.FileId.String())
}

func doneUploadKey(fileId domain.FullFileId) (key []byte) {
	return []byte(keyPrefix + "done/upload/" + fileId.SpaceId + "/" + fileId.FileId.String())
}

func doneRemoveKey(fileId domain.FullFileId) (key []byte) {
	return []byte(keyPrefix + "done/remove/" + fileId.SpaceId + "/" + fileId.FileId.String())
}

func binTime(timestamp int64) []byte {
	return binary.LittleEndian.AppendUint64(nil, uint64(timestamp))
}

func getTimestamp(item *badger.Item) (uint64, error) {
	var ts uint64
	err := item.Value(func(raw []byte) error {
		ts = binary.LittleEndian.Uint64(raw)
		return nil
	})
	return ts, err
}

func getQueueItem(item *badger.Item) (*QueueItem, error) {
	var it QueueItem
	err := item.Value(func(raw []byte) error {
		return json.Unmarshal(raw, &it)
	})
	if err != nil {
		return nil, err
	}
	return &it, nil
}
