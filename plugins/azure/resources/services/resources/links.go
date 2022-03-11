package resources

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-09-01/links"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ResourcesLinks() *schema.Table {
	return &schema.Table{
		Name:          "azure_resources_links",
		Description:   "Azure resource links",
		Resolver:      fetchResourcesLinks,
		Multiplex:     client.SubscriptionMultiplex,
		DeleteFilter:  client.DeleteSubscriptionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "id",
				Description: "The fully qualified ID of the resource link.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the resource link.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The resource link object.",
				Type:        schema.TypeString,
				Resolver:    resolveResourceLinksType,
			},
			{
				Name:        "source_id",
				Description: "The fully qualified ID of the source resource in the link.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.SourceID"),
			},
			{
				Name:        "target_id",
				Description: "The fully qualified ID of the target resource in the link.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.TargetID"),
			},
			{
				Name:        "notes",
				Description: "Notes about the resource link.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.Notes"),
			},
		},
	}
}

func fetchResourcesLinks(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Resources.Links
	response, err := svc.ListAtSubscription(ctx, "")
	if err != nil {
		return diag.WrapError(err)
	}
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}

func resolveResourceLinksType(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	link := r.Item.(links.ResourceLink)
	if link.Type == nil {
		return nil
	}
	return r.Set(c.Name, fmt.Sprintf("%s", link.Type))
}
