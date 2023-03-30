package frauddetector

import (
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Labels() *schema.Table {
	tableName := "aws_frauddetector_labels"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/frauddetector/latest/api/API_Label.html`,
		Resolver:    fetchFrauddetectorLabels,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "frauddetector"),
		Transform:   client.TransformWithStruct(&types.Label{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveResourceTags,
			},
		},
	}
}
