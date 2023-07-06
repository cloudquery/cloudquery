package files

import (
	"context"
	"fmt"
	"strconv"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/stripe/stripe-go/v74"
)

func Files() *schema.Table {
	return &schema.Table{
		Name:        "stripe_files",
		Description: `https://stripe.com/docs/api/files`,
		Transform:   client.TransformWithStruct(&stripe.File{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchFiles,

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

func fetchFiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.FileListParams{}

	const key = "files"

	if cl.Backend != nil {
		value, err := cl.Backend.GetKey(ctx, key)
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

	it := cl.Services.Files.List(lp)
	for it.Next() {
		data := it.File()
		lp.Created = client.MaxInt64(lp.Created, &data.Created)
		res <- data
	}

	err := it.Err()
	if cl.Backend != nil && err == nil && lp.Created != nil {
		return cl.Backend.SetKey(ctx, key, strconv.FormatInt(*lp.Created, 10))
	}
	return err
}
