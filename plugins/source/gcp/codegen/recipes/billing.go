package recipes

import (
	billing "cloud.google.com/go/billing/apiv1"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	pb "google.golang.org/genproto/googleapis/cloud/billing/v1"
)



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
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:    "name",
					Type:    schema.TypeString,
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
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
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:    "name",
					Type:    schema.TypeString,
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
	}

	for _, resource := range resources {
		resource.Service = "billing"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.MockImports = []string{"cloud.google.com/go/billing/apiv1"}
		resource.ProtobufImport = "google.golang.org/genproto/googleapis/cloud/billing/v1"
	}

	Resources = append(Resources, resources...)
}
