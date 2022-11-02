package container

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
)

func fetchClusters(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	parent := fmt.Sprintf("projects/%s/locations/-", c.ProjectId)
	output, err := c.Services.Container.Projects.Locations.Clusters.List(parent).Do()
	if err != nil {
		return errors.WithStack(err)
	}
	res <- output.Clusters
	return nil
}
