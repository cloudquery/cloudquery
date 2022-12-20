package cloudhsmv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchCloudhsmv2Clusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudhsmv2
	var input cloudhsmv2.DescribeClustersInput
	paginator := cloudhsmv2.NewDescribeClustersPaginator(svc, &input)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.Clusters
	}
	return nil
}

func resolveClusterArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.Cluster)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "hsm",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "cluster/" + aws.ToString(item.ClusterId),
	}
	return resource.Set(c.Name, a.String())
}
