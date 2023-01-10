package redshift

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchRedshiftSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Redshift
	cluster := parent.Item.(types.Cluster)
	params := redshift.DescribeClusterSnapshotsInput{
		ClusterExists:     aws.Bool(true),
		ClusterIdentifier: cluster.ClusterIdentifier,
		MaxRecords:        aws.Int32(100),
	}
	for {
		result, err := svc.DescribeClusterSnapshots(ctx, &params)
		if err != nil {
			return err
		}
		res <- result.Snapshots
		if aws.ToString(result.Marker) == "" {
			break
		}
		params.Marker = result.Marker
	}
	return nil
}

func resolveSnapshotARN(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	snapshot := resource.Item.(types.Snapshot)
	return resource.Set(c.Name, snapshotARN(cl, *snapshot.ClusterIdentifier, *snapshot.SnapshotIdentifier))
}

func snapshotARN(cl *client.Client, clusterName, snapshotName string) string {
	return arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.RedshiftService),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("snapshot:%s/%s", clusterName, snapshotName),
	}.String()
}
