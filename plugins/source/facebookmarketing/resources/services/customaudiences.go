package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Customaudiences() *schema.Table {
	return &schema.Table{
		Name:        "facebookmarketing_customaudiences",
		Resolver:    fetchCustomaudiences,
		Transform:   client.TransformWithStruct(&rest.Customaudience{}, transformers.WithPrimaryKeys("Id")),
		Description: "https://developers.facebook.com/docs/marketing-api/reference/custom-audience#Reading",
	}
}
