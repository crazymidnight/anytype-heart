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
	// 806 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x97, 0xd1, 0x4e, 0x13, 0x41,
	0x14, 0x86, 0x21, 0x26, 0x12, 0x47, 0x24, 0x3a, 0xc6, 0x84, 0x90, 0x50, 0x10, 0x50, 0xa1, 0xc6,
	0x21, 0xd1, 0x27, 0x80, 0x1a, 0x4a, 0x13, 0x40, 0xa4, 0x08, 0x09, 0xd1, 0x98, 0xed, 0xf6, 0x50,
	0xd6, 0x6e, 0x67, 0xd6, 0x9d, 0xa1, 0xda, 0xb7, 0xf0, 0xb1, 0xbc, 0xe4, 0xd2, 0x4b, 0x03, 0x6f,
	0xe0, 0x13, 0x98, 0x99, 0x9d, 0xdd, 0x76, 0x76, 0x67, 0xa7, 0x7b, 0x01, 0x24, 0xfb, 0x7f, 0xe7,
	0xff, 0xcf, 0xd9, 0x3d, 0x3b, 0xb4, 0x68, 0x25, 0xea, 0x6c, 0x47, 0x31, 0x13, 0x8c, 0x6f, 0x73,
	0x88, 0x87, 0x81, 0x0f, 0xe9, 0x5f, 0xa2, 0x2e, 0xe3, 0x39, 0x8f, 0x8e, 0xc4, 0x28, 0x82, 0xa5,
	0xc5, 0x31, 0xe9, 0xb3, 0xc1, 0xc0, 0xa3, 0x5d, 0x9e, 0x20, 0x6f, 0xff, 0x2d, 0xa2, 0x85, 0x46,
	0x18, 0x00, 0x15, 0x0d, 0x2d, 0xe0, 0x73, 0x34, 0x7f, 0xee, 0x85, 0x21, 0x88, 0x46, 0x0c, 0x9e,
	0x00, 0xbc, 0x46, 0xb4, 0x0d, 0x39, 0x89, 0x7c, 0x92, 0x48, 0x24, 0xd1, 0xc8, 0x09, 0x7c, 0xbf,
	0x06, 0x2e, 0x96, 0xd6, 0x9d, 0x0c, 0x8f, 0x18, 0xe5, 0x80, 0x2f, 0xd0, 0xa3, 0x44, 0x39, 0x01,
	0x9f, 0x0d, 0x21, 0xc6, 0xd6, 0x2a, 0x2d, 0x66, 0xd6, 0x1b, 0x6e, 0x48, 0x7b, 0x7f, 0x41, 0x0b,
	0x3b, 0xbe, 0xcf, 0xae, 0x69, 0x66, 0x6e, 0xd6, 0x69, 0xb1, 0xe0, 0xfe, 0x62, 0x0a, 0x35, 0x6e,
	0x5d, 0x6b, 0xfa, 0xa6, 0xac, 0x5b, 0xeb, 0x72, 0x77, 0x65, 0xc3, 0x0d, 0x15, 0xbc, 0xdb, 0x10,
	0x82, 0x2f, 0x4a, 0xbc, 0x13, 0x71, 0x8a, 0x77, 0x06, 0x69, 0x6f, 0x1f, 0xcd, 0xb7, 0x06, 0x5e,
	0x0f, 0x9a, 0x20, 0x76, 0x43, 0xd6, 0xc1, 0x9b, 0x46, 0x55, 0x2b, 0xba, 0xe4, 0x44, 0xe9, 0xa4,
	0x09, 0x82, 0x48, 0x22, 0xf3, 0xdf, 0xaa, 0x40, 0xea, 0x90, 0x8f, 0x08, 0x9d, 0x41, 0xcc, 0x03,
	0x46, 0x9b, 0x20, 0xf0, 0xaa, 0x51, 0xa8, 0x05, 0x55, 0x95, 0x5a, 0x3f, 0x77, 0x10, 0xda, 0x72,
	0x1f, 0xcd, 0x1d, 0xb0, 0x5e, 0x1b, 0x68, 0x17, 0x2f, 0x1b, 0xf4, 0x01, 0xeb, 0x11, 0x79, 0x39,
	0x33, 0xab, 0x95, 0xc9, 0xda, 0xe9, 0x08, 0x3d, 0x68, 0x30, 0x7a, 0x19, 0xf4, 0x64, 0x6f, 0x2b,
	0x06, 0x9c, 0x5c, 0x37, 0x5a, 0x5b, 0x2d, 0x07, 0xb4, 0xdf, 0x29, 0x7a, 0xb8, 0x1b, 0x32, 0xbf,
	0xff, 0x29, 0x0a, 0x99, 0xd7, 0xc5, 0xe6, 0x2c, 0x4a, 0x21, 0x89, 0x94, 0x79, 0xae, 0xb9, 0x90,
	0x71, 0x97, 0x4a, 0xf8, 0x10, 0x01, 0xcd, 0x75, 0x99, 0x14, 0x48, 0xa1, 0xa4, 0x4b, 0x03, 0xc8,
	0x75, 0xa9, 0xb7, 0xd5, 0xd6, 0x65, 0x6e, 0x57, 0xd7, 0x5c, 0x48, 0x7e, 0x76, 0x1a, 0x06, 0xb4,
	0x6f, 0x9f, 0x5d, 0x49, 0xee, 0xd9, 0x53, 0x64, 0xbc, 0x3e, 0x49, 0xaf, 0x21, 0xe3, 0x80, 0x6d,
	0xb3, 0x29, 0xa5, 0x64, 0x7d, 0x4c, 0x62, 0xfc, 0x4a, 0xa9, 0xeb, 0xef, 0xd9, 0x0f, 0xaa, 0x1e,
	0xd3, 0xba, 0xa5, 0x26, 0x15, 0x4b, 0x5e, 0xa9, 0x02, 0xa4, 0xbd, 0x3f, 0x6b, 0xef, 0x26, 0x88,
	0x43, 0x2f, 0xee, 0x73, 0x6c, 0x2b, 0x93, 0x2b, 0xa3, 0xd4, 0x92, 0x83, 0xa6, 0x48, 0x69, 0x77,
	0x40, 0x8f, 0x95, 0xb6, 0x1f, 0x70, 0xc1, 0xe2, 0xd1, 0x21, 0x1b, 0x02, 0x7e, 0x65, 0x29, 0xd5,
	0x3a, 0x91, 0x40, 0x96, 0xb1, 0x39, 0x1d, 0xd4, 0x31, 0x5f, 0xd1, 0x82, 0x92, 0xdb, 0x20, 0xf6,
	0x02, 0x08, 0xbb, 0x1c, 0xdb, 0xfa, 0x6b, 0x83, 0x20, 0x89, 0x9c, 0x45, 0xbc, 0x9c, 0x86, 0xe9,
	0x00, 0x8a, 0x9e, 0xa6, 0x01, 0xc7, 0x10, 0x0f, 0x02, 0x2e, 0xdf, 0x71, 0x8e, 0xeb, 0x25, 0xe5,
	0x13, 0x4c, 0x16, 0xf5, 0xba, 0x12, 0xab, 0xf3, 0xfa, 0x08, 0xa7, 0x79, 0x2d, 0xbe, 0x13, 0xfb,
	0x57, 0xc1, 0x10, 0xba, 0x78, 0xab, 0xc4, 0x62, 0x8c, 0x64, 0x69, 0xf5, 0x2a, 0x68, 0x6e, 0xbd,
	0x0e, 0x02, 0x2e, 0xd4, 0x13, 0xb2, 0xac, 0x97, 0xd4, 0xcc, 0xa7, 0xb3, 0xe1, 0x86, 0xb4, 0xf7,
	0x10, 0x3d, 0xcb, 0xa4, 0x36, 0x88, 0x53, 0xf8, 0x29, 0xda, 0x62, 0x14, 0x02, 0x7e, 0x53, 0x52,
	0x2e, 0x9b, 0x94, 0x14, 0x51, 0x58, 0x96, 0x46, 0xaa, 0xe2, 0x3a, 0xb7, 0xa7, 0x17, 0x4f, 0x67,
	0xca, 0x1f, 0xbc, 0x59, 0x72, 0x4f, 0x54, 0xbd, 0xfa, 0x65, 0xff, 0x6f, 0x61, 0x27, 0x75, 0xd0,
	0x37, 0xf4, 0x64, 0x32, 0x28, 0x19, 0xce, 0x59, 0x6f, 0x0e, 0x56, 0xaf, 0x82, 0xda, 0xb3, 0x1a,
	0x57, 0xe0, 0xf7, 0xdd, 0x59, 0x0a, 0xa9, 0x96, 0x95, 0xa2, 0xb9, 0x23, 0xbc, 0xc1, 0xa2, 0x91,
	0xf5, 0x08, 0x97, 0x82, 0xf3, 0x08, 0xd7, 0x40, 0xee, 0x58, 0x3c, 0xf6, 0xb8, 0xb0, 0x1f, 0x8b,
	0x4a, 0x71, 0x1e, 0x8b, 0x29, 0x51, 0x7c, 0xc6, 0x7b, 0x41, 0x08, 0x47, 0xde, 0x00, 0x4a, 0x9f,
	0xb1, 0x04, 0x88, 0x24, 0xa6, 0x3e, 0xe3, 0x49, 0xb2, 0x78, 0xdf, 0xd5, 0x87, 0x06, 0x95, 0x54,
	0xfa, 0x32, 0xaa, 0x8f, 0x15, 0x46, 0x54, 0xbd, 0x0a, 0xaa, 0xb3, 0x06, 0x13, 0x6f, 0xbe, 0x94,
	0xcf, 0x83, 0xae, 0xb8, 0xc2, 0x6e, 0x07, 0xc5, 0x4c, 0x3d, 0x68, 0x4c, 0xb6, 0x38, 0xda, 0x59,
	0xd0, 0x05, 0xe6, 0x1c, 0x4d, 0x11, 0xd5, 0x46, 0x33, 0xd0, 0xe2, 0x68, 0x4a, 0x76, 0x8f, 0x96,
	0x38, 0x54, 0x1b, 0xcd, 0x64, 0x8b, 0xeb, 0xd1, 0xf2, 0x19, 0x75, 0xae, 0x87, 0x04, 0xaa, 0xad,
	0xc7, 0x24, 0x99, 0x04, 0xed, 0x2e, 0xff, 0xbe, 0xad, 0xcd, 0xde, 0xdc, 0xd6, 0x66, 0xff, 0xde,
	0xd6, 0x66, 0x7f, 0xdd, 0xd5, 0x66, 0x6e, 0xee, 0x6a, 0x33, 0x7f, 0xee, 0x6a, 0x33, 0x17, 0xf7,
	0xc2, 0xa0, 0xd3, 0xb9, 0xaf, 0xbe, 0x9a, 0xbc, 0xfb, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x84,
	0xe7, 0x00, 0xe0, 0x0c, 0x00, 0x00,
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
	ConfigGet(*pb.RpcConfigGetRequest) *pb.RpcConfigGetResponse
	BlockUpload(*pb.RpcBlockUploadRequest) *pb.RpcBlockUploadResponse
	BlockOpen(*pb.RpcBlockOpenRequest) *pb.RpcBlockOpenResponse
	BlockCreate(*pb.RpcBlockCreateRequest) *pb.RpcBlockCreateResponse
	BlockUnlink(*pb.RpcBlockUnlinkRequest) *pb.RpcBlockUnlinkResponse
	BlockClose(*pb.RpcBlockCloseRequest) *pb.RpcBlockCloseResponse
	BlockDownload(*pb.RpcBlockDownloadRequest) *pb.RpcBlockDownloadResponse
	BlockGetMarks(*pb.RpcBlockGetMarksRequest) *pb.RpcBlockGetMarksResponse
	BlockHistoryMove(*pb.RpcBlockHistoryMoveRequest) *pb.RpcBlockHistoryMoveResponse
	BlockSetFields(*pb.RpcBlockSetFieldsRequest) *pb.RpcBlockSetFieldsResponse
	BlockSetPermissions(*pb.RpcBlockSetPermissionsRequest) *pb.RpcBlockSetPermissionsResponse
	BlockSetIsArchived(*pb.RpcBlockSetIsArchivedRequest) *pb.RpcBlockSetIsArchivedResponse
	BlockListMove(*pb.RpcBlockListMoveRequest) *pb.RpcBlockListMoveResponse
	BlockListSetTextStyle(*pb.RpcBlockListSetTextStyleRequest) *pb.RpcBlockListSetTextStyleResponse
	BlockSetTextText(*pb.RpcBlockSetTextTextRequest) *pb.RpcBlockSetTextTextResponse
	BlockSetTextStyle(*pb.RpcBlockSetTextStyleRequest) *pb.RpcBlockSetTextStyleResponse
	BlockSetTextCheck(*pb.RpcBlockSetTextCheckRequest) *pb.RpcBlockSetTextCheckResponse
	BlockCopy(*pb.RpcBlockCopyRequest) *pb.RpcBlockCopyResponse
	BlockPaste(*pb.RpcBlockPasteRequest) *pb.RpcBlockPasteResponse
	BlockSetFileName(*pb.RpcBlockSetFileNameRequest) *pb.RpcBlockSetFileNameResponse
	BlockSetImageName(*pb.RpcBlockSetImageNameRequest) *pb.RpcBlockSetImageNameResponse
	BlockSetImageWidth(*pb.RpcBlockSetImageWidthRequest) *pb.RpcBlockSetImageWidthResponse
	BlockSetVideoName(*pb.RpcBlockSetVideoNameRequest) *pb.RpcBlockSetVideoNameResponse
	BlockSetVideoWidth(*pb.RpcBlockSetVideoWidthRequest) *pb.RpcBlockSetVideoWidthResponse
	BlockSetIconName(*pb.RpcBlockSetIconNameRequest) *pb.RpcBlockSetIconNameResponse
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

