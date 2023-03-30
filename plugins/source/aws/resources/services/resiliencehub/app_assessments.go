package resiliencehub

import (
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func appAssesments() *schema.Table {
	tableName := "aws_resiliencehub_app_assessments"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_AppAssessment.html`,
		Resolver:            fetchAppAssessments,
		PreResourceResolver: describeAppAssessments,
		Transform:           client.TransformWithStruct(&types.AppAssessment{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "resiliencehub"),
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
