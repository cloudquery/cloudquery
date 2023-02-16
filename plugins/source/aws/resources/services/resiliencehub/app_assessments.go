package resiliencehub

import (
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func appAssesments() *schema.Table {
	return &schema.Table{
		Name:                "aws_resiliencehub_app_assessments",
		Description:         `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_AppAssessment.html`,
		Resolver:            fetchAppAssessments,
		PreResourceResolver: describeAppAssessments,
		Transform:           transformers.TransformWithStruct(&types.AppAssessment{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer("resiliencehub"),
		Columns:             []schema.Column{client.DefaultAccountIDColumn(false), client.DefaultRegionColumn(false), appARNTop, arnColumn("AssessmentArn")},
		Relations: []*schema.Table{
			appComponentCompliances(),
			appComponentRecommendations(),
			testRecommendations(),
			alarmRecommendations(),
			recommendationTemplates(),
			sopRecommendations(),
		},
	}
}
