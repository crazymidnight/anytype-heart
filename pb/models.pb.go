// Code generated by protoc-gen-go. DO NOT EDIT.
// source: models.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type StatusType int32

const (
	StatusType_SUCCESS        StatusType = 0
	StatusType_FAILURE        StatusType = 1
	StatusType_WRONG_MNEMONIC StatusType = 2
	StatusType_NOT_FOUND      StatusType = 3
)

var StatusType_name = map[int32]string{
	0: "SUCCESS",
	1: "FAILURE",
	2: "WRONG_MNEMONIC",
	3: "NOT_FOUND",
}

var StatusType_value = map[string]int32{
	"SUCCESS":        0,
	"FAILURE":        1,
	"WRONG_MNEMONIC": 2,
	"NOT_FOUND":      3,
}

func (x StatusType) String() string {
	return proto.EnumName(StatusType_name, int32(x))
}

func (StatusType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0b5431a010549573, []int{0}
}

type MarkType int32

const (
	MarkType_B   MarkType = 0
	MarkType_I   MarkType = 1
	MarkType_S   MarkType = 2
	MarkType_KBD MarkType = 3
	MarkType_A   MarkType = 4
)

var MarkType_name = map[int32]string{
	0: "B",
	1: "I",
	2: "S",
	3: "KBD",
	4: "A",
}

var MarkType_value = map[string]int32{
	"B":   0,
	"I":   1,
	"S":   2,
	"KBD": 3,
	"A":   4,
}

func (x MarkType) String() string {
	return proto.EnumName(MarkType_name, int32(x))
}

func (MarkType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0b5431a010549573, []int{1}
}

type ContentType int32

const (
	ContentType_P      ContentType = 0
	ContentType_CODE   ContentType = 1
	ContentType_H_1    ContentType = 2
	ContentType_H_2    ContentType = 3
	ContentType_H_3    ContentType = 4
	ContentType_H_4    ContentType = 5
	ContentType_OL     ContentType = 6
	ContentType_UL     ContentType = 7
	ContentType_QUOTE  ContentType = 8
	ContentType_TOGGLE ContentType = 9
	ContentType_CHECK  ContentType = 10
)

var ContentType_name = map[int32]string{
	0:  "P",
	1:  "CODE",
	2:  "H_1",
	3:  "H_2",
	4:  "H_3",
	5:  "H_4",
	6:  "OL",
	7:  "UL",
	8:  "QUOTE",
	9:  "TOGGLE",
	10: "CHECK",
}

var ContentType_value = map[string]int32{
	"P":      0,
	"CODE":   1,
	"H_1":    2,
	"H_2":    3,
	"H_3":    4,
	"H_4":    5,
	"OL":     6,
	"UL":     7,
	"QUOTE":  8,
	"TOGGLE": 9,
	"CHECK":  10,
}

func (x ContentType) String() string {
	return proto.EnumName(ContentType_name, int32(x))
}

func (ContentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0b5431a010549573, []int{2}
}

type BlockType int32

const (
	BlockType_H_GRID    BlockType = 0
	BlockType_V_GRID    BlockType = 1
	BlockType_EDITABLE  BlockType = 2
	BlockType_DIV       BlockType = 3
	BlockType_VIDEO     BlockType = 4
	BlockType_IMAGE     BlockType = 5
	BlockType_PAGE      BlockType = 6
	BlockType_NEW_PAGE  BlockType = 7
	BlockType_BOOK_MARK BlockType = 8
	BlockType_FILE      BlockType = 9
	BlockType_DATA_VIEW BlockType = 10
)

var BlockType_name = map[int32]string{
	0:  "H_GRID",
	1:  "V_GRID",
	2:  "EDITABLE",
	3:  "DIV",
	4:  "VIDEO",
	5:  "IMAGE",
	6:  "PAGE",
	7:  "NEW_PAGE",
	8:  "BOOK_MARK",
	9:  "FILE",
	10: "DATA_VIEW",
}

