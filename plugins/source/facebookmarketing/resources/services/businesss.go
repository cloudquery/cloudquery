package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Businesses() *schema.Table {
	return &schema.Table{
		Name:        "facebookmarketing_businesses",
		Resolver:    fetchBusinesss,
		Transform:   client.TransformWithStruct(&rest.Business{}, transformers.WithPrimaryKeys("Id")),
		Description: "https://developers.facebook.com/docs/marketing-api/reference/business/#Reading",
	}
}
