package iam

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/iam/admin/apiv1/adminpb"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	admin "cloud.google.com/go/iam/admin/apiv1"
)

func ServiceAccounts() *schema.Table {
	return &schema.Table{
		Name:        "gcp_iam_service_accounts",
		Description: `https://cloud.google.com/iam/docs/reference/rest/v1/projects.serviceAccounts#ServiceAccount`,
		Resolver:    fetchServiceAccounts,
		Multiplex:   client.ProjectMultiplexEnabledServices("iam.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.ServiceAccount{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveProject,
			},
			{
				Name:       "unique_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("UniqueId"),
				PrimaryKey: true,
			},
		},
		Relations: []*schema.Table{
			ServiceAccountKeys(),
		},
	}
}

func fetchServiceAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListServiceAccountsRequest{
		Name: "projects/" + c.ProjectId,
	}
	gcpClient, err := admin.NewIamClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListServiceAccounts(ctx, req, c.CallOptions...)
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
