package block

import (
	"crypto/rand"
	"fmt"

	"github.com/anyproto/any-sync/commonspace/object/tree/objecttree"
	"github.com/anyproto/any-sync/util/crypto"

	"github.com/anyproto/anytype-heart/core/domain"
	coresb "github.com/anyproto/anytype-heart/pkg/lib/core/smartblock"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	spaceservice "github.com/anyproto/anytype-heart/space"
)

type internalKeyGetter interface {
	InternalKey() string
}

func createChangePayload(sbType coresb.SmartBlockType, key domain.UniqueKey) (data []byte, err error) {
	var keyStr string
	if key != nil {
		keyInternalGetter, ok := key.(internalKeyGetter)
		if !ok {
			return nil, fmt.Errorf("uniquekey format invalid")
		}
		if key.SmartblockType() != sbType.ToProto() {
			return nil, fmt.Errorf("uniquekey smartblocktype mismatch")
		}
		keyStr = keyInternalGetter.InternalKey()
	}

	payload := &model.ObjectChangePayload{SmartBlockType: model.SmartBlockType(sbType), Key: keyStr}
	return payload.Marshal()
}

func derivePayload(spaceId string, signKey crypto.PrivKey, changePayload []byte) objecttree.ObjectTreeCreatePayload {
	return objecttree.ObjectTreeCreatePayload{
		PrivKey:       signKey,
		ChangeType:    spaceservice.ChangeType,
		ChangePayload: changePayload,
		SpaceId:       spaceId,
		IsEncrypted:   true,
	}
}

func createPayload(spaceId string, signKey crypto.PrivKey, changePayload []byte, timestamp int64) (objecttree.ObjectTreeCreatePayload, error) {
	seed := make([]byte, 32)
	if _, err := rand.Read(seed); err != nil {
		return objecttree.ObjectTreeCreatePayload{}, err
	}
	return objecttree.ObjectTreeCreatePayload{
		PrivKey:       signKey,
		ChangeType:    spaceservice.ChangeType,
		ChangePayload: changePayload,
		SpaceId:       spaceId,
		IsEncrypted:   true,
		Timestamp:     timestamp,
		Seed:          seed,
	}, nil
}
