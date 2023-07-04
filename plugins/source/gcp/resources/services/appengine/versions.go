package appengine

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	appengine "cloud.google.com/go/appengine/apiv1"
)

func Versions() *schema.Table {
	return &schema.Table{
		Name:        "gcp_appengine_versions",
		Description: `https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#Version`,
		Resolver:    fetchVersions,
		Multiplex:   client.ProjectMultiplexEnabledServices("appengine.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Version{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},
		Relations: []*schema.Table{
			Instances(),
		},
	}
}

func fetchVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListVersionsRequest{
		Parent: parent.Item.(*pb.Service).Name,
	}
	gcpClient, err := appengine.NewVersionsClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListVersions(ctx, req, c.CallOptions...)
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
