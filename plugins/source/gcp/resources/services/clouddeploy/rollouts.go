package clouddeploy

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/deploy/apiv1/deploypb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	deploy "cloud.google.com/go/deploy/apiv1"
)

func Rollouts() *schema.Table {
	return &schema.Table{
		Name:        "gcp_clouddeploy_rollouts",
		Description: `https://cloud.google.com/deploy/docs/api/reference/rest/v1/projects.locations.deliveryPipelines.releases.rollouts#Rollout`,
		Resolver:    fetchRollouts,
		Multiplex:   client.ProjectMultiplexEnabledServices("clouddeploy.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.Rollout{}, client.Options()...),
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
		},
		Relations: []*schema.Table{
			JobRuns(),
		},
	}
}

func fetchRollouts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListRolloutsRequest{
		Parent: parent.Item.(*pb.Release).Name,
	}
	gcpClient, err := deploy.NewCloudDeployClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListRollouts(ctx, req, c.CallOptions...)
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
