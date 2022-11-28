package applications

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/cloudquery/plugins/source/okta/resources/services/groups/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
)

func fetchApplicationUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	api := meta.(*client.Client)
	app := parent.Item.(*okta.Application)
	items, resp, err := api.Okta.Application.ListApplicationUsers(ctx, app.Id, query.NewQueryParams(query.WithLimit(200), query.WithAfter("")))
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return nil
	}
	res <- convertApplicationUsers(items)

	for resp != nil && resp.HasNextPage() {
		var nextItems []*okta.AppUser
		resp, err = resp.Next(ctx, &nextItems)
		if err != nil {
			return err
		}
		res <- convertApplicationUsers(nextItems)
	}
	return nil
}

func convertApplicationUsers(list []*okta.AppUser) []*models.ApplicationUser {
	res := make([]*models.ApplicationUser, len(list))
	for i := range list {
		res[i] = &models.ApplicationUser{
			AppUser: list[i],
		}
	}
	return res
}
