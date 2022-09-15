// Code generated by codegen; DO NOT EDIT.

package glue

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func JobRuns() *schema.Table {
	return &schema.Table{
		Name:      "aws_glue_job_runs",
		Resolver:  fetchGlueJobRuns,
		Multiplex: client.ServiceAccountRegionMultiplexer("glue"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "job_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("arn"),
			},
			{
				Name:     "allocated_capacity",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("AllocatedCapacity"),
			},
			{
				Name:     "arguments",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Arguments"),
			},
			{
				Name:     "attempt",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Attempt"),
			},
			{
				Name:     "completed_on",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CompletedOn"),
			},
			{
				Name:     "dpu_seconds",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("DPUSeconds"),
			},
			{
				Name:     "error_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ErrorMessage"),
			},
			{
				Name:     "execution_time",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ExecutionTime"),
			},
			{
				Name:     "glue_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GlueVersion"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "job_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("JobName"),
			},
			{
				Name:     "job_run_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("JobRunState"),
			},
			{
				Name:     "last_modified_on",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastModifiedOn"),
			},
			{
				Name:     "log_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LogGroupName"),
			},
			{
				Name:     "max_capacity",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("MaxCapacity"),
			},
			{
				Name:     "notification_property",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NotificationProperty"),
			},
			{
				Name:     "number_of_workers",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumberOfWorkers"),
			},
			{
				Name:     "predecessor_runs",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PredecessorRuns"),
			},
			{
				Name:     "previous_run_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PreviousRunId"),
			},
			{
				Name:     "security_configuration",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecurityConfiguration"),
			},
			{
				Name:     "started_on",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("StartedOn"),
			},
			{
				Name:     "timeout",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Timeout"),
			},
			{
				Name:     "trigger_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TriggerName"),
			},
			{
				Name:     "worker_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WorkerType"),
			},
		},
	}
}
