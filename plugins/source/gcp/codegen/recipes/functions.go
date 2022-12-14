package recipes

import (
	functions "cloud.google.com/go/functions/apiv1"
	"cloud.google.com/go/functions/apiv1/functionspb"
)



func init() {
	resources := []*Resource{
		{
			SubService:          "functions",
			Struct:              &functionspb.CloudFunction{},
			NewFunction:         functions.NewCloudFunctionsClient,
			RequestStruct:       &functionspb.ListFunctionsRequest{},
			ResponseStruct:      &functionspb.ListFunctionsResponse{},
			RegisterServer:      functionspb.RegisterCloudFunctionsServiceServer,
			ListFunction:        (&functionspb.UnimplementedCloudFunctionsServiceServer{}).ListFunctions,
			RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/-",`,
			UnimplementedServer: &functionspb.UnimplementedCloudFunctionsServiceServer{},
			SkipFields:          []string{"SourceCode", "Trigger"},
		},
	}

	for _, resource := range resources {
		resource.Service = "functions"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.MockImports = []string{"cloud.google.com/go/functions/apiv1"}
		resource.ProtobufImport = "cloud.google.com/go/functions/apiv1/functionspb"
	}

	Resources = append(Resources, resources...)
}
