package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Adrules() *schema.Table {
	return &schema.Table{
		Name:        "facebookmarketing_adrules",
		Resolver:    fetchAdrules,
		Transform:   client.TransformWithStruct(&rest.Adrule{}, transformers.WithPrimaryKeys("Id")),
		Description: "https://developers.facebook.com/docs/marketing-api/reference/ad-rule#Reading",
	}
}
