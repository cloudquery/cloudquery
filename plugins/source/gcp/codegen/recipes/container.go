package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	pb "google.golang.org/genproto/googleapis/container/v1"
)



func init() {
	resources := []*Resource{
		{
			SubService: "clusters",
			Struct:     &pb.Cluster{},
			SkipFetch:  true,
			SkipMock:   true,
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:    "self_link",
					Type:    schema.TypeString,
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
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
