package quicksight

import (
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DataSets() *schema.Table {
	tableName := "aws_quicksight_data_sets"
	return &schema.Table{
		Name:        tableName,
		Description: "https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DataSetSummary.html",
		Resolver:    fetchQuicksightDataSets,
		Transform:   transformers.TransformWithStruct(&types.DataSetSummary{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "quicksight"),
		Columns:     []schema.Column{client.DefaultAccountIDColumn(true), client.DefaultRegionColumn(true), tagsCol},
		Relations:   []*schema.Table{ingestions()},
	}
}
