package campaigns

import (
	"github.com/cloudquery/cloudquery/plugins/source/googleads/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

func campaignLabels() *schema.Table {
	return &schema.Table{
		Name:        "googleads_campaign_labels",
		Description: client.APIRef + "/CampaignLabel",
		Transform: client.TransformWithStruct(new(resources.CampaignLabel),
			transformers.WithPrimaryKeys("Campaign", "ResourceName"),
		),
		Columns:  schema.ColumnList{client.CustomerID},
		Resolver: client.Fetcher((*services.GoogleAdsRow).GetCampaignLabel),
	}
}
