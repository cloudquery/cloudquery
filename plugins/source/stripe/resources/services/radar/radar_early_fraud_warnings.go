package radar

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func RadarEarlyFraudWarnings() *schema.Table {
	return &schema.Table{
		Name:        "stripe_radar_early_fraud_warnings",
		Description: `https://stripe.com/docs/api/radar_early_fraud_warnings`,
		Transform:   transformers.TransformWithStruct(&stripe.RadarEarlyFraudWarning{}, client.SharedTransformers(transformers.WithSkipFields("APIResource", "ID"))...),
		Resolver:    fetchRadarEarlyFraudWarnings,

		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchRadarEarlyFraudWarnings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.RadarEarlyFraudWarningListParams{}

	it := cl.Services.RadarEarlyFraudWarnings.List(lp)
	for it.Next() {
		res <- it.RadarEarlyFraudWarning()
	}

	return it.Err()
}
