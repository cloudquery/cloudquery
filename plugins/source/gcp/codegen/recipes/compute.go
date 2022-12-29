package recipes

import (
	"reflect"
	"runtime"
	"strings"

	compute "cloud.google.com/go/compute/apiv1"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/iancoleman/strcase"
)

var computeResourcesAggList = []*Resource{
	{
		SubService:   "addresses",
		Struct:       &pb.Address{},
		NewFunction:  compute.NewAddressesRESTClient,
		ListFunction: (&compute.AddressesClient{}).AggregatedList,
		OutputField:  "Value.Addresses",
		Description:  "https://cloud.google.com/compute/docs/reference/rest/v1/addresses#Address",
	},
	{
		SubService:   "autoscalers",
		Struct:       &pb.Autoscaler{},
		NewFunction:  compute.NewAutoscalersRESTClient,
		ListFunction: (&compute.AutoscalersClient{}).AggregatedList,
		OutputField:  "Value.Autoscalers",
		Description:  "https://cloud.google.com/compute/docs/reference/rest/v1/autoscalers#Autoscaler",
	},
	{
		SubService:      "backend_services",
		Struct:          &pb.BackendService{},
		NewFunction:     compute.NewBackendServicesRESTClient,
		ListFunction:    (&compute.BackendServicesClient{}).AggregatedList,
		OutputField:     "Value.BackendServices",
		NameTransformer: CreateReplaceTransformer(map[string]string{"c_d_n": "cdn"}),
		Description:     "https://cloud.google.com/compute/docs/reference/rest/v1/backendServices#BackendService",
	},
	{
		SubService:   "disk_types",
		Struct:       &pb.DiskType{},
		NewFunction:  compute.NewDiskTypesRESTClient,
		ListFunction: (&compute.DiskTypesClient{}).AggregatedList,
		OutputField:  "Value.DiskTypes",
		Description:  "https://cloud.google.com/compute/docs/reference/rest/v1/diskTypes#DiskType",
	},
	{
		SubService:   "disks",
		Struct:       &pb.Disk{},
		NewFunction:  compute.NewDisksRESTClient,
		ListFunction: (&compute.DisksClient{}).AggregatedList,
		OutputField:  "Value.Disks",
		Description:  "https://cloud.google.com/compute/docs/reference/rest/v1/disks#Disk",
	},
	{
		SubService:      "forwarding_rules",
		Struct:          &pb.ForwardingRule{},
		NewFunction:     compute.NewForwardingRulesRESTClient,
		ListFunction:    (&compute.ForwardingRulesClient{}).AggregatedList,
		OutputField:     "Value.ForwardingRules",
		NameTransformer: CreateReplaceTransformer(map[string]string{"i_p_": "ip_"}),
		Description:     "https://cloud.google.com/compute/docs/reference/rest/v1/forwardingRules#ForwardingRule",
	},
	{
		SubService:      "instances",
		Struct:          &pb.Instance{},
		NewFunction:     compute.NewInstancesRESTClient,
		ListFunction:    (&compute.InstancesClient{}).AggregatedList,
		OutputField:     "Value.Instances",
		NameTransformer: CreateReplaceTransformer(map[string]string{"ipv_6": "ipv6"}),
		Description:     "https://cloud.google.com/compute/docs/reference/rest/v1/instances#Instance",
	},
	{
		SubService:   "ssl_certificates",
		Struct:       &pb.SslCertificate{},
		NewFunction:  compute.NewSslCertificatesRESTClient,
		ListFunction: (&compute.SslCertificatesClient{}).AggregatedList,
		OutputField:  "Value.SslCertificates",
		Description:  "https://cloud.google.com/compute/docs/reference/rest/v1/sslCertificates#SslCertificate",
	},
	{
		SubService:      "subnetworks",
		Struct:          &pb.Subnetwork{},
		NewFunction:     compute.NewSubnetworksRESTClient,
		ListFunction:    (&compute.SubnetworksClient{}).AggregatedList,
		OutputField:     "Value.Subnetworks",
		NameTransformer: CreateReplaceTransformer(map[string]string{"ipv_6": "ipv6", "i_pv_4": "ipv4"}),
		Description:     "https://cloud.google.com/compute/docs/reference/rest/v1/subnetworks#Subnetwork",
	},
	{
		SubService:   "target_http_proxies",
		Struct:       &pb.TargetHttpProxy{},
		NewFunction:  compute.NewTargetHttpProxiesRESTClient,
		ListFunction: (&compute.TargetHttpProxiesClient{}).AggregatedList,
		OutputField:  "Value.TargetHttpProxies",
		Description:  "https://cloud.google.com/compute/docs/reference/rest/v1/targetHttpProxies#TargetHttpProxy",
	},
	{
		SubService:     "url_maps",
		Struct:         &pb.UrlMap{},
		NewFunction:    compute.NewUrlMapsRESTClient,
		ResponseStruct: &pb.UrlMapsAggregatedList{},
		ListFunction:   (&compute.UrlMapsClient{}).AggregatedList,
		OutputField:    "Value.UrlMaps",
		Description:    "https://cloud.google.com/compute/docs/reference/rest/v1/urlMaps#UrlMap",
	},
	{
		SubService:   "vpn_gateways",
		Struct:       &pb.VpnGateway{},
		NewFunction:  compute.NewVpnGatewaysRESTClient,
		ListFunction: (&compute.VpnGatewaysClient{}).AggregatedList,
		OutputField:  "Value.VpnGateways",
		Description:  "https://cloud.google.com/compute/docs/reference/rest/v1/vpnGateways#VpnGateway",
	},
	{
		SubService:   "instance_groups",
		Struct:       &pb.InstanceGroup{},
		NewFunction:  compute.NewInstanceGroupsRESTClient,
		ListFunction: (&compute.InstanceGroupsClient{}).AggregatedList,
		OutputField:  "Value.InstanceGroups",
		Description:  "https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroups#InstanceGroup",
	},
}

