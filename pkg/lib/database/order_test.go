package database

import (
	"bytes"
	"testing"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fastjson"
	"golang.org/x/text/collate"
	"golang.org/x/text/language"

	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/util/pbtypes"
)

func assertCompare(t *testing.T, order Order, a *types.Struct, b *types.Struct, expected int) {
	assert.Equal(t, expected, order.Compare(a, b))
	arena := &fastjson.Arena{}
	aJson := pbtypes.ProtoToJson(arena, a)
	bJson := pbtypes.ProtoToJson(arena, b)
	s := order.Compile()
	aBytes := s.AppendKey(nil, aJson)
	bBytes := s.AppendKey(nil, bJson)
	got := bytes.Compare(aBytes, bBytes)
	assert.Equal(t, expected, got)
}

func TestTextSort(t *testing.T) {
	t.Run("note layout, not empty name", func(t *testing.T) {
		a := &types.Struct{
			Fields: map[string]*types.Value{
				bundle.RelationKeyName.String(): pbtypes.String("b"),
			},
		}
		b := &types.Struct{
			Fields: map[string]*types.Value{
				bundle.RelationKeyName.String():    pbtypes.String("a"),
				bundle.RelationKeySnippet.String(): pbtypes.String("b"),
				bundle.RelationKeyLayout.String():  pbtypes.Int64(int64(model.ObjectType_note)),
			},
		}
		asc := &KeyOrder{Key: bundle.RelationKeyName.String(), Type: model.BlockContentDataviewSort_Asc, RelationFormat: model.RelationFormat_shorttext}
		assertCompare(t, asc, a, b, 1)
		desc := &KeyOrder{Key: bundle.RelationKeyName.String(), Type: model.BlockContentDataviewSort_Desc, RelationFormat: model.RelationFormat_shorttext}
		assertCompare(t, desc, a, b, -1)
	})
	t.Run("note layout, empty name", func(t *testing.T) {
		t.Run("one with name, one with snippet, not equal", func(t *testing.T) {
			a := &types.Struct{
				Fields: map[string]*types.Value{
					bundle.RelationKeyName.String(): pbtypes.String("a"),
				},
			}
			b := &types.Struct{
				Fields: map[string]*types.Value{
					bundle.RelationKeySnippet.String(): pbtypes.String("b"),
					bundle.RelationKeyLayout.String():  pbtypes.Int64(int64(model.ObjectType_note)),
				},
			}
			asc := &KeyOrder{Key: bundle.RelationKeyName.String(), Type: model.BlockContentDataviewSort_Asc, RelationFormat: model.RelationFormat_shorttext}
			assertCompare(t, asc, a, b, -1)
			desc := &KeyOrder{Key: bundle.RelationKeyName.String(), Type: model.BlockContentDataviewSort_Desc, RelationFormat: model.RelationFormat_shorttext}
			assertCompare(t, desc, a, b, 1)
		})
		t.Run("one with name, one with snippet, equal", func(t *testing.T) {
			a := &types.Struct{
				Fields: map[string]*types.Value{
					bundle.RelationKeyName.String(): pbtypes.String("a"),
				},
			}
			b := &types.Struct{
				Fields: map[string]*types.Value{
					bundle.RelationKeySnippet.String(): pbtypes.String("a"),
					bundle.RelationKeyLayout.String():  pbtypes.Int64(int64(model.ObjectType_note)),
				},
			}
			asc := &KeyOrder{Key: bundle.RelationKeyName.String(), Type: model.BlockContentDataviewSort_Asc, RelationFormat: model.RelationFormat_shorttext}
			assertCompare(t, asc, a, b, 0)
			desc := &KeyOrder{Key: bundle.RelationKeyName.String(), Type: model.BlockContentDataviewSort_Desc, RelationFormat: model.RelationFormat_shorttext}
			assertCompare(t, desc, a, b, 0)
		})
	})
}

