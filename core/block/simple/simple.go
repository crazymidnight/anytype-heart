package simple

import (
	"github.com/anytypeio/go-anytype-library/pb/model"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple/base"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple/text"
	"github.com/mohae/deepcopy"
)

type Block interface {
	Virtual() bool
	Model() *model.Block
}

func New(block *model.Block) (b Block) {
	if block.Content == nil {
		return base.NewBase(block)
	}
	switch block.Content.(type) {
	case *model.BlockContentOfText:
		return text.NewText(block)
	default:
		return base.NewBase(block)
	}
}

func NewVirtual(block *model.Block) Block {
	return base.NewVirtual(block)
}

func Copy(b Block) Block {
	return deepcopy.Copy(b).(Block)
}
