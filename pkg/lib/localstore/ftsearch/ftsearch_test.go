package ftsearch

import (
	"context"
	"encoding/json"
	"github.com/anytypeio/go-anytype-middleware/app/testapp"
	"github.com/anytypeio/go-anytype-middleware/core/wallet"
	"github.com/blevesearch/bleve/v2"
	"github.com/golang/mock/gomock"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fixture struct {
	ft   FTSearch
	ta   *testapp.TestApp
	ctrl *gomock.Controller
}

func newFixture(path string, t *testing.T) *fixture {
	ft := New()
	ta := testapp.New().
		With(wallet.NewWithRepoPathAndKeys(path, nil, nil)).
		With(ft)

	require.NoError(t, ta.Start(context.Background()))
	return &fixture{
		ft: ft,
		ta: ta,
	}
}

func TestNewFTSearch(t *testing.T) {
	tmpDir, _ := os.MkdirTemp("", "")
	assertSearch(t, tmpDir)
	assertThaiSubstrFound(t, tmpDir)
	assertChineseFound(t, tmpDir)
	assertFoundPartsOfTheWords(t, tmpDir)
}

func assertChineseFound(t *testing.T, tmpDir string) {
	fixture := newFixture(tmpDir, t)
	ft := fixture.ft
	require.NoError(t, ft.Index(SearchDoc{
		Id:    "1",
		Title: "",
		Text:  "你好",
	}))
	require.NoError(t, ft.Index(SearchDoc{
		Id:    "2",
		Title: "",
		Text:  "交代",
	}))
	require.NoError(t, ft.Index(SearchDoc{
		Id:    "3",
		Title: "",
		Text:  "长江大桥",
	}))

	queries := []string{
		"你好世界",
		"亲口交代",
		"长江",
	}

	for _, qry := range queries {
		validateSearch(t, ft, qry, 1)
	}

	_ = ft.Close(nil)
}

func assertThaiSubstrFound(t *testing.T, tmpDir string) {
	fixture := newFixture(tmpDir, t)
	ft := fixture.ft
	require.NoError(t, ft.Index(SearchDoc{
		Id:    "test",
		Title: "ตัวอย่าง",
		Text:  "พรระเจ้า \n kumamon",
	}))

	validateSearch(t, ft, "ระเ", 1)
	validateSearch(t, ft, "ระเ ma", 1)

	_ = ft.Close(nil)
}

func assertSearch(t *testing.T, tmpDir string) {
	fixture := newFixture(tmpDir, t)
	ft := fixture.ft
	require.NoError(t, ft.Index(SearchDoc{
		Id:    "test",
		Title: "one",
		Text:  "two",
	}))

	validateSearch(t, ft, "one", 1)
	validateSearch(t, ft, "two", 1)

	_ = ft.Close(nil)
}

func assertFoundPartsOfTheWords(t *testing.T, tmpDir string) {
	fixture := newFixture(tmpDir, t)
	ft := fixture.ft
	require.NoError(t, ft.Index(SearchDoc{
		Id:    "1",
		Title: "This is the title",
		Text:  "two",
	}))
	require.NoError(t, ft.Index(SearchDoc{
		Id:    "2",
		Title: "is the title",
		Text:  "two",
	}))

	validateSearch(t, ft, "this", 1)
	validateSearch(t, ft, "his", 1)
	validateSearch(t, ft, "is", 2)
	validateSearch(t, ft, "i t", 2)

	_ = ft.Close(nil)
}

func validateSearch(t *testing.T, ft FTSearch, qry string, times int) {
	res, err := ft.Search(qry)
	require.NoError(t, err)
	assert.Len(t, res, times)
}

func TestChineseSearch(t *testing.T) {
	//given
	index := givenPrefilledChineseIndex()
	defer func() { _ = index.Close() }()

	expected := givenExpectedChinese()

	//when
	queries := []string{
		"你好世界",
		"亲口交代",
		"长江",
	}

	//then
	result := validateChinese(queries, index)
	assert.Equal(t, expected, result)
}

func prettify(res *bleve.SearchResult) string {
	type Result struct {
		Id    string  `json:"id"`
		Score float64 `json:"score"`
	}
	results := []Result{}
	for _, item := range res.Hits {
		results = append(results, Result{item.ID, item.Score})
	}
	b, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func validateChinese(queries []string, index bleve.Index) [3]string {
	result := [3]string{}
	for i, q := range queries {
		req := bleve.NewSearchRequest(bleve.NewQueryStringQuery(q))
		req.Highlight = bleve.NewHighlight()
		res, err := index.Search(req)
		if err != nil {
			panic(err)
		}
		result[i] = prettify(res)
	}
	return result
}

func givenExpectedChinese() [3]string {
	return [3]string{
		`[{"id":"1","score":0.3192794660708729}]`,
		`[{"id":"2","score":0.3192794660708729}]`,
		`[{"id":"3","score":0.8888941720598743}]`,
	}
}

func givenPrefilledChineseIndex() bleve.Index {
	tmpDir, _ := os.MkdirTemp("", "")
	messages := []struct {
		Id   string
		Text string
	}{
		{
			Id:   "1",
			Text: "你好",
		},
		{
			Id:   "2",
			Text: "交代",
		},
		{
			Id:   "3",
			Text: "长江大桥",
		},
	}

	indexMapping := makeMapping()

	index, err := bleve.New(tmpDir, indexMapping)
	if err != nil {
		panic(err)
	}
	for _, msg := range messages {
		if err := index.Index(msg.Id, msg); err != nil {
			panic(err)
		}
	}
	return index
}
