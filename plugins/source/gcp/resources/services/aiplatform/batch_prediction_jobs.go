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

func BatchPredictionJobs() *schema.Table {
	return &schema.Table{
		Name:        "gcp_aiplatform_batch_prediction_jobs",
		Description: `https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.batchPredictionJobs#BatchPredictionJob`,
		Resolver:    fetchBatchPredictionJobs,
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
				Name:     "model",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Model"),
			},
			{
				Name:     "model_version_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ModelVersionId"),
			},
			{
				Name:     "unmanaged_container_model",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("UnmanagedContainerModel"),
			},
			{
				Name:     "input_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("InputConfig"),
			},
			{
				Name:          "model_parameters",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("ModelParameters"),
				IgnoreInTests: true,
			},
			{
				Name:     "output_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OutputConfig"),
			},
			{
				Name:     "dedicated_resources",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DedicatedResources"),
			},
			{
				Name:     "service_account",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceAccount"),
			},
			{
				Name:     "manual_batch_tuning_parameters",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ManualBatchTuningParameters"),
			},
			{
				Name:     "generate_explanation",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("GenerateExplanation"),
			},
			{
				Name:     "explanation_spec",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ExplanationSpec"),
			},
			{
				Name:     "output_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OutputInfo"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("State"),
			},
			{
				Name:     "error",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Error"),
			},
			{
				Name:     "partial_failures",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PartialFailures"),
			},
			{
				Name:     "resources_consumed",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResourcesConsumed"),
			},
			{
				Name:     "completion_stats",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CompletionStats"),
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
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "encryption_spec",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EncryptionSpec"),
			},
		},
	}
}

func fetchBatchPredictionJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListBatchPredictionJobsRequest{
		Parent: parent.Item.(*location.Location).Name,
	}
	if filterLocation(parent) {
		return nil
	}

	clientOptions := c.ClientOptions
	clientOptions = append([]option.ClientOption{option.WithEndpoint(parent.Item.(*location.Location).LocationId + "-aiplatform.googleapis.com:443")}, clientOptions...)
	gcpClient, err := aiplatform.NewJobClient(ctx, clientOptions...)

	if err != nil {
		return err
	}
	it := gcpClient.ListBatchPredictionJobs(ctx, req, c.CallOptions...)
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