func ConfigGet(b []byte) []byte {
	in := new(pb.RpcConfigGetRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcConfigGetResponse{Error: &pb.RpcConfigGetResponseError{Code: pb.RpcConfigGetResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.ConfigGet(in).Marshal()
	return resp
}

func BlockUpload(b []byte) []byte {
	in := new(pb.RpcBlockUploadRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockUploadResponse{Error: &pb.RpcBlockUploadResponseError{Code: pb.RpcBlockUploadResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockUpload(in).Marshal()
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

func BlockUnlink(b []byte) []byte {
	in := new(pb.RpcBlockUnlinkRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockUnlinkResponse{Error: &pb.RpcBlockUnlinkResponseError{Code: pb.RpcBlockUnlinkResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockUnlink(in).Marshal()
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

func BlockDownload(b []byte) []byte {
	in := new(pb.RpcBlockDownloadRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockDownloadResponse{Error: &pb.RpcBlockDownloadResponseError{Code: pb.RpcBlockDownloadResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockDownload(in).Marshal()
	return resp
}

func BlockGetMarks(b []byte) []byte {
	in := new(pb.RpcBlockGetMarksRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockGetMarksResponse{Error: &pb.RpcBlockGetMarksResponseError{Code: pb.RpcBlockGetMarksResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockGetMarks(in).Marshal()
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

func BlockSetFields(b []byte) []byte {
	in := new(pb.RpcBlockSetFieldsRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetFieldsResponse{Error: &pb.RpcBlockSetFieldsResponseError{Code: pb.RpcBlockSetFieldsResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetFields(in).Marshal()
	return resp
}

func BlockSetPermissions(b []byte) []byte {
	in := new(pb.RpcBlockSetPermissionsRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetPermissionsResponse{Error: &pb.RpcBlockSetPermissionsResponseError{Code: pb.RpcBlockSetPermissionsResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetPermissions(in).Marshal()
	return resp
}

func BlockSetIsArchived(b []byte) []byte {
	in := new(pb.RpcBlockSetIsArchivedRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetIsArchivedResponse{Error: &pb.RpcBlockSetIsArchivedResponseError{Code: pb.RpcBlockSetIsArchivedResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetIsArchived(in).Marshal()
	return resp
}

func BlockListMove(b []byte) []byte {
	in := new(pb.RpcBlockListMoveRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockListMoveResponse{Error: &pb.RpcBlockListMoveResponseError{Code: pb.RpcBlockListMoveResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockListMove(in).Marshal()
	return resp
}

func BlockListSetTextStyle(b []byte) []byte {
	in := new(pb.RpcBlockListSetTextStyleRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockListSetTextStyleResponse{Error: &pb.RpcBlockListSetTextStyleResponseError{Code: pb.RpcBlockListSetTextStyleResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockListSetTextStyle(in).Marshal()
	return resp
}

func BlockSetTextText(b []byte) []byte {
	in := new(pb.RpcBlockSetTextTextRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetTextTextResponse{Error: &pb.RpcBlockSetTextTextResponseError{Code: pb.RpcBlockSetTextTextResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetTextText(in).Marshal()
	return resp
}

func BlockSetTextStyle(b []byte) []byte {
	in := new(pb.RpcBlockSetTextStyleRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetTextStyleResponse{Error: &pb.RpcBlockSetTextStyleResponseError{Code: pb.RpcBlockSetTextStyleResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetTextStyle(in).Marshal()
	return resp
}

func BlockSetTextCheck(b []byte) []byte {
	in := new(pb.RpcBlockSetTextCheckRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetTextCheckResponse{Error: &pb.RpcBlockSetTextCheckResponseError{Code: pb.RpcBlockSetTextCheckResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetTextCheck(in).Marshal()
	return resp
}

func BlockCopy(b []byte) []byte {
	in := new(pb.RpcBlockCopyRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockCopyResponse{Error: &pb.RpcBlockCopyResponseError{Code: pb.RpcBlockCopyResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockCopy(in).Marshal()
	return resp
}

func BlockPaste(b []byte) []byte {
	in := new(pb.RpcBlockPasteRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockPasteResponse{Error: &pb.RpcBlockPasteResponseError{Code: pb.RpcBlockPasteResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockPaste(in).Marshal()
	return resp
}

func BlockSetFileName(b []byte) []byte {
	in := new(pb.RpcBlockSetFileNameRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetFileNameResponse{Error: &pb.RpcBlockSetFileNameResponseError{Code: pb.RpcBlockSetFileNameResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetFileName(in).Marshal()
	return resp
}

func BlockSetImageName(b []byte) []byte {
	in := new(pb.RpcBlockSetImageNameRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetImageNameResponse{Error: &pb.RpcBlockSetImageNameResponseError{Code: pb.RpcBlockSetImageNameResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetImageName(in).Marshal()
	return resp
}

func BlockSetImageWidth(b []byte) []byte {
	in := new(pb.RpcBlockSetImageWidthRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetImageWidthResponse{Error: &pb.RpcBlockSetImageWidthResponseError{Code: pb.RpcBlockSetImageWidthResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetImageWidth(in).Marshal()
	return resp
}

func BlockSetVideoName(b []byte) []byte {
	in := new(pb.RpcBlockSetVideoNameRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetVideoNameResponse{Error: &pb.RpcBlockSetVideoNameResponseError{Code: pb.RpcBlockSetVideoNameResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetVideoName(in).Marshal()
	return resp
}

func BlockSetVideoWidth(b []byte) []byte {
	in := new(pb.RpcBlockSetVideoWidthRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetVideoWidthResponse{Error: &pb.RpcBlockSetVideoWidthResponseError{Code: pb.RpcBlockSetVideoWidthResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetVideoWidth(in).Marshal()
	return resp
}

func BlockSetIconName(b []byte) []byte {
	in := new(pb.RpcBlockSetIconNameRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetIconNameResponse{Error: &pb.RpcBlockSetIconNameResponseError{Code: pb.RpcBlockSetIconNameResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetIconName(in).Marshal()
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
		case "ConfigGet":
			cd = ConfigGet(data)
		case "BlockUpload":
			cd = BlockUpload(data)
		case "BlockOpen":
			cd = BlockOpen(data)
		case "BlockCreate":
			cd = BlockCreate(data)
		case "BlockUnlink":
			cd = BlockUnlink(data)
		case "BlockClose":
			cd = BlockClose(data)
		case "BlockDownload":
			cd = BlockDownload(data)
		case "BlockGetMarks":
			cd = BlockGetMarks(data)
		case "BlockHistoryMove":
			cd = BlockHistoryMove(data)
		case "BlockSetFields":
			cd = BlockSetFields(data)
		case "BlockSetPermissions":
			cd = BlockSetPermissions(data)
		case "BlockSetIsArchived":
			cd = BlockSetIsArchived(data)
		case "BlockListMove":
			cd = BlockListMove(data)
		case "BlockListSetTextStyle":
			cd = BlockListSetTextStyle(data)
		case "BlockSetTextText":
			cd = BlockSetTextText(data)
		case "BlockSetTextStyle":
			cd = BlockSetTextStyle(data)
		case "BlockSetTextCheck":
			cd = BlockSetTextCheck(data)
		case "BlockCopy":
			cd = BlockCopy(data)
		case "BlockPaste":
			cd = BlockPaste(data)
		case "BlockSetFileName":
			cd = BlockSetFileName(data)
		case "BlockSetImageName":
			cd = BlockSetImageName(data)
		case "BlockSetImageWidth":
			cd = BlockSetImageWidth(data)
		case "BlockSetVideoName":
			cd = BlockSetVideoName(data)
		case "BlockSetVideoWidth":
			cd = BlockSetVideoWidth(data)
		case "BlockSetIconName":
			cd = BlockSetIconName(data)
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
