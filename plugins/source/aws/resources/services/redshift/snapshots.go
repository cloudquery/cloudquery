package redshift

import (
	"context"
	"fmt"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func snapshots() *schema.Table {
	tableName := "aws_redshift_snapshots"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/redshift/latest/APIReference/API_Snapshot.html`,
		Resolver:    fetchSnapshots,
		Transform:   transformers.TransformWithStruct(&types.Snapshot{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "redshift"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:        "arn",
				Type:        arrow.BinaryTypes.String,
				Resolver:    resolveSnapshotARN,
				Description: `ARN of the snapshot.`,
				PrimaryKey:  true,
			},
			{
				Name:        "tags",
				Type:        sdkTypes.ExtensionTypes.JSON,
				Resolver:    client.ResolveTags,
				Description: `Tags consisting of a name/value pair for a resource.`,
			},
		},
	}
}

func fetchSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Redshift
	cluster := parent.Item.(types.Cluster)
	params := redshift.DescribeClusterSnapshotsInput{
		ClusterExists:     aws.Bool(true),
		ClusterIdentifier: cluster.ClusterIdentifier,
		MaxRecords:        aws.Int32(100),
	}
	paginator := redshift.NewDescribeClusterSnapshotsPaginator(svc, &params)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *redshift.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Snapshots
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
