package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ec2/models"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func RegionalConfigs() *schema.Table {
	tableName := "aws_ec2_regional_configs"
	return &schema.Table{
		Name:      tableName,
		Resolver:  fetchEc2RegionalConfigs,
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Transform: transformers.TransformWithStruct(&models.RegionalConfig{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}

func fetchEc2RegionalConfigs(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	svc := c.Services().Ec2
	var regionalConfig models.RegionalConfig
	resp, err := svc.GetEbsDefaultKmsKeyId(ctx, &ec2.GetEbsDefaultKmsKeyIdInput{}, func(options *ec2.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}
	regionalConfig.EbsDefaultKmsKeyId = resp.KmsKeyId

	ebsResp, err := svc.GetEbsEncryptionByDefault(ctx, &ec2.GetEbsEncryptionByDefaultInput{}, func(options *ec2.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}

	if ebsResp.EbsEncryptionByDefault != nil {
		regionalConfig.EbsEncryptionEnabledByDefault = *ebsResp.EbsEncryptionByDefault
	}
	res <- regionalConfig
	return nil
}
