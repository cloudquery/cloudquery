package recipes

import (
	billing "cloud.google.com/go/billing/apiv1"
	pb "cloud.google.com/go/billing/apiv1/billingpb"
)

func init() {
	resources := []*Resource{
		{
			SubService:     "billing_accounts",
			Struct:         &pb.BillingAccount{},
			NewFunction:    billing.NewCloudBillingClient,
			ListFunction:   (&billing.CloudBillingClient{}).ListBillingAccounts,
			RegisterServer: pb.RegisterCloudBillingServer,
			PrimaryKeys:    []string{"name"},
			Description:    "https://cloud.google.com/billing/docs/reference/rest/v1/billingAccounts#BillingAccount",
		},
		{
			SubService:     "services",
			Struct:         &pb.Service{},
			NewFunction:    billing.NewCloudCatalogClient,
			ListFunction:   (&billing.CloudCatalogClient{}).ListServices,
			RegisterServer: pb.RegisterCloudCatalogServer,
			PrimaryKeys:    []string{"name"},
			Description:    "https://cloud.google.com/billing/docs/reference/rest/v1/services/list#Service",
		},
	}

	for _, resource := range resources {
		resource.Service = "billing"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.MockImports = []string{"cloud.google.com/go/billing/apiv1"}
		resource.ProtobufImport = "cloud.google.com/go/billing/apiv1/billingpb"
	}

	Resources = append(Resources, resources...)
}