var BlockType_value = map[string]int32{
	"H_GRID":    0,
	"V_GRID":    1,
	"EDITABLE":  2,
	"DIV":       3,
	"VIDEO":     4,
	"IMAGE":     5,
	"PAGE":      6,
	"NEW_PAGE":  7,
	"BOOK_MARK": 8,
	"FILE":      9,
	"DATA_VIEW": 10,
}

func (x BlockType) String() string {
	return proto.EnumName(BlockType_name, int32(x))
}

func (BlockType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0b5431a010549573, []int{3}
}

type Account struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Avatar               string   `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b5431a010549573, []int{0}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Account) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Account) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

type Accounts struct {
	Accounts             []*Account `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Accounts) Reset()         { *m = Accounts{} }
func (m *Accounts) String() string { return proto.CompactTextString(m) }
func (*Accounts) ProtoMessage()    {}
func (*Accounts) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b5431a010549573, []int{1}
}

func (m *Accounts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Accounts.Unmarshal(m, b)
}
func (m *Accounts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Accounts.Marshal(b, m, deterministic)
}
func (m *Accounts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Accounts.Merge(m, src)
}
func (m *Accounts) XXX_Size() int {
	return xxx_messageInfo_Accounts.Size(m)
}
func (m *Accounts) XXX_DiscardUnknown() {
	xxx_messageInfo_Accounts.DiscardUnknown(m)
}

var xxx_messageInfo_Accounts proto.InternalMessageInfo

func (m *Accounts) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

type State struct {
	DocumentHeaders      []*DocumentHeader `protobuf:"bytes,1,rep,name=documentHeaders,proto3" json:"documentHeaders,omitempty"`
	CurrentDocumentId    string            `protobuf:"bytes,2,opt,name=currentDocumentId,proto3" json:"currentDocumentId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b5431a010549573, []int{2}
}

func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (m *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(m, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetDocumentHeaders() []*DocumentHeader {
	if m != nil {
		return m.DocumentHeaders
	}
	return nil
}

func (m *State) GetCurrentDocumentId() string {
	if m != nil {
		return m.CurrentDocumentId
	}
	return ""
}

type DocumentHeader struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Icon                 string   `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DocumentHeader) Reset()         { *m = DocumentHeader{} }
func (m *DocumentHeader) String() string { return proto.CompactTextString(m) }
func (*DocumentHeader) ProtoMessage()    {}
func (*DocumentHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b5431a010549573, []int{3}
}

func (m *DocumentHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocumentHeader.Unmarshal(m, b)
}
func (m *DocumentHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocumentHeader.Marshal(b, m, deterministic)
}
func (m *DocumentHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocumentHeader.Merge(m, src)
}
func (m *DocumentHeader) XXX_Size() int {
	return xxx_messageInfo_DocumentHeader.Size(m)
}
func (m *DocumentHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_DocumentHeader.DiscardUnknown(m)
}

var xxx_messageInfo_DocumentHeader proto.InternalMessageInfo

func (m *DocumentHeader) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DocumentHeader) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DocumentHeader) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *DocumentHeader) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

