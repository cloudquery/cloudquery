package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Adimages() *schema.Table {
	return &schema.Table{
		Name:        "facebookmarketing_adimages",
		Resolver:    fetchAdimages,
		Transform:   transformers.TransformWithStruct(&rest.Adimage{}, append(client.TransformerOptions(), transformers.WithPrimaryKeys("Id"))...),
		Description: "https://developers.facebook.com/docs/marketing-api/reference/ad-image#Reading",
	}
}
