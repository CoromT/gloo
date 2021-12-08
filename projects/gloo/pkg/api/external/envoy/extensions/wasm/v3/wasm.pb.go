// Copied from https://github.com/envoyproxy/envoy/blob/26eaa2e85cee69e5c32ab6bf4c5ae3d338fa462f/api/envoy/extensions/wasm/v3/wasm.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/wasm/v3/wasm.proto

package v3

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	any "github.com/golang/protobuf/ptypes/any"
	v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/core/v3"
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

// Configuration for a Wasm VM.
// [#next-free-field: 7]
type VmConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An ID which will be used along with a hash of the wasm code (or the name of the registered Null
	// VM plugin) to determine which VM will be used for the plugin. All plugins which use the same
	// *vm_id* and code will use the same VM. May be left blank. Sharing a VM between plugins can
	// reduce memory utilization and make sharing of data easier which may have security implications.
	// See ref: "TODO: add ref" for details.
	VmId string `protobuf:"bytes,1,opt,name=vm_id,json=vmId,proto3" json:"vm_id,omitempty"`
	// The Wasm runtime type (either "v8" or "null" for code compiled into Envoy).
	Runtime string `protobuf:"bytes,2,opt,name=runtime,proto3" json:"runtime,omitempty"`
	// The Wasm code that Envoy will execute.
	Code *v3.AsyncDataSource `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	// The Wasm configuration used in initialization of a new VM
	// (proxy_on_start). `google.protobuf.Struct` is serialized as JSON before
	// passing it to the plugin. `google.protobuf.BytesValue` and
	// `google.protobuf.StringValue` are passed directly without the wrapper.
	Configuration *any.Any `protobuf:"bytes,4,opt,name=configuration,proto3" json:"configuration,omitempty"`
	// Allow the wasm file to include pre-compiled code on VMs which support it.
	// Warning: this should only be enable for trusted sources as the precompiled code is not
	// verified.
	AllowPrecompiled bool `protobuf:"varint,5,opt,name=allow_precompiled,json=allowPrecompiled,proto3" json:"allow_precompiled,omitempty"`
	// If true and the code needs to be remotely fetched and it is not in the cache then NACK the configuration
	// update and do a background fetch to fill the cache, otherwise fetch the code asynchronously and enter
	// warming state.
	NackOnCodeCacheMiss bool `protobuf:"varint,6,opt,name=nack_on_code_cache_miss,json=nackOnCodeCacheMiss,proto3" json:"nack_on_code_cache_miss,omitempty"`
}

func (x *VmConfig) Reset() {
	*x = VmConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VmConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VmConfig) ProtoMessage() {}

func (x *VmConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VmConfig.ProtoReflect.Descriptor instead.
func (*VmConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_rawDescGZIP(), []int{0}
}

func (x *VmConfig) GetVmId() string {
	if x != nil {
		return x.VmId
	}
	return ""
}

func (x *VmConfig) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *VmConfig) GetCode() *v3.AsyncDataSource {
	if x != nil {
		return x.Code
	}
	return nil
}

func (x *VmConfig) GetConfiguration() *any.Any {
	if x != nil {
		return x.Configuration
	}
	return nil
}

func (x *VmConfig) GetAllowPrecompiled() bool {
	if x != nil {
		return x.AllowPrecompiled
	}
	return false
}

func (x *VmConfig) GetNackOnCodeCacheMiss() bool {
	if x != nil {
		return x.NackOnCodeCacheMiss
	}
	return false
}

// Base Configuration for Wasm Plugins e.g. filters and services.
// [#next-free-field: 6]
type PluginConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique name for a filters/services in a VM for use in identifying the filter/service if
	// multiple filters/services are handled by the same *vm_id* and *root_id* and for
	// logging/debugging.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A unique ID for a set of filters/services in a VM which will share a RootContext and Contexts
	// if applicable (e.g. an Wasm HttpFilter and an Wasm AccessLog). If left blank, all
	// filters/services with a blank root_id with the same *vm_id* will share Context(s).
	RootId string `protobuf:"bytes,2,opt,name=root_id,json=rootId,proto3" json:"root_id,omitempty"`
	// Configuration for finding or starting VM.
	//
	// Types that are assignable to Vm:
	//	*PluginConfig_VmConfig
	Vm isPluginConfig_Vm `protobuf_oneof:"vm"`
	// Filter/service configuration used to configure or reconfigure a plugin
	// (proxy_on_configuration).
	// `google.protobuf.Struct` is serialized as JSON before
	// passing it to the plugin. `google.protobuf.BytesValue` and
	// `google.protobuf.StringValue` are passed directly without the wrapper.
	Configuration *any.Any `protobuf:"bytes,4,opt,name=configuration,proto3" json:"configuration,omitempty"`
	// If there is a fatal error on the VM (e.g. exception, abort(), on_start or on_configure return false),
	// then all plugins associated with the VM will either fail closed (by default), e.g. by returning an HTTP 503 error,
	// or fail open (if 'fail_open' is set to true) by bypassing the filter. Note: when on_start or on_configure return false
	// during xDS updates the xDS configuration will be rejected and when on_start or on_configuration return false on initial
	// startup the proxy will not start.
	FailOpen bool `protobuf:"varint,5,opt,name=fail_open,json=failOpen,proto3" json:"fail_open,omitempty"`
}

func (x *PluginConfig) Reset() {
	*x = PluginConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginConfig) ProtoMessage() {}

func (x *PluginConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginConfig.ProtoReflect.Descriptor instead.
func (*PluginConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_rawDescGZIP(), []int{1}
}

func (x *PluginConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PluginConfig) GetRootId() string {
	if x != nil {
		return x.RootId
	}
	return ""
}

func (m *PluginConfig) GetVm() isPluginConfig_Vm {
	if m != nil {
		return m.Vm
	}
	return nil
}

func (x *PluginConfig) GetVmConfig() *VmConfig {
	if x, ok := x.GetVm().(*PluginConfig_VmConfig); ok {
		return x.VmConfig
	}
	return nil
}

func (x *PluginConfig) GetConfiguration() *any.Any {
	if x != nil {
		return x.Configuration
	}
	return nil
}

func (x *PluginConfig) GetFailOpen() bool {
	if x != nil {
		return x.FailOpen
	}
	return false
}

type isPluginConfig_Vm interface {
	isPluginConfig_Vm()
}

type PluginConfig_VmConfig struct {
	VmConfig *VmConfig `protobuf:"bytes,3,opt,name=vm_config,json=vmConfig,proto3,oneof"` // TODO: add referential VM configurations.
}

func (*PluginConfig_VmConfig) isPluginConfig_Vm() {}

// WasmService is configured as a built-in *envoy.wasm_service* :ref:`WasmService
// <config_wasm_service>` This opaque configuration will be used to create a Wasm Service.
type WasmService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// General plugin configuration.
	Config *PluginConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	// If true, create a single VM rather than creating one VM per worker. Such a singleton can
	// not be used with filters.
	Singleton bool `protobuf:"varint,2,opt,name=singleton,proto3" json:"singleton,omitempty"`
}

