package resiliencehub

import (
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func appComponentRecommendations() *schema.Table {
	return &schema.Table{
		Name:        "aws_resiliencehub_component_recommendations",
		Description: `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_ComponentRecommendation.html`,
		Resolver:    fetchComponentRecommendations,
		Transform:   transformers.TransformWithStruct(&types.ComponentRecommendation{}, transformers.WithPrimaryKeys("AppComponentName")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("resiliencehub"),
		Columns:     []schema.Column{client.DefaultAccountIDColumn(false), client.DefaultRegionColumn(false), appARN, assessmentARN},
	}
}
