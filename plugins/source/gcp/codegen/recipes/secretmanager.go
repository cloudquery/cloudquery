package recipes

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	pb "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

func init() {
	resources := []*Resource{
		{
			SubService:     "secrets",
			Struct:         &pb.Secret{},
			NewFunction:    secretmanager.NewClient,
			RegisterServer: pb.RegisterSecretManagerServiceServer,
			PrimaryKeys:    []string{"name"},
			SkipFields:     []string{"Expiration"},
			Description:    "https://cloud.google.com/secret-manager/docs/reference/rest/v1/projects.secrets#Secret",
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
