package emr

import (
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func BlockPublicAccessConfigs() *schema.Table {
	return &schema.Table{
		Name:      "aws_emr_block_public_access_configs",
		Resolver:  fetchEmrBlockPublicAccessConfigs,
		Multiplex: client.ServiceAccountRegionMultiplexer("elasticmapreduce"),
		Transform: transformers.TransformWithStruct(&emr.GetBlockPublicAccessConfigurationOutput{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
