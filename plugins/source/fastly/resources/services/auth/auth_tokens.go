package auth

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/fastly/go-fastly/v7/fastly"
)

func AuthTokens() *schema.Table {
	return &schema.Table{
		Name:        "fastly_auth_tokens",
		Description: `https://developer.fastly.com/reference/api/auth-tokens/`,
		Resolver:    fetchAuthTokens,
		Transform:   transformers.TransformWithStruct(&fastly.Token{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "ip",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IP"),
			},
		},
	}
}
