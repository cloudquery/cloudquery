package users

import (
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/xanzy/go-gitlab"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:      "gitlab_users",
		Resolver:  fetchUsers,
		Transform: client.TransformWithStruct(&gitlab.User{}, transformers.WithPrimaryKeys("ID")),
		Columns:   schema.ColumnList{client.BaseURLColumn},
	}
}
