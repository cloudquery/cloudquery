package users

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func UserPermissions() *schema.Table {
	return &schema.Table{
		Name:      "datadog_user_permissions",
		Resolver:  fetchUserPermissions,
		Transform: transformers.TransformWithStruct(&datadogV2.Permission{}),
		Columns: []schema.Column{
			{
				Name:     "account_name",
				Type:     schema.TypeString,
				Resolver: client.ResolveAccountName,
			},
		},
	}
}
