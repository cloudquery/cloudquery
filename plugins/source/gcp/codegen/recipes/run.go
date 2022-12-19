package recipes

import (
	run "cloud.google.com/go/run/apiv2"
	pb "cloud.google.com/go/run/apiv2/runpb"
)

func init() {
	resources := []*Resource{
		{
			SubService:          "services",
			Struct:              &pb.Service{},
			NewFunction:         run.NewServicesClient,
			RequestStruct:       &pb.ListServicesRequest{},
			ResponseStruct:      &pb.ListServicesResponse{},
			RegisterServer:      pb.RegisterServicesServer,
			ListFunction:        (&pb.UnimplementedServicesServer{}).ListServices,
			UnimplementedServer: &pb.UnimplementedServicesServer{},
			Description:         "https://cloud.google.com/run/docs/reference/rest/v2/projects.locations.services#Service",
		},
	}

	for _, resource := range resources {
		resource.Service = "run"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.MockImports = []string{"cloud.google.com/go/run/apiv2"}
		resource.ProtobufImport = "cloud.google.com/go/run/apiv2/runpb"
		resource.RequestStructFields = `Parent: "projects/" + c.ProjectId + "locations/-",`
	}

	Resources = append(Resources, resources...)
}
