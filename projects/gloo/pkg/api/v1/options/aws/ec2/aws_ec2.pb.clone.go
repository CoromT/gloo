// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/aws/ec2/aws_ec2.proto

package ec2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = clone.Cloner(nil)
	_ = proto.Message(nil)
)

// Clone function
func (m *UpstreamSpec) Clone() proto.Message {
	var target *UpstreamSpec
	if m == nil {
		return target
	}
	target = &UpstreamSpec{}

	target.Region = m.GetRegion()

	if h, ok := interface{}(m.GetSecretRef()).(clone.Cloner); ok {
		target.SecretRef = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	} else {
		target.SecretRef = proto.Clone(m.GetSecretRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	}

	target.RoleArn = m.GetRoleArn()

	if m.GetFilters() != nil {
		target.Filters = make([]*TagFilter, len(m.GetFilters()))
		for idx, v := range m.GetFilters() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Filters[idx] = h.Clone().(*TagFilter)
			} else {
				target.Filters[idx] = proto.Clone(v).(*TagFilter)
			}

		}
	}

	target.PublicIp = m.GetPublicIp()

	target.Port = m.GetPort()

	return target
}

// Clone function
func (m *TagFilter) Clone() proto.Message {
	var target *TagFilter
	if m == nil {
		return target
	}
	target = &TagFilter{}

	switch m.Spec.(type) {

	case *TagFilter_Key:

		target.Spec = &TagFilter_Key{
			Key: m.GetKey(),
		}

	case *TagFilter_KvPair_:

		if h, ok := interface{}(m.GetKvPair()).(clone.Cloner); ok {
			target.Spec = &TagFilter_KvPair_{
				KvPair: h.Clone().(*TagFilter_KvPair),
			}
		} else {
			target.Spec = &TagFilter_KvPair_{
				KvPair: proto.Clone(m.GetKvPair()).(*TagFilter_KvPair),
			}
		}

	}

	return target
}

// Clone function
func (m *TagFilter_KvPair) Clone() proto.Message {
	var target *TagFilter_KvPair
	if m == nil {
		return target
	}
	target = &TagFilter_KvPair{}

	target.Key = m.GetKey()

	target.Value = m.GetValue()

	return target
}
