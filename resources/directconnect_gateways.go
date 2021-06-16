package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

// todo implement tags
func DirectconnectGateways() *schema.Table {
	return &schema.Table{
		Name:         "aws_directconnect_gateways",
		Description:  "Information about a Direct Connect gateway, which enables you to connect virtual interfaces and virtual private gateway or transit gateways.",
		Resolver:     fetchDirectconnectGateways,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
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
				Name:        "amazon_side_asn",
				Description: "The autonomous system number (ASN) for the Amazon side of the connection.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "direct_connect_gateway_id",
				Description: "The ID of the Direct Connect gateway.",
				Type:        schema.TypeString,
			},
			{
				Name:        "direct_connect_gateway_name",
				Description: "The name of the Direct Connect gateway.",
				Type:        schema.TypeString,
			},
			{
				Name:        "direct_connect_gateway_state",
				Description: "The state of the Direct Connect gateway.",
				Type:        schema.TypeString,
			},
			{
				Name:        "owner_account",
				Description: "The ID of the AWS account that owns the Direct Connect gateway.",
				Type:        schema.TypeString,
			},
			{
				Name:        "state_change_error",
				Description: "The error message if the state of an object failed to advance.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchDirectconnectGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config directconnect.DescribeDirectConnectGatewaysInput
	c := meta.(*client.Client)
	svc := c.Services().Directconnect
	for {
		output, err := svc.DescribeDirectConnectGateways(ctx, &config, func(options *directconnect.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- output.DirectConnectGateways
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
