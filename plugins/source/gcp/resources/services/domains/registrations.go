package domains

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/domains/apiv1beta1/domainspb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	domains "cloud.google.com/go/domains/apiv1beta1"

	"fmt"
)

func Registrations() *schema.Table {
	return &schema.Table{
		Name:        "gcp_domains_registrations",
		Description: `https://cloud.google.com/domains/docs/reference/rest/v1beta1/projects.locations.registrations#Registration`,
		Resolver:    fetchRegistrations,
		Multiplex:   client.ProjectMultiplexEnabledServices("domains.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.Registration{}, append(client.Options(), transformers.WithPrimaryKeys("Name"))...),
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

func fetchRegistrations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListRegistrationsRequest{
		Parent: fmt.Sprintf("projects/%s/locations/-", c.ProjectId),
	}
	gcpClient, err := domains.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListRegistrations(ctx, req, c.CallOptions...)
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
