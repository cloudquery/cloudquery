package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
)

func Applications() *schema.Table {
	return &schema.Table{
		Name:     "okta_applications",
		Resolver: fetchApplications,
		Columns: []schema.Column{
			{
				Name:     "accessibility",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Accessibility"),
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
				Name:     "features",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Features"),
			},
			{
				Name:            "id",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:     "label",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Label"),
			},
			{
				Name:     "last_updated",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastUpdated"),
			},
			{
				Name:     "licensing",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Licensing"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Profile"),
			},
			{
				Name:     "settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Settings"),
			},
			{
				Name:     "sign_on_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SignOnMode"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "visibility",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Visibility"),
			},
		},
		Relations: []*schema.Table{
			ApplicationUsers(),
			ApplicationGroupAssignments(),
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchApplications(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	api := meta.(*client.Client)
	items, resp, err := api.Okta.Application.ListApplications(ctx, query.NewQueryParams(query.WithLimit(200), query.WithAfter("")))
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return nil
	}
	res <- items
	for resp != nil && resp.HasNextPage() {
		var nextItems []*okta.App
		resp, err = resp.Next(ctx, &nextItems)
		if err != nil {
			return err
		}
		res <- nextItems
	}
	return nil
}
