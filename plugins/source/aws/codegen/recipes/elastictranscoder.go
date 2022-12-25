package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"
	"github.com/cloudquery/plugin-sdk/codegen"
)

func ElastictranscoderResources() []*Resource {
	return []*Resource{
		{
			TableDefinition: codegen.TableDefinition{
				Service:      "elastictranscoder",
				SubService:   "pipeline_jobs",
				Struct:       new(types.Job),
				Description:  "https://docs.aws.amazon.com/elastictranscoder/latest/developerguide/list-jobs-by-pipeline.html",
				Multiplex:    "", // relation for pipelines
				PKColumns:    []string{"arn"},
				ExtraColumns: defaultRegionalColumns,
			},
		},
		{
			TableDefinition: codegen.TableDefinition{
				Service:      "elastictranscoder",
				SubService:   "pipelines",
				Struct:       new(types.Pipeline),
				Description:  "https://docs.aws.amazon.com/elastictranscoder/latest/developerguide/list-pipelines.html",
				Multiplex:    `client.ServiceAccountRegionMultiplexer("elastictranscoder")`,
				PKColumns:    []string{"arn"},
				ExtraColumns: defaultRegionalColumns,
				Relations:    []string{"PipelineJobs()"},
			},
		},
		{
			TableDefinition: codegen.TableDefinition{
				Service:      "elastictranscoder",
				SubService:   "presets",
				Struct:       new(types.Preset),
				Description:  "https://docs.aws.amazon.com/elastictranscoder/latest/developerguide/list-presets.html",
				Multiplex:    `client.ServiceAccountRegionMultiplexer("elastictranscoder")`,
				PKColumns:    []string{"arn"},
				ExtraColumns: defaultRegionalColumns,
			},
		},
	}
}
