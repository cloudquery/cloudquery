package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SpotInstanceRequests() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_spot_instance_requests",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SpotInstanceRequest.html`,
		Resolver:    fetchEC2SpotInstanceRequests,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
		Transform:   transformers.TransformWithStruct(&types.SpotInstanceRequest{}, transformers.WithPrimaryKeys("SpotInstanceRequestId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchEC2SpotInstanceRequests(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Ec2
	pag := ec2.NewDescribeSpotInstanceRequestsPaginator(svc, &ec2.DescribeSpotInstanceRequestsInput{})
	for pag.HasMorePages() {
		resp, err := pag.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- resp.SpotInstanceRequests
	}
	return nil
}
