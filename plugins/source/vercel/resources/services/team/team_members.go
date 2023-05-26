package team

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func TeamMembers() *schema.Table {
	return &schema.Table{
		Name:          "vercel_team_members",
		Resolver:      fetchTeamMembers,
		Transform:     client.TransformWithStruct(&vercel.TeamMember{}),
		IsIncremental: true,
		Columns: []schema.Column{
			{
				Name:       "team_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("id"),
				PrimaryKey: true,
			},
			{
				Name:       "uid",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("UID"),
				PrimaryKey: true,
			},
		},
	}
}
