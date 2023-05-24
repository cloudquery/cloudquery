// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	
)

//go:generate mockgen -package=mocks -destination=../mocks/elastictranscoder.go -source=elastictranscoder.go ElastictranscoderClient
type ElastictranscoderClient interface {
	ListJobsByPipeline(context.Context, *elastictranscoder.ListJobsByPipelineInput, ...func(*elastictranscoder.Options)) (*elastictranscoder.ListJobsByPipelineOutput, error)
	ListJobsByStatus(context.Context, *elastictranscoder.ListJobsByStatusInput, ...func(*elastictranscoder.Options)) (*elastictranscoder.ListJobsByStatusOutput, error)
	ListPipelines(context.Context, *elastictranscoder.ListPipelinesInput, ...func(*elastictranscoder.Options)) (*elastictranscoder.ListPipelinesOutput, error)
	ListPresets(context.Context, *elastictranscoder.ListPresetsInput, ...func(*elastictranscoder.Options)) (*elastictranscoder.ListPresetsOutput, error)
}
