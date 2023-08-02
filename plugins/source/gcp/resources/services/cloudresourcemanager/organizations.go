package cloudresourcemanager

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	crmv1 "google.golang.org/api/cloudresourcemanager/v1"
)

func Organizations() *schema.Table {
	return &schema.Table{
		Name:        "gcp_cloudresourcemanager_organizations",
		Description: `https://cloud.google.com/resource-manager/reference/rest/v1/organizations#Organization`,
		Resolver:    fetchOrganizations,
		Multiplex:   client.OrgMultiplex,
		Transform:   client.TransformWithStruct(&crmv1.Organization{}, transformers.WithPrimaryKeys("Name")),
	}
}

func fetchOrganizations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	res <- c.Org
	return nil
}
