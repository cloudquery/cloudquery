package project

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchProjects(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	projects, _, err := c.Projects.List(ctx, c.OrganizationID)
	if err != nil {
		return err
	}

	res <- projects

	return nil
}
