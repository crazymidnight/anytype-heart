package change

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestStateBuildCases(t *testing.T) {
	require.NoError(t, filepath.Walk("./testcases", func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(info.Name()) == ".yml" {
			t.Run(strings.ReplaceAll(info.Name(), ".yml", ""), func(t *testing.T) {
				doTestCaseFile(t, path)
			})
		}
		return nil
	}))
}

type TestCase struct {
	Name     string
	Init     *DocStruct
	Changes  []*TestChange
	Expected *DocStruct
}

type DocStruct struct {
	Id    string
	Child []*DocStruct
}

type TestChange struct {
	Type  string
	Error bool
	Data  *TestChangeContent
}

type TestChangeContent struct {
	unmarshal func(interface{}) error
}

func (t *TestChangeContent) UnmarshalYAML(unmarshal func(interface{}) error) error {
	t.unmarshal = unmarshal
	return nil
}

func (t *TestChangeContent) GetChange(tp string) *pb.ChangeContent {
	var value pb.IsChangeContentValue
	switch tp {
	case "move":
		bm := &pb.ChangeBlockMove{}
		t.unmarshal(&bm)
		value = &pb.ChangeContentValueOfBlockMove{
			BlockMove: bm,
		}
	}
	return &pb.ChangeContent{
		Value: value,
	}
}

func doTestCaseFile(t *testing.T, filename string) {
	data, err := ioutil.ReadFile(filename)
	require.NoError(t, err)
	var cases []*TestCase
	err = yaml.Unmarshal(data, &cases)
	require.NoError(t, err)
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			doTestCase(t, tc)
		})
	}
}

func doTestCase(t *testing.T, tc *TestCase) {
	d := state.NewDoc("root", nil).(*state.State)
	if tc.Init.Id == "" {
		tc.Init.Id = "root"
	}
	var fillStruct func(s *state.State, d *DocStruct)
	fillStruct = func(s *state.State, d *DocStruct) {
		b := &model.Block{Id: d.Id}
		for _, ch := range d.Child {
			fillStruct(s, ch)
			b.ChildrenIds = append(b.ChildrenIds, ch.Id)
		}
		s.Add(simple.New(b))
	}
	fillStruct(d, tc.Init)

	for i, ch := range tc.Changes {
		s := d.NewState()
		chc := ch.Data.GetChange(ch.Type)
		err := s.ApplyChange(chc)
		if !ch.Error {
			require.NoError(t, err, fmt.Sprintf("index: %d; change: %s", i, chc.String()))
		}

		_, act, err := state.ApplyState(s, false)
		require.NoError(t, err)
		t.Log(d.String(), act)
	}

	exp := state.NewDoc("root", nil).(*state.State)
	if tc.Expected.Id == "" {
		tc.Expected.Id = "root"
	}
	fillStruct(exp, tc.Expected)
	t.Log(d.String())
	assert.Equal(t, stateToTestString(exp), stateToTestString(d))
}

func stateToTestString(s *state.State) string {
	var repl = make(map[string]string)
	s.Iterate(func(b simple.Block) (isContinue bool) {
		if layout := b.Model().GetLayout(); layout != nil {
			repl[b.Model().Id] = "/" + strings.ToLower(layout.Style.String()) + "/"
		}
		return true
	})
	str := s.String()
	for k, v := range repl {
		str = strings.ReplaceAll(str, k, v)
	}
	return str
}
