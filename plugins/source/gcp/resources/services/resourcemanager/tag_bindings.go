package resourcemanager

import (
	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"context"
	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"google.golang.org/api/iterator"
)

func TagBindings() *schema.Table {
	return &schema.Table{
		Name:        "gcp_resourcemanager_tag_bindings",
		Description: `https://cloud.google.com/resource-manager/reference/rest/v3/tagBindings`,
		Resolver:    fetchTagBindings,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudresourcemanager.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.TagValue{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},
	}
}

func fetchTagBindings(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	fClient, err := resourcemanager.NewTagBindingsClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	req := &pb.ListTagBindingsRequest{
		Parent: "//cloudresourcemanager.googleapis.com/projects/" + c.ProjectId,
	}

	it := fClient.ListTagBindings(ctx, req)
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
