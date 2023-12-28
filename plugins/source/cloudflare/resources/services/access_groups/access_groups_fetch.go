package access_groups

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchAccessGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client)
	rc := cloudflare.ZoneIdentifier(svc.ZoneId)

	params := cloudflare.ListAccessGroupsParams{
		ResultInfo: cloudflare.ResultInfo{
			Page:    1,
			PerPage: client.MaxItemsPerPage,
		},
	}

	for {
		resp, info, err := svc.ClientApi.ListAccessGroups(ctx, rc, params)
		if err != nil {
			return err
		}
		res <- resp

		if !info.HasMorePages() {
			break
		}
		params.ResultInfo.Page++
	}
	return nil
}
