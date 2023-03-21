package plugin

import (
	"strings"

	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/resources/services"
	"github.com/cloudquery/plugin-sdk/caser"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

var customExceptions = map[string]string{
	"facebookmarketing": "Facebook Marketing",
	"adaccounts":        "Ad Accounts",
	"adcreatives":       "Ad Creatives",
	"adimages":          "Ad Images",
	"adlabels":          "Ad Labels",
	"adplayables":       "Ad Playables",
	"adcloudplayables":  "Ad Cloud Playables",
	"advideos":          "Ad Videos",
	"adspixels":         "Ads Pixels",
	"adsets":            "Ad Sets",
	"customaudiences":   "Custom Audiences",
	"customconversions": "Custom Conversions",
}

func titleTransformer(table *schema.Table) string {
	if table.Title != "" {
		return table.Title
	}
	exceptions := make(map[string]string)
	for k, v := range source.DefaultTitleExceptions {
		exceptions[k] = v
	}
	for k, v := range customExceptions {
		exceptions[k] = v
	}
	csr := caser.New(caser.WithCustomExceptions(exceptions))
	t := csr.ToTitle(table.Name)
	return strings.Trim(strings.ReplaceAll(t, "  ", " "), " ")
}

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"facebookmarketing",
		Version,
		schema.Tables{
			services.Campaigns(),
			services.Adsets(),
			services.Ads(),
			services.Adcreatives(),
			services.Adimages(),
			services.Advideos(),
			services.AdStudies(),
			services.Customaudiences(),
			services.Users(),
			services.Adaccounts(),
			services.AdPlacePageSets(),
			services.Adcloudplayables(),
			services.Adlabels(),
			services.Adplayables(),
			services.Adrules(),
			services.Adspixels(),
			services.AdvertisableApplications(),
			services.Businesses(),
			services.BroadTargetingCategoriess(),
			services.ConnectedInstagramAccounts(),
			services.Customconversions(),
			services.MaxBids(),
			services.OfflineConversionDataSets(),
			services.PromotePages(),
			services.PublisherBlockLists(),
			services.ReachFrequencyPredictions(),
			services.SavedAudiences(),
		},
		client.New,
		source.WithTitleTransformer(titleTransformer),
	)
}
