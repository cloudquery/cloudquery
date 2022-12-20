package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchRedshiftClusterParameterGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cluster := parent.Item.(types.Cluster)
	res <- cluster.ClusterParameterGroups
	return nil
}
