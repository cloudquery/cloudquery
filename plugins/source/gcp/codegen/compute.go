package codegen

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/iancoleman/strcase"
	"google.golang.org/api/compute/v1"
)

var computeResourcesAggList = []Resource{
	{
		GCPSubService: "addresses",
		GCPStruct:     &compute.Address{},
	},
	{
		GCPSubService: "autoscalers",
		GCPStruct:     &compute.Autoscaler{},
	},
	{
		GCPSubService: "backend_services",
		GCPStruct:     &compute.BackendService{},
	},
	{
		GCPSubService: "disk_types",
		GCPStruct:     &compute.DiskType{},
	},
	{
		GCPSubService: "forwarding_rules",
		GCPStruct:     &compute.ForwardingRule{},
	},
	{
		GCPSubService: "instances",
		GCPStruct:     &compute.Instance{},
	},
	{
		GCPSubService: "ssl_certificates",
		GCPStruct:     &compute.SslCertificate{},
	},
	{
		GCPSubService: "subnetworks",
		GCPStruct:     &compute.Subnetwork{},
	},
	{
		GCPSubService: "target_http_proxies",
		GCPStruct:     &compute.TargetHttpProxy{},
	},
	{
		GCPSubService:  "url_maps",
		GCPStruct:      &compute.UrlMap{},
		MockListStruct: "UrlMaps",
	},
	{
		GCPSubService: "vpn_gateways",
		GCPStruct:     &compute.VpnGateway{},
	},
	{
		GCPSubService: "instance_groups",
		GCPStruct:     &compute.InstanceGroup{},
	},
}

var computeResourcesList = []Resource{
	{
		GCPSubService: "images",
		GCPStruct:     &compute.Image{},
	},
	{
		GCPSubService: "firewalls",
		GCPStruct:     &compute.Firewall{},
	},
	{
		GCPSubService: "networks",
		GCPStruct:     &compute.Network{},
	},
	{
		GCPSubService:  "ssl_policies",
		GCPStruct:      &compute.SslPolicy{},
		MockListStruct: "SslPolicies",
	},
	{
		GCPSubService: "interconnects",
		GCPStruct:     &compute.Interconnect{},
	},
	{
		GCPSubService: "target_ssl_proxies",
		GCPStruct:     &compute.TargetSslProxy{},
	},
}

var computeResourcesGet = []Resource{
	{
		GCPSubService: "projects",
		GCPStruct:     &compute.Project{},
	},
}

func ComputeResources() []Resource {
	for i := range computeResourcesList {
		if computeResourcesList[i].Template == "" {
			computeResourcesList[i].Template = "resource_list"
		}
	}
	for i := range computeResourcesAggList {
		if computeResourcesAggList[i].Template == "" {
			computeResourcesAggList[i].Template = "resource_agg_list"
		}
		if len(computeResourcesAggList[i].Imports) == 0 {
			computeResourcesAggList[i].Imports = []string{"google.golang.org/api/compute/v1"}
		}
	}
	for i := range computeResourcesGet {
		if computeResourcesGet[i].Template == "" {
			computeResourcesGet[i].Template = "resource_get"
		}
	}
	resources := computeResourcesAggList
	resources = append(resources, computeResourcesList...)
	resources = append(resources, computeResourcesGet...)
	// add all shared properties
	for i := range resources {
		resources[i].GCPService = "compute"
		resources[i].DefaultColumns = []codegen.ColumnDefinition{ProjectIdColumn}
		resources[i].GCPStructName = reflect.TypeOf(resources[i].GCPStruct).Elem().Name()
		resources[i].SkipFields = []string{"SelfLink", "NullFields", "ForceSendFields"}
		resources[i].MockImports = []string{"google.golang.org/api/compute/v1"}
		if resources[i].MockListStruct == "" {
			resources[i].MockListStruct = strcase.ToCamel(resources[i].GCPStructName)
		}
	}

	return resources
}
