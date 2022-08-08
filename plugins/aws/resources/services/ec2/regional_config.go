package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

type ec2RegionalConfig struct {
	EbsEncryptionEnabledByDefault bool
	EbsDefaultKmsKeyId            *string
}

func Ec2RegionalConfig() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_regional_config",
		Description:  "Ec2 Regional Config defines common default configuration for ec2 service",
		Resolver:     fetchEc2RegionalConfig,
		Multiplex:    client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "region"}},
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
				Name:        "ebs_encryption_enabled_by_default",
				Type:        schema.TypeBool,
				Description: "Indicates whether EBS encryption by default is enabled for your account in the current Region.",
			},
			{
				Name:        "ebs_default_kms_key_id",
				Type:        schema.TypeString,
				Description: "The Amazon Resource Name (ARN) of the default CMK for encryption by default.",
			},
		},
	}
}

func fetchEc2RegionalConfig(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	svc := c.Services().EC2
	var regionalConfig ec2RegionalConfig
	resp, err := svc.GetEbsDefaultKmsKeyId(ctx, &ec2.GetEbsDefaultKmsKeyIdInput{}, func(options *ec2.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	regionalConfig.EbsDefaultKmsKeyId = resp.KmsKeyId

	ebsResp, err := svc.GetEbsEncryptionByDefault(ctx, &ec2.GetEbsEncryptionByDefaultInput{}, func(options *ec2.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}

	if ebsResp.EbsEncryptionByDefault != nil {
		regionalConfig.EbsEncryptionEnabledByDefault = *ebsResp.EbsEncryptionByDefault
	}
	res <- regionalConfig
	return nil
}
