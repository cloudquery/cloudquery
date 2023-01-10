package groups

import (
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/xanzy/go-gitlab"
)

func Groups() *schema.Table {
	return &schema.Table{
		Name:      "gitlab_groups",
		Resolver:  fetchGroups,
		Transform: transformers.TransformWithStruct(&gitlab.Group{}, client.SharedTransformers()...),
		Columns: []schema.Column{
			{
				Name:     "base_url",
				Type:     schema.TypeString,
				Resolver: client.ResolveURL,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "marked_for_deletion_on",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("MarkedForDeletionOn"),
			},
		},

		Relations: []*schema.Table{
			GroupMembers(),
		},
	}
}
