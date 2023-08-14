package services

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func OfflineConversionDataSets() *schema.Table {
	return &schema.Table{
		Name:      "facebookmarketing_offline_conversion_data_sets",
		Resolver:  fetchOfflineConversionDataSets,
		Transform: client.TransformWithStruct(&rest.OfflineConversionDataSet{}, transformers.WithPrimaryKeys("Id")),
		Columns: []schema.Column{
			{
				Name:       "account_id",
				Resolver:   client.ResolveAccountId,
				Type:       arrow.BinaryTypes.String,
				PrimaryKey: true,
			},
		},
		Description: "https://developers.facebook.com/docs/marketing-api/reference/offline-conversion-data-set/#Reading",
	}
}
