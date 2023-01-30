package rds

import (
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func EngineVersions() *schema.Table {
	return &schema.Table{
		Name:        "aws_rds_engine_versions",
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBEngineVersion.html`,
		Resolver:    fetchRdsEngineVersions,
		Transform:   transformers.TransformWithStruct(&types.DBEngineVersion{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("rds"),
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
			{
				Name:     "_engine_version_hash",
				Type:     schema.TypeString,
				Resolver: calculateUniqueHash,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tag_list",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTagField("TagList"),
			},
		},

		Relations: []*schema.Table{
			ClusterParameters(),
		},
	}
}
