package editor

import (
	"testing"

	"github.com/anytypeio/go-anytype-middleware/core/block/editor/smartblock/smarttest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArchive_Init(t *testing.T) {
	a := NewArchive()
	a.SmartBlock = smarttest.New("root")
	require.NoError(t, a.Init(nil))
	assert.Len(t, a.Blocks(), 1)
}

func TestArchive_Archive(t *testing.T) {
	t.Run("archive", func(t *testing.T) {
		a := NewArchive()
		a.SmartBlock = smarttest.New("root")
		require.NoError(t, a.Init(nil))

		require.NoError(t, a.Archive("1"))
		require.NoError(t, a.Archive("2"))

		s := a.NewState()
		chIds := s.Get(s.RootId()).Model().ChildrenIds
		require.Len(t, chIds, 2)
		require.Equal(t, "2", s.Get(chIds[0]).Model().GetLink().TargetBlockId)
		require.Equal(t, "1", s.Get(chIds[1]).Model().GetLink().TargetBlockId)
	})
	t.Run("archive archived", func(t *testing.T) {
		a := NewArchive()
		a.SmartBlock = smarttest.New("root")
		require.NoError(t, a.Init(nil))

		require.NoError(t, a.Archive("1"))
		require.NoError(t, a.Archive("1"))

		s := a.NewState()
		chIds := s.Get(s.RootId()).Model().ChildrenIds
		require.Len(t, chIds, 1)
	})
}

func TestArchive_UnArchive(t *testing.T) {
	t.Run("unarchive", func(t *testing.T) {
		a := NewArchive()
		a.SmartBlock = smarttest.New("root")
		require.NoError(t, a.Init(nil))

		require.NoError(t, a.Archive("1"))
		require.NoError(t, a.Archive("2"))

		require.NoError(t, a.UnArchive("2"))
		s := a.NewState()
		chIds := s.Get(s.RootId()).Model().ChildrenIds
		require.Len(t, chIds, 1)
	})
	t.Run("unarchived", func(t *testing.T) {
		a := NewArchive()
		a.SmartBlock = smarttest.New("root")
		require.NoError(t, a.Init(nil))

		require.NoError(t, a.Archive("1"))

		require.NoError(t, a.UnArchive("2"))

		s := a.NewState()
		chIds := s.Get(s.RootId()).Model().ChildrenIds
		require.Len(t, chIds, 1)
	})
}
