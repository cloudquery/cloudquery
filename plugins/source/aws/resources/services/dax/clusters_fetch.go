package dax

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/aws/aws-sdk-go-v2/service/dax/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDaxClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Dax

	config := dax.DescribeClustersInput{}
	for {
		output, err := svc.DescribeClusters(ctx, &config)
		if err != nil {
			return err
		}

		res <- output.Clusters

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}

	return nil
}
func resolveClusterTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cluster := resource.Item.(types.Cluster)

	cl := meta.(*client.Client)
	svc := cl.Services().Dax
	response, err := svc.ListTags(ctx, &dax.ListTagsInput{
		ResourceName: cluster.ClusterArn,
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(response.Tags))
}
