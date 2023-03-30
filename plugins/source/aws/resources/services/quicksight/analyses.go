package quicksight

import (
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Analyses() *schema.Table {
	tableName := "aws_quicksight_analyses"
	return &schema.Table{
		Name:                tableName,
		Description:         "https://docs.aws.amazon.com/quicksight/latest/APIReference/API_Analysis.html",
		Resolver:            fetchQuicksightAnalyses,
		PreResourceResolver: getAnalysis,
		Transform:           client.TransformWithStruct(&types.Analysis{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "quicksight"),
		Columns:             []schema.Column{client.DefaultAccountIDColumn(true), client.DefaultRegionColumn(true), tagsCol},
	}
}
