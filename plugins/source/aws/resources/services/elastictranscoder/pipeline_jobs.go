package elastictranscoder

import (
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func PipelineJobs() *schema.Table {
	return &schema.Table{
		Name:        "aws_elastictranscoder_pipeline_jobs",
		Description: `https://docs.aws.amazon.com/elastictranscoder/latest/developerguide/list-jobs-by-pipeline.html`,
		Resolver:    fetchElastictranscoderPipelineJobs,
		Transform:   transformers.TransformWithStruct(&types.Job{}),
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
