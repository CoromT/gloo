// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/grpc/version/version.proto

package version

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// type of gloo server instance
type GlooType int32

const (
	GlooType_Unknown GlooType = 0
	GlooType_Gateway GlooType = 1
	GlooType_Ingress GlooType = 2
	GlooType_Knative GlooType = 3
)

// Enum value maps for GlooType.
var (
	GlooType_name = map[int32]string{
		0: "Unknown",
		1: "Gateway",
		2: "Ingress",
		3: "Knative",
	}
	GlooType_value = map[string]int32{
		"Unknown": 0,
		"Gateway": 1,
		"Ingress": 2,
		"Knative": 3,
	}
)

func (x GlooType) Enum() *GlooType {
	p := new(GlooType)
	*p = x
	return p
}

func (x GlooType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GlooType) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_enumTypes[0].Descriptor()
}

func (GlooType) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_enumTypes[0]
}

func (x GlooType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GlooType.Descriptor instead.
func (GlooType) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_rawDescGZIP(), []int{0}
}

type ServerVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type GlooType `protobuf:"varint,1,opt,name=type,proto3,enum=gloo.solo.io.GlooType" json:"type,omitempty"`
	// Whether or not this is an enterprise distribution
	Enterprise bool `protobuf:"varint,2,opt,name=enterprise,proto3" json:"enterprise,omitempty"`
	// The type of server distribution
	// Currently only kubernetes is supported
	//
	// Types that are assignable to VersionType:
	//	*ServerVersion_Kubernetes
	VersionType isServerVersion_VersionType `protobuf_oneof:"version_type"`
}

func (x *ServerVersion) Reset() {
	*x = ServerVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerVersion) ProtoMessage() {}

func (x *ServerVersion) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerVersion.ProtoReflect.Descriptor instead.
func (*ServerVersion) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_rawDescGZIP(), []int{0}
}

func (x *ServerVersion) GetType() GlooType {
	if x != nil {
		return x.Type
	}
	return GlooType_Unknown
}

func (x *ServerVersion) GetEnterprise() bool {
	if x != nil {
		return x.Enterprise
	}
	return false
}

func (m *ServerVersion) GetVersionType() isServerVersion_VersionType {
	if m != nil {
		return m.VersionType
	}
	return nil
}

func (x *ServerVersion) GetKubernetes() *Kubernetes {
	if x, ok := x.GetVersionType().(*ServerVersion_Kubernetes); ok {
		return x.Kubernetes
	}
	return nil
}

type isServerVersion_VersionType interface {
	isServerVersion_VersionType()
}

type ServerVersion_Kubernetes struct {
	Kubernetes *Kubernetes `protobuf:"bytes,3,opt,name=kubernetes,proto3,oneof"`
}

func (*ServerVersion_Kubernetes) isServerVersion_VersionType() {}

type Kubernetes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Array of containers comprising a single distribution of gloo
	Containers []*Kubernetes_Container `protobuf:"bytes,1,rep,name=containers,proto3" json:"containers,omitempty"`
	// namespace gloo is running in
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *Kubernetes) Reset() {
	*x = Kubernetes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kubernetes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kubernetes) ProtoMessage() {}

func (x *Kubernetes) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kubernetes.ProtoReflect.Descriptor instead.
func (*Kubernetes) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_rawDescGZIP(), []int{1}
}

func (x *Kubernetes) GetContainers() []*Kubernetes_Container {
	if x != nil {
		return x.Containers
	}
	return nil
}

func (x *Kubernetes) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type ClientVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ClientVersion) Reset() {
	*x = ClientVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientVersion) ProtoMessage() {}

func (x *ClientVersion) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientVersion.ProtoReflect.Descriptor instead.
func (*ClientVersion) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_rawDescGZIP(), []int{2}
}

func (x *ClientVersion) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client *ClientVersion `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
	// This field is an array of server versions because although there can only be 1 client version, there can
	// potentially be many instances of gloo running on a single cluster
	Server []*ServerVersion `protobuf:"bytes,2,rep,name=server,proto3" json:"server,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_rawDescGZIP(), []int{3}
}

func (x *Version) GetClient() *ClientVersion {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *Version) GetServer() []*ServerVersion {
	if x != nil {
		return x.Server
	}
	return nil
}

type Kubernetes_Container struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag      string `protobuf:"bytes,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Registry string `protobuf:"bytes,3,opt,name=Registry,proto3" json:"Registry,omitempty"`
}

func (x *Kubernetes_Container) Reset() {
	*x = Kubernetes_Container{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kubernetes_Container) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kubernetes_Container) ProtoMessage() {}

func (x *Kubernetes_Container) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kubernetes_Container.ProtoReflect.Descriptor instead.
func (*Kubernetes_Container) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Kubernetes_Container) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Kubernetes_Container) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Kubernetes_Container) GetRegistry() string {
	if x != nil {
		return x.Registry
	}
	return ""
}

var File_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_rawDesc = []byte{
	0x0a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65,
	0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xac, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x47, 0x6c, 0x6f, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x12,
	0x3a, 0x0a, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x48, 0x00, 0x52,
	0x0a, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x42, 0x13, 0x0a, 0x0c, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x03, 0xf8, 0x42, 0x01,
	0x22, 0xbd, 0x01, 0x0a, 0x0a, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x12,
	0x42, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x1a, 0x4d, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x54, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x54, 0x61, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x22, 0x29, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x73, 0x0a, 0x07, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x06, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2a, 0x3e, 0x0a, 0x08, 0x47, 0x6c, 0x6f, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4b, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x10, 0x03,
	0x42, 0x44, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0xc0, 0xf5,
	0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_goTypes = []interface{}{
	(GlooType)(0),                // 0: gloo.solo.io.GlooType
	(*ServerVersion)(nil),        // 1: gloo.solo.io.ServerVersion
	(*Kubernetes)(nil),           // 2: gloo.solo.io.Kubernetes
	(*ClientVersion)(nil),        // 3: gloo.solo.io.ClientVersion
	(*Version)(nil),              // 4: gloo.solo.io.Version
	(*Kubernetes_Container)(nil), // 5: gloo.solo.io.Kubernetes.Container
}
var file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_depIdxs = []int32{
	0, // 0: gloo.solo.io.ServerVersion.type:type_name -> gloo.solo.io.GlooType
	2, // 1: gloo.solo.io.ServerVersion.kubernetes:type_name -> gloo.solo.io.Kubernetes
	5, // 2: gloo.solo.io.Kubernetes.containers:type_name -> gloo.solo.io.Kubernetes.Container
	3, // 3: gloo.solo.io.Version.client:type_name -> gloo.solo.io.ClientVersion
	1, // 4: gloo.solo.io.Version.server:type_name -> gloo.solo.io.ServerVersion
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerVersion); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kubernetes); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientVersion); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kubernetes_Container); i {
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
	file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ServerVersion_Kubernetes)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_grpc_version_version_proto_depIdxs = nil
}
