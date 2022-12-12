package recipes

import (
	"reflect"
	"runtime"
	"strings"

	compute "cloud.google.com/go/compute/apiv1"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
	pb "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

var computeResourcesAggList = []*Resource{
	{
		SubService:     "addresses",
		Struct:         &pb.Address{},
		NewFunction:    compute.NewAddressesRESTClient,
		RequestStruct:  &pb.AggregatedListAddressesRequest{},
		ResponseStruct: &pb.AddressAggregatedList{},
		ListFunction:   (&compute.AddressesClient{}).AggregatedList,
		OutputField:    "Value.Addresses",
	},
	{
		SubService:     "autoscalers",
		Struct:         &pb.Autoscaler{},
		NewFunction:    compute.NewAutoscalersRESTClient,
		RequestStruct:  &pb.AggregatedListAutoscalersRequest{},
		ResponseStruct: &pb.AutoscalerAggregatedList{},
		ListFunction:   (&compute.AutoscalersClient{}).AggregatedList,
		OutputField:    "Value.Autoscalers",
	},
	{
		SubService:      "backend_services",
		Struct:          &pb.BackendService{},
		NewFunction:     compute.NewBackendServicesRESTClient,
		RequestStruct:   &pb.AggregatedListBackendServicesRequest{},
		ResponseStruct:  &pb.BackendServiceAggregatedList{},
		ListFunction:    (&compute.BackendServicesClient{}).AggregatedList,
		OutputField:     "Value.BackendServices",
		NameTransformer: CreateReplaceTransformer(map[string]string{"c_d_n": "cdn"}),
	},
	{
		SubService:     "disk_types",
		Struct:         &pb.DiskType{},
		NewFunction:    compute.NewDiskTypesRESTClient,
		RequestStruct:  &pb.AggregatedListDiskTypesRequest{},
		ResponseStruct: &pb.DiskTypeAggregatedList{},
		ListFunction:   (&compute.DiskTypesClient{}).AggregatedList,
		OutputField:    "Value.DiskTypes",
	},
	{
		SubService:     "disks",
		Struct:         &pb.Disk{},
		NewFunction:    compute.NewDisksRESTClient,
		RequestStruct:  &pb.AggregatedListDisksRequest{},
		ResponseStruct: &pb.DiskAggregatedList{},
		ListFunction:   (&compute.DisksClient{}).AggregatedList,
		OutputField:    "Value.Disks",
	},
	{
		SubService:      "forwarding_rules",
		Struct:          &pb.ForwardingRule{},
		NewFunction:     compute.NewForwardingRulesRESTClient,
		RequestStruct:   &pb.AggregatedListForwardingRulesRequest{},
		ResponseStruct:  &pb.ForwardingRuleAggregatedList{},
		ListFunction:    (&compute.ForwardingRulesClient{}).AggregatedList,
		OutputField:     "Value.ForwardingRules",
		NameTransformer: CreateReplaceTransformer(map[string]string{"i_p_": "ip_"}),
	},
	{
		SubService:      "instances",
		Struct:          &pb.Instance{},
		NewFunction:     compute.NewInstancesRESTClient,
		RequestStruct:   &pb.AggregatedListInstancesRequest{},
		ResponseStruct:  &pb.InstanceAggregatedList{},
		ListFunction:    (&compute.InstancesClient{}).AggregatedList,
		OutputField:     "Value.Instances",
		NameTransformer: CreateReplaceTransformer(map[string]string{"ipv_6": "ipv6"}),
	},
	{
		SubService:     "ssl_certificates",
		Struct:         &pb.SslCertificate{},
		NewFunction:    compute.NewSslCertificatesRESTClient,
		RequestStruct:  &pb.AggregatedListSslCertificatesRequest{},
		ResponseStruct: &pb.SslCertificateAggregatedList{},
		ListFunction:   (&compute.SslCertificatesClient{}).AggregatedList,
		OutputField:    "Value.SslCertificates",
	},
	{
		SubService:      "subnetworks",
		Struct:          &pb.Subnetwork{},
		NewFunction:     compute.NewSubnetworksRESTClient,
		RequestStruct:   &pb.AggregatedListSubnetworksRequest{},
		ResponseStruct:  &pb.SubnetworkAggregatedList{},
		ListFunction:    (&compute.SubnetworksClient{}).AggregatedList,
		OutputField:     "Value.Subnetworks",
		NameTransformer: CreateReplaceTransformer(map[string]string{"ipv_6": "ipv6", "i_pv_4": "ipv4"}),
	},
	{
		SubService:     "target_http_proxies",
		Struct:         &pb.TargetHttpProxy{},
		NewFunction:    compute.NewTargetHttpProxiesRESTClient,
		RequestStruct:  &pb.AggregatedListTargetHttpProxiesRequest{},
		ResponseStruct: &pb.TargetHttpProxyAggregatedList{},
		ListFunction:   (&compute.TargetHttpProxiesClient{}).AggregatedList,
		OutputField:    "Value.TargetHttpProxies",
	},
	{
		SubService:     "url_maps",
		Struct:         &pb.UrlMap{},
		NewFunction:    compute.NewUrlMapsRESTClient,
		RequestStruct:  &pb.AggregatedListUrlMapsRequest{},
		ResponseStruct: &pb.UrlMapsAggregatedList{},
		ListFunction:   (&compute.UrlMapsClient{}).AggregatedList,
		OutputField:    "Value.UrlMaps",
	},
	{
		SubService:     "vpn_gateways",
		Struct:         &pb.VpnGateway{},
		NewFunction:    compute.NewVpnGatewaysRESTClient,
		RequestStruct:  &pb.AggregatedListVpnGatewaysRequest{},
		ResponseStruct: &pb.VpnGatewayAggregatedList{},
		ListFunction:   (&compute.VpnGatewaysClient{}).AggregatedList,
		OutputField:    "Value.VpnGateways",
	},
	{
		SubService:     "instance_groups",
		Struct:         &pb.InstanceGroup{},
		NewFunction:    compute.NewInstanceGroupsRESTClient,
		RequestStruct:  &pb.AggregatedListInstanceGroupsRequest{},
		ResponseStruct: &pb.InstanceGroupAggregatedList{},
		ListFunction:   (&compute.InstanceGroupsClient{}).AggregatedList,
		OutputField:    "Value.InstanceGroups",
	},
}

var computeResourcesList = []*Resource{
	{
		SubService:     "images",
		Struct:         &pb.Image{},
		NewFunction:    compute.NewImagesRESTClient,
		RequestStruct:  &pb.ListImagesRequest{},
		ResponseStruct: &pb.ImageList{},
		ListFunction:   (&compute.ImagesClient{}).List,
	},
	{
		SubService:     "firewalls",
		Struct:         &pb.Firewall{},
		NewFunction:    compute.NewFirewallsRESTClient,
		RequestStruct:  &pb.ListFirewallsRequest{},
		ResponseStruct: &pb.FirewallList{},
		ListFunction:   (&compute.FirewallsClient{}).List,
	},
	{
		SubService:      "networks",
		Struct:          &pb.Network{},
		NewFunction:     compute.NewNetworksRESTClient,
		RequestStruct:   &pb.ListNetworksRequest{},
		ResponseStruct:  &pb.NetworkList{},
		ListFunction:    (&compute.NetworksClient{}).List,
		NameTransformer: CreateReplaceTransformer(map[string]string{"ipv_6": "ipv6", "i_pv4": "ipv4"}),
	},
	{
		SubService:     "ssl_policies",
		Struct:         &pb.SslPolicy{},
		NewFunction:    compute.NewSslPoliciesRESTClient,
		RequestStruct:  &pb.ListSslPoliciesRequest{},
		ResponseStruct: &pb.SslPoliciesList{},
		ListFunction:   (&compute.InterconnectsClient{}).List,
	},
	{
		SubService:     "interconnects",
		Struct:         &pb.Interconnect{},
		NewFunction:    compute.NewInterconnectsRESTClient,
		RequestStruct:  &pb.ListInterconnectsRequest{},
		ResponseStruct: &pb.InterconnectList{},
		ListFunction:   (&compute.InterconnectsClient{}).List,
	},
	{
		SubService:     "target_ssl_proxies",
		Struct:         &pb.TargetSslProxy{},
		NewFunction:    compute.NewTargetSslProxiesRESTClient,
		RequestStruct:  &pb.ListTargetSslProxiesRequest{},
		ResponseStruct: &pb.TargetSslProxyList{},
		ListFunction:   (&compute.TargetSslProxiesClient{}).List,
	},
	{
		SubService:     "projects",
		Struct:         &pb.Project{},
		NewFunction:    compute.NewProjectsRESTClient,
		RequestStruct:  &pb.GetProjectRequest{},
		ResponseStruct: &pb.Project{},
		ListFunction:   (&compute.ProjectsClient{}).Get,
		SkipFetch:      true,
		SkipMock:       true,
	},
}

func init() {
	for _, resource := range computeResourcesList {
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_rest_mock"
		resource.RequestStructName = "List" + strcase.ToCamel(resource.SubService) + "Request"
	}
	for _, resource := range computeResourcesAggList {
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_rest_mock"
	}
	resources := computeResourcesAggList
	resources = append(resources, computeResourcesList...)

	// add all shared properties
	for _, resource := range resources {
		resource.RequestStructFields = `Project: c.ProjectId,`
		resource.Service = "compute"
		if resource.NewFunction != nil {
			newFunctionNamePath := strings.Split(runtime.FuncForPC(reflect.ValueOf(resource.NewFunction).Pointer()).Name(), ".")
			resource.NewFunctionName = newFunctionNamePath[len(newFunctionNamePath)-1]
		}
		if resource.ResponseStruct != nil {
			resource.ResponseStructName = reflect.TypeOf(resource.ResponseStruct).Elem().Name()
		}
		if resource.RequestStruct != nil {
			resource.RequestStructName = reflect.TypeOf(resource.RequestStruct).Elem().Name()
		}
		resource.MockImports = []string{"cloud.google.com/go/compute/apiv1"}
		resource.ProtobufImport = "google.golang.org/genproto/googleapis/cloud/compute/v1"
		if resource.ExtraColumns == nil {
			resource.ExtraColumns = []codegen.ColumnDefinition{
				{
					Name:    "self_link",
					Type:    schema.TypeString,
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}
		}
	}

	Resources = append(Resources, resources...)
}