func TestKeyOrder_Compare(t *testing.T) {
	t.Run("eq", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("a")}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("a")}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Asc, RelationFormat: model.RelationFormat_shorttext}
		assertCompare(t, asc, a, b, 0)
	})
	t.Run("asc", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("a")}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("b")}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Asc, RelationFormat: model.RelationFormat_shorttext}
		assertCompare(t, asc, a, b, -1)
	})

	t.Run("desc_float", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Float64(1)}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Float64(2)}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Desc, EmptyPlacement: model.BlockContentDataviewSort_End, RelationFormat: model.RelationFormat_number}
		assertCompare(t, asc, a, b, 1)
	})

	t.Run("asc_float", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Float64(1)}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Float64(2)}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Asc, EmptyPlacement: model.BlockContentDataviewSort_Start, RelationFormat: model.RelationFormat_number}
		assertCompare(t, asc, a, b, -1)
	})

	t.Run("asc_emptylast", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("a")}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("")}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Asc, EmptyPlacement: model.BlockContentDataviewSort_End, RelationFormat: model.RelationFormat_shorttext}
		assertCompare(t, asc, a, b, -1)
	})

	t.Run("asc_emptylast_float", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Float64(1)}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Null()}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Asc, EmptyPlacement: model.BlockContentDataviewSort_End, RelationFormat: model.RelationFormat_number}
		assertCompare(t, asc, a, b, -1)
	})

	t.Run("asc_emptylast_str", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("a")}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("")}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Asc, EmptyPlacement: model.BlockContentDataviewSort_End, RelationFormat: model.RelationFormat_shorttext}
		assertCompare(t, asc, a, b, -1)
	})

	t.Run("desc_emptylast_str", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("a")}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("")}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Desc, EmptyPlacement: model.BlockContentDataviewSort_End, RelationFormat: model.RelationFormat_shorttext}
		assertCompare(t, asc, a, b, -1)
	})

	t.Run("asc_emptyfirst_str", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("a")}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("")}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Asc, EmptyPlacement: model.BlockContentDataviewSort_Start, RelationFormat: model.RelationFormat_shorttext}
		assertCompare(t, asc, a, b, 1)
	})

	t.Run("desc_emptyfirst_str", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("a")}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("")}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Desc, EmptyPlacement: model.BlockContentDataviewSort_Start, RelationFormat: model.RelationFormat_shorttext}
		assertCompare(t, asc, a, b, 1)
	})

	t.Run("asc_str_end", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("a")}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("b")}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Asc, EmptyPlacement: model.BlockContentDataviewSort_End, RelationFormat: model.RelationFormat_shorttext}
		assertCompare(t, asc, a, b, -1)
	})

	t.Run("desc_str_end", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("a")}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("b")}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Desc, EmptyPlacement: model.BlockContentDataviewSort_End, RelationFormat: model.RelationFormat_shorttext}
		assertCompare(t, asc, a, b, 1)
	})

	t.Run("asc_str_start", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("a")}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("b")}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Asc, EmptyPlacement: model.BlockContentDataviewSort_Start, RelationFormat: model.RelationFormat_shorttext}
		assertCompare(t, asc, a, b, -1)
	})

	t.Run("desc_str_start", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("a")}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("b")}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Desc, EmptyPlacement: model.BlockContentDataviewSort_Start, RelationFormat: model.RelationFormat_shorttext}
		assertCompare(t, asc, a, b, 1)
	})

	date := time.Unix(-1, 0)

	t.Run("asc date, no time", func(t *testing.T) {
		date1 := time.Date(2020, 2, 4, 15, 22, 0, 0, time.UTC)
		date2 := time.Date(2020, 2, 4, 16, 26, 0, 0, time.UTC)

		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Int64(date1.Unix())}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Int64(date2.Unix())}}
		asc := &KeyOrder{
			Key:            "k",
			Type:           model.BlockContentDataviewSort_Asc,
			EmptyPlacement: model.BlockContentDataviewSort_End,
			IncludeTime:    false,
			RelationFormat: model.RelationFormat_date,
		}
		assertCompare(t, asc, a, b, 0)
	})

	t.Run("asc_date_end_empty", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Int64(date.Unix())}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": nil}}
		asc := &KeyOrder{
			Key:            "k",
			Type:           model.BlockContentDataviewSort_Asc,
			EmptyPlacement: model.BlockContentDataviewSort_End,
			IncludeTime:    false,
			RelationFormat: model.RelationFormat_date,
		}
		assertCompare(t, asc, a, b, -1)
	})

	t.Run("desc_date_end_empty", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Int64(date.Unix())}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": nil}}
		asc := &KeyOrder{
			Key:            "k",
			Type:           model.BlockContentDataviewSort_Desc,
			EmptyPlacement: model.BlockContentDataviewSort_End,
			IncludeTime:    false,
			RelationFormat: model.RelationFormat_date,
		}
		assertCompare(t, asc, a, b, -1)
	})

	t.Run("asc_date_start_empty", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Int64(date.Unix())}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": nil}}
		asc := &KeyOrder{
			Key:            "k",
			Type:           model.BlockContentDataviewSort_Asc,
			EmptyPlacement: model.BlockContentDataviewSort_Start,
			IncludeTime:    false,
			RelationFormat: model.RelationFormat_date,
		}
		assertCompare(t, asc, a, b, 1)
	})

	t.Run("desc_date_start", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Int64(date.Unix())}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": nil}}
		asc := &KeyOrder{
			Key:            "k",
			Type:           model.BlockContentDataviewSort_Desc,
			EmptyPlacement: model.BlockContentDataviewSort_Start,
			IncludeTime:    false,
			RelationFormat: model.RelationFormat_date,
		}
		assertCompare(t, asc, a, b, 1)
	})

	t.Run("asc_nil_emptylast", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("a")}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": nil}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Asc, EmptyPlacement: model.BlockContentDataviewSort_End, RelationFormat: model.RelationFormat_shorttext}
		assertCompare(t, asc, a, b, -1)
	})

	t.Run("asc_nil_emptylast_float", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Float64(1)}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": nil}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Asc, EmptyPlacement: model.BlockContentDataviewSort_End, RelationFormat: model.RelationFormat_number}
		assertCompare(t, asc, a, b, -1)
	})

	t.Run("desc_nil_emptylast", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("a")}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": nil}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Desc, EmptyPlacement: model.BlockContentDataviewSort_End, RelationFormat: model.RelationFormat_shorttext}
		assertCompare(t, asc, a, b, -1)
	})

	t.Run("desc_nil_emptylast_float", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Float64(1)}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": nil}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Desc, EmptyPlacement: model.BlockContentDataviewSort_End, RelationFormat: model.RelationFormat_number}
		assertCompare(t, asc, a, b, -1)
	})

	t.Run("asc_nil_emptyfirst", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("a")}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": nil}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Asc, EmptyPlacement: model.BlockContentDataviewSort_Start, RelationFormat: model.RelationFormat_shorttext}
		assertCompare(t, asc, a, b, 1)
	})

	t.Run("asc_nil_emptyfirst_float", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Float64(1)}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": nil}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Asc, EmptyPlacement: model.BlockContentDataviewSort_Start, RelationFormat: model.RelationFormat_number}
		assertCompare(t, asc, a, b, 1)
	})

	t.Run("desc_nil_emptyfirst", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("a")}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": nil}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Desc, EmptyPlacement: model.BlockContentDataviewSort_Start, RelationFormat: model.RelationFormat_shorttext}
		assertCompare(t, asc, a, b, 1)
	})

	t.Run("desc_nil_emptyfirst_float", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Float64(1)}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": nil}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Desc, EmptyPlacement: model.BlockContentDataviewSort_Start, RelationFormat: model.RelationFormat_number}
		assertCompare(t, asc, a, b, 1)
	})

	t.Run("nil and empty string -- equal", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": nil}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("")}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Asc, RelationFormat: model.RelationFormat_shorttext}
		assertCompare(t, asc, a, b, 0)
	})

	t.Run("asc_nil", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": nil}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Float64(0)}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Asc, RelationFormat: model.RelationFormat_number}
		assertCompare(t, asc, a, b, -1)
	})

	t.Run("asc_notspecified", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("a")}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("")}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Asc, RelationFormat: model.RelationFormat_shorttext}
		assertCompare(t, asc, a, b, 1)
	})

	t.Run("asc_notspecified", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Float64(1)}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Float64(0)}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Asc, RelationFormat: model.RelationFormat_number}
		assertCompare(t, asc, a, b, 1)
	})

	t.Run("desc_notspecified", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("a")}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("b")}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Desc, RelationFormat: model.RelationFormat_shorttext}
		assertCompare(t, asc, a, b, 1)
	})

	t.Run("desc_notspecified_float", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Float64(0)}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Float64(1)}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Desc, RelationFormat: model.RelationFormat_number}
		assertCompare(t, asc, a, b, 1)
	})
}

