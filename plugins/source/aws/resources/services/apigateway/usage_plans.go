package apigateway

import (
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func UsagePlans() *schema.Table {
	return &schema.Table{
		Name:        "aws_apigateway_usage_plans",
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_UsagePlan.html`,
		Resolver:    fetchApigatewayUsagePlans,
		Multiplex:   client.ServiceAccountRegionMultiplexer("apigateway"),
		Transform:  transformers.TransformWithStruct(&types.UsagePlan{}),
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
				Resolver: resolveApigatewayUsagePlanArn,
			},
		},

		Relations: []*schema.Table{
			UsagePlanKeys(),
		},
	}
}
