// Code generated by codegen; DO NOT EDIT.

package batch

import (
	"context"
	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/batch/apiv1/batchpb"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"cloud.google.com/go/batch/apiv1"
)

func Jobs() *schema.Table {
	return &schema.Table{
		Name:        "gcp_batch_jobs",
		Description: `https://cloud.google.com/batch/docs/reference/rest/v1/projects.locations.jobs#Job`,
		Resolver:    fetchJobs,
		Multiplex:   client.ProjectMultiplexEnabledServices("batch.googleapis.com"),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "uid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Uid"),
			},
			{
				Name:     "priority",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Priority"),
			},
			{
				Name:     "task_groups",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TaskGroups"),
			},
			{
				Name:     "allocation_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AllocationPolicy"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "create_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("CreateTime"),
			},
			{
				Name:     "update_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("UpdateTime"),
			},
			{
				Name:     "logs_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LogsPolicy"),
			},
			{
				Name:     "notifications",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Notifications"),
			},
		},
	}
}

func fetchJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListJobsRequest{
		Parent: "projects/" + c.ProjectId + "/locations/-",
	}
	gcpClient, err := batch.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListJobs(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp

	}
	return nil
}
