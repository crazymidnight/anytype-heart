/*
Code generated by pkg/lib/bundle/generator. DO NOT EDIT.
source: pkg/lib/bundle/relations.json
*/
package bundle

import "github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"

const RelationChecksum = "43ff13d2e00a3c65dd58a705924ec7734574024f6a15a19ca70e68d21faad45d"

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
	RelationKeyLayoutAlign           RelationKey = "layoutAlign"
	RelationKeyStatus                RelationKey = "status"
	RelationKeyDurationInSeconds     RelationKey = "durationInSeconds"
	RelationKeyIsHidden              RelationKey = "isHidden"
	RelationKeyAperture              RelationKey = "aperture"
	RelationKeyLastModifiedDate      RelationKey = "lastModifiedDate"
	RelationKeyRecommendedRelations  RelationKey = "recommendedRelations"
	RelationKeyCreator               RelationKey = "creator"
	RelationKeyRecommendedLayout     RelationKey = "recommendedLayout"
	RelationKeyLastOpenedDate        RelationKey = "lastOpenedDate"
	RelationKeyArtist                RelationKey = "artist"
	RelationKeyDueDate               RelationKey = "dueDate"
	RelationKeyIconEmoji             RelationKey = "iconEmoji"
	RelationKeyCoverType             RelationKey = "coverType"
	RelationKeyCoverY                RelationKey = "coverY"
	RelationKeySizeInBytes           RelationKey = "sizeInBytes"
	RelationKeyCollectionOf          RelationKey = "collectionOf"
	RelationKeyAddedDate             RelationKey = "addedDate"
	RelationKeyAssignee              RelationKey = "assignee"
	RelationKeyExposure              RelationKey = "exposure"
	RelationKeyTemplateName          RelationKey = "templateName"
	RelationKeyTargetObjectType      RelationKey = "targetObjectType"
	RelationKeyAudioGenre            RelationKey = "audioGenre"
	RelationKeyName                  RelationKey = "name"
	RelationKeyFocalRatio            RelationKey = "focalRatio"
	RelationKeyPriority              RelationKey = "priority"
	RelationKeyFileMimeType          RelationKey = "fileMimeType"
	RelationKeyType                  RelationKey = "type"
	RelationKeyRelationFormat        RelationKey = "relationFormat"
	RelationKeyLayout                RelationKey = "layout"
	RelationKeyAudioAlbumTrackNumber RelationKey = "audioAlbumTrackNumber"
	RelationKeyPlaceOfBirth          RelationKey = "placeOfBirth"
	RelationKeyComposer              RelationKey = "composer"
	RelationKeyCoverX                RelationKey = "coverX"
	RelationKeyDescription           RelationKey = "description"
	RelationKeyId                    RelationKey = "id"
	RelationKeyCameraIso             RelationKey = "cameraIso"
	RelationKeyCoverId               RelationKey = "coverId"
	RelationKeyLastModifiedBy        RelationKey = "lastModifiedBy"
	RelationKeyWidthInPixels         RelationKey = "widthInPixels"
	RelationKeySetOf                 RelationKey = "setOf"
	RelationKeyGender                RelationKey = "gender"
	RelationKeyIsArchived            RelationKey = "isArchived"
	RelationKeyFileExt               RelationKey = "fileExt"
	RelationKeyMpAddedToLibrary      RelationKey = "mpAddedToLibrary"
	RelationKeyFeaturedRelations     RelationKey = "featuredRelations"
)

