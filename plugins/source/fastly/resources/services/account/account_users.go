package account

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/fastly/go-fastly/v7/fastly"
)

func AccountUsers() *schema.Table {
	return &schema.Table{
		Name:        "fastly_account_users",
		Description: `https://developer.fastly.com/reference/api/account/user/`,
		Resolver:    fetchAccountUsers,
		Transform:   transformers.TransformWithStruct(&fastly.User{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
