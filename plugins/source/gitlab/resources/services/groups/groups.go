package groups

import (
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/xanzy/go-gitlab"
)

func Groups() *schema.Table {
	return &schema.Table{
		Name:      "gitlab_groups",
		Resolver:  fetchGroups,
		Transform: client.TransformWithStruct(&gitlab.Group{}, transformers.WithPrimaryKeys("ID", "Name")),
		Columns:   schema.ColumnList{client.BaseURLColumn},
		Relations: schema.Tables{GroupMembers()},
	}
}
