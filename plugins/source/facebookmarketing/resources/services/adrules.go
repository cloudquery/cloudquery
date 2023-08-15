package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Adrules() *schema.Table {
	return &schema.Table{
		Name:        "facebookmarketing_adrules",
		Resolver:    fetchAdrules,
		Transform:   client.TransformWithStruct(&rest.Adrule{}, transformers.WithPrimaryKeys("Id")),
		Description: "https://developers.facebook.com/docs/marketing-api/reference/ad-rule#Reading",
	}
}
