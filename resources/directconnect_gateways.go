package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func DirectconnectGateways() *schema.Table {
	return &schema.Table{
		Name:         "aws_directconnect_gateways",
		Resolver:     fetchDirectconnectGateways,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name: "amazon_side_asn",
				Type: schema.TypeBigInt,
			},
			{
				Name: "direct_connect_gateway_id",
				Type: schema.TypeString,
			},
			{
				Name: "direct_connect_gateway_name",
				Type: schema.TypeString,
			},
			{
				Name: "direct_connect_gateway_state",
				Type: schema.TypeString,
			},
			{
				Name: "owner_account",
				Type: schema.TypeString,
			},
			{
				Name: "state_change_error",
				Type: schema.TypeString,
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
