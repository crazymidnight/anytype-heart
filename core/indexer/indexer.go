package indexer

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/anytypeio/any-sync/app"
	"github.com/gogo/protobuf/types"
	ds "github.com/ipfs/go-datastore"
	"github.com/textileio/go-threads/core/thread"
	"golang.org/x/exp/slices"

	"github.com/anytypeio/go-anytype-middleware/core/anytype/config"
	"github.com/anytypeio/go-anytype-middleware/core/block"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor"
	smartblock2 "github.com/anytypeio/go-anytype-middleware/core/block/editor/smartblock"
	"github.com/anytypeio/go-anytype-middleware/core/block/source"
	"github.com/anytypeio/go-anytype-middleware/core/relation/relationutils"
	"github.com/anytypeio/go-anytype-middleware/metrics"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/bundle"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core/smartblock"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/database"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/addr"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/filestore"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/ftsearch"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/objectstore"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/logging"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/space"
	"github.com/anytypeio/go-anytype-middleware/space/typeprovider"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/anytypeio/go-anytype-middleware/util/slice"
)

const (
	CName = "indexer"

	// ### Increasing counters below will trigger existing account to reindex their

	// ForceThreadsObjectsReindexCounter reindex thread-based objects
	ForceThreadsObjectsReindexCounter int32 = 8
	// ForceFilesReindexCounter reindex ipfs-file-based objects
	ForceFilesReindexCounter int32 = 9 //
	// ForceBundledObjectsReindexCounter reindex objects like anytypeProfile
	ForceBundledObjectsReindexCounter int32 = 5 // reindex objects like anytypeProfile
	// ForceIdxRebuildCounter erases localstore indexes and reindex all type of objects
	// (no need to increase ForceThreadsObjectsReindexCounter & ForceFilesReindexCounter)
	ForceIdxRebuildCounter int32 = 38
	// ForceFulltextIndexCounter  performs fulltext indexing for all type of objects (useful when we change fulltext config)
	ForceFulltextIndexCounter int32 = 4
	// ForceFilestoreKeysReindexCounter reindex filestore keys in all objects
	ForceFilestoreKeysReindexCounter int32 = 2
)

var log = logging.Logger("anytype-doc-indexer")

var (
	ftIndexInterval         = time.Minute
	ftIndexForceMinInterval = time.Second * 10
)

func New() Indexer {
	return &indexer{}
}

type Indexer interface {
	ForceFTIndex()
	app.ComponentRunnable
}

type ThreadLister interface {
	Threads() (thread.IDSlice, error)
}

type Hasher interface {
	Hash() string
}

type subObjectCreator interface {
	CreateSubObjectsInWorkspace(details []*types.Struct) (ids []string, objects []*types.Struct, err error)
}

type indexer struct {
	store            objectstore.ObjectStore
	fileStore        filestore.FileStore
	anytype          core.Service
	source           source.Service
	picker           block.Picker
	ftsearch         ftsearch.FTSearch
	subObjectCreator subObjectCreator

	quit        chan struct{}
	mu          sync.Mutex
	btHash      Hasher
	archivedMap map[string]struct{}
	favoriteMap map[string]struct{}
	newAccount  bool
	forceFt     chan struct{}

	typeProvider typeprovider.ObjectTypeProvider
	spaceService space.Service
}

func (i *indexer) Init(a *app.App) (err error) {
	i.newAccount = a.MustComponent(config.CName).(*config.Config).NewAccount
	i.anytype = a.MustComponent(core.CName).(core.Service)
	i.store = a.MustComponent(objectstore.CName).(objectstore.ObjectStore)
	i.typeProvider = a.MustComponent(typeprovider.CName).(typeprovider.ObjectTypeProvider)
	i.source = a.MustComponent(source.CName).(source.Service)
	i.btHash = a.MustComponent("builtintemplate").(Hasher)
	i.picker = app.MustComponent[block.Picker](a)
	i.fileStore = app.MustComponent[filestore.FileStore](a)
	i.spaceService = app.MustComponent[space.Service](a)
	i.ftsearch = app.MustComponent[ftsearch.FTSearch](a)
	i.subObjectCreator = app.MustComponent[subObjectCreator](a)
	i.quit = make(chan struct{})
	i.archivedMap = make(map[string]struct{}, 100)
	i.favoriteMap = make(map[string]struct{}, 100)
	i.forceFt = make(chan struct{})
	return
}

