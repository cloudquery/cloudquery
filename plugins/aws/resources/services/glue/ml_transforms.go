package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource ml_transforms --config ml_transforms.hcl --output .
func MlTransforms() *schema.Table {
	return &schema.Table{
		Name:         "aws_glue_ml_transforms",
		Description:  "A structure for a machine learning transform",
		Resolver:     fetchGlueMlTransforms,
		Multiplex:    client.ServiceAccountRegionMultiplexer("glue"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
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
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the workflow.",
				Type:        schema.TypeString,
				Resolver:    resolveGlueMlTransformArn,
			},
			{
				Name:        "tags",
				Description: "Resource tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveGlueMlTransformTags,
			},
			{
				Name:        "created_on",
				Description: "A timestamp",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "description",
				Description: "A user-defined, long-form description text for the machine learning transform Descriptions are not guaranteed to be unique and can be changed at any time",
				Type:        schema.TypeString,
			},
			{
				Name:        "evaluation_metrics_transform_type",
				Description: "The type of machine learning transform",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EvaluationMetrics.TransformType"),
			},
			{
				Name:        "evaluation_metrics_find_matches_metrics_area_under_pr_curve",
				Description: "The area under the precision/recall curve (AUPRC) is a single number measuring the overall quality of the transform, that is independent of the choice made for precision vs",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("EvaluationMetrics.FindMatchesMetrics.AreaUnderPRCurve"),
			},
			{
				Name:        "evaluation_metrics_find_matches_metrics_column_importances",
				Description: "A list of ColumnImportance structures containing column importance metrics, sorted in order of descending importance",
				Type:        schema.TypeJSON,
				Resolver:    resolveMlTransformsEvaluationMetricsFindMatchesMetricsColumnImportances,
			},
			{
				Name:        "evaluation_metrics_find_matches_metrics_confusion_matrix",
				Description: "The confusion matrix shows you what your transform is predicting accurately and what types of errors it is making",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("EvaluationMetrics.FindMatchesMetrics.ConfusionMatrix"),
			},
			{
				Name:        "evaluation_metrics_find_matches_metrics_f1",
				Description: "The maximum F1 metric indicates the transform's accuracy between 0 and 1, where 1 is the best accuracy",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("EvaluationMetrics.FindMatchesMetrics.F1"),
			},
			{
				Name:        "evaluation_metrics_find_matches_metrics_precision",
				Description: "The precision metric indicates when often your transform is correct when it predicts a match",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("EvaluationMetrics.FindMatchesMetrics.Precision"),
			},
			{
				Name:        "evaluation_metrics_find_matches_metrics_recall",
				Description: "The recall metric indicates that for an actual match, how often your transform predicts the match",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("EvaluationMetrics.FindMatchesMetrics.Recall"),
			},
			{
				Name:        "glue_version",
				Description: "This value determines which version of Glue this machine learning transform is compatible with",
				Type:        schema.TypeString,
			},
			{
				Name:        "label_count",
				Description: "A count identifier for the labeling files generated by Glue for this transform As you create a better transform, you can iteratively download, label, and upload the labeling file",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "last_modified_on",
				Description: "A timestamp",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "max_capacity",
				Description: "The number of Glue data processing units (DPUs) that are allocated to task runs for this transform",
				Type:        schema.TypeFloat,
			},
			{
				Name:        "max_retries",
				Description: "The maximum number of times to retry after an MLTaskRun of the machine learning transform fails",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "name",
				Description: "A user-defined name for the machine learning transform",
				Type:        schema.TypeString,
			},
			{
				Name:        "number_of_workers",
				Description: "The number of workers of a defined workerType that are allocated when a task of the transform runs",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "parameters_transform_type",
				Description: "The type of machine learning transform",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Parameters.TransformType"),
			},
			{
				Name:        "parameters_find_matches_parameters_accuracy_cost_tradeoff",
				Description: "The value that is selected when tuning your transform for a balance between accuracy and cost",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("Parameters.FindMatchesParameters.AccuracyCostTradeoff"),
			},
			{
				Name:        "parameters_find_matches_parameters_enforce_provided_labels",
				Description: "The value to switch on or off to force the output to match the provided labels from users",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Parameters.FindMatchesParameters.EnforceProvidedLabels"),
			},
			{
				Name:        "parameters_find_matches_parameters_precision_recall_tradeoff",
				Description: "The value selected when tuning your transform for a balance between precision and recall",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("Parameters.FindMatchesParameters.PrecisionRecallTradeoff"),
			},
			{
				Name:        "parameters_find_matches_parameters_primary_key_column_name",
				Description: "The name of a column that uniquely identifies rows in the source table",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Parameters.FindMatchesParameters.PrimaryKeyColumnName"),
			},
			{
				Name:        "role",
				Description: "The name or Amazon Resource Name (ARN) of the IAM role with the required permissions",
				Type:        schema.TypeString,
			},
			{
				Name:        "schema",
				Description: "A map of key-value pairs representing the columns and data types that this transform can run against",
				Type:        schema.TypeJSON,
				Resolver:    resolveMlTransformsSchema,
			},
			{
				Name:        "status",
				Description: "The current status of the machine learning transform",
				Type:        schema.TypeString,
			},
			{
				Name:        "timeout",
				Description: "The timeout in minutes of the machine learning transform",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "transform_encryption_user_data_encryption_mode",
				Description: "The encryption mode applied to user data",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TransformEncryption.MlUserDataEncryption.MlUserDataEncryptionMode"),
			},
			{
				Name:        "transform_encryption_ml_user_data_encryption_kms_key_id",
				Description: "The ID for the customer-provided KMS key",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TransformEncryption.MlUserDataEncryption.KmsKeyId"),
			},
			{
				Name:        "transform_encryption_task_run_security_configuration_name",
				Description: "The name of the security configuration",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TransformEncryption.TaskRunSecurityConfigurationName"),
			},
			{
				Name:        "id",
				Description: "The unique transform ID that is generated for the machine learning transform The ID is guaranteed to be unique and does not change",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TransformId"),
			},
			{
				Name:        "worker_type",
				Description: "The type of predefined worker that is allocated when a task of this transform runs",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_glue_ml_transform_input_record_tables",
				Description: "The database and table in the Glue Data Catalog that is used for input or output data",
				Resolver:    schema.PathTableResolver("InputRecordTables"),
				Columns: []schema.Column{
					{
						Name:        "ml_transform_cq_id",
						Description: "Unique CloudQuery ID of aws_glue_ml_transforms table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "database_name",
						Description: "A database name in the Glue Data Catalog",
						Type:        schema.TypeString,
					},
					{
						Name:        "table_name",
						Description: "A table name in the Glue Data Catalog",
						Type:        schema.TypeString,
					},
					{
						Name:        "catalog_id",
						Description: "A unique identifier for the Glue Data Catalog",
						Type:        schema.TypeString,
					},
					{
						Name:        "connection_name",
						Description: "The name of the connection to the Glue Data Catalog",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_glue_ml_transform_task_runs",
				Description: "The sampling parameters that are associated with the machine learning transform",
				Resolver:    fetchGlueMlTransformTaskRuns,
				Columns: []schema.Column{
					{
						Name:        "ml_transform_cq_id",
						Description: "Unique CloudQuery ID of aws_glue_ml_transforms table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "completed_on",
						Description: "The last point in time that the requested task run was completed",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "error_string",
						Description: "The list of error strings associated with this task run",
						Type:        schema.TypeString,
					},
					{
						Name:        "execution_time",
						Description: "The amount of time (in seconds) that the task run consumed resources",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "last_modified_on",
						Description: "The last point in time that the requested task run was updated",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "log_group_name",
						Description: "The names of the log group for secure logging, associated with this task run",
						Type:        schema.TypeString,
					},
					{
						Name:        "export_labels_task_run_properties_output_s3_path",
						Description: "The Amazon Simple Storage Service (Amazon S3) path where you will export the labels",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.ExportLabelsTaskRunProperties.OutputS3Path"),
					},
					{
						Name:        "find_matches_task_run_properties_job_id",
						Description: "The job ID for the Find Matches task run",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.FindMatchesTaskRunProperties.JobId"),
					},
					{
						Name:        "find_matches_task_run_properties_job_name",
						Description: "The name assigned to the job for the Find Matches task run",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.FindMatchesTaskRunProperties.JobName"),
					},
					{
						Name:        "find_matches_task_run_properties_job_run_id",
						Description: "The job run ID for the Find Matches task run",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.FindMatchesTaskRunProperties.JobRunId"),
					},
					{
						Name:        "import_labels_task_run_properties_input_s3_path",
						Description: "The Amazon Simple Storage Service (Amazon S3) path from where you will import the labels",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.ImportLabelsTaskRunProperties.InputS3Path"),
					},
					{
						Name:        "import_labels_task_run_properties_replace",
						Description: "Indicates whether to overwrite your existing labels",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Properties.ImportLabelsTaskRunProperties.Replace"),
					},
					{
						Name:        "labeling_set_generation_task_run_properties_output_s3_path",
						Description: "The Amazon Simple Storage Service (Amazon S3) path where you will generate the labeling set",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.LabelingSetGenerationTaskRunProperties.OutputS3Path"),
					},
					{
						Name:        "task_type",
						Description: "The type of task run",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.TaskType"),
					},
					{
						Name:        "started_on",
						Description: "The date and time that this task run started",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "status",
						Description: "The current status of the requested task run",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "The unique identifier for this task run",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TaskRunId"),
					},
					{
						Name:        "transform_id",
						Description: "The unique identifier for the transform",
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

func fetchGlueMlTransforms(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.GetMLTransformsInput{}
	for {
		result, err := svc.GetMLTransforms(ctx, &input)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- result.Transforms
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}
func resolveGlueMlTransformArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	r := resource.Item.(types.MLTransform)
	arn := aws.String(mlTransformARN(cl, &r))
	return diag.WrapError(resource.Set(c.Name, arn))
}
func resolveGlueMlTransformTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	r := resource.Item.(types.MLTransform)
	result, err := svc.GetTags(ctx, &glue.GetTagsInput{
		ResourceArn: aws.String(mlTransformARN(cl, &r)),
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, result.Tags))
}
func resolveMlTransformsEvaluationMetricsFindMatchesMetricsColumnImportances(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.MLTransform)
	j := make(map[string]float64)
	if r.EvaluationMetrics == nil || r.EvaluationMetrics.FindMatchesMetrics == nil {
		return nil
	}
	for _, c := range r.EvaluationMetrics.FindMatchesMetrics.ColumnImportances {
		j[*c.ColumnName] = *c.Importance
	}
	return diag.WrapError(resource.Set(c.Name, j))
}
func resolveMlTransformsSchema(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.MLTransform)
	j := make(map[string]string)
	for _, c := range r.Schema {
		j[*c.Name] = *c.DataType
	}
	return diag.WrapError(resource.Set(c.Name, j))
}
func fetchGlueMlTransformTaskRuns(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.MLTransform)
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.GetMLTaskRunsInput{
		TransformId: r.TransformId,
	}
	for {
		result, err := svc.GetMLTaskRuns(ctx, &input)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- result.TaskRuns
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func mlTransformARN(cl *client.Client, tr *types.MLTransform) string {
	return cl.ARN(client.GlueService, "mlTransform", *tr.TransformId)
}
