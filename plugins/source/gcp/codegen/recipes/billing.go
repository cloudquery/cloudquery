package recipes

import (
	billing "cloud.google.com/go/billing/apiv1"
	pb "cloud.google.com/go/billing/apiv1/billingpb"
	budgets "cloud.google.com/go/billing/budgets/apiv1"
	budgetspb "cloud.google.com/go/billing/budgets/apiv1/budgetspb"
)

var ProjectMultiplex = "client.ProjectMultiplex"

func init() {
	resources := []*Resource{
		{
			SubService:     "billing_accounts",
			Struct:         &pb.BillingAccount{},
			NewFunction:    billing.NewCloudBillingClient,
			RegisterServer: pb.RegisterCloudBillingServer,
			PrimaryKeys:    []string{"name"},
			Description:    "https://cloud.google.com/billing/docs/reference/rest/v1/billingAccounts#BillingAccount",
			Relations:      []string{"Budgets()"},
			ServiceDNS:     "cloudbilling.googleapis.com",
			MockImports:    []string{"cloud.google.com/go/billing/apiv1"},
			ProtobufImport: "cloud.google.com/go/billing/apiv1/billingpb",
			SkipMock:       true,
		},
		{
			SubService:     "services",
			Struct:         &pb.Service{},
			NewFunction:    billing.NewCloudCatalogClient,
			RegisterServer: pb.RegisterCloudCatalogServer,
			PrimaryKeys:    []string{"name"},
			Description:    "https://cloud.google.com/billing/docs/reference/rest/v1/services/list#Service",
			Multiplex:      &ProjectMultiplex,
			ServiceDNS:     "cloudbilling.googleapis.com",
			MockImports:    []string{"cloud.google.com/go/billing/apiv1"},
			ProtobufImport: "cloud.google.com/go/billing/apiv1/billingpb",
		},
		{
			SubService:          "budgets",
			Struct:              &budgetspb.Budget{},
			RequestStructFields: "Parent: parent.Item.(*pb.Repository).Name,",
			NewFunction:         budgets.NewBudgetClient,
			RegisterServer:      budgetspb.RegisterBudgetServiceServer,
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			Description:         "https://cloud.google.com/billing/docs/reference/budget/rest/v1/billingAccounts.budgets#Budget",
			Multiplex:           &ProjectMultiplex,
			ServiceDNS:          "billingbudgets.googleapis.com",
			ChildTable:          true,
			SkipMock:            true,
			SkipFetch:           true,
			MockImports:         []string{"cloud.google.com/go/billing/budgets/apiv1"},
			ProtobufImport:      "cloud.google.com/go/billing/budgets/apiv1/budgetspb",
			ServiceAPIOverride:  "budgets",
		},
	}

	for _, resource := range resources {
		resource.Service = "billing"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
	}

	Resources = append(Resources, resources...)
}
