package recipes

import (
	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
)

func init() {
	resources := []*Resource{
		{
			SubService:     "alert_policies",
			Struct:         &pb.AlertPolicy{},
			NewFunction:    monitoring.NewAlertPolicyClient,
			RegisterServer: pb.RegisterAlertPolicyServiceServer,
			PrimaryKeys:    []string{"name"},
			Description:    "https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.alertPolicies#AlertPolicy",
		},
	}

	for _, resource := range resources {
		resource.Service = "monitoring"
		resource.MockImports = []string{"cloud.google.com/go/monitoring/apiv3/v2"}
		resource.ProtobufImport = "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.RequestStructFields = `Name: "projects/" + c.ProjectId,`
	}

	Resources = append(Resources, resources...)
}
