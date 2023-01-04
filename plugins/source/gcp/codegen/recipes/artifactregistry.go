package recipes

import (
	artifactregistry "cloud.google.com/go/artifactregistry/apiv1"
	pb "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
	registerV1 "google.golang.org/api/artifactregistry/v1"
)

func init() {
	resources := []*Resource{
		{
			SubService:  "locations",
			PrimaryKeys: []string{ProjectIdColumn.Name, "name"},
			Struct:      &registerV1.Location{},
			SkipFetch:   true,
			SkipMock:    true,
			Description: "https://cloud.google.com/artifact-registry/docs/reference/rest/Shared.Types/ListLocationsResponse#Location",
			Relations:   []string{"Repositories()"},
		},
		{
			SubService:  "repositories",
			Struct:      &pb.Repository{},
			PrimaryKeys: []string{ProjectIdColumn.Name, "name"},
			Description: "https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories#Repository",
			ChildTable:  true,
			SkipFetch:   true,
			SkipMock:    true,
			Relations:   []string{"DockerImages()", "Files()", "Packages()"},
		},
		{
			SubService:          "docker_images",
			Struct:              &pb.DockerImage{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: parent.Item.(*pb.Repository).Name,`,
			Description:         "https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories.dockerImages#DockerImage",
			ChildTable:          true,
			SkipMock:            true,
		},
		{
			SubService:          "files",
			Struct:              &pb.File{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: parent.Item.(*pb.Repository).Name,`,
			Description:         "https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories.files#File",
			ChildTable:          true,
			SkipMock:            true,
		},
		{
			SubService:          "packages",
			Struct:              &pb.Package{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: parent.Item.(*pb.Repository).Name,`,
			Description:         "https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories.packages#Package",
			ChildTable:          true,
			SkipMock:            true,
			Relations:           []string{"Tags()", "Versions()"},
		},
		{
			SubService:          "tags",
			Struct:              &pb.Tag{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: parent.Item.(*pb.Package).Name,`,
			Description:         "https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories.packages.tags#Tag",
			ChildTable:          true,
			SkipMock:            true,
		},
		{
			SubService:          "versions",
			Struct:              &pb.Version{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: parent.Item.(*pb.Package).Name,`,
			Description:         "https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories.packages.versions#Version",
			ChildTable:          true,
			SkipMock:            true,
		},
	}

	for _, resource := range resources {
		resource.Service = "artifactregistry"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		if resource.ProtobufImport == "" {
			resource.ProtobufImport = "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
		}
		resource.MockImports = []string{"cloud.google.com/go/artifactregistry/apiv1"}
		resource.NewFunction = artifactregistry.NewClient
		resource.RegisterServer = pb.RegisterArtifactRegistryServer
	}

	Resources = append(Resources, resources...)
}
