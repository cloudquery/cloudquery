package groups

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/cloudquery/plugins/source/okta/resources/services/groups/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
)

func fetchGroupUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	api := meta.(*client.Client)
	grp := parent.Item.(*okta.Group)

	items, resp, err := api.Okta.Group.ListGroupUsers(ctx, grp.Id, query.NewQueryParams(query.WithLimit(200), query.WithAfter("")))
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return nil
	}

	res <- convertGroupUsers(items)

	for resp != nil && resp.HasNextPage() {
		var nextItems []*okta.User
		resp, err = resp.Next(ctx, &nextItems)
		if err != nil {
			return err
		}
		res <- convertGroupUsers(nextItems)
	}
	return nil
}

func convertGroupUsers(list []*okta.User) []*models.GroupUser {
	res := make([]*models.GroupUser, len(list))
	for i := range list {
		res[i] = &models.GroupUser{
			Id: list[i].Id,
		}
	}
	return res
}
