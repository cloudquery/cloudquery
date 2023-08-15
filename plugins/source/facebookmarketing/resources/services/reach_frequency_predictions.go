package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ReachFrequencyPredictions() *schema.Table {
	return &schema.Table{
		Name:        "facebookmarketing_reach_frequency_predictions",
		Resolver:    fetchReachFrequencyPredictions,
		Transform:   client.TransformWithStruct(&rest.ReachFrequencyPrediction{}, transformers.WithPrimaryKeys("Id")),
		Description: "https://developers.facebook.com/docs/marketing-api/reference/reach-frequency-prediction/#Reading",
	}
}