var computeResourcesList = []*Resource{
	{
		SubService:   "images",
		Struct:       &pb.Image{},
		NewFunction:  compute.NewImagesRESTClient,
		ListFunction: (&compute.ImagesClient{}).List,
	},
	{
		SubService:   "firewalls",
		Struct:       &pb.Firewall{},
		NewFunction:  compute.NewFirewallsRESTClient,
		ListFunction: (&compute.FirewallsClient{}).List,
	},
	{
		SubService:      "networks",
		Struct:          &pb.Network{},
		NewFunction:     compute.NewNetworksRESTClient,
		ListFunction:    (&compute.NetworksClient{}).List,
		NameTransformer: CreateReplaceTransformer(map[string]string{"ipv_6": "ipv6", "i_pv4": "ipv4"}),
	},
	{
		SubService:     "ssl_policies",
		Struct:         &pb.SslPolicy{},
		NewFunction:    compute.NewSslPoliciesRESTClient,
		ResponseStruct: &pb.SslPoliciesList{},
		ListFunction:   (&compute.SslPoliciesClient{}).List,
	},
	{
		SubService:   "interconnects",
		Struct:       &pb.Interconnect{},
		NewFunction:  compute.NewInterconnectsRESTClient,
		ListFunction: (&compute.InterconnectsClient{}).List,
	},
	{
		SubService:   "target_ssl_proxies",
		Struct:       &pb.TargetSslProxy{},
		NewFunction:  compute.NewTargetSslProxiesRESTClient,
		ListFunction: (&compute.TargetSslProxiesClient{}).List,
	},
	{
		SubService:   "projects",
		Struct:       &pb.Project{},
		NewFunction:  compute.NewProjectsRESTClient,
		ListFunction: (&compute.ProjectsClient{}).Get,
		SkipFetch:    true,
		SkipMock:     true,
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
		resource.MockImports = []string{"cloud.google.com/go/compute/apiv1"}
		resource.ProtobufImport = "google.golang.org/genproto/googleapis/cloud/compute/v1"
		if resource.PrimaryKeys == nil {
			resource.PrimaryKeys = []string{"self_link"}
		}

	}

	Resources = append(Resources, resources...)
}
