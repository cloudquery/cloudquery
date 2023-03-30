package athena

import (
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func WorkGroups() *schema.Table {
	tableName := "aws_athena_work_groups"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/athena/latest/APIReference/API_WorkGroup.html`,
		Resolver:            fetchAthenaWorkGroups,
		PreResourceResolver: getWorkGroup,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "athena"),
		Transform:           transformers.TransformWithStruct(&types.WorkGroup{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveAthenaWorkGroupArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveAthenaWorkGroupTags,
			},
		},

		Relations: []*schema.Table{
			WorkGroupPreparedStatements(),
			WorkGroupQueryExecutions(),
			WorkGroupNamedQueries(),
		},
	}
}
