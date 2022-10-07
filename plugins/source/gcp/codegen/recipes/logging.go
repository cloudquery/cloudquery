package recipes

import (
	logging "cloud.google.com/go/logging/apiv2"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	pb "google.golang.org/genproto/googleapis/logging/v2"
)

var loggingResources = []*Resource{
	{
		SubService:          "metrics",
		Struct:              &pb.LogMetric{},
		NewFunction:         logging.NewMetricsClient,
		RequestStruct:       &pb.ListLogMetricsRequest{},
		ResponseStruct:      &pb.ListLogMetricsResponse{},
		RegisterServer:      pb.RegisterMetricsServiceV2Server,
		ListFunction:        (&pb.UnimplementedMetricsServiceV2Server{}).ListLogMetrics,
		UnimplementedServer: &pb.UnimplementedMetricsServiceV2Server{},
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "name",
				Type:     schema.TypeString,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				Resolver: `schema.PathResolver("Name")`,
			},
		},
	},
	{
		SubService:          "sinks",
		Struct:              &pb.LogSink{},
		NewFunction:         logging.NewConfigClient,
		RequestStruct:       &pb.ListSinksRequest{},
		ResponseStruct:      &pb.ListSinksResponse{},
		RegisterServer:      pb.RegisterConfigServiceV2Server,
		ListFunction:        (&pb.UnimplementedConfigServiceV2Server{}).ListSinks,
		UnimplementedServer: &pb.UnimplementedConfigServiceV2Server{},
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "name",
				Type:     schema.TypeString,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				Resolver: `schema.PathResolver("Name")`,
			},
		},
		SkipFields: []string{"Options"},
	},
}

func LoggingResources() []*Resource {
	var resources []*Resource
	resources = append(resources, loggingResources...)

	for _, resource := range resources {
		resource.Service = "logging"
		resource.RequestStructFields = `Parent: "projects/" + c.ProjectId,`
		resource.MockImports = []string{"cloud.google.com/go/logging/apiv2"}
		resource.ProtobufImport = "google.golang.org/genproto/googleapis/logging/v2"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
	}

	return resources
}
