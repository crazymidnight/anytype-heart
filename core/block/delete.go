package block

import (
	"context"
	"fmt"
	"github.com/anyproto/any-sync/commonspace"
	"github.com/anyproto/anytype-heart/core/block/editor/smartblock"
	"github.com/anyproto/anytype-heart/core/domain"
	"github.com/anyproto/anytype-heart/pb"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	coresb "github.com/anyproto/anytype-heart/pkg/lib/core/smartblock"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/util/pbtypes"
	"go.uber.org/zap"
)

func (s *Service) DeleteObject(objectID string) (err error) {
	spaceID, err := s.spaceService.ResolveSpaceID(objectID)
	if err != nil {
		return fmt.Errorf("resolve spaceID: %w", err)
	}
	err = Do(s, objectID, func(b smartblock.SmartBlock) error {
		if err = b.Restrictions().Object.Check(model.Restrictions_Delete); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}

	id := domain.FullID{
		SpaceID:  spaceID,
		ObjectID: objectID,
	}
	sbt, _ := s.sbtProvider.Type(spaceID, objectID)
	switch sbt {
	case coresb.SmartBlockTypeSubObject:
		return fmt.Errorf("subobjects deprecated")
	case coresb.SmartBlockTypeFile:
		err = s.OnDelete(id, func() error {
			if err := s.fileStore.DeleteFile(objectID); err != nil {
				return err
			}
			if err := s.fileSync.RemoveFile(spaceID, objectID); err != nil {
				return fmt.Errorf("failed to remove file from sync: %w", err)
			}
			_, err = s.fileService.FileOffload(context.Background(), objectID, true)
			if err != nil {
				return err
			}
			return nil
		})
	default:
		var space commonspace.Space
		space, err = s.spaceService.GetSpace(context.Background(), spaceID)
		if err != nil {
			return
		}
		// this will call DeleteTree asynchronously in the end
		return space.DeleteTree(context.Background(), objectID)
	}
	if err != nil {
		return
	}

	s.sendOnRemoveEvent(objectID)
	_, err = s.cache.Remove(context.Background(), objectID)
	return
}

func (s *Service) OnDelete(id domain.FullID, workspaceRemove func() error) error {
	var (
		isFavorite bool
	)

	err := Do(s, id.ObjectID, func(b smartblock.SmartBlock) error {
		b.ObjectCloseAllSessions()
		st := b.NewState()
		isFavorite = pbtypes.GetBool(st.LocalDetails(), bundle.RelationKeyIsFavorite.String())
		if isFavorite {
			_ = s.SetPageIsFavorite(pb.RpcObjectSetIsFavoriteRequest{IsFavorite: false, ContextId: id.ObjectID})
		}
		b.SetIsDeleted()
		if workspaceRemove != nil {
			return workspaceRemove()
		}
		return nil
	})
	if err != nil {
		log.Error("failed to perform delete operation on object", zap.Error(err))
	}
	if err := s.objectStore.DeleteObject(id.ObjectID); err != nil {
		return fmt.Errorf("delete object from local store: %w", err)
	}

	return nil
}

func (s *Service) DeleteSpace(ctx context.Context, spaceID string) error {
	log.Debug("space deleted", zap.String("spaceID", spaceID))
	return nil
}

func (s *Service) sendOnRemoveEvent(ids ...string) {
	s.eventSender.Broadcast(&pb.Event{
		Messages: []*pb.EventMessage{
			{
				Value: &pb.EventMessageValueOfObjectRemove{
					ObjectRemove: &pb.EventObjectRemove{
						Ids: ids,
					},
				},
			},
		},
	})
}
