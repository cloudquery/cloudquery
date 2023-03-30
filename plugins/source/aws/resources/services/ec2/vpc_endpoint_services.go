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

func VpcEndpointServices() *schema.Table {
	tableName := "aws_ec2_vpc_endpoint_services"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ServiceDetail.html`,
		Resolver:    fetchEc2VpcEndpointServices,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Transform:   transformers.TransformWithStruct(&types.ServiceDetail{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveVpcEndpointServiceArn,
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
		Relations: []*schema.Table{
			vpcEndpointServicePermissions(),
		},
	}
}
func fetchEc2VpcEndpointServices(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	var config ec2.DescribeVpcEndpointServicesInput
	c := meta.(*client.Client)
	svc := c.Services().Ec2
	for {
		output, err := svc.DescribeVpcEndpointServices(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.ServiceDetails
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func resolveVpcEndpointServiceArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.ServiceDetail)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "ec2",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "vpc-endpoint-service/" + aws.ToString(item.ServiceId),
	}
	return resource.Set(c.Name, a.String())
}
