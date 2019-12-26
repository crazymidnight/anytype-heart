package block

import (
	"testing"

	"github.com/anytypeio/go-anytype-library/core"
	"github.com/anytypeio/go-anytype-library/pb/model"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCommonSmart_Open(t *testing.T) {
	t.Run("should send fullscreen event on open", func(t *testing.T) {
		fx := newServiceFixture(t, "")
		defer fx.ctrl.Finish()
		defer fx.tearDown()

		sb := &commonSmart{
			s:              fx.Service.(*service),
			versionsChange: func(vers []core.BlockVersion) {},
		}

		block, _ := fx.newMockBlockWithContent(
			"1",
			&model.BlockContentOfPage{Page: &model.BlockContentPage{}},
			[]string{"2", "3"},
			map[string]core.BlockVersion{
				"2": fx.newMockVersion(&model.Block{Id: "2"}),
				"3": fx.newMockVersion(&model.Block{Id: "3"}),
			},
		)
		block.EXPECT().Close()
		err := sb.Open(block)
		require.NoError(t, err)
		sb.Init()

		defer func() {
			err := sb.Close()
			require.NoError(t, err)
			assert.True(t, block.cancelBlockVersionsCalled)
			assert.True(t, block.cancelClientEventsCalled)
		}()

		require.Len(t, fx.events, 1)
		event := fx.events[0]
		require.IsType(t, (*pb.EventMessageValueOfBlockShow)(nil), event.Messages[0].Value)
		show := event.Messages[0].Value.(*pb.EventMessageValueOfBlockShow).BlockShow
		assert.Equal(t, show.RootId, "1")
		assert.Len(t, show.Blocks, 3)
	})
}

func TestCommonSmart_Create(t *testing.T) {
	t.Run("should create block", func(t *testing.T) {
		fx := newServiceFixture(t, "")
		defer fx.ctrl.Finish()
		defer fx.tearDown()

		sb := &commonSmart{
			s: fx.Service.(*service),
		}

		block, _ := fx.newMockBlockWithContent(
			"1",
			&model.BlockContentOfPage{Page: &model.BlockContentPage{}},
			[]string{"2", "3"},
			map[string]core.BlockVersion{
				"2": fx.newMockVersion(&model.Block{Id: "2"}),
				"3": fx.newMockVersion(&model.Block{Id: "3"}),
			},
		)
		block.EXPECT().Close()
		err := sb.Open(block)
		require.NoError(t, err)
		sb.Init()
		defer func() {
			err := sb.Close()
			require.NoError(t, err)
		}()

		req := pb.RpcBlockCreateRequest{
			Block: &model.Block{
				Content: &model.BlockContentOfText{Text: &model.BlockContentText{}},
			},
			TargetId:  "3",
			Position:  model.Block_Top,
			ContextId: "1",
		}
		newBlockId := "23"
		newBlock, _ := fx.newMockBlockWithContent(newBlockId, req.Block.Content, nil, nil)
		block.EXPECT().NewBlock(*req.Block).Return(newBlock, nil)
		newBlockVer, _ := newBlock.GetCurrentVersion()

		var versToSave []*model.Block
		block.EXPECT().AddVersions(&matcher{name: "AddVersions", f: func(x interface{}) bool {
			versToSave = x.([]*model.Block)
			return true
		}}).Return([]core.BlockVersion{newBlockVer}, nil)

		id, err := sb.Create(req)
		require.NoError(t, err)
		assert.Equal(t, newBlockId, id)

		require.Len(t, versToSave, 2)
		//assert.Equal(t, []string{"2", "23", "3"}, versToSave[0].ChildrenIds)
		t.Log(versToSave)
		assert.Len(t, fx.events, 2)
	})
	t.Run("create smart block", func(t *testing.T) {
		fx := newPageFixture(t)
		defer fx.ctrl.Finish()
		defer fx.tearDown()
		b := &model.Block{
			Fields: &types.Struct{
				Fields: map[string]*types.Value{
					"key": testStringValue("value"),
				},
			},
			Content: &model.BlockContentOfPage{
				Page: &model.BlockContentPage{},
			},
		}
		fx.block.EXPECT().AddVersion(b)
		newId, err := fx.Create(pb.RpcBlockCreateRequest{
			ContextId: fx.pageId,
			Block:     b,
		})
		require.NoError(t, err)

		require.Len(t, fx.savedBlocks, 2)
		assert.Equal(t, b.Id, fx.savedBlocks[newId].GetLink().TargetBlockId)
	})
}

func TestCommonSmart_Duplicate(t *testing.T) {
	t.Run("should duplicate block with child", func(t *testing.T) {
		// initial blocks on page
		pageBlocks := []*model.Block{
			{Id: "b1"},
			{Id: "b2", ChildrenIds: []string{"c1"}},
			{Id: "c1"},
		}
		fx := newPageFixture(t, pageBlocks...)
		defer fx.ctrl.Finish()
		defer fx.tearDown()

		require.Len(t, fx.versions[fx.GetId()].Model().ChildrenIds, 2)

		newId, err := fx.Duplicate(pb.RpcBlockDuplicateRequest{
			TargetId: "b1",
			BlockId:  "b2",
			Position: model.Block_Top,
		})
		require.NoError(t, err)

		// plus one block in page
		require.Len(t, fx.versions[fx.GetId()].Model().ChildrenIds, 3)
		// have new copied block as first page child
		assert.Equal(t, newId, fx.versions[fx.GetId()].Model().ChildrenIds[0])
		// copied block have children
		require.Len(t, fx.versions[newId].Model().ChildrenIds, 1)
		// copied child have new id
		assert.NotEqual(t, "c1", fx.versions[newId].Model().ChildrenIds[0])

		// have 2 events: 1 - show, 2 - update for duplicate
		require.Len(t, fx.serviceFx.events, 2)
		// check we have 3 messages: 2 add + 1 change children
		assert.Len(t, fx.serviceFx.events[1].Messages, 3)
	})

}
