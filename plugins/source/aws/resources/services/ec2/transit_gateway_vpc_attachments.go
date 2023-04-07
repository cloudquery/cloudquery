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

func transitGatewayVpcAttachments() *schema.Table {
	tableName := "aws_ec2_transit_gateway_vpc_attachments"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayVpcAttachment.html`,
		Resolver:    fetchEc2TransitGatewayVpcAttachments,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Transform:   transformers.TransformWithStruct(&types.TransitGatewayVpcAttachment{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "transit_gateway_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchEc2TransitGatewayVpcAttachments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.TransitGateway)

	config := ec2.DescribeTransitGatewayVpcAttachmentsInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("transit-gateway-id"),
				Values: []string{*r.TransitGatewayId},
			},
		},
	}
	c := meta.(*client.Client)
	svc := c.Services().Ec2
	paginator := ec2.NewDescribeTransitGatewayVpcAttachmentsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.TransitGatewayVpcAttachments
	}
	return nil
}
