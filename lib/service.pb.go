// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service.proto

package lib

import (
	fmt "fmt"
	pb "github.com/anytypeio/go-anytype-middleware/pb"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcf, 0x4e, 0xfa, 0x40,
	0x10, 0xc7, 0x21, 0xbf, 0x04, 0xf2, 0x5b, 0xfe, 0x1c, 0xf6, 0xa2, 0x16, 0x58, 0xdf, 0x80, 0x83,
	0x5e, 0xbd, 0x58, 0x62, 0xd0, 0x60, 0x34, 0x42, 0xd4, 0xc4, 0x5b, 0x5b, 0x46, 0xd2, 0xb8, 0xdd,
	0xad, 0xdd, 0x85, 0x84, 0xb7, 0xf0, 0xb1, 0x8c, 0x27, 0x8e, 0x1e, 0x0d, 0xbc, 0x88, 0x91, 0x9d,
	0x2d, 0x94, 0xb6, 0xb7, 0xed, 0xf7, 0x33, 0xf3, 0x69, 0x27, 0xd3, 0x25, 0x2d, 0x05, 0xc9, 0x22,
	0x0c, 0xa0, 0x1f, 0x27, 0x52, 0x4b, 0x5a, 0xf7, 0xc4, 0x52, 0x2f, 0x63, 0x70, 0x5a, 0x5e, 0x10,
	0xc8, 0xb9, 0xd0, 0x26, 0x77, 0xc8, 0x6b, 0xc8, 0xc1, 0x9e, 0xa3, 0x50, 0x05, 0x78, 0x6e, 0xf8,
	0x5c, 0x06, 0x6f, 0xe6, 0xe1, 0xec, 0xab, 0x46, 0xda, 0x03, 0x1e, 0x82, 0xd0, 0x03, 0x19, 0x45,
	0x9e, 0x98, 0x2a, 0x3a, 0x22, 0xcd, 0x67, 0x8f, 0x73, 0xd0, 0x83, 0x04, 0x3c, 0x0d, 0xb4, 0xdb,
	0xc7, 0x17, 0xf4, 0xf7, 0xe3, 0x31, 0xbc, 0xcf, 0x41, 0x69, 0xa7, 0x57, 0x42, 0x55, 0x2c, 0x85,
	0x02, 0x7a, 0x47, 0x5a, 0x26, 0x1f, 0x43, 0x20, 0x17, 0x90, 0xd0, 0xc3, 0x7a, 0xcc, 0xad, 0x8e,
	0x95, 0x61, 0xf4, 0x3d, 0x90, 0xf6, 0xa5, 0x99, 0xd2, 0x0a, 0x77, 0x1d, 0x59, 0x60, 0x8d, 0xa7,
	0xa5, 0x7c, 0xf7, 0x89, 0x48, 0x70, 0xe0, 0xde, 0x61, 0x47, 0x76, 0x62, 0x56, 0x86, 0x73, 0xbe,
	0x09, 0x70, 0x08, 0x74, 0xde, 0x67, 0xf2, 0x52, 0x9f, 0xc5, 0xe8, 0x1b, 0x91, 0xe6, 0x4d, 0xe4,
	0xcd, 0x60, 0x08, 0xda, 0xe5, 0xd2, 0xdf, 0xdb, 0xc7, 0x7e, 0x9c, 0xdf, 0x47, 0x96, 0xa2, 0xec,
	0x8a, 0x90, 0x27, 0x48, 0x54, 0x28, 0xc5, 0x10, 0x34, 0x75, 0xd2, 0xe2, 0x5d, 0x68, 0x45, 0x9d,
	0x42, 0x86, 0x9a, 0x0b, 0x52, 0xbf, 0x95, 0xb3, 0x09, 0x88, 0x29, 0x3d, 0x4a, 0xeb, 0x30, 0xb1,
	0x82, 0xe3, 0x3c, 0xc0, 0x6e, 0x97, 0xfc, 0x77, 0xff, 0xfe, 0xc1, 0xfb, 0x18, 0x04, 0x3d, 0x49,
	0xcb, 0xd2, 0xcc, 0x1a, 0x9c, 0x22, 0x84, 0x8e, 0x6b, 0xd2, 0xd8, 0x86, 0xb8, 0xb3, 0x4e, 0xb6,
	0x34, 0xbb, 0xb1, 0x6e, 0x31, 0x3c, 0x30, 0x3d, 0xc6, 0xd3, 0x02, 0x93, 0x49, 0x4b, 0x4c, 0x16,
	0x1a, 0x93, 0xdb, 0xfb, 0x5c, 0xb3, 0xea, 0x6a, 0xcd, 0xaa, 0x3f, 0x6b, 0x56, 0xfd, 0xd8, 0xb0,
	0xca, 0x6a, 0xc3, 0x2a, 0xdf, 0x1b, 0x56, 0x79, 0xf9, 0xc7, 0x43, 0xdf, 0xaf, 0x6d, 0xaf, 0xdc,
	0xf9, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x87, 0x48, 0xbc, 0xc0, 0x03, 0x00, 0x00,
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the gomobile package it is being compiled against.

// ClientCommandsHandler is the handler API for ClientCommands service.
var clientCommandsHandler ClientCommandsHandler

type ClientCommandsHandler interface {
	WalletCreate(*pb.WalletCreateRequest) *pb.WalletCreateResponse
	WalletRecover(*pb.WalletRecoverRequest) *pb.WalletRecoverResponse
	AccountRecover(*pb.AccountRecoverRequest) *pb.AccountRecoverResponse
	AccountCreate(*pb.AccountCreateRequest) *pb.AccountCreateResponse
	AccountSelect(*pb.AccountSelectRequest) *pb.AccountSelectResponse
	ImageGetBlob(*pb.ImageGetBlobRequest) *pb.ImageGetBlobResponse
	VersionGet(*pb.VersionGetRequest) *pb.VersionGetResponse
	LogSend(*pb.LogSendRequest) *pb.LogSendResponse
	BlockOpen(*pb.BlockOpenRequest) *pb.BlockOpenResponse
	BlockCreate(*pb.BlockCreateRequest) *pb.BlockCreateResponse
	BlockUpdate(*pb.BlockUpdateRequest) *pb.BlockUpdateResponse
}

func registerClientCommandsHandler(srv ClientCommandsHandler) {
	clientCommandsHandler = srv
}

func WalletCreate(b []byte) []byte {
	in := new(pb.WalletCreateRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.WalletCreateResponse{Error: &pb.WalletCreateResponse_Error{Code: pb.WalletCreateResponse_Error_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.WalletCreate(in).Marshal()
	return resp
}

func WalletRecover(b []byte) []byte {
	in := new(pb.WalletRecoverRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.WalletRecoverResponse{Error: &pb.WalletRecoverResponse_Error{Code: pb.WalletRecoverResponse_Error_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.WalletRecover(in).Marshal()
	return resp
}

func AccountRecover(b []byte) []byte {
	in := new(pb.AccountRecoverRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.AccountRecoverResponse{Error: &pb.AccountRecoverResponse_Error{Code: pb.AccountRecoverResponse_Error_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.AccountRecover(in).Marshal()
	return resp
}

func AccountCreate(b []byte) []byte {
	in := new(pb.AccountCreateRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.AccountCreateResponse{Error: &pb.AccountCreateResponse_Error{Code: pb.AccountCreateResponse_Error_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.AccountCreate(in).Marshal()
	return resp
}

func AccountSelect(b []byte) []byte {
	in := new(pb.AccountSelectRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.AccountSelectResponse{Error: &pb.AccountSelectResponse_Error{Code: pb.AccountSelectResponse_Error_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.AccountSelect(in).Marshal()
	return resp
}

func ImageGetBlob(b []byte) []byte {
	in := new(pb.ImageGetBlobRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.ImageGetBlobResponse{Error: &pb.ImageGetBlobResponse_Error{Code: pb.ImageGetBlobResponse_Error_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.ImageGetBlob(in).Marshal()
	return resp
}

func VersionGet(b []byte) []byte {
	in := new(pb.VersionGetRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.VersionGetResponse{Error: &pb.VersionGetResponse_Error{Code: pb.VersionGetResponse_Error_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.VersionGet(in).Marshal()
	return resp
}

func LogSend(b []byte) []byte {
	in := new(pb.LogSendRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.LogSendResponse{Error: &pb.LogSendResponse_Error{Code: pb.LogSendResponse_Error_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.LogSend(in).Marshal()
	return resp
}

func BlockOpen(b []byte) []byte {
	in := new(pb.BlockOpenRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.BlockOpenResponse{Error: &pb.BlockOpenResponse_Error{Code: pb.BlockOpenResponse_Error_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockOpen(in).Marshal()
	return resp
}

func BlockCreate(b []byte) []byte {
	in := new(pb.BlockCreateRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.BlockCreateResponse{Error: &pb.BlockCreateResponse_Error{Code: pb.BlockCreateResponse_Error_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockCreate(in).Marshal()
	return resp
}

func BlockUpdate(b []byte) []byte {
	in := new(pb.BlockUpdateRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.BlockUpdateResponse{Error: &pb.BlockUpdateResponse_Error{Code: pb.BlockUpdateResponse_Error_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockUpdate(in).Marshal()
	return resp
}

func CommandAsync(cmd string, data []byte, callback func(data []byte)) {
	go func() {
		var cd []byte
		switch cmd {
		case "WalletCreate":
			cd = WalletCreate(data)
		case "WalletRecover":
			cd = WalletRecover(data)
		case "AccountRecover":
			cd = AccountRecover(data)
		case "AccountCreate":
			cd = AccountCreate(data)
		case "AccountSelect":
			cd = AccountSelect(data)
		case "ImageGetBlob":
			cd = ImageGetBlob(data)
		case "VersionGet":
			cd = VersionGet(data)
		case "LogSend":
			cd = LogSend(data)
		case "BlockOpen":
			cd = BlockOpen(data)
		case "BlockCreate":
			cd = BlockCreate(data)
		case "BlockUpdate":
			cd = BlockUpdate(data)
		default:
			log.Errorf("unknown command type: %s\n", cmd)
		}
		if callback != nil {
			callback(cd)
		}
	}()
}

type MessageHandler interface {
	Handle(b []byte)
}

func CommandMobile(cmd string, data []byte, callback MessageHandler) {
	CommandAsync(cmd, data, callback.Handle)
}