func (i *indexer) Name() (name string) {
	return CName
}

func (i *indexer) Run(context.Context) (err error) {
	if ftErr := i.ftInit(); ftErr != nil {
		log.Errorf("can't init ft: %v", ftErr)
	}
	err = i.reindexIfNeeded()
	if err != nil {
		return err
	}
	i.migrateRemoveNonindexableObjects()
	go i.ftLoop()
	return
}

func (i *indexer) migrateRemoveNonindexableObjects() {
	ids, err := i.getIdsForTypes(
		smartblock.SmartblockTypeMarketplaceType, smartblock.SmartblockTypeMarketplaceRelation,
		smartblock.SmartblockTypeMarketplaceTemplate, smartblock.SmartBlockTypeDate, smartblock.SmartBlockTypeBreadcrumbs,
	)
	if err != nil {
		log.Errorf("migrateRemoveNonindexableObjects: failed to get ids: %s", err.Error())
	}

	for _, id := range ids {
		err = i.store.DeleteDetails(id)
		if err != nil {
			log.Errorf("migrateRemoveNonindexableObjects: failed to get ids: %s", err.Error())
		}
	}
}

func (i *indexer) Close(ctx context.Context) (err error) {
	i.mu.Lock()
	quit := i.quit
	i.mu.Unlock()
	if quit != nil {
		close(quit)
		i.mu.Lock()
		i.quit = nil
		i.mu.Unlock()
	}
	return nil
}

func (i *indexer) Index(ctx context.Context, info smartblock2.DocInfo) error {
	startTime := time.Now()
	sbType, err := i.typeProvider.Type(info.Id)
	if err != nil {
		sbType = smartblock.SmartBlockTypePage
	}
	saveIndexedHash := func() {
		if headsHash := headsHash(info.Heads); headsHash != "" {
			err = i.store.SaveLastIndexedHeadsHash(info.Id, headsHash)
			if err != nil {
				log.With("thread", info.Id).Errorf("failed to save indexed heads hash: %v", err)
			}
		}
	}

	indexDetails, indexLinks := sbType.Indexable()
	if !indexDetails && !indexLinks {
		saveIndexedHash()
		return nil
	}

	details := info.State.CombinedDetails()
	details.Fields[bundle.RelationKeyLinks.String()] = pbtypes.StringList(info.Links)
	setCreator := pbtypes.GetString(info.State.LocalDetails(), bundle.RelationKeyCreator.String())
	if setCreator == "" {
		setCreator = i.anytype.ProfileID()
	}
	indexSetTime := time.Now()
	var hasError bool
	if indexLinks {
		if err = i.store.UpdateObjectLinks(info.Id, info.Links); err != nil {
			hasError = true
			log.With("thread", info.Id).Errorf("failed to save object links: %v", err)
		}
	}

	indexLinksTime := time.Now()
	if indexDetails {
		if err := i.store.UpdateObjectDetails(info.Id, details, false); err != nil {
			hasError = true
			log.With("thread", info.Id).Errorf("can't update object store: %v", err)
		}
		if err := i.store.AddToIndexQueue(info.Id); err != nil {
			log.With("thread", info.Id).Errorf("can't add id to index queue: %v", err)
		} else {
			log.With("thread", info.Id).Debugf("to index queue")
		}

		go i.indexLinkedFiles(ctx, info.FileHashes)
	} else {
		_ = i.store.DeleteDetails(info.Id)
	}
	indexDetailsTime := time.Now()
	detailsCount := 0
	if details.GetFields() != nil {
		detailsCount = len(details.GetFields())
	}

	if !hasError {
		saveIndexedHash()
	}

	metrics.SharedClient.RecordEvent(metrics.IndexEvent{
		ObjectId:                info.Id,
		IndexLinksTimeMs:        indexLinksTime.Sub(indexSetTime).Milliseconds(),
		IndexDetailsTimeMs:      indexDetailsTime.Sub(indexLinksTime).Milliseconds(),
		IndexSetRelationsTimeMs: indexSetTime.Sub(startTime).Milliseconds(),
		RelationsCount:          len(info.State.PickRelationLinks()),
		DetailsCount:            detailsCount,
	})

	return nil
}

