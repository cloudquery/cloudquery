package settings

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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

func fetchSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	if c.BaseURL == "" {
		c.Logger().Info().Str("table", "gitlab_settings").Msg("not supported for GitLab SaaS, skipping...")
		return nil
	}
	setting, _, err := c.Gitlab.Settings.GetSettings(gitlab.WithContext(ctx))
	if err != nil {
		return err
	}

	res <- setting

	return nil
}
