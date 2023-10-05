/*
Code generated by pkg/lib/bundle/generator. DO NOT EDIT.
source: pkg/lib/bundle/systemTypes.json
*/
package bundle

import domain "github.com/anyproto/anytype-heart/core/domain"

const SystemTypesChecksum = "4855568ff0cdde0fbd2e8cc227a2a0ec4b87f97472f51f6676237740d56b3cf9"

// SystemTypes contains types that have some special biz logic depends on them in some objects
// they shouldn't be removed or edited in any way
var SystemTypes = append(InternalTypes, []domain.TypeKey{
	TypeKeyPage,
	TypeKeyNote,
	TypeKeyTask,
	TypeKeyCollection,
	TypeKeySet,
	TypeKeyProfile,
	TypeKeyBookmark,
}...)
