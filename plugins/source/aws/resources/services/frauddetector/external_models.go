package frauddetector

import (
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ExternalModels() *schema.Table {
	tableName := "aws_frauddetector_external_models"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/frauddetector/latest/api/API_ExternalModel.html`,
		Resolver:    fetchFrauddetectorExternalModels,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "frauddetector"),
		Transform:   client.TransformWithStruct(&types.ExternalModel{}),
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
		},
	}
}
