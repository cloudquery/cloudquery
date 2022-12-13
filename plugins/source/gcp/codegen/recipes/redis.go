package recipes

import (
	redis "cloud.google.com/go/redis/apiv1"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	pb "google.golang.org/genproto/googleapis/cloud/redis/v1"
)

func init() {
	resources := []*Resource{
		{
			SubService:          "instances",
			Struct:              &pb.Instance{},
			NewFunction:         redis.NewCloudRedisClient,
			RequestStruct:       &pb.ListInstancesRequest{},
			ResponseStruct:      &pb.ListInstancesResponse{},
			RegisterServer:      pb.RegisterCloudRedisServer,
			UnimplementedServer: &pb.UnimplementedCloudRedisServer{},
			ListFunction:        (&pb.UnimplementedCloudRedisServer{}).ListInstances,
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "name",
					Type:     schema.TypeString,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					Resolver: `schema.PathResolver("Name")`,
				},
			},
		},
	}

	for _, resource := range resources {
		resource.Service = "redis"
		resource.MockImports = []string{"cloud.google.com/go/redis/apiv1"}
		resource.ProtobufImport = "google.golang.org/genproto/googleapis/cloud/redis/v1"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.RequestStructFields = `Parent: "projects/" + c.ProjectId + "/locations/-",`
	}

	Resources = append(Resources, resources...)
}
