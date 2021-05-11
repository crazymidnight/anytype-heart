/*
Code generated by pkg/lib/bundle/generator. DO NOT EDIT.
source: pkg/lib/bundle/types.json
*/
package bundle

import "github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"

const TypeChecksum = "b3a0c2a45eb4299bfbee55bd5f89f4d3420ce9e5847e8d2b85562cee42092451"

type TypeKey string

func (tk TypeKey) String() string {
	return string(tk)
}
func (tk TypeKey) URL() string {
	return string(TypePrefix + tk)
}

const (
	TypePrefix = "_ot"
)
const (
	TypeKeyNote       TypeKey = "note"
	TypeKeyContact    TypeKey = "contact"
	TypeKeyIdea       TypeKey = "idea"
	TypeKeyTask       TypeKey = "task"
	TypeKeyRelation   TypeKey = "relation"
	TypeKeyVideo      TypeKey = "video"
	TypeKeyDashboard  TypeKey = "dashboard"
	TypeKeyObjectType TypeKey = "objectType"
	TypeKeyTemplate   TypeKey = "template"
	TypeKeySet        TypeKey = "set"
	TypeKeyPage       TypeKey = "page"
	TypeKeyImage      TypeKey = "image"
	TypeKeyProfile    TypeKey = "profile"
	TypeKeyAudio      TypeKey = "audio"
	TypeKeyDocument   TypeKey = "document"
	TypeKeyFile       TypeKey = "file"
	TypeKeyProject    TypeKey = "project"
)

