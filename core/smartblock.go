package core

import (
	"context"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"

	"github.com/anytypeio/go-anytype-library/pb/model"
	"github.com/anytypeio/go-anytype-library/pb/storage"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"github.com/ipfs/go-cid"
	cbornode "github.com/ipfs/go-ipld-cbor"
	mh "github.com/multiformats/go-multihash"
	"github.com/textileio/go-threads/cbor"
	"github.com/textileio/go-threads/core/service"
	"github.com/textileio/go-threads/core/thread"
)

type SmartBlockType uint64

const (
	SmartBlockTypePage      SmartBlockType = 0x10
	SmartBlockTypeDashboard SmartBlockType = 0x20
)

type SmartBlockState map[string]uint64
func (state SmartBlockState) Hash() string{
	var sum [32]byte
	for k, v := range state {
		sum = sha256.Sum256(append(sum[:], []byte(k+"="+strconv.Itoa(int(v)))...))
	}
	return fmt.Sprintf("%x", sum)
}

func (state SmartBlockState) VectorCounterPerPeer() map[string]uint64{
	return state
}

func (state SmartBlockState) ShouldCreateSnapshot() bool{
	// some deterministic alg depends on the hash
	return true
}

type SmartBlockContentChanges struct {
	// to be discussed
}

type SmartBlockMeta struct {
	Details *types.Struct
}

type SmartBlockMetaChanges struct {
	SmartBlockMeta
	state SmartBlockState
}

func (meta *SmartBlockMetaChanges) State() SmartBlockState{
	return meta.state
}

type SmartBlockChanges interface {
	Content() SmartBlockContentChanges
	Meta() SmartBlockMetaChanges
}

type SmartBlock interface {
	ID() string
	Type() SmartBlockType
	Creator() (string, error)
	GetLastSnapshot() (SmartBlockSnapshot, error)
	GetSnapshotBefore(state SmartBlockState) (SmartBlockSnapshot, error)

	GetCurrentState() (SmartBlockState, error)
	PushChanges(content *SmartBlockContentChanges, meta *SmartBlockMetaChanges) (state SmartBlockState, err error)
	PushSnapshot(state SmartBlockState, meta *SmartBlockMeta, blocks []*model.Block) (SmartBlockSnapshot, error)
	SubscribeForChangesSinceState(state SmartBlockState, ch chan SmartBlockChanges) (cancel func(), err error)
	SubscribeForMetaChangesSinceState(state SmartBlockState, ch chan SmartBlockMetaChanges) (cancel func(), err error)

	// SubscribeClientEvents provide a way to subscribe for the client-side events e.g. carriage position change
	SubscribeClientEvents(event chan<- proto.Message) (cancelFunc func(), err error)
	// PublishClientEvent gives a way to push the new client-side event e.g. carriage position change
	// notice that you will also get this event in SubscribeForEvents
	PublishClientEvent(event proto.Message) error
}

type SmartBlockSnapshot interface {
	State() SmartBlockState
	Creator() (string, error)
	CreatedDate() *time.Time
	ReceivedDate() *time.Time
	Blocks() ([]*model.Block, error)
	Meta() (*SmartBlockMeta, error)
}

type smartBlock struct {
	thread thread.Info
	node   *Anytype
}

func (block *smartBlock) Creator() (string, error) {
	// todo: to be implemented
	return "", fmt.Errorf("to be implemented")
}

func (block *smartBlock) GetCurrentState() (SmartBlockState, error) {
	// todo: to be implemented
	return nil, fmt.Errorf("to be implemented")
}

func (block *smartBlock) PushChanges(content *SmartBlockContentChanges, meta *SmartBlockMetaChanges) (state SmartBlockState, err error) {
	// todo: to be implemented
	return nil, fmt.Errorf("to be implemented")
}

func (block *smartBlock) GetThread() thread.Info {
	return block.thread
}

func (block *smartBlock) Type() SmartBlockType {
	id := block.thread.ID.KeyString()
	v := block.thread.ID.Variant()
	fmt.Println(v)
	// skip version
	_, n := uvarint(id)
	// skip variant
	_, n2 := uvarint(id[n:])
	blockType, _ := uvarint(id[n+n2:])

	return SmartBlockType(blockType)
}

func (block *smartBlock) ID() string {
	return block.thread.ID.String()
}

func (block *smartBlock) GetLastSnapshot() (SmartBlockSnapshot, error) {
	versions, err := block.GetSnapshots("", 1, false)
	if err != nil {
		return nil, err
	}
	if len(versions) == 0 {
		return nil, fmt.Errorf("no block versions found")
	}

	return versions[0], nil
}

