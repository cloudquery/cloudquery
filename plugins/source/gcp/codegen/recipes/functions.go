package recipes

import (
	functions "cloud.google.com/go/functions/apiv1"
	pb "google.golang.org/genproto/googleapis/cloud/functions/v1"
)

var functionsResources = []*Resource{
	{
		SubService:          "functions",
		Struct:              &pb.CloudFunction{},
		NewFunction:         functions.NewCloudFunctionsClient,
		RequestStruct:       &pb.ListFunctionsRequest{},
		ResponseStruct:      &pb.ListFunctionsResponse{},
		RegisterServer:      pb.RegisterCloudFunctionsServiceServer,
		ListFunction:        (&pb.UnimplementedCloudFunctionsServiceServer{}).ListFunctions,
		RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/-",`,
		UnimplementedServer: &pb.UnimplementedCloudFunctionsServiceServer{},
		FakerFieldsToIgnore: []string{"SourceCode", "Trigger"},
		// Skipping Timeout because `TypeInterval` is broken right now, and breaks the plugin completely.
		SkipFields: []string{"SourceCode", "Trigger", "Timeout"},
	},
}

func FunctionsResources() []*Resource {
	var resources []*Resource
	resources = append(resources, functionsResources...)

	for _, resource := range resources {
		resource.Service = "functions"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.MockImports = []string{"cloud.google.com/go/functions/apiv1"}
		resource.ProtobufImport = "google.golang.org/genproto/googleapis/cloud/functions/v1"
	}

	return resources
}
