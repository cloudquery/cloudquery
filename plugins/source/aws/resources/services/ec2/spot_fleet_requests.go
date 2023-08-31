package ec2

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func SpotFleetRequests() *schema.Table {
	tableName := "aws_ec2_spot_fleet_requests"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SpotFleetRequestConfig.html`,
		Resolver:    fetchEC2SpotFleetRequests,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Transform:   transformers.TransformWithStruct(&types.SpotFleetRequestConfig{}, transformers.WithPrimaryKeys("SpotFleetRequestId")),
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
			spotFleetInstances(),
		},
	}
}

func fetchEC2SpotFleetRequests(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEc2).Ec2
	pag := ec2.NewDescribeSpotFleetRequestsPaginator(svc, &ec2.DescribeSpotFleetRequestsInput{})
	for pag.HasMorePages() {
		resp, err := pag.NextPage(ctx, func(options *ec2.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- resp.SpotFleetRequestConfigs
	}
	return nil
}
