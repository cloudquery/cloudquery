package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ReachFrequencyPredictions() *schema.Table {
	return &schema.Table{
		Name:        "facebookmarketing_reach_frequency_predictions",
		Resolver:    fetchReachFrequencyPredictions,
		Transform:   transformers.TransformWithStruct(&rest.ReachFrequencyPrediction{}, append(client.TransformerOptions(), transformers.WithPrimaryKeys("Id"))...),
		Description: "https://developers.facebook.com/docs/marketing-api/reference/reach-frequency-prediction/#Reading",
	}
}
