package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IotJobs() *schema.Table {
	return &schema.Table{
		Name:         "aws_iot_jobs",
		Description:  "The Job object contains details about a job.",
		Resolver:     fetchIotJobs,
		Multiplex:    client.ServiceAccountRegionMultiplexer("iot"),
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
				Name:        "tags",
				Description: "Tags of the resource",
				Type:        schema.TypeJSON,
				Resolver:    ResolveIotJobTags,
			},
			{
				Name:        "comment",
				Description: "If the job was updated, describes the reason for the update.",
				Type:        schema.TypeString,
			},
			{
				Name:        "completed_at",
				Description: "The time, in seconds since the epoch, when the job was completed.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "created_at",
				Description: "The time, in seconds since the epoch, when the job was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "description",
				Description: "A short text description of the job.",
				Type:        schema.TypeString,
			},
			{
				Name:        "document_parameters",
				Description: "A key-value map that pairs the patterns that need to be replaced in a managed template job document schema",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "force_canceled",
				Description: "Will be true if the job was canceled with the optional force parameter set to true.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "arn",
				Description: "An ARN identifying the job with format \"arn:aws:iot:region:account:job/jobId\".",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("JobArn"),
			},
			{
				Name:        "base_rate_per_minute",
				Description: "The minimum number of things that will be notified of a pending job, per minute at the start of job rollout",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("JobExecutionsRolloutConfig.ExponentialRate.BaseRatePerMinute"),
			},
			{
				Name:        "increment_factor",
				Description: "The exponential factor to increase the rate of rollout for a job",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("JobExecutionsRolloutConfig.ExponentialRate.IncrementFactor"),
			},
			{
				Name:        "rollout_config_rate_increase_criteria_number_of_notified_things",
				Description: "The threshold for number of notified things that will initiate the increase in rate of rollout.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("JobExecutionsRolloutConfig.ExponentialRate.RateIncreaseCriteria.NumberOfNotifiedThings"),
			},
			{
				Name:        "rate_increase_criteria_number_of_succeeded_things",
				Description: "The threshold for number of succeeded things that will initiate the increase in rate of rollout.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("JobExecutionsRolloutConfig.ExponentialRate.RateIncreaseCriteria.NumberOfSucceededThings"),
			},
			{
				Name:        "maximum_per_minute",
				Description: "The maximum number of things that will be notified of a pending job, per minute. This parameter allows you to create a staged rollout.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("JobExecutionsRolloutConfig.MaximumPerMinute"),
			},
			{
				Name:        "id",
				Description: "The unique identifier you assigned to this job when it was created.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("JobId"),
			},
			{
				Name:        "process_details_number_of_canceled_things",
				Description: "The number of things that cancelled the job.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("JobProcessDetails.NumberOfCanceledThings"),
			},
			{
				Name:        "process_details_number_of_failed_things",
				Description: "The number of things that failed executing the job.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("JobProcessDetails.NumberOfFailedThings"),
			},
			{
				Name:        "process_details_number_of_in_progress_things",
				Description: "The number of things currently executing the job.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("JobProcessDetails.NumberOfInProgressThings"),
			},
			{
				Name:        "process_details_number_of_queued_things",
				Description: "The number of things that are awaiting execution of the job.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("JobProcessDetails.NumberOfQueuedThings"),
			},
			{
				Name:        "process_details_number_of_rejected_things",
				Description: "The number of things that rejected the job.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("JobProcessDetails.NumberOfRejectedThings"),
			},
			{
				Name:        "process_details_number_of_removed_things",
				Description: "The number of things that are no longer scheduled to execute the job because they have been deleted or have been removed from the group that was a target of the job.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("JobProcessDetails.NumberOfRemovedThings"),
			},
			{
				Name:        "process_details_number_of_succeeded_things",
				Description: "The number of things which successfully completed the job.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("JobProcessDetails.NumberOfSucceededThings"),
			},
			{
				Name:        "process_details_number_of_timed_out_things",
				Description: "The number of things whose job execution status is TIMED_OUT.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("JobProcessDetails.NumberOfTimedOutThings"),
			},
			{
				Name:        "process_details_processing_targets",
				Description: "The target devices to which the job execution is being rolled out",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("JobProcessDetails.ProcessingTargets"),
			},
			{
				Name:        "job_template_arn",
				Description: "The ARN of the job template used to create the job.",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_updated_at",
				Description: "The time, in seconds since the epoch, when the job was last updated.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "namespace_id",
				Description: "The namespace used to indicate that a job is a customer-managed job",
				Type:        schema.TypeString,
			},
			{
				Name:        "presigned_url_config_expires_in_sec",
				Description: "How long (in seconds) pre-signed URLs are valid",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("PresignedUrlConfig.ExpiresInSec"),
			},
			{
				Name:        "presigned_url_config_role_arn",
				Description: "The ARN of an IAM role that grants grants permission to download files from the S3 bucket where the job data/updates are stored",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PresignedUrlConfig.RoleArn"),
			},
			{
				Name:        "reason_code",
				Description: "If the job was updated, provides the reason code for the update.",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The status of the job, one of IN_PROGRESS, CANCELED, DELETION_IN_PROGRESS or COMPLETED.",
				Type:        schema.TypeString,
			},
			{
				Name:        "target_selection",
				Description: "Specifies whether the job will continue to run (CONTINUOUS), or will be complete after all those things specified as targets have completed the job (SNAPSHOT). If continuous, the job may also be run on a thing when a change is detected in a target",
				Type:        schema.TypeString,
			},
			{
				Name:        "targets",
				Description: "A list of IoT things and thing groups to which the job should be sent.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "timeout_config_in_progress_timeout_in_minutes",
				Description: "Specifies the amount of time, in minutes, this device has to finish execution of this job",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("TimeoutConfig.InProgressTimeoutInMinutes"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_iot_job_abort_config_criteria_list",
				Description: "The criteria that determine when and how a job abort takes place.",
				Resolver:    fetchIotJobAbortConfigCriteriaLists,
				Columns: []schema.Column{
					{
						Name:        "job_cq_id",
						Description: "Unique CloudQuery ID of aws_iot_jobs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "action",
						Description: "The type of job action to take to initiate the job abort.",
						Type:        schema.TypeString,
					},
					{
						Name:        "failure_type",
						Description: "The type of job execution failures that can initiate a job abort.",
						Type:        schema.TypeString,
					},
					{
						Name:        "min_number_of_executed_things",
						Description: "The minimum number of things which must receive job execution notifications before the job can be aborted.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "threshold_percentage",
						Description: "The minimum percentage of job execution failures that must occur to initiate the job abort",
						Type:        schema.TypeFloat,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchIotJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	client := meta.(*client.Client)
	svc := client.Services().IOT
	input := iot.ListJobsInput{
		MaxResults: aws.Int32(250),
	}

	for {
		response, err := svc.ListJobs(ctx, &input, func(options *iot.Options) {
			options.Region = client.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}

		for _, s := range response.Jobs {
			job, err := svc.DescribeJob(ctx, &iot.DescribeJobInput{
				JobId: s.JobId,
			}, func(options *iot.Options) {
				options.Region = client.Region
			})
			if err != nil {
				return diag.WrapError(err)
			}
			res <- job.Job
		}

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
func ResolveIotJobTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*types.Job)
	cl := meta.(*client.Client)
	svc := cl.Services().IOT
	input := iot.ListTagsForResourceInput{
		ResourceArn: i.JobArn,
	}
	tags := make(map[string]string)

	for {
		response, err := svc.ListTagsForResource(ctx, &input, func(options *iot.Options) {
			options.Region = cl.Region
		})

		if err != nil {
			return diag.WrapError(err)
		}

		client.TagsIntoMap(response.Tags, tags)

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return diag.WrapError(resource.Set(c.Name, tags))
}
func fetchIotJobAbortConfigCriteriaLists(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	i := parent.Item.(*types.Job)
	if i.AbortConfig == nil {
		return nil
	}
	res <- i.AbortConfig.CriteriaList
	return nil
}
