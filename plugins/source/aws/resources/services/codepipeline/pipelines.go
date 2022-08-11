package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource pipelines --config gen.hcl --output .
func Pipelines() *schema.Table {
	return &schema.Table{
		Name:         "aws_codepipeline_pipelines",
		Description:  "Represents the output of a GetPipeline action",
		Resolver:     fetchCodepipelinePipelines,
		Multiplex:    client.ServiceAccountRegionMultiplexer("codepipeline"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "tags",
				Description: "The tags associated with the pipeline.",
				Type:        schema.TypeJSON,
				Resolver:    resolveCodepipelinePipelineTags,
			},
			{
				Name:        "created",
				Description: "The date and time the pipeline was created, in timestamp format",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Metadata.Created"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the pipeline",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Metadata.PipelineArn"),
			},
			{
				Name:        "updated",
				Description: "The date and time the pipeline was last updated, in timestamp format",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Metadata.Updated"),
			},
			{
				Name:        "name",
				Description: "The name of the pipeline",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Pipeline.Name"),
			},
			{
				Name:        "role_arn",
				Description: "The Amazon Resource Name (ARN) for AWS CodePipeline to use to either perform actions with no actionRoleArn, or to use to assume roles for actions with an actionRoleArn",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Pipeline.RoleArn"),
			},
			{
				Name:        "artifact_store_location",
				Description: "The S3 bucket used for storing the artifacts for a pipeline",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Pipeline.ArtifactStore.Location"),
			},
			{
				Name:        "artifact_store_type",
				Description: "The type of the artifact store, such as S3",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Pipeline.ArtifactStore.Type"),
			},
			{
				Name:        "artifact_store_encryption_key_id",
				Description: "The ID used to identify the key",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Pipeline.ArtifactStore.EncryptionKey.Id"),
			},
			{
				Name:        "artifact_store_encryption_key_type",
				Description: "The type of encryption key, such as an AWS Key Management Service (AWS KMS) key When creating or updating a pipeline, the value must be set to 'KMS'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Pipeline.ArtifactStore.EncryptionKey.Type"),
			},
			{
				Name:        "artifact_stores",
				Description: "A mapping of artifactStore objects and their corresponding AWS Regions",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Pipeline.ArtifactStores"),
			},
			{
				Name:        "version",
				Description: "The version number of the pipeline",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Pipeline.Version"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_codepipeline_pipeline_stages",
				Description: "Represents information about a stage and its definition",
				Resolver:    fetchCodepipelinePipelineStages,
				Columns: []schema.Column{
					{
						Name:        "pipeline_cq_id",
						Description: "Unique CloudQuery ID of aws_codepipeline_pipelines table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "stage_order",
						Description: "The stage order in the pipeline.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "name",
						Description: "The name of the stage",
						Type:        schema.TypeString,
					},
					{
						Name:          "blockers",
						Description:   "Reserved for future use",
						Type:          schema.TypeJSON,
						Resolver:      schema.PathResolver("Blockers"),
						IgnoreInTests: true,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_codepipeline_pipeline_stage_actions",
						Description: "Represents information about an action declaration",
						Resolver:    schema.PathTableResolver("Actions"),
						Columns: []schema.Column{
							{
								Name:        "pipeline_stage_cq_id",
								Description: "Unique CloudQuery ID of aws_codepipeline_pipeline_stages table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "category",
								Description: "A category defines what kind of action can be taken in the stage, and constrains the provider type for the action",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ActionTypeId.Category"),
							},
							{
								Name:        "owner",
								Description: "The creator of the action being called",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ActionTypeId.Owner"),
							},
							{
								Name:        "provider",
								Description: "The provider of the service being called by the action",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ActionTypeId.Provider"),
							},
							{
								Name:        "version",
								Description: "A string that describes the action version",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ActionTypeId.Version"),
							},
							{
								Name:        "name",
								Description: "The action declaration's name",
								Type:        schema.TypeString,
							},
							{
								Name:        "configuration",
								Description: "The action's configuration",
								Type:        schema.TypeJSON,
							},
							{
								Name:        "input_artifacts",
								Description: "The name or ID of the artifact consumed by the action, such as a test or build artifact",
								Type:        schema.TypeStringArray,
								Resolver:    schema.PathResolver("InputArtifacts.Name"),
							},
							{
								Name:          "namespace",
								Description:   "The variable namespace associated with the action",
								Type:          schema.TypeString,
								IgnoreInTests: true,
							},
							{
								Name:        "output_artifacts",
								Description: "The name or ID of the result of the action declaration, such as a test or build artifact",
								Type:        schema.TypeStringArray,
								Resolver:    schema.PathResolver("OutputArtifacts.Name"),
							},
							{
								Name:          "region",
								Description:   "The action declaration's AWS Region, such as us-east-1",
								Type:          schema.TypeString,
								IgnoreInTests: true,
							},
							{
								Name:          "role_arn",
								Description:   "The ARN of the IAM service role that performs the declared action",
								Type:          schema.TypeString,
								IgnoreInTests: true,
							},
							{
								Name:        "run_order",
								Description: "The order in which actions are run",
								Type:        schema.TypeBigInt,
							},
						},
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchCodepipelinePipelines(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().CodePipeline
	config := codepipeline.ListPipelinesInput{}
	for {
		response, err := svc.ListPipelines(ctx, &config, func(options *codepipeline.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		for i := range response.Pipelines {
			response, err := svc.GetPipeline(ctx, &codepipeline.GetPipelineInput{Name: response.Pipelines[i].Name}, func(o *codepipeline.Options) {
				o.Region = c.Region
			})
			if err != nil {
				if c.IsNotFoundError(err) {
					continue
				}
				return diag.WrapError(err)
			}
			res <- response
		}

		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func resolveCodepipelinePipelineTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	pipeline := resource.Item.(*codepipeline.GetPipelineOutput)

	cl := meta.(*client.Client)
	svc := cl.Services().CodePipeline
	response, err := svc.ListTagsForResource(ctx, &codepipeline.ListTagsForResourceInput{
		ResourceArn: pipeline.Metadata.PipelineArn,
	}, func(options *codepipeline.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}

	return diag.WrapError(resource.Set(c.Name, client.TagsToMap(response.Tags)))
}
func fetchCodepipelinePipelineStages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	type StageWrapper struct {
		types.StageDeclaration
		StageOrder int32
	}

	r := parent.Item.(*codepipeline.GetPipelineOutput)
	for i, stage := range r.Pipeline.Stages {
		res <- StageWrapper{
			StageDeclaration: stage,
			StageOrder:       int32(i),
		}
	}
	return nil
}
