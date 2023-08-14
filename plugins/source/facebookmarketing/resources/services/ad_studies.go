package services

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func AdStudies() *schema.Table {
	return &schema.Table{
		Name:        "facebookmarketing_ad_studies",
		Resolver:    fetchAdStudies,
		Transform:   client.TransformWithStruct(&rest.AdStudy{}, transformers.WithPrimaryKeys("Id")),
		Description: "https://developers.facebook.com/docs/marketing-api/reference/ad-study/#Reading",
		Columns: []schema.Column{
			{
				Name:       "account_id",
				Resolver:   client.ResolveAccountId,
				Type:       arrow.BinaryTypes.String,
				PrimaryKey: true,
			},
		},
	}
}
