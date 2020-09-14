package simple

import (
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/globalsign/mgo/bson"
)

type BlockCreator = func(m *model.Block) Block

var (
	registry []BlockCreator
	fallback BlockCreator
)

func RegisterCreator(c BlockCreator) {
	registry = append(registry, c)
}

func RegisterFallback(c BlockCreator) {
	fallback = c
}

type Block interface {
	Model() *model.Block
	Diff(block Block) (msgs []*pb.EventMessage, err error)
	String() string
	Copy() Block
}

type FileHashes interface {
	FillFileHashes(hashes []string) []string
}

func New(block *model.Block) (b Block) {
	if block.Id == "" {
		block.Id = bson.NewObjectId().Hex()
	}
	for _, c := range registry {
		if b = c(block); b != nil {
			return
		}
	}
	return fallback(block)
}
