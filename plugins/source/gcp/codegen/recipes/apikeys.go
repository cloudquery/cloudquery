package recipes

import (
	apikeys "cloud.google.com/go/apikeys/apiv2"
	pb "google.golang.org/genproto/googleapis/api/apikeys/v2"
)

func init() {
	resources := []*Resource{
		{
			SubService:          "keys",
			Struct:              &pb.Key{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "uid"},
			ListFunction:        (&apikeys.Client{}).ListKeys,
			RequestStruct:       &pb.ListKeysRequest{},
			ResponseStruct:      &pb.ListKeysResponse{},
			RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/global",`,
		},
	}

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

	Resources = append(Resources, resources...)
}
