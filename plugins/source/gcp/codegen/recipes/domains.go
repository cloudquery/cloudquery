package recipes

import (
	domains "cloud.google.com/go/domains/apiv1beta1"
	pb "cloud.google.com/go/domains/apiv1beta1/domainspb"
)

func init() {
	resources := []*Resource{
		{
			SubService:          "registrations",
			Struct:              &pb.Registration{},
			NewFunction:         domains.NewClient,
			RequestStruct:       &pb.ListRegistrationsRequest{},
			ResponseStruct:      &pb.ListRegistrationsResponse{},
			RegisterServer:      pb.RegisterDomainsServer,
			ListFunction:        (&pb.UnimplementedDomainsServer{}).ListRegistrations,
			UnimplementedServer: &pb.UnimplementedDomainsServer{},
			RequestStructFields: `Parent: fmt.Sprintf("projects/%s/locations/-", c.ProjectId),`,
			Imports:             []string{"fmt"},
			Description:         "https://cloud.google.com/domains/docs/reference/rest/v1beta1/projects.locations.registrations#Registration",
		},
	}

	for _, resource := range resources {
		resource.Service = "domains"
		resource.MockImports = []string{"cloud.google.com/go/domains/apiv1beta1"}
		resource.ProtobufImport = "cloud.google.com/go/domains/apiv1beta1/domainspb"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
	}

	Resources = append(Resources, resources...)
}
