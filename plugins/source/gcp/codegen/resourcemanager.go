package codegen

import (
	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	pb "google.golang.org/genproto/googleapis/cloud/resourcemanager/v3"
)

var resourceManagerResources = []*Resource{
	{
		SubService:          "folders",
		Struct:              &pb.Folder{},
		NewFunction:         resourcemanager.NewFoldersClient,
		RequestStruct:       &pb.ListFoldersRequest{},
		ResponseStruct:      &pb.ListFoldersResponse{},
		RegisterServer:      pb.RegisterFoldersServer,
		ListFunction:        (&pb.UnimplementedFoldersServer{}).ListFolders,
		UnimplementedServer: &pb.UnimplementedFoldersServer{},
	},
	{
		SubService:          "projects",
		Struct:              &pb.Project{},
		NewFunction:         resourcemanager.NewProjectsClient,
		RequestStruct:       &pb.ListProjectsRequest{},
		ResponseStruct:      &pb.ListProjectsResponse{},
		RegisterServer:      pb.RegisterProjectsServer,
		ListFunction:        (&pb.UnimplementedProjectsServer{}).ListProjects,
		UnimplementedServer: &pb.UnimplementedProjectsServer{},
		SkipFetch:           true,
		SkipMock:            true,
		SkipFields:          []string{"ProjectId"},
	},
}

func ResourceManagerResources() []*Resource {
	var resources []*Resource
	resources = append(resources, resourceManagerResources...)

	for _, resource := range resources {
		resource.Service = "resourcemanager"
		resource.MockImports = []string{"cloud.google.com/go/resourcemanager/apiv3"}
		resource.ProtobufImport = "google.golang.org/genproto/googleapis/cloud/resourcemanager/v3"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
	}

	return resources
}
