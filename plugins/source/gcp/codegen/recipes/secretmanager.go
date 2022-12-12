package recipes

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	pb "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)



func init() {
	resources := []*Resource{
		{
			SubService:          "secrets",
			Struct:              &pb.Secret{},
			NewFunction:         secretmanager.NewClient,
			RequestStruct:       &pb.ListSecretsRequest{},
			ResponseStruct:      &pb.ListSecretsResponse{},
			RegisterServer:      pb.RegisterSecretManagerServiceServer,
			ListFunction:        (&pb.UnimplementedSecretManagerServiceServer{}).ListSecrets,
			UnimplementedServer: &pb.UnimplementedSecretManagerServiceServer{},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "name",
					Type:     schema.TypeString,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					Resolver: `schema.PathResolver("Name")`,
				},
			},
			SkipFields: []string{"Expiration"},
		},
	}

	for _, resource := range resources {
		resource.Service = "secretmanager"
		resource.MockImports = []string{"cloud.google.com/go/secretmanager/apiv1"}
		resource.ProtobufImport = "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.RequestStructFields = `Parent: "projects/" + c.ProjectId,`
	}

	Resources = append(Resources, resources...)
}
