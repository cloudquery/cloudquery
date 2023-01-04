package athena

import (
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func WorkGroupPreparedStatements() *schema.Table {
	return &schema.Table{
		Name:                "aws_athena_work_group_prepared_statements",
		Description:         `https://docs.aws.amazon.com/athena/latest/APIReference/API_PreparedStatement.html`,
		Resolver:            fetchAthenaWorkGroupPreparedStatements,
		PreResourceResolver: getWorkGroupPreparedStatement,
		Multiplex:           client.ServiceAccountRegionMultiplexer("athena"),
		Transform:           transformers.TransformWithStruct(&types.PreparedStatement{}),
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
				Name:     "work_group_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
