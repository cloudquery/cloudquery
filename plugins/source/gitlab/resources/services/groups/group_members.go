package groups

import (
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/xanzy/go-gitlab"
)

func GroupMembers() *schema.Table {
	return &schema.Table{
		Name:      "gitlab_group_members",
		Resolver:  fetchGroupMembers,
		Transform: transformers.TransformWithStruct(&gitlab.GroupMember{}, client.SharedTransformers()...),
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
				Name:     "group_id",
				Type:     schema.TypeInt,
				Resolver: resolveGroupID,
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
				Name:     "expires_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ExpiresAt"),
			},
		},
	}
}
