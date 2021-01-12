/*
Code generated by pkg/lib/bundle/generator. DO NOT EDIT.
source: pkg/lib/bundle/relations.json
*/
package bundle

import "github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/relation"

type RelationKey string

func (rk RelationKey) String() string {
	return string(rk)
}

const (
	RelationKeyTag                   RelationKey = "tag"
	RelationKeyCamera                RelationKey = "camera"
	RelationKeyHeightInPixels        RelationKey = "heightInPixels"
	RelationKeyCreatedDate           RelationKey = "createdDate"
	RelationKeyToBeDeletedDate       RelationKey = "toBeDeletedDate"
	RelationKeyDone                  RelationKey = "done"
	RelationKeyDateOfBirth           RelationKey = "dateOfBirth"
	RelationKeyThumbnailImage        RelationKey = "thumbnailImage"
	RelationKeyAttachments           RelationKey = "attachments"
	RelationKeyLinkedTasks           RelationKey = "linkedTasks"
	RelationKeyIconImage             RelationKey = "iconImage"
	RelationKeyReleasedYear          RelationKey = "releasedYear"
	RelationKeyCoverScale            RelationKey = "coverScale"
	RelationKeyLinkedProjects        RelationKey = "linkedProjects"
	RelationKeyAudioAlbum            RelationKey = "audioAlbum"
	RelationKeyStatus                RelationKey = "status"
	RelationKeyDurationInSeconds     RelationKey = "durationInSeconds"
	RelationKeyAperture              RelationKey = "aperture"
	RelationKeyLastModifiedDate      RelationKey = "lastModifiedDate"
	RelationKeyRecommendedRelations  RelationKey = "recommendedRelations"
	RelationKeyCreator               RelationKey = "creator"
	RelationKeyLastOpenedDate        RelationKey = "lastOpenedDate"
	RelationKeyArtist                RelationKey = "artist"
	RelationKeyDueDate               RelationKey = "dueDate"
	RelationKeyIconEmoji             RelationKey = "iconEmoji"
	RelationKeyCoverY                RelationKey = "coverY"
	RelationKeySizeInBytes           RelationKey = "sizeInBytes"
	RelationKeyCollectionOf          RelationKey = "collectionOf"
	RelationKeyDoneStatus            RelationKey = "doneStatus"
	RelationKeyAssignee              RelationKey = "assignee"
	RelationKeyExposure              RelationKey = "exposure"
	RelationKeyAudioGenre            RelationKey = "audioGenre"
	RelationKeyName                  RelationKey = "name"
	RelationKeyPriority              RelationKey = "priority"
	RelationKeyFileMimeType          RelationKey = "fileMimeType"
	RelationKeyType                  RelationKey = "type"
	RelationKeyAudioAlbumTrackNumber RelationKey = "audioAlbumTrackNumber"
	RelationKeyPlaceOfBirth          RelationKey = "placeOfBirth"
	RelationKeyComposer              RelationKey = "composer"
	RelationKeyCoverX                RelationKey = "coverX"
	RelationKeyDescription           RelationKey = "description"
	RelationKeyId                    RelationKey = "id"
	RelationKeyCameraIso             RelationKey = "cameraIso"
	RelationKeyCoverImage            RelationKey = "coverImage"
	RelationKeyLastModifiedBy        RelationKey = "lastModifiedBy"
	RelationKeyWidthInPixels         RelationKey = "widthInPixels"
	RelationKeySetOf                 RelationKey = "setOf"
	RelationKeyGender                RelationKey = "gender"
	RelationKeyFeaturedRelations     RelationKey = "featuredRelations"
)

