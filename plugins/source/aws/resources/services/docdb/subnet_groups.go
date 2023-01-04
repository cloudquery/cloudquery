package docdb

import (
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SubnetGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_docdb_subnet_groups",
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBSubnetGroup.html`,
		Resolver:    fetchDocdbSubnetGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer("docdb"),
		Transform: transformers.TransformWithStruct(&types.DBSubnetGroup{}),
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveDBSubnetGroupTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBSubnetGroupArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
