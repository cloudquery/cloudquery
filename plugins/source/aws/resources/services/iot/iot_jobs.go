package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func IotJobs() *schema.Table {
	return &schema.Table{
		Name:        "aws_iot_jobs",
		Description: "The Job object contains details about a job.",
		Resolver:    fetchIotJobs,
		Multiplex:   client.ServiceAccountRegionMultiplexer("iot"),

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
				Name:            "arn",
				Description:     "An ARN identifying the job with format \"arn:aws:iot:region:account:job/jobId\".",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("JobArn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:     "job_executions_rollout_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("JobExecutionsRolloutConfig"),
			},
			{
				Name:        "id",
				Description: "The unique identifier you assigned to this job when it was created.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("JobId"),
			},
			{
				Name:     "job_process_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("JobProcessDetails"),
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
				Name:        "presigned_url_config",
				Type: 			schema.TypeJSON,
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
				Name:     "timeout_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TimeoutConfig"),
			},
			{
				Name: "abort_config",
				Type: schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchIotJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().IOT
	input := iot.ListJobsInput{
		MaxResults: aws.Int32(250),
	}

	for {
		response, err := svc.ListJobs(ctx, &input)
		if err != nil {
			return err
		}

		for _, s := range response.Jobs {
			job, err := svc.DescribeJob(ctx, &iot.DescribeJobInput{
				JobId: s.JobId,
			}, func(options *iot.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				return err
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
		response, err := svc.ListTagsForResource(ctx, &input)

		if err != nil {
			return err
		}

		client.TagsIntoMap(response.Tags, tags)

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return resource.Set(c.Name, tags)
}