var (
	Relations = map[RelationKey]*relation.Relation{
		RelationKeyAperture: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_title,
			Hidden:     false,
			Key:        "aperture",
			Name:       "Camera Aperture",
			ReadOnly:   false,
		},
		RelationKeyArtist: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_title,
			Hidden:     false,
			Key:        "artist",
			Name:       "Artist",
			ReadOnly:   false,
		},
		RelationKeyAssignee: {

			DataSource:  relation.Relation_details,
			Format:      relation.RelationFormat_object,
			Hidden:      false,
			Key:         "assignee",
			Name:        "Assignee",
			ObjectTypes: []string{TypePrefix + "profile"},
			ReadOnly:    false,
		},
		RelationKeyAttachments: {

			DataSource:  relation.Relation_details,
			Format:      relation.RelationFormat_object,
			Hidden:      false,
			Key:         "attachments",
			Name:        "Attachments",
			ObjectTypes: []string{TypePrefix + "file", TypePrefix + "image", TypePrefix + "video", TypePrefix + "audio"},
			ReadOnly:    false,
		},
		RelationKeyAudioAlbum: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_title,
			Hidden:     false,
			Key:        "audioAlbum",
			Name:       "Album",
			ReadOnly:   false,
		},
		RelationKeyAudioAlbumTrackNumber: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_number,
			Hidden:     false,
			Key:        "audioAlbumTrackNumber",
			Name:       "Track #",
			ReadOnly:   false,
		},
		RelationKeyAudioGenre: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_title,
			Hidden:     false,
			Key:        "audioGenre",
			Name:       "Genre",
			ReadOnly:   false,
		},
		RelationKeyCamera: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_title,
			Hidden:     false,
			Key:        "camera",
			Name:       "Camera",
			ReadOnly:   false,
		},
		RelationKeyCameraIso: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_number,
			Hidden:     false,
			Key:        "cameraIso",
			Name:       "ISO",
			ReadOnly:   false,
		},
		RelationKeyCollectionOf: {

			DataSource:  relation.Relation_details,
			Format:      relation.RelationFormat_object,
			Hidden:      false,
			Key:         "collectionOf",
			Name:        "Collection of",
			ObjectTypes: []string{TypePrefix + "objectType"},
			ReadOnly:    false,
		},
		RelationKeyComposer: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_title,
			Hidden:     false,
			Key:        "composer",
			Name:       "Composer",
			ReadOnly:   false,
		},
		RelationKeyCoverImage: {

			DataSource:  relation.Relation_details,
			Format:      relation.RelationFormat_object,
			Hidden:      false,
			Key:         "coverImage",
			Name:        "Cover image",
			ObjectTypes: []string{TypePrefix + "image"},
			ReadOnly:    false,
		},
		RelationKeyCoverScale: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_number,
			Hidden:     true,
			Key:        "coverScale",
			Name:       "Cover scale",
			ReadOnly:   false,
		},
		RelationKeyCoverX: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_number,
			Hidden:     true,
			Key:        "coverX",
			Name:       "Cover x offset",
			ReadOnly:   false,
		},
		RelationKeyCoverY: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_number,
			Hidden:     true,
			Key:        "coverY",
			Name:       "Cover y offset",
			ReadOnly:   false,
		},
		RelationKeyCreatedDate: {

			DataSource: relation.Relation_derived,
			Format:     relation.RelationFormat_date,
			Hidden:     false,
			Key:        "createdDate",
			Name:       "Creation date",
			ReadOnly:   true,
		},
		RelationKeyCreator: {

			DataSource:  relation.Relation_derived,
			Format:      relation.RelationFormat_object,
			Hidden:      false,
			Key:         "creator",
			Name:        "Created by",
			ObjectTypes: []string{TypePrefix + "profile"},
			ReadOnly:    true,
		},
		RelationKeyDateOfBirth: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_date,
			Hidden:     false,
			Key:        "dateOfBirth",
			Name:       "Date of birth",
			ReadOnly:   false,
		},
		RelationKeyDescription: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_description,
			Hidden:     false,
			Key:        "description",
			Name:       "Description",
			ReadOnly:   false,
		},
		RelationKeyDone: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_checkbox,
			Hidden:     false,
			Key:        "done",
			Name:       "Done",
			ReadOnly:   false,
		},
		RelationKeyDoneStatus: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_select,
			Hidden:     false,
			Key:        "doneStatus",
			Name:       "Done status",
			ReadOnly:   false,
		},
		RelationKeyDueDate: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_date,
			Hidden:     false,
			Key:        "dueDate",
			Name:       "Due date",
			ReadOnly:   false,
		},
		RelationKeyDurationInSeconds: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_number,
			Hidden:     false,
			Key:        "durationInSeconds",
			Name:       "Duration(sec)",
			ReadOnly:   false,
		},
		RelationKeyExposure: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_title,
			Hidden:     false,
			Key:        "exposure",
			Name:       "Camera Exposure",
			ReadOnly:   false,
		},
		RelationKeyFeaturedRelations: {

			DataSource:  relation.Relation_details,
			Format:      relation.RelationFormat_object,
			Hidden:      true,
			Key:         "featuredRelations",
			Name:        "Featured relations management will be \u2028implemented later.",
			ObjectTypes: []string{TypePrefix + "relation"},
			ReadOnly:    false,
		},
		RelationKeyFileMimeType: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_title,
			Hidden:     true,
			Key:        "fileMimeType",
			Name:       "Mime type",
			ReadOnly:   false,
		},
		RelationKeyGender: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_select,
			Hidden:     false,
			Key:        "gender",
			Name:       "Gender",
			ReadOnly:   false,
		},
		RelationKeyHeightInPixels: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_number,
			Hidden:     false,
			Key:        "heightInPixels",
			Name:       "Height(px)",
			ReadOnly:   true,
		},
		RelationKeyIconEmoji: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_emoji,
			Hidden:     false,
			Key:        "iconEmoji",
			Name:       "Emoji",
			ReadOnly:   false,
		},
		RelationKeyIconImage: {

			DataSource:  relation.Relation_details,
			Format:      relation.RelationFormat_object,
			Hidden:      false,
			Key:         "iconImage",
			Name:        "Image",
			ObjectTypes: []string{TypePrefix + "image"},
			ReadOnly:    false,
		},
		RelationKeyId: {

			DataSource: relation.Relation_derived,
			Format:     relation.RelationFormat_object,
			Hidden:     true,
			Key:        "id",
			Name:       "Anytype ID",
			ReadOnly:   false,
		},
		RelationKeyLastModifiedBy: {

			DataSource:  relation.Relation_derived,
			Format:      relation.RelationFormat_object,
			Hidden:      false,
			Key:         "lastModifiedBy",
			Name:        "Last modified by",
			ObjectTypes: []string{TypePrefix + "profile"},
			ReadOnly:    true,
		},
		RelationKeyLastModifiedDate: {

			DataSource: relation.Relation_derived,
			Format:     relation.RelationFormat_date,
			Hidden:     false,
			Key:        "lastModifiedDate",
			Name:       "Last modified date",
			ReadOnly:   true,
		},
		RelationKeyLastOpenedDate: {

			DataSource: relation.Relation_account,
			Format:     relation.RelationFormat_date,
			Hidden:     false,
			Key:        "lastOpenedDate",
			Name:       "Last opened date",
			ReadOnly:   true,
		},
		RelationKeyLinkedProjects: {

			DataSource:  relation.Relation_details,
			Format:      relation.RelationFormat_object,
			Hidden:      false,
			Key:         "linkedProjects",
			Name:        "Linked Projects",
			ObjectTypes: []string{TypePrefix + "project"},
			ReadOnly:    false,
		},
		RelationKeyLinkedTasks: {

			DataSource:  relation.Relation_details,
			Format:      relation.RelationFormat_object,
			Hidden:      false,
			Key:         "linkedTasks",
			Name:        "Linked tasks",
			ObjectTypes: []string{TypePrefix + "task"},
			ReadOnly:    false,
		},
		RelationKeyName: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_title,
			Hidden:     false,
			Key:        "name",
			Name:       "Name",
			ReadOnly:   false,
		},
		RelationKeyPlaceOfBirth: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_select,
			Hidden:     false,
			Key:        "placeOfBirth",
			Name:       "Place of birth",
			ReadOnly:   false,
		},
		RelationKeyPriority: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_number,
			Hidden:     true,
			Key:        "priority",
			Name:       "Priority",
			ReadOnly:   false,
		},
		RelationKeyRecommendedRelations: {

			DataSource:  relation.Relation_details,
			Format:      relation.RelationFormat_object,
			Hidden:      false,
			Key:         "recommendedRelations",
			Name:        "Recommended relations",
			ObjectTypes: []string{TypePrefix + "relation"},
			ReadOnly:    false,
		},
		RelationKeyReleasedYear: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_number,
			Hidden:     false,
			Key:        "releasedYear",
			Name:       "Released year",
			ReadOnly:   false,
		},
		RelationKeySetOf: {

			DataSource:  relation.Relation_details,
			Format:      relation.RelationFormat_object,
			Hidden:      false,
			Key:         "setOf",
			Name:        "Set of",
			ObjectTypes: []string{TypePrefix + "objectType"},
			ReadOnly:    false,
		},
		RelationKeySizeInBytes: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_number,
			Hidden:     false,
			Key:        "sizeInBytes",
			Name:       "Size(bytes)",
			ReadOnly:   false,
		},
		RelationKeyStatus: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_select,
			Hidden:     false,
			Key:        "status",
			Name:       "Status",
			ReadOnly:   false,
		},
		RelationKeyTag: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_select,
			Hidden:     false,
			Key:        "tag",
			Name:       "Tag",
			ReadOnly:   false,
		},
		RelationKeyThumbnailImage: {

			DataSource:  relation.Relation_details,
			Format:      relation.RelationFormat_object,
			Hidden:      false,
			Key:         "thumbnailImage",
			Name:        "Thumbnail image",
			ObjectTypes: []string{TypePrefix + "image"},
			ReadOnly:    false,
		},
		RelationKeyToBeDeletedDate: {

			DataSource: relation.Relation_account,
			Format:     relation.RelationFormat_date,
			Hidden:     true,
			Key:        "toBeDeletedDate",
			Name:       "Date to delete",
			ReadOnly:   true,
		},
		RelationKeyType: {

			DataSource:  relation.Relation_details,
			Format:      relation.RelationFormat_object,
			Hidden:      false,
			Key:         "type",
			Name:        "Object type",
			ObjectTypes: []string{TypePrefix + "objectType"},
			ReadOnly:    false,
		},
		RelationKeyWidthInPixels: {

			DataSource: relation.Relation_details,
			Format:     relation.RelationFormat_number,
			Hidden:     false,
			Key:        "widthInPixels",
			Name:       "Width(px)",
			ReadOnly:   true,
		},
	}
)
