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

func OrganizationTagKeys() *schema.Table {
	return &schema.Table{
		Name:        "gcp_resourcemanager_organization_tag_keys",
		Description: `https://cloud.google.com/resource-manager/reference/rest/v3/tagKeys/list`,
		Resolver:    fetchOrganizationTagKeys,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudresourcemanager.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.TagKey{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "organization_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveOrganization,
				PrimaryKey: true,
			},
		},
		Relations: []*schema.Table{
			OrganizationTagValues(),
		},
	}
}

func fetchOrganizationTagKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	fClient, err := resourcemanager.NewTagKeysClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	req := &pb.ListTagKeysRequest{
		Parent: "organizations/" + c.OrgId,
	}
	it := fClient.ListTagKeys(ctx, req)
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
