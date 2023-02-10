package applicationautoscaling

import (
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Policies() *schema.Table {
	return &schema.Table{
		Name:        "aws_applicationautoscaling_policies",
		Description: `https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_ScalingPolicy.html`,
		Resolver:    fetchApplicationautoscalingPolicies,
		Multiplex:   client.ServiceAccountRegionNamespaceMultiplexer("application-autoscaling"),
		Transform:   transformers.TransformWithStruct(&types.ScalingPolicy{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PolicyARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
