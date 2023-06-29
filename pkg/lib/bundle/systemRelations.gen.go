/*
Code generated by pkg/lib/bundle/generator. DO NOT EDIT.
source: pkg/lib/bundle/systemRelations.json
*/
package bundle

const SystemRelationsChecksum = "c3e623682ba3262ab08447ae73bdba7a3cc35463c28c5f15546106d0e256aa50"

// SystemRelations contains relations that have some special biz logic depends on them in some objects
// in case EVERY object depend on the relation please add it to RequiredInternalRelations
var SystemRelations = append(RequiredInternalRelations, []RelationKey{
	RelationKeyAddedDate,
	RelationKeySource,
	RelationKeySourceObject,
	RelationKeySetOf,
	RelationKeyRelationFormat,
	RelationKeyRelationKey,
	RelationKeyRelationReadonlyValue,
	RelationKeyRelationDefaultValue,
	RelationKeyRelationMaxCount,
	RelationKeyRelationOptionColor,
	RelationKeyRelationFormatObjectTypes,
	RelationKeyIsReadonly,
	RelationKeyIsDeleted,
	RelationKeyIsHidden,
	RelationKeyDone,
	RelationKeyIsArchived,
	RelationKeyTemplateIsBundled,
	RelationKeyTag,
	RelationKeySmartblockTypes,
	RelationKeyTargetObjectType,
	RelationKeyRecommendedLayout,
	RelationKeyFileExt,
	RelationKeyFileMimeType,
	RelationKeySizeInBytes,
	RelationKeyOldAnytypeID,
	RelationKeySpaceDashboardId,
	RelationKeyRecommendedRelations,
	RelationKeyIconOption,
	RelationKeySpaceAccessibility,
	RelationKeyWidthInPixels,
	RelationKeyHeightInPixels,
	RelationKeyFileExt,
	RelationKeySizeInBytes,
	RelationKeySourceFilePath,
	RelationKeyFileSyncStatus,
}...)
