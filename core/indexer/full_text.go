package indexer

import (
	"context"
	"fmt"
	"time"

	"github.com/anyproto/anytype-heart/metrics"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	"github.com/anyproto/anytype-heart/pkg/lib/core/smartblock"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/ftsearch"
	"github.com/anyproto/anytype-heart/util/pbtypes"
)

var (
	ftIndexInterval         = 10 * time.Second
	ftIndexForceMinInterval = time.Second * 10
)

func (i *indexer) ForceFTIndex() {
	select {
	case i.forceFt <- struct{}{}:
	default:
	}
}

func (i *indexer) ftLoop() {
	ticker := time.NewTicker(ftIndexInterval)
	i.runFullTextIndexer()
	var lastForceIndex time.Time
	for {
		select {
		case <-i.quit:
			return
		case <-ticker.C:
			i.runFullTextIndexer()
		case <-i.forceFt:
			if time.Since(lastForceIndex) > ftIndexForceMinInterval {
				i.runFullTextIndexer()
				lastForceIndex = time.Now()
			}
		}
	}
}

// TODO maybe use two queues? One for objects, one for files
func (i *indexer) runFullTextIndexer() {
	ids, err := i.store.ListIDsFromFullTextQueue()
	if err != nil {
		log.Errorf("list ids from full-text queue: %v", err)
		return
	}

	var docs []ftsearch.SearchDoc
	for _, id := range ids {
		doc, err := i.prepareSearchDocument(id)
		if err != nil {
			log.With("id", id).Errorf("prepare document for full-text indexing: %s", err)
			continue
		}
		docs = append(docs, doc)
	}

	err = i.ftsearch.BatchIndex(docs)
	if err != nil {
		log.Errorf("full-text indexing: %v", err)
		return
	}

	i.store.RemoveIDsFromFullTextQueue(ids)
}

func (i *indexer) prepareSearchDocument(id string) (ftDoc ftsearch.SearchDoc, err error) {
	// ctx := context.WithValue(context.Background(), ocache.CacheTimeout, cacheTimeout)
	ctx := context.WithValue(context.Background(), metrics.CtxKeyEntrypoint, "index_fulltext")
	info, err := i.getObjectInfo(ctx, id)
	if err != nil {
		return ftDoc, fmt.Errorf("get object info: %w", err)
	}
	// TODO Parametrize with actual SpaceID: GO-1625
	sbType, err := i.typeProvider.Type(i.provider.PersonalSpaceID(), info.Id)
	if err != nil {
		sbType = smartblock.SmartBlockTypePage
	}
	indexDetails, _ := sbType.Indexable()
	if !indexDetails {
		return ftsearch.SearchDoc{}, nil
	}

	if err = i.store.UpdateObjectSnippet(id, info.State.Snippet()); err != nil {
		return
	}

	title := pbtypes.GetString(info.State.Details(), bundle.RelationKeyName.String())
	if info.State.ObjectTypeKey() == bundle.TypeKeyNote || title == "" {
		title = info.State.Snippet()
	}

	spaceID := pbtypes.GetString(info.State.LocalDetails(), bundle.RelationKeySpaceId.String())
	ftDoc = ftsearch.SearchDoc{
		Id:      id,
		SpaceID: spaceID,
		Title:   title,
		Text:    info.State.SearchText(),
	}
	return
}

func (i *indexer) ftInit() error {
	if ft := i.store.FTSearch(); ft != nil {
		docCount, err := ft.DocCount()
		if err != nil {
			return err
		}
		if docCount == 0 {
			ids, err := i.store.ListIds()
			if err != nil {
				return err
			}
			for _, id := range ids {
				if err := i.store.AddToIndexQueue(id); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
