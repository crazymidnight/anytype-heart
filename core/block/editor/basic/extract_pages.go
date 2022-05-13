package basic

import (
	"fmt"

	"github.com/anytypeio/go-anytype-middleware/core/block/editor/smartblock"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple/base"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/bundle"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/globalsign/mgo/bson"
	"github.com/gogo/protobuf/types"
)

type PageCreator interface {
	CreatePageFromState(ctx *state.Context, contextBlock smartblock.SmartBlock, groupId string, req pb.RpcBlockCreatePageRequest, state *state.State) (linkId string, pageId string, err error)
}

// ExtractBlocksToPages extracts child blocks from the page to separate pages and
// replaces these blocks to the links to these pages
func (bs *basic) ExtractBlocksToPages(s PageCreator, req pb.RpcBlockListConvertChildrenToPagesRequest) (linkIds []string, err error) {
	st := bs.NewState()

	roots := listRoots(st, req.BlockIds)
	for _, root := range roots {
		children := listChildren(st, root)
		newRoot, newBlocks := reassignSubtreeIds(root.Model().Id, append([]simple.Block{root}, children...))

		// Remove children
		for _, b := range children {
			st.Unlink(b.Model().Id)
		}

		// Build a state for the new page from child blocks
		pageState := state.NewDoc("", nil).NewState()
		for _, b := range newBlocks {
			pageState.Add(b)
		}
		pageState.Add(base.NewBase(&model.Block{
			// This id will be replaced by id of the new page
			Id:          "_root",
			ChildrenIds: []string{newRoot},
		}))

		fields := map[string]*types.Value{
			"name": pbtypes.String(root.Model().GetText().Text),
		}
		if req.ObjectType != "" {
			fields[bundle.RelationKeyType.String()] = pbtypes.String(req.ObjectType)
		}
		_, pageId, err := s.CreatePageFromState(nil, bs, "", pb.RpcBlockCreatePageRequest{
			ContextId: req.ContextId,
			Details: &types.Struct{
				Fields: fields,
			},
		}, pageState)
		if err != nil {
			return nil, fmt.Errorf("create child page: %w", err)
		}

		linkId, err := CreateBlock(st, "", pb.RpcBlockCreateRequest{
			TargetId: root.Model().Id,
			Block: &model.Block{
				Content: &model.BlockContentOfLink{
					Link: &model.BlockContentLink{
						TargetBlockId: pageId,
						Style:         model.BlockContentLink_Page,
					},
				},
			},
			Position: model.Block_Replace,
		})
		if err != nil {
			return nil, fmt.Errorf("create link to page %s: %w", pageId, err)
		}

		linkIds = append(linkIds, linkId)
	}

	return linkIds, bs.Apply(st)
}

// listRoots returns unique root blocks that are listed in blockIds
func listRoots(st *state.State, blockIds []string) []simple.Block {
	visited := map[string]struct{}{}

	queue := blockIds
	for len(queue) > 0 {
		id := queue[0]
		queue = queue[1:]

		b := st.Pick(id)
		if b == nil {
			continue
		}

		childrenIds := b.Model().ChildrenIds
		for _, chId := range childrenIds {
			visited[chId] = struct{}{}
			queue = append(queue, childrenIds...)
		}
	}

	var roots []simple.Block
	for _, id := range blockIds {
		if _, ok := visited[id]; ok {
			continue
		}
		b := st.Pick(id)
		if b == nil {
			continue
		}

		roots = append(roots, b)
	}
	return roots
}

func listChildren(st *state.State, root simple.Block) []simple.Block {
	var (
		queue    = []simple.Block{root}
		children []simple.Block
	)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, id := range cur.Model().ChildrenIds {
			b := st.Pick(id)
			if b == nil {
				continue
			}
			children = append(children, b)
			queue = append(queue, b)
		}
	}

	return children
}

// reassignSubtreeIds makes a copy of a subtree of blocks and assign a new id for each block
func reassignSubtreeIds(rootId string, blocks []simple.Block) (string, []simple.Block) {
	res := make([]simple.Block, 0, len(blocks))
	mapping := map[string]string{}
	for _, b := range blocks {
		newId := bson.NewObjectId().Hex()
		mapping[b.Model().Id] = newId

		newBlock := b.Copy()
		newBlock.Model().Id = newId
		res = append(res, newBlock)
	}

	for _, b := range res {
		for i, id := range b.Model().ChildrenIds {
			b.Model().ChildrenIds[i] = mapping[id]
		}
	}
	return mapping[rootId], res
}
