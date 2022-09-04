package codegen

import (
	domains "cloud.google.com/go/domains/apiv1beta1"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	pb "google.golang.org/genproto/googleapis/cloud/domains/v1beta1"
)

var domainsResources = []*Resource{
	{
		SubService:          "registrations",
		Struct:              &pb.Registration{},
		NewFunction:         domains.NewClient,
		RequestStruct:       &pb.ListRegistrationsRequest{},
		ResponseStruct:      &pb.ListRegistrationsResponse{},
		RegisterServer:      pb.RegisterDomainsServer,
		ListFunction:        (&pb.UnimplementedDomainsServer{}).ListRegistrations,
		UnimplementedServer: &pb.UnimplementedDomainsServer{},
		OverrideColumns: []codegen.ColumnDefinition{
			{
				Name:    "self_link",
				Type:    schema.TypeString,
				Options: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	},
}

func DomainsResources() []*Resource {
	var resources []*Resource
	resources = append(resources, domainsResources...)

	for _, resource := range resources {
		resource.Service = "domains"
		resource.MockImports = []string{"cloud.google.com/go/domains/apiv1beta1"}
		resource.ProtobufImport = "google.golang.org/genproto/googleapis/cloud/domains/v1beta1"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		// resource.OutputField = strcase.ToCamel(resource.SubService)
	}

	return resources
}
