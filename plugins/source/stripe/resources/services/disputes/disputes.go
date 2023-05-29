package disputes

import (
	"context"
	"fmt"
	"strconv"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/stripe/stripe-go/v74"
)

func Disputes() *schema.Table {
	return &schema.Table{
		Name:        "stripe_disputes",
		Description: `https://stripe.com/docs/api/disputes`,
		Transform:   client.TransformWithStruct(&stripe.Dispute{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchDisputes,

		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
			{
				Name:           "created",
				Type:           arrow.FixedWidthTypes.Timestamp_us,
				Resolver:       schema.PathResolver("Created"),
				IncrementalKey: true,
			},
		},
		IsIncremental: true,
	}
}

func fetchDisputes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.DisputeListParams{}

	const key = "disputes"

	if cl.Backend != nil {
		value, err := cl.Backend.Get(ctx, key, cl.ID())
		if err != nil {
			return fmt.Errorf("failed to retrieve state from backend: %w", err)
		}
		if value != "" {
			vi, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return fmt.Errorf("retrieved invalid state value: %q %w", value, err)
			}
			lp.Created = &vi
		}
	}

	it := cl.Services.Disputes.List(lp)
	for it.Next() {
		data := it.Dispute()
		lp.Created = client.MaxInt64(lp.Created, &data.Created)
		res <- data
	}

	err := it.Err()
	if cl.Backend != nil && err == nil && lp.Created != nil {
		return cl.Backend.Set(ctx, key, cl.ID(), strconv.FormatInt(*lp.Created, 10))
	}
	return err
}
