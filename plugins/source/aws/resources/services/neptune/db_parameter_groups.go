package neptune

import (
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DbParameterGroups() *schema.Table {
	tableName := "aws_neptune_db_parameter_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/neptune/latest/userguide/api-parameters.html#DescribeDBClusterParameterGroups`,
		Resolver:    fetchNeptuneDbParameterGroups,
		Transform:   client.TransformWithStruct(&types.DBParameterGroup{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "neptune"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBParameterGroupArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveNeptuneDbParameterGroupTags,
			},
		},

		Relations: []*schema.Table{
			DbParameterGroupDbParameters(),
		},
	}
}
