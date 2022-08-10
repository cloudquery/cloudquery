package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource key_pairs --config gen.hcl --output .
func Ec2KeyPairs() *schema.Table {
	return &schema.Table{
		Name:          "aws_ec2_key_pairs",
		Description:   "Describes an EC2 Key Pair.",
		Resolver:      fetchEc2KeyPairs,
		Multiplex:     client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"key_pair_id"}},
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
				Name:        "create_time",
				Description: "The date and time when the key was created in ISO 8601 date-time format.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "key_fingerprint",
				Description: "The fingerprint of the private key digest.",
				Type:        schema.TypeString,
			},
			{
				Name:        "key_name",
				Description: "The name of the key pair.",
				Type:        schema.TypeString,
			},
			{
				Name:        "key_pair_id",
				Description: "The ID of the key pair.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Any tags assigned to the key pair.",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEc2KeyPairs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeKeyPairsInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	output, err := svc.DescribeKeyPairs(ctx, &config)
	if err != nil {
		return diag.WrapError(err)
	}
	res <- output.KeyPairs
	return nil
}
