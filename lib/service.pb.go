// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/protos/service/service.proto

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

func init() { proto.RegisterFile("pb/protos/service/service.proto", fileDescriptor_93a29dc403579097) }

var fileDescriptor_93a29dc403579097 = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdf, 0x6a, 0x13, 0x41,
	0x14, 0xc6, 0x13, 0x04, 0x8b, 0x63, 0x2c, 0x32, 0x57, 0x52, 0xe8, 0xf6, 0xaf, 0xd8, 0x8a, 0x4e,
	0x41, 0x9f, 0xa0, 0x89, 0x90, 0x16, 0x5a, 0xc5, 0x44, 0x2d, 0x14, 0xbc, 0xd8, 0x4c, 0x8e, 0x61,
	0xe9, 0x64, 0x66, 0x9c, 0x99, 0x06, 0xf2, 0x16, 0x3e, 0x96, 0x97, 0xbd, 0xf4, 0x52, 0x12, 0xf0,
	0x39, 0x64, 0x67, 0x4e, 0x27, 0x6c, 0xba, 0xbb, 0xc9, 0x55, 0xe0, 0x7c, 0xbf, 0xef, 0x77, 0x66,
	0x77, 0x96, 0x90, 0x1d, 0x3d, 0x38, 0xd1, 0x46, 0x39, 0x65, 0x4f, 0x2c, 0x98, 0x49, 0xc6, 0xe1,
	0xfe, 0x97, 0xf9, 0x31, 0xdd, 0x48, 0xe5, 0xd4, 0x4d, 0x35, 0x6c, 0xbd, 0x58, 0x90, 0x5c, 0x8d,
	0xc7, 0xa9, 0x1c, 0xda, 0x80, 0xbc, 0xfb, 0x47, 0xc8, 0x66, 0x47, 0x64, 0x20, 0x5d, 0x07, 0x03,
	0x7a, 0x45, 0x5a, 0x57, 0xa9, 0x10, 0xe0, 0x3a, 0x06, 0x52, 0x07, 0x74, 0x9f, 0xa1, 0x86, 0xf5,
	0x34, 0x67, 0x21, 0x62, 0x21, 0x63, 0x3d, 0xf8, 0x79, 0x0b, 0xd6, 0x6d, 0x1d, 0xd4, 0x32, 0x56,
	0x2b, 0x69, 0x81, 0x5e, 0x93, 0x67, 0x21, 0xe9, 0x01, 0x57, 0x13, 0x30, 0xb4, 0xb4, 0x85, 0x61,
	0x54, 0x1f, 0xd6, 0x43, 0xe8, 0xfe, 0x4e, 0x36, 0x4f, 0x39, 0x57, 0xb7, 0x32, 0xca, 0x8b, 0x3d,
	0x0c, 0x1f, 0xd8, 0x5f, 0xae, 0xa0, 0x16, 0x47, 0xc7, 0x0c, 0x5f, 0xca, 0x41, 0x69, 0x6f, 0xe9,
	0xad, 0x1c, 0xd6, 0x43, 0x0f, 0xdc, 0x7d, 0x10, 0xc0, 0x5d, 0x85, 0x3b, 0x84, 0x2b, 0xdc, 0x11,
	0x42, 0x37, 0x27, 0xad, 0xf3, 0x71, 0x3a, 0x82, 0x2e, 0xb8, 0xb6, 0x50, 0x03, 0x7a, 0x54, 0x68,
	0x9d, 0xeb, 0x1f, 0x96, 0xf9, 0x9c, 0x75, 0xc1, 0xb1, 0x9c, 0x88, 0xfe, 0xe3, 0x35, 0x48, 0x5c,
	0xf2, 0x99, 0x90, 0x6f, 0x60, 0x6c, 0xa6, 0x64, 0x17, 0x1c, 0xdd, 0x2d, 0x14, 0x31, 0xf0, 0xad,
	0x7b, 0xf5, 0x5e, 0x0d, 0x81, 0xca, 0x33, 0xb2, 0x71, 0xa1, 0x46, 0x7d, 0x90, 0x43, 0xba, 0x5d,
	0xa0, 0x2f, 0xd4, 0x88, 0xe5, 0xe3, 0x28, 0x4b, 0xaa, 0x62, 0x34, 0x7d, 0x24, 0x4f, 0xda, 0x42,
	0xf1, 0x9b, 0x4f, 0x1a, 0x24, 0xdd, 0x29, 0xc0, 0x7e, 0xce, 0xf2, 0x20, 0xda, 0x76, 0xab, 0x01,
	0xf4, 0x7d, 0x21, 0x4f, 0xfd, 0x18, 0xbf, 0x83, 0xbd, 0x92, 0xc2, 0xd2, 0x57, 0xb0, 0x5f, 0x87,
	0x2c, 0x59, 0xbf, 0xea, 0x61, 0x95, 0x35, 0x44, 0xb5, 0xd6, 0x88, 0x2c, 0x2e, 0x26, 0x9c, 0x55,
	0x28, 0x0b, 0xb4, 0xec, 0xd9, 0x7c, 0x52, 0x71, 0x31, 0x45, 0x02, 0x95, 0x40, 0x5a, 0x7e, 0x7e,
	0x99, 0x9a, 0x9b, 0x3e, 0x38, 0x7a, 0x5c, 0x52, 0x39, 0xe5, 0x2e, 0xbf, 0x50, 0x44, 0xa2, 0xfd,
	0xf5, 0x3a, 0x28, 0xae, 0xc9, 0xf0, 0xe4, 0xf6, 0x83, 0x51, 0x9a, 0xbe, 0xa9, 0x6e, 0x2e, 0xa8,
	0xb8, 0xe7, 0xed, 0x9a, 0x74, 0x7c, 0xa2, 0xe7, 0x7e, 0x7c, 0x96, 0x59, 0xa7, 0xcc, 0xf4, 0x52,
	0x4d, 0x80, 0xbe, 0x2a, 0x51, 0x60, 0xce, 0x72, 0x20, 0xee, 0x3a, 0x5a, 0x0d, 0x86, 0x35, 0xed,
	0xed, 0xdf, 0xb3, 0xa4, 0x79, 0x37, 0x4b, 0x9a, 0x7f, 0x67, 0x49, 0xf3, 0xd7, 0x3c, 0x69, 0xdc,
	0xcd, 0x93, 0xc6, 0x9f, 0x79, 0xd2, 0xb8, 0x7e, 0x24, 0xb2, 0xc1, 0xe0, 0xb1, 0xff, 0x3b, 0x7e,
	0xff, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x38, 0xad, 0x19, 0xd4, 0x05, 0x00, 0x00,
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the gomobile package it is being compiled against.

