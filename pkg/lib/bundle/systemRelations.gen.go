/*
Code generated by pkg/lib/bundle/generator. DO NOT EDIT.
source: pkg/lib/bundle/systemRelations.json
*/
package bundle

import domain "github.com/anyproto/anytype-heart/core/domain"

const SystemRelationsChecksum = "6d3b58dd79958eb6fe47a13bed541f0ecea32fb918ff0ede3d6ecd2bd1ee0009"

// SystemRelations contains relations that have some special biz logic depends on them in some objects
// in case EVERY object depend on the relation please add it to RequiredInternalRelations
var SystemRelations = append(RequiredInternalRelations, []domain.RelationKey{
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
	RelationKeySpaceShareableStatus,
	RelationKeyIsHiddenDiscovery,
	RelationKeyDone,
	RelationKeyIsArchived,
	RelationKeyTemplateIsBundled,
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
	RelationKeyWidthInPixels,
	RelationKeyHeightInPixels,
	RelationKeyFileExt,
	RelationKeySizeInBytes,
	RelationKeySourceFilePath,
	RelationKeyFileSyncStatus,
	RelationKeyDefaultTemplateId,
	RelationKeyUniqueKey,
	RelationKeyBacklinks,
	RelationKeyProfileOwnerIdentity,
	RelationKeyFileBackupStatus,
	RelationKeyFileId,
	RelationKeyFileIndexingStatus,
	RelationKeyOrigin,
	RelationKeyRevision,
	RelationKeyImageKind,
	RelationKeyImportType,
	RelationKeySpaceAccessType,
	RelationKeySpaceInviteFileCid,
	RelationKeySpaceInviteFileKey,
	RelationKeyReadersLimit,
	RelationKeyWritersLimit,
	RelationKeyParticipantPermissions,
	RelationKeyParticipantStatus,
	RelationKeyLatestAclHeadId,
	RelationKeyIdentity,
	RelationKeyGlobalName,
}...)