func (i *indexer) indexLinkedFiles(ctx context.Context, fileHashes []string) {
	if len(fileHashes) == 0 {
		return
	}
	existingIDs, err := i.store.HasIDs(fileHashes...)
	if err != nil {
		log.Errorf("failed to get existing file ids : %s", err.Error())
	}
	newIDs := slice.Difference(fileHashes, existingIDs)
	for _, id := range newIDs {
		// file's hash is id
		err = i.reindexDoc(ctx, id)
		if err != nil {
			log.With("id", id).Errorf("failed to reindex file: %s", err.Error())
		}

		err = i.store.AddToIndexQueue(id)
		if err != nil {
			log.With("id", id).Error(err.Error())
		}
	}
}

func (i *indexer) reindexIfNeeded() error {
	checksums, err := i.store.GetChecksums()
	if err != nil && err != ds.ErrNotFound {
		return err
	}
	if checksums == nil {
		checksums = &model.ObjectStoreChecksums{
			// do no add bundled relations checksums, because we want to index them for new accounts
			ObjectsForceReindexCounter:       ForceThreadsObjectsReindexCounter,
			FilesForceReindexCounter:         ForceFilesReindexCounter,
			IdxRebuildCounter:                ForceIdxRebuildCounter,
			FilestoreKeysForceReindexCounter: ForceFilestoreKeysReindexCounter,
		}
	}

	var flags reindexFlags
	if checksums.BundledRelations != bundle.RelationChecksum {
		flags.bundledRelations = true
	}
	if checksums.BundledObjectTypes != bundle.TypeChecksum {
		flags.bundledTypes = true
	}
	if checksums.ObjectsForceReindexCounter != ForceThreadsObjectsReindexCounter {
		flags.threadObjects = true
	}
	if checksums.FilestoreKeysForceReindexCounter != ForceFilestoreKeysReindexCounter {
		flags.fileKeys = true
	}
	if checksums.FilesForceReindexCounter != ForceFilesReindexCounter {
		flags.fileObjects = true
	}
	if checksums.FulltextRebuild != ForceFulltextIndexCounter {
		flags.fulltext = true
	}
	if checksums.BundledTemplates != i.btHash.Hash() {
		flags.bundledTemplates = true
	}
	if checksums.BundledObjects != ForceBundledObjectsReindexCounter {
		flags.bundledObjects = true
	}
	if checksums.IdxRebuildCounter != ForceIdxRebuildCounter {
		flags.enableAll()
	}
	return i.reindex(context.WithValue(context.TODO(), metrics.CtxKeyRequest, "reindex_forced"), flags)
}

