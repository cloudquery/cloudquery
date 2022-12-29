package recipes

import (
	logging "cloud.google.com/go/logging/apiv2"
	pb "cloud.google.com/go/logging/apiv2/loggingpb"
)

func init() {
	resources := []*Resource{
		{
			SubService:     "metrics",
			Struct:         &pb.LogMetric{},
			NewFunction:    logging.NewMetricsClient,
			RegisterServer: pb.RegisterMetricsServiceV2Server,
			PrimaryKeys:    []string{"name"},
			Description:    "https://cloud.google.com/logging/docs/reference/v2/rest/v2/projects.metrics#LogMetric",
		},
		{
			SubService:     "sinks",
			Struct:         &pb.LogSink{},
			NewFunction:    logging.NewConfigClient,
			RegisterServer: pb.RegisterConfigServiceV2Server,
			PrimaryKeys:    []string{"name"},
			SkipFields:     []string{"Options"},
			Description:    "https://cloud.google.com/logging/docs/reference/v2/rest/v2/projects.sinks#LogSink",
		},
	}

	for _, resource := range resources {
		resource.Service = "logging"
		resource.RequestStructFields = `Parent: "projects/" + c.ProjectId,`
		resource.MockImports = []string{"cloud.google.com/go/logging/apiv2"}
		resource.ProtobufImport = "google.golang.org/genproto/googleapis/logging/v2"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
	}

	Resources = append(Resources, resources...)
}
