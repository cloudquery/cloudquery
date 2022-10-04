package recipes

import (
	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	cloudresourcemanager "google.golang.org/api/cloudresourcemanager/v3"
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
		SubService: "projects",
		Struct:     &pb.Project{},
		SkipFetch:  true,
		SkipMock:   true,
		SkipFields: []string{"ProjectId"},
	},
	{
		SubService: "project_policies",
		Struct:     &cloudresourcemanager.Policy{},
		SkipFetch:  true,
		SkipMock:   true,
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