func (i *indexer) reindex(ctx context.Context, flags reindexFlags) (err error) {
	if flags.any() {
		log.Infof("start store reindex (%s)", flags.String())
	}

	if flags.fileKeys {
		err = i.fileStore.RemoveEmpty()
		if err != nil {
			log.Errorf("reindex failed to RemoveEmpty filekeys: %v", err.Error())
		} else {
			log.Infof("RemoveEmpty filekeys succeed")
		}
	}

	if flags.removeAllIndexedObjects {
		ids, err := i.store.ListIds()
		if err != nil {
			log.Errorf("reindex failed to get all ids(removeAllIndexedObjects): %v", err.Error())
		}
		for _, id := range ids {
			err = i.store.DeleteDetails(id)
			if err != nil {
				log.Errorf("reindex failed to delete details(removeAllIndexedObjects): %v", err.Error())
			}
		}
	}
	var indexesWereRemoved bool
	if flags.eraseIndexes {
		err = i.store.EraseIndexes()
		if err != nil {
			log.Errorf("reindex failed to erase indexes: %v", err.Error())
		} else {
			log.Infof("all store indexes succesfully erased")
			indexesWereRemoved = true
		}
	}

	// We derive or init predefined blocks here in order to ensure consistency of object store.
	// If we call this method before removing objects from store, we will end up with inconsistent state
	// because indexing of predefined objects will not run again
	err = i.anytype.EnsurePredefinedBlocks(ctx)
	if err != nil {
		return err
	}

	err = i.ensurePreinstalledObjects()
	if err != nil {
		return fmt.Errorf("ensure preinstalled objects: %w", err)
	}

	if flags.any() {
		d, err := i.getObjectInfo(ctx, i.anytype.PredefinedBlocks().Archive)
		if err != nil {
			log.Errorf("reindex failed to open archive: %s", err.Error())
		} else {
			for _, target := range d.Links {
				i.archivedMap[target] = struct{}{}
			}
		}

		d, err = i.getObjectInfo(ctx, i.anytype.PredefinedBlocks().Home)
		if err != nil {
			log.Errorf("reindex failed to open home: %s", err.Error())
		} else {
			for _, b := range d.Links {
				i.favoriteMap[b] = struct{}{}
			}
		}
	}

	// for all ids except home and archive setting cache timeout for reindexing
	// ctx = context.WithValue(ctx, ocache.CacheTimeout, cacheTimeout)
	if flags.threadObjects {
		ids, err := i.getIdsForTypes(
			smartblock.SmartBlockTypePage,
			smartblock.SmartBlockTypeSet,
			smartblock.SmartBlockTypeObjectType,
			smartblock.SmartBlockTypeProfilePage,
			smartblock.SmartBlockTypeTemplate,
			smartblock.SmartblockTypeMarketplaceType,
			smartblock.SmartblockTypeMarketplaceTemplate,
			smartblock.SmartblockTypeMarketplaceRelation,
			smartblock.SmartBlockTypeArchive,
			smartblock.SmartBlockTypeHome,
			smartblock.SmartBlockTypeWorkspaceOld,
		)
		if err != nil {
			return err
		}
		start := time.Now()
		successfullyReindexed := i.reindexIdsIgnoreErr(ctx, ids...)
		if metrics.Enabled {
			metrics.SharedClient.RecordEvent(metrics.ReindexEvent{
				ReindexType:    metrics.ReindexTypeThreads,
				Total:          len(ids),
				Success:        successfullyReindexed,
				SpentMs:        int(time.Since(start).Milliseconds()),
				IndexesRemoved: indexesWereRemoved,
			})
		}
		log.Infof("%d/%d objects have been successfully reindexed", successfullyReindexed, len(ids))
	} else {
		go func() {
			start := time.Now()
			total, success, err := i.reindexOutdatedThreads()
			if err != nil {
				log.Infof("failed to reindex outdated objects: %s", err.Error())
			} else {
				log.Infof("%d/%d outdated objects have been successfully reindexed", success, total)
			}
			if metrics.Enabled && total > 0 {
				metrics.SharedClient.RecordEvent(metrics.ReindexEvent{
					ReindexType:    metrics.ReindexTypeOutdatedHeads,
					Total:          total,
					Success:        success,
					SpentMs:        int(time.Since(start).Milliseconds()),
					IndexesRemoved: indexesWereRemoved,
				})
			}
		}()
	}

	if flags.fileObjects {
		err = i.reindexIDsForSmartblockTypes(ctx, metrics.ReindexTypeFiles, indexesWereRemoved, smartblock.SmartBlockTypeFile)
		if err != nil {
			return err
		}
	}
	if flags.bundledRelations {
		err = i.reindexIDsForSmartblockTypes(ctx, metrics.ReindexTypeBundledRelations, indexesWereRemoved, smartblock.SmartBlockTypeBundledRelation)
		if err != nil {
			return err
		}
	}
	if flags.bundledTypes {
		err = i.reindexIDsForSmartblockTypes(ctx, metrics.ReindexTypeBundledTypes, indexesWereRemoved, smartblock.SmartBlockTypeBundledObjectType, smartblock.SmartBlockTypeAnytypeProfile)
		if err != nil {
			return err
		}
	}
	if flags.bundledObjects {
		// hardcoded for now
		ids := []string{addr.AnytypeProfileId}
		err = i.reindexIDs(ctx, metrics.ReindexTypeBundledObjects, false, ids)
		if err != nil {
			return err
		}
	}

	if flags.bundledTemplates {
		existing, _, err := i.store.QueryObjectIds(database.Query{}, []smartblock.SmartBlockType{smartblock.SmartBlockTypeBundledTemplate})
		if err != nil {
			return err
		}
		for _, id := range existing {
			i.store.DeleteObject(id)
		}

		err = i.reindexIDsForSmartblockTypes(ctx, metrics.ReindexTypeBundledTemplates, indexesWereRemoved, smartblock.SmartBlockTypeBundledTemplate)
		if err != nil {
			return err
		}
	}
	if flags.fulltext {
		ids, err := i.getIdsForTypes(smartblock.SmartBlockTypePage, smartblock.SmartBlockTypeFile, smartblock.SmartBlockTypeBundledRelation, smartblock.SmartBlockTypeBundledObjectType, smartblock.SmartBlockTypeAnytypeProfile)
		if err != nil {
			return err
		}

		var addedToQueue int
		for _, id := range ids {
			if err := i.store.AddToIndexQueue(id); err != nil {
				log.Errorf("failed to add to index queue: %v", err)
			} else {
				addedToQueue++
			}
		}
		msg := fmt.Sprintf("%d/%d objects have been successfully added to the fulltext queue", addedToQueue, len(ids))
		if len(ids)-addedToQueue != 0 {
			log.Error(msg)
		} else {
			log.Info(msg)
		}
	}

	return i.saveLatestChecksums()
}

