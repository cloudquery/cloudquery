package applications

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/okta/okta-sdk-golang/v3/okta"
)

func groupAssignments() *schema.Table {
	return &schema.Table{
		Name:      "okta_application_group_assignments",
		Resolver:  fetchGroupAssignments,
		Transform: client.TransformWithStruct(&okta.ApplicationGroupAssignment{}),
		Columns:   schema.ColumnList{appIDColumn},
	}
}

func fetchGroupAssignments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) (err error) {
	defer func() {
		err = client.ProcessOktaAPIError(err)
	}()

	cl := meta.(*client.Client)
	app := parent.Item.(*okta.Application)

	req := cl.ApplicationApi.ListApplicationGroupAssignments(ctx, *app.Id).Limit(200)
	items, resp, err := cl.ApplicationApi.ListApplicationGroupAssignmentsExecute(req)
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return nil
	}
	res <- items

	for resp != nil && resp.HasNextPage() {
		var nextItems []okta.ApplicationGroupAssignment
		resp, err = resp.Next(&nextItems)
		if err != nil {
			return err
		}
		res <- nextItems
	}
	return nil
}
