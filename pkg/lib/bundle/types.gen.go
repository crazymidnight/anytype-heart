/*
Code generated by pkg/lib/bundle/generator. DO NOT EDIT.
source: pkg/lib/bundle/types.json
*/
package bundle

import (
	"github.com/anyproto/anytype-heart/core/domain"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
)

const TypeChecksum = "e166a24cf6f7afa3be1882069fc2df72200d75eb151b423f4d8f364433102344"
const (
	TypePrefix = "_ot"
)
const (
	TypeKeyRecipe         domain.TypeKey = "recipe"
	TypeKeyNote           domain.TypeKey = "note"
	TypeKeyContact        domain.TypeKey = "contact"
	TypeKeyBookmark       domain.TypeKey = "bookmark"
	TypeKeyDate           domain.TypeKey = "date"
	TypeKeyIdea           domain.TypeKey = "idea"
	TypeKeyTask           domain.TypeKey = "task"
	TypeKeyRelation       domain.TypeKey = "relation"
	TypeKeyBook           domain.TypeKey = "book"
	TypeKeyVideo          domain.TypeKey = "video"
	TypeKeyDashboard      domain.TypeKey = "dashboard"
	TypeKeyDailyPlan      domain.TypeKey = "dailyPlan"
	TypeKeyMovie          domain.TypeKey = "movie"
	TypeKeyObjectType     domain.TypeKey = "objectType"
	TypeKeyRelationOption domain.TypeKey = "relationOption"
	TypeKeySpace          domain.TypeKey = "space"
	TypeKeySpaceView      domain.TypeKey = "spaceView"
	TypeKeyParticipant    domain.TypeKey = "participant"
	TypeKeyTemplate       domain.TypeKey = "template"
	TypeKeySet            domain.TypeKey = "set"
	TypeKeyCollection     domain.TypeKey = "collection"
	TypeKeyDiaryEntry     domain.TypeKey = "diaryEntry"
	TypeKeyPage           domain.TypeKey = "page"
	TypeKeyImage          domain.TypeKey = "image"
	TypeKeyProfile        domain.TypeKey = "profile"
	TypeKeyAudio          domain.TypeKey = "audio"
	TypeKeyGoal           domain.TypeKey = "goal"
	TypeKeyFile           domain.TypeKey = "file"
	TypeKeyProject        domain.TypeKey = "project"
)

