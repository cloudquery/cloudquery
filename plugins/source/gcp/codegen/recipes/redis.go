package recipes

import (
	redis "cloud.google.com/go/redis/apiv1"
	pb "cloud.google.com/go/redis/apiv1/redispb"
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
			PrimaryKeys:         []string{"name"},
		},
	}

	for _, resource := range resources {
		resource.Service = "redis"
		resource.MockImports = []string{"cloud.google.com/go/redis/apiv1"}
		resource.ProtobufImport = "cloud.google.com/go/redis/apiv1/redispb"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.RequestStructFields = `Parent: "projects/" + c.ProjectId + "/locations/-",`
	}

	Resources = append(Resources, resources...)
}
