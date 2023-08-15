package services

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func DependenciesResolver(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cqClient := meta.(*client.Client)

	item := resource.Item.(pagerduty.Service)

	response, err := cqClient.PagerdutyClient.ListTechnicalServiceDependenciesWithContext(ctx, item.ID)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, response.Relationships)
}
