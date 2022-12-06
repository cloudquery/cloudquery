package recipes

import (
	apikeys "cloud.google.com/go/apikeys/apiv2"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	pb "google.golang.org/genproto/googleapis/api/apikeys/v2"
)

var resources = []*Resource{
	{
		SubService: "keys",
		Struct:     &pb.Key{},
		SkipFields: []string{"Uid"},
		ExtraColumns: []codegen.ColumnDefinition{
			ProjectIdColumnPk,
			{
				Name:     "uid",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("Uid")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},

		ListFunction:        (&apikeys.Client{}).ListKeys,
		RequestStruct:       &pb.ListKeysRequest{},
		ResponseStruct:      &pb.ListKeysResponse{},
		RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/global",`,
	},
}

func ApiKeysResources() []*Resource {
	for _, resource := range resources {
		resource.Service = "apikeys"
		resource.Template = "newapi_list"

		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.ProtobufImport = "google.golang.org/genproto/googleapis/api/apikeys/v2"
		resource.MockImports = []string{"cloud.google.com/go/apikeys/apiv2"}
		resource.NewFunction = apikeys.NewClient
		resource.RegisterServer = pb.RegisterApiKeysServer
		resource.UnimplementedServer = &pb.UnimplementedApiKeysServer{}
	}

	return resources
}
