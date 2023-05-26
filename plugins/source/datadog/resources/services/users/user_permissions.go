package users

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func UserPermissions() *schema.Table {
	return &schema.Table{
		Name:      "datadog_user_permissions",
		Resolver:  fetchUserPermissions,
		Transform: transformers.TransformWithStruct(&datadogV2.Permission{}),
		Columns: []schema.Column{
			{
				Name:     "account_name",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveAccountName,
			},
		},
	}
}

func fetchUserPermissions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(datadogV2.User)
	c := meta.(*client.Client)
	ctx = c.BuildContextV2(ctx)
	resp, _, err := c.DDServices.UsersAPI.ListUserPermissions(ctx, *p.Id)
	if err != nil {
		return err
	}
	res <- resp.GetData()
	return nil
}
