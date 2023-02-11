package athena

import (
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func WorkGroupNamedQueries() *schema.Table {
	return &schema.Table{
		Name:                "aws_athena_work_group_named_queries",
		Description:         `https://docs.aws.amazon.com/athena/latest/APIReference/API_NamedQuery.html`,
		Resolver:            fetchAthenaWorkGroupNamedQueries,
		PreResourceResolver: getWorkGroupNamedQuery,
		Multiplex:           client.ServiceAccountRegionMultiplexer("athena"),
		Transform:           transformers.TransformWithStruct(&types.NamedQuery{}),
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
