package emr

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func clusterInstances() *schema.Table {
	tableName := "aws_emr_cluster_instances"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/emr/latest/APIReference/API_Instance.html`,
		Resolver:    fetchClusterInstances,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticmapreduce"),
		Transform:   transformers.TransformWithStruct(&types.Instance{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveClusterInstanceArn,
				PrimaryKey: true,
			},
		},
	}
}

func fetchClusterInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*types.Cluster)
	svc := cl.Services().Emr
	paginator := emr.NewListInstancesPaginator(svc, &emr.ListInstancesInput{ClusterId: p.Id})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *emr.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Instances
	}
	return nil
}

func resolveClusterInstanceArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.Instance)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "ec2",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "instance/" + aws.ToString(item.Ec2InstanceId),
	}
	return resource.Set(c.Name, a.String())
}
