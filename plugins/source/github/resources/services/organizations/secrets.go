package organizations

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func secrets() *schema.Table {
	return &schema.Table{
		Name:     "github_organization_dependabot_secrets",
		Resolver: fetchSecrets,
		Transform: transformers.TransformWithStruct(&github.Secret{},
			append(client.SharedTransformers(), transformers.WithPrimaryKeys("Name"))...),
		Columns: []schema.Column{client.OrgColumn},
	}
}

func fetchSecrets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	secrets, _, err := c.Github.Dependabot.ListOrgSecrets(ctx, c.Org, nil)
	if err != nil {
		return err
	}

	res <- secrets.Secrets

	return nil
}
