package directconnect

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func gatewayAttachments() *schema.Table {
	tableName := "aws_directconnect_gateway_attachments"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DirectConnectGatewayAttachment.html`,
		Resolver:    fetchDirectconnectGatewayAttachments,
		Transform: transformers.TransformWithStruct(&types.DirectConnectGatewayAttachment{},
			transformers.WithPrimaryKeyComponents("VirtualInterfaceOwnerAccount", "VirtualInterfaceRegion", "VirtualInterfaceId"),
		),
		Columns: []schema.Column{
			client.RequestAccountIDColumn(true),
			client.RequestRegionColumn(false),
			{
				Name:                "gateway_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "gateway_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}

func fetchDirectconnectGatewayAttachments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	gateway := parent.Item.(types.DirectConnectGateway)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceDirectconnect).Directconnect
	config := directconnect.DescribeDirectConnectGatewayAttachmentsInput{DirectConnectGatewayId: gateway.DirectConnectGatewayId}
	// No paginator available
	for {
		output, err := svc.DescribeDirectConnectGatewayAttachments(ctx, &config, func(options *directconnect.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.DirectConnectGatewayAttachments
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
