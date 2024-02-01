package xray

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func EncryptionConfigs() *schema.Table {
	tableName := "aws_xray_encryption_configs"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/xray/latest/api/API_EncryptionConfig.html`,
		Resolver:    fetchXrayEncryptionConfigs,
		Transform:   transformers.TransformWithStruct(&types.EncryptionConfig{}, transformers.WithPrimaryKeyComponents("KeyId", "Type")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "xray"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}

func fetchXrayEncryptionConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceXray).Xray
	input := xray.GetEncryptionConfigInput{}
	output, err := svc.GetEncryptionConfig(ctx, &input, func(o *xray.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	if output.EncryptionConfig.KeyId == nil {
		// Can't have nil for PK: when Type is `NONE`, KeyId is nil
		output.EncryptionConfig.KeyId = aws.String("")
	}
	res <- output.EncryptionConfig
	return nil
}
