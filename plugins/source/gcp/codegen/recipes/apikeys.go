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
			RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/global",`,
			Description:         "https://cloud.google.com/api-keys/docs/reference/rest/v2/projects.locations.keys#Key",
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
	}

	Resources = append(Resources, resources...)
}
