// Code generated by codegen; DO NOT EDIT.

package xray

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func EncryptionConfig() *schema.Table {
	return &schema.Table{
		Name:        "aws_xray_encryption_config",
		Description: "https://docs.aws.amazon.com/xray/latest/api/API_EncryptionConfig.html",
		Resolver:    fetchXrayEncryptionConfig,
		Multiplex:   client.ServiceAccountRegionMultiplexer("xray"),
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
				Name:     "key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeyId"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}
