package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func SavedAudiences() *schema.Table {
	return &schema.Table{
		Name:        "facebookmarketing_saved_audiences",
		Resolver:    fetchSavedAudiences,
		Transform:   client.TransformWithStruct(&rest.SavedAudience{}, transformers.WithPrimaryKeys("Id")),
		Description: "https://developers.facebook.com/docs/marketing-api/reference/saved-audience/#Reading",
	}
}
