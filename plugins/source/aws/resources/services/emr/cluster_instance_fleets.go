package emr

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func clusterInstanceFleets() *schema.Table {
	tableName := "aws_emr_cluster_instance_fleets"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceFleet.html`,
		Resolver:    fetchClusterInstanceFleets,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticmapreduce"),
		Transform:   transformers.TransformWithStruct(&types.InstanceFleet{}, transformers.WithPrimaryKeys("Id")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "cluster_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchClusterInstanceFleets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cluster := parent.Item.(*types.Cluster)
	// instance fleets and instance groups are mutually exclusive
	if cluster.InstanceCollectionType != types.InstanceCollectionTypeInstanceFleet {
		return nil
	}
	config := emr.ListInstanceFleetsInput{
		ClusterId: cluster.Id,
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Emr
	paginator := emr.NewListInstanceFleetsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *emr.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.InstanceFleets
	}
	return nil
}
