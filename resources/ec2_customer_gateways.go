package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Ec2CustomerGateways() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_customer_gateways",
		Resolver:     fetchEc2CustomerGateways,
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
				Name: "bgp_asn",
				Type: schema.TypeString,
			},
			{
				Name: "certificate_arn",
				Type: schema.TypeString,
			},
			{
				Name: "customer_gateway_id",
				Type: schema.TypeString,
			},
			{
				Name: "device_name",
				Type: schema.TypeString,
			},
			{
				Name: "ip_address",
				Type: schema.TypeString,
			},
			{
				Name: "state",
				Type: schema.TypeString,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveEc2customerGatewayTags,
			},
			{
				Name: "type",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEc2CustomerGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().EC2
	response, err := svc.DescribeCustomerGateways(ctx, nil, func(options *ec2.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}
	res <- response.CustomerGateways
	return nil

}
func resolveEc2customerGatewayTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.CustomerGateway)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
