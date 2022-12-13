package recipes

import (
	pb "google.golang.org/genproto/googleapis/container/v1"
)

func init() {
	resources := []*Resource{
		{
			SubService:      "clusters",
			Struct:          &pb.Cluster{},
			SkipFetch:       true,
			SkipMock:        true,
			PrimaryKeys:     []string{"self_link"},
			NameTransformer: CreateReplaceTransformer(map[string]string{"ipv_4": "ipv4"}),
		},
	}

	for _, resource := range resources {
		resource.Service = "container"
		resource.MockImports = []string{"cloud.google.com/go/container/apiv1"}
		resource.ProtobufImport = "google.golang.org/genproto/googleapis/container/v1"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
	}

	Resources = append(Resources, resources...)
}
