package users

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:      "datadog_users",
		Resolver:  fetchUsers,
		Multiplex: client.AccountMultiplex,
		Transform: transformers.TransformWithStruct(&datadogV2.User{}),
		Columns: []schema.Column{
			{
				Name:       "account_name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAccountName,
				PrimaryKey: true,
			},
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Id"),
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			UserPermissions(),
		},
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
