// Code generated by codegen; DO NOT EDIT.

package aiplatform

import (
	"context"
	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"google.golang.org/api/option"

	"google.golang.org/genproto/googleapis/cloud/location"

	"cloud.google.com/go/aiplatform/apiv1"
)

func PipelineJobs() *schema.Table {
	return &schema.Table{
		Name:        "gcp_aiplatform_pipeline_jobs",
		Description: `https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.pipelineJobs#PipelineJob`,
		Resolver:    fetchPipelineJobs,
		Multiplex:   client.ProjectMultiplexEnabledServices("aiplatform.googleapis.com"),
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
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "create_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("CreateTime"),
			},
			{
				Name:     "start_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("StartTime"),
			},
			{
				Name:     "end_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("EndTime"),
			},
			{
				Name:     "update_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("UpdateTime"),
			},
			{
				Name:          "pipeline_spec",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("PipelineSpec"),
				IgnoreInTests: true,
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("State"),
			},
			{
				Name:     "job_detail",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("JobDetail"),
			},
			{
				Name:     "error",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Error"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:          "runtime_config",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("RuntimeConfig"),
				IgnoreInTests: true,
			},
			{
				Name:     "encryption_spec",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EncryptionSpec"),
			},
			{
				Name:     "service_account",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceAccount"),
			},
			{
				Name:     "network",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Network"),
			},
			{
				Name:     "template_uri",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TemplateUri"),
			},
			{
				Name:     "template_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TemplateMetadata"),
			},
		},
	}
}

func fetchPipelineJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListPipelineJobsRequest{
		Parent: parent.Item.(*location.Location).Name,
	}
	if filterLocation(parent) {
		return nil
	}

	clientOptions := c.ClientOptions
	clientOptions = append([]option.ClientOption{option.WithEndpoint(parent.Item.(*location.Location).LocationId + "-aiplatform.googleapis.com:443")}, clientOptions...)
	gcpClient, err := aiplatform.NewPipelineClient(ctx, clientOptions...)

	if err != nil {
		return err
	}
	it := gcpClient.ListPipelineJobs(ctx, req, c.CallOptions...)
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
