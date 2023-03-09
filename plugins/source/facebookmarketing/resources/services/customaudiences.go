package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Customaudiences() *schema.Table {
	return &schema.Table{
		Name:        "facebookmarketing_customaudiences",
		Resolver:    fetchCustomaudiences,
		Transform:   transformers.TransformWithStruct(&rest.Customaudience{}, append(client.TransformerOptions(), transformers.WithPrimaryKeys("Id"))...),
		Description: "https://developers.facebook.com/docs/marketing-api/reference/custom-audience#Reading",
	}
}
