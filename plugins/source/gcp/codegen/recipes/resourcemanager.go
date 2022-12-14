package recipes

import (
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	cloudresourcemanager "google.golang.org/api/cloudresourcemanager/v3"
)

var OrgMultiplex = "client.OrgMultiplex"

func init() {
	resources := []*Resource{
		{
			SubService: "folders",
			Struct:     &pb.Folder{},
			SkipFetch:  true,
			SkipMock:   true,
			Multiplex:  &OrgMultiplex,
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "organization_id",
					Type:     schema.TypeString,
					Resolver: "resolveOrganizationId",
				},
			},
			Description: "https://cloud.google.com/resource-manager/reference/rest/v3/folders#Folder",
		},
		{
			SubService:  "projects",
			Struct:      &pb.Project{},
			SkipFetch:   true,
			SkipMock:    true,
			SkipFields:  []string{"ProjectId"},
			PrimaryKeys: []string{"project_id", "name"},
			Description: "https://cloud.google.com/resource-manager/reference/rest/v3/projects#Project",
		},
		{
			SubService:  "project_policies",
			Struct:      &cloudresourcemanager.Policy{},
			SkipFetch:   true,
			SkipMock:    true,
			Description: "https://cloud.google.com/resource-manager/reference/rest/Shared.Types/Policy",
		},
	}

	for _, resource := range resources {
		resource.Service = "resourcemanager"
		resource.MockImports = []string{"cloud.google.com/go/resourcemanager/apiv3"}
		resource.ProtobufImport = "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.ServiceDNS = "cloudresourcemanager.googleapis.com"
	}

	Resources = append(Resources, resources...)
}
