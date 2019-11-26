package block

import "github.com/gogo/protobuf/types"

func findPosInSlice(s []string, v string) int {
	for i, sv := range s {
		if sv == v {
			return i
		}
	}
	return -1
}

func insertToSlice(s []string, v string, pos int) []string {
	if len(s) <= pos {
		return append(s, v)
	}
	if pos == 0 {
		return append([]string{v}, s[pos:]...)
	}
	return append(s[:pos], append([]string{v}, s[pos:]...)...)
}

func fieldsGetString(field *types.Struct, key string) (value string, ok bool) {
	if field != nil && field.Fields != nil {
		if value, ok := field.Fields[key]; ok {
			if s, ok := value.Kind.(*types.Value_StringValue); ok {
				return s.StringValue, true
			}
		}
	}
	return
}