func (i *indexer) reindexIDsForSmartblockTypes(ctx context.Context, reindexType metrics.ReindexType, indexesWereRemoved bool, sbTypes ...smartblock.SmartBlockType) error {
	ids, err := i.getIdsForTypes(sbTypes...)
	if err != nil {
		return err
	}
	return i.reindexIDs(ctx, reindexType, indexesWereRemoved, ids)
}

func (i *indexer) reindexIDs(ctx context.Context, reindexType metrics.ReindexType, indexesWereRemoved bool, ids []string) error {
	start := time.Now()
	successfullyReindexed := i.reindexIdsIgnoreErr(ctx, ids...)
	if metrics.Enabled && len(ids) > 0 {
		metrics.SharedClient.RecordEvent(metrics.ReindexEvent{
			ReindexType:    reindexType,
			Total:          len(ids),
			Success:        successfullyReindexed,
			SpentMs:        int(time.Since(start).Milliseconds()),
			IndexesRemoved: indexesWereRemoved,
		})
	}
	msg := fmt.Sprintf("%d/%d %s have been successfully reindexed", successfullyReindexed, len(ids), reindexType)
	if len(ids)-successfullyReindexed != 0 {
		log.Error(msg)
	} else {
		log.Info(msg)
	}
	return nil
}

func (i *indexer) ensurePreinstalledObjects() error {
	var objects []*types.Struct

	for _, ot := range bundle.SystemTypes {
		t, err := bundle.GetTypeByUrl(ot.BundledURL())
		if err != nil {
			continue
		}
		objects = append(objects, (&relationutils.ObjectType{ObjectType: t}).ToStruct())
	}

	for _, rk := range bundle.SystemRelations {
		rel := bundle.MustGetRelation(rk)
		for _, opt := range rel.SelectDict {
			opt.RelationKey = rel.Key
			objects = append(objects, (&relationutils.Option{RelationOption: opt}).ToStruct())
		}
		objects = append(objects, (&relationutils.Relation{Relation: rel}).ToStruct())
	}

	_, _, err := i.subObjectCreator.CreateSubObjectsInWorkspace(objects)
	if errors.Is(err, editor.ErrSubObjectAlreadyExists) {
		return nil
	}
	return err
}

func (i *indexer) saveLatestChecksums() error {
	// todo: add layout indexing when needed
	checksums := model.ObjectStoreChecksums{
		BundledObjectTypes:         bundle.TypeChecksum,
		BundledRelations:           bundle.RelationChecksum,
		BundledTemplates:           i.btHash.Hash(),
		ObjectsForceReindexCounter: ForceThreadsObjectsReindexCounter,
		FilesForceReindexCounter:   ForceFilesReindexCounter,

		IdxRebuildCounter:                ForceIdxRebuildCounter,
		FulltextRebuild:                  ForceFulltextIndexCounter,
		BundledObjects:                   ForceBundledObjectsReindexCounter,
		FilestoreKeysForceReindexCounter: ForceFilestoreKeysReindexCounter,
	}
	return i.store.SaveChecksums(&checksums)
}

