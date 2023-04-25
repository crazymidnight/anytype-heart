/*
Code generated by pkg/lib/bundle/generator. DO NOT EDIT.
source: pkg/lib/bundle/internalRelations.json
*/
package bundle

const InternalRelationsChecksum = "3c4d85cc0506984e7c1e7441ce0e02bd54c6b520944fc2760f1ee3f6bb800cb9"

// RequiredInternalRelations contains internal relations that will be added to EVERY new or existing object
// if this relation only needs SPECIFIC objects(e.g. of some type) add it to the SystemRelations
var RequiredInternalRelations = []RelationKey{
	RelationKeyId,
	RelationKeyName,
	RelationKeyDescription,
	RelationKeySnippet,
	RelationKeyIconEmoji,
	RelationKeyIconImage,
	RelationKeyType,
	RelationKeyLayout,
	RelationKeyLayoutAlign,
	RelationKeyCoverId,
	RelationKeyCoverScale,
	RelationKeyCoverType,
	RelationKeyCoverX,
	RelationKeyCoverY,
	RelationKeyCreatedDate,
	RelationKeyCreator,
	RelationKeyLastModifiedDate,
	RelationKeyLastModifiedBy,
	RelationKeyLastOpenedDate,
	RelationKeyFeaturedRelations,
	RelationKeyIsFavorite,
	RelationKeyWorkspaceId,
	RelationKeyLinks,
	RelationKeyInternalFlags,
	RelationKeyRestrictions,
}
