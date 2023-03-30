package athena

import (
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func WorkGroupQueryExecutions() *schema.Table {
	tableName := "aws_athena_work_group_query_executions"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/athena/latest/APIReference/API_QueryExecution.html`,
		Resolver:            fetchAthenaWorkGroupQueryExecutions,
		PreResourceResolver: getWorkGroupQueryExecution,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "athena"),
		Transform:           transformers.TransformWithStruct(&types.QueryExecution{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "work_group_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
