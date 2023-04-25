// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/protos/snapshot.proto

package pb

import (
	fmt "fmt"
	model "github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	proto "github.com/gogo/protobuf/proto"
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

type SnapshotWithType struct {
	SbType   model.SmartBlockType `protobuf:"varint,1,opt,name=sbType,proto3,enum=anytype.model.SmartBlockType" json:"sbType,omitempty"`
	Snapshot *ChangeSnapshot      `protobuf:"bytes,2,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
}

func (m *SnapshotWithType) Reset()         { *m = SnapshotWithType{} }
func (m *SnapshotWithType) String() string { return proto.CompactTextString(m) }
func (*SnapshotWithType) ProtoMessage()    {}
func (*SnapshotWithType) Descriptor() ([]byte, []int) {
	return fileDescriptor_022f1596c727bff6, []int{0}
}
func (m *SnapshotWithType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SnapshotWithType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SnapshotWithType.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SnapshotWithType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotWithType.Merge(m, src)
}
func (m *SnapshotWithType) XXX_Size() int {
	return m.Size()
}
func (m *SnapshotWithType) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotWithType.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotWithType proto.InternalMessageInfo

func (m *SnapshotWithType) GetSbType() model.SmartBlockType {
	if m != nil {
		return m.SbType
	}
	return model.SmartBlockType_AccountOld
}

func (m *SnapshotWithType) GetSnapshot() *ChangeSnapshot {
	if m != nil {
		return m.Snapshot
	}
	return nil
}

type Profile struct {
	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Avatar  string `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_022f1596c727bff6, []int{1}
}
func (m *Profile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(m, src)
}
func (m *Profile) XXX_Size() int {
	return m.Size()
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

func (m *Profile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Profile) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *Profile) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*SnapshotWithType)(nil), "anytype.SnapshotWithType")
	proto.RegisterType((*Profile)(nil), "anytype.Profile")
}

func init() { proto.RegisterFile("pb/protos/snapshot.proto", fileDescriptor_022f1596c727bff6) }

var fileDescriptor_022f1596c727bff6 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0x9b, 0x31, 0x5a, 0x17, 0x41, 0x24, 0x07, 0x0d, 0x43, 0xc3, 0x18, 0x1e, 0x76, 0x4a,
	0x60, 0xea, 0x0b, 0xcc, 0x07, 0x50, 0x3a, 0x41, 0xf0, 0x96, 0xac, 0x71, 0x2d, 0x6b, 0x9b, 0x90,
	0x04, 0xa1, 0x27, 0x5f, 0xc1, 0xc7, 0xf2, 0xb8, 0xa3, 0x47, 0x69, 0x5f, 0x44, 0x96, 0xa5, 0xdd,
	0xed, 0xfb, 0xf3, 0xfd, 0xbf, 0xdf, 0x2f, 0x81, 0x58, 0x0b, 0xa6, 0x8d, 0x72, 0xca, 0x32, 0x5b,
	0x73, 0x6d, 0x73, 0xe5, 0xa8, 0xcf, 0x28, 0xe1, 0x75, 0xe3, 0x1a, 0x2d, 0xa7, 0x77, 0x7a, 0xb7,
	0x65, 0x65, 0x21, 0x98, 0x16, 0xac, 0x52, 0x99, 0x2c, 0xfb, 0x03, 0x1f, 0xec, 0xb1, 0x3e, 0xbd,
	0x3e, 0x81, 0x36, 0x39, 0xaf, 0xb7, 0x32, 0x2c, 0xe6, 0x5f, 0xf0, 0x72, 0x1d, 0xc8, 0x6f, 0x85,
	0xcb, 0x5f, 0x1b, 0x2d, 0xd1, 0x23, 0x8c, 0xad, 0x38, 0x4c, 0x18, 0xcc, 0xc0, 0xe2, 0x62, 0x79,
	0x4b, 0x83, 0x8c, 0x7a, 0x26, 0x5d, 0x57, 0xdc, 0xb8, 0x55, 0xa9, 0x36, 0xbb, 0x43, 0x29, 0x0d,
	0x65, 0xf4, 0x00, 0xcf, 0xfa, 0x47, 0xe2, 0xd1, 0x0c, 0x2c, 0xce, 0x97, 0x78, 0x38, 0x7c, 0xf2,
	0x52, 0xda, 0xab, 0xd2, 0xa1, 0x39, 0x7f, 0x86, 0xc9, 0x8b, 0x51, 0x1f, 0x45, 0x29, 0x11, 0x82,
	0xe3, 0x9a, 0x57, 0x47, 0xeb, 0x24, 0xf5, 0x33, 0xba, 0x82, 0x31, 0xff, 0xe4, 0x8e, 0x1b, 0x8f,
	0x9c, 0xa4, 0x21, 0x21, 0x0c, 0x13, 0x9e, 0x65, 0x46, 0x5a, 0x8b, 0xc7, 0x7e, 0xd1, 0xc7, 0xd5,
	0xcd, 0x4f, 0x4b, 0xc0, 0xbe, 0x25, 0xe0, 0xaf, 0x25, 0xe0, 0xbb, 0x23, 0xd1, 0xbe, 0x23, 0xd1,
	0x6f, 0x47, 0xa2, 0xf7, 0x91, 0x16, 0x22, 0xf6, 0xdf, 0xbe, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff,
	0x69, 0xb8, 0x0e, 0xc3, 0x5a, 0x01, 0x00, 0x00,
}

func (m *SnapshotWithType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SnapshotWithType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SnapshotWithType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Snapshot != nil {
		{
			size, err := m.Snapshot.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSnapshot(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.SbType != 0 {
		i = encodeVarintSnapshot(dAtA, i, uint64(m.SbType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Profile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Profile) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Profile) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Avatar) > 0 {
		i -= len(m.Avatar)
		copy(dAtA[i:], m.Avatar)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.Avatar)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSnapshot(dAtA []byte, offset int, v uint64) int {
	offset -= sovSnapshot(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SnapshotWithType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SbType != 0 {
		n += 1 + sovSnapshot(uint64(m.SbType))
	}
	if m.Snapshot != nil {
		l = m.Snapshot.Size()
		n += 1 + l + sovSnapshot(uint64(l))
	}
	return n
}

func (m *Profile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	l = len(m.Avatar)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	return n
}

func sovSnapshot(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSnapshot(x uint64) (n int) {
	return sovSnapshot(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SnapshotWithType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
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
			return fmt.Errorf("proto: SnapshotWithType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SnapshotWithType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SbType", wireType)
			}
			m.SbType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SbType |= model.SmartBlockType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Snapshot", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Snapshot == nil {
				m.Snapshot = &ChangeSnapshot{}
			}
			if err := m.Snapshot.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSnapshot
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
func (m *Profile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
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
			return fmt.Errorf("proto: Profile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Profile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Avatar", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Avatar = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSnapshot
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
func skipSnapshot(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSnapshot
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
					return 0, ErrIntOverflowSnapshot
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
					return 0, ErrIntOverflowSnapshot
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
				return 0, ErrInvalidLengthSnapshot
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSnapshot
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSnapshot
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSnapshot        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSnapshot          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSnapshot = fmt.Errorf("proto: unexpected end of group")
)
