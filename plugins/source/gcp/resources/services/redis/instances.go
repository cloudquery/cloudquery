package redis

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/redis/apiv1/redispb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	redis "cloud.google.com/go/redis/apiv1"
)

func Instances() *schema.Table {
	return &schema.Table{
		Name:        "gcp_redis_instances",
		Description: `https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance`,
		Resolver:    fetchInstances,
		Multiplex:   client.ProjectMultiplexEnabledServices("redis.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.Instance{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
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
	}
}

func fetchInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListInstancesRequest{
		Parent: "projects/" + c.ProjectId + "/locations/-",
	}
	gcpClient, err := redis.NewCloudRedisClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListInstances(ctx, req, c.CallOptions...)
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
