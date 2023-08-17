package editor

import (
	"strings"

	"github.com/anyproto/anytype-heart/core/block/editor/bookmark"
	"github.com/anyproto/anytype-heart/core/block/editor/converter"
	"github.com/anyproto/anytype-heart/core/block/editor/file"
	"github.com/anyproto/anytype-heart/core/block/editor/smartblock"
	"github.com/anyproto/anytype-heart/core/block/editor/state"
	"github.com/anyproto/anytype-heart/core/block/editor/template"
	"github.com/anyproto/anytype-heart/core/block/getblock"
	"github.com/anyproto/anytype-heart/core/block/migration"
	"github.com/anyproto/anytype-heart/core/event"
	"github.com/anyproto/anytype-heart/core/files"
	"github.com/anyproto/anytype-heart/core/relation"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	"github.com/anyproto/anytype-heart/pkg/lib/core"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/objectstore"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/space/typeprovider"
	"github.com/anyproto/anytype-heart/util/pbtypes"
	"fmt"
)

type Template struct {
	*Page
}

func NewTemplate(
	sb smartblock.SmartBlock,
	objectStore objectstore.ObjectStore,
	anytype core.Service,
	fileBlockService file.BlockService,
	picker getblock.Picker,
	bookmarkService bookmark.BookmarkService,
	relationService relation.Service,
	tempDirProvider core.TempDirProvider,
	sbtProvider typeprovider.SmartBlockTypeProvider,
	layoutConverter converter.LayoutConverter,
	fileService files.Service,
	eventSender event.Sender,
) *Template {
	return &Template{Page: NewPage(
		sb,
		objectStore,
		anytype,
		fileBlockService,
		picker,
		bookmarkService,
		relationService,
		tempDirProvider,
		sbtProvider,
		layoutConverter,
		fileService,
		eventSender,
	)}
}

func (t *Template) Init(ctx *smartblock.InitContext) (err error) {
	if err = t.Page.Init(ctx); err != nil {
		return
	}

	return
}

func (t *Template) CreationStateMigration(ctx *smartblock.InitContext) migration.Migration {
	parent := t.Page.CreationStateMigration(ctx)

	return migration.Compose(parent, migration.Migration{
		Version: 1,
		Proc: func(s *state.State) {
			// TODO What is fixOt???
			var fixOt bool
			for _, ot := range t.ObjectTypeKeys() {
				if strings.HasPrefix(string(ot), "&") {
					fixOt = true
					break
				}
			}

			if t.Type() == model.SmartBlockType_Template && (len(t.ObjectTypeKeys()) != 2 || fixOt) {
				targetObjectTypeID := pbtypes.GetString(s.Details(), bundle.RelationKeyTargetObjectType.String())
				if targetObjectTypeID != "" {
					targetObjectType, err := t.objectStore.GetObjectType(targetObjectTypeID)
					if err != nil {
						log.Errorf("template createion state: failed to get target object type %s: %s", targetObjectTypeID, err)
						return
					}
					s.SetObjectTypes([]bundle.TypeKey{bundle.TypeKeyTemplate, bundle.TypeKey(targetObjectType.Key)})
				}
			}
		},
	})
}

// GetNewPageState returns state that can be safely used to create the new document
// it has not localDetails set
func (t *Template) GetNewPageState(name string) (st *state.State, err error) {
	st = t.NewState().Copy()
	objectType, err := t.objectStore.GetObjectType(pbtypes.GetString(st.Details(), bundle.RelationKeyTargetObjectType.String()))
	if err != nil {
		return nil, fmt.Errorf("get target object type: %w", err)
	}
	st.SetObjectType(bundle.TypeKey(objectType.Key))
	st.RemoveDetail(bundle.RelationKeyTargetObjectType.String(), bundle.RelationKeyTemplateIsBundled.String())
	// clean-up local details from the template state
	st.SetLocalDetails(nil)

	if name != "" {
		st.SetDetail(bundle.RelationKeyName.String(), pbtypes.String(name))
		if title := st.Get(template.TitleBlockId); title != nil {
			title.Model().GetText().Text = ""
		}
	}
	return
}