var (
	types = map[TypeKey]*model.ObjectType{
		TypeKeyAudio: {

			Description: "",
			Hidden:      true,
			IconEmoji:   "🎵",
			Layout:      model.ObjectType_basic,
			Name:        "Audio",
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyArtist], relations[RelationKeyAudioAlbum], relations[RelationKeyAudioAlbumTrackNumber], relations[RelationKeyAudioGenre], relations[RelationKeyReleasedYear], relations[RelationKeyThumbnailImage], relations[RelationKeyComposer], relations[RelationKeyDurationInSeconds], relations[RelationKeySizeInBytes], relations[RelationKeyFileMimeType], relations[RelationKeyAddedDate], relations[RelationKeyFileExt]},
			Types:       []model.SmartBlockType{model.SmartBlockType_File},
			Url:         TypePrefix + "audio",
		},
		TypeKeyContact: {

			Description: "",
			IconEmoji:   "📇",
			Layout:      model.ObjectType_profile,
			Name:        "Contact",
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTag]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Page},
			Url:         TypePrefix + "contact",
		},
		TypeKeyDashboard: {

			Description: "Internal home dashboard",
			Hidden:      true,
			Layout:      model.ObjectType_dashboard,
			Name:        "Dashboard",
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Home},
			Url:         TypePrefix + "dashboard",
		},
		TypeKeyDocument: {

			Description: "",
			IconEmoji:   "📋",
			Layout:      model.ObjectType_basic,
			Name:        "Document",
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTag]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Page},
			Url:         TypePrefix + "document",
		},
		TypeKeyFile: {

			Description: "",
			IconEmoji:   "🗂️",
			Layout:      model.ObjectType_basic,
			Name:        "File",
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyFileMimeType], relations[RelationKeySizeInBytes], relations[RelationKeyAddedDate], relations[RelationKeyFileExt]},
			Types:       []model.SmartBlockType{model.SmartBlockType_File},
			Url:         TypePrefix + "file",
		},
		TypeKeyIdea: {

			Description: "",
			IconEmoji:   "💡",
			Layout:      model.ObjectType_basic,
			Name:        "Idea",
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTag]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Page},
			Url:         TypePrefix + "idea",
		},
		TypeKeyImage: {

			Description: "",
			IconEmoji:   "🌅",
			Layout:      model.ObjectType_image,
			Name:        "Image",
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyFileMimeType], relations[RelationKeyWidthInPixels], relations[RelationKeyCamera], relations[RelationKeyHeightInPixels], relations[RelationKeySizeInBytes], relations[RelationKeyCameraIso], relations[RelationKeyAperture], relations[RelationKeyExposure], relations[RelationKeyAddedDate], relations[RelationKeyFocalRatio], relations[RelationKeyFileExt]},
			Types:       []model.SmartBlockType{model.SmartBlockType_File},
			Url:         TypePrefix + "image",
		},
		TypeKeyNote: {

			Description: "",
			IconEmoji:   "🗒️",
			Layout:      model.ObjectType_basic,
			Name:        "Note",
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTag]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Page},
			Url:         TypePrefix + "note",
		},
		TypeKeyObjectType: {

			Description: "Object that contains a definition of some object type",
			Hidden:      true,
			IconEmoji:   "🔮",
			Layout:      model.ObjectType_objectType,
			Name:        "Type",
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyRecommendedRelations], relations[RelationKeyRecommendedLayout], relations[RelationKeyMpAddedToLibrary], relations[RelationKeyIsHidden]},
			Types:       []model.SmartBlockType{model.SmartBlockType_STObjectType, model.SmartBlockType_BundledObjectType},
			Url:         TypePrefix + "objectType",
		},
		TypeKeyPage: {

			Description: "Base type to start with",
			IconEmoji:   "📄",
			Layout:      model.ObjectType_basic,
			Name:        "Page",
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTag]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Page},
			Url:         TypePrefix + "page",
		},
		TypeKeyProfile: {

			Description: "",
			IconEmoji:   "🧍",
			Layout:      model.ObjectType_profile,
			Name:        "Human",
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Page, model.SmartBlockType_ProfilePage},
			Url:         TypePrefix + "profile",
		},
		TypeKeyProject: {

			Description: "",
			IconEmoji:   "🔨",
			Layout:      model.ObjectType_basic,
			Name:        "Project",
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTag]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Page},
			Url:         TypePrefix + "project",
		},
		TypeKeyRelation: {

			Description: "",
			Hidden:      true,
			IconEmoji:   "🔗",
			Layout:      model.ObjectType_relation,
			Name:        "Relation",
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyLayout], relations[RelationKeyDescription], relations[RelationKeyCreator], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyMpAddedToLibrary], relations[RelationKeyRelationFormat], relations[RelationKeyIsHidden]},
			Types:       []model.SmartBlockType{model.SmartBlockType_IndexedRelation, model.SmartBlockType_BundledRelation},
			Url:         TypePrefix + "relation",
		},
		TypeKeySet: {

			Description: "",
			IconEmoji:   "🗂️",
			Layout:      model.ObjectType_set,
			Name:        "Set",
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTag], relations[RelationKeySetOf]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Set},
			Url:         TypePrefix + "set",
		},
		TypeKeyTask: {

			Description: "",
			IconEmoji:   "✔️",
			Layout:      model.ObjectType_todo,
			Name:        "Task",
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTag], relations[RelationKeyAssignee], relations[RelationKeyDueDate], relations[RelationKeyAttachments], relations[RelationKeyStatus], relations[RelationKeyDone], relations[RelationKeyPriority], relations[RelationKeyLinkedTasks], relations[RelationKeyLinkedProjects]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Page},
			Url:         TypePrefix + "task",
		},
		TypeKeyTemplate: {

			Description: "Special type to create objects from",
			Hidden:      true,
			IconEmoji:   "✨",
			Layout:      model.ObjectType_basic,
			Name:        "Template",
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTargetObjectType], relations[RelationKeyBuiltinTemplateId]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Template},
			Url:         TypePrefix + "template",
		},
		TypeKeyVideo: {

			Description: "",
			IconEmoji:   "📹",
			Layout:      model.ObjectType_basic,
			Name:        "Video",
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyDurationInSeconds], relations[RelationKeySizeInBytes], relations[RelationKeyFileMimeType], relations[RelationKeyCamera], relations[RelationKeyThumbnailImage], relations[RelationKeyHeightInPixels], relations[RelationKeyWidthInPixels], relations[RelationKeyCameraIso], relations[RelationKeyAperture], relations[RelationKeyExposure], relations[RelationKeyAddedDate], relations[RelationKeyFileExt]},
			Types:       []model.SmartBlockType{model.SmartBlockType_File},
			Url:         TypePrefix + "video",
		},
	}
)
