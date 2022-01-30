package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Ec2ByoipCidrs() *schema.Table {
	return &schema.Table{
		Name:          "aws_ec2_byoip_cidrs",
		Description:   "Information about an address range that is provisioned for use with your AWS resources through bring your own IP addresses (BYOIP).",
		Resolver:      fetchEc2ByoipCidrs,
		Multiplex:     client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreError:   client.IgnoreWithInvalidAction,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "region", "cidr"}},
		IgnoreInTests: true,
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
				Name:        "cidr",
				Description: "The address range, in CIDR notation.",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "The description of the address range.",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "The state of the address pool.",
				Type:        schema.TypeString,
			},
			{
				Name:        "status_message",
				Description: "Upon success, contains the ID of the address pool.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEc2ByoipCidrs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	config := ec2.DescribeByoipCidrsInput{
		MaxResults: aws.Int32(100),
	}

	c := meta.(*client.Client)
	// DescribeByoipCidrs does not work in next regions, so we ignore them.
	if _, ok := map[string]struct{}{
		"ap-northeast-2": {},
		"af-south-1":     {},
		"cn-north-1":     {},
		"cn-northwest-1": {},
		"eu-west-3":      {},
		"eu-south-1":     {},
		"eu-north-1":     {},
		"me-south-1":     {},
	}[c.Region]; ok {
		return nil
	}
	svc := c.Services().EC2
	for {
		response, err := svc.DescribeByoipCidrs(ctx, &config, func(options *ec2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.ByoipCidrs
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
