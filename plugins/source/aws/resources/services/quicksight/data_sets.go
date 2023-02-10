package quicksight

import (
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DataSets() *schema.Table {
	return &schema.Table{
		Name:        "aws_quicksight_data_sets",
		Description: "https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DataSetSummary.html",
		Resolver:    fetchQuicksightDataSets,
		Transform:   transformers.TransformWithStruct(&types.DataSetSummary{}),

		Multiplex: client.ServiceAccountRegionMultiplexer("quicksight"),
		Columns: []schema.Column{
			client.AccountPKColumn(false),
			client.RegionPKColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveTags(),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			Ingestions(),
		},
	}
}
