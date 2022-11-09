package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
)

func Groups() *schema.Table {
	return &schema.Table{
		Name:     "okta_groups",
		Resolver: fetchGroups,
		Columns: []schema.Column{
			{
				Name:     "created",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Created"),
			},
			{
				Name:            "id",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:     "last_membership_updated",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastMembershipUpdated"),
			},
			{
				Name:     "last_updated",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastUpdated"),
			},
			{
				Name:     "object_class",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ObjectClass"),
			},
			{
				Name:     "profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Profile"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
		Relations: []*schema.Table{
			GroupUsers(),
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	api := meta.(*client.Client)
	items, resp, err := api.Okta.Group.ListGroups(ctx, query.NewQueryParams(query.WithLimit(200), query.WithAfter("")))
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return nil
	}
	res <- items
	for resp != nil && resp.HasNextPage() {
		var nextItems []*okta.Group
		resp, err = resp.Next(ctx, &nextItems)
		if err != nil {
			return err
		}
		res <- nextItems
	}
	return nil
}