func TestKeyUnicodeOrder_Compare(t *testing.T) {
	t.Run("asc", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("Єгипет")}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("Японія")}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Asc, RelationFormat: model.RelationFormat_shorttext}
		assertCompare(t, asc, a, b, -1)
	})

	t.Run("dsc", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("Ürkmez")}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("Zurich")}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Desc, RelationFormat: model.RelationFormat_shorttext}
		assertCompare(t, asc, a, b, 1)
	})
}

func TestCollate(t *testing.T) {
	t.Run("asc", func(t *testing.T) {
		col := collate.New(language.Und, collate.IgnoreCase)

		s1 := "Єгипет"
		s2 := "Японія"

		buf := &collate.Buffer{}
		k1 := col.KeyFromString(buf, s1)
		k2 := col.KeyFromString(buf, s2)
		assert.Equal(t, -1, bytes.Compare(k1, k2))
		assert.Equal(t, 1, bytes.Compare(k2, k1))
		buf.Reset()
	})
}

func TestSetOrder_Compare(t *testing.T) {
	so := SetOrder{
		&KeyOrder{Key: "a", Type: model.BlockContentDataviewSort_Asc, RelationFormat: model.RelationFormat_shorttext},
		&KeyOrder{Key: "b", Type: model.BlockContentDataviewSort_Desc, RelationFormat: model.RelationFormat_shorttext},
	}
	t.Run("eq", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"a": pbtypes.String("a"), "b": pbtypes.String("b")}}
		b := &types.Struct{Fields: map[string]*types.Value{"a": pbtypes.String("a"), "b": pbtypes.String("b")}}
		assertCompare(t, so, a, b, 0)
	})
	t.Run("first order", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"a": pbtypes.String("b"), "b": pbtypes.String("a")}}
		b := &types.Struct{Fields: map[string]*types.Value{"a": pbtypes.String("a"), "b": pbtypes.String("b")}}
		assertCompare(t, so, a, b, 1)

	})
	t.Run("second order", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"a": pbtypes.String("b"), "b": pbtypes.String("b")}}
		b := &types.Struct{Fields: map[string]*types.Value{"a": pbtypes.String("b"), "b": pbtypes.String("a")}}
		assertCompare(t, so, a, b, -1)
	})
}

