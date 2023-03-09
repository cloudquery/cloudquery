package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SavedAudiences() *schema.Table {
	return &schema.Table{
		Name:        "facebookmarketing_saved_audiences",
		Resolver:    fetchSavedAudiences,
		Transform:   transformers.TransformWithStruct(&rest.SavedAudience{}, append(client.TransformerOptions(), transformers.WithPrimaryKeys("Id"))...),
		Description: "https://developers.facebook.com/docs/marketing-api/reference/saved-audience/#Reading",
	}
}