// ClientCommandsHandler is the handler API for ClientCommands service.
var clientCommandsHandler ClientCommandsHandler

type ClientCommandsHandler interface {
	WalletCreate(*pb.RpcWalletCreateRequest) *pb.RpcWalletCreateResponse
	WalletRecover(*pb.RpcWalletRecoverRequest) *pb.RpcWalletRecoverResponse
	AccountRecover(*pb.RpcAccountRecoverRequest) *pb.RpcAccountRecoverResponse
	AccountCreate(*pb.RpcAccountCreateRequest) *pb.RpcAccountCreateResponse
	AccountSelect(*pb.RpcAccountSelectRequest) *pb.RpcAccountSelectResponse
	ImageGetBlob(*pb.RpcIpfsImageGetBlobRequest) *pb.RpcIpfsImageGetBlobResponse
	VersionGet(*pb.RpcVersionGetRequest) *pb.RpcVersionGetResponse
	LogSend(*pb.RpcLogSendRequest) *pb.RpcLogSendResponse
	BlockOpen(*pb.RpcBlockOpenRequest) *pb.RpcBlockOpenResponse
	BlockCreate(*pb.RpcBlockCreateRequest) *pb.RpcBlockCreateResponse
	BlockUpdate(*pb.RpcBlockUpdateRequest) *pb.RpcBlockUpdateResponse
	// TODO: rpc BlockDelete (anytype.Rpc.Block.Delete.Request) returns (anytype.Rpc.Block.Delete.Response);
	BlockClose(*pb.RpcBlockCloseRequest) *pb.RpcBlockCloseResponse
	BlockMarkSet(*pb.RpcBlockActionMarkSetRequest) *pb.RpcBlockActionMarkSetResponse
	BlocksDrop(*pb.RpcBlockActionBlocksDropRequest) *pb.RpcBlockActionBlocksDropResponse
	BlockHistoryMove(*pb.RpcBlockHistoryMoveRequest) *pb.RpcBlockHistoryMoveResponse
}

func registerClientCommandsHandler(srv ClientCommandsHandler) {
	clientCommandsHandler = srv
}

