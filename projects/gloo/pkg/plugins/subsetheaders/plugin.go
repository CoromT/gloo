package subsetheaders

import (
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	envoy_header_to_metadata_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/header_to_metadata/v3"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/utils"
	"github.com/solo-io/solo-kit/pkg/errors"
	"google.golang.org/protobuf/types/known/anypb"
)

var (
	MissingHeaderValueError = errors.Errorf("header section of header value option cannot be nil")
	// this should go to the wellknown class in the control-plan module (see projects/gloo/pkg/plugins/cors/plugin.go)
	HeaderToMetadata = "envoy.filters.http.header_to_metadata"
)

// Puts Header Manipulation config on Routes, VirtualHosts, and Weighted Clusters
type Plugin struct{}

var _ plugins.RoutePlugin = NewPlugin()
var _ plugins.HttpFilterPlugin = NewPlugin()

// unclear exactly where in the filter chain this should go.
// putting it earlier to reuse header metadata later in the chain if needed.
var pluginStage = plugins.DuringStage(plugins.CorsStage)

func NewPlugin() *Plugin {
	return &Plugin{}
}

func (p *Plugin) Init(_ plugins.InitParams) error {
	return nil
}

func (p *Plugin) ProcessRoute(params plugins.RouteParams, in *v1.Route, out *envoy_config_route_v3.Route) error {
	var myconfig = &envoy_header_to_metadata_v3.Config{
		RequestRules: []*envoy_header_to_metadata_v3.Config_Rule{
			{
				Header: "azs-indexname",
				OnHeaderPresent: &envoy_header_to_metadata_v3.Config_KeyValuePair{
					MetadataNamespace: "envoy.lb",
					Key:               "azs-indexname",
					Type:              envoy_header_to_metadata_v3.Config_STRING,
				},
				Remove: false,
			},
			{
				Header: "azs-role",
				OnHeaderPresent: &envoy_header_to_metadata_v3.Config_KeyValuePair{
					MetadataNamespace: "envoy.lb",
					Key:               "azs-role",
					Type:              envoy_header_to_metadata_v3.Config_STRING,
				},
				Remove: false,
			},
			/*{
				Header: "azs-read",
				OnHeaderPresent: &envoy_header_to_metadata_v3.Config_KeyValuePair{
					MetadataNamespace: "envoy.lb",
					Key:               "azs-read",
					Type:              envoy_header_to_metadata_v3.Config_STRING,
				},
				Remove: false,
			},
			{
				Header: "azs-write",
				OnHeaderPresent: &envoy_header_to_metadata_v3.Config_KeyValuePair{
					MetadataNamespace: "envoy.lb",
					Key:               "azs-write",
					Type:              envoy_header_to_metadata_v3.Config_STRING,
				},
				Remove: false,
			},*/
		},
	}

	config, err := utils.MessageToAny(myconfig)
	if err != nil {
		return err
	}

	out.TypedPerFilterConfig = map[string]*anypb.Any{
		"envoy.filters.http.header_to_metadata": config,
	}

	return nil
}

func (p *Plugin) HttpFilters(params plugins.Params, l *v1.HttpListener) ([]plugins.StagedHttpFilter, error) {
	return []plugins.StagedHttpFilter{
		plugins.NewStagedFilter(HeaderToMetadata, pluginStage),
	}, nil
}
