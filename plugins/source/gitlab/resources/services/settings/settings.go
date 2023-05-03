package settings

import (
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/xanzy/go-gitlab"
)

func Settings() *schema.Table {
	return &schema.Table{
		Name:      "gitlab_settings",
		Resolver:  fetchSettings,
		Transform: client.TransformWithStruct(&gitlab.Settings{}, transformers.WithPrimaryKeys("ID")),
		Columns:   schema.ColumnList{client.BaseURLColumn},
	}
}
