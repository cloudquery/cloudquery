package recipes

import (
	errorreporting "cloud.google.com/go/errorreporting/apiv1beta1"
	pb "cloud.google.com/go/errorreporting/apiv1beta1/errorreportingpb"
)

func init() {
	resources := []*Resource{
		{
			SubService:          "error_group_stats",
			Struct:              &pb.ErrorGroupStats{},
			RequestStructFields: `ProjectName: "projects/" + c.ProjectId,`,
			Description:         "https://cloud.google.com/error-reporting/reference/rest/v1beta1/projects.groupStats/list#ErrorGroupStats",
			SkipMock:            true,
			Relations:           []string{"ErrorEvents()"},
		},
		{
			SubService:          "error_events",
			Struct:              &pb.ErrorEvent{},
			RequestStructFields: `ProjectName: "projects/" + c.ProjectId, GroupId: parent.Item.(*pb.ErrorGroupStats).Group.GroupId,`,
			Description:         "https://cloud.google.com/error-reporting/reference/rest/v1beta1/ErrorEvent",
			ChildTable:          true,
			SkipMock:            true,
		},
	}

	for _, resource := range resources {
		resource.Service = "clouderrorreporting"
		resource.ServiceAPIOverride = "errorreporting"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.ProtobufImport = "cloud.google.com/go/errorreporting/apiv1beta1/errorreportingpb"
		resource.MockImports = []string{"cloud.google.com/go/errorreporting/apiv1beta1"}
		resource.NewFunction = errorreporting.NewErrorStatsClient
		resource.RegisterServer = pb.RegisterErrorStatsServiceServer
	}

	Resources = append(Resources, resources...)
}
