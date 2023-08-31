package extension_schemas

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchExtensionSchemas(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	more := true
	var offset uint
	for more {
		response, err := cqClient.PagerdutyClient.ListExtensionSchemasWithContext(ctx, pagerduty.ListExtensionSchemaOptions{
			Limit:  client.MaxPaginationLimit,
			Offset: offset,
		})
		if err != nil {
			return err
		}

		if len(response.ExtensionSchemas) == 0 {
			return nil
		}

		res <- response.ExtensionSchemas

		offset += uint(len(response.ExtensionSchemas))
		more = response.More
	}

	return nil
}
