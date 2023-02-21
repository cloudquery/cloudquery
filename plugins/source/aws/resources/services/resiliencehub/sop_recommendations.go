package resiliencehub

import (
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func sopRecommendations() *schema.Table {
	return &schema.Table{
		Name:        "aws_resiliencehub_sop_recommendations",
		Description: `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_SopRecommendation.html`,
		Resolver:    fetchSopRecommendations,
		Transform:   transformers.TransformWithStruct(&types.SopRecommendation{}, transformers.WithPrimaryKeys("RecommendationId")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("resiliencehub"),
		Columns:     []schema.Column{client.DefaultAccountIDColumn(false), client.DefaultRegionColumn(false), appARN, assessmentARN},
	}
}