var (
	relations = map[RelationKey]*model.Relation{
		RelationKeyAddedDate: {

			DataSource:  model.Relation_details,
			Description: "Date when the file were added into the anytype",
			Format:      model.RelationFormat_date,
			Key:         "addedDate",
			MaxCount:    1,
			Name:        "Added date",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyAperture: {

			DataSource:  model.Relation_details,
			Description: "",
			Format:      model.RelationFormat_shorttext,
			Key:         "aperture",
			MaxCount:    1,
			Name:        "Camera Aperture",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyArtist: {

			DataSource:  model.Relation_details,
			Description: "Name of artist",
			Format:      model.RelationFormat_shorttext,
			Key:         "artist",
			MaxCount:    1,
			Name:        "Artist",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyAssignee: {

			DataSource:  model.Relation_details,
			Description: "Person who is responsible for this task or object",
			Format:      model.RelationFormat_object,
			Key:         "assignee",
			Name:        "Assignee",
			ObjectTypes: []string{TypePrefix + "profile"},
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyAttachments: {

			DataSource:  model.Relation_details,
			Description: "",
			Format:      model.RelationFormat_file,
			Key:         "attachments",
			Name:        "Attachments",
			ObjectTypes: []string{TypePrefix + "file", TypePrefix + "image", TypePrefix + "video", TypePrefix + "audio"},
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyAudioAlbum: {

			DataSource:  model.Relation_details,
			Description: "Audio record's album name",
			Format:      model.RelationFormat_shorttext,
			Key:         "audioAlbum",
			MaxCount:    1,
			Name:        "Album",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyAudioAlbumTrackNumber: {

			DataSource:  model.Relation_details,
			Description: "Number of the track in the",
			Format:      model.RelationFormat_number,
			Key:         "audioAlbumTrackNumber",
			MaxCount:    1,
			Name:        "Track #",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyAudioGenre: {

			DataSource:  model.Relation_details,
			Description: "Audio record's genre name",
			Format:      model.RelationFormat_shorttext,
			Key:         "audioGenre",
			MaxCount:    1,
			Name:        "Genre",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyCamera: {

			DataSource:  model.Relation_details,
			Description: "Camera used to capture image or video",
			Format:      model.RelationFormat_shorttext,
			Key:         "camera",
			MaxCount:    1,
			Name:        "Camera",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyCameraIso: {

			DataSource:  model.Relation_details,
			Description: "",
			Format:      model.RelationFormat_number,
			Key:         "cameraIso",
			MaxCount:    1,
			Name:        "ISO",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyCollectionOf: {

			DataSource:  model.Relation_details,
			Description: "Point to the object types that can be added to collection. Empty means any object type can be added to the collection",
			Format:      model.RelationFormat_object,
			Key:         "collectionOf",
			Name:        "Collection of",
			ObjectTypes: []string{TypePrefix + "objectType"},
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyComposer: {

			DataSource:  model.Relation_details,
			Description: "",
			Format:      model.RelationFormat_shorttext,
			Key:         "composer",
			MaxCount:    1,
			Name:        "Composer",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyCoverId: {

			DataSource:  model.Relation_details,
			Description: "Can contains image hash, color or prebuild bg id, depends on coverType relation",
			Format:      model.RelationFormat_shorttext,
			Hidden:      true,
			Key:         "coverId",
			MaxCount:    1,
			Name:        "Cover image or color",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyCoverScale: {

			DataSource:  model.Relation_details,
			Description: "",
			Format:      model.RelationFormat_number,
			Hidden:      true,
			Key:         "coverScale",
			MaxCount:    1,
			Name:        "Cover scale",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyCoverType: {

			DataSource:  model.Relation_details,
			Description: "1-image, 2-color, 3-gradient, 4-prebuilt bg image. Value stored in coverId",
			Format:      model.RelationFormat_number,
			Hidden:      true,
			Key:         "coverType",
			MaxCount:    1,
			Name:        "Cover type",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyCoverX: {

			DataSource:  model.Relation_details,
			Description: "Image x offset of the provided image",
			Format:      model.RelationFormat_number,
			Hidden:      true,
			Key:         "coverX",
			MaxCount:    1,
			Name:        "Cover x offset",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyCoverY: {

			DataSource:  model.Relation_details,
			Description: "Image y offset of the provided image",
			Format:      model.RelationFormat_number,
			Hidden:      true,
			Key:         "coverY",
			MaxCount:    1,
			Name:        "Cover y offset",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyCreatedDate: {

			DataSource:  model.Relation_derived,
			Description: "Date when the object was initially created",
			Format:      model.RelationFormat_date,
			Key:         "createdDate",
			MaxCount:    1,
			Name:        "Creation date",
			ReadOnly:    true,
			Scope:       model.Relation_type,
		},
		RelationKeyCreator: {

			DataSource:  model.Relation_derived,
			Description: "Human which created this object",
			Format:      model.RelationFormat_object,
			Hidden:      true,
			Key:         "creator",
			MaxCount:    1,
			Name:        "Created by",
			ObjectTypes: []string{TypePrefix + "profile"},
			ReadOnly:    true,
			Scope:       model.Relation_type,
		},
		RelationKeyDateOfBirth: {

			DataSource:  model.Relation_details,
			Description: "",
			Format:      model.RelationFormat_date,
			Key:         "dateOfBirth",
			MaxCount:    1,
			Name:        "Date of birth",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyDescription: {

			DataSource:  model.Relation_details,
			Description: "",
			Format:      model.RelationFormat_longtext,
			Key:         "description",
			MaxCount:    1,
			Name:        "Description",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyDone: {

			DataSource:  model.Relation_details,
			Description: "Done checkbox used to render action layout. ",
			Format:      model.RelationFormat_checkbox,
			Hidden:      true,
			Key:         "done",
			MaxCount:    1,
			Name:        "Done",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyDueDate: {

			DataSource:  model.Relation_details,
			Description: "",
			Format:      model.RelationFormat_date,
			Key:         "dueDate",
			MaxCount:    1,
			Name:        "Due date",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyDurationInSeconds: {

			DataSource:  model.Relation_details,
			Description: "Duration of audio/video file in seconds",
			Format:      model.RelationFormat_number,
			Key:         "durationInSeconds",
			MaxCount:    1,
			Name:        "Duration",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyExposure: {

			DataSource:  model.Relation_details,
			Description: "",
			Format:      model.RelationFormat_shorttext,
			Key:         "exposure",
			MaxCount:    1,
			Name:        "Camera Exposure",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyFeaturedRelations: {

			DataSource:  model.Relation_details,
			Description: "Important relations that always appear at the top of the object",
			Format:      model.RelationFormat_object,
			Hidden:      true,
			Key:         "featuredRelations",
			Name:        "Featured Relations",
			ObjectTypes: []string{TypePrefix + "relation"},
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyFileExt: {

			DataSource:  model.Relation_derived,
			Description: "",
			Format:      model.RelationFormat_shorttext,
			Key:         "fileExt",
			MaxCount:    1,
			Name:        "File extension",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyFileMimeType: {

			DataSource:  model.Relation_details,
			Description: "Mime type of object",
			Format:      model.RelationFormat_shorttext,
			Hidden:      true,
			Key:         "fileMimeType",
			MaxCount:    1,
			Name:        "Mime type",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyFocalRatio: {

			DataSource:  model.Relation_details,
			Description: "",
			Format:      model.RelationFormat_number,
			Key:         "focalRatio",
			MaxCount:    1,
			Name:        "Focal ratio",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyGender: {

			DataSource:  model.Relation_details,
			Description: "",
			Format:      model.RelationFormat_status,
			Key:         "gender",
			MaxCount:    1,
			Name:        "Gender",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyHeightInPixels: {

			DataSource:  model.Relation_details,
			Description: "Height of image/video in pixels",
			Format:      model.RelationFormat_number,
			Key:         "heightInPixels",
			MaxCount:    1,
			Name:        "Height",
			ReadOnly:    true,
			Scope:       model.Relation_type,
		},
		RelationKeyIconEmoji: {

			DataSource:  model.Relation_details,
			Description: "1 emoji(can contains multiple UTF symbols) used as an icon",
			Format:      model.RelationFormat_emoji,
			Hidden:      true,
			Key:         "iconEmoji",
			MaxCount:    1,
			Name:        "Emoji",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyIconImage: {

			DataSource:  model.Relation_details,
			Description: "Image icon",
			Format:      model.RelationFormat_file,
			Hidden:      true,
			Key:         "iconImage",
			MaxCount:    1,
			Name:        "Image",
			ObjectTypes: []string{TypePrefix + "image"},
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyId: {

			DataSource:  model.Relation_derived,
			Description: "Link to itself. Used in databases",
			Format:      model.RelationFormat_object,
			Hidden:      true,
			Key:         "id",
			MaxCount:    1,
			Name:        "Anytype ID",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyIsArchived: {

			DataSource:  model.Relation_account,
			Description: "Hides the object",
			Format:      model.RelationFormat_checkbox,
			Hidden:      true,
			Key:         "isArchived",
			MaxCount:    1,
			Name:        "Archived",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyIsHidden: {

			DataSource:  model.Relation_details,
			Description: "Specify if object is hidden",
			Format:      model.RelationFormat_checkbox,
			Hidden:      true,
			Key:         "isHidden",
			MaxCount:    1,
			Name:        "Hidden",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyLastModifiedBy: {

			DataSource:  model.Relation_derived,
			Description: "Human which updates the object last time",
			Format:      model.RelationFormat_object,
			Key:         "lastModifiedBy",
			MaxCount:    1,
			Name:        "Last modified by",
			ObjectTypes: []string{TypePrefix + "profile"},
			ReadOnly:    true,
			Scope:       model.Relation_type,
		},
		RelationKeyLastModifiedDate: {

			DataSource:  model.Relation_derived,
			Description: "Date when the object was modified last time",
			Format:      model.RelationFormat_date,
			Key:         "lastModifiedDate",
			MaxCount:    1,
			Name:        "Last modified date",
			ReadOnly:    true,
			Scope:       model.Relation_type,
		},
		RelationKeyLastOpenedDate: {

			DataSource:  model.Relation_account,
			Description: "Date when the object was modified last opened",
			Format:      model.RelationFormat_date,
			Key:         "lastOpenedDate",
			MaxCount:    1,
			Name:        "Last opened date",
			ReadOnly:    true,
			Scope:       model.Relation_type,
		},
		RelationKeyLayout: {

			DataSource:  model.Relation_details,
			Description: "Anytype layout ID(from pb enum)",
			Format:      model.RelationFormat_number,
			Hidden:      true,
			Key:         "layout",
			MaxCount:    1,
			Name:        "Layout",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyLayoutAlign: {

			DataSource:  model.Relation_details,
			Description: "Specify visual align of the layout",
			Format:      model.RelationFormat_number,
			Hidden:      true,
			Key:         "layoutAlign",
			MaxCount:    1,
			Name:        "Layout align",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyLinkedProjects: {

			DataSource:  model.Relation_details,
			Description: "",
			Format:      model.RelationFormat_object,
			Key:         "linkedProjects",
			Name:        "Linked Projects",
			ObjectTypes: []string{TypePrefix + "project"},
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyLinkedTasks: {

			DataSource:  model.Relation_details,
			Description: "",
			Format:      model.RelationFormat_object,
			Key:         "linkedTasks",
			Name:        "Linked tasks",
			ObjectTypes: []string{TypePrefix + "task"},
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyMpAddedToLibrary: {

			DataSource:  model.Relation_account,
			Description: "Have been added to library from marketplace",
			Format:      model.RelationFormat_checkbox,
			Hidden:      true,
			Key:         "mpAddedToLibrary",
			MaxCount:    1,
			Name:        "Added to library",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyName: {

			DataSource:  model.Relation_details,
			Description: "Name of the object",
			Format:      model.RelationFormat_shorttext,
			Key:         "name",
			MaxCount:    1,
			Name:        "Name",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyPlaceOfBirth: {

			DataSource:  model.Relation_details,
			Description: "",
			Format:      model.RelationFormat_shorttext,
			Key:         "placeOfBirth",
			MaxCount:    1,
			Name:        "Place of birth",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyPriority: {

			DataSource:  model.Relation_details,
			Description: "Used to order tasks in list/canban",
			Format:      model.RelationFormat_number,
			Key:         "priority",
			MaxCount:    1,
			Name:        "Priority",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyRecommendedLayout: {

			DataSource:  model.Relation_details,
			Description: "Recommended layout for new templates and objects of specific objec",
			Format:      model.RelationFormat_number,
			Hidden:      true,
			Key:         "recommendedLayout",
			MaxCount:    1,
			Name:        "Recommended layout",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyRecommendedRelations: {

			DataSource:  model.Relation_details,
			Description: "List of recommended relations",
			Format:      model.RelationFormat_object,
			Hidden:      true,
			Key:         "recommendedRelations",
			Name:        "Recommended relations",
			ObjectTypes: []string{TypePrefix + "relation"},
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyRelationFormat: {

			DataSource:  model.Relation_details,
			Description: "Type of the underlying value",
			Format:      model.RelationFormat_number,
			Hidden:      true,
			Key:         "relationFormat",
			MaxCount:    1,
			Name:        "Relation Format",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyReleasedYear: {

			DataSource:  model.Relation_details,
			Description: "Year when this object were released",
			Format:      model.RelationFormat_number,
			Key:         "releasedYear",
			MaxCount:    1,
			Name:        "Released year",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeySetOf: {

			DataSource:  model.Relation_details,
			Description: "Point to the object types used to aggregate the set. Empty means object of all types will be aggregated ",
			Format:      model.RelationFormat_object,
			Key:         "setOf",
			Name:        "Set of",
			ObjectTypes: []string{TypePrefix + "objectType"},
			ReadOnly:    true,
			Scope:       model.Relation_type,
		},
		RelationKeySizeInBytes: {

			DataSource:  model.Relation_details,
			Description: "Size of file/image in bytes",
			Format:      model.RelationFormat_number,
			Key:         "sizeInBytes",
			MaxCount:    1,
			Name:        "Size",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyStatus: {

			DataSource:  model.Relation_details,
			Description: "Task status?",
			Format:      model.RelationFormat_status,
			Key:         "status",
			MaxCount:    1,
			Name:        "Status",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyTag: {

			DataSource:  model.Relation_details,
			Description: "",
			Format:      model.RelationFormat_tag,
			Key:         "tag",
			Name:        "Tag",
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyTargetObjectType: {

			DataSource:  model.Relation_details,
			Description: "Type that is used for templating",
			Format:      model.RelationFormat_object,
			Hidden:      true,
			Key:         "targetObjectType",
			MaxCount:    1,
			Name:        "Template's Type",
			ObjectTypes: []string{TypePrefix + "objectType"},
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyTemplateName: {

			DataSource:  model.Relation_details,
			Description: "Type indicated the name of template",
			Format:      model.RelationFormat_object,
			Hidden:      true,
			Key:         "templateName",
			MaxCount:    1,
			Name:        "Template name",
			ObjectTypes: []string{TypePrefix + "objectType"},
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyThumbnailImage: {

			DataSource:  model.Relation_details,
			Description: "",
			Format:      model.RelationFormat_file,
			Key:         "thumbnailImage",
			Name:        "Thumbnail image",
			ObjectTypes: []string{TypePrefix + "image"},
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyToBeDeletedDate: {

			DataSource:  model.Relation_account,
			Description: "Date when the object will be deleted from your device",
			Format:      model.RelationFormat_date,
			Hidden:      true,
			Key:         "toBeDeletedDate",
			MaxCount:    1,
			Name:        "Date to delete",
			ReadOnly:    true,
			Scope:       model.Relation_type,
		},
		RelationKeyType: {

			DataSource:  model.Relation_derived,
			Description: "",
			Format:      model.RelationFormat_object,
			Hidden:      true,
			Key:         "type",
			MaxCount:    1,
			Name:        "Object type",
			ObjectTypes: []string{TypePrefix + "objectType"},
			ReadOnly:    false,
			Scope:       model.Relation_type,
		},
		RelationKeyWidthInPixels: {

			DataSource:  model.Relation_details,
			Description: "Width of image/video in pixels",
			Format:      model.RelationFormat_number,
			Key:         "widthInPixels",
			MaxCount:    1,
			Name:        "Width",
			ReadOnly:    true,
			Scope:       model.Relation_type,
		},
	}
)
