package file_links

import (
	"context"

	"fmt"
	"strconv"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func FileLinks() *schema.Table {
	return &schema.Table{
		Name:        "stripe_file_links",
		Description: `https://stripe.com/docs/api/file_links`,
		Transform:   transformers.TransformWithStruct(&stripe.FileLink{}, client.SharedTransformers(transformers.WithSkipFields("APIResource", "ID"))...),
		Resolver:    fetchFileLinks("file_links"),

		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "created",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Created"),
				CreationOptions: schema.ColumnCreationOptions{
					IncrementalKey: true,
				},
			},
		},
		IsIncremental: true,
	}
}

func fetchFileLinks(tableName string) schema.TableResolver {
	return func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
		cl := meta.(*client.Client)

		lp := &stripe.FileLinkListParams{}

		if cl.Backend != nil {
			value, err := cl.Backend.Get(ctx, tableName, cl.ID())
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

		it := cl.Services.FileLinks.List(lp)
		for it.Next() {
			data := it.FileLink()
			lp.Created = client.MaxInt64(lp.Created, &data.Created)
			res <- data
		}

		err := it.Err()
		if cl.Backend != nil && err == nil && lp.Created != nil {
			return cl.Backend.Set(ctx, tableName, cl.ID(), strconv.FormatInt(*lp.Created, 10))
		}
		return err
	}
}
