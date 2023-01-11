package neptune

import (
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DbParameterGroupDbParameters() *schema.Table {
	return &schema.Table{
		Name:        "aws_neptune_db_parameter_group_db_parameters",
		Description: `https://docs.aws.amazon.com/neptune/latest/userguide/api-parameters.html#DescribeDBClusterParameters`,
		Resolver:    fetchNeptuneDbParameterGroupDbParameters,
		Transform:   transformers.TransformWithStruct(&types.Parameter{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("neptune"),
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
				Name:     "db_parameter_group_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
