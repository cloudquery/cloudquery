package container

import (
	"context"
	"fmt"

	"cloud.google.com/go/container/apiv1/containerpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
)

func fetchClusters(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &containerpb.ListClustersRequest{
		Parent: fmt.Sprintf("projects/%s/locations/-", c.ProjectId),
	}
	output, err := c.Services.ContainerClusterManagerClient.ListClusters(ctx, req)
	if err != nil {
		return errors.WithStack(err)
	}
	res <- output.Clusters
	return nil
}
