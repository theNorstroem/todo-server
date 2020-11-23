// Code generated by furo-proto-gen. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: auth/reqmsgs.proto

package authpb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// request message for CreateAuthSession
type CreateAuthSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body with auth.Credentials
	Body *Credentials `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *CreateAuthSessionRequest) Reset() {
	*x = CreateAuthSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_reqmsgs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAuthSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthSessionRequest) ProtoMessage() {}

func (x *CreateAuthSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_reqmsgs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthSessionRequest.ProtoReflect.Descriptor instead.
func (*CreateAuthSessionRequest) Descriptor() ([]byte, []int) {
	return file_auth_reqmsgs_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAuthSessionRequest) GetBody() *Credentials {
	if x != nil {
		return x.Body
	}
	return nil
}

// request message for DeleteAuthSession
type DeleteAuthSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body with google.protobuf.Empty
	Body *emptypb.Empty `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *DeleteAuthSessionRequest) Reset() {
	*x = DeleteAuthSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_reqmsgs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAuthSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAuthSessionRequest) ProtoMessage() {}

func (x *DeleteAuthSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_reqmsgs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAuthSessionRequest.ProtoReflect.Descriptor instead.
func (*DeleteAuthSessionRequest) Descriptor() ([]byte, []int) {
	return file_auth_reqmsgs_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteAuthSessionRequest) GetBody() *emptypb.Empty {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_auth_reqmsgs_proto protoreflect.FileDescriptor

var file_auth_reqmsgs_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72, 0x65, 0x71, 0x6d, 0x73, 0x67, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x0f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x46, 0x0a, 0x18, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x42, 0x5b, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x75, 0x72, 0x6f, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x61, 0x75, 0x74, 0x68, 0x42, 0x0c, 0x52, 0x65, 0x71, 0x6d, 0x73, 0x67,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x4e, 0x6f, 0x72, 0x73, 0x74, 0x72, 0x6f, 0x65,
	0x6d, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x64, 0x69, 0x73,
	0x74, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_reqmsgs_proto_rawDescOnce sync.Once
	file_auth_reqmsgs_proto_rawDescData = file_auth_reqmsgs_proto_rawDesc
)

func file_auth_reqmsgs_proto_rawDescGZIP() []byte {
	file_auth_reqmsgs_proto_rawDescOnce.Do(func() {
		file_auth_reqmsgs_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_reqmsgs_proto_rawDescData)
	})
	return file_auth_reqmsgs_proto_rawDescData
}

var file_auth_reqmsgs_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_auth_reqmsgs_proto_goTypes = []interface{}{
	(*CreateAuthSessionRequest)(nil), // 0: auth.CreateAuthSessionRequest
	(*DeleteAuthSessionRequest)(nil), // 1: auth.DeleteAuthSessionRequest
	(*Credentials)(nil),              // 2: auth.Credentials
	(*emptypb.Empty)(nil),            // 3: google.protobuf.Empty
}
var file_auth_reqmsgs_proto_depIdxs = []int32{
	2, // 0: auth.CreateAuthSessionRequest.body:type_name -> auth.Credentials
	3, // 1: auth.DeleteAuthSessionRequest.body:type_name -> google.protobuf.Empty
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_auth_reqmsgs_proto_init() }
func file_auth_reqmsgs_proto_init() {
	if File_auth_reqmsgs_proto != nil {
		return
	}
	file_auth_auth_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_auth_reqmsgs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAuthSessionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_reqmsgs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAuthSessionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_reqmsgs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_reqmsgs_proto_goTypes,
		DependencyIndexes: file_auth_reqmsgs_proto_depIdxs,
		MessageInfos:      file_auth_reqmsgs_proto_msgTypes,
	}.Build()
	File_auth_reqmsgs_proto = out.File
	file_auth_reqmsgs_proto_rawDesc = nil
	file_auth_reqmsgs_proto_goTypes = nil
	file_auth_reqmsgs_proto_depIdxs = nil
}
