package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Clusters() *schema.Table {
	tableName := "aws_eks_clusters"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/eks/latest/APIReference/API_Cluster.html`,
		Resolver:            fetchEksClusters,
		PreResourceResolver: getEksCluster,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "eks"),
		Transform:           transformers.TransformWithStruct(&types.Cluster{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{
			nodeGroups(),
			fargateProfiles(),
		},
	}
}

func fetchEksClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config eks.ListClustersInput
	c := meta.(*client.Client)
	svc := c.Services().Eks
	// TODO: replace with paginator
	for {
		listClustersOutput, err := svc.ListClusters(ctx, &config)
		if err != nil {
			return err
		}
		res <- listClustersOutput.Clusters
		if listClustersOutput.NextToken == nil {
			break
		}
		config.NextToken = listClustersOutput.NextToken
	}
	return nil
}

func getEksCluster(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Eks
	name := resource.Item.(string)
	output, err := svc.DescribeCluster(
		ctx, &eks.DescribeClusterInput{Name: &name}, func(options *eks.Options) {
			options.Region = c.Region
		})
	if err != nil {
		return err
	}
	resource.Item = output.Cluster
	return nil
}
