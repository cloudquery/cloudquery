package athena

import (
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func WorkGroupPreparedStatements() *schema.Table {
	tableName := "aws_athena_work_group_prepared_statements"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/athena/latest/APIReference/API_PreparedStatement.html`,
		Resolver:            fetchAthenaWorkGroupPreparedStatements,
		PreResourceResolver: getWorkGroupPreparedStatement,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "athena"),
		Transform:           client.TransformWithStruct(&types.PreparedStatement{}),
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
