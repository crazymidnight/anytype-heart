/*
Code generated by pkg/lib/bundle/generator. DO NOT EDIT.
source: pkg/lib/bundle/types.json
*/
package bundle

import "github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/relation"

type TypeKey string

func (tk TypeKey) String() string {
	return string(tk)
}
func (tk TypeKey) URL() string {
	return string(TypePrefix + tk)
}

const (
	TypePrefix = "https://anytype.io/schemas/object/bundled/"
)
const (
	TypeKeyNote       TypeKey = "note"
	TypeKeyDashboard  TypeKey = "dashboard"
	TypeKeyContact    TypeKey = "contact"
	TypeKeyIdea       TypeKey = "idea"
	TypeKeyTask       TypeKey = "task"
	TypeKeyRelation   TypeKey = "relation"
	TypeKeyVideo      TypeKey = "video"
	TypeKeyObjectType TypeKey = "objectType"
	TypeKeySet        TypeKey = "set"
	TypeKeyPage       TypeKey = "page"
	TypeKeyImage      TypeKey = "image"
	TypeKeyProfile    TypeKey = "profile"
	TypeKeyAudio      TypeKey = "audio"
	TypeKeyDocument   TypeKey = "document"
	TypeKeyFile       TypeKey = "file"
	TypeKeyProject    TypeKey = "project"
	TypeKeyCollection TypeKey = "collection"
)

var (
	types = map[TypeKey]*relation.ObjectType{
		TypeKeyAudio: {

			Layout:    relation.ObjectType_basic,
			Name:      "Audio",
			Relations: []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverImage], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyArtist], relations[RelationKeyAudioAlbum], relations[RelationKeyAudioAlbumTrackNumber], relations[RelationKeyAudioGenre], relations[RelationKeyReleasedYear], relations[RelationKeyThumbnailImage], relations[RelationKeyComposer], relations[RelationKeyDurationInSeconds], relations[RelationKeySizeInBytes], relations[RelationKeyFileMimeType]},
			Url:       TypePrefix + "audio",
		},
		TypeKeyCollection: {

			Layout:    relation.ObjectType_database,
			Name:      "Collection",
			Relations: []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverImage], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyCollectionOf]},
			Url:       TypePrefix + "collection",
		},
		TypeKeyContact: {

			Layout:    relation.ObjectType_profile,
			Name:      "Contact",
			Relations: []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverImage], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate]},
			Url:       TypePrefix + "contact",
		},
		TypeKeyDashboard: {

			Layout:    relation.ObjectType_dashboard,
			Name:      "Dashboard",
			Relations: []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverImage], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate]},
			Url:       TypePrefix + "dashboard",
		},
		TypeKeyDocument: {

			Layout:    relation.ObjectType_basic,
			Name:      "Document",
			Relations: []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverImage], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate]},
			Url:       TypePrefix + "document",
		},
		TypeKeyFile: {

			Layout:    relation.ObjectType_basic,
			Name:      "File",
			Relations: []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverImage], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFileMimeType], relations[RelationKeySizeInBytes]},
			Url:       TypePrefix + "file",
		},
		TypeKeyIdea: {

			Layout:    relation.ObjectType_basic,
			Name:      "Idea",
			Relations: []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverImage], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyTag]},
			Url:       TypePrefix + "idea",
		},
		TypeKeyImage: {

			Layout:    relation.ObjectType_basic,
			Name:      "Image",
			Relations: []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverImage], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFileMimeType], relations[RelationKeyWidthInPixels], relations[RelationKeyCamera], relations[RelationKeyHeightInPixels], relations[RelationKeySizeInBytes], relations[RelationKeyCameraIso], relations[RelationKeyAperture], relations[RelationKeyExposure]},
			Url:       TypePrefix + "image",
		},
		TypeKeyNote: {

			Layout:    relation.ObjectType_basic,
			Name:      "Note",
			Relations: []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverImage], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate]},
			Url:       TypePrefix + "note",
		},
		TypeKeyObjectType: {

			Layout:    relation.ObjectType_objectType,
			Name:      "Type",
			Relations: []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverImage], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyRecommendedRelations]},
			Url:       TypePrefix + "objectType",
		},
		TypeKeyPage: {

			Layout:    relation.ObjectType_basic,
			Name:      "Undefined",
			Relations: []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverImage], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate]},
			Url:       TypePrefix + "page",
		},
		TypeKeyProfile: {

			Layout:    relation.ObjectType_profile,
			Name:      "Human",
			Relations: []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverImage], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate]},
			Url:       TypePrefix + "profile",
		},
		TypeKeyProject: {

			Layout:    relation.ObjectType_basic,
			Name:      "Project",
			Relations: []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverImage], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate]},
			Url:       TypePrefix + "project",
		},
		TypeKeyRelation: {

			Layout:    relation.ObjectType_relation,
			Name:      "Relation",
			Relations: []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyLayout]},
			Url:       TypePrefix + "relation",
		},
		TypeKeySet: {

			Layout:    relation.ObjectType_set,
			Name:      "Set of objects",
			Relations: []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverImage], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeySetOf]},
			Url:       TypePrefix + "set",
		},
		TypeKeyTask: {

			Layout:    relation.ObjectType_action,
			Name:      "Task",
			Relations: []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverImage], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyAssignee], relations[RelationKeyDueDate], relations[RelationKeyDescription], relations[RelationKeyAttachments], relations[RelationKeyStatus], relations[RelationKeyDone], relations[RelationKeyPriority], relations[RelationKeyLinkedTasks], relations[RelationKeyLinkedProjects], relations[RelationKeyTag]},
			Url:       TypePrefix + "task",
		},
		TypeKeyVideo: {

			Layout:    relation.ObjectType_basic,
			Name:      "Video",
			Relations: []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverImage], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyDurationInSeconds], relations[RelationKeySizeInBytes], relations[RelationKeyFileMimeType], relations[RelationKeyCamera], relations[RelationKeyThumbnailImage], relations[RelationKeyHeightInPixels], relations[RelationKeyWidthInPixels], relations[RelationKeyCameraIso], relations[RelationKeyAperture], relations[RelationKeyExposure]},
			Url:       TypePrefix + "video",
		},
	}
)
