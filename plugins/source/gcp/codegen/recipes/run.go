package recipes

import (
	run "cloud.google.com/go/run/apiv2"
	pb "google.golang.org/genproto/googleapis/cloud/run/v2"
)

var runResources = []*Resource{
	{
		SubService:          "services",
		Struct:              &pb.Service{},
		NewFunction:         run.NewServicesClient,
		RequestStruct:       &pb.ListServicesRequest{},
		ResponseStruct:      &pb.ListServicesResponse{},
		RegisterServer:      pb.RegisterServicesServer,
		ListFunction:        (&pb.UnimplementedServicesServer{}).ListServices,
		UnimplementedServer: &pb.UnimplementedServicesServer{},
	},
}

func RunResources() []*Resource {
	var resources []*Resource
	resources = append(resources, runResources...)

	for _, resource := range resources {
		resource.Service = "run"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.MockImports = []string{"cloud.google.com/go/run/apiv2"}
		resource.ProtobufImport = "google.golang.org/genproto/googleapis/cloud/run/v2"
		resource.RequestStructFields = `Parent: "projects/" + c.ProjectId + "locations/-",`
	}

	return resources
}
