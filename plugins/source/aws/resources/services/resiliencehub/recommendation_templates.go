package resiliencehub

import (
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func recommendationTemplates() *schema.Table {
	tableName := "aws_resiliencehub_recommendation_templates"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_RecommendationTemplate.html`,
		Resolver:    fetchRecommendationTemplates,
		Transform:   transformers.TransformWithStruct(&types.RecommendationTemplate{}, transformers.WithPrimaryKeys("AppArn", "AssessmentArn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "resiliencehub"),
		Columns:     []schema.Column{client.DefaultAccountIDColumn(false), client.DefaultRegionColumn(false), arnColumn("RecommendationTemplateArn")},
	}
}
