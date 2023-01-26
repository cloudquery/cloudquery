package appengine

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	appengine "cloud.google.com/go/appengine/apiv1"
)

func Instances() *schema.Table {
	return &schema.Table{
		Name:        "gcp_appengine_instances",
		Description: `https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions.instances#Instance`,
		Resolver:    fetchInstances,
		Multiplex:   client.ProjectMultiplexEnabledServices("appengine.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.Instance{}, append(client.Options(), transformers.WithPrimaryKeys("Name"))...),
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

func fetchInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListInstancesRequest{
		Parent: parent.Item.(*pb.Version).Name,
	}
	gcpClient, err := appengine.NewInstancesClient(ctx, c.ClientOptions...)
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
