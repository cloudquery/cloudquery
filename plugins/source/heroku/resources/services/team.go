// Code generated by codegen; DO NOT EDIT.

package services

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/pkg/errors"
)

func Teams() *schema.Table {
	return &schema.Table{
		Name:      "heroku_teams",
		Resolver:  fetchTeams,
		Multiplex: client.NoMultiplex,
		Columns: []schema.Column{
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "credit_card_collections",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("CreditCardCollections"),
			},
			{
				Name:     "default",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Default"),
			},
			{
				Name:     "enterprise_account",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EnterpriseAccount"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "identity_provider",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IdentityProvider"),
			},
			{
				Name:     "membership_limit",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("MembershipLimit"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "provisioned_licenses",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ProvisionedLicenses"),
			},
			{
				Name:     "role",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Role"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
		},
	}
}

func fetchTeams(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	v, err := c.Heroku.TeamList(ctx, nil)
	if err != nil {
		return errors.WithStack(err)
	}
	res <- v
	return nil
}
