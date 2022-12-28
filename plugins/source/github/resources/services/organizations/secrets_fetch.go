package organizations

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchSecrets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	secrets, _, err := c.Github.Dependabot.ListOrgSecrets(ctx, c.Org, nil)
	if err != nil {
		return err
	}

	res <- secrets.Secrets

	return nil
}
