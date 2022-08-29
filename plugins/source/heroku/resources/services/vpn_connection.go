// Code generated by codegen; DO NOT EDIT.

package services

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/pkg/errors"
)

func VPNConnections() *schema.Table {
	return &schema.Table{
		Name:        "heroku_vpn_connections",
		Description: "https://devcenter.heroku.com/articles/platform-api-reference#vpn-connection-attributes",
		Resolver:    fetchVPNConnections,
		Multiplex:   client.NoMultiplex,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "ike_version",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("IKEVersion"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "public_ip",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PublicIP"),
			},
			{
				Name:     "routable_cidrs",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("RoutableCidrs"),
			},
			{
				Name:     "space_cidr_block",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SpaceCIDRBlock"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "status_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StatusMessage"),
			},
			{
				Name:     "tunnels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tunnels"),
			},
		},
	}
}

func fetchVPNConnections(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextRange := &heroku.ListRange{
		Field: "id",
		Max:   1000,
	}
	items := make([]heroku.Space, 0, 10)
	// Roundtripper middleware in client/pagination.go
	// sets the nextRange value after each request
	for nextRange.Max != 0 {
		ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
		v, err := c.Heroku.SpaceList(ctxWithRange, nextRange)
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
			v, err := c.Heroku.VPNConnectionList(ctxWithRange, it.ID, nextRange)
			if err != nil {
				return errors.WithStack(err)
			}
			res <- v
		}
	}
	return nil
}
