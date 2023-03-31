package elastictranscoder

import (
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Pipelines() *schema.Table {
	tableName := "aws_elastictranscoder_pipelines"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/elastictranscoder/latest/developerguide/list-pipelines.html`,
		Resolver:    fetchElastictranscoderPipelines,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elastictranscoder"),
		Transform:   transformers.TransformWithStruct(&types.Pipeline{}),
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

		Relations: []*schema.Table{
			pipelineJobs(),
		},
	}
}
