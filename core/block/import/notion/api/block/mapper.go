package block

import (
	"sync"

	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
)

// MapRequest is a data object with all necessary structures for blocks
type MapRequest struct {
	Blocks []interface{}
	// Need all these maps for correct mapping of pages and databases from notion to anytype
	// for such blocks as mentions or links to objects
	NotionPageIdsToAnytype     map[string]string
	NotionDatabaseIdsToAnytype map[string]string
	PageNameToID               map[string]string
	DatabaseNameToID           map[string]string
	relMap                     sync.RWMutex
	RelationsIdsToAnytypeID    map[string]*model.SmartBlockSnapshotBase
	relOptMap                  sync.RWMutex
	RelationsIdsToOptions      map[string][]*model.SmartBlockSnapshotBase
}

func (m *MapRequest) ReadRelationsMap(key string) *model.SmartBlockSnapshotBase {
	m.relMap.RLock()
	defer m.relMap.RUnlock()
	if snapshot, ok := m.RelationsIdsToAnytypeID[key]; ok {
		return snapshot
	}
	return nil
}

func (m *MapRequest) WriteToRelationsMap(key string, relation *model.SmartBlockSnapshotBase) {
	m.relMap.Lock()
	defer m.relMap.Unlock()
	m.RelationsIdsToAnytypeID[key] = relation
}

func (m *MapRequest) ReadRelationsOptionsMap(key string) []*model.SmartBlockSnapshotBase {
	m.relOptMap.RLock()
	defer m.relOptMap.RUnlock()
	if snapshot, ok := m.RelationsIdsToOptions[key]; ok {
		return snapshot
	}
	return nil
}

func (m *MapRequest) WriteToRelationsOptionsMap(key string, relationOptions []*model.SmartBlockSnapshotBase) {
	m.relOptMap.Lock()
	defer m.relOptMap.Unlock()
	m.RelationsIdsToOptions[key] = append(m.RelationsIdsToOptions[key], relationOptions...)
}

type MapResponse struct {
	Blocks   []*model.Block
	BlockIDs []string
}

func (m *MapResponse) Merge(mergedResp *MapResponse) {
	if mergedResp != nil {
		m.BlockIDs = append(m.BlockIDs, mergedResp.BlockIDs...)
		m.Blocks = append(m.Blocks, mergedResp.Blocks...)
	}
}

func MapBlocks(req *MapRequest) *MapResponse {
	resp := &MapResponse{}
	for _, bl := range req.Blocks {
		if ba, ok := bl.(Getter); ok {
			textResp := ba.GetBlocks(req)
			resp.Merge(textResp)
			continue
		}
	}
	return resp
}