func TestCustomOrder_Compare(t *testing.T) {
	idxIndices := map[string]int{
		"b": 0,
		"c": 1,
		"d": 2,
		"a": 3,
	}
	co := NewCustomOrder("ID", idxIndices, KeyOrder{Key: "ID", Type: model.BlockContentDataviewSort_Asc, RelationFormat: model.RelationFormat_shorttext})

	t.Run("gt", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"ID": pbtypes.String("c")}}
		b := &types.Struct{Fields: map[string]*types.Value{"ID": pbtypes.String("a")}}
		assertCompare(t, co, a, b, -1)
	})

	t.Run("eq", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"ID": pbtypes.String("a")}}
		b := &types.Struct{Fields: map[string]*types.Value{"ID": pbtypes.String("a")}}
		assertCompare(t, co, a, b, 0)
	})

	t.Run("lt", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"ID": pbtypes.String("a")}}
		b := &types.Struct{Fields: map[string]*types.Value{"ID": pbtypes.String("b")}}
		assertCompare(t, co, a, b, 1)
	})

	t.Run("first found second not", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"ID": pbtypes.String("a")}}
		b := &types.Struct{Fields: map[string]*types.Value{"ID": pbtypes.String("x")}}
		assertCompare(t, co, a, b, -1)
	})

	t.Run("first not found second yes", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"ID": pbtypes.String("x")}}
		b := &types.Struct{Fields: map[string]*types.Value{"ID": pbtypes.String("a")}}
		assertCompare(t, co, a, b, 1)
	})

	t.Run("both not found gt", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"ID": pbtypes.String("y")}}
		b := &types.Struct{Fields: map[string]*types.Value{"ID": pbtypes.String("z")}}
		assertCompare(t, co, a, b, -1)
	})

	t.Run("both not found eq", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"ID": pbtypes.String("z")}}
		b := &types.Struct{Fields: map[string]*types.Value{"ID": pbtypes.String("z")}}
		assertCompare(t, co, a, b, 0)
	})

	t.Run("both not found lt", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"ID": pbtypes.String("z")}}
		b := &types.Struct{Fields: map[string]*types.Value{"ID": pbtypes.String("y")}}
		assertCompare(t, co, a, b, 1)
	})
}

