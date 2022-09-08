package xray

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)


func EncryptionConfigs() *schema.Table {
	return &schema.Table{
		Name:         "aws_xray_encryption_config",
		Description:  "A configuration document that specifies encryption configuration settings",
		Resolver:     fetchXrayEncryptionConfigs,
		Multiplex:    client.ServiceAccountRegionMultiplexer("xray"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "key_id",
				Description: "The ID of the KMS key used for encryption, if applicable",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The encryption status",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The type of encryption",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchXrayEncryptionConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
