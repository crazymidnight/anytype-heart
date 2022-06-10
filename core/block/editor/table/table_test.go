package table

import (
	"strconv"
	"testing"

	"github.com/anytypeio/go-anytype-middleware/core/block/editor/smartblock/smarttest"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTable_TableCreate(t *testing.T) {
	sb := smarttest.New("root")
	sb.AddBlock(simple.New(&model.Block{
		Id: "root",
	}))

	tb := New(sb)

	id, err := tb.TableCreate(nil, pb.RpcBlockTableCreateRequest{
		ContextId: "",
		TargetId:  "root",
		Position:  model.Block_Inner,
		Columns:   3,
		Rows:      2,
	})

	s := sb.NewState()

	assert.NoError(t, err)
	assert.True(t, s.Exists(id))

	want := mkTestTable([]string{"col1", "col2", "col3"}, []string{"row1", "row2"}, [][]string{})

	assertIsomorphic(t, want, s, map[string]string{}, map[string]string{})
}

func TestTable_FillRows(t *testing.T) {
	ctx := newTableTestContext(t, 2, 2, false,
		mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2"}, [][]string{}))

	tb, err := newTableBlockFromState(ctx.s, ctx.id)
	require.NoError(t, err)

	err = ctx.editor.RowListFill(nil, pb.RpcBlockTableRowListFillRequest{
		ContextId: "",
		BlockIds:  tb.rows().ChildrenIds,
	})

	require.NoError(t, err)

	want := mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2"}, [][]string{{"row1-col1", "row1-col2"}, {"row2-col1", "row2-col2"}})

	assertIsomorphic(t, want, ctx.s, ctx.wantMapping, ctx.gotMapping)
}

func TestTable_TableRowCreate(t *testing.T) {
	ctx := newTableTestContext(t, 2, 2, true,
		mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2"}, [][]string{{"row1-col1", "row1-col2"}, {"row2-col1", "row2-col2"}}))

	t.Run("to the top of the target", func(t *testing.T) {
		tb, err := newTableBlockFromState(ctx.s, ctx.id)
		require.NoError(t, err)

		err = ctx.editor.RowCreate(nil, pb.RpcBlockTableRowCreateRequest{
			TargetId: tb.rows().ChildrenIds[0],
			Position: model.Block_Top,
		})

		require.NoError(t, err)

		// Cells are not created automatically
		want := mkTestTable([]string{"col1", "col2"}, []string{"row3", "row1", "row2"}, [][]string{{"row1-col1", "row1-col2"}, {"row2-col1", "row2-col2"}})

		assertIsomorphic(t, want, ctx.s, ctx.wantMapping, ctx.gotMapping)
	})

	t.Run("to the bottom of the target", func(t *testing.T) {
		tb, err := newTableBlockFromState(ctx.s, ctx.id)
		require.NoError(t, err)

		err = ctx.editor.RowCreate(nil, pb.RpcBlockTableRowCreateRequest{
			TargetId: tb.rows().ChildrenIds[0],
			Position: model.Block_Bottom,
		})

		require.NoError(t, err)

		// Cells are not created automatically
		want := mkTestTable([]string{"col1", "col2"}, []string{"row3", "row4", "row1", "row2"}, [][]string{{"row1-col1", "row1-col2"}, {"row2-col1", "row2-col2"}})

		assertIsomorphic(t, want, ctx.s, ctx.wantMapping, ctx.gotMapping)
	})
}

func TestTable_TableRowDelete(t *testing.T) {
	ctx := newTableTestContext(t, 2, 2, true,
		mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2"}, [][]string{{"row1-col1", "row1-col2"}, {"row2-col1", "row2-col2"}}))

	tb, err := newTableBlockFromState(ctx.s, ctx.id)
	require.NoError(t, err)

	err = ctx.editor.RowDelete(nil, pb.RpcBlockTableRowDeleteRequest{
		TargetId: tb.rows().ChildrenIds[1],
	})

	require.NoError(t, err)

	want := mkTestTable([]string{"col1", "col2"}, []string{"row1"}, [][]string{{"row1-col1", "row1-col2"}})

	assertIsomorphic(t, want, ctx.s, ctx.wantMapping, ctx.gotMapping)
}

func TestTable_TableRowMove(t *testing.T) {
	t.Run("to the top of the target", func(t *testing.T) {
		ctx := newTableTestContext(t, 2, 3, true,
			mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2", "row3"}, [][]string{{"row1-col1", "row1-col2"}, {"row2-col1", "row2-col2"}, {"row3-col1", "row3-col2"}}))

		tb, err := newTableBlockFromState(ctx.s, ctx.id)
		require.NoError(t, err)

		err = ctx.editor.RowMove(nil, pb.RpcBlockTableRowMoveRequest{
			TargetId:     tb.rows().ChildrenIds[0],
			DropTargetId: tb.rows().ChildrenIds[2],
			Position:     model.Block_Top,
		})

		require.NoError(t, err)

		want := mkTestTable([]string{"col1", "col2"}, []string{"row2", "row1", "row3"}, [][]string{{"row2-col1", "row2-col2"}, {"row1-col1", "row1-col2"}, {"row3-col1", "row3-col2"}})

		assertIsomorphic(t, want, ctx.s, ctx.wantMapping, ctx.gotMapping)
	})

	t.Run("to the bottom of the target", func(t *testing.T) {
		ctx := newTableTestContext(t, 2, 3, true,
			mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2", "row3"}, [][]string{{"row1-col1", "row1-col2"}, {"row2-col1", "row2-col2"}, {"row3-col1", "row3-col2"}}))

		tb, err := newTableBlockFromState(ctx.s, ctx.id)
		require.NoError(t, err)

		err = ctx.editor.RowMove(nil, pb.RpcBlockTableRowMoveRequest{
			TargetId:     tb.rows().ChildrenIds[2],
			DropTargetId: tb.rows().ChildrenIds[0],
			Position:     model.Block_Bottom,
		})

		require.NoError(t, err)

		want := mkTestTable([]string{"col1", "col2"}, []string{"row1", "row3", "row2"}, [][]string{{"row1-col1", "row1-col2"}, {"row3-col1", "row3-col2"}, {"row2-col1", "row2-col2"}})

		assertIsomorphic(t, want, ctx.s, ctx.wantMapping, ctx.gotMapping)
	})
}

func TestTable_TableColumnCreate(t *testing.T) {
	ctx := newTableTestContext(t, 2, 2, true,
		mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2"}, [][]string{{"row1-col1", "row1-col2"}, {"row2-col1", "row2-col2"}}))

	t.Run("to the right of target", func(t *testing.T) {
		tb, err := newTableBlockFromState(ctx.s, ctx.id)
		require.NoError(t, err)

		target := tb.columns().ChildrenIds[0]
		err = ctx.editor.ColumnCreate(nil, pb.RpcBlockTableColumnCreateRequest{
			TargetId: target,
			Position: model.Block_Right,
		})

		require.NoError(t, err)

		want := mkTestTable([]string{"col1", "col3", "col2"}, []string{"row1", "row2"}, [][]string{{"row1-col1", "row1-col3", "row1-col2"}, {"row2-col1", "row2-col3", "row2-col2"}})

		assertIsomorphic(t, want, ctx.s, ctx.wantMapping, ctx.gotMapping)
	})

	t.Run("to the left of target", func(t *testing.T) {
		tb, err := newTableBlockFromState(ctx.s, ctx.id)
		require.NoError(t, err)

		target := tb.columns().ChildrenIds[0]
		err = ctx.editor.ColumnCreate(nil, pb.RpcBlockTableColumnCreateRequest{
			TargetId: target,
			Position: model.Block_Left,
		})

		require.NoError(t, err)

		// Remember that we operate under the same table, so previous modifications preserved
		want := mkTestTable([]string{"col4", "col1", "col3", "col2"}, []string{"row1", "row2"}, [][]string{{"row1-col4", "row1-col1", "row1-col3", "row1-col2"}, {"row2-col4", "row2-col1", "row2-col3", "row2-col2"}})

		assertIsomorphic(t, want, ctx.s, ctx.wantMapping, ctx.gotMapping)
	})
}

func TestTable_TableColumnDuplicate(t *testing.T) {
	t.Run("to the right of the target", func(t *testing.T) {
		ctx := newTableTestContext(t, 2, 2, true,
			mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2"}, [][]string{{"row1-col1", "row1-col2"}, {"row2-col1", "row2-col2"}}))

		tb, err := newTableBlockFromState(ctx.s, ctx.id)
		require.NoError(t, err)

		id, err := ctx.editor.ColumnDuplicate(nil, pb.RpcBlockTableColumnDuplicateRequest{
			BlockId:  tb.columns().ChildrenIds[0],
			TargetId: tb.columns().ChildrenIds[0],
			Position: model.Block_Right,
		})

		require.NoError(t, err)
		require.True(t, ctx.s.Exists(id))

		want := mkTestTable([]string{"col1", "col3", "col2"}, []string{"row1", "row2"}, [][]string{{"row1-col1", "row1-col3", "row1-col2"}, {"row2-col1", "row2-col3", "row2-col2"}})

		assertIsomorphic(t, want, ctx.s, ctx.wantMapping, ctx.gotMapping)
	})

	t.Run("to the left of the target", func(t *testing.T) {
		ctx := newTableTestContext(t, 2, 2, true,
			mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2"}, [][]string{{"row1-col1", "row1-col2"}, {"row2-col1", "row2-col2"}}))

		tb, err := newTableBlockFromState(ctx.s, ctx.id)
		require.NoError(t, err)

		id, err := ctx.editor.ColumnDuplicate(nil, pb.RpcBlockTableColumnDuplicateRequest{
			BlockId:  tb.columns().ChildrenIds[1],
			TargetId: tb.columns().ChildrenIds[0],
			Position: model.Block_Left,
		})

		require.NoError(t, err)
		require.True(t, ctx.s.Exists(id))

		want := mkTestTable([]string{"col3", "col1", "col2"}, []string{"row1", "row2"}, [][]string{{"row1-col3", "row1-col1", "row1-col2"}, {"row2-col3", "row2-col1", "row2-col2"}})

		assertIsomorphic(t, want, ctx.s, ctx.wantMapping, ctx.gotMapping)
	})
}

func TestTable_TableColumnMove(t *testing.T) {
	t.Run("to the right of the drop target", func(t *testing.T) {
		ctx := newTableTestContext(t, 3, 2, true,
			mkTestTable([]string{"col1", "col2", "col3"}, []string{"row1", "row2"}, [][]string{{"row1-col1", "row1-col2", "row1-col3"}, {"row2-col1", "row2-col2", "row2-col3"}}))

		tb, err := newTableBlockFromState(ctx.s, ctx.id)
		require.NoError(t, err)

		target := tb.columns().ChildrenIds[2]
		err = ctx.editor.ColumnMove(nil, pb.RpcBlockTableColumnMoveRequest{
			TargetId:     target,
			DropTargetId: tb.columns().ChildrenIds[0],
			Position:     model.Block_Right,
		})

		require.NoError(t, err)

		want := mkTestTable([]string{"col1", "col3", "col2"}, []string{"row1", "row2"}, [][]string{{"row1-col1", "row1-col3", "row1-col2"}, {"row2-col1", "row2-col3", "row2-col2"}})

		assertIsomorphic(t, want, ctx.s, ctx.wantMapping, ctx.gotMapping)
	})

	t.Run("to the left of the drop target", func(t *testing.T) {
		ctx := newTableTestContext(t, 3, 2, true,
			mkTestTable([]string{"col1", "col2", "col3"}, []string{"row1", "row2"}, [][]string{{"row1-col1", "row1-col2", "row1-col3"}, {"row2-col1", "row2-col2", "row2-col3"}}))

		tb, err := newTableBlockFromState(ctx.s, ctx.id)
		require.NoError(t, err)

		err = ctx.editor.ColumnMove(nil, pb.RpcBlockTableColumnMoveRequest{
			TargetId:     tb.columns().ChildrenIds[2],
			DropTargetId: tb.columns().ChildrenIds[0],
			Position:     model.Block_Left,
		})

		require.NoError(t, err)

		want := mkTestTable([]string{"col3", "col1", "col2"}, []string{"row1", "row2"}, [][]string{{"row1-col3", "row1-col1", "row1-col2"}, {"row2-col3", "row2-col1", "row2-col2"}})

		assertIsomorphic(t, want, ctx.s, ctx.wantMapping, ctx.gotMapping)
	})
}

func TestTable_TableColumnDelete(t *testing.T) {
	ctx := newTableTestContext(t, 3, 2, true,
		mkTestTable([]string{"col1", "col2", "col3"}, []string{"row1", "row2"}, [][]string{{"row1-col1", "row1-col2", "row1-col3"}, {"row2-col1", "row2-col2", "row2-col3"}}))

	tb, err := newTableBlockFromState(ctx.s, ctx.id)
	require.NoError(t, err)

	err = ctx.editor.ColumnDelete(nil, pb.RpcBlockTableColumnDeleteRequest{
		TargetId: tb.columns().ChildrenIds[0],
	})
	require.NoError(t, err)

	want := mkTestTable([]string{"col2", "col3"}, []string{"row1", "row2"}, [][]string{{"row1-col2", "row1-col3"}, {"row2-col2", "row2-col3"}})

	assertIsomorphic(t, want, ctx.s, ctx.wantMapping, ctx.gotMapping)
}

func TestTable_TableExpand(t *testing.T) {
	t.Run("columns only", func(t *testing.T) {
		ctx := newTableTestContext(t, 2, 2, true,
			mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2"}, [][]string{{"row1-col1", "row1-col2"}, {"row2-col1", "row2-col2"}}))

		err := ctx.editor.Expand(nil, pb.RpcBlockTableExpandRequest{
			TargetId: ctx.id,
			Columns:  2,
			Rows:     0,
		})
		require.NoError(t, err)

		want := mkTestTable([]string{"col1", "col2", "col3", "col4"}, []string{"row1", "row2"},
			[][]string{{"row1-col1", "row1-col2", "row1-col3", "row1-col4"}, {"row2-col1", "row2-col2", "row2-col3", "row2-col4"}})

		assertIsomorphic(t, want, ctx.s, ctx.wantMapping, ctx.gotMapping)
	})

	t.Run("rows only", func(t *testing.T) {
		ctx := newTableTestContext(t, 2, 2, true,
			mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2"}, [][]string{{"row1-col1", "row1-col2"}, {"row2-col1", "row2-col2"}}))

		err := ctx.editor.Expand(nil, pb.RpcBlockTableExpandRequest{
			TargetId: ctx.id,
			Columns:  0,
			Rows:     2,
		})
		require.NoError(t, err)

		tb, err := newTableBlockFromState(ctx.s, ctx.id)
		require.NoError(t, err)

		err = ctx.editor.RowListFill(nil, pb.RpcBlockTableRowListFillRequest{
			BlockIds: tb.rows().ChildrenIds,
		})
		require.NoError(t, err)

		want := mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2", "row3", "row4"},
			[][]string{{"row1-col1", "row1-col2"}, {"row2-col1", "row2-col2"}, {"row3-col1", "row3-col2"}, {"row4-col1", "row4-col2"}})

		assertIsomorphic(t, want, ctx.s, ctx.wantMapping, ctx.gotMapping)
	})

	t.Run("cols and rows", func(t *testing.T) {
		ctx := newTableTestContext(t, 2, 2, true,
			mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2"}, [][]string{{"row1-col1", "row1-col2"}, {"row2-col1", "row2-col2"}}))

		err := ctx.editor.Expand(nil, pb.RpcBlockTableExpandRequest{
			TargetId: ctx.id,
			Columns:  2,
			Rows:     2,
		})
		require.NoError(t, err)

		tb, err := newTableBlockFromState(ctx.s, ctx.id)
		require.NoError(t, err)

		err = ctx.editor.RowListFill(nil, pb.RpcBlockTableRowListFillRequest{
			BlockIds: tb.rows().ChildrenIds,
		})
		require.NoError(t, err)

		want := mkTestTable([]string{"col1", "col2", "col3", "col4"}, []string{"row1", "row2", "row3", "row4"},
			[][]string{{"row1-col1", "row1-col2", "row1-col3", "row1-col4"}, {"row2-col1", "row2-col2", "row2-col3", "row2-col4"}, {"row3-col1", "row3-col2", "row3-col3", "row3-col4"}, {"row4-col1", "row4-col2", "row4-col3", "row4-col4"}})

		assertIsomorphic(t, want, ctx.s, ctx.wantMapping, ctx.gotMapping)
	})
}

type tableTestContext struct {
	id          string
	editor      Table
	s           *state.State
	wantMapping map[string]string
	gotMapping  map[string]string
}

func newTableTestContext(t *testing.T, columnsCount, rowsCount uint32, filled bool, wantTable *state.State) tableTestContext {
	sb := smarttest.New("root")
	sb.AddBlock(simple.New(&model.Block{
		Id: "root",
	}))

	ctx := tableTestContext{}

	ctx.editor = New(sb)

	id, err := ctx.editor.TableCreate(nil, pb.RpcBlockTableCreateRequest{
		ContextId: "",
		TargetId:  "root",
		Position:  model.Block_Inner,
		Columns:   columnsCount,
		Rows:      rowsCount,
	})
	ctx.id = id
	ctx.s = sb.NewState()

	assert.NoError(t, err)
	assert.True(t, ctx.s.Exists(id))

	if filled {
		tb, err := newTableBlockFromState(ctx.s, id)
		require.NoError(t, err)

		err = ctx.editor.RowListFill(nil, pb.RpcBlockTableRowListFillRequest{
			BlockIds: tb.rows().ChildrenIds,
		})
		require.NoError(t, err)
	}

	ctx.wantMapping = map[string]string{}
	ctx.gotMapping = map[string]string{}

	assertIsomorphic(t, wantTable, ctx.s, ctx.wantMapping, ctx.gotMapping)

	return ctx
}

func reassignIds(s *state.State, mapping map[string]string) (*state.State, error) {
	err := s.Iterate(func(b simple.Block) bool {
		if _, ok := mapping[b.Model().Id]; !ok {
			id := strconv.Itoa(len(mapping) + 1)
			mapping[b.Model().Id] = id
		}
		return true
	})
	if err != nil {
		return nil, err
	}

	res := state.NewDoc("root", nil).NewState()
	err = s.Iterate(func(b simple.Block) bool {
		b = b.Copy()

		b.Model().Id = mapping[b.Model().Id]
		// Don't care about restrictions here
		b.Model().Restrictions = nil
		for i, id := range b.Model().ChildrenIds {
			b.Model().ChildrenIds[i] = mapping[id]
		}

		res.Add(b)
		return true
	})
	if err != nil {
		return nil, err
	}

	res.SetRootId(mapping["root"])

	return res, nil
}

// assertIsomorphic checks that two states have same structure
// Preserves mappings for tracking structure changes
func assertIsomorphic(t *testing.T, want, got *state.State, wantMapping, gotMapping map[string]string) {
	var err error
	want, err = reassignIds(want, wantMapping)
	require.NoError(t, err)
	got, err = reassignIds(got, gotMapping)
	require.NoError(t, err)

	var gotBlocks []simple.Block
	got.Iterate(func(b simple.Block) bool {
		gotBlocks = append(gotBlocks, b)
		return true
	})

	var wantBlocks []simple.Block
	want.Iterate(func(b simple.Block) bool {
		wantBlocks = append(wantBlocks, b)
		return true
	})

	assert.Equal(t, wantBlocks, gotBlocks)
}

func mkTestTable(columns []string, rows []string, cells [][]string) *state.State {
	s := state.NewDoc("root", nil).NewState()
	blocks := []*model.Block{
		{
			Id:          "root",
			ChildrenIds: []string{"table"},
		},
		{
			Id:          "table",
			ChildrenIds: []string{"columns", "rows"},
			Content:     &model.BlockContentOfTable{Table: &model.BlockContentTable{}},
		},
		{
			Id:          "columns",
			ChildrenIds: columns,
			Content: &model.BlockContentOfLayout{
				Layout: &model.BlockContentLayout{
					Style: model.BlockContentLayout_TableColumns,
				},
			},
		},
		{
			Id:          "rows",
			ChildrenIds: rows,
			Content: &model.BlockContentOfLayout{
				Layout: &model.BlockContentLayout{
					Style: model.BlockContentLayout_TableRows,
				},
			},
		},
	}

	for _, c := range columns {
		blocks = append(blocks, &model.Block{
			Id:      c,
			Content: &model.BlockContentOfTableColumn{TableColumn: &model.BlockContentTableColumn{}},
		})
	}

	cellsByRow := map[string][]string{}
	for _, cc := range cells {
		rowId, _, err := parseCellId(cc[0])
		if err != nil {
			panic(err)
		}
		cellsByRow[rowId] = cc

		for _, c := range cc {
			blocks = append(blocks, &model.Block{
				Id:      c,
				Content: &model.BlockContentOfText{Text: &model.BlockContentText{}},
			})
		}
	}

	for _, r := range rows {
		blocks = append(blocks, &model.Block{
			Id:          r,
			ChildrenIds: cellsByRow[r],
			Content:     &model.BlockContentOfTableRow{TableRow: &model.BlockContentTableRow{}},
		})
	}

	for _, b := range blocks {
		s.Add(simple.New(b))
	}
	return s
}
