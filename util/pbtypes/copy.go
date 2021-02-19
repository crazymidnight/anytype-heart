package pbtypes

import (
	"github.com/anytypeio/go-anytype-middleware/util/slice"
	"sync"

	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	pbrelation "github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/relation"
	"github.com/gogo/protobuf/types"
)

var bytesPool = &sync.Pool{
	New: func() interface{} {
		return []byte{}
	},
}

func CopyBlock(in *model.Block) (out *model.Block) {
	buf := bytesPool.Get().([]byte)
	size := in.Size()
	if cap(buf) < size {
		buf = make([]byte, 0, size*2)
	}
	size, _ = in.MarshalToSizedBuffer(buf[:size])
	out = &model.Block{}
	_ = out.Unmarshal(buf[:size])
	bytesPool.Put(buf)
	return
}

func CopyStruct(in *types.Struct) (out *types.Struct) {
	if in == nil {
		return nil
	}
	buf := bytesPool.Get().([]byte)
	size := in.Size()
	if cap(buf) < size {
		buf = make([]byte, 0, size*2)
	}
	size, _ = in.MarshalToSizedBuffer(buf[:size])
	out = &types.Struct{}
	_ = out.Unmarshal(buf[:size])
	if out.Fields == nil && in.Fields != nil {
		out.Fields = make(map[string]*types.Value)
	}
	bytesPool.Put(buf)
	return
}

func CopyVal(in *types.Value) (out *types.Value) {
	if in == nil {
		return nil
	}
	buf := bytesPool.Get().([]byte)
	size := in.Size()
	if cap(buf) < size {
		buf = make([]byte, 0, size*2)
	}
	size, _ = in.MarshalToSizedBuffer(buf[:size])
	out = &types.Value{}
	_ = out.Unmarshal(buf[:size])

	bytesPool.Put(buf)
	return
}

func CopyRelation(in *pbrelation.Relation) (out *pbrelation.Relation) {
	if in == nil {
		return nil
	}
	buf := bytesPool.Get().([]byte)
	size := in.Size()
	if cap(buf) < size {
		buf = make([]byte, 0, size*2)
	}
	size, _ = in.MarshalToSizedBuffer(buf[:size])
	out = &pbrelation.Relation{}
	_ = out.Unmarshal(buf[:size])

	bytesPool.Put(buf)
	return out
}

func CopyObjectType(in *pbrelation.ObjectType) (out *pbrelation.ObjectType) {
	if in == nil {
		return nil
	}

	buf := bytesPool.Get().([]byte)
	size := in.Size()
	if cap(buf) < size {
		buf = make([]byte, 0, size*2)
	}
	size, _ = in.MarshalToSizedBuffer(buf[:size])
	out = &pbrelation.ObjectType{}
	_ = out.Unmarshal(buf[:size])

	bytesPool.Put(buf)
	return out
}

func CopyRelations(in []*pbrelation.Relation) (out []*pbrelation.Relation) {
	if in == nil {
		return nil
	}
	buf := bytesPool.Get().([]byte)
	inWrapped := pbrelation.Relations{Relations: in}
	size := inWrapped.Size()
	if cap(buf) < size {
		buf = make([]byte, 0, size*2)
	}
	size, _ = inWrapped.MarshalToSizedBuffer(buf[:size])
	outWrapped := &pbrelation.Relations{}
	_ = outWrapped.Unmarshal(buf[:size])

	bytesPool.Put(buf)
	return outWrapped.Relations
}

func CopyOptions(in []*pbrelation.RelationOption) (out []*pbrelation.RelationOption) {
	if in == nil {
		return nil
	}

	for _, inO := range in {
		inCopy := *inO
		out = append(out, &inCopy)
	}
	return
}

func CopyRelationsToMap(in []*pbrelation.Relation) (out map[string]*pbrelation.Relation) {
	out = make(map[string]*pbrelation.Relation, len(in))
	rels := CopyRelations(in)
	for _, rel := range rels {
		out[rel.Key] = rel
	}

	return
}

func RelationsFilterKeys(in []*pbrelation.Relation, keys []string) (out []*pbrelation.Relation) {
	for i, inRel := range in {
		if slice.FindPos(keys, inRel.Key) >= 0 {
			out = append(out, in[i])
		}
	}
	return
}

func StructNotNilKeys(st *types.Struct) (keys []string) {
	if st == nil || st.Fields == nil {
		return nil
	}

	for k, v := range st.Fields {
		if v != nil {
			keys = append(keys, k)
		}
	}
	return
}
