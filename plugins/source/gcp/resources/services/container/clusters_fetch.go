package container

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	pb "google.golang.org/genproto/googleapis/container/v1"
)

func fetchClusters(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.ListClustersRequest{
		Parent: fmt.Sprintf("projects/%s/locations/-", c.ProjectId),
	}
	output, err := c.Services.ContainerClusterManagerClient.ListClusters(ctx, req)
	if err != nil {
		return errors.WithStack(err)
	}
	res <- output.Clusters
	return nil
}
