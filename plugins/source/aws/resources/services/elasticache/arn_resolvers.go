package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func resolveCacheNodesOfferingArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.ReservedCacheNodesOffering)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "elasticache",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "elasticache/" + aws.ToString(item.ReservedCacheNodesOfferingId),
	}
	return resource.Set(c.Name, a.String())
}

func resolveServiceUpdateArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.ServiceUpdate)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "elasticache",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "elasticache/" + aws.ToString(item.ServiceUpdateName),
	}
	return resource.Set(c.Name, a.String())
}
