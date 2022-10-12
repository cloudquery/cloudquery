// Code generated by codegen; DO NOT EDIT.

package organizations

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Accounts() *schema.Table {
	return &schema.Table{
		Name:      "aws_organizations_accounts",
		Resolver:  fetchOrganizationsAccounts,
		Multiplex: client.ServiceAccountRegionMultiplexer("organizations"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveAccountTags,
			},
			{
				Name:     "email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Email"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "joined_method",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("JoinedMethod"),
			},
			{
				Name:     "joined_timestamp",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("JoinedTimestamp"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
		},
	}
}
