/*
Code generated by pkg/lib/bundle/generator. DO NOT EDIT.
source: pkg/lib/bundle/types.json
*/
package bundle

import "github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"

const TypeChecksum = "beffb54845886c177f4cf308a5e9672fae5c464652c9a711635dfa07c79aba5e"

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
	TypeKeyDailyReflection TypeKey = "dailyReflection"
	TypeKeyRecipe          TypeKey = "recipe"
	TypeKeyNote            TypeKey = "note"
	TypeKeyResume          TypeKey = "resume"
	TypeKeyContact         TypeKey = "contact"
	TypeKeyBookmark        TypeKey = "bookmark"
	TypeKeyWeeklyPlan      TypeKey = "weeklyPlan"
	TypeKeyDate            TypeKey = "date"
	TypeKeyInvoice         TypeKey = "invoice"
	TypeKeyIdea            TypeKey = "idea"
	TypeKeyTask            TypeKey = "task"
	TypeKeyRelation        TypeKey = "relation"
	TypeKeyBook            TypeKey = "book"
	TypeKeyVideo           TypeKey = "video"
	TypeKeyCompany         TypeKey = "company"
	TypeKeyDashboard       TypeKey = "dashboard"
	TypeKeyDailyPlan       TypeKey = "dailyPlan"
	TypeKeyMeetingNote     TypeKey = "meetingNote"
	TypeKeyMovie           TypeKey = "movie"
	TypeKeyArticle         TypeKey = "article"
	TypeKeyObjectType      TypeKey = "objectType"
	TypeKeyRelationOption  TypeKey = "relationOption"
	TypeKeySpace           TypeKey = "space"
	TypeKeyTemplate        TypeKey = "template"
	TypeKeyHabitTrack      TypeKey = "habitTrack"
	TypeKeySet             TypeKey = "set"
	TypeKeyClassNote       TypeKey = "classNote"
	TypeKeyDiaryEntry      TypeKey = "diaryEntry"
	TypeKeyPage            TypeKey = "page"
	TypeKeyImage           TypeKey = "image"
	TypeKeyBug             TypeKey = "bug"
	TypeKeyProfile         TypeKey = "profile"
	TypeKeyAudio           TypeKey = "audio"
	TypeKeyActionPlan      TypeKey = "actionPlan"
	TypeKeyGoal            TypeKey = "goal"
	TypeKeyFeature         TypeKey = "feature"
	TypeKeyDocument        TypeKey = "document"
	TypeKeyFile            TypeKey = "file"
	TypeKeyProject         TypeKey = "project"
)

