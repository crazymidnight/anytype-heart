package typeprovider

import (
	"context"
	"errors"
	"strings"
	"sync"

	"github.com/anytypeio/any-sync/app"
	"github.com/anytypeio/any-sync/commonspace/object/tree/treechangeproto"
	"github.com/globalsign/mgo/bson"
	"github.com/gogo/protobuf/proto"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
	"go.uber.org/zap"

	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core/smartblock"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/addr"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/logging"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/space"
)

const CName = "space.typeprovider"

var log = logging.Logger(CName)

var ErrUnknownSmartBlockType = errors.New("error unknown smartblock type")

type SmartBlockTypeProvider interface {
	app.Component
	Type(id string) (smartblock.SmartBlockType, error)
}

type provider struct {
	sync.Mutex
	spaceService space.Service
	cache        map[string]smartblock.SmartBlockType
}

func New(spaceService space.Service) SmartBlockTypeProvider {
	return &provider{
		spaceService: spaceService,
	}
}

func (p *provider) Init(a *app.App) (err error) {
	p.cache = map[string]smartblock.SmartBlockType{}
	return
}

func (p *provider) Name() (name string) {
	return CName
}

func (p *provider) Type(id string) (tp smartblock.SmartBlockType, err error) {
	tp, err = smartBlockTypeFromID(id)
	if err != nil || tp != smartblock.SmartBlockTypePage {
		return
	}
	return p.objectTypeFromSpace(id)
}

func smartBlockTypeFromID(id string) (smartblock.SmartBlockType, error) {
	if strings.HasPrefix(id, addr.BundledRelationURLPrefix) {
		return smartblock.SmartBlockTypeBundledRelation, nil
	}

	if strings.HasPrefix(id, addr.BundledObjectTypeURLPrefix) {
		return smartblock.SmartBlockTypeBundledObjectType, nil
	}

	if len(strings.Split(id, addr.SubObjectCollectionIdSeparator)) == 2 {
		return smartblock.SmartBlockTypeSubObject, nil
	}

	// workaround for options that have no prefix
	// todo: remove this after migration to the new records format
	if bson.IsObjectIdHex(id) {
		return smartblock.SmartBlockTypeSubObject, nil
	}

	if strings.HasPrefix(id, addr.AnytypeProfileId) {
		return smartblock.SmartBlockTypeProfilePage, nil
	}
	if strings.HasPrefix(id, addr.VirtualPrefix) {
		sbt, err := addr.ExtractVirtualSourceType(id)
		if err != nil {
			return 0, err
		}
		return smartblock.SmartBlockType(sbt), nil
	}
	if strings.HasPrefix(id, addr.DatePrefix) {
		return smartblock.SmartBlockTypeDate, nil
	}

	c, err := cid.Decode(id)
	if err != nil {
		return smartblock.SmartBlockTypePage, err
	}
	// TODO: discard this fragile condition as soon as we will move to the multiaddr with prefix
	if c.Prefix().Codec == cid.DagProtobuf && c.Prefix().MhType == multihash.SHA2_256 {
		return smartblock.SmartBlockTypeFile, nil
	}
	if c.Prefix().Codec == cid.DagCBOR {
		return smartblock.SmartBlockTypePage, nil
	}

	return smartblock.SmartBlockTypePage, smartblock.ErrNoSuchSmartblock
}

func (p *provider) objectTypeFromSpace(id string) (tp smartblock.SmartBlockType, err error) {
	p.Lock()
	tp, exists := p.cache[id]
	if exists {
		p.Unlock()
		return
	}
	p.Unlock()

	sp, err := p.spaceService.AccountSpace(context.Background())
	if err != nil {
		return
	}

	store := sp.Storage()
	rawRoot, err := store.TreeRoot(id)
	if err != nil {
		return
	}

	root, err := p.unmarshallRoot(rawRoot)
	if err != nil {
		return
	}

	tp, err = p.objectType(root.ChangeType)
	if err != nil {
		return
	}
	p.Lock()
	defer p.Unlock()
	p.cache[id] = tp
	return
}

func (p *provider) objectType(changeType string) (smartblock.SmartBlockType, error) {
	log.With(zap.String("changeType", changeType)).Warn("getting change type")
	if v, exists := model.SmartBlockType_value[changeType]; exists {
		return smartblock.SmartBlockType(v), nil
	}

	return smartblock.SmartBlockTypePage, nil
}

func (p *provider) unmarshallRoot(rawRoot *treechangeproto.RawTreeChangeWithId) (root *treechangeproto.RootChange, err error) {
	raw := &treechangeproto.RawTreeChange{}
	err = proto.Unmarshal(rawRoot.GetRawChange(), raw)
	if err != nil {
		return
	}

	root = &treechangeproto.RootChange{}
	err = proto.Unmarshal(raw.Payload, root)
	if err != nil {
		return
	}
	return
}
