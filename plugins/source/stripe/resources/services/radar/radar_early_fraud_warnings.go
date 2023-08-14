package radar

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/stripe/stripe-go/v74"
)

func RadarEarlyFraudWarnings() *schema.Table {
	return &schema.Table{
		Name:        "stripe_radar_early_fraud_warnings",
		Description: `https://stripe.com/docs/api/radar/early_fraud_warnings`,
		Transform:   client.TransformWithStruct(&stripe.RadarEarlyFraudWarning{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchRadarEarlyFraudWarnings,

		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
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
