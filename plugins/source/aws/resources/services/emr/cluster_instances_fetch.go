package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchClusterInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	p := parent.Item.(*types.Cluster)
	config := emr.ListInstancesInput{ClusterId: p.Id}
	svc := c.Services().Emr
	for {
		response, err := svc.ListInstances(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.Instances

		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
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
