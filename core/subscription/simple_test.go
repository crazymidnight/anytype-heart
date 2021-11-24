package subscription

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimpleSub_Changes(t *testing.T) {
	t.Run("add to set", func(t *testing.T) {
		sub := &simpleSub{
			keys:             []string{"id", "order"},
			cache:            newCache(),
		}
		require.NoError(t, sub.init(genEntries(10, false)))
		ctx := &opCtx{}
		sub.onChangeBatch(ctx, genEntry("id5", 109))
		assertCtxChange(t, ctx, "id5")
	})
	t.Run("miss set", func(t *testing.T) {
		sub := &simpleSub{
			keys:             []string{"id", "order"},
			cache:            newCache(),
		}
		require.NoError(t, sub.init(genEntries(10, false)))
		ctx := &opCtx{}
		sub.onChangeBatch(ctx, genEntry("id50", 100))
		assertCtxEmpty(t, ctx)
	})
}

func TestSimpleSub_Refill(t *testing.T) {
	sub := &simpleSub{
		keys:             []string{"id", "order"},
		cache:            newCache(),
	}
	require.NoError(t, sub.init(genEntries(3, false)))
	ctx := &opCtx{}
	sub.refill(ctx, []*entry{genEntry("id3", 100), genEntry("id20", 200)})
	assertCtxChange(t, ctx, "id3")
	assertCtxRemove(t, ctx, "id1", "id2")
	assertCtxAdd(t, ctx, "id20", "")
}