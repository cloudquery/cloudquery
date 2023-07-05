package access_groups

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchAccessGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client)
	zoneID := svc.ZoneId

	pagination := cloudflare.PaginationOptions{
		Page:    1,
		PerPage: client.MaxItemsPerPage,
	}

	for {
		resp, info, err := svc.ClientApi.ZoneLevelAccessGroups(ctx, zoneID, pagination)
		if err != nil {
			return err
		}
		res <- resp

		if !info.HasMorePages() {
			break
		}
		pagination.Page++
	}
	return nil
}
