package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Businesses() *schema.Table {
	return &schema.Table{
		Name:        "facebookmarketing_businesses",
		Resolver:    fetchBusinesss,
		Transform:   transformers.TransformWithStruct(&rest.Business{}, append(client.TransformerOptions(), transformers.WithPrimaryKeys("Id"))...),
		Description: "https://developers.facebook.com/docs/marketing-api/reference/business/#Reading",
	}
}
