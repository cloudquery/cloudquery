package xray

import (
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
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
