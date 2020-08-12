// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/extensions/transport_sockets/quic/v3/quic_transport.proto

package envoy_extensions_transport_sockets_quic_v3

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/extensions/transport_sockets/tls/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type QuicDownstreamTransport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DownstreamTlsContext *v3.DownstreamTlsContext `protobuf:"bytes,1,opt,name=downstream_tls_context,json=downstreamTlsContext,proto3" json:"downstream_tls_context,omitempty"`
}

func (x *QuicDownstreamTransport) Reset() {
	*x = QuicDownstreamTransport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuicDownstreamTransport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuicDownstreamTransport) ProtoMessage() {}

func (x *QuicDownstreamTransport) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuicDownstreamTransport.ProtoReflect.Descriptor instead.
func (*QuicDownstreamTransport) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_rawDescGZIP(), []int{0}
}

func (x *QuicDownstreamTransport) GetDownstreamTlsContext() *v3.DownstreamTlsContext {
	if x != nil {
		return x.DownstreamTlsContext
	}
	return nil
}

type QuicUpstreamTransport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpstreamTlsContext *v3.UpstreamTlsContext `protobuf:"bytes,1,opt,name=upstream_tls_context,json=upstreamTlsContext,proto3" json:"upstream_tls_context,omitempty"`
}

func (x *QuicUpstreamTransport) Reset() {
	*x = QuicUpstreamTransport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuicUpstreamTransport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuicUpstreamTransport) ProtoMessage() {}

func (x *QuicUpstreamTransport) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuicUpstreamTransport.ProtoReflect.Descriptor instead.
func (*QuicUpstreamTransport) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_rawDescGZIP(), []int{1}
}

func (x *QuicUpstreamTransport) GetUpstreamTlsContext() *v3.UpstreamTlsContext {
	if x != nil {
		return x.UpstreamTlsContext
	}
	return nil
}

var File_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto protoreflect.FileDescriptor

var file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x2f, 0x71, 0x75, 0x69, 0x63, 0x2f, 0x76, 0x33, 0x2f, 0x71, 0x75, 0x69,
	0x63, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x2a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x71, 0x75, 0x69, 0x63, 0x2e, 0x76, 0x33, 0x1a, 0x33, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x2f, 0x74, 0x6c, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x74, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01,
	0x0a, 0x17, 0x51, 0x75, 0x69, 0x63, 0x44, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x7f, 0x0a, 0x16, 0x64, 0x6f, 0x77,
	0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x74,
	0x6c, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x54, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x14, 0x64, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x54, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x92, 0x01, 0x0a, 0x15, 0x51,
	0x75, 0x69, 0x63, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x79, 0x0a, 0x14, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x5f, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x55,
	0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x12, 0x75, 0x70, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x42,
	0x58, 0x0a, 0x38, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x2e, 0x71, 0x75, 0x69, 0x63, 0x2e, 0x76, 0x33, 0x42, 0x12, 0x51, 0x75, 0x69,
	0x63, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_rawDescOnce sync.Once
	file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_rawDescData = file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_rawDesc
)

func file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_rawDescGZIP() []byte {
	file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_rawDescData)
	})
	return file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_rawDescData
}

var file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_goTypes = []interface{}{
	(*QuicDownstreamTransport)(nil), // 0: envoy.extensions.transport_sockets.quic.v3.QuicDownstreamTransport
	(*QuicUpstreamTransport)(nil),   // 1: envoy.extensions.transport_sockets.quic.v3.QuicUpstreamTransport
	(*v3.DownstreamTlsContext)(nil), // 2: envoy.extensions.transport_sockets.tls.v3.DownstreamTlsContext
	(*v3.UpstreamTlsContext)(nil),   // 3: envoy.extensions.transport_sockets.tls.v3.UpstreamTlsContext
}
var file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.transport_sockets.quic.v3.QuicDownstreamTransport.downstream_tls_context:type_name -> envoy.extensions.transport_sockets.tls.v3.DownstreamTlsContext
	3, // 1: envoy.extensions.transport_sockets.quic.v3.QuicUpstreamTransport.upstream_tls_context:type_name -> envoy.extensions.transport_sockets.tls.v3.UpstreamTlsContext
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_init() }
func file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_init() {
	if File_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuicDownstreamTransport); i {
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
		file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuicUpstreamTransport); i {
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
			RawDescriptor: file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_msgTypes,
	}.Build()
	File_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto = out.File
	file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_rawDesc = nil
	file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_goTypes = nil
	file_envoy_extensions_transport_sockets_quic_v3_quic_transport_proto_depIdxs = nil
}
