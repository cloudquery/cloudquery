package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func AdPlacePageSets() *schema.Table {
	return &schema.Table{
		Name:        "facebookmarketing_ad_place_page_sets",
		Resolver:    fetchAdPlacePageSets,
		Transform:   transformers.TransformWithStruct(&rest.AdPlacePageSet{}, append(client.TransformerOptions(), transformers.WithPrimaryKeys("Id"))...),
		Description: "https://developers.facebook.com/docs/marketing-api/reference/ad-place-page-set#Reading",
	}
}
