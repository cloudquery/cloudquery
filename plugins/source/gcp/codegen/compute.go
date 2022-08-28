package codegen

import (
	"fmt"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
	"google.golang.org/api/compute/v1"
)

var computeResourcesAggList = []*Resource{
	{
		SubService: "addresses",
		Struct:     &compute.Address{},
	},
	{
		SubService: "autoscalers",
		Struct:     &compute.Autoscaler{},
	},
	{
		SubService: "backend_services",
		Struct:     &compute.BackendService{},
	},
	{
		SubService: "disk_types",
		Struct:     &compute.DiskType{},
		OverrideColumns: []codegen.ColumnDefinition{
			{
				Name:    "self_link",
				Type:    schema.TypeString,
				Options: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		// Id doesn't return anything
		SkipFields: []string{"ServerResponse", "NullFields", "ForceSendFields", "Id"},
	},
	{
		SubService: "forwarding_rules",
		Struct:     &compute.ForwardingRule{},
	},
	{
		SubService: "instances",
		Struct:     &compute.Instance{},
	},
	{
		SubService: "ssl_certificates",
		Struct:     &compute.SslCertificate{},
	},
	{
		SubService: "subnetworks",
		Struct:     &compute.Subnetwork{},
	},
	{
		SubService: "target_http_proxies",
		Struct:     &compute.TargetHttpProxy{},
	},
	{
		SubService: "url_maps",
		Struct:     &compute.UrlMap{},
	},
	{
		SubService: "vpn_gateways",
		Struct:     &compute.VpnGateway{},
	},
	{
		SubService: "instance_groups",
		Struct:     &compute.InstanceGroup{},
	},
}

var computeResourcesList = []*Resource{
	{
		SubService: "images",
		Struct:     &compute.Image{},
	},
	{
		SubService: "firewalls",
		Struct:     &compute.Firewall{},
	},
	{
		SubService: "networks",
		Struct:     &compute.Network{},
	},
	{
		SubService: "ssl_policies",
		Struct:     &compute.SslPolicy{},
	},
	{
		SubService: "interconnects",
		Struct:     &compute.Interconnect{},
	},
	{
		SubService: "target_ssl_proxies",
		Struct:     &compute.TargetSslProxy{},
	},
}

var computeResourcesGet = []*Resource{
	{
		SubService: "projects",
		Struct:     &compute.Project{},
	},
}

func ComputeResources() []*Resource {
	for _, resource := range computeResourcesList {
		if resource.Template == "" {
			resource.Template = "resource_list"
		}
		if resource.ListFunction == "" {
			resource.ListFunction = fmt.Sprintf("c.Services.Compute.%s.List(c.ProjectId).PageToken(nextPageToken).Do()", strcase.ToCamel(resource.SubService))
		}
	}
	for _, resource := range computeResourcesAggList {
		if resource.Template == "" {
			resource.Template = "resource_agg_list"
		}
		if len(resource.Imports) == 0 {
			resource.Imports = []string{"google.golang.org/api/compute/v1"}
		}
		if resource.ListFunction == "" {
			resource.ListFunction = fmt.Sprintf("c.Services.Compute.%s.AggregatedList(c.ProjectId).PageToken(nextPageToken).Do()", strcase.ToCamel(resource.SubService))
		}
	}
	for _, resource := range computeResourcesGet {
		if resource.Template == "" {
			resource.Template = "resource_get"
		}
		if resource.ListFunction == "" {
			resource.ListFunction = fmt.Sprintf("c.Services.Compute.%s.Get(c.ProjectId).Do()", strcase.ToCamel(resource.SubService))
		}
	}
	resources := computeResourcesAggList
	resources = append(resources, computeResourcesList...)
	resources = append(resources, computeResourcesGet...)
	// add all shared properties
	for i := range resources {
		resources[i].Service = "compute"
		if resources[i].OverrideColumns == nil {
			resources[i].OverrideColumns = []codegen.ColumnDefinition{
				{
					Name:    "self_link",
					Type:    schema.TypeString,
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}
		}
	}

	return resources
}
