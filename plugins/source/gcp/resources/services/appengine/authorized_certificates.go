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

func AuthorizedCertificates() *schema.Table {
	return &schema.Table{
		Name:        "gcp_appengine_authorized_certificates",
		Description: `https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.authorizedCertificates#AuthorizedCertificate`,
		Resolver:    fetchAuthorizedCertificates,
		Multiplex:   client.ProjectMultiplexEnabledServices("appengine.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.AuthorizedCertificate{}, transformers.WithPrimaryKeys("Name")),
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

func fetchAuthorizedCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListAuthorizedCertificatesRequest{
		Parent: "apps/" + c.ProjectId,
	}
	gcpClient, err := appengine.NewAuthorizedCertificatesClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListAuthorizedCertificates(ctx, req, c.CallOptions...)
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
