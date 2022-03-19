package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource stacks --config gen.hcl --output .
func Stacks() *schema.Table {
	return &schema.Table{
		Name:         "aws_cloudformation_stacks",
		Description:  "The Stack data type.",
		Resolver:     fetchCloudformationStacks,
		Multiplex:    client.ServiceAccountRegionMultiplexer("cloudformation"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StackId"),
			},
			{
				Name:        "creation_time",
				Description: "The time at which the stack was created.  This member is required.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "stack",
				Description: "The name associated with the stack.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StackName"),
			},
			{
				Name:        "status",
				Description: "Current status of the stack.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StackStatus"),
			},
			{
				Name:          "capabilities",
				Description:   "The capabilities allowed in the stack.",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:          "change_set_id",
				Description:   "The unique ID of the change set.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "deletion_time",
				Description:   "The time the stack was deleted.",
				Type:          schema.TypeTimestamp,
				IgnoreInTests: true,
			},
			{
				Name:          "description",
				Description:   "A user-defined description associated with the stack.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "disable_rollback",
				Description: "Boolean to enable or disable rollback on stack creation failures:  * true: disable rollback.  * false: enable rollback.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "stack_drift_status",
				Description: "Status of the stack's actual configuration compared to its expected template configuration.  * DRIFTED: The stack differs from its expected template configuration",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DriftInformation.StackDriftStatus"),
			},
			{
				Name:          "drift_last_check_timestamp",
				Description:   "Most recent time when a drift detection operation was initiated on the stack, or any of its individual resources that support drift detection.",
				Type:          schema.TypeTimestamp,
				Resolver:      schema.PathResolver("DriftInformation.LastCheckTimestamp"),
				IgnoreInTests: true,
			},
			{
				Name:          "enable_termination_protection",
				Description:   "Whether termination protection is enabled for the stack",
				Type:          schema.TypeBool,
				IgnoreInTests: true,
			},
			{
				Name:          "last_updated_time",
				Description:   "The time the stack was last updated",
				Type:          schema.TypeTimestamp,
				IgnoreInTests: true,
			},
			{
				Name:        "notification_arns",
				Description: "Amazon SNS topic Amazon Resource Names (ARNs) to which stack related events are published.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("NotificationARNs"),
			},
			{
				Name:        "parameters",
				Description: "A list of Parameter structures.",
				Type:        schema.TypeJSON,
			},
			{
				Name:          "parent_id",
				Description:   "For nested stacks--stacks created as resources for another stack--the stack ID of the direct parent of this stack",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "role_arn",
				Description:   "The Amazon Resource Name (ARN) of an Identity and Access Management (IAM) role that's associated with the stack",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("RoleARN"),
				IgnoreInTests: true,
			},
			{
				Name:          "rollback_configuration_monitoring_time_in_minutes",
				Description:   "The amount of time, in minutes, during which CloudFormation should monitor all the rollback triggers after the stack creation or update operation deploys all necessary resources",
				Type:          schema.TypeInt,
				Resolver:      schema.PathResolver("RollbackConfiguration.MonitoringTimeInMinutes"),
				IgnoreInTests: true,
			},
			{
				Name:          "rollback_configuration_rollback_triggers",
				Description:   "The triggers to monitor during stack creation or update actions",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:          "root_id",
				Description:   "For nested stacks--stacks created as resources for another stack--the stack ID of the top-level stack to which the nested stack ultimately belongs",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "id",
				Description: "Unique identifier of the stack.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StackId"),
			},
			{
				Name:          "stack_status_reason",
				Description:   "Success/failure message associated with the stack status.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "tags",
				Description: "A list of Tags that specify information about the stack.",
				Type:        schema.TypeJSON,
				Resolver:    resolveStacksTags,
			},
			{
				Name:        "timeout_in_minutes",
				Description: "The amount of time within which stack creation should complete.",
				Type:        schema.TypeInt,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_cloudformation_stack_outputs",
				Description:   "The Output data type.",
				Resolver:      fetchCloudformationStackOutputs,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "stack_cq_id",
						Description: "Unique CloudQuery ID of aws_cloudformation_stacks table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "description",
						Description: "User defined description associated with the output.",
						Type:        schema.TypeString,
					},
					{
						Name:        "export_name",
						Description: "The name of the export associated with the output.",
						Type:        schema.TypeString,
					},
					{
						Name:        "output_key",
						Description: "The key associated with the output.",
						Type:        schema.TypeString,
					},
					{
						Name:        "output_value",
						Description: "The value associated with the output.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_cloudformation_stack_resources",
				Description: "Contains high-level information about the specified stack resource.",
				Resolver:    fetchCloudformationStackResources,
				Columns: []schema.Column{
					{
						Name:        "stack_cq_id",
						Description: "Unique CloudQuery ID of aws_cloudformation_stacks table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "last_updated_timestamp",
						Description: "Time the status was updated.  This member is required.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "logical_resource_id",
						Description: "The logical name of the resource specified in the template.  This member is required.",
						Type:        schema.TypeString,
					},
					{
						Name:        "resource_status",
						Description: "Current status of the resource.  This member is required.",
						Type:        schema.TypeString,
					},
					{
						Name:        "resource_type",
						Description: "Type of resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "stack_resource_drift_status",
						Description: "Status of the resource's actual configuration compared to its expected configuration.  * DELETED: The resource differs from its expected configuration in that it has been deleted.  * MODIFIED: The resource differs from its expected configuration.  * NOT_CHECKED: CloudFormation hasn't checked if the resource differs from its expected configuration",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DriftInformation.StackResourceDriftStatus"),
					},
					{
						Name:          "drift_last_check_timestamp",
						Description:   "When CloudFormation last checked if the resource had drifted from its expected configuration.",
						Type:          schema.TypeTimestamp,
						Resolver:      schema.PathResolver("DriftInformation.LastCheckTimestamp"),
						IgnoreInTests: true,
					},
					{
						Name:          "module_info_logical_id_hierarchy",
						Description:   "A concatenated list of the logical IDs of the module or modules containing the resource",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("ModuleInfo.LogicalIdHierarchy"),
						IgnoreInTests: true,
					},
					{
						Name:          "module_info_type_hierarchy",
						Description:   "A concatenated list of the module type or types containing the resource",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("ModuleInfo.TypeHierarchy"),
						IgnoreInTests: true,
					},
					{
						Name:        "physical_resource_id",
						Description: "The name or unique identifier that corresponds to a physical instance ID of the resource.",
						Type:        schema.TypeString,
					},
					{
						Name:          "resource_status_reason",
						Description:   "Success/failure message associated with the resource.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchCloudformationStacks(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	var config cloudformation.DescribeStacksInput
	c := meta.(*client.Client)
	svc := c.Services().Cloudformation
	for {
		output, err := svc.DescribeStacks(ctx, &config, func(options *cloudformation.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.Stacks
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func resolveStacksTags(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(types.Stack)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
func fetchCloudformationStackOutputs(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.Stack)
	res <- r.Outputs
	return nil
}
func fetchCloudformationStackResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	stack := parent.Item.(types.Stack)
	config := cloudformation.ListStackResourcesInput{
		StackName: stack.StackName,
	}
	c := meta.(*client.Client)
	svc := c.Services().Cloudformation
	for {
		output, err := svc.ListStackResources(ctx, &config, func(options *cloudformation.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.StackResourceSummaries
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil

}
