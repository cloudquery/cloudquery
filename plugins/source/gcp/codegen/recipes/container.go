package recipes

import (
	pb "cloud.google.com/go/container/apiv1/containerpb"
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
			Description:     "https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#Cluster",
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
