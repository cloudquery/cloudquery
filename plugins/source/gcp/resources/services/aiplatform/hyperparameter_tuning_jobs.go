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

func HyperparameterTuningJobs() *schema.Table {
	return &schema.Table{
		Name:        "gcp_aiplatform_hyperparameter_tuning_jobs",
		Description: `https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.hyperparameterTuningJobs#HyperparameterTuningJob`,
		Resolver:    fetchHyperparameterTuningJobs,
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
				Name:     "study_spec",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StudySpec"),
			},
			{
				Name:     "max_trial_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxTrialCount"),
			},
			{
				Name:     "parallel_trial_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ParallelTrialCount"),
			},
			{
				Name:     "max_failed_trial_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxFailedTrialCount"),
			},
			{
				Name:     "trial_job_spec",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TrialJobSpec"),
			},
			{
				Name:     "trials",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Trials"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("State"),
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
				Name:     "encryption_spec",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EncryptionSpec"),
			},
		},
	}
}

func fetchHyperparameterTuningJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListHyperparameterTuningJobsRequest{
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
	it := gcpClient.ListHyperparameterTuningJobs(ctx, req, c.CallOptions...)
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
