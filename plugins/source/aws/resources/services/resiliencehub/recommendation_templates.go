package resiliencehub

import (
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func recommendationTemplates() *schema.Table {
	return &schema.Table{
		Name:        "aws_resiliencehub_recommendation_templates",
		Description: `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_RecommendationTemplate.html`,
		Resolver:    fetchRecommendationTemplates,
		Transform:   transformers.TransformWithStruct(&types.RecommendationTemplate{}, transformers.WithPrimaryKeys("Name")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("resiliencehub"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