func TestTagStatusOrder_Compare(t *testing.T) {

	for _, relation := range []model.RelationFormat{model.RelationFormat_tag, model.RelationFormat_status} {
		t.Run("eq", func(t *testing.T) {
			a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("a")}}
			b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("a")}}
			asc := &KeyOrder{
				Key:            "k",
				Type:           model.BlockContentDataviewSort_Asc,
				RelationFormat: relation,
				Options:        map[string]string{"a": "a"},
			}
			assertCompare(t, asc, a, b, 0)
		})

		t.Run("asc", func(t *testing.T) {
			a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("b")}}
			b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.String("a")}}
			asc := &KeyOrder{
				Key:            "k",
				Type:           model.BlockContentDataviewSort_Asc,
				RelationFormat: relation,
				Options: map[string]string{
					"b": "a",
					"a": "b",
				},
			}
			assertCompare(t, asc, a, b, -1)
		})
	}
}

func TestIncludeTime_Compare(t *testing.T) {
	date := time.Unix(1672012800, 0)

	t.Run("date only eq", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Int64(date.Add(time.Second * 5).Unix())}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Int64(date.Add(time.Second * 10).Unix())}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Asc,
			IncludeTime: false, RelationFormat: model.RelationFormat_date}
		assertCompare(t, asc, a, b, 0)
	})

	t.Run("only date lt", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Int64(date.Unix())}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Int64(date.Add(time.Hour * 24).Unix())}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Asc,
			IncludeTime: false, RelationFormat: model.RelationFormat_date}
		assertCompare(t, asc, a, b, -1)
	})

	t.Run("date includeTime eq", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Int64(date.Add(time.Second * 10).Unix())}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Int64(date.Add(time.Second * 10).Unix())}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Asc,
			IncludeTime: true, RelationFormat: model.RelationFormat_date}
		assertCompare(t, asc, a, b, 0)
	})

	t.Run("date includeTime lt", func(t *testing.T) {
		a := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Int64(date.Add(time.Second * 5).Unix())}}
		b := &types.Struct{Fields: map[string]*types.Value{"k": pbtypes.Int64(date.Add(time.Second * 10).Unix())}}
		asc := &KeyOrder{Key: "k", Type: model.BlockContentDataviewSort_Asc,
			IncludeTime: true, RelationFormat: model.RelationFormat_date}
		assertCompare(t, asc, a, b, -1)
	})

}
