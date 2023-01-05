package recipes

import (
	batch "cloud.google.com/go/batch/apiv1"
	pb "cloud.google.com/go/batch/apiv1/batchpb"
)

func init() {
	resources := []*Resource{
		{
			SubService:          "jobs",
			Struct:              &pb.Job{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/-",`,
			Description:         "https://cloud.google.com/batch/docs/reference/rest/v1/projects.locations.jobs#Job",
		},
		{
			SubService:          "task_groups",
			Struct:              &pb.TaskGroup{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			ListFunction:        (&batch.Client{}).ListJobs,
			RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/-",`,
			Description:         "https://cloud.google.com/batch/docs/reference/rest/v1/projects.locations.jobs#TaskGroup",
			OutputField:         "TaskGroups",
			Relations:           []string{"Tasks()"},
			SkipMock:            true,
		},
		{
			SubService:          "tasks",
			Struct:              &pb.Task{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: parent.Item.(*pb.TaskGroup).Name,`,
			Description:         "https://cloud.google.com/batch/docs/reference/rest/v1/projects.locations.jobs.taskGroups.tasks/list",
			ChildTable:          true,
			SkipMock:            true,
		},
	}

	for _, resource := range resources {
		resource.Service = "batch"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.ProtobufImport = "cloud.google.com/go/batch/apiv1/batchpb"
		resource.MockImports = []string{"cloud.google.com/go/batch/apiv1"}
		resource.NewFunction = batch.NewClient
		resource.RegisterServer = pb.RegisterBatchServiceServer
	}

	Resources = append(Resources, resources...)
}
