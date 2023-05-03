package groups

import (
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/xanzy/go-gitlab"
)

func GroupMembers() *schema.Table {
	return &schema.Table{
		Name:      "gitlab_group_members",
		Resolver:  fetchGroupMembers,
		Transform: client.TransformWithStruct(&gitlab.GroupMember{}),
		Columns: []schema.Column{
			client.BaseURLColumn,
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
