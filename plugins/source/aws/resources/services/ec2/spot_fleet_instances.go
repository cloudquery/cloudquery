package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func spotFleetInstances() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_spot_fleet_instances",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ActiveInstance.html`,
		Resolver:    fetchEC2SpotFleetInstances,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
		Transform:   transformers.TransformWithStruct(&types.ActiveInstance{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveActiveInstanceArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "spot_fleet_request_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("spot_fleet_request_id"),
			},
		},
	}
}

func fetchEC2SpotFleetInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(types.SpotFleetRequestConfig)

	config := ec2.DescribeSpotFleetInstancesInput{
		SpotFleetRequestId: p.SpotFleetRequestId,
	}
	c := meta.(*client.Client)
	svc := c.Services().Ec2
	for {
		output, err := svc.DescribeSpotFleetInstances(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.ActiveInstances
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func resolveActiveInstanceArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.ActiveInstance)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "ec2",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "instance/" + aws.ToString(item.InstanceId),
	}
	return resource.Set(c.Name, a.String())
}
