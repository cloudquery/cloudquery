package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/cloudquery/plugin-sdk/codegen"
)

func init() {
	add(&Resource{
		DefaultColumns:     []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
		AWSStruct:          &codepipeline.GetPipelineOutput{},
		AWSService:         "CodePipeline",
		Template:           "resource_list_describe",
		PaginatorStruct:    &codepipeline.ListPipelinesOutput{},
		PaginatorGetStruct: &codepipeline.GetPipelineInput{},
		ItemsStruct:        &codepipeline.GetPipelineOutput{},
		SkipFields:         []string{"ResultMetadata"},
		// TODO get tags
	},
		&Resource{
			DefaultColumns: []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
			AWSStruct:      &types.ListWebhookItem{},
			AWSService:     "CodePipeline",
			Template:       "resource_get",
			ItemsStruct:    &codepipeline.ListWebhooksOutput{},
		},
	)
}
