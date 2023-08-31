package integration

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
)

func Integrations() *schema.Table {
	return &schema.Table{
		Name:                "snyk_integrations",
		Description:         `https://snyk.docs.apiary.io/#reference/integrations/integrations/list`,
		Resolver:            fetchIntegrations,
		PreResourceResolver: getIntegration,
		Multiplex:           client.ByOrganization,
		Transform:           transformers.TransformWithStruct(&snyk.Integration{}, transformers.WithPrimaryKeys("ID")),
		Columns: []schema.Column{
			client.OrganizationID,
			{
				Name:     "settings",
				Type:     types.ExtensionTypes.JSON,
				Resolver: getIntegrationSettings,
			},
		},
	}
}

func fetchIntegrations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	integrations, _, err := c.Integrations.List(ctx, c.OrganizationID)
	if err != nil {
		return err
	}

	for typ, id := range integrations {
		if len(id) > 0 {
			res <- typ
		}
	}

	return nil
}

func getIntegration(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)

	integration, _, err := c.Integrations.GetByType(ctx, c.OrganizationID, resource.Item.(snyk.IntegrationType))
	if err != nil {
		return err
	}

	resource.SetItem(integration)

	return nil
}

func getIntegrationSettings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, column schema.Column) error {
	c := meta.(*client.Client)

	settings, _, err := c.Integrations.GetSettings(ctx, c.OrganizationID, resource.Item.(*snyk.Integration).ID)
	if err != nil {
		return err
	}

	return resource.Set(column.Name, settings)
}
