package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
)

func ApplicationGroupAssignments() *schema.Table {
	return &schema.Table{
		Name:     "okta_application_group_assignments",
		Resolver: fetchApplicationGroupAssignments,
		Columns: []schema.Column{
			{
				Name:            "app_id",
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
			{
				Name:     "last_updated",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastUpdated"),
			},
			{
				Name:     "priority",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Priority"),
			},
			{
				Name:     "profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Profile"),
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchApplicationGroupAssignments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	api := meta.(*client.Client)
	app := parent.Item.(*okta.Application)
	items, resp, err := api.Okta.Application.ListApplicationGroupAssignments(ctx, app.Id, query.NewQueryParams(query.WithLimit(200), query.WithAfter("")))
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return nil
	}
	res <- items
	for resp != nil && resp.HasNextPage() {
		var nextItems []*okta.ApplicationGroupAssignment
		resp, err = resp.Next(ctx, &nextItems)
		if err != nil {
			return err
		}
		res <- nextItems
	}
	return nil
}
