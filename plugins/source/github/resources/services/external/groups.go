package external

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func Groups() *schema.Table {
	return &schema.Table{
		Name:      "github_external_groups",
		Resolver:  fetchGroups,
		Multiplex: client.OrgMultiplex,
		Transform: transformers.TransformWithStruct(&github.ExternalGroup{}),
		Columns: []schema.Column{
			{
				Name:        "org",
				Type:        schema.TypeString,
				Resolver:    client.ResolveOrg,
				Description: `The Github Organization of the resource.`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "group_id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("GroupID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GroupName"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "teams",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Teams"),
			},
			{
				Name:     "members",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Members"),
			},
		},
	}
}
