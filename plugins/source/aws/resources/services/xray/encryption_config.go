package xray

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource encryption_config --config gen.hcl --output .
func EncryptionConfigs() *schema.Table {
	return &schema.Table{
		Name:         "aws_xray_encryption_config",
		Description:  "A configuration document that specifies encryption configuration settings.",
		Resolver:     fetchXrayEncryptionConfigs,
		Multiplex:    client.ServiceAccountRegionMultiplexer("xray"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "region"}},
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
				Name:          "key_id",
				Description:   "The ID of the KMS key used for encryption, if applicable.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
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
//                                               Table Resolver Functions1
// ====================================================================================================================

func fetchXrayEncryptionConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Xray
	input := xray.GetEncryptionConfigInput{}
	output, err := svc.GetEncryptionConfig(ctx, &input, func(o *xray.Options) {
		o.Region = c.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	res <- output.EncryptionConfig
	return nil
}
