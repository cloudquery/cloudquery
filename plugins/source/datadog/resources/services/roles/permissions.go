package roles

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Permissions() *schema.Table {
	return &schema.Table{
		Name:      "datadog_permissions",
		Resolver:  fetchPermissions,
		Multiplex: client.AccountMultiplex,
		Transform: client.TransformWithStruct(&datadogV2.Permission{}, transformers.WithPrimaryKeys("Id")),
		Columns:   schema.ColumnList{client.AccountNameColumn},
	}
}

func fetchPermissions(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	ctx = c.BuildContextV2(ctx)
	resp, _, err := c.DDServices.RolesAPI.ListPermissions(ctx)
	if err != nil {
		return err
	}
	res <- resp.GetData()
	return nil
}
