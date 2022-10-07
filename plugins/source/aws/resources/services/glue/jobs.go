// Code generated by codegen; DO NOT EDIT.

package glue

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Jobs() *schema.Table {
	return &schema.Table{
		Name:      "aws_glue_jobs",
		Resolver:  fetchGlueJobs,
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveGlueJobArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveGlueJobTags,
			},
			{
				Name:     "allocated_capacity",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("AllocatedCapacity"),
			},
			{
				Name:     "code_gen_configuration_nodes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CodeGenConfigurationNodes"),
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
				Name:     "created_on",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedOn"),
			},
			{
				Name:     "default_arguments",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DefaultArguments"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "execution_class",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ExecutionClass"),
			},
			{
				Name:     "execution_property",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ExecutionProperty"),
			},
			{
				Name:     "glue_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GlueVersion"),
			},
			{
				Name:     "last_modified_on",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastModifiedOn"),
			},
			{
				Name:     "log_uri",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LogUri"),
			},
			{
				Name:     "max_capacity",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("MaxCapacity"),
			},
			{
				Name:     "max_retries",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxRetries"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "non_overridable_arguments",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NonOverridableArguments"),
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
				Name:     "role",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Role"),
			},
			{
				Name:     "security_configuration",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecurityConfiguration"),
			},
			{
				Name:     "source_control_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SourceControlDetails"),
			},
			{
				Name:     "timeout",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Timeout"),
			},
			{
				Name:     "worker_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WorkerType"),
			},
		},

		Relations: []*schema.Table{
			JobRuns(),
		},
	}
}
