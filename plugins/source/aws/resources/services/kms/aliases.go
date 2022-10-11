// Code generated by codegen; DO NOT EDIT.

package kms

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Aliases() *schema.Table {
	return &schema.Table{
		Name:      "aws_kms_aliases",
		Resolver:  fetchKmsAliases,
		Multiplex: client.ServiceAccountRegionMultiplexer("kms"),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AliasArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "alias_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AliasName"),
			},
			{
				Name:     "creation_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationDate"),
			},
			{
				Name:     "last_updated_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastUpdatedDate"),
			},
			{
				Name:     "target_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TargetKeyId"),
			},
		},
	}
}
