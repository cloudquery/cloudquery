package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/resources/services"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

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
	)
}
