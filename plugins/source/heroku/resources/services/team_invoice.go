// Code generated by codegen; DO NOT EDIT.

package services

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/pkg/errors"
)

func TeamInvoices() *schema.Table {
	return &schema.Table{
		Name:      "heroku_team_invoices",
		Resolver:  fetchTeamInvoices,
		Multiplex: client.NoMultiplex,
		Columns: []schema.Column{
			{
				Name:     "addons_total",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("AddonsTotal"),
			},
			{
				Name:     "charges_total",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ChargesTotal"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "credits_total",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CreditsTotal"),
			},
			{
				Name:     "database_total",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DatabaseTotal"),
			},
			{
				Name:     "dyno_units",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("DynoUnits"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "number",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Number"),
			},
			{
				Name:     "payment_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PaymentStatus"),
			},
			{
				Name:     "period_end",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PeriodEnd"),
			},
			{
				Name:     "period_start",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PeriodStart"),
			},
			{
				Name:     "platform_total",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PlatformTotal"),
			},
			{
				Name:     "state",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "total",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Total"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "weighted_dyno_hours",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("WeightedDynoHours"),
			},
		},
	}
}

func fetchTeamInvoices(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextRange := &heroku.ListRange{
		Field: "id",
		Max:   1000,
	}
	items := make([]heroku.Team, 0, 10)
	// Roundtripper middleware in client/pagination.go
	// sets the nextRange value after each request
	for nextRange.Max != 0 {
		ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
		v, err := c.Heroku.TeamList(ctxWithRange, nextRange)
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
			v, err := c.Heroku.TeamInvoiceList(ctxWithRange, it.ID, nextRange)
			if err != nil {
				return errors.WithStack(err)
			}
			res <- v
		}
	}
	return nil
}