type Document struct {
	Header               *DocumentHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Blocks               []*Block        `protobuf:"bytes,2,rep,name=blocks,proto3" json:"blocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Document) Reset()         { *m = Document{} }
func (m *Document) String() string { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()    {}
func (*Document) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b5431a010549573, []int{4}
}

func (m *Document) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Document.Unmarshal(m, b)
}
func (m *Document) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Document.Marshal(b, m, deterministic)
}
func (m *Document) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Document.Merge(m, src)
}
func (m *Document) XXX_Size() int {
	return xxx_messageInfo_Document.Size(m)
}
func (m *Document) XXX_DiscardUnknown() {
	xxx_messageInfo_Document.DiscardUnknown(m)
}

var xxx_messageInfo_Document proto.InternalMessageInfo

func (m *Document) GetHeader() *DocumentHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Document) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

type Text struct {
	Content              string      `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Marks                []*Mark     `protobuf:"bytes,2,rep,name=marks,proto3" json:"marks,omitempty"`
	ContentType          ContentType `protobuf:"varint,3,opt,name=contentType,proto3,enum=anytype.ContentType" json:"contentType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Text) Reset()         { *m = Text{} }
func (m *Text) String() string { return proto.CompactTextString(m) }
func (*Text) ProtoMessage()    {}
func (*Text) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b5431a010549573, []int{5}
}

func (m *Text) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Text.Unmarshal(m, b)
}
func (m *Text) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Text.Marshal(b, m, deterministic)
}
func (m *Text) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Text.Merge(m, src)
}
func (m *Text) XXX_Size() int {
	return xxx_messageInfo_Text.Size(m)
}
func (m *Text) XXX_DiscardUnknown() {
	xxx_messageInfo_Text.DiscardUnknown(m)
}

var xxx_messageInfo_Text proto.InternalMessageInfo

func (m *Text) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Text) GetMarks() []*Mark {
	if m != nil {
		return m.Marks
	}
	return nil
}

func (m *Text) GetContentType() ContentType {
	if m != nil {
		return m.ContentType
	}
	return ContentType_P
}

type Mark struct {
	Type                 MarkType `protobuf:"varint,1,opt,name=type,proto3,enum=anytype.MarkType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mark) Reset()         { *m = Mark{} }
func (m *Mark) String() string { return proto.CompactTextString(m) }
func (*Mark) ProtoMessage()    {}
func (*Mark) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b5431a010549573, []int{6}
}

func (m *Mark) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mark.Unmarshal(m, b)
}
func (m *Mark) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mark.Marshal(b, m, deterministic)
}
func (m *Mark) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mark.Merge(m, src)
}
func (m *Mark) XXX_Size() int {
	return xxx_messageInfo_Mark.Size(m)
}
func (m *Mark) XXX_DiscardUnknown() {
	xxx_messageInfo_Mark.DiscardUnknown(m)
}

var xxx_messageInfo_Mark proto.InternalMessageInfo

func (m *Mark) GetType() MarkType {
	if m != nil {
		return m.Type
	}
	return MarkType_B
}

