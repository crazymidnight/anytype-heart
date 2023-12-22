package objecttype

import (
	"strings"
	"time"

	"github.com/gogo/protobuf/types"

	"github.com/anyproto/anytype-heart/core/block/editor/smartblock"
	"github.com/anyproto/anytype-heart/core/domain"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/addr"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/objectstore"
	"github.com/anyproto/anytype-heart/pkg/lib/logging"
	"github.com/anyproto/anytype-heart/util/pbtypes"
)

const maxInstallationTime = 5 * time.Minute

var log = logging.Logger("update-last-used-date")

func UpdateLastUsedDate(spc smartblock.Space, store objectstore.ObjectStore, key domain.TypeKey) {
	uk, err := domain.UnmarshalUniqueKey(key.URL())
	if err != nil {
		log.Errorf("failed to unmarshall type key '%s': %w", key.String(), err)
		return
	}
	details, err := store.GetObjectByUniqueKey(spc.Id(), uk)
	if err != nil {
		log.Errorf("failed to get details of type object '%s': %w", key.String(), err)
		return
	}
	id := pbtypes.GetString(details.Details, bundle.RelationKeyId.String())
	if id == "" {
		log.Errorf("failed to get id from details of type object '%s': %w", key.String(), err)
		return
	}
	if err = spc.Do(id, func(sb smartblock.SmartBlock) error {
		st := sb.NewState()
		st.SetLocalDetail(bundle.RelationKeyLastUsedDate.String(), pbtypes.Int64(time.Now().Unix()))
		return sb.Apply(st)
	}); err != nil {
		log.Errorf("failed to set lastUsedDate to type object '%s': %w", key.String(), err)
	}
}

func SetLastUsedDateForInitialObjectType(id string, details *types.Struct) {
	if !strings.HasPrefix(id, addr.BundledObjectTypeURLPrefix) || details == nil || details.Fields == nil {
		return
	}

	var decrement time.Duration
	switch id {
	case bundle.TypeKeyNote.BundledURL():
		decrement = -1 * maxInstallationTime
	case bundle.TypeKeyPage.BundledURL():
		decrement = -2 * maxInstallationTime
	case bundle.TypeKeyTask.BundledURL():
		decrement = -3 * maxInstallationTime
	case bundle.TypeKeySet.BundledURL():
		decrement = -4 * maxInstallationTime
	case bundle.TypeKeyCollection.BundledURL():
		decrement = -5 * maxInstallationTime
	default:
		decrement = -7 * maxInstallationTime
	}

	// we do this trick to order crucial Anytype object types by last date
	lastUsed := time.Now().Add(decrement).Unix()
	details.Fields[bundle.RelationKeyLastUsedDate.String()] = pbtypes.Int64(lastUsed)
}