var (
	types = map[TypeKey]*model.ObjectType{
		TypeKeyActionPlan: {

			Description:   "Is a detailed plan outlining actions needed to reach one or more goals",
			IconEmoji:     "🤝",
			Layout:        model.ObjectType_todo,
			Name:          "Action Plan",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyTasks), MustGetRelationLink(RelationKeyResources), MustGetRelationLink(RelationKeyResult), MustGetRelationLink(RelationKeyDueDate), MustGetRelationLink(RelationKeyResponsible)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "actionPlan",
		},
		TypeKeyArticle: {

			Description:   "A piece of writing included with others in a newspaper, magazine, or other publication",
			IconEmoji:     "📰",
			Layout:        model.ObjectType_basic,
			Name:          "Article",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "article",
		},
		TypeKeyAudio: {

			Description:   "Auto-generated object from .wav, .mp3, .ogg files added to Anytype. Sound when recorded, with ability to reproduce",
			IconEmoji:     "🎵",
			Layout:        model.ObjectType_file,
			Name:          "Audio",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyArtist), MustGetRelationLink(RelationKeyAudioAlbum), MustGetRelationLink(RelationKeyAudioAlbumTrackNumber), MustGetRelationLink(RelationKeyAudioGenre), MustGetRelationLink(RelationKeyReleasedYear), MustGetRelationLink(RelationKeyThumbnailImage), MustGetRelationLink(RelationKeyComposer), MustGetRelationLink(RelationKeyDurationInSeconds), MustGetRelationLink(RelationKeySizeInBytes), MustGetRelationLink(RelationKeyFileMimeType), MustGetRelationLink(RelationKeyAddedDate), MustGetRelationLink(RelationKeyFileExt), MustGetRelationLink(RelationKeyAudioArtist), MustGetRelationLink(RelationKeyAudioLyrics)},
			Types:         []model.SmartBlockType{model.SmartBlockType_File},
			Url:           TypePrefix + "audio",
		},
		TypeKeyBook: {

			Description:   "A book is a medium for recording information in the form of writing or images, typically composed of many pages bound together and protected by a cover",
			IconEmoji:     "📘",
			Layout:        model.ObjectType_basic,
			Name:          "Book",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyAuthor), MustGetRelationLink(RelationKeyCategory), MustGetRelationLink(RelationKeyRating), MustGetRelationLink(RelationKeyStatus)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "book",
		},
		TypeKeyBookmark: {

			Description:   "URL that is stored as Object and may be categorised and linked with objects",
			IconEmoji:     "🔖",
			Layout:        model.ObjectType_bookmark,
			Name:          "Bookmark",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeySource), MustGetRelationLink(RelationKeyPicture), MustGetRelationLink(RelationKeyNotes), MustGetRelationLink(RelationKeyQuote)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "bookmark",
		},
		TypeKeyBug: {

			Description:   "An error, fault or flaw in any computer program or a hardware system. A bug produces unexpected results or causes a system to behave unexpectedly",
			IconEmoji:     "🐞",
			Layout:        model.ObjectType_todo,
			Name:          "Bug (Software)",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyHowToReproduce), MustGetRelationLink(RelationKeyResult), MustGetRelationLink(RelationKeyAdditional), MustGetRelationLink(RelationKeyAttachments), MustGetRelationLink(RelationKeyAssignee), MustGetRelationLink(RelationKeyDueDate), MustGetRelationLink(RelationKeyPriority)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "bug",
		},
		TypeKeyClassNote: {

			Description:   "Note for recording lectures or seminars",
			IconEmoji:     "👨🏻\u200d🏫",
			Layout:        model.ObjectType_basic,
			Name:          "Class Note",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyClass), MustGetRelationLink(RelationKeyClassType), MustGetRelationLink(RelationKeyRecords), MustGetRelationLink(RelationKeyQuestions), MustGetRelationLink(RelationKeyMaterials), MustGetRelationLink(RelationKeyTasks), MustGetRelationLink(RelationKeyReflection)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "classNote",
		},
		TypeKeyCompany: {

			Description:   "A company, abbreviated as co., is a legal entity representing an association of people, whether natural, legal or a mixture of both, with a specific objective",
			IconEmoji:     "🏢",
			Layout:        model.ObjectType_profile,
			Name:          "Company",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyStory), MustGetRelationLink(RelationKeyFounded), MustGetRelationLink(RelationKeyCeo), MustGetRelationLink(RelationKeyFounders), MustGetRelationLink(RelationKeyOwner), MustGetRelationLink(RelationKeyNumberOfEmployees), MustGetRelationLink(RelationKeyHeadquarters), MustGetRelationLink(RelationKeyWebsite), MustGetRelationLink(RelationKeySocialProfile), MustGetRelationLink(RelationKeyStockprice), MustGetRelationLink(RelationKeyTickerSymbol), MustGetRelationLink(RelationKeyAddress), MustGetRelationLink(RelationKeySubsidiaries)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "company",
		},
		TypeKeyContact: {

			Description:   "Information to make action of communicating or meeting with Human or Company",
			IconEmoji:     "📇",
			Layout:        model.ObjectType_profile,
			Name:          "Contact",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyPhone), MustGetRelationLink(RelationKeyEmail), MustGetRelationLink(RelationKeyDateOfBirth), MustGetRelationLink(RelationKeyPlaceOfBirth), MustGetRelationLink(RelationKeyCompany), MustGetRelationLink(RelationKeySocialProfile), MustGetRelationLink(RelationKeyJob), MustGetRelationLink(RelationKeyLinkedContacts), MustGetRelationLink(RelationKeyOccupation), MustGetRelationLink(RelationKeyInstagram), MustGetRelationLink(RelationKeyGender), MustGetRelationLink(RelationKeyFacebook)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "contact",
		},
		TypeKeyDailyPlan: {

			Description:   "A detailed proposal for doing or achieving something for the day\n",
			IconEmoji:     "📆",
			Layout:        model.ObjectType_todo,
			Name:          "Daily Plan",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyNotes), MustGetRelationLink(RelationKeyTasks), MustGetRelationLink(RelationKeyEvents)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "dailyPlan",
		},
		TypeKeyDailyReflection: {

			Description:   "Serious thought or consideration",
			IconEmoji:     "💭",
			Layout:        model.ObjectType_basic,
			Name:          "Daily Reflection",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyIntentions), MustGetRelationLink(RelationKeyHappenings), MustGetRelationLink(RelationKeyGratefulFor), MustGetRelationLink(RelationKeyMood), MustGetRelationLink(RelationKeyWorriedAbout), MustGetRelationLink(RelationKeyTasks)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "dailyReflection",
		},
		TypeKeyDashboard: {

			Description: "Internal home dashboard with favourite objects",
			Hidden:      true,
			Layout:      model.ObjectType_dashboard,
			Name:        "Dashboard",
			Readonly:    true,
			Types:       []model.SmartBlockType{model.SmartBlockType_Home},
			Url:         TypePrefix + "dashboard",
		},
		TypeKeyDate: {

			Description:   "Gregorian calendar date",
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
		TypeKeyDocument: {

			Description:   "A piece of matter that provides information or evidence or that serves as an official record",
			IconEmoji:     "📃",
			Layout:        model.ObjectType_basic,
			Name:          "Document",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "document",
		},
		TypeKeyFeature: {

			Description:   "A distinguishing characteristic of a software item (e.g., performance, portability, or functionality)",
			IconEmoji:     "🪁",
			Layout:        model.ObjectType_todo,
			Name:          "Feature",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyHypothesisAssumptions), MustGetRelationLink(RelationKeyProblem), MustGetRelationLink(RelationKeyUserStories), MustGetRelationLink(RelationKeyLogic), MustGetRelationLink(RelationKeyMeasureOfSuccess), MustGetRelationLink(RelationKeyAttachments), MustGetRelationLink(RelationKeyAssignee), MustGetRelationLink(RelationKeyDueDate), MustGetRelationLink(RelationKeyPriority)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "feature",
		},
		TypeKeyFile: {

			Description:   "Auto-generated object from files added to Anytype. Computer resource for recording data in a computer storage device",
			IconEmoji:     "🗂️",
			Layout:        model.ObjectType_file,
			Name:          "File",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyFileMimeType), MustGetRelationLink(RelationKeySizeInBytes), MustGetRelationLink(RelationKeyAddedDate), MustGetRelationLink(RelationKeyFileExt)},
			Types:         []model.SmartBlockType{model.SmartBlockType_File},
			Url:           TypePrefix + "file",
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
		TypeKeyHabitTrack: {

			Description:   "A habit track is a simple way to measure whether you did a habit",
			IconEmoji:     "🥕",
			Layout:        model.ObjectType_todo,
			Name:          "Habit Track",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyRunning), MustGetRelationLink(RelationKeyMeditation), MustGetRelationLink(RelationKey7hourssleep), MustGetRelationLink(RelationKeyJournaling), MustGetRelationLink(RelationKeyHealthyEating)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "habitTrack",
		},
		TypeKeyIdea: {

			Description:   "A thought or suggestion as to a possible course of action",
			IconEmoji:     "💡",
			Layout:        model.ObjectType_basic,
			Name:          "Idea",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyProblem), MustGetRelationLink(RelationKeySolution), MustGetRelationLink(RelationKeyAlternative)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "idea",
		},
		TypeKeyImage: {

			Description:   "Auto-generated object from .jpg & .png files added to Anytype. A representation of the external form of a person or thing in art",
			IconEmoji:     "🏞",
			Layout:        model.ObjectType_image,
			Name:          "Image",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyFileMimeType), MustGetRelationLink(RelationKeyWidthInPixels), MustGetRelationLink(RelationKeyCamera), MustGetRelationLink(RelationKeyHeightInPixels), MustGetRelationLink(RelationKeySizeInBytes), MustGetRelationLink(RelationKeyCameraIso), MustGetRelationLink(RelationKeyAperture), MustGetRelationLink(RelationKeyExposure), MustGetRelationLink(RelationKeyAddedDate), MustGetRelationLink(RelationKeyFocalRatio), MustGetRelationLink(RelationKeyFileExt)},
			Types:         []model.SmartBlockType{model.SmartBlockType_File},
			Url:           TypePrefix + "image",
		},
		TypeKeyInvoice: {

			Description:   "A list of goods sent or services provided, with a statement of the sum due for these; a bill",
			IconEmoji:     "🧾",
			Layout:        model.ObjectType_todo,
			Name:          "Invoice",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyCompany), MustGetRelationLink(RelationKeyAddress), MustGetRelationLink(RelationKeyBillTo), MustGetRelationLink(RelationKeyBillToAddress), MustGetRelationLink(RelationKeyShipTo), MustGetRelationLink(RelationKeyShipToAddress), MustGetRelationLink(RelationKeyDueDate), MustGetRelationLink(RelationKeyTotal), MustGetRelationLink(RelationKeyNumber)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "invoice",
		},
		TypeKeyMeetingNote: {

			Description:   "Quick references to ideas, goals, deadlines, data, and anything else important that's covered in your meeting",
			IconEmoji:     "✏️",
			Layout:        model.ObjectType_basic,
			Name:          "Meeting Note",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyParticipants), MustGetRelationLink(RelationKeyAgenda), MustGetRelationLink(RelationKeyTasks)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "meetingNote",
		},
		TypeKeyMovie: {

			Description:   "Motion picture or Moving picture, is a work of visual art used to simulate experiences that communicate ideas, stories, perceptions, feelings, beauty, or atmosphere through the use of moving images",
			IconEmoji:     "🍿",
			Layout:        model.ObjectType_basic,
			Name:          "Movie",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyDirector), MustGetRelationLink(RelationKeyStars), MustGetRelationLink(RelationKeyGenre), MustGetRelationLink(RelationKeyTrailer), MustGetRelationLink(RelationKeyRating), MustGetRelationLink(RelationKeyImdbRating), MustGetRelationLink(RelationKeyRottenTomatoesRating), MustGetRelationLink(RelationKeyStatus)},
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
			Types:         []model.SmartBlockType{model.SmartBlockType_STObjectType, model.SmartBlockType_BundledObjectType},
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
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyObjectives), MustGetRelationLink(RelationKeyScope), MustGetRelationLink(RelationKeyTimeframe), MustGetRelationLink(RelationKeyBudget), MustGetRelationLink(RelationKeyStakeholders), MustGetRelationLink(RelationKeyTasks)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "project",
		},
		TypeKeyRecipe: {

			Description:   "A recipe is a set of instructions that describes how to prepare or make something, especially a dish of prepared food",
			IconEmoji:     "🍲",
			Layout:        model.ObjectType_basic,
			Name:          "Recipe",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyTime), MustGetRelationLink(RelationKeyServings), MustGetRelationLink(RelationKeyIngredients), MustGetRelationLink(RelationKeyInstructions), MustGetRelationLink(RelationKeyDifficulty)},
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
		TypeKeyResume: {

			Description:   "A resume is a formal document that a job applicant creates to itemize his or her qualifications for a position",
			IconEmoji:     "👋",
			Layout:        model.ObjectType_profile,
			Name:          "Resume",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyCandidate), MustGetRelationLink(RelationKeyJob), MustGetRelationLink(RelationKeyPhone), MustGetRelationLink(RelationKeyEmail), MustGetRelationLink(RelationKeyLocation), MustGetRelationLink(RelationKeyWebsite), MustGetRelationLink(RelationKeySocialProfile)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "resume",
		},
		TypeKeySet: {

			Description:   "Collection of objects with equal types and relations. Database experience based on all objects in Anytype",
			IconEmoji:     "🗂️",
			Layout:        model.ObjectType_set,
			Name:          "Set",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeySetOf)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Set},
			Url:           TypePrefix + "set",
		},
		TypeKeySpace: {

			Description:   "Space for sharing",
			Hidden:        true,
			IconEmoji:     "🌎",
			Layout:        model.ObjectType_space,
			Name:          "Space",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Workspace},
			Url:           TypePrefix + "space",
		},
		TypeKeyTask: {

			Description:   "A piece of work to be done or undertaken",
			IconEmoji:     "✅",
			Layout:        model.ObjectType_todo,
			Name:          "Task",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyAssignee), MustGetRelationLink(RelationKeyDueDate), MustGetRelationLink(RelationKeyAttachments), MustGetRelationLink(RelationKeyStatus), MustGetRelationLink(RelationKeyDone), MustGetRelationLink(RelationKeyPriority), MustGetRelationLink(RelationKeyTasks), MustGetRelationLink(RelationKeyLinkedProjects)},
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

			Description:   "Auto-generated object from .mpeg-4 files added to Anytype. The recording of moving visual images",
			IconEmoji:     "📽",
			Layout:        model.ObjectType_file,
			Name:          "Video",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyDurationInSeconds), MustGetRelationLink(RelationKeySizeInBytes), MustGetRelationLink(RelationKeyFileMimeType), MustGetRelationLink(RelationKeyCamera), MustGetRelationLink(RelationKeyThumbnailImage), MustGetRelationLink(RelationKeyHeightInPixels), MustGetRelationLink(RelationKeyWidthInPixels), MustGetRelationLink(RelationKeyCameraIso), MustGetRelationLink(RelationKeyAperture), MustGetRelationLink(RelationKeyExposure), MustGetRelationLink(RelationKeyAddedDate), MustGetRelationLink(RelationKeyFileExt)},
			Types:         []model.SmartBlockType{model.SmartBlockType_File},
			Url:           TypePrefix + "video",
		},
		TypeKeyWeeklyPlan: {

			Description:   "The act of organizing your activities and tasks for the week",
			IconEmoji:     "🗓️",
			Layout:        model.ObjectType_todo,
			Name:          "Weekly Plan",
			Readonly:      true,
			RelationLinks: []*model.RelationLink{MustGetRelationLink(RelationKeyTag), MustGetRelationLink(RelationKeyNotes), MustGetRelationLink(RelationKeyEvents), MustGetRelationLink(RelationKeyTasks)},
			Types:         []model.SmartBlockType{model.SmartBlockType_Page},
			Url:           TypePrefix + "weeklyPlan",
		},
	}
)
