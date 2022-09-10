package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Jobs() *schema.Table {
	return &schema.Table{
		Name:        "aws_glue_jobs",
		Description: "Specifies a job definition",
		Resolver:    fetchGlueJobs,
		Multiplex:   client.ServiceAccountRegionMultiplexer("glue"),
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
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) of the workflow.",
				Type:            schema.TypeString,
				Resolver:        resolveGlueJobArn,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
				Resolver:    resolveGlueJobTags,
			},
			{
				Name:        "allocated_capacity",
				Description: "This field is deprecated",
				Type:        schema.TypeInt,
			},
			{
				Name:        "code_gen_configuration_nodes",
				Description: "The representation of a directed acyclic graph on which both the Glue Studio visual component and Glue Studio code generation is based",
				Type:        schema.TypeJSON,
			},
			{
				Name:     "command",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Command"),
			},
			{
				Name:     "connections",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Connections"),
			},
			{
				Name:        "created_on",
				Description: "The time and date that this job definition was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "default_arguments",
				Description: "The default arguments for this job, specified as name-value pairs",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "description",
				Description: "A description of the job",
				Type:        schema.TypeString,
			},
			{
				Name:     "execution_property",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ExecutionProperty"),
			},
			{
				Name:        "glue_version",
				Description: "Glue version determines the versions of Apache Spark and Python that Glue supports",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_modified_on",
				Description: "The last point in time when this job definition was modified",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "log_uri",
				Description: "This field is reserved for future use",
				Type:        schema.TypeString,
			},
			{
				Name:        "max_capacity",
				Description: "For Glue version 10 or earlier jobs, using the standard worker type, the number of Glue data processing units (DPUs) that can be allocated when this job runs",
				Type:        schema.TypeFloat,
			},
			{
				Name:        "max_retries",
				Description: "The maximum number of times to retry this job after a JobRun fails",
				Type:        schema.TypeInt,
			},
			{
				Name:        "name",
				Description: "The name you assign to this job definition",
				Type:        schema.TypeString,
			},
			{
				Name:        "non_overridable_arguments",
				Description: "Non-overridable arguments for this job, specified as name-value pairs",
				Type:        schema.TypeJSON,
			},
			{
				Name:     "notification_property",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NotificationProperty"),
			},
			{
				Name:        "number_of_workers",
				Description: "The number of workers of a defined workerType that are allocated when a job runs",
				Type:        schema.TypeInt,
			},
			{
				Name:        "role",
				Description: "The name or Amazon Resource Name (ARN) of the IAM role associated with this job",
				Type:        schema.TypeString,
			},
			{
				Name:        "security_configuration",
				Description: "The name of the SecurityConfiguration structure to be used with this job",
				Type:        schema.TypeString,
			},
			{
				Name:        "timeout",
				Description: "The job timeout in minutes",
				Type:        schema.TypeInt,
			},
			{
				Name:        "worker_type",
				Description: "The type of predefined worker that is allocated when a job runs",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_glue_job_runs",
				Description: "Contains information about a job run",
				Resolver:    fetchGlueJobRuns,
				Columns: []schema.Column{
					{
						Name:        "job_cq_id",
						Description: "Unique CloudQuery ID of aws_glue_jobs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "allocated_capacity",
						Description: "This field is deprecated",
						Type:        schema.TypeInt,
					},
					{
						Name:        "arguments",
						Description: "The job arguments associated with this run",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "attempt",
						Description: "The number of the attempt to run this job",
						Type:        schema.TypeInt,
					},
					{
						Name:        "completed_on",
						Description: "The date and time that this job run completed",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "dpu_seconds",
						Description: "This field populates only for Auto Scaling job runs, and represents the total time each executor ran during the lifecycle of a job run in seconds, multiplied by a DPU factor (1 for G1X, 2 for G2X, or 025 for G025X workers)",
						Type:        schema.TypeFloat,
						Resolver:    schema.PathResolver("DPUSeconds"),
					},
					{
						Name:        "error_message",
						Description: "An error message associated with this job run",
						Type:        schema.TypeString,
					},
					{
						Name:        "execution_time",
						Description: "The amount of time (in seconds) that the job run consumed resources",
						Type:        schema.TypeInt,
					},
					{
						Name:        "glue_version",
						Description: "Glue version determines the versions of Apache Spark and Python that Glue supports",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "The ID of this job run",
						Type:        schema.TypeString,
					},
					{
						Name:        "job_name",
						Description: "The name of the job definition being used in this run",
						Type:        schema.TypeString,
					},
					{
						Name:        "job_run_state",
						Description: "The current state of the job run",
						Type:        schema.TypeString,
					},
					{
						Name:        "last_modified_on",
						Description: "The last time that this job run was modified",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "log_group_name",
						Description: "The name of the log group for secure logging that can be server-side encrypted in Amazon CloudWatch using KMS",
						Type:        schema.TypeString,
					},
					{
						Name:        "max_capacity",
						Description: "The number of Glue data processing units (DPUs) that can be allocated when this job runs",
						Type:        schema.TypeFloat,
					},
					{
						Name:     "notification_property",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("NotificationProperty"),
					},
					{
						Name:        "number_of_workers",
						Description: "The number of workers of a defined workerType that are allocated when a job runs",
						Type:        schema.TypeInt,
					},
					{
						Name:        "predecessor_runs",
						Description: "A list of predecessors to this job run",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "previous_run_id",
						Description: "The ID of the previous run of this job",
						Type:        schema.TypeString,
					},
					{
						Name:        "security_configuration",
						Description: "The name of the SecurityConfiguration structure to be used with this job run",
						Type:        schema.TypeString,
					},
					{
						Name:        "started_on",
						Description: "The date and time at which this job run was started",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "timeout",
						Description: "The JobRun timeout in minutes",
						Type:        schema.TypeInt,
					},
					{
						Name:        "trigger_name",
						Description: "The name of the trigger that started this job run",
						Type:        schema.TypeString,
					},
					{
						Name:        "worker_type",
						Description: "The type of predefined worker that is allocated when a job runs",
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

func fetchGlueJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.GetJobsInput{}
	for {
		result, err := svc.GetJobs(ctx, &input)
		if err != nil {
			return err
		}
		res <- result.Jobs
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}
func resolveGlueJobArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	arn := aws.String(jobARN(cl, aws.ToString(resource.Item.(types.Job).Name)))
	return resource.Set(c.Name, arn)
}
func resolveGlueJobTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	result, err := svc.GetTags(ctx, &glue.GetTagsInput{
		ResourceArn: aws.String(jobARN(cl, aws.ToString(resource.Item.(types.Job).Name))),
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, result.Tags)
}
func fetchGlueJobRuns(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.GetJobRunsInput{
		JobName: parent.Item.(types.Job).Name,
	}
	for {
		result, err := svc.GetJobRuns(ctx, &input)
		if err != nil {
			return err
		}
		res <- result.JobRuns
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

func jobARN(cl *client.Client, name string) string {
	return cl.ARN(client.GlueService, "job", name)
}
