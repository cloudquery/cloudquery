package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func AdPlacePageSets() *schema.Table {
	return &schema.Table{
		Name:        "facebookmarketing_ad_place_page_sets",
		Resolver:    fetchAdPlacePageSets,
		Transform:   client.TransformWithStruct(&rest.AdPlacePageSet{}, transformers.WithPrimaryKeys("Id")),
		Description: "https://developers.facebook.com/docs/marketing-api/reference/ad-place-page-set#Reading",
	}
}
