package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchEksClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config eks.ListClustersInput
	c := meta.(*client.Client)
	svc := c.Services().Eks
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
