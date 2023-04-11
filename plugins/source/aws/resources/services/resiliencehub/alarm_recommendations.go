package resiliencehub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func alarmRecommendations() *schema.Table {
	tableName := "aws_resiliencehub_alarm_recommendations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_AlarmRecommendation.html`,
		Resolver:    fetchAlarmRecommendations,
		Transform:   transformers.TransformWithStruct(&types.AlarmRecommendation{}, transformers.WithPrimaryKeys("RecommendationId")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "resiliencehub"),
		Columns:     []schema.Column{client.DefaultAccountIDColumn(false), client.DefaultRegionColumn(false), appARN, assessmentARN},
	}
}

func fetchAlarmRecommendations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Resiliencehub
	p := resiliencehub.NewListAlarmRecommendationsPaginator(svc, &resiliencehub.ListAlarmRecommendationsInput{AssessmentArn: parent.Item.(*types.AppAssessment).AppArn})
	for p.HasMorePages() {
		out, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- out.AlarmRecommendations
	}
	return nil
}
