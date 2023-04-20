package aiplatform

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	"google.golang.org/api/option"

	"google.golang.org/genproto/googleapis/cloud/location"

	aiplatform "cloud.google.com/go/aiplatform/apiv1"
)

func HyperparameterTuningJobs() *schema.Table {
	return &schema.Table{
		Name:        "gcp_aiplatform_hyperparameter_tuning_jobs",
		Description: `https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.hyperparameterTuningJobs#HyperparameterTuningJob`,
		Resolver:    fetchHyperparameterTuningJobs,
		Multiplex:   client.ProjectMultiplexEnabledServices("aiplatform.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.HyperparameterTuningJob{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
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
