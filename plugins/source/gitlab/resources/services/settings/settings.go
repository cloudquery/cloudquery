package settings

import (
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/xanzy/go-gitlab"
)

func Settings() *schema.Table {
	return &schema.Table{
		Name:      "gitlab_settings",
		Resolver:  fetchSettings,
		Transform: client.TransformWithStruct(&gitlab.Settings{}),
		Columns: []schema.Column{
			client.BaseURLColumn,
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
