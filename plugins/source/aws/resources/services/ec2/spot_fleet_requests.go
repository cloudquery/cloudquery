package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SpotFleetRequests() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_spot_fleet_requests",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SpotFleetRequestConfig.html`,
		Resolver:    fetchEC2SpotFleetRequests,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
		Transform:   transformers.TransformWithStruct(&types.SpotFleetRequestConfig{}, transformers.WithPrimaryKeys("SpotFleetRequestId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
		Relations: []*schema.Table{
			spotFleetInstances(),
		},
	}
}

func fetchEC2SpotFleetRequests(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config ec2.DescribeSpotFleetRequestsInput
	c := meta.(*client.Client)
	svc := c.Services().Ec2
	for {
		output, err := svc.DescribeSpotFleetRequests(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.SpotFleetRequestConfigs
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
