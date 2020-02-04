package block

import (
	"fmt"
	"sync"

	"github.com/anytypeio/go-anytype-library/pb/model"
	"github.com/anytypeio/go-anytype-middleware/core/anytype"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/google/uuid"
)

func newBreadcrumbs(s *service) smartBlock {
	return &breadcrumbs{
		s:  s,
		id: uuid.New().String(),
	}
}

type breadcrumbs struct {
	emptySmart
	id     string
	s      *service
	blocks map[string]simple.Block
	mu     sync.Mutex
}

func (b *breadcrumbs) Open(_ anytype.Block) error {
	b.blocks = map[string]simple.Block{
		b.id: simple.New(&model.Block{
			Id: b.id,
			Content: &model.BlockContentOfPage{
				Page: &model.BlockContentPage{
					Style: model.BlockContentPage_Breadcrumbs,
				},
			},
		}),
	}
	return nil
}

func (b *breadcrumbs) Init() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.show()
}

func (b *breadcrumbs) Show() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.show()
	return nil
}

func (b *breadcrumbs) show() {
	blocks := make([]*model.Block, 0, len(b.blocks))
	for _, b := range b.blocks {
		blocks = append(blocks, b.Model())
	}
	b.s.sendEvent(&pb.Event{
		Messages: []*pb.EventMessage{
			{
				Value: &pb.EventMessageValueOfBlockShow{
					BlockShow: &pb.EventBlockShow{
						RootId: b.id,
						Blocks: blocks,
					},
				},
			},
		},
		ContextId: b.id,
	})
}

func (b *breadcrumbs) UpdateBlock(ids []string, hist bool, apply func(b simple.Block) error) (err error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	event := &pb.Event{
		ContextId: b.id,
	}

	for _, id := range ids {
		if block, ok := b.blocks[id]; ok {
			copy := block.Copy()
			if err = apply(copy); err != nil {
				return
			}
			msgs, e := block.Diff(copy)
			if e != nil {
				return e
			}
			if len(msgs) > 0 {
				event.Messages = append(event.Messages, msgs...)
				b.blocks[id] = copy
			}
		}
	}

	b.s.sendEvent(event)
	return
}

func (b *breadcrumbs) createLink(targetId string) simple.Block {
	return simple.New(&model.Block{
		Id: uuid.New().String(),
		Content: &model.BlockContentOfLink{
			Link: &model.BlockContentLink{
				TargetBlockId: targetId,
				Style:         model.BlockContentLink_Page,
			},
		},
	})
}

func (b *breadcrumbs) GetId() string {
	return b.id
}

func (b *breadcrumbs) OnSmartOpen(id string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	linkIds := b.blocks[b.id].Model().ChildrenIds
	newLink := b.createLink(id)
	b.blocks[newLink.Model().Id] = newLink
	b.blocks[b.id].Model().ChildrenIds = append(linkIds, newLink.Model().Id)
	b.s.ls.onCreate(b, newLink)

	event := &pb.Event{
		ContextId: b.id,
	}
	event.Messages = append(event.Messages, &pb.EventMessage{
		Value: &pb.EventMessageValueOfBlockAdd{
			BlockAdd: &pb.EventBlockAdd{
				Blocks: []*model.Block{newLink.Model()},
			},
		},
	})
	event.Messages = append(event.Messages, &pb.EventMessage{
		Value: &pb.EventMessageValueOfBlockSetChildrenIds{
			BlockSetChildrenIds: &pb.EventBlockSetChildrenIds{
				Id:          b.id,
				ChildrenIds: b.blocks[b.id].Model().ChildrenIds,
			},
		},
	})

	b.s.sendEvent(event)
	return
}

func (b *breadcrumbs) Cut(index int) (err error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	linkIds := b.blocks[b.id].Model().ChildrenIds
	if len(linkIds) < index {
		return fmt.Errorf("index out of range: %d vs %d", index, len(linkIds))
	}
	if len(linkIds) == index {
		return
	}

	toRemove := linkIds[index:]
	b.blocks[b.id].Model().ChildrenIds = linkIds[:index]

	event := &pb.Event{
		ContextId: b.id,
	}
	for _, removeId := range toRemove {
		event.Messages = append(event.Messages, &pb.EventMessage{
			Value: &pb.EventMessageValueOfBlockDelete{
				BlockDelete: &pb.EventBlockDelete{
					BlockId: removeId,
				},
			},
		})
	}
	event.Messages = append(event.Messages, &pb.EventMessage{
		Value: &pb.EventMessageValueOfBlockSetChildrenIds{
			BlockSetChildrenIds: &pb.EventBlockSetChildrenIds{
				Id:          b.id,
				ChildrenIds: b.blocks[b.id].Model().ChildrenIds,
			},
		},
	})

	b.s.sendEvent(event)
	return
}

func (b *breadcrumbs) Close() error {
	for _, l := range b.blocks {
		b.s.ls.onDelete(b, l)
	}
	return nil
}

func (b *breadcrumbs) Anytype() anytype.Anytype {
	return b.s.anytype
}
