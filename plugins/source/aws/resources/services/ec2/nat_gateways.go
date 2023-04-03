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

func NatGateways() *schema.Table {
	tableName := "aws_ec2_nat_gateways"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_NatGateway.html`,
		Resolver:    fetchEc2NatGateways,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Transform:   transformers.TransformWithStruct(&types.NatGateway{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveNatGatewayArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchEc2NatGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config ec2.DescribeNatGatewaysInput
	c := meta.(*client.Client)
	svc := c.Services().Ec2
	for {
		output, err := svc.DescribeNatGateways(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.NatGateways
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func resolveNatGatewayArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.NatGateway)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "ec2",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "natgateway/" + aws.ToString(item.NatGatewayId),
	}
	return resource.Set(c.Name, a.String())
}