type Block struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Content:
	//	*Block_Text
	Content              isBlock_Content `protobuf_oneof:"content"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b5431a010549573, []int{7}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type isBlock_Content interface {
	isBlock_Content()
}

type Block_Text struct {
	Text *Text `protobuf:"bytes,11,opt,name=text,proto3,oneof"`
}

func (*Block_Text) isBlock_Content() {}

func (m *Block) GetContent() isBlock_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Block) GetText() *Text {
	if x, ok := m.GetContent().(*Block_Text); ok {
		return x.Text
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Block) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Block_Text)(nil),
	}
}

func init() {
	proto.RegisterEnum("anytype.StatusType", StatusType_name, StatusType_value)
	proto.RegisterEnum("anytype.MarkType", MarkType_name, MarkType_value)
	proto.RegisterEnum("anytype.ContentType", ContentType_name, ContentType_value)
	proto.RegisterEnum("anytype.BlockType", BlockType_name, BlockType_value)
	proto.RegisterType((*Account)(nil), "anytype.Account")
	proto.RegisterType((*Accounts)(nil), "anytype.Accounts")
	proto.RegisterType((*State)(nil), "anytype.State")
	proto.RegisterType((*DocumentHeader)(nil), "anytype.DocumentHeader")
	proto.RegisterType((*Document)(nil), "anytype.Document")
	proto.RegisterType((*Text)(nil), "anytype.Text")
	proto.RegisterType((*Mark)(nil), "anytype.Mark")
	proto.RegisterType((*Block)(nil), "anytype.Block")
}

func init() { proto.RegisterFile("models.proto", fileDescriptor_0b5431a010549573) }

var fileDescriptor_0b5431a010549573 = []byte{
	// 628 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0xb5, 0xbe, 0xa5, 0x71, 0xe2, 0x6e, 0x96, 0xd2, 0xea, 0x18, 0x14, 0x5a, 0x82, 0x49, 0x13,
	0xea, 0x94, 0xd2, 0xab, 0xbe, 0x62, 0x0b, 0x7f, 0x28, 0x95, 0x65, 0x07, 0x7a, 0x11, 0xb2, 0xbc,
	0x50, 0x93, 0x58, 0x32, 0xf2, 0x3a, 0x24, 0x87, 0xfe, 0x87, 0xfe, 0xe4, 0xb2, 0x2b, 0x59, 0x49,
	0x1a, 0x28, 0x3d, 0xed, 0x9b, 0x99, 0x37, 0x33, 0x6f, 0x1e, 0x42, 0x70, 0xb0, 0x2e, 0x96, 0xe4,
	0x6e, 0x7b, 0xbe, 0x29, 0x0b, 0x5a, 0x60, 0x2d, 0xcd, 0x1f, 0xe9, 0xe3, 0x86, 0x58, 0x3e, 0x68,
	0x76, 0x96, 0x15, 0xbb, 0x9c, 0xe2, 0x0e, 0x88, 0xab, 0xa5, 0x29, 0x1c, 0x0b, 0xa7, 0x46, 0x24,
	0xae, 0x96, 0x18, 0x83, 0x9c, 0xa7, 0x6b, 0x62, 0x8a, 0x3c, 0xc3, 0x31, 0x7e, 0x07, 0x6a, 0x7a,
	0x9f, 0xd2, 0xb4, 0x34, 0x25, 0x9e, 0xad, 0x23, 0xeb, 0x1b, 0xe8, 0xf5, 0x98, 0x2d, 0x3e, 0x03,
	0x3d, 0xad, 0xb1, 0x29, 0x1c, 0x4b, 0xa7, 0xed, 0x1e, 0x3a, 0xaf, 0xd7, 0x9d, 0xd7, 0xa4, 0xa8,
	0x61, 0x58, 0x0f, 0xa0, 0x4c, 0x69, 0x4a, 0x09, 0xb6, 0xe1, 0xcd, 0xb2, 0xc8, 0x76, 0x6b, 0x92,
	0xd3, 0x01, 0x49, 0x97, 0xa4, 0xdc, 0x77, 0xbf, 0x6f, 0xba, 0xbd, 0x17, 0xf5, 0xe8, 0x6f, 0x3e,
	0x3e, 0x83, 0xa3, 0x6c, 0x57, 0x96, 0x24, 0xa7, 0x7b, 0x66, 0xb0, 0xac, 0xe5, 0xbf, 0x2e, 0x58,
	0x0b, 0xe8, 0xbc, 0x1c, 0xf8, 0x5f, 0x0e, 0x98, 0xa0, 0xdd, 0x93, 0x72, 0xbb, 0x2a, 0xf2, 0xda,
	0x82, 0x7d, 0xc8, 0xd8, 0xab, 0xac, 0xc8, 0x4d, 0xb9, 0x62, 0x33, 0x6c, 0x65, 0xa0, 0xef, 0x77,
	0xe0, 0x0b, 0x50, 0x7f, 0xf2, 0x3d, 0x7c, 0xc3, 0x3f, 0xee, 0xaa, 0x69, 0xf8, 0x23, 0xa8, 0x8b,
	0xbb, 0x22, 0xbb, 0xdd, 0x9a, 0x22, 0x37, 0xa2, 0xd3, 0x34, 0x38, 0x2c, 0x1d, 0xd5, 0x55, 0xeb,
	0x17, 0xc8, 0x31, 0x79, 0xa0, 0x4c, 0x5a, 0x56, 0xe4, 0x94, 0xe4, 0xb4, 0xbe, 0x61, 0x1f, 0xe2,
	0x13, 0x50, 0xd6, 0x69, 0xd9, 0x0c, 0x3a, 0x6c, 0x06, 0x8d, 0xd3, 0xf2, 0x36, 0xaa, 0x6a, 0xf8,
	0x2b, 0xb4, 0x6b, 0x7e, 0xfc, 0xb8, 0x21, 0xfc, 0xba, 0x4e, 0xef, 0x6d, 0x43, 0x75, 0x9f, 0x6a,
	0xd1, 0x73, 0xa2, 0xf5, 0x09, 0x64, 0x36, 0x06, 0x7f, 0x00, 0x99, 0x11, 0xf9, 0xee, 0x4e, 0xef,
	0xe8, 0xc5, 0x0e, 0xde, 0xc5, 0xcb, 0x56, 0x1f, 0x14, 0x2e, 0xff, 0x95, 0xdb, 0x27, 0x20, 0x53,
	0xf2, 0x40, 0xcd, 0x36, 0x77, 0xe7, 0x49, 0x23, 0xbb, 0x6d, 0xd0, 0x8a, 0x78, 0xd1, 0x31, 0x9a,
	0x1b, 0xbb, 0x01, 0x00, 0xfb, 0x72, 0x76, 0x5b, 0x36, 0x1c, 0xb7, 0x41, 0x9b, 0xce, 0x5c, 0xd7,
	0x9f, 0x4e, 0x51, 0x8b, 0x05, 0x57, 0x76, 0x30, 0x9a, 0x45, 0x3e, 0x12, 0x30, 0x86, 0xce, 0x4d,
	0x14, 0x4e, 0xfa, 0xc9, 0x78, 0xe2, 0x8f, 0xc3, 0x49, 0xe0, 0x22, 0x11, 0x1f, 0x82, 0x31, 0x09,
	0xe3, 0xe4, 0x2a, 0x9c, 0x4d, 0x3c, 0x24, 0x75, 0x2f, 0x40, 0xdf, 0xab, 0xc4, 0x0a, 0x08, 0x0e,
	0x6a, 0xb1, 0x27, 0x40, 0x02, 0x7b, 0xa6, 0x48, 0xc4, 0x1a, 0x48, 0x43, 0xc7, 0x43, 0x12, 0x8b,
	0x6d, 0x24, 0x77, 0x29, 0xb4, 0x9f, 0xf9, 0xc1, 0xb2, 0xd7, 0xa8, 0x85, 0x75, 0x90, 0xdd, 0xd0,
	0x63, 0x3b, 0x35, 0x90, 0x06, 0xc9, 0xe7, 0xaa, 0x71, 0x90, 0xf4, 0x90, 0x54, 0x81, 0x4b, 0x24,
	0x57, 0xe0, 0x0b, 0x52, 0xb0, 0x0a, 0x62, 0x38, 0x42, 0x2a, 0x7b, 0x67, 0x23, 0xa4, 0x61, 0x03,
	0x94, 0xef, 0xb3, 0x30, 0xf6, 0x91, 0x8e, 0x01, 0xd4, 0x38, 0xec, 0xf7, 0x47, 0x3e, 0x32, 0x58,
	0xda, 0x1d, 0xf8, 0xee, 0x10, 0x41, 0xf7, 0xb7, 0x00, 0x06, 0xf7, 0x8e, 0x2f, 0x05, 0x50, 0x07,
	0x49, 0x3f, 0x0a, 0x3c, 0xd4, 0x62, 0x78, 0x5e, 0x61, 0x01, 0x1f, 0x80, 0xee, 0x7b, 0x41, 0x6c,
	0x3b, 0x23, 0xbf, 0x12, 0xe0, 0x05, 0x73, 0x24, 0xb1, 0x39, 0xf3, 0xc0, 0xf3, 0x43, 0x24, 0x33,
	0x18, 0x8c, 0xed, 0xbe, 0x8f, 0x14, 0x26, 0xf9, 0x9a, 0x21, 0x95, 0xb5, 0x4d, 0xfc, 0x9b, 0x84,
	0x47, 0x1a, 0x33, 0xc8, 0x09, 0xc3, 0x61, 0x32, 0xb6, 0xa3, 0x21, 0xd2, 0x19, 0xed, 0x2a, 0xe0,
	0x72, 0x0e, 0xc1, 0xf0, 0xec, 0xd8, 0x4e, 0xe6, 0x81, 0x7f, 0x83, 0xc0, 0x91, 0x7f, 0x88, 0x9b,
	0xc5, 0x42, 0xe5, 0x7f, 0x95, 0xcb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xde, 0x4d, 0xc7, 0x04,
	0x65, 0x04, 0x00, 0x00,
}