func (block *smartBlock) GetSnapshotBefore(state SmartBlockState) (SmartBlockSnapshot, error) {
	return nil, fmt.Errorf("not implemented")
}

func (block *smartBlock) getSnapshotTime(event service.Event) (*types.Timestamp, error) {
	header, err := event.GetHeader(context.TODO(), block.node.ts, block.thread.ReadKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get headers: %w", err)
	}

	versionTime, err := header.Time()
	if err != nil {
		return nil, fmt.Errorf("failed to get record time from headers: %w", err)
	}

	versionTimePB, err := types.TimestampProto(*versionTime)
	if err != nil {
		return nil, err
	}

	return versionTimePB, nil
}

func (block *smartBlock) getSnapshotSnapshotEvent(id string) (service.Event, error) {
	vid, err := cid.Parse(id)
	if err != nil {
		return nil, err
	}

	rec, err := block.node.ts.GetRecord(context.TODO(), block.thread.ID, vid)
	if err != nil {
		return nil, err
	}

	if block.thread.ReadKey == nil {
		return nil, fmt.Errorf("no read key")
	}
	event, err := cbor.EventFromRecord(context.TODO(), block.node.ts, rec)
	if err != nil {
		return nil, fmt.Errorf("failed to get event: %w", err)

	}

	return event, nil
}

/*func (block *smartBlock) GetSnapshotMeta(id string) (Sm, error) {
	event, err := block.getSnapshotSnapshotEvent(id)
	if err != nil {
		return nil, err
	}

	node, err := event.GetBody(context.TODO(), block.node.ts, block.thread.ReadKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get record body: %w", err)
	}
	m := new(threadSnapshot)
	err = cbornode.DecodeInto(node.RawData(), m)
	if err != nil {
		return nil, fmt.Errorf("incorrect record type: %w", err)
	}

	model, err := m.()
	if err != nil {
		return nil, fmt.Errorf("failed to decode pb block version: %w", err)
	}

	time, err := block.getSnapshotTime(event)
	if err != nil {
		return nil, fmt.Errorf("failed to decode pb block version: %w", err)
	}

	// todo: how to get creator peer id?
	version := &smartBlockSnapshotMeta{model: model, date: time, user: "<todo>"}

	return version, nil
}*/

func (block *smartBlock) GetSnapshots(offset string, limit int, metaOnly bool) (snapshots []smartBlockSnapshot, err error) {
	var head cid.Cid

	var offsetTime *time.Time
	if offset != "" {
		head, err = cid.Decode(offset)
		if err != nil {
			return nil, err
		}
		rec, err2 := block.node.ts.GetRecord(context.TODO(), block.thread.ID, head)
		if err2 != nil {
			err = err2
			return nil, err
		}
		event, err2 := cbor.EventFromRecord(context.TODO(), block.node.ts, rec)
		if err2 != nil {
			err = err2
			return
		}

		header, err2 := event.GetHeader(context.TODO(), block.node.ts, block.thread.ReadKey)
		if err2 != nil {
			err = err2
			return
		}

		offsetTime, err = header.Time()
		if err != nil {
			return
		}
	}

	records, err := block.node.traverseLogs(context.TODO(), block.thread, offsetTime, limit)
	if err != nil {
		return
	}

	for _, rec := range records {
		event, err := cbor.EventFromRecord(context.TODO(), block.node.ts, rec)
		if err != nil {
			return nil, fmt.Errorf("failed to get event: %w", err)
		}

		node, err := event.GetBody(context.TODO(), block.node.ts, block.thread.ReadKey)
		if err != nil {
			return nil, fmt.Errorf("failed to get record body: %w", err)
		}
		m := new(threadSnapshot)
		err = cbornode.DecodeInto(node.RawData(), m)
		if err != nil {
			return nil, fmt.Errorf("incorrect record type: %w", err)
		}

		model, err := m.BlockWithMeta()
		if err != nil {
			return nil, fmt.Errorf("failed to decode pb block version: %w", err)
		}

		t, err := types.TimestampProto(rec.Date)
		if err != nil {
			return nil, fmt.Errorf("can't convert tme to pb: %w", err)
		}

		snapshots = append(snapshots, smartBlockSnapshot{model: model, date: t, user: "<todo>"})
	}

	return
}

