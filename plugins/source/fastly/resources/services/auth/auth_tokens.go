package auth

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
			{
				Name:     "ip",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("IP"),
			},
		},
	}
}
