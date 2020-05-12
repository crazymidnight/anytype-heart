package schema

// Code generated by go generate; DO NOT EDIT.
//go:generate go run embed/embed.go

const (
	Page = `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "id": "https://anytype.io/schemas/page",
  "title": "Anytype Page",
  "description": "This schema contains the base properties of Anytype Page and should be refereed if you want to extend it",

  "definitions": {
    "image": {
      "$id": "https://anytype.io/schemas/image",
      "type": "string",
      "description": "CID of image node in the IPFS"
    },
    "file": {
      "$id": "https://anytype.io/schemas/file",
      "type": "string",
      "description": "CID of file node in the IPFS"
    },
    "emoji": {
      "$id": "https://anytype.io/schemas/emoji",
      "type": "string",
      "description": "Unicode emoji"
    }
  },

  "type": "object",
  "properties": {
    "name": { "type": "string", "default": "Untitled" },
    "iconEmoji": { "$ref": "#/definitions/emoji" },
    "iconImage":  { "$ref": "#/definitions/image" },
    "isArchived":  { "type": "boolean", "comment": "scope:account" },
    "isDeleted":  { "type": "boolean", "comment": "scope:local", "readOnly": true}
  },
  "required": ["name"]
}
`
)
