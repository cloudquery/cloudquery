package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
)

func ApplicationUsers() *schema.Table {
	return &schema.Table{
		Name:     "okta_application_users",
		Resolver: fetchApplicationUsers,
		Columns: []schema.Column{
			{
				Name:            "app_id",
				Type:            schema.TypeString,
				Resolver:        schema.ParentColumnResolver("id"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:     "created",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Created"),
			},
			{
				Name:     "credentials",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Credentials"),
			},
			{
				Name:     "external_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ExternalId"),
			},
			{
				Name:            "id",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:     "last_sync",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastSync"),
			},
			{
				Name:     "last_updated",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastUpdated"),
			},
			{
				Name:     "password_changed",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("PasswordChanged"),
			},
			{
				Name:     "profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Profile"),
			},
			{
				Name:     "scope",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Scope"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "status_changed",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("StatusChanged"),
			},
			{
				Name:     "sync_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SyncState"),
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
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
	res <- items
	for resp != nil && resp.HasNextPage() {
		var nextItems []*okta.AppUser
		resp, err = resp.Next(ctx, &nextItems)
		if err != nil {
			return err
		}
		res <- nextItems
	}
	return nil
}
