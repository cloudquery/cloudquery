package resiliencehub

import (
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SuggestedResiliencyPolicies() *schema.Table {
	return &schema.Table{
		Name:        "aws_resiliencehub_suggested_resiliency_policies",
		Description: `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_ResiliencyPolicy.html`,
		Resolver:    fetchSuggestedResiliencyPolicies,
		Transform:   transformers.TransformWithStruct(&types.ResiliencyPolicy{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("resiliencehub"),
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
				Resolver: schema.PathResolver("PolicyArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
