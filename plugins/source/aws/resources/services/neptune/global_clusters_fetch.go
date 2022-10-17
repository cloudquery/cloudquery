package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchNeptuneGlobalClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config neptune.DescribeGlobalClustersInput
	c := meta.(*client.Client)
	svc := c.Services().Neptune
	for {
		response, err := svc.DescribeGlobalClusters(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.GlobalClusters
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}

func resolveNeptuneGlobalClusterTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	s := resource.Item.(types.GlobalCluster)
	cl := meta.(*client.Client)
	svc := cl.Services().Neptune
	out, err := svc.ListTagsForResource(ctx, &neptune.ListTagsForResourceInput{ResourceName: s.GlobalClusterArn})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(out.TagList))
}
