package resiliencehub

import (
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func sopRecommendations() *schema.Table {
	tableName := "aws_resiliencehub_sop_recommendations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_SopRecommendation.html`,
		Resolver:    fetchSopRecommendations,
		Transform:   client.TransformWithStruct(&types.SopRecommendation{}, transformers.WithPrimaryKeys("RecommendationId")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "resiliencehub"),
		Columns:     []schema.Column{client.DefaultAccountIDColumn(false), client.DefaultRegionColumn(false), appARN, assessmentARN},
	}
}