func (i *indexer) saveLatestCounters() error {
	// todo: add layout indexing when needed
	checksums := model.ObjectStoreChecksums{
		BundledObjectTypes:               bundle.TypeChecksum,
		BundledRelations:                 bundle.RelationChecksum,
		BundledTemplates:                 i.btHash.Hash(),
		ObjectsForceReindexCounter:       ForceThreadsObjectsReindexCounter,
		FilesForceReindexCounter:         ForceFilesReindexCounter,
		IdxRebuildCounter:                ForceIdxRebuildCounter,
		FulltextRebuild:                  ForceFulltextIndexCounter,
		BundledObjects:                   ForceBundledObjectsReindexCounter,
		FilestoreKeysForceReindexCounter: ForceFilestoreKeysReindexCounter,
	}
	return i.store.SaveChecksums(&checksums)
}

func (i *indexer) reindexOutdatedThreads() (toReindex, success int, err error) {
	spc, err := i.spaceService.AccountSpace(context.Background())
	if err != nil {
		return
	}

	tids := spc.StoredIds()
	var idsToReindex []string
	for _, tid := range tids {
		logErr := func(err error) {
			log.With("tree", tid).Errorf("reindexOutdatedThreads failed to get tree to reindex: %s", err.Error())
		}

		lastHash, err := i.store.GetLastIndexedHeadsHash(tid)
		if err != nil {
			logErr(err)
			continue
		}
		info, err := spc.Storage().TreeStorage(tid)
		if err != nil {
			logErr(err)
			continue
		}
		heads, err := info.Heads()
		if err != nil {
			logErr(err)
			continue
		}

		hh := headsHash(heads)
		if lastHash != hh {
			log.With("tree", tid).Warnf("not equal indexed heads hash: %s!=%s (%d logs)", lastHash, hh, len(heads))
			idsToReindex = append(idsToReindex, tid)
		}
	}

	ctx := context.WithValue(context.Background(), metrics.CtxKeyRequest, "reindexOutdatedThreads")
	success = i.reindexIdsIgnoreErr(ctx, idsToReindex...)
	return len(idsToReindex), success, nil
}

func (i *indexer) reindexDoc(ctx context.Context, id string) error {
	_, isArchived := i.archivedMap[id]
	_, isFavorite := i.favoriteMap[id]

	err := i.store.UpdatePendingLocalDetails(id, func(pending *types.Struct) (*types.Struct, error) {
		pending.Fields[bundle.RelationKeyIsArchived.String()] = pbtypes.Bool(isArchived)
		pending.Fields[bundle.RelationKeyIsFavorite.String()] = pbtypes.Bool(isFavorite)
		return pending, nil
	})
	if err != nil {
		log.Errorf("failed to update isArchived and isFavorite details for %s: %s", id, err)
	}

	// Touch the object to initiate indexing
	return block.DoWithContext(ctx, i.picker, id, func(sb smartblock2.SmartBlock) error {
		return sb.Apply(sb.NewState(), smartblock2.NoHistory, smartblock2.NoEvent, smartblock2.NoRestrictions)
	})
}

func (i *indexer) reindexIdsIgnoreErr(ctx context.Context, ids ...string) (successfullyReindexed int) {
	for _, id := range ids {
		err := i.reindexDoc(ctx, id)
		if err != nil {
			log.With("thread", id).Errorf("failed to reindex: %v", err)
		} else {
			successfullyReindexed++
		}
	}
	return
}

func (i *indexer) getObjectInfo(ctx context.Context, id string) (info smartblock2.DocInfo, err error) {
	err = block.DoWithContext(ctx, i.picker, id, func(sb smartblock2.SmartBlock) error {
		info = sb.GetDocInfo()
		return nil
	})
	return
}

func (i *indexer) getIdsForTypes(sbt ...smartblock.SmartBlockType) ([]string, error) {
	var ids []string
	for _, t := range sbt {
		st, err := i.source.SourceTypeBySbType(t)
		if err != nil {
			return nil, err
		}
		idsT, err := st.ListIds()
		if err != nil {
			return nil, err
		}
		ids = append(ids, idsT...)
	}
	return ids, nil
}

func headsHash(heads []string) string {
	if len(heads) == 0 {
		return ""
	}
	slices.Sort(heads)

	sum := sha256.Sum256([]byte(strings.Join(heads, ",")))
	return fmt.Sprintf("%x", sum)
}
