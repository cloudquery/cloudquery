package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func VpcEndpointServices() *schema.Table {
	tableName := "aws_ec2_vpc_endpoint_services"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ServiceDetail.html`,
		Resolver:    fetchEc2VpcEndpointServices,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Transform:   transformers.TransformWithStruct(&types.ServiceDetail{}, transformers.WithPrimaryKeyComponents("ServiceId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
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
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEc2).Ec2
	// No paginator available
	for {
		output, err := svc.DescribeVpcEndpointServices(ctx, &config, func(options *ec2.Options) {
			options.Region = cl.Region
		})
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
