// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/lib/pb/relation/protos/relation.proto

package relation

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// RelationFormat describes how the underlying data is stored in the google.protobuf.Value and how it should be validated/sanitized
type RelationFormat int32

const (
	RelationFormat_shortText   RelationFormat = 0
	RelationFormat_longText    RelationFormat = 1
	RelationFormat_number      RelationFormat = 2
	RelationFormat_select      RelationFormat = 3
	RelationFormat_multiSelect RelationFormat = 4
	RelationFormat_date        RelationFormat = 5
	RelationFormat_file        RelationFormat = 6
	RelationFormat_checkbox    RelationFormat = 7
	RelationFormat_url         RelationFormat = 8
	RelationFormat_email       RelationFormat = 9
	RelationFormat_phone       RelationFormat = 10
	RelationFormat_emoji       RelationFormat = 11
	RelationFormat_objectId    RelationFormat = 100
)

var RelationFormat_name = map[int32]string{
	0:   "shortText",
	1:   "longText",
	2:   "number",
	3:   "select",
	4:   "multiSelect",
	5:   "date",
	6:   "file",
	7:   "checkbox",
	8:   "url",
	9:   "email",
	10:  "phone",
	11:  "emoji",
	100: "objectId",
}

var RelationFormat_value = map[string]int32{
	"shortText":   0,
	"longText":    1,
	"number":      2,
	"select":      3,
	"multiSelect": 4,
	"date":        5,
	"file":        6,
	"checkbox":    7,
	"url":         8,
	"email":       9,
	"phone":       10,
	"emoji":       11,
	"objectId":    100,
}

func (x RelationFormat) String() string {
	return proto.EnumName(RelationFormat_name, int32(x))
}

func (RelationFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d7f8dc729d740ffd, []int{0}
}

type RelationRelationDataSource int32

const (
	Relation_details RelationRelationDataSource = 0
	Relation_local   RelationRelationDataSource = 1
)

var RelationRelationDataSource_name = map[int32]string{
	0: "details",
	1: "local",
}

var RelationRelationDataSource_value = map[string]int32{
	"details": 0,
	"local":   1,
}

func (x RelationRelationDataSource) String() string {
	return proto.EnumName(RelationRelationDataSource_name, int32(x))
}

func (RelationRelationDataSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d7f8dc729d740ffd, []int{2, 0}
}

type ObjectType struct {
	Name      string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Relations []*Relation `protobuf:"bytes,2,rep,name=relations,proto3" json:"relations,omitempty"`
}

func (m *ObjectType) Reset()         { *m = ObjectType{} }
func (m *ObjectType) String() string { return proto.CompactTextString(m) }
func (*ObjectType) ProtoMessage()    {}
func (*ObjectType) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7f8dc729d740ffd, []int{0}
}
func (m *ObjectType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ObjectType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ObjectType.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ObjectType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectType.Merge(m, src)
}
func (m *ObjectType) XXX_Size() int {
	return m.Size()
}
func (m *ObjectType) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectType.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectType proto.InternalMessageInfo

func (m *ObjectType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ObjectType) GetRelations() []*Relation {
	if m != nil {
		return m.Relations
	}
	return nil
}

