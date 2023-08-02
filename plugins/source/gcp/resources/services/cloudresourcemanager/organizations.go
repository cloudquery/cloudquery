package cloudresourcemanager

import (
	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Organizations() *schema.Table {
	return &schema.Table{
		Name:        "gcp_cloudresourcemanager_organizations",
		Description: `https://cloud.google.com/resource-manager/reference/rest/v1/organizations#Organization`,
		Resolver:    fetchOrganizations,
		Multiplex:   client.OrgMultiplex,
		Transform:   client.TransformWithStruct(&resourcemanagerpb.Organization{}, transformers.WithPrimaryKeys("Name")),
	}
}

func fetchOrganizations(_ context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	res <- c.Org
	return nil
}
