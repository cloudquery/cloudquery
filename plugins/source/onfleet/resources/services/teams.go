package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/onfleet/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/keplr-team/go-onfleet/onfleet"
)

func Teams() *schema.Table {
	return &schema.Table{
		Name:        "onfleet_teams",
		Resolver:    fetchTeams,
		Description: "https://docs.onfleet.com/reference/list-teams",
		Transform: transformers.TransformWithStruct(&onfleet.Team{},
			client.OnfleetSharedTransformers()...),
		Columns: []schema.Column{
			{
				Name:     "organization_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveOrganizationId,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
