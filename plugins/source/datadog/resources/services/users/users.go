// Code generated by codegen; DO NOT EDIT.

package users

import (
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:      "datadog_users",
		Resolver:  fetchUsers,
		Multiplex: client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_name",
				Type:     schema.TypeString,
				Resolver: client.ResolveAccountName,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "attributes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Attributes"),
			},
			{
				Name:     "relationships",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Relationships"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},

		Relations: []*schema.Table{
			UserPermissions(),
			UserOrganizations(),
		},
	}
}
