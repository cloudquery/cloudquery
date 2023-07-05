package users

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:      "datadog_users",
		Resolver:  fetchUsers,
		Multiplex: client.AccountMultiplex,
		Transform: client.TransformWithStruct(&datadogV2.User{}, transformers.WithPrimaryKeys("Id")),
		Columns:   schema.ColumnList{client.AccountNameColumn},
		Relations: schema.Tables{permissions()},
	}
}

func fetchUsers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	ctx = c.BuildContextV2(ctx)
	resp, _, err := c.DDServices.UsersAPI.ListUsers(ctx)
	if err != nil {
		return err
	}
	res <- resp.GetData()
	return nil
}
