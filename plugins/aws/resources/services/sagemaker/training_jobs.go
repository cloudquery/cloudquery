package sagemaker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SagemakerTrainingJobs() *schema.Table {
	return &schema.Table{
		Name:          "aws_sagemaker_training_jobs",
		Description:   "Provides summary information about a training job.",
		Resolver:      fetchSagemakerTrainingJobs,
		Multiplex:     client.ServiceAccountRegionMultiplexer("api.sagemaker"),
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Resolver:    resolveSagemakerTrainingJobCheckpointConfig,
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
				Resolver:    resolveSagemakerTrainingJobExperimentConfig,
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
				Resolver:    resolveSagemakerTrainingJobModelArtifacts,
			},
			{
				Name:        "output_data_config",
				Description: "The S3 path where model artifacts that you configured when creating the job are stored.",
				Type:        schema.TypeJSON,
				Resolver:    resolveSagemakerTrainingJobOutputDataConfig,
			},
			{
				Name:        "profiler_config",
				Description: "Configuration information for Debugger system monitoring, framework profiling, and storage paths.",
				Type:        schema.TypeJSON,
				Resolver:    resolveSagemakerTrainingJobProfilerConfig,
			},
			{
				Name:        "resource_config",
				Description: "Resources, including ML compute instances and ML storage volumes, that are configured for model training.",
				Type:        schema.TypeJSON,
				Resolver:    resolveSagemakerTrainingJobResourceConfig,
			},
			{
				Name:        "stopping_condition",
				Description: "Specifies a limit to how long a model training job can run.",
				Type:        schema.TypeJSON,
				Resolver:    resolveSagemakerTrainingJobStoppingCondition,
			},
			{
				Name:        "tensor_board_output_config",
				Description: "Configuration of storage locations for the Debugger TensorBoard output data.",
				Type:        schema.TypeJSON,
				Resolver:    resolveSagemakerTrainingJobTensorBoardOutputConfig,
			},
			{
				Name:        "vpc_config",
				Description: "A VpcConfig object that specifies the VPC that this training job has access to.",
				Type:        schema.TypeJSON,
				Resolver:    resolveSagemakerTrainingJobVpcConfig,
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
				Resolver: resolveSagemakerTrainingJobSecondaryStatusTransitions,
			},
			{
				Name:     "final_metric_data_list",
				Type:     schema.TypeJSON,
				Resolver: resolveSagemakerTrainingJobFinalMetricDataList,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_sagemaker_training_job_algorithm_specification",
				Description: "Specifies the training algorithm to use in a CreateTrainingJob request",
				Resolver:    fetchSagemakerTrainingJobAlgorithmSpecifications,
				Columns: []schema.Column{
					{
						Name:        "training_job_cq_id",
						Description: "Unique CloudQuery ID of aws_sagemaker_training_jobs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "training_input_mode",
						Description: "The training input mode that the algorithm supports",
						Type:        schema.TypeString,
					},
					{
						Name:        "algorithm_name",
						Description: "The name of the algorithm resource to use for the training job",
						Type:        schema.TypeString,
					},
					{
						Name:        "enable_sage_maker_metrics_time_series",
						Description: "To generate and save time-series metrics during training, set to true",
						Type:        schema.TypeBool,
					},
					{
						Name:        "metric_definitions",
						Description: "A list of metric definition objects",
						Type:        schema.TypeJSON,
						Resolver:    resolveSagemakerTrainingJobAlgorithmSpecificationsMetricDefinitions,
					},
					{
						Name:        "training_image",
						Description: "The registry path of the Docker image that contains the training algorithm",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_sagemaker_training_job_debug_hook_config",
				Description: "Configuration information for the Debugger hook parameters, metric and tensor collections, and storage paths",
				Resolver:    fetchSagemakerTrainingJobDebugHookConfigs,
				Columns: []schema.Column{
					{
						Name:        "training_job_cq_id",
						Description: "Unique CloudQuery ID of aws_sagemaker_training_jobs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "s3_output_path",
						Description: "Path to Amazon S3 storage location for metrics and tensors. ",
						Type:        schema.TypeString,
					},
					{
						Name:        "collection_configurations",
						Description: "Configuration information for Debugger tensor collections",
						Type:        schema.TypeJSON,
						Resolver:    resolveSagemakerTrainingJobDebugHookConfigsCollectionConfigurations,
					},
					{
						Name:        "hook_parameters",
						Description: "Configuration information for the Debugger hook parameters.",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "local_path",
						Description: "Path to local storage location for metrics and tensors",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_sagemaker_training_job_debug_rule_configurations",
				Description: "Configuration information for SageMaker Debugger rules for debugging",
				Resolver:    fetchSagemakerTrainingJobDebugRuleConfigurations,
				Columns: []schema.Column{
					{
						Name:        "training_job_cq_id",
						Description: "Unique CloudQuery ID of aws_sagemaker_training_jobs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "rule_configuration_name",
						Description: "The name of the rule configuration",
						Type:        schema.TypeString,
					},
					{
						Name:        "rule_evaluator_image",
						Description: "The Amazon Elastic Container (ECR) Image for the managed rule evaluation. ",
						Type:        schema.TypeString,
					},
					{
						Name:        "instance_type",
						Description: "The instance type to deploy a Debugger custom rule for debugging a training job.",
						Type:        schema.TypeString,
					},
					{
						Name:        "local_path",
						Description: "Path to local storage location for output of rules",
						Type:        schema.TypeString,
					},
					{
						Name:        "rule_parameters",
						Description: "Runtime configuration for rule container.",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "s3_output_path",
						Description: "Path to Amazon S3 storage location for rules.",
						Type:        schema.TypeString,
					},
					{
						Name:        "volume_size_in_gb",
						Description: "The size, in GB, of the ML storage volume attached to the processing instance.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("VolumeSizeInGB"),
					},
				},
			},
			{
				Name:        "aws_sagemaker_training_job_debug_rule_evaluation_statuses",
				Description: "Information about the status of the rule evaluation.",
				Resolver:    fetchSagemakerTrainingJobDebugRuleEvaluationStatuses,
				Columns: []schema.Column{
					{
						Name:        "training_job_cq_id",
						Description: "Unique CloudQuery ID of aws_sagemaker_training_jobs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "last_modified_time",
						Description: "Timestamp when the rule evaluation status was last modified.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "rule_configuration_name",
						Description: "The name of the rule configuration.",
						Type:        schema.TypeString,
					},
					{
						Name:        "rule_evaluation_job_arn",
						Description: "The Amazon Resource Name (ARN) of the rule evaluation job.",
						Type:        schema.TypeString,
					},
					{
						Name:        "rule_evaluation_status",
						Description: "Status of the rule evaluation.",
						Type:        schema.TypeString,
					},
					{
						Name:        "status_details",
						Description: "Details from the rule evaluation.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_sagemaker_training_job_input_data_config",
				Description: "A channel is a named input source that training algorithms can consume.",
				Resolver:    fetchSagemakerTrainingJobInputDataConfigs,
				Columns: []schema.Column{
					{
						Name:        "training_job_cq_id",
						Description: "Unique CloudQuery ID of aws_sagemaker_training_jobs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "channel_name",
						Description: "The name of the channel. ",
						Type:        schema.TypeString,
					},
					{
						Name:        "data_source_file_directory_path",
						Description: "The full path to the directory to associate with the channel. ",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DataSource.FileSystemDataSource.DirectoryPath"),
					},
					{
						Name:        "data_source_file_system_access_mode",
						Description: "The access mode of the mount of the directory associated with the channel",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DataSource.FileSystemDataSource.FileSystemAccessMode"),
					},
					{
						Name:        "data_source_file_system_id",
						Description: "The file system id. ",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DataSource.FileSystemDataSource.FileSystemId"),
					},
					{
						Name:        "data_source_file_system_type",
						Description: "The file system type. ",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DataSource.FileSystemDataSource.FileSystemType"),
					},
					{
						Name:        "data_source_s3_data_type",
						Description: "If you choose S3Prefix, S3Uri identifies a key name prefix",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DataSource.S3DataSource.S3DataType"),
					},
					{
						Name:        "data_source_s3_uri",
						Description: "Depending on the value specified for the S3DataType, identifies either a key name prefix or a manifest",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DataSource.S3DataSource.S3Uri"),
					},
					{
						Name:        "data_source_attribute_names",
						Description: "A list of one or more attribute names to use that are found in a specified augmented manifest file.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("DataSource.S3DataSource.AttributeNames"),
					},
					{
						Name:        "data_source_s3_data_distribution_type",
						Description: "If you want Amazon SageMaker to replicate the entire dataset on each ML compute instance that is launched for model training, specify FullyReplicated",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DataSource.S3DataSource.S3DataDistributionType"),
					},
					{
						Name:        "compression_type",
						Description: "If training data is compressed, the compression type",
						Type:        schema.TypeString,
					},
					{
						Name:        "content_type",
						Description: "The MIME type of the data.",
						Type:        schema.TypeString,
					},
					{
						Name:        "input_mode",
						Description: "(Optional) The input mode to use for the data channel in a training job",
						Type:        schema.TypeString,
					},
					{
						Name:        "record_wrapper_type",
						Description: "Specify RecordIO as the value when input data is in raw format but the training algorithm requires the RecordIO format",
						Type:        schema.TypeString,
					},
					{
						Name:        "shuffle_config_seed",
						Description: "Determines the shuffling order in ShuffleConfig value. ",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ShuffleConfig.Seed"),
					},
				},
			},
			{
				Name:        "aws_sagemaker_training_job_profiler_rule_configurations",
				Description: "Configuration information for profiling rules.",
				Resolver:    fetchSagemakerTrainingJobProfilerRuleConfigurations,
				Columns: []schema.Column{
					{
						Name:        "training_job_cq_id",
						Description: "Unique CloudQuery ID of aws_sagemaker_training_jobs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "rule_configuration_name",
						Description: "The name of the rule configuration",
						Type:        schema.TypeString,
					},
					{
						Name:        "rule_evaluator_image",
						Description: "The Amazon Elastic Container (ECR) Image for the managed rule evaluation. ",
						Type:        schema.TypeString,
					},
					{
						Name:        "instance_type",
						Description: "The instance type to deploy a Debugger custom rule for profiling a training job.",
						Type:        schema.TypeString,
					},
					{
						Name:        "local_path",
						Description: "Path to local storage location for output of rules",
						Type:        schema.TypeString,
					},
					{
						Name:        "rule_parameters",
						Description: "Runtime configuration for rule container.",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "s3_output_path",
						Description: "Path to Amazon S3 storage location for rules.",
						Type:        schema.TypeString,
					},
					{
						Name:        "volume_size_in_gb",
						Description: "The size, in GB, of the ML storage volume attached to the processing instance.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("VolumeSizeInGB"),
					},
				},
			},
			{
				Name:        "aws_sagemaker_training_job_profiler_rule_evaluation_statuses",
				Description: "Information about the status of the rule evaluation.",
				Resolver:    fetchSagemakerTrainingJobProfilerRuleEvaluationStatuses,
				Columns: []schema.Column{
					{
						Name:        "training_job_cq_id",
						Description: "Unique CloudQuery ID of aws_sagemaker_training_jobs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "last_modified_time",
						Description: "Timestamp when the rule evaluation status was last modified.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "rule_configuration_name",
						Description: "The name of the rule configuration.",
						Type:        schema.TypeString,
					},
					{
						Name:        "rule_evaluation_job_arn",
						Description: "The Amazon Resource Name (ARN) of the rule evaluation job.",
						Type:        schema.TypeString,
					},
					{
						Name:        "rule_evaluation_status",
						Description: "Status of the rule evaluation.",
						Type:        schema.TypeString,
					},
					{
						Name:        "status_details",
						Description: "Details from the rule evaluation.",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchSagemakerTrainingJobs(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().SageMaker
	config := sagemaker.ListTrainingJobsInput{}
	for {
		response, err := svc.ListTrainingJobs(ctx, &config, func(options *sagemaker.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}

		// get more details about the notebook instance
		for _, n := range response.TrainingJobSummaries {

			config := sagemaker.DescribeTrainingJobInput{
				TrainingJobName: n.TrainingJobName,
			}
			response, err := svc.DescribeTrainingJob(ctx, &config, func(options *sagemaker.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return err
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

func fetchSagemakerTrainingJobAlgorithmSpecifications(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r, ok := parent.Item.(*sagemaker.DescribeTrainingJobOutput)
	if !ok {
		return fmt.Errorf("expected DescribeTrainingJobOutput but got %T", parent.Item)
	}
	if r.AlgorithmSpecification == nil {
		return nil
	}
	res <- r.AlgorithmSpecification
	return nil
}
func resolveSagemakerTrainingJobAlgorithmSpecificationsMetricDefinitions(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(*types.AlgorithmSpecification)
	if !ok {
		return fmt.Errorf("expected AlgorithmSpecification but got %T", resource.Item)
	}
	if len(r.MetricDefinitions) == 0 {
		return nil
	}

	var metricDefinitions = make([]map[string]interface{}, len(r.MetricDefinitions))

	for i, metric := range r.MetricDefinitions {
		metricDefinitions[i] = map[string]interface{}{
			"name":  aws.ToString(metric.Name),
			"regex": aws.ToString(metric.Regex),
		}
	}
	b, err := json.Marshal(metricDefinitions)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func fetchSagemakerTrainingJobDebugHookConfigs(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r, ok := parent.Item.(*sagemaker.DescribeTrainingJobOutput)
	if !ok {
		return fmt.Errorf("expected DescribeTrainingJobOutput but got %T", parent.Item)
	}
	if r.DebugHookConfig == nil {
		return nil
	}

	res <- r.DebugHookConfig
	return nil
}
func resolveSagemakerTrainingJobDebugHookConfigsCollectionConfigurations(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(*types.DebugHookConfig)
	if !ok {
		return fmt.Errorf("expected DebugHookConfig but got %T", resource.Item)
	}
	if len(r.CollectionConfigurations) == 0 {
		return nil
	}

	var collectionConfigurations = make([]map[string]interface{}, len(r.CollectionConfigurations))

	for i, config := range r.CollectionConfigurations {
		collectionConfigurations[i] = map[string]interface{}{
			"collection_name":       aws.ToString(config.CollectionName),
			"collection_parameters": config.CollectionParameters,
		}
	}
	b, err := json.Marshal(collectionConfigurations)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func fetchSagemakerTrainingJobDebugRuleConfigurations(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r, ok := parent.Item.(*sagemaker.DescribeTrainingJobOutput)
	if !ok {
		return fmt.Errorf("expected DescribeTrainingJobOutput but got %T", parent.Item)
	}
	res <- r.DebugRuleConfigurations
	return nil
}
func fetchSagemakerTrainingJobDebugRuleEvaluationStatuses(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r, ok := parent.Item.(*sagemaker.DescribeTrainingJobOutput)
	if !ok {
		return fmt.Errorf("expected DescribeTrainingJobOutput but got %T", parent.Item)
	}
	res <- r.DebugRuleEvaluationStatuses
	return nil
}
func fetchSagemakerTrainingJobInputDataConfigs(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r, ok := parent.Item.(*sagemaker.DescribeTrainingJobOutput)
	if !ok {
		return fmt.Errorf("expected DescribeTrainingJobOutput but got %T", parent.Item)
	}
	res <- r.InputDataConfig
	return nil
}
func fetchSagemakerTrainingJobProfilerRuleConfigurations(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r, ok := parent.Item.(*sagemaker.DescribeTrainingJobOutput)
	if !ok {
		return fmt.Errorf("expected DescribeTrainingJobOutput but got %T", parent.Item)
	}
	res <- r.ProfilerRuleConfigurations
	return nil
}
func fetchSagemakerTrainingJobProfilerRuleEvaluationStatuses(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r, ok := parent.Item.(*sagemaker.DescribeTrainingJobOutput)
	if !ok {
		return fmt.Errorf("expected DescribeTrainingJobOutput but got %T", parent.Item)
	}
	res <- r.ProfilerRuleEvaluationStatuses
	return nil
}
func resolveSagemakerTrainingJobCheckpointConfig(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(*sagemaker.DescribeTrainingJobOutput)
	if !ok {
		return fmt.Errorf("expected DescribeTrainingJobOutput but got %T", resource.Item)
	}
	if r.CheckpointConfig == nil {
		return nil
	}

	checkpointConfig := map[string]interface{}{
		"s3_uri":     aws.ToString(r.CheckpointConfig.S3Uri),
		"local_path": aws.ToString(r.CheckpointConfig.LocalPath),
	}
	return resource.Set(c.Name, checkpointConfig)
}
func resolveSagemakerTrainingJobExperimentConfig(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(*sagemaker.DescribeTrainingJobOutput)
	if !ok {
		return fmt.Errorf("expected DescribeTrainingJobOutput but got %T", resource.Item)
	} else if r.ExperimentConfig == nil {
		return nil
	}

	experimentConfig := map[string]interface{}{
		"experiment_name":              aws.ToString(r.ExperimentConfig.ExperimentName),
		"trial_component_display_name": aws.ToString(r.ExperimentConfig.TrialComponentDisplayName),
		"trial_name":                   aws.ToString(r.ExperimentConfig.TrialName),
	}
	return resource.Set(c.Name, experimentConfig)
}
func resolveSagemakerTrainingJobModelArtifacts(__ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(*sagemaker.DescribeTrainingJobOutput)
	if !ok {
		return fmt.Errorf("expected DescribeTrainingJobOutput but got %T", resource.Item)
	}
	if r.ModelArtifacts == nil {
		return nil
	}

	modelArtifacts := map[string]interface{}{
		"s3_model_artifacts": aws.ToString(r.ModelArtifacts.S3ModelArtifacts),
	}
	return resource.Set(c.Name, modelArtifacts)
}
func resolveSagemakerTrainingJobOutputDataConfig(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(*sagemaker.DescribeTrainingJobOutput)
	if !ok {
		return fmt.Errorf("expected DescribeTrainingJobOutput but got %T", resource.Item)
	}
	if r.OutputDataConfig == nil {
		return nil
	}

	outputDataConfig := map[string]interface{}{
		"s3_output_path": aws.ToString(r.OutputDataConfig.S3OutputPath),
		"kms_key_id":     aws.ToString(r.OutputDataConfig.KmsKeyId),
	}
	return resource.Set(c.Name, outputDataConfig)
}
func resolveSagemakerTrainingJobProfilerConfig(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(*sagemaker.DescribeTrainingJobOutput)
	if !ok {
		return fmt.Errorf("expected DescribeTrainingJobOutput but got %T", resource.Item)
	}
	if r.ProfilerConfig == nil {
		return nil
	}

	profilerConfig := map[string]interface{}{
		"s3_output_path":           aws.ToString(r.ProfilerConfig.S3OutputPath),
		"profiling_interval_in_ms": aws.ToInt64(r.ProfilerConfig.ProfilingIntervalInMilliseconds),
		"profiling_parameters":     r.ProfilerConfig.ProfilingParameters,
	}
	return resource.Set(c.Name, profilerConfig)
}
func resolveSagemakerTrainingJobResourceConfig(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(*sagemaker.DescribeTrainingJobOutput)
	if !ok {
		return fmt.Errorf("expected DescribeTrainingJobOutput but got %T", resource.Item)
	}
	if r.ResourceConfig == nil {
		return nil
	}

	resourceConfig := map[string]interface{}{
		"instance_count":    r.ResourceConfig.InstanceCount,
		"instance_type":     r.ResourceConfig.InstanceType,
		"volume_size_in_gb": r.ResourceConfig.VolumeSizeInGB,
		"volume_kms_key_id": aws.ToString(r.ResourceConfig.VolumeKmsKeyId),
	}
	return resource.Set(c.Name, resourceConfig)
}
func resolveSagemakerTrainingJobStoppingCondition(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(*sagemaker.DescribeTrainingJobOutput)
	if !ok {
		return fmt.Errorf("expected DescribeTrainingJobOutput but got %T", resource.Item)
	}
	if r.StoppingCondition == nil {
		return nil
	}

	stoppingCondition := map[string]interface{}{
		"max_runtime_in_seconds":   r.StoppingCondition.MaxRuntimeInSeconds,
		"max_wait_time_in_seconds": r.StoppingCondition.MaxWaitTimeInSeconds,
	}
	return resource.Set(c.Name, stoppingCondition)
}
func resolveSagemakerTrainingJobTensorBoardOutputConfig(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(*sagemaker.DescribeTrainingJobOutput)
	if !ok {
		return fmt.Errorf("expected DescribeTrainingJobOutput but got %T", resource.Item)
	}
	if r.TensorBoardOutputConfig == nil {
		return nil
	}

	tensorBoardOutputConfig := map[string]interface{}{
		"s3_output_path": r.TensorBoardOutputConfig.S3OutputPath,
		"local_path":     r.TensorBoardOutputConfig.LocalPath,
	}
	return resource.Set(c.Name, tensorBoardOutputConfig)
}
func resolveSagemakerTrainingJobVpcConfig(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(*sagemaker.DescribeTrainingJobOutput)
	if !ok {
		return fmt.Errorf("expected DescribeTrainingJobOutput but got %T", resource.Item)
	}
	if r.VpcConfig == nil {
		return nil
	}

	vpcConfig := map[string]interface{}{
		"subnets":            r.VpcConfig.Subnets,
		"security_group_ids": r.VpcConfig.SecurityGroupIds,
	}
	return resource.Set(c.Name, vpcConfig)
}
func resolveSagemakerTrainingJobTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r, ok := resource.Item.(*sagemaker.DescribeTrainingJobOutput)
	if !ok {
		return fmt.Errorf("expected DescribeTrainingJobOutput but got %T", resource.Item)
	}
	if r == nil {
		return nil
	}

	c := meta.(*client.Client)
	svc := c.Services().SageMaker
	config := sagemaker.ListTagsInput{
		ResourceArn: r.TrainingJobArn,
	}
	response, err := svc.ListTags(ctx, &config, func(options *sagemaker.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}

	tags := make(map[string]*string, len(response.Tags))
	for _, t := range response.Tags {
		tags[*t.Key] = t.Value
	}

	return resource.Set("tags", tags)
}
func resolveSagemakerTrainingJobSecondaryStatusTransitions(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(*sagemaker.DescribeTrainingJobOutput)
	if !ok {
		return fmt.Errorf("expected DescribeEndpointConfigOutput but got %T", resource.Item)
	}
	if len(r.SecondaryStatusTransitions) == 0 {
		return nil
	}

	var secondaryStatusTransitions = make([]map[string]interface{}, len(r.SecondaryStatusTransitions))

	for i, status := range r.SecondaryStatusTransitions {
		secondaryStatusTransitions[i] = map[string]interface{}{
			"start_time":     status.StartTime,
			"end_time":       status.EndTime,
			"status":         status.Status,
			"status_message": status.StatusMessage,
		}
	}
	b, err := json.Marshal(secondaryStatusTransitions)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveSagemakerTrainingJobFinalMetricDataList(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(*sagemaker.DescribeTrainingJobOutput)
	if !ok {
		return fmt.Errorf("expected DescribeEndpointConfigOutput but got %T", resource.Item)
	}
	if len(r.FinalMetricDataList) == 0 {
		return nil
	}

	var finalMetricDataList = make([]map[string]interface{}, len(r.FinalMetricDataList))
	for i, config := range r.FinalMetricDataList {
		finalMetricDataList[i] = map[string]interface{}{
			"metric_name": config.MetricName,
			"value":       config.Value,
			"timestamp":   config.Timestamp,
		}
	}
	b, err := json.Marshal(finalMetricDataList)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
