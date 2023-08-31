package container

import (
	"context"
	"fmt"

	container "cloud.google.com/go/container/apiv1"
	"cloud.google.com/go/container/apiv1/containerpb"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func fetchClusters(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &containerpb.ListClustersRequest{
		Parent: fmt.Sprintf("projects/%s/locations/-", c.ProjectId),
	}
	containerClient, err := container.NewClusterManagerClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	output, err := containerClient.ListClusters(ctx, req)
	if err != nil {
		return err
	}
	res <- output.Clusters
	return nil
}
