package dependency

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDependencies(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	result, _, err := c.Dependencies.List(ctx, c.OrganizationID)
	if err != nil {
		return err
	}
	for _, dep := range result {
		res <- dep
		fmt.Println(dep)
	}

	res <- result

	return nil
}
