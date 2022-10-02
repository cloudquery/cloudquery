package recipes

import (
	serviceusage "cloud.google.com/go/serviceusage/apiv1"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	pb "google.golang.org/genproto/googleapis/api/serviceusage/v1"
)

var serviceusageResources = []*Resource{
	{
		SubService:          "services",
		Struct:              &pb.Service{},
		NewFunction:         serviceusage.NewClient,
		RequestStruct:       &pb.ListServicesRequest{},
		ResponseStruct:      &pb.ListServicesResponse{},
		RegisterServer:      pb.RegisterServiceUsageServer,
		ListFunction:        (&pb.UnimplementedServiceUsageServer{}).ListServices,
		UnimplementedServer: &pb.UnimplementedServiceUsageServer{},
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:    "name",
				Type:    schema.TypeString,
				Options: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	},
}

func ServiceusageResources() []*Resource {
	var resources []*Resource
	resources = append(resources, serviceusageResources...)

	for _, resource := range resources {
		resource.Service = "serviceusage"
		resource.MockImports = []string{"cloud.google.com/go/serviceusage/apiv1"}
		resource.ProtobufImport = "google.golang.org/genproto/googleapis/api/serviceusage/v1"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.RequestStructFields = `Parent: "projects/" + c.ProjectId,`
	}

	return resources
}
