package recipes

import (
	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	pb "google.golang.org/genproto/googleapis/monitoring/v3"
)



func init() {
	resources :=  []*Resource{
		{
			SubService:          "alert_policies",
			Struct:              &pb.AlertPolicy{},
			NewFunction:         monitoring.NewAlertPolicyClient,
			RequestStruct:       &pb.ListAlertPoliciesRequest{},
			ResponseStruct:      &pb.ListAlertPoliciesResponse{},
			RegisterServer:      pb.RegisterAlertPolicyServiceServer,
			ListFunction:        (&pb.UnimplementedAlertPolicyServiceServer{}).ListAlertPolicies,
			UnimplementedServer: &pb.UnimplementedAlertPolicyServiceServer{},
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
		resource.Service = "monitoring"
		resource.MockImports = []string{"cloud.google.com/go/monitoring/apiv3/v2"}
		resource.ProtobufImport = "google.golang.org/genproto/googleapis/monitoring/v3"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.RequestStructFields = `Name: "projects/" + c.ProjectId,`
	}

	Resources = append(Resources, resources...)
}
