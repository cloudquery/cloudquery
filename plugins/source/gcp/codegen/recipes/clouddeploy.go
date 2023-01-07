package recipes

import (
	deploy "cloud.google.com/go/deploy/apiv1"
	pb "cloud.google.com/go/deploy/apiv1/deploypb"
)

func init() {
	resources := []*Resource{
		{
			SubService:          "targets",
			Struct:              &pb.Target{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/-",`,
			Description:         "https://cloud.google.com/deploy/docs/api/reference/rest/v1/projects.locations.targets#Target",
		},
		{
			SubService:          "delivery_pipelines",
			Struct:              &pb.DeliveryPipeline{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/-",`,
			Description:         "https://cloud.google.com/deploy/docs/api/reference/rest/v1/projects.locations.deliveryPipelines#DeliveryPipeline",
			Relations:           []string{"Releases()"},
			SkipMock:            true,
		},
		{
			SubService:          "releases",
			Struct:              &pb.Release{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: parent.Item.(*pb.DeliveryPipeline).Name,`,
			Description:         "https://cloud.google.com/deploy/docs/api/reference/rest/v1/projects.locations.deliveryPipelines.releases#Release",
			ChildTable:          true,
			SkipMock:            true,
			Relations:           []string{"Rollouts()"},
		},
		{
			SubService:          "rollouts",
			Struct:              &pb.Rollout{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: parent.Item.(*pb.Release).Name,`,
			Description:         "https://cloud.google.com/deploy/docs/api/reference/rest/v1/projects.locations.deliveryPipelines.releases.rollouts#Rollout",
			ChildTable:          true,
			SkipMock:            true,
			Relations:           []string{"JobRuns()"},
		},
		{
			SubService:          "job_runs",
			Struct:              &pb.JobRun{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: parent.Item.(*pb.Rollout).Name,`,
			Description:         "https://cloud.google.com/deploy/docs/api/reference/rest/v1/projects.locations.deliveryPipelines.releases.rollouts.jobRuns#JobRun",
			ChildTable:          true,
			SkipMock:            true,
		},
	}

	for _, resource := range resources {
		resource.Service = "clouddeploy"
		resource.ServiceAPIOverride = "deploy"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.ProtobufImport = "cloud.google.com/go/deploy/apiv1/deploypb"
		resource.MockImports = []string{"cloud.google.com/go/deploy/apiv1"}
		resource.NewFunction = deploy.NewCloudDeployClient
		resource.RegisterServer = pb.RegisterCloudDeployServer
	}

	Resources = append(Resources, resources...)
}
