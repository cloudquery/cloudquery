package athena

import (
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func WorkGroups() *schema.Table {
	return &schema.Table{
		Name:                "aws_athena_work_groups",
		Description:         `https://docs.aws.amazon.com/athena/latest/APIReference/API_WorkGroup.html`,
		Resolver:            fetchAthenaWorkGroups,
		PreResourceResolver: getWorkGroup,
		Multiplex:           client.ServiceAccountRegionMultiplexer("athena"),
		Transform: transformers.TransformWithStruct(&types.WorkGroup{}),
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
