package roles

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func RolePermissions() *schema.Table {
	return &schema.Table{
		Name:      "datadog_role_permissions",
		Resolver:  fetchRolePermissions,
		Transform: transformers.TransformWithStruct(&datadogV2.Permission{}),
		Columns: []schema.Column{
			{
				Name:     "account_name",
				Type:     schema.TypeString,
				Resolver: client.ResolveAccountName,
			},
			{
				Name:     "attributes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Attributes"),
			},
		},
	}
}

func fetchRolePermissions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(datadogV2.Role)
	c := meta.(*client.Client)
	ctx = c.BuildContextV2(ctx)
	resp, _, err := c.DDServices.RolesAPI.ListRolePermissions(ctx, *p.Id)
	if err != nil {
		return err
	}
	res <- resp.GetData()
	return nil
}