func (block *smartBlock) PushSnapshot(state SmartBlockState, meta *SmartBlockMeta, blocks []*model.Block) (SmartBlockSnapshot, error) {
	model := &storage.SmartBlockWithMeta{}
	if meta != nil && meta.Details != nil{
		model.Details = meta.Details
	}

	if blocks != nil {
		model.Blocks = blocks
	}

	var err error
	_, user, date, err := block.pushSnapshot(model)
	if err != nil {
		return nil, err
	}


	return &smartBlockSnapshot{model: model, user: user, date: date, node: block.node}, nil
}

func (block *smartBlock) pushSnapshot(newSnapshot *storage.SmartBlockWithMeta) (versionId string, user string, date *types.Timestamp, err error) {
	var newSnapshotB []byte
	newSnapshotB, err = proto.Marshal(newSnapshot)
	if err != nil {
		return
	}

	body, err2 := cbornode.WrapObject(&threadSnapshot{Data: newSnapshotB}, mh.SHA2_256, -1)
	if err2 != nil {
		err = err2
		return
	}

	rec, err2 := block.node.ts.CreateRecord(context.TODO(), block.thread.ID, body)
	if err2 != nil {
		err = err2
		return
	}

	event, err2 := cbor.EventFromRecord(context.TODO(), block.node.ts, rec.Value())
	if err2 != nil {
		err = err2
		return
	}

	header, err2 := event.GetHeader(context.TODO(), block.node.ts, block.thread.ReadKey)
	if err2 != nil {
		err = err2
		return
	}

	msgTime, err2 := header.Time()
	if err2 != nil {
		err = err2
		return
	}

	versionId = rec.LogID().String()
	log.Debugf("SmartBlock.addSnapshot: blockId = %s newSnapshotId = %s", block.ID(), versionId)
	user = block.node.account.Address()
	date, err = types.TimestampProto(*msgTime)
	if err != nil {
		return
	}

	return
}


func (block *smartBlock) EmptySnapshot() SmartBlockSnapshot {
	return &smartBlockSnapshot{
		node: block.node,
		model: &storage.SmartBlockWithMeta{
			Blocks: []*model.Block{
				// todo: add title and icon blocks
			},
		},
	}
}

func (block *smartBlock) SubscribeForChangesSinceState(state SmartBlockState, ch chan SmartBlockChanges) (cancel func(), err error) {
	chCloseFn := func() { close(ch) }

	//todo: to be implemented
	return chCloseFn, nil
}

func (block *smartBlock) SubscribeForMetaChangesSinceState(state SmartBlockState, ch chan SmartBlockMetaChanges) (cancel func(), err error) {
	chCloseFn := func() { close(ch) }

	/*// temporary just sent the last version
	if sinceSnapshotId == "" {
		// it must be set to ensure no versions were skipped in between
		return nil, fmt.Errorf("sinceSnapshotId must be set")
	}
	var closeChan = make(chan struct{})
	chCloseFn := func() {
		close(closeChan)
	}

	// todo: implement with chan from textile events feed
	if includeSinceSnapshot {
		versionMeta, err := block.GetSnapshotMeta(sinceSnapshotId)
		if err != nil {
			return chCloseFn, err
		}
		go func() {
			select {
			case blockMeta <- versionMeta:
			case <-closeChan:
			}
			close(blockMeta)
		}()
	}
*/
	//todo: to be implemented
	return chCloseFn, nil
}

func (block *smartBlock) SubscribeClientEvents(events chan<- proto.Message) (cancelFunc func(), err error) {
	//todo: to be implemented
	return func() { close(events) }, nil
}

func (block *smartBlock) PublishClientEvent(event proto.Message) error {
	//todo: to be implemented
	return fmt.Errorf("not implemented")
}

// Snapshot of varint function that work with a string rather than
// []byte to avoid unnecessary allocation

// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license as given at https://golang.org/LICENSE

// uvarint decodes a uint64 from buf and returns that value and the
// number of characters read (> 0). If an error occurred, the value is 0
// and the number of bytes n is <= 0 meaning:
//
// 	n == 0: buf too small
// 	n  < 0: value larger than 64 bits (overflow)
// 	        and -n is the number of bytes read
//
func uvarint(buf string) (uint64, int) {
	var x uint64
	var s uint
	// we have a binary string so we can't use a range loope
	for i := 0; i < len(buf); i++ {
		b := buf[i]
		if b < 0x80 {
			if i > 9 || i == 9 && b > 1 {
				return 0, -(i + 1) // overflow
			}
			return x | uint64(b)<<s, i + 1
		}
		x |= uint64(b&0x7f) << s
		s += 7
	}
	return 0, 0
}
