package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Ads() *schema.Table {
	return &schema.Table{
		Name:        "facebookmarketing_ads",
		Resolver:    fetchAds,
		Transform:   client.TransformWithStruct(&rest.Ad{}, transformers.WithPrimaryKeys("Id")),
		Description: "https://developers.facebook.com/docs/marketing-api/reference/adgroup",
	}
}
