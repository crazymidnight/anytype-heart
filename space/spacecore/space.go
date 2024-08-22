package spacecore

import (
	"context"
	"time"

	"github.com/anyproto/any-sync/commonspace"

	"github.com/anyproto/anytype-heart/core/syncstatus/objectsyncstatus"
)

func newAnySpace(cc commonspace.Space) (*AnySpace, error) {
	return &AnySpace{Space: cc}, nil
}

type AnySpace struct {
	commonspace.Space
	statusWatcher objectsyncstatus.StatusWatcher
}

func (s *AnySpace) Init(ctx context.Context) (err error) {
	err = s.Space.Init(ctx)
	if err != nil {
		return
	}
	return
}

func (s *AnySpace) TryClose(objectTTL time.Duration) (close bool, err error) {
	return false, nil
}

func (s *AnySpace) Close() (err error) {
	return s.Space.Close()
}
