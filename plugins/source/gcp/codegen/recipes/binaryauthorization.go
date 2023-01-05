package recipes

import (
	binaryauthorization "cloud.google.com/go/binaryauthorization/apiv1"
	pb "cloud.google.com/go/binaryauthorization/apiv1/binaryauthorizationpb"
)

func init() {
	resources := []*Resource{
		{
			SubService:          "assertors",
			Struct:              &pb.Attestor{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: "projects/" + c.ProjectId,`,
			Description:         "https://cloud.google.com/binary-authorization/docs/reference/rest/v1/projects.attestors#Attestor",
		},
	}

	for _, resource := range resources {
		resource.Service = "binaryauthorization"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.ProtobufImport = "cloud.google.com/go/binaryauthorization/apiv1/binaryauthorizationpb"
		resource.MockImports = []string{"cloud.google.com/go/binaryauthorization/apiv1"}
		resource.NewFunction = binaryauthorization.NewBinauthzManagementClient
		resource.RegisterServer = pb.RegisterBinauthzManagementServiceV1Server
	}

	Resources = append(Resources, resources...)
}
