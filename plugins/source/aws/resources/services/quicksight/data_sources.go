package quicksight

import (
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DataSources() *schema.Table {
	return &schema.Table{
		Name:        "aws_quicksight_data_sources",
		Description: "https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DataSource.html",
		Resolver:    fetchQuicksightDataSources,
		Transform: transformers.TransformWithStruct(&types.DataSource{},
			transformers.WithPrimaryKeys("Arn"),
			transformers.WithSkipFields("AlternateDataSourceParameters"),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer("quicksight"),
		Columns:   []schema.Column{client.DefaultAccountIDColumn(true), client.DefaultRegionColumn(true), tagsCol},
	}
}
