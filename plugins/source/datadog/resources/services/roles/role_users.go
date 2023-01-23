package roles

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func RoleUsers() *schema.Table {
	return &schema.Table{
		Name:      "datadog_role_users",
		Resolver:  fetchRoleUsers,
		Transform: transformers.TransformWithStruct(&datadogV2.User{}),
		Columns: []schema.Column{
			{
				Name:     "account_name",
				Type:     schema.TypeString,
				Resolver: client.ResolveAccountName,
			},
		},
	}
}
