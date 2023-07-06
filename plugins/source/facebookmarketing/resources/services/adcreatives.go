package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Adcreatives() *schema.Table {
	return &schema.Table{
		Name:        "facebookmarketing_adcreatives",
		Resolver:    fetchAdcreatives,
		Transform:   client.TransformWithStruct(&rest.Adcreative{}, transformers.WithPrimaryKeys("Id")),
		Description: "https://developers.facebook.com/docs/marketing-api/reference/ad-creative#Reading",
	}
}
