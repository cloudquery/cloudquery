package recipes

import (
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	cloudresourcemanager "google.golang.org/api/cloudresourcemanager/v3"
)

func init() {
	resources := []*Resource{
		{
			SubService: "folders",
			Struct:     &pb.Folder{},
			SkipFetch:  true,
			SkipMock:   true,
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

	for _, resource := range resources {
		resource.Service = "resourcemanager"
		resource.MockImports = []string{"cloud.google.com/go/resourcemanager/apiv3"}
		resource.ProtobufImport = "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
	}

	Resources = append(Resources, resources...)
}
