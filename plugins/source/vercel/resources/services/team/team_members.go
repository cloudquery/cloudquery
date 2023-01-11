package team

import (
	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func TeamMembers() *schema.Table {
	return &schema.Table{
		Name:      "vercel_team_members",
		Resolver:  fetchTeamMembers,
		Transform: transformers.TransformWithStruct(&vercel.TeamMember{}, client.SharedTransformers()...),
		Columns: []schema.Column{
			{
				Name:     "team_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "uid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
