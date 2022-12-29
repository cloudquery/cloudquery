package recipes

import (
	functions "cloud.google.com/go/functions/apiv1"
	"cloud.google.com/go/functions/apiv1/functionspb"
)

func init() {
	resources := []*Resource{
		{
			SubService:          "functions",
			Struct:              &functionspb.CloudFunction{},
			NewFunction:         functions.NewCloudFunctionsClient,
			RegisterServer:      functionspb.RegisterCloudFunctionsServiceServer,
			RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/-",`,
			SkipFields:          []string{"SourceCode", "Trigger"},
			Description:         "https://cloud.google.com/functions/docs/reference/rest/v1/projects.locations.functions#CloudFunction",
		},
	}

	for _, resource := range resources {
		resource.Service = "functions"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.MockImports = []string{"cloud.google.com/go/functions/apiv1"}
		resource.ProtobufImport = "cloud.google.com/go/functions/apiv1/functionspb"
		resource.ServiceDNS = "cloudfunctions.googleapis.com"
	}

	Resources = append(Resources, resources...)
}
