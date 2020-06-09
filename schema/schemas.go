package schema

// Code generated by go generate; DO NOT EDIT.
//go:generate go run embed/embed.go

var SchemaByURL = map[string]string{

	"https://anytype.io/schemas/page": `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "id": "https://anytype.io/schemas/page",
  "title": "Anytype Page",
  "description": "This schema contains the base properties of Anytype Page and should be refereed if you want to extend it",
  "type": "array",
  "items": {
    "$ref": "https://anytype.io/schemas/relation"
  },
  "uniqueItems": true,
  "default": [
    {
      "id": "id",
      "name": "ID",
      "isHided": true,
      "type": "https://anytype.io/schemas/types/page"
    },
    {
      "id": "name",
      "name": "Name",
      "type": "https://anytype.io/schemas/types/title"
    },
    {
      "id": "iconEmoji",
      "name": "Emoji",
      "isHided": true,
      "type": "https://anytype.io/schemas/types/emoji"
    },
    {
      "id": "iconImage",
      "name": "Image",
      "isHided": true,
      "type": "https://anytype.io/schemas/types/image"
    },
    {
      "id": "isArchived",
      "name": "Archived",
      "type": "https://anytype.io/schemas/types/checkbox"
    },
    {
      "id": "coverType",
      "name": "Cover Type",
      "isHided": true,
      "type": "https://anytype.io/schemas/types/coverType"
    },
    {
      "id": "coverId",
      "name": "Predefined ID or Image",
      "isHided": true,
      "type": "https://anytype.io/schemas/types/image"
    },
    {
      "id": "coverX",
      "name": "Cover x offset",
      "isHided": true,
      "type": "https://anytype.io/schemas/types/number"
    },
    {
      "id": "coverY",
      "name": "Cover y offset",
      "isHided": true,
      "type": "https://anytype.io/schemas/types/number"
    },
    {
      "id": "coverScale",
      "name": "Cover scale",
      "isHided": true,
      "type": "https://anytype.io/schemas/types/number"
    }
  ]
}
`,
	"https://anytype.io/schemas/person": `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "id": "https://anytype.io/schemas/person",
  "title": "Anytype Page",
  "description": "This schema contains the base properties of Anytype Person",

  "allOf": [
    { "$ref": "https://anytype.io/schemas/page" }
  ],

  "type": "object"
}
`,
	"https://anytype.io/schemas/relation": `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "id": "https://anytype.io/schemas/relation",
  "title": "Anytype Page",
  "description": "This schema contains relation and all type definitions",
  "type": "object",
  "$comment": "fills relation from specific details",
  "properties": {
    "id": {
      "type": "string",
      "$comment": "detail's ID"
    },
    "name": {
      "type": "string"
    },
    "type": {
      "type": "string",
      "$comment": "json schema $id for the relation type, starting from https://anytype.io/schemas/types/"
    },
    "objectType": {
      "type": "string",
      "$comment": "json schema $id of the object type for relations with the object type, e.g. https://anytype.io/schemas/page"
    },
    "isMulti": {
      "type": "boolean",
      "$comment": "multiple fields of the same type grouped in the array. allowed for: select, image, file, object"
    },
    "isHided": {
      "type": "boolean",
      "$comment": "presented in the dataset, may be rendered with some view types but should be hided in the relations list"
    }
  },
  "definitions": {
    "title": {
      "$id": "https://anytype.io/schemas/types/title",
      "type": "string",
      "description": "Title renders name plus first emoji/image relation for the same relation"
    },
    "description": {
      "$id": "https://anytype.io/schemas/types/description",
      "type": "string"
    },
    "select": {
      "$id": "https://anytype.io/schemas/types/select",
      "type": "string"
    },
    "number": {
      "$id": "https://anytype.io/schemas/types/number",
      "type": "number"
    },
    "url": {
      "$id": "https://anytype.io/schemas/types/url",
      "type": "string",
      "description": "External URL",
      "format": "uri"
    },
    "email": {
      "$id": "https://anytype.io/schemas/types/email",
      "type": "string",
      "format": "email"
    },
    "phone": {
      "$id": "https://anytype.io/schemas/types/phone",
      "type": "string"
    },
    "date": {
      "$id": "https://anytype.io/schemas/types/date",
      "type": "string",
      "description": "UNIX timestamp as a string"
    },
    "checkbox": {
      "$id": "https://anytype.io/schemas/types/checkbox",
      "type": "boolean"
    },
    "object": {
      "$id": "https://anytype.io/schemas/types/object",
      "type": "string",
      "description": "ID of the object, e.g. page or person"
    },
    "image": {
      "$id": "https://anytype.io/schemas/types/image",
      "type": "string",
      "description": "CID of image node in the IPFS"
    },
    "file": {
      "$id": "https://anytype.io/schemas/types/file",
      "type": "string",
      "description": "CID of file node in the IPFS"
    },
    "emoji": {
      "$id": "https://anytype.io/schemas/types/emoji",
      "type": "string",
      "description": "One emoji as unicode"
    },
    "coverType": {
      "$id": "https://anytype.io/schemas/types/coverType",
      "enum": [
        0, "None",
        1, "Image",
        2, "Color",
        3, "Gradient",
        4, "Upload",
        5, "BgImage"
      ],
      "description": "Page cover type"
    }
  }
}
`}
