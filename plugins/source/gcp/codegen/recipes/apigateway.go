package recipes

import (
	apigateway "cloud.google.com/go/apigateway/apiv1"
	pb "cloud.google.com/go/apigateway/apiv1/apigatewaypb"
)

func init() {
	resources := []*Resource{
		{
			SubService:          "apis",
			Struct:              &pb.Api{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/-",`,
			Description:         "https://cloud.google.com/api-gateway/docs/reference/rest/v1/projects.locations.apis#Api",
		},
		{
			SubService:          "gateways",
			Struct:              &pb.Gateway{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/-",`,
			Description:         "https://cloud.google.com/api-gateway/docs/reference/rest/v1/projects.locations.gateways#Gateway",
		},
	}

	for _, resource := range resources {
		resource.Service = "apigateway"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.ProtobufImport = "cloud.google.com/go/apigateway/apiv1/apigatewaypb"
		resource.MockImports = []string{"cloud.google.com/go/apigateway/apiv1"}
		resource.NewFunction = apigateway.NewClient
		resource.RegisterServer = pb.RegisterApiGatewayServiceServer
	}

	Resources = append(Resources, resources...)
}
