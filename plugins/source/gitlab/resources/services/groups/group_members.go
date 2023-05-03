package groups

import (
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/xanzy/go-gitlab"
)

func GroupMembers() *schema.Table {
	return &schema.Table{
		Name:      "gitlab_group_members",
		Resolver:  fetchGroupMembers,
		Transform: client.TransformWithStruct(&gitlab.GroupMember{}, transformers.WithPrimaryKeys("ID")),
		Columns: schema.ColumnList{client.BaseURLColumn,
			{
				Name:            "group_id",
				Type:            schema.TypeInt,
				Resolver:        resolveGroupID,
				CreationOptions: schema.ColumnCreationOptions{NotNull: true, PrimaryKey: true},
			},
		},
	}
}
