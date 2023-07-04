package storage

import (
	"context"
	"errors"

	"github.com/anyproto/any-sync/commonspace/object/acl/liststorage"
	"github.com/anyproto/any-sync/consensus/consensusproto"
	"github.com/dgraph-io/badger/v3"
)

var ErrIncorrectKey = errors.New("key format is incorrect")

type listStorage struct {
	db   *badger.DB
	keys aclKeys
	id   string
	root *consensusproto.RawRecordWithId
}

func newListStorage(spaceId string, db *badger.DB, txn *badger.Txn) (ls liststorage.ListStorage, err error) {
	keys := newAclKeys(spaceId)
	rootId, err := getTxn(txn, keys.RootIdKey())
	if err != nil {
		return
	}

	stringId := string(rootId)
	value, err := getTxn(txn, keys.RawRecordKey(stringId))
	if err != nil {
		return
	}

	rootWithId := &consensusproto.RawRecordWithId{
		Payload: value,
		Id:      stringId,
	}

	ls = &listStorage{
		db:   db,
		keys: keys,
		id:   stringId,
		root: rootWithId,
	}
	return
}

func createListStorage(spaceId string, db *badger.DB, txn *badger.Txn, root *consensusproto.RawRecordWithId) (ls liststorage.ListStorage, err error) {
	keys := newAclKeys(spaceId)
	_, err = getTxn(txn, keys.RootIdKey())
	if err != badger.ErrKeyNotFound {
		if err == nil {
			return newListStorage(spaceId, db, txn)
		}
		return
	}

	err = txn.Set(keys.HeadIdKey(), []byte(root.Id))
	if err != nil {
		return
	}

	err = txn.Set(keys.RawRecordKey(root.Id), root.Payload)
	if err != nil {
		return
	}
	err = txn.Set(keys.RootIdKey(), []byte(root.Id))
	if err != nil {
		return
	}

	ls = &listStorage{
		db:   db,
		keys: keys,
		id:   root.Id,
		root: root,
	}
	return
}

func (l *listStorage) Id() string {
	return l.id
}

func (l *listStorage) Root() (*consensusproto.RawRecordWithId, error) {
	return l.root, nil
}

func (l *listStorage) Head() (head string, err error) {
	bytes, err := getDB(l.db, l.keys.HeadIdKey())
	if err != nil {
		return
	}
	head = string(bytes)
	return
}

func (l *listStorage) GetRawRecord(ctx context.Context, id string) (raw *consensusproto.RawRecordWithId, err error) {
	res, err := getDB(l.db, l.keys.RawRecordKey(id))
	if err != nil {
		if err == badger.ErrKeyNotFound {
			err = liststorage.ErrUnknownRecord
		}
		return
	}

	raw = &consensusproto.RawRecordWithId{
		Payload: res,
		Id:      id,
	}
	return
}

func (l *listStorage) SetHead(headId string) (err error) {
	return putDB(l.db, l.keys.HeadIdKey(), []byte(headId))
}

func (l *listStorage) AddRawRecord(ctx context.Context, rec *consensusproto.RawRecordWithId) error {
	return putDB(l.db, l.keys.RawRecordKey(rec.Id), rec.Payload)
}
