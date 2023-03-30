package ads

import (
	"github.com/cloudquery/cloudquery/plugins/source/googleads/client"
	"github.com/cloudquery/cloudquery/plugins/source/googleads/gaql"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

var groupOptions = &gaql.Options{
	Skip:   []string{"TargetingSetting"},
	Expand: []string{"AudienceSetting"},
}

func Groups() *schema.Table {
	return &schema.Table{
		Name:        "googleads_ad_groups",
		Description: client.APIRef + "/AdGroup",
		Multiplex:   client.MultiplexByCustomer,
		Transform: client.TransformWithStruct(new(resources.AdGroup),
			transformers.WithPrimaryKeys("Id", "ResourceName"),
			transformers.WithSkipFields(groupOptions.Skip...),
		),
		Columns:   schema.ColumnList{client.CustomerID},
		Resolver:  client.Fetcher((*services.GoogleAdsRow).GetAdGroup, groupOptions),
		Relations: schema.Tables{groupCriteria(), groupAds(), groupLabels()},
	}
}