func WalletCreate(b []byte) []byte {
	in := new(pb.RpcWalletCreateRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcWalletCreateResponse{Error: &pb.RpcWalletCreateResponseError{Code: pb.RpcWalletCreateResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.WalletCreate(in).Marshal()
	return resp
}

func WalletRecover(b []byte) []byte {
	in := new(pb.RpcWalletRecoverRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcWalletRecoverResponse{Error: &pb.RpcWalletRecoverResponseError{Code: pb.RpcWalletRecoverResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.WalletRecover(in).Marshal()
	return resp
}

func AccountRecover(b []byte) []byte {
	in := new(pb.RpcAccountRecoverRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcAccountRecoverResponse{Error: &pb.RpcAccountRecoverResponseError{Code: pb.RpcAccountRecoverResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.AccountRecover(in).Marshal()
	return resp
}

func AccountCreate(b []byte) []byte {
	in := new(pb.RpcAccountCreateRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcAccountCreateResponse{Error: &pb.RpcAccountCreateResponseError{Code: pb.RpcAccountCreateResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.AccountCreate(in).Marshal()
	return resp
}

func AccountSelect(b []byte) []byte {
	in := new(pb.RpcAccountSelectRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcAccountSelectResponse{Error: &pb.RpcAccountSelectResponseError{Code: pb.RpcAccountSelectResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.AccountSelect(in).Marshal()
	return resp
}

func ImageGetBlob(b []byte) []byte {
	in := new(pb.RpcIpfsImageGetBlobRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcIpfsImageGetBlobResponse{Error: &pb.RpcIpfsImageGetBlobResponseError{Code: pb.RpcIpfsImageGetBlobResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.ImageGetBlob(in).Marshal()
	return resp
}

func VersionGet(b []byte) []byte {
	in := new(pb.RpcVersionGetRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcVersionGetResponse{Error: &pb.RpcVersionGetResponseError{Code: pb.RpcVersionGetResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.VersionGet(in).Marshal()
	return resp
}

func LogSend(b []byte) []byte {
	in := new(pb.RpcLogSendRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcLogSendResponse{Error: &pb.RpcLogSendResponseError{Code: pb.RpcLogSendResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.LogSend(in).Marshal()
	return resp
}

func BlockOpen(b []byte) []byte {
	in := new(pb.RpcBlockOpenRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockOpenResponse{Error: &pb.RpcBlockOpenResponseError{Code: pb.RpcBlockOpenResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockOpen(in).Marshal()
	return resp
}

func BlockCreate(b []byte) []byte {
	in := new(pb.RpcBlockCreateRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockCreateResponse{Error: &pb.RpcBlockCreateResponseError{Code: pb.RpcBlockCreateResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockCreate(in).Marshal()
	return resp
}

func BlockUpdate(b []byte) []byte {
	in := new(pb.RpcBlockUpdateRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockUpdateResponse{Error: &pb.RpcBlockUpdateResponseError{Code: pb.RpcBlockUpdateResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockUpdate(in).Marshal()
	return resp
}

func BlockClose(b []byte) []byte {
	in := new(pb.RpcBlockCloseRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockCloseResponse{Error: &pb.RpcBlockCloseResponseError{Code: pb.RpcBlockCloseResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockClose(in).Marshal()
	return resp
}

func BlockMarkSet(b []byte) []byte {
	in := new(pb.RpcBlockActionMarkSetRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockActionMarkSetResponse{Error: &pb.RpcBlockActionMarkSetResponseError{Code: pb.RpcBlockActionMarkSetResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockMarkSet(in).Marshal()
	return resp
}

func BlocksDrop(b []byte) []byte {
	in := new(pb.RpcBlockActionBlocksDropRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockActionBlocksDropResponse{Error: &pb.RpcBlockActionBlocksDropResponseError{Code: pb.RpcBlockActionBlocksDropResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlocksDrop(in).Marshal()
	return resp
}

func BlockHistoryMove(b []byte) []byte {
	in := new(pb.RpcBlockHistoryMoveRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockHistoryMoveResponse{Error: &pb.RpcBlockHistoryMoveResponseError{Code: pb.RpcBlockHistoryMoveResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockHistoryMove(in).Marshal()
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
		case "BlockClose":
			cd = BlockClose(data)
		case "BlockMarkSet":
			cd = BlockMarkSet(data)
		case "BlocksDrop":
			cd = BlocksDrop(data)
		case "BlockHistoryMove":
			cd = BlockHistoryMove(data)
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
