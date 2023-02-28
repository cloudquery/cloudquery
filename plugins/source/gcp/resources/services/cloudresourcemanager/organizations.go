package cloudresourcemanager

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	crmv1 "google.golang.org/api/cloudresourcemanager/v1"
)

func Organizations() *schema.Table {
	return &schema.Table{
		Name:        "gcp_cloudresourcemanager_organizations",
		Description: `https://cloud.google.com/resource-manager/reference/rest/v1/organizations#Organization`,
		Resolver:    fetchOrganizations,
		Multiplex:   client.OrgMultiplex,
		Transform:   transformers.TransformWithStruct(&crmv1.Organization{}, append(client.Options(), transformers.WithPrimaryKeys("Name"))...),
	}
}

func fetchOrganizations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	res <- c.Org
	return nil
}
