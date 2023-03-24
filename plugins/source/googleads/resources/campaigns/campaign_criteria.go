package campaigns

import (
	"github.com/cloudquery/cloudquery/plugins/source/googleads/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

func campaignCriteria() *schema.Table {
	return &schema.Table{
		Name:        "googleads_campaign_criteria",
		Description: client.APIRef + "/CampaignCriterion",
		Transform: client.TransformWithStruct(new(resources.CampaignCriterion),
			transformers.WithPrimaryKeys("Campaign", "ResourceName"),
		),
		Columns:  schema.ColumnList{client.CustomerID, client.IDColumn("CriterionId")},
		Resolver: client.Fetcher((*services.GoogleAdsRow).GetCampaignCriterion),
	}
}
