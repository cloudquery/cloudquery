package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SagemakerTrainingJobs() *schema.Table {
	return &schema.Table{
		Name:        "aws_sagemaker_training_jobs",
		Description: "Provides summary information about a training job.",
		Resolver: func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
			return diag.WrapError(client.ListAndDetailResolver(ctx, meta, res, listSagemakerTrainingJobs, sagemakerTrainingJobsDetail))
		},
		Multiplex:     client.ServiceAccountRegionMultiplexer("api.sagemaker"),
		IgnoreInTests: true,
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
				Name:        "auto_ml_job_arn",
				Description: "The Amazon Resource Name (ARN) of an AutoML job.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AutoMLJobArn"),
			},
			{
				Name:        "billable_time_in_seconds",
				Description: "The billable time in seconds. Billable time refers to the absolute wall-clock time.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "enable_managed_spot_training",
				Description: "A Boolean indicating whether managed spot training is enabled (True) or not (False).",
				Type:        schema.TypeBool,
			},
			{
				Name:        "enable_network_isolation",
				Description: "If you want to allow inbound or outbound network calls, except for calls between peers within a training cluster for distributed training, choose True. If you enable network isolation for training jobs that are configured to use a VPC, Amazon SageMaker downloads and uploads customer data and model artifacts through the specified VPC, but the training container does not have network access.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "enable_inter_container_traffic_encryption",
				Description: "To encrypt all communications between ML compute instances in distributed training, choose True. Encryption provides greater security for distributed training, but training might take longer. How long it takes depends on the amount of communication between compute instances, especially if you use a deep learning algorithms in distributed training.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "failure_reason",
				Description: "If the training job failed, the reason it failed.",
				Type:        schema.TypeString,
			},
			{
				Name:        "labeling_job_arn",
				Description: "The Amazon Resource Name (ARN) of the Amazon SageMaker Ground Truth labeling job that created the transform or training job.",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_modified_time",
				Description: "A timestamp that indicates when the status of the training job was last modified.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "profiling_status",
				Description: "Profiling status of a training job.",
				Type:        schema.TypeString,
			},
			{
				Name:        "role_arn",
				Description: "The Amazon Web Services Identity and Access Management (IAM) role configured for the training job.",
				Type:        schema.TypeString,
			},
			{
				Name:        "secondary_status",
				Description: "Provides detailed information about the state of the training job.",
				Type:        schema.TypeString,
			},
			{
				Name:        "training_end_time",
				Description: "Indicates the time when the training job ends on training instances.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "training_start_time",
				Description: "Indicates the time when the training job starts on training instances.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "training_time_in_seconds",
				Description: "The training time in seconds.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "tuning_job_arn",
				Description: "The Amazon Resource Name (ARN) of the associated hyperparameter tuning job if the training job was launched by a hyperparameter tuning job.",
				Type:        schema.TypeString,
			},
			{
				Name:        "checkpoint_config",
				Description: "Contains information about the output location for managed spot training checkpoint data.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "environment",
				Description: "The environment variables to set in the Docker container.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "experiment_config",
				Description: "Associates a SageMaker job as a trial component with an experiment and trial.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "hyper_parameters",
				Description: "Algorithm-specific parameters.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "model_artifacts",
				Description: "Information about the Amazon S3 location that is configured for storing model artifacts.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "output_data_config",
				Description: "The S3 path where model artifacts that you configured when creating the job are stored.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "profiler_config",
				Description: "Configuration information for Debugger system monitoring, framework profiling, and storage paths.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "resource_config",
				Description: "Resources, including ML compute instances and ML storage volumes, that are configured for model training.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "stopping_condition",
				Description: "Specifies a limit to how long a model training job can run.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "tensor_board_output_config",
				Description: "Configuration of storage locations for the Debugger TensorBoard output data.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "vpc_config",
				Description: "A VpcConfig object that specifies the VPC that this training job has access to.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "tags",
				Description: "The tags associated with the model.",
				Type:        schema.TypeJSON,
				Resolver:    resolveSagemakerTrainingJobTags,
			},
			{
				Name:        "creation_time",
				Description: "A timestamp that shows when the training job was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the training job.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TrainingJobArn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "name",
				Description: "The name of the training job.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TrainingJobName"),
			},
			{
				Name:        "training_job_status",
				Description: "The status of the training job.",
				Type:        schema.TypeString,
			},
			{
				Name:     "secondary_status_transitions",
				Type:     schema.TypeJSON,
			},
			{
				Name:     "final_metric_data_list",
				Type:     schema.TypeJSON,
			},
			{
				Name:     "algorithm_specification",
				Description: "Specifies the training algorithm to use in a CreateTrainingJob request",
				Type: 	 schema.TypeJSON,
			},
			{
				Name:     "debug_hook_config",
				Description: "Configuration information for the Debugger hook parameters, metric and tensor collections, and storage paths",
				Type: 	 schema.TypeJSON,
			},
			{
				Name:     "debug_rule_configurations",
				Description: "Configuration information for SageMaker Debugger rules for debugging",
				Type: 	 schema.TypeJSON,
			},
			{
				Name: "debug_rule_evaluation_statuses",
				Description: "Information about the status of the rule evaluation.",
				Type: 	 schema.TypeJSON,
			},
			{
				Name: "input_data_config",
				Description: "A channel is a named input source that training algorithms can consume.",
				Type: 	 schema.TypeJSON,
			},
			{
				Name: "profiler_rule_configurations",
				Description: "Configuration information for profiling rules.",
				Type: 	 schema.TypeJSON,
			},
			{
				Name: "profiler_rule_evaluation_statuses",
				Description: "Information about the status of the rule evaluation.",
				Type: 	 schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func sagemakerTrainingJobsDetail(ctx context.Context, meta schema.ClientMeta, resultsChan chan<- interface{}, errorChan chan<- error, detail interface{}) {
	c := meta.(*client.Client)
	svc := c.Services().SageMaker
	n := detail.(types.TrainingJobSummary)
	config := sagemaker.DescribeTrainingJobInput{
		TrainingJobName: n.TrainingJobName,
	}
	response, err := svc.DescribeTrainingJob(ctx, &config)
	if err != nil {
		errorChan <- err
		return
	}
	resultsChan <- response
}

func listSagemakerTrainingJobs(ctx context.Context, meta schema.ClientMeta, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().SageMaker
	config := sagemaker.ListTrainingJobsInput{}

	for {
		response, err := svc.ListTrainingJobs(ctx, &config)
		if err != nil {
			return err
		}
		for _, d := range response.TrainingJobSummaries {
			res <- d
		}
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}


func resolveSagemakerTrainingJobTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(*sagemaker.DescribeTrainingJobOutput)
	if r == nil {
		return nil
	}

	c := meta.(*client.Client)
	svc := c.Services().SageMaker
	config := sagemaker.ListTagsInput{
		ResourceArn: r.TrainingJobArn,
	}
	response, err := svc.ListTags(ctx, &config)
	if err != nil {
		return err
	}

	tags := make(map[string]*string, len(response.Tags))
	for _, t := range response.Tags {
		tags[*t.Key] = t.Value
	}

	return diag.WrapError(resource.Set("tags", tags))
}