var (
	types = map[domain.TypeKey]*model.ObjectType{
		TypeKeyAudio: {

			Description:            "Sound when recorded, with ability to reproduce",
			IconEmoji:              "🎵",
			Layout:                 model.ObjectType_file,
			Name:                   "Audio",
			Readonly:               true,
			RelationLinks:          []*model.RelationLink{MustGetRelationLink(RelationKeyArtist), MustGetRelationLink(RelationKeyAudioAlbum), MustGetRelationLink(RelationKeyAudioAlbumTrackNumber), MustGetRelationLink(RelationKeyAudioGenre), MustGetRelationLink(RelationKeyAudioLyrics), MustGetRelationLink(RelationKeyReleasedYear), MustGetRelationLink(RelationKeySizeInBytes), MustGetRelationLink(RelationKeyFileMimeType), MustGetRelationLink(RelationKeyAddedDate), MustGetRelationLink(RelationKeyFileExt), MustGetRelationLink(RelationKeyOrigin)},
			RestrictObjectCreation: true,
			Types:                  []model.SmartBlockType{model.SmartBlockType_File},
			Url:                    TypePrefix + "audio",
		},
		TypeKeyBook: {

			Description:   "A book is a medium for recording information in the form of writing or images, typically composed of many pages bound together and protected by a cover",
			IconEmoji:     "📘",
			Layout:        model.ObjectType_basic,
			Name:          "Book",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyAuthor), MustGetRelationLink(RelationKeyStatus), MustGetRelationLink(RelationKeyStarred), MustGetRelationLink(RelationKeyUrl)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "book",
		},
		TypeKeyBookmark: {

			Description:   "URL that is stored as Object and may be categorised and linked with objects",
			IconEmoji:     "🔖",
			Layout:        model.ObjectType_bookmark,
			Name:          "Bookmark",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeySource), MustGetRelationLink(RelationKeyPicture)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "bookmark",
		},
		TypeKeyCollection: {

			Description:   "Collect objects in one place, use different views to organize them",
			IconEmoji:     "🗂️",
			Layout:        model.ObjectType_collection,
			Name:          "Collection",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "collection",
		},
		TypeKeyContact: {

			Description:   "Information to make action of communicating or meeting with Human or Company",
			IconEmoji:     "📇",
			Layout:        model.ObjectType_profile,
			Name:          "Contact",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyPhone), MustGetRelationLink(RelationKeyEmail), MustGetRelationLink(RelationKeyCompany), MustGetRelationLink(RelationKeySocialProfile), MustGetRelationLink(RelationKeyOccupation)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "contact",
		},
		TypeKeyDailyPlan: {

			Description:   "A detailed proposal for doing or achieving something for the day\n",
			IconEmoji:     "📆",
			Layout:        model.ObjectType_todo,
			Name:          "Daily Plan",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyTasks)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "dailyPlan",
		},
		TypeKeyDashboard: {

			Description:            "Internal home dashboard with favourite objects",
			Hidden:                 true,
			Layout:                 model.ObjectType_dashboard,
			Name:                   "Dashboard",
			Readonly:               true,
			RestrictObjectCreation: true,
			Types:                  []model.SmartBlockType{model.SmartBlockType_Home},
			Url:                    TypePrefix + "dashboard",
		},
		TypeKeyDate: {

			Description:   "Gregorian calendar date",
			Hidden:        true,
			IconEmoji:     "📅",
			Layout:        model.ObjectType_basic,
			Name:          "Date",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Date},
			Url:           TypePrefix + "date",
		},
		TypeKeyDiaryEntry: {

			Description:   "Record of events and experiences",
			IconEmoji:     "✨",
			Layout:        model.ObjectType_basic,
			Name:          "Diary Entry",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyMood)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "diaryEntry",
		},
		TypeKeyFile: {

			Description:            "Computer resource for recording data in a computer storage device",
			IconEmoji:              "📎",
			Layout:                 model.ObjectType_file,
			Name:                   "File",
			Readonly:               true,
			RelationLinks:          []*model.RelationLink{MustGetRelationLink(RelationKeyFileMimeType), MustGetRelationLink(RelationKeySizeInBytes), MustGetRelationLink(RelationKeyAddedDate), MustGetRelationLink(RelationKeyFileExt), MustGetRelationLink(RelationKeyOrigin)},
			RestrictObjectCreation: true,
			Types:                  []model.SmartBlockType{model.SmartBlockType_File},
			Url:                    TypePrefix + "file",
		},
		TypeKeyGoal: {

			Description:   "The object of a person's ambition or effort; an aim or desired result",
			IconEmoji:     "🎯",
			Layout:        model.ObjectType_todo,
			Name:          "Goal",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyProgress), MustGetRelationLink(RelationKeyStatus), MustGetRelationLink(RelationKeyDueDate), MustGetRelationLink(RelationKeyTasks)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "goal",
		},
		TypeKeyIdea: {

			Description:   "A thought or suggestion as to a possible course of action",
			IconEmoji:     "💡",
			Layout:        model.ObjectType_basic,
			Name:          "Idea",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "idea",
		},
		TypeKeyImage: {

			Description:            "A representation of the external form of a person or thing in art",
			IconEmoji:              "🏞",
			Layout:                 model.ObjectType_image,
			Name:                   "Image",
			Readonly:               true,
			RelationLinks:          []*model.RelationLink{MustGetRelationLink(RelationKeyFileMimeType), MustGetRelationLink(RelationKeyWidthInPixels), MustGetRelationLink(RelationKeyCamera), MustGetRelationLink(RelationKeyHeightInPixels), MustGetRelationLink(RelationKeySizeInBytes), MustGetRelationLink(RelationKeyCameraIso), MustGetRelationLink(RelationKeyAperture), MustGetRelationLink(RelationKeyExposure), MustGetRelationLink(RelationKeyAddedDate), MustGetRelationLink(RelationKeyFocalRatio), MustGetRelationLink(RelationKeyFileExt), MustGetRelationLink(RelationKeyOrigin)},
			RestrictObjectCreation: true,
			Types:                  []model.SmartBlockType{model.SmartBlockType_File},
			Url:                    TypePrefix + "image",
		},
		TypeKeyMovie: {

			Description:   "Motion picture or Moving picture, is a work of visual art used to simulate experiences that communicate ideas, stories, perceptions, feelings, beauty, or atmosphere through the use of moving images",
			IconEmoji:     "🍿",
			Layout:        model.ObjectType_basic,
			Name:          "Movie",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyDirector), MustGetRelationLink(RelationKeyGenre), MustGetRelationLink(RelationKeyStatus)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "movie",
		},
		TypeKeyNote: {

			Description:   "Blank canvas with no Title. A brief record of points written down as an aid to memory",
			IconEmoji:     "📝",
			Layout:        model.ObjectType_note,
			Name:          "Note",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "note",
		},
		TypeKeyObjectType: {

			Description:   "Object that contains a definition of some object type",
			IconEmoji:     "🥚",
			Layout:        model.ObjectType_objectType,
			Name:          "Type",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyRecommendedRelations), MustGetRelationLink(RelationKeyRecommendedLayout)},
			Types:         []model.SmartBlockType{model.SmartBlockType_SubObject, model.SmartBlockType_BundledObjectType},
			Url:           TypePrefix + "objectType",
		},
		TypeKeyPage: {

			Description:   "Blank canvas with Title",
			IconEmoji:     "📄",
			Layout:        model.ObjectType_basic,
			Name:          "Page",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "page",
		},
		TypeKeyParticipant: {

			Description:            "Anytype identity as a space participant",
			IconEmoji:              "🧑",
			Layout:                 model.ObjectType_participant,
			Name:                   "Space member",
			Readonly:               true,
			RelationLinks:          []*model.RelationLink{MustGetRelationLink(RelationKeyTag)},
			RestrictObjectCreation: true,
			Revision:               1,
			Types:                  []model.SmartBlockType{model.SmartBlockType_Participant},
			Url:                    TypePrefix + "participant",
		},
		TypeKeyProfile: {

			Description:   "Homo sapiens",
			IconEmoji:     "🧍",
			Layout:        model.ObjectType_profile,
			Name:          "Human",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page, model.SmartBlockType_ProfilePage},
			Url:           TypePrefix + "profile",
		},
		TypeKeyProject: {

			Description:   "An individual or collaborative enterprise that is carefully planned to achieve a particular aim",
			IconEmoji:     "🔨",
			Layout:        model.ObjectType_basic,
			Name:          "Project",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyTasks)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "project",
		},
		TypeKeyRecipe: {

			Description:   "A recipe is a set of instructions that describes how to prepare or make something, especially a dish of prepared food",
			IconEmoji:     "🍲",
			Layout:        model.ObjectType_basic,
			Name:          "Recipe",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyTime), MustGetRelationLink(RelationKeyIngredients)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "recipe",
		},
		TypeKeyRelation: {

			Description:   "Meaningful connection between objects",
			Hidden:        true,
			IconEmoji:     "🔗",
			Layout:        model.ObjectType_relation,
			Name:          "Relation",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyRelationFormat), MustGetRelationLink(RelationKeyRelationMaxCount), MustGetRelationLink(RelationKeyRelationDefaultValue), MustGetRelationLink(RelationKeyRelationFormatObjectTypes)},
			Types:         []model.SmartBlockType{model.SmartBlockType_SubObject, model.SmartBlockType_BundledRelation},
			Url:           TypePrefix + "relation",
		},
		TypeKeyRelationOption: {

			Description: "Object that contains a relation option",
			Hidden:      true,
			IconEmoji:   "🥚",
			Layout:      model.ObjectType_relationOption,
			Name:        "Relation option",
			Readonly:    true,
			Types:       []model.SmartBlockType{model.SmartBlockType_SubObject},
			Url:         TypePrefix + "relationOption",
		},
		TypeKeySet: {

			Description:   "Query all objects in your space based on types and relations",
			IconEmoji:     "🔎",
			Layout:        model.ObjectType_set,
			Name:          "Set",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeySetOf)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "set",
		},
		TypeKeySpace: {

			Description:            "Workspace",
			Hidden:                 true,
			IconEmoji:              "🌎",
			Layout:                 model.ObjectType_space,
			Name:                   "Space",
			Readonly:               true,
			RelationLinks:          []*model.RelationLink{MustGetRelationLink(RelationKeyTag)},
			RestrictObjectCreation: true,
			Types:                  []model.SmartBlockType{model.SmartBlockType_Workspace},
			Url:                    TypePrefix + "space",
		},
		TypeKeySpaceView: {

			Description:            "Space",
			Hidden:                 true,
			IconEmoji:              "🌎",
			Layout:                 model.ObjectType_spaceView,
			Name:                   "Space",
			Readonly:               true,
			RelationLinks:          []*model.RelationLink{MustGetRelationLink(RelationKeyTag)},
			RestrictObjectCreation: true,
			Types:                  []model.SmartBlockType{model.SmartBlockType_SpaceView},
			Url:                    TypePrefix + "spaceView",
		},
		TypeKeyTask: {

			Description:   "A piece of work to be done or undertaken",
			IconEmoji:     "✅",
			Layout:        model.ObjectType_todo,
			Name:          "Task",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyAssignee), MustGetRelationLink(RelationKeyDueDate), MustGetRelationLink(RelationKeyStatus), MustGetRelationLink(RelationKeyDone), MustGetRelationLink(RelationKeyPriority), MustGetRelationLink(RelationKeyTasks), MustGetRelationLink(RelationKeyLinkedProjects)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "task",
		},
		TypeKeyTemplate: {

			Description:   "Sample object that has already some details in place and used to create objects from",
			Hidden:        true,
			Layout:        model.ObjectType_basic,
			Name:          "Template",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTargetObjectType), MustGetRelationLink(RelationKeyTemplateIsBundled)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Template},
			Url:           TypePrefix + "template",
		},
		TypeKeyVideo: {

			Description:            "The recording of moving visual images",
			IconEmoji:              "📽",
			Layout:                 model.ObjectType_file,
			Name:                   "Video",
			Readonly:               true,
			RelationLinks:          []*model.RelationLink{MustGetRelationLink(RelationKeySizeInBytes), MustGetRelationLink(RelationKeyFileMimeType), MustGetRelationLink(RelationKeyCamera), MustGetRelationLink(RelationKeyHeightInPixels), MustGetRelationLink(RelationKeyWidthInPixels), MustGetRelationLink(RelationKeyCameraIso), MustGetRelationLink(RelationKeyAperture), MustGetRelationLink(RelationKeyExposure), MustGetRelationLink(RelationKeyAddedDate), MustGetRelationLink(RelationKeyFileExt), MustGetRelationLink(RelationKeyOrigin)},
			RestrictObjectCreation: true,
			Types:                  []model.SmartBlockType{model.SmartBlockType_File},
			Url:                    TypePrefix + "video",
		},
	}
)
