// Code generated by codegen; DO NOT EDIT.

package glue

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Workflows() *schema.Table {
	return &schema.Table{
		Name:                "aws_glue_workflows",
		Resolver:            fetchGlueWorkflows,
		PreResourceResolver: getWorkflow,
		Multiplex:           client.ServiceAccountRegionMultiplexer("glue"),
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
				Resolver: resolveGlueWorkflowArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveGlueWorkflowTags,
			},
			{
				Name:     "blueprint_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BlueprintDetails"),
			},
			{
				Name:     "created_on",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedOn"),
			},
			{
				Name:     "default_run_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DefaultRunProperties"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "graph",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Graph"),
			},
			{
				Name:     "last_modified_on",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastModifiedOn"),
			},
			{
				Name:     "last_run",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LastRun"),
			},
			{
				Name:     "max_concurrent_runs",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxConcurrentRuns"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
		},
	}
}
