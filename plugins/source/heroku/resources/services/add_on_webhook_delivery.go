// Code generated by codegen; DO NOT EDIT.

package services

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/plugin-sdk/schema"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/pkg/errors"
)

func AddOnWebhookDeliveries() *schema.Table {
	return &schema.Table{
		Name:        "heroku_add_on_webhook_deliveries",
		Description: `https://devcenter.heroku.com/articles/platform-api-reference#add-on-webhook-delivery-attributes`,
		Resolver:    fetchAddOnWebhookDeliveries,
		Columns: []schema.Column{
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "event",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Event"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "last_attempt",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LastAttempt"),
			},
			{
				Name:     "next_attempt_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("NextAttemptAt"),
			},
			{
				Name:     "num_attempts",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumAttempts"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "webhook",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Webhook"),
			},
		},
	}
}

func fetchAddOnWebhookDeliveries(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextRange := &heroku.ListRange{
		Field: "id",
		Max:   1000,
	}
	items := make([]heroku.AddOn, 0, 10)
	// Roundtripper middleware in client/pagination.go
	// sets the nextRange value after each request
	for nextRange.Max != 0 {
		ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
		v, err := c.Heroku.AddOnList(ctxWithRange, nextRange)
		if err != nil {
			return errors.WithStack(err)
		}
		items = append(items, v...)
	}

	for _, it := range items {
		nextRange = &heroku.ListRange{
			Field: "id",
			Max:   1000,
		}
		// Roundtripper middleware in client/pagination.go
		// sets the nextRange value after each request
		for nextRange.Max != 0 {
			ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
			v, err := c.Heroku.AddOnWebhookDeliveryList(ctxWithRange, it.ID, nextRange)
			if err != nil {
				return errors.WithStack(err)
			}
			res <- v
		}
	}
	return nil
}
