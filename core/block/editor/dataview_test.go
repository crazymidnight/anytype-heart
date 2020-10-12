package editor

import (
	"testing"

	"github.com/anytypeio/go-anytype-middleware/core/block/editor/smartblock/smarttest"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/stretchr/testify/require"
)

func TestDataview_SetDetails(t *testing.T) {
	var event *pb.Event
	p := NewProfile(nil, nil, nil, nil, func(e *pb.Event) {
		event = e
	})
	p.SmartBlock = smarttest.New("1")

	err := p.SetDetails(nil, []*pb.RpcBlockSetDetailsDetail{
		{
			Key:   "key",
			Value: pbtypes.String("value"),
		},
	})
	require.NoError(t, err)
	require.NotNil(t, event)
}