type RelationWithValue struct {
	Relation *Relation    `protobuf:"bytes,1,opt,name=relation,proto3" json:"relation,omitempty"`
	Value    *types.Value `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *RelationWithValue) Reset()         { *m = RelationWithValue{} }
func (m *RelationWithValue) String() string { return proto.CompactTextString(m) }
func (*RelationWithValue) ProtoMessage()    {}
func (*RelationWithValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7f8dc729d740ffd, []int{1}
}
func (m *RelationWithValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RelationWithValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RelationWithValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RelationWithValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelationWithValue.Merge(m, src)
}
func (m *RelationWithValue) XXX_Size() int {
	return m.Size()
}
func (m *RelationWithValue) XXX_DiscardUnknown() {
	xxx_messageInfo_RelationWithValue.DiscardUnknown(m)
}

var xxx_messageInfo_RelationWithValue proto.InternalMessageInfo

func (m *RelationWithValue) GetRelation() *Relation {
	if m != nil {
		return m.Relation
	}
	return nil
}

func (m *RelationWithValue) GetValue() *types.Value {
	if m != nil {
		return m.Value
	}
	return nil
}

// Relation describe the human-interpreted relation type. It may be something like "Date of creation, format=date" or "Assignee, format=objectId, objectType=person"
type Relation struct {
	Format       RelationFormat             `protobuf:"varint,1,opt,name=format,proto3,enum=anytype.relation.RelationFormat" json:"format,omitempty"`
	DefaultName  string                     `protobuf:"bytes,2,opt,name=defaultName,proto3" json:"defaultName,omitempty"`
	DefaultValue *types.Value               `protobuf:"bytes,3,opt,name=defaultValue,proto3" json:"defaultValue,omitempty"`
	DataKey      string                     `protobuf:"bytes,4,opt,name=dataKey,proto3" json:"dataKey,omitempty"`
	DataSource   RelationRelationDataSource `protobuf:"varint,5,opt,name=dataSource,proto3,enum=anytype.relation.RelationRelationDataSource" json:"dataSource,omitempty"`
	Hidden       bool                       `protobuf:"varint,6,opt,name=hidden,proto3" json:"hidden,omitempty"`
	ReadOnly     bool                       `protobuf:"varint,7,opt,name=readOnly,proto3" json:"readOnly,omitempty"`
	ObjectType   string                     `protobuf:"bytes,10,opt,name=objectType,proto3" json:"objectType,omitempty"`
	SelectDict   []string                   `protobuf:"bytes,11,rep,name=selectDict,proto3" json:"selectDict,omitempty"`
}

func (m *Relation) Reset()         { *m = Relation{} }
func (m *Relation) String() string { return proto.CompactTextString(m) }
func (*Relation) ProtoMessage()    {}
func (*Relation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7f8dc729d740ffd, []int{2}
}
func (m *Relation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Relation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Relation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Relation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Relation.Merge(m, src)
}
func (m *Relation) XXX_Size() int {
	return m.Size()
}
func (m *Relation) XXX_DiscardUnknown() {
	xxx_messageInfo_Relation.DiscardUnknown(m)
}

var xxx_messageInfo_Relation proto.InternalMessageInfo

func (m *Relation) GetFormat() RelationFormat {
	if m != nil {
		return m.Format
	}
	return RelationFormat_shortText
}

func (m *Relation) GetDefaultName() string {
	if m != nil {
		return m.DefaultName
	}
	return ""
}

func (m *Relation) GetDefaultValue() *types.Value {
	if m != nil {
		return m.DefaultValue
	}
	return nil
}

func (m *Relation) GetDataKey() string {
	if m != nil {
		return m.DataKey
	}
	return ""
}

func (m *Relation) GetDataSource() RelationRelationDataSource {
	if m != nil {
		return m.DataSource
	}
	return Relation_details
}

func (m *Relation) GetHidden() bool {
	if m != nil {
		return m.Hidden
	}
	return false
}

func (m *Relation) GetReadOnly() bool {
	if m != nil {
		return m.ReadOnly
	}
	return false
}

func (m *Relation) GetObjectType() string {
	if m != nil {
		return m.ObjectType
	}
	return ""
}

func (m *Relation) GetSelectDict() []string {
	if m != nil {
		return m.SelectDict
	}
	return nil
}

func init() {
	proto.RegisterEnum("anytype.relation.RelationFormat", RelationFormat_name, RelationFormat_value)
	proto.RegisterEnum("anytype.relation.RelationRelationDataSource", RelationRelationDataSource_name, RelationRelationDataSource_value)
	proto.RegisterType((*ObjectType)(nil), "anytype.relation.ObjectType")
	proto.RegisterType((*RelationWithValue)(nil), "anytype.relation.RelationWithValue")
	proto.RegisterType((*Relation)(nil), "anytype.relation.Relation")
}

func init() {
	proto.RegisterFile("pkg/lib/pb/relation/protos/relation.proto", fileDescriptor_d7f8dc729d740ffd)
}

var fileDescriptor_d7f8dc729d740ffd = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xd1, 0x6a, 0xdb, 0x3c,
	0x18, 0x8d, 0xe3, 0xc4, 0x71, 0x3e, 0xf7, 0xef, 0xaf, 0x7d, 0x83, 0x62, 0xca, 0x30, 0x26, 0x57,
	0xdd, 0x68, 0x1d, 0xe8, 0x60, 0x94, 0x5d, 0x8e, 0x32, 0x18, 0x63, 0x2b, 0xb8, 0x65, 0x83, 0xde,
	0xc9, 0xb6, 0x92, 0xb8, 0x55, 0x2c, 0x63, 0xcb, 0xa3, 0x7e, 0x8b, 0xbd, 0xc8, 0x60, 0x8f, 0xb1,
	0xcb, 0x5e, 0xee, 0x72, 0x24, 0x2f, 0x32, 0x24, 0xc7, 0x6e, 0xba, 0xad, 0xbd, 0x3b, 0xe7, 0xe8,
	0xe8, 0x7c, 0xc7, 0x92, 0x0c, 0xcf, 0xf3, 0xeb, 0xf9, 0x94, 0xa7, 0xd1, 0x34, 0x8f, 0xa6, 0x05,
	0xe3, 0x54, 0xa6, 0x22, 0x9b, 0xe6, 0x85, 0x90, 0xa2, 0xec, 0x78, 0xa0, 0x39, 0x12, 0x9a, 0xd5,
	0xb2, 0xce, 0x59, 0xd0, 0xea, 0xfb, 0xcf, 0xe6, 0x42, 0xcc, 0x39, 0x6b, 0xfc, 0x51, 0x35, 0x9b,
	0x96, 0xb2, 0xa8, 0x62, 0xd9, 0xf8, 0x27, 0x97, 0x00, 0x67, 0xd1, 0x15, 0x8b, 0xe5, 0x45, 0x9d,
	0x33, 0x44, 0x18, 0x64, 0x74, 0xc9, 0x5c, 0xc3, 0x37, 0x0e, 0xc6, 0xa1, 0xc6, 0x78, 0x02, 0xe3,
	0x36, 0xab, 0x74, 0xfb, 0xbe, 0x79, 0xe0, 0x1c, 0xef, 0x07, 0x7f, 0x4e, 0x09, 0xc2, 0x0d, 0x08,
	0xef, 0xcc, 0x93, 0x1a, 0x9e, 0xb4, 0xf2, 0xe7, 0x54, 0x2e, 0x3e, 0x51, 0x5e, 0x31, 0x7c, 0x05,
	0x76, 0xeb, 0xd0, 0x63, 0x1e, 0x4f, 0xeb, 0xbc, 0x78, 0x08, 0xc3, 0x2f, 0x2a, 0xc0, 0xed, 0xeb,
	0x4d, 0x7b, 0x41, 0xf3, 0x59, 0x41, 0xfb, 0x59, 0x81, 0x8e, 0x0f, 0x1b, 0xd3, 0xe4, 0x9b, 0x09,
	0x76, 0x1b, 0x82, 0x27, 0x60, 0xcd, 0x44, 0xb1, 0xa4, 0x52, 0x0f, 0xdc, 0x3d, 0xf6, 0x1f, 0x1e,
	0xf8, 0x56, 0xfb, 0xc2, 0x8d, 0x1f, 0x7d, 0x70, 0x12, 0x36, 0xa3, 0x15, 0x97, 0x1f, 0xd5, 0xb1,
	0xf4, 0xf5, 0xb1, 0x6c, 0x4b, 0xf8, 0x1a, 0x76, 0x36, 0x54, 0xcf, 0x77, 0xcd, 0x47, 0xdb, 0xdd,
	0xf3, 0xa2, 0x0b, 0xa3, 0x84, 0x4a, 0xfa, 0x9e, 0xd5, 0xee, 0x40, 0x27, 0xb7, 0x14, 0x3f, 0x00,
	0x28, 0x78, 0x2e, 0xaa, 0x22, 0x66, 0xee, 0x50, 0xb7, 0x3e, 0x7a, 0xb8, 0x75, 0x07, 0x4e, 0xbb,
	0x4d, 0xe1, 0x56, 0x00, 0xee, 0x81, 0xb5, 0x48, 0x93, 0x84, 0x65, 0xae, 0xe5, 0x1b, 0x07, 0x76,
	0xb8, 0x61, 0xb8, 0xaf, 0xee, 0x82, 0x26, 0x67, 0x19, 0xaf, 0xdd, 0x91, 0x5e, 0xe9, 0x38, 0x7a,
	0x00, 0xa2, 0x7b, 0x18, 0x2e, 0xe8, 0x7e, 0x5b, 0x8a, 0x5a, 0x2f, 0x19, 0x67, 0xb1, 0x3c, 0x4d,
	0x63, 0xe9, 0x3a, 0xbe, 0xa9, 0xd6, 0xef, 0x94, 0xc9, 0x21, 0xe0, 0xdf, 0xad, 0xd0, 0x81, 0x51,
	0xc2, 0x24, 0x4d, 0x79, 0x49, 0x7a, 0x38, 0x86, 0x21, 0x17, 0x31, 0xe5, 0xc4, 0x78, 0xf1, 0xdd,
	0x80, 0xdd, 0xfb, 0x77, 0x80, 0xff, 0xc1, 0xb8, 0x5c, 0x88, 0x42, 0x5e, 0xb0, 0x1b, 0x49, 0x7a,
	0xb8, 0x03, 0x36, 0x17, 0xd9, 0x5c, 0x33, 0x03, 0x01, 0xac, 0xac, 0x5a, 0x46, 0xac, 0x20, 0x7d,
	0x85, 0x9b, 0xb9, 0xc4, 0xc4, 0xff, 0xc1, 0x59, 0x56, 0x5c, 0xa6, 0xe7, 0x8d, 0x30, 0x40, 0x1b,
	0x06, 0x09, 0x95, 0x8c, 0x0c, 0x15, 0x9a, 0xa5, 0x9c, 0x11, 0x4b, 0x45, 0xc5, 0x0b, 0x16, 0x5f,
	0x47, 0xe2, 0x86, 0x8c, 0x70, 0x04, 0x66, 0x55, 0x70, 0x62, 0xab, 0x3a, 0x6c, 0x49, 0x53, 0x4e,
	0xc6, 0x0a, 0xe6, 0x0b, 0x91, 0x31, 0x02, 0x8d, 0x2a, 0xae, 0x52, 0xe2, 0xa8, 0x7d, 0xcd, 0x01,
	0xbc, 0x4b, 0x48, 0xf2, 0xe6, 0xe8, 0xc7, 0xca, 0x33, 0x6e, 0x57, 0x9e, 0xf1, 0x6b, 0xe5, 0x19,
	0x5f, 0xd7, 0x5e, 0xef, 0x76, 0xed, 0xf5, 0x7e, 0xae, 0xbd, 0xde, 0xe5, 0xd3, 0x7f, 0xfc, 0xae,
	0x91, 0xa5, 0x9f, 0xc2, 0xcb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd1, 0xb4, 0xb0, 0x09, 0xcc,
	0x03, 0x00, 0x00,
}

func (m *ObjectType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ObjectType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ObjectType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Relations) > 0 {
		for iNdEx := len(m.Relations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Relations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRelation(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintRelation(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RelationWithValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RelationWithValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RelationWithValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != nil {
		{
			size, err := m.Value.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRelation(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Relation != nil {
		{
			size, err := m.Relation.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRelation(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Relation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Relation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Relation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SelectDict) > 0 {
		for iNdEx := len(m.SelectDict) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SelectDict[iNdEx])
			copy(dAtA[i:], m.SelectDict[iNdEx])
			i = encodeVarintRelation(dAtA, i, uint64(len(m.SelectDict[iNdEx])))
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.ObjectType) > 0 {
		i -= len(m.ObjectType)
		copy(dAtA[i:], m.ObjectType)
		i = encodeVarintRelation(dAtA, i, uint64(len(m.ObjectType)))
		i--
		dAtA[i] = 0x52
	}
	if m.ReadOnly {
		i--
		if m.ReadOnly {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.Hidden {
		i--
		if m.Hidden {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.DataSource != 0 {
		i = encodeVarintRelation(dAtA, i, uint64(m.DataSource))
		i--
		dAtA[i] = 0x28
	}
	if len(m.DataKey) > 0 {
		i -= len(m.DataKey)
		copy(dAtA[i:], m.DataKey)
		i = encodeVarintRelation(dAtA, i, uint64(len(m.DataKey)))
		i--
		dAtA[i] = 0x22
	}
	if m.DefaultValue != nil {
		{
			size, err := m.DefaultValue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRelation(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DefaultName) > 0 {
		i -= len(m.DefaultName)
		copy(dAtA[i:], m.DefaultName)
		i = encodeVarintRelation(dAtA, i, uint64(len(m.DefaultName)))
		i--
		dAtA[i] = 0x12
	}
	if m.Format != 0 {
		i = encodeVarintRelation(dAtA, i, uint64(m.Format))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRelation(dAtA []byte, offset int, v uint64) int {
	offset -= sovRelation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ObjectType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRelation(uint64(l))
	}
	if len(m.Relations) > 0 {
		for _, e := range m.Relations {
			l = e.Size()
			n += 1 + l + sovRelation(uint64(l))
		}
	}
	return n
}

func (m *RelationWithValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Relation != nil {
		l = m.Relation.Size()
		n += 1 + l + sovRelation(uint64(l))
	}
	if m.Value != nil {
		l = m.Value.Size()
		n += 1 + l + sovRelation(uint64(l))
	}
	return n
}

func (m *Relation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Format != 0 {
		n += 1 + sovRelation(uint64(m.Format))
	}
	l = len(m.DefaultName)
	if l > 0 {
		n += 1 + l + sovRelation(uint64(l))
	}
	if m.DefaultValue != nil {
		l = m.DefaultValue.Size()
		n += 1 + l + sovRelation(uint64(l))
	}
	l = len(m.DataKey)
	if l > 0 {
		n += 1 + l + sovRelation(uint64(l))
	}
	if m.DataSource != 0 {
		n += 1 + sovRelation(uint64(m.DataSource))
	}
	if m.Hidden {
		n += 2
	}
	if m.ReadOnly {
		n += 2
	}
	l = len(m.ObjectType)
	if l > 0 {
		n += 1 + l + sovRelation(uint64(l))
	}
	if len(m.SelectDict) > 0 {
		for _, s := range m.SelectDict {
			l = len(s)
			n += 1 + l + sovRelation(uint64(l))
		}
	}
	return n
}

func sovRelation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRelation(x uint64) (n int) {
	return sovRelation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ObjectType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRelation
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ObjectType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ObjectType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRelation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRelation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Relations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRelation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRelation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Relations = append(m.Relations, &Relation{})
			if err := m.Relations[len(m.Relations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRelation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRelation
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRelation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RelationWithValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRelation
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RelationWithValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RelationWithValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Relation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRelation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRelation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Relation == nil {
				m.Relation = &Relation{}
			}
			if err := m.Relation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRelation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRelation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Value == nil {
				m.Value = &types.Value{}
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRelation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRelation
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRelation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Relation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRelation
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Relation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Relation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Format", wireType)
			}
			m.Format = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Format |= RelationFormat(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRelation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRelation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefaultName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRelation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRelation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DefaultValue == nil {
				m.DefaultValue = &types.Value{}
			}
			if err := m.DefaultValue.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRelation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRelation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataSource", wireType)
			}
			m.DataSource = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DataSource |= RelationRelationDataSource(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hidden", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Hidden = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReadOnly", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ReadOnly = bool(v != 0)
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRelation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRelation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObjectType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SelectDict", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRelation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRelation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SelectDict = append(m.SelectDict, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRelation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRelation
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRelation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRelation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRelation
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRelation
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRelation
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthRelation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRelation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRelation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRelation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRelation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRelation = fmt.Errorf("proto: unexpected end of group")
)
