package quicksight

import (
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Templates() *schema.Table {
	return &schema.Table{
		Name:        "aws_quicksight_templates",
		Description: "https://docs.aws.amazon.com/quicksight/latest/APIReference/API_TemplateSummary.html",
		Resolver:    fetchQuicksightTemplates,
		Transform:   transformers.TransformWithStruct(&types.TemplateSummary{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("quicksight"),
		Columns:     []schema.Column{client.DefaultAccountIDColumn(true), client.DefaultRegionColumn(true), tagsCol},
	}
}
