package account

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
	}
}
