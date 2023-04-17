package xray

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func EncryptionConfigs() *schema.Table {
	tableName := "aws_xray_encryption_configs"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/xray/latest/api/API_EncryptionConfig.html`,
		Resolver:    fetchXrayEncryptionConfigs,
		Transform:   transformers.TransformWithStruct(&types.EncryptionConfig{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "xray"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}

func fetchXrayEncryptionConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Xray
	input := xray.GetEncryptionConfigInput{}
	output, err := svc.GetEncryptionConfig(ctx, &input)
	if err != nil {
		return err
	}
	res <- output.EncryptionConfig
	return nil
}
