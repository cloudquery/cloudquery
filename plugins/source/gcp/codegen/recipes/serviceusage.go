package recipes

import (
	serviceusage "cloud.google.com/go/serviceusage/apiv1"
	pb "cloud.google.com/go/serviceusage/apiv1/serviceusagepb"
)

func init() {
	resources := []*Resource{
		{
			SubService:          "services",
			Struct:              &pb.Service{},
			NewFunction:         serviceusage.NewClient,
			RequestStruct:       &pb.ListServicesRequest{},
			ResponseStruct:      &pb.ListServicesResponse{},
			RegisterServer:      pb.RegisterServiceUsageServer,
			ListFunction:        (&pb.UnimplementedServiceUsageServer{}).ListServices,
			UnimplementedServer: &pb.UnimplementedServiceUsageServer{},
			PrimaryKeys:         []string{"name"},
			Description:         "https://cloud.google.com/service-usage/docs/reference/rest/v1/services#Service",
		},
	}

	for _, resource := range resources {
		resource.Service = "serviceusage"
		resource.MockImports = []string{"cloud.google.com/go/serviceusage/apiv1"}
		resource.ProtobufImport = "cloud.google.com/go/serviceusage/apiv1/serviceusagepb"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.RequestStructFields = `Parent: "projects/" + c.ProjectId,
		PageSize: 200,
		Filter: "state:ENABLED",`
	}

	Resources = append(Resources, resources...)
}
