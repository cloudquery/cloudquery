package groups

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/cloudquery/plugins/source/okta/resources/services/groups/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/okta/okta-sdk-golang/v3/okta"
)

func users() *schema.Table {
	return &schema.Table{
		Name:      "okta_group_users",
		Resolver:  fetchUsers,
		Transform: client.TransformWithStruct(&models.GroupUser{}),
		Columns: []schema.Column{
			{
				Name:       "group_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("id"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) (err error) {
	defer func() {
		err = client.ProcessOktaAPIError(err)
	}()

	cl := meta.(*client.Client)
	grp := parent.Item.(okta.Group)

	req := cl.GroupApi.ListGroupUsers(ctx, *grp.Id).Limit(200)

	items, resp, err := cl.GroupApi.ListGroupUsersExecute(req)
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return nil
	}

	res <- convertUsers(items)

	for resp != nil && resp.HasNextPage() {
		var nextItems []okta.User
		resp, err = resp.Next(&nextItems)
		if err != nil {
			return err
		}
		res <- convertUsers(nextItems)
	}
	return nil
}

func convertUsers(list []okta.User) []models.GroupUser {
	res := make([]models.GroupUser, len(list))
	for i := range list {
		res[i] = models.GroupUser{Id: *list[i].Id}
	}
	return res
}
