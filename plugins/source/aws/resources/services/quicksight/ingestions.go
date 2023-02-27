package quicksight

import (
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ingestions() *schema.Table {
	return &schema.Table{
		Name:        "aws_quicksight_ingestions",
		Description: "https://docs.aws.amazon.com/quicksight/latest/APIReference/API_Ingestion.html",
		Resolver:    fetchQuicksightIngestions,
		Transform:   transformers.TransformWithStruct(&types.Ingestion{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("quicksight"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			tagsCol,
			{
				Name:            "data_set_arn",
				Type:            schema.TypeString,
				Resolver:        schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	}
}
