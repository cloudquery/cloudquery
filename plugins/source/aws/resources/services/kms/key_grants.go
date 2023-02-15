package kms

import (
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func KeyGrants() *schema.Table {
	return &schema.Table{
		Name:        "aws_kms_key_grants",
		Description: `https://docs.aws.amazon.com/kms/latest/APIReference/API_GrantListEntry.html`,
		Resolver:    fetchKmsKeyGrants,
		Transform:   transformers.TransformWithStruct(&types.GrantListEntry{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("kms"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "key_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "grant_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GrantId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
