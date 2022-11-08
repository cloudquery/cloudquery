package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
)

func GroupUsers() *schema.Table {
	return &schema.Table{
		Name:     "okta_group_users",
		Resolver: fetchGroupUsers,
		Columns: []schema.Column{
			{
				Name:            "group_id",
				Type:            schema.TypeString,
				Resolver:        schema.ParentColumnResolver("id"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "id",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
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
	res <- items
	for resp != nil && resp.HasNextPage() {
		var nextItems []*okta.User
		resp, err = resp.Next(ctx, &nextItems)
		if err != nil {
			return err
		}
		res <- nextItems
	}
	return nil
}
