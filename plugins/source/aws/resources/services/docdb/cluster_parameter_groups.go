package docdb

import (
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ClusterParameterGroups() *schema.Table {
	tableName := "aws_docdb_cluster_parameter_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBClusterParameterGroup.html`,
		Resolver:    fetchDocdbClusterParameterGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "docdb"),
		Transform:   transformers.TransformWithStruct(&types.DBClusterParameterGroup{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveDBClusterParameterGroupTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterParameterGroupArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "parameters",
				Type:     schema.TypeJSON,
				Resolver: resolveDocdbClusterParameterGroupParameters,
			},
			{
				Name:     "db_cluster_parameter_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterParameterGroupName"),
			},
			{
				Name:     "db_parameter_group_family",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBParameterGroupFamily"),
			},
		},
	}
}