func (x *WasmService) Reset() {
	*x = WasmService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WasmService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WasmService) ProtoMessage() {}

func (x *WasmService) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WasmService.ProtoReflect.Descriptor instead.
func (*WasmService) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_rawDescGZIP(), []int{2}
}

func (x *WasmService) GetConfig() *PluginConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *WasmService) GetSingleton() bool {
	if x != nil {
		return x.Singleton
	}
	return false
}

var File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_rawDesc = []byte{
	0x0a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x77, 0x61, 0x73, 0x6d, 0x2f, 0x76, 0x33, 0x2f, 0x77, 0x61,
	0x73, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x33, 0x1a, 0x52, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x02, 0x0a, 0x08, 0x56, 0x6d, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x13, 0x0a, 0x05, 0x76, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x76, 0x6d, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x20, 0x01, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x3a,
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0d, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x5f, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x72, 0x65, 0x63,
	0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x17, 0x6e, 0x61, 0x63, 0x6b, 0x5f,
	0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x6d, 0x69,
	0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x6e, 0x61, 0x63, 0x6b, 0x4f, 0x6e,
	0x43, 0x6f, 0x64, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4d, 0x69, 0x73, 0x73, 0x22, 0xe5, 0x01,
	0x0a, 0x0c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x49, 0x0a, 0x09, 0x76,
	0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x76,
	0x33, 0x2e, 0x56, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x08, 0x76, 0x6d,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3a, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61, 0x69, 0x6c, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x66, 0x61, 0x69, 0x6c, 0x4f, 0x70, 0x65, 0x6e, 0x42,
	0x04, 0x0a, 0x02, 0x76, 0x6d, 0x22, 0x73, 0x0a, 0x0b, 0x57, 0x61, 0x73, 0x6d, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x77, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x42, 0x92, 0x01, 0x0a, 0x26, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x61,
	0x73, 0x6d, 0x2e, 0x76, 0x33, 0x42, 0x09, 0x57, 0x61, 0x73, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x77, 0x61, 0x73, 0x6d,
	0x2f, 0x76, 0x33, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_goTypes = []interface{}{
	(*VmConfig)(nil),           // 0: solo.io.envoy.extensions.wasm.v3.VmConfig
	(*PluginConfig)(nil),       // 1: solo.io.envoy.extensions.wasm.v3.PluginConfig
	(*WasmService)(nil),        // 2: solo.io.envoy.extensions.wasm.v3.WasmService
	(*v3.AsyncDataSource)(nil), // 3: solo.io.envoy.config.core.v3.AsyncDataSource
	(*any.Any)(nil),            // 4: google.protobuf.Any
}
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_depIdxs = []int32{
	3, // 0: solo.io.envoy.extensions.wasm.v3.VmConfig.code:type_name -> solo.io.envoy.config.core.v3.AsyncDataSource
	4, // 1: solo.io.envoy.extensions.wasm.v3.VmConfig.configuration:type_name -> google.protobuf.Any
	0, // 2: solo.io.envoy.extensions.wasm.v3.PluginConfig.vm_config:type_name -> solo.io.envoy.extensions.wasm.v3.VmConfig
	4, // 3: solo.io.envoy.extensions.wasm.v3.PluginConfig.configuration:type_name -> google.protobuf.Any
	1, // 4: solo.io.envoy.extensions.wasm.v3.WasmService.config:type_name -> solo.io.envoy.extensions.wasm.v3.PluginConfig
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VmConfig); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginConfig); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WasmService); i {
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
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*PluginConfig_VmConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_wasm_v3_wasm_proto_depIdxs = nil
}
