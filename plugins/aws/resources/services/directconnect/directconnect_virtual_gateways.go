package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func DirectconnectVirtualGateways() *schema.Table {
	return &schema.Table{
		Name:         "aws_directconnect_virtual_gateways",
		Description:  "Information about a virtual private gateway for a private virtual interface.",
		Resolver:     fetchDirectconnectVirtualGateways,
		Multiplex:    client.ServiceAccountRegionMultiplexer("directconnect"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "id",
				Description: "The ID of the virtual private gateway.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualGatewayId"),
			},
			{
				Name:        "state",
				Description: "The state of the virtual private gateway.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualGatewayState"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchDirectconnectVirtualGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config directconnect.DescribeVirtualGatewaysInput
	c := meta.(*client.Client)
	svc := c.Services().Directconnect
	output, err := svc.DescribeVirtualGateways(ctx, &config, func(options *directconnect.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}
	res <- output.VirtualGateways
	return nil
}
