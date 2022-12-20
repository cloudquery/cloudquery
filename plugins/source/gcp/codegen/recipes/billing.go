package recipes

import (
	billing "cloud.google.com/go/billing/apiv1"
	pb "cloud.google.com/go/billing/apiv1/billingpb"
)

var ProjectMultiplex = "client.ProjectMultiplex"

func init() {
	resources := []*Resource{
		{
			SubService:          "billing_accounts",
			Struct:              &pb.BillingAccount{},
			NewFunction:         billing.NewCloudBillingClient,
			RequestStruct:       &pb.ListBillingAccountsRequest{},
			ResponseStruct:      &pb.ListBillingAccountsResponse{},
			ListFunction:        (&billing.CloudBillingClient{}).ListBillingAccounts,
			RegisterServer:      pb.RegisterCloudBillingServer,
			UnimplementedServer: &pb.UnimplementedCloudBillingServer{},
			PrimaryKeys:         []string{"name"},
			Description:         "https://cloud.google.com/billing/docs/reference/rest/v1/billingAccounts#BillingAccount",
		},
		{
			SubService:          "services",
			Struct:              &pb.Service{},
			NewFunction:         billing.NewCloudCatalogClient,
			RequestStruct:       &pb.ListServicesRequest{},
			ResponseStruct:      &pb.ListServicesResponse{},
			ListFunction:        (&billing.CloudCatalogClient{}).ListServices,
			RegisterServer:      pb.RegisterCloudCatalogServer,
			UnimplementedServer: &pb.UnimplementedCloudCatalogServer{},
			PrimaryKeys:         []string{"name"},
			Description:         "https://cloud.google.com/billing/docs/reference/rest/v1/services/list#Service",
			Multiplex:           &ProjectMultiplex,
		},
	}

	for _, resource := range resources {
		resource.Service = "billing"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.MockImports = []string{"cloud.google.com/go/billing/apiv1"}
		resource.ProtobufImport = "cloud.google.com/go/billing/apiv1/billingpb"
		resource.ServiceDNS = "cloudbilling.googleapis.com"
	}

	Resources = append(Resources, resources...)
}
