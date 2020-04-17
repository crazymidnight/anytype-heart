package core

import (
	"context"
	"encoding/base64"
	"fmt"
	"sort"
	"time"

	"github.com/anytypeio/go-anytype-library/pb/model"
	"github.com/anytypeio/go-anytype-library/pb/storage"
	"github.com/anytypeio/go-anytype-library/vclock"
	"github.com/gogo/protobuf/types"
	cid "github.com/ipfs/go-cid"
	cbornode "github.com/ipfs/go-ipld-cbor"
	"github.com/textileio/go-threads/cbor"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/go-threads/crypto"
)

type SmartBlockSnapshot interface {
	State() vclock.VClock
	Creator() (string, error)
	CreatedDate() *time.Time
	ReceivedDate() *time.Time
	Blocks() ([]*model.Block, error)
	Meta() (*SmartBlockMeta, error)
	PublicWebURL() (string, error)
}

type smartBlockSnapshot struct {
	blocks  []*model.Block
	details *types.Struct
	state   vclock.VClock

	threadID thread.ID
	recordID cid.Cid
	eventID  cid.Cid
	key      crypto.DecryptionKey
	creator  string
	date     *types.Timestamp
	node     *Anytype
}

func (snapshot smartBlockSnapshot) State() vclock.VClock {
	return snapshot.state
}

func (snapshot smartBlockSnapshot) Creator() (string, error) {
	return snapshot.creator, nil
}

func (snapshot smartBlockSnapshot) CreatedDate() *time.Time {
	return nil
}

func (snapshot smartBlockSnapshot) ReceivedDate() *time.Time {
	return nil
}

func (snapshot smartBlockSnapshot) Blocks() ([]*model.Block, error) {
	// todo: blocks lazy loading
	return snapshot.blocks, nil
}

func (snapshot smartBlockSnapshot) Meta() (*SmartBlockMeta, error) {
	return &SmartBlockMeta{Details: snapshot.details}, nil
}

func (snapshot smartBlockSnapshot) PublicWebURL() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	ipfs := snapshot.node.Ipfs()
	if snapshot.eventID == cid.Undef {
		// todo: extract from recordID?
		return "", fmt.Errorf("eventID is empty")
	}

	event, err := cbor.GetEvent(ctx, ipfs, snapshot.eventID)
	if err != nil {
		return "", fmt.Errorf("failed to get snapshot event: %w", err)
	}

	header, err := event.GetHeader(ctx, ipfs, snapshot.key)
	if err != nil {
		return "", fmt.Errorf("failed to get snapshot event header: %w", err)
	}

	bodyKey, err := header.Key()
	if err != nil {
		return "", fmt.Errorf("failed to get body decryption key: %w", err)
	}

	bodyKeyBin, err := bodyKey.MarshalBinary()
	if err != nil {
		return "", fmt.Errorf("failed to get marshal decryption key: %w", err)
	}

	return fmt.Sprintf(
		WebGatewayHost+WebGatewaySnapshotURI,
		snapshot.threadID.String(),
		event.BodyID().String(),
		base64.RawURLEncoding.EncodeToString(bodyKeyBin),
	), nil
}

type SnapshotWithMetadata struct {
	storage.SmartBlockSnapshot
	Creator  string
	RecordID cid.Cid
	EventID  cid.Cid
}

func (a *Anytype) snapshotTraverseFromCid(ctx context.Context, thrd thread.Info, li thread.LogInfo, before vclock.VClock, limit int) ([]SnapshotWithMetadata, error) {
	var snapshots []SnapshotWithMetadata
	// todo: filter by record type
	var m = make(map[cid.Cid]struct{})

	pubKey, err := li.ID.ExtractPublicKey()
	if err != nil {
		return nil, err
	}
	rid := li.Head
	if rid == cid.Undef {
		return []SnapshotWithMetadata{}, nil
	}

	for {
		if _, exists := m[rid]; exists {
			break
		}
		m[rid] = struct{}{}
		rec, err := a.t.GetRecord(ctx, thrd.ID, rid)
		if err != nil {
			return nil, err
		}

		event, err := cbor.EventFromRecord(ctx, a.t, rec)
		if err != nil {
			return nil, err
		}

		node, err := event.GetBody(context.TODO(), a.t, thrd.Key.Read())
		if err != nil {
			return nil, fmt.Errorf("failed to get record body: %w", err)
		}
		m := new(SignedPbPayload)
		err = cbornode.DecodeInto(node.RawData(), m)
		if err != nil {
			return nil, fmt.Errorf("incorrect record type: %w", err)
		}

		err = m.Verify(pubKey)
		if err != nil {
			return nil, err
		}

		var snapshot = storage.SmartBlockSnapshot{}
		err = m.Unmarshal(&snapshot)
		if err != nil {
			return nil, fmt.Errorf("failed to decode pb block snapshot: %w", err)
		}

		if !before.IsNil() && vclock.NewFromMap(snapshot.State).Compare(before, vclock.Descendant) {
			log.Debugf("snapshotTraverseFromCid skip Descendant: %+v > %+v", snapshot.State, before)
			continue
		}

		snapshots = append(snapshots, SnapshotWithMetadata{
			SmartBlockSnapshot: snapshot,
			Creator:            m.AccAddr,
			RecordID:           rec.Cid(),
			EventID:            event.Cid(),
		})
		if len(snapshots) == limit {
			break
		}

		if !rec.PrevID().Defined() {
			break
		}

		rid = rec.PrevID()
	}

	return snapshots, nil
}

func (a *Anytype) snapshotTraverseLogs(ctx context.Context, thrdId thread.ID, before vclock.VClock, limit int) ([]SnapshotWithMetadata, error) {
	var allSnapshots []SnapshotWithMetadata
	thrd, err := a.t.GetThread(context.Background(), thrdId)
	if err != nil {
		return nil, err
	}

	for _, log := range thrd.Logs {
		snapshots, err := a.snapshotTraverseFromCid(ctx, thrd, log, before, limit)
		if err != nil {
			continue
		}

		allSnapshots = append(allSnapshots, snapshots...)
	}

	sort.Slice(allSnapshots, func(i, j int) bool {
		// sort from the newest to the oldest snapshot
		stateI := vclock.NewFromMap(allSnapshots[i].State)
		stateJ := vclock.NewFromMap(allSnapshots[j].State)
		anc := stateI.Compare(stateJ, vclock.Ancestor)
		if anc {
			return true
		}

		if stateI.Compare(stateJ, vclock.Descendant) {
			return false
		}

		// in case of concurrent changes choose the hash with greater hash first
		return stateI.Hash() > stateJ.Hash()
	})

	if len(allSnapshots) < limit {
		limit = len(allSnapshots)
	}

	return allSnapshots[0:limit], nil
}
