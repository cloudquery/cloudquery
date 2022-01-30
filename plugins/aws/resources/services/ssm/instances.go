package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SsmInstances() *schema.Table {
	return &schema.Table{
		Name:          "aws_ssm_instances",
		Description:   "Describes a filter for a specific list of instances.",
		Resolver:      fetchSsmInstances,
		Multiplex:     client.ServiceAccountRegionMultiplexer("ssm"),
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
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the managed instance.",
				Type:        schema.TypeString,
				Resolver:    resolveSSMInstanceARN,
			},
			{
				Name:        "activation_id",
				Description: "The activation ID created by Amazon Web Services Systems Manager when the server or virtual machine (VM) was registered.",
				Type:        schema.TypeString,
			},
			{
				Name:        "agent_version",
				Description: "The version of SSM Agent running on your Linux instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "association_overview_detailed_status",
				Description: "Detailed status information about the aggregated associations.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AssociationOverview.DetailedStatus"),
			},
			{
				Name:        "association_instance_status_aggregated_count",
				Description: "The number of associations for the instance(s).",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("AssociationOverview.InstanceAssociationStatusAggregatedCount"),
			},
			{
				Name:        "association_status",
				Description: "The status of the association.",
				Type:        schema.TypeString,
			},
			{
				Name:        "computer_name",
				Description: "The fully qualified host name of the managed instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "ip_address",
				Description: "The IP address of the managed instance.",
				Type:        schema.TypeInet,
				Resolver:    schema.IPAddressResolver("IPAddress"),
			},
			{
				Name:        "iam_role",
				Description: "The Identity and Access Management (IAM) role assigned to the on-premises Systems Manager managed instance",
				Type:        schema.TypeString,
			},
			{
				Name:        "instance_id",
				Description: "The instance ID.",
				Type:        schema.TypeString,
			},
			{
				Name:        "is_latest_version",
				Description: "Indicates whether the latest version of SSM Agent is running on your Linux Managed Instance",
				Type:        schema.TypeBool,
			},
			{
				Name:        "last_association_execution_date",
				Description: "The date the association was last run.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "last_ping_date_time",
				Description: "The date and time when the agent last pinged the Systems Manager service.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "last_successful_association_execution_date",
				Description: "The last date the association was successfully run.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "name",
				Description: "The name assigned to an on-premises server or virtual machine (VM) when it is activated as a Systems Manager managed instance",
				Type:        schema.TypeString,
			},
			{
				Name:        "ping_status",
				Description: "Connection status of SSM Agent",
				Type:        schema.TypeString,
			},
			{
				Name:        "platform_name",
				Description: "The name of the operating system platform running on your instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "platform_type",
				Description: "The operating system platform type.",
				Type:        schema.TypeString,
			},
			{
				Name:        "platform_version",
				Description: "The version of the OS platform running on your instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "registration_date",
				Description: "The date the server or VM was registered with Amazon Web Services as a managed instance.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "resource_type",
				Description: "The type of instance",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_ssm_instance_compliance_items",
				Description: "Information about the compliance as defined by the resource type",
				Resolver:    fetchSsmInstanceComplianceItems,
				IgnoreError: client.IgnoreAccessDeniedServiceDisabled,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"instance_cq_id", "resource_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_ssm_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "compliance_type",
						Description: "The compliance type",
						Type:        schema.TypeString,
					},
					{
						Name:        "details",
						Description: "A \"Key\": \"Value\" tag combination for the compliance item.",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "execution_summary_execution_time",
						Description: "The time the execution ran as a datetime object that is saved in the following format: yyyy-MM-dd'T'HH:mm:ss'Z'.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("ExecutionSummary.ExecutionTime"),
					},
					{
						Name:        "execution_summary_execution_id",
						Description: "An ID created by the system when PutComplianceItems was called",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExecutionSummary.ExecutionId"),
					},
					{
						Name:        "execution_summary_execution_type",
						Description: "The type of execution",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExecutionSummary.ExecutionType"),
					},
					{
						Name:        "id",
						Description: "An ID for the compliance item",
						Type:        schema.TypeString,
					},
					{
						Name:        "resource_id",
						Description: "An ID for the resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "resource_type",
						Description: "The type of resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "severity",
						Description: "The severity of the compliance status",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "The status of the compliance item",
						Type:        schema.TypeString,
					},
					{
						Name:        "title",
						Description: "A title for the compliance item",
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
func fetchSsmInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	client := meta.(*client.Client)
	svc := client.Services().SSM
	optsFn := func(o *ssm.Options) {
		o.Region = client.Region
	}
	var input ssm.DescribeInstanceInformationInput
	for {
		output, err := svc.DescribeInstanceInformation(ctx, &input, optsFn)
		if err != nil {
			return err
		}
		res <- output.InstanceInformationList
		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	return nil
}

func fetchSsmInstanceComplianceItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance, ok := parent.Item.(types.InstanceInformation)
	if !ok {
		return fmt.Errorf("not a %T instance: %T", instance, parent.Item)
	}
	client := meta.(*client.Client)
	svc := client.Services().SSM
	optsFn := func(o *ssm.Options) {
		o.Region = client.Region
	}
	input := ssm.ListComplianceItemsInput{
		ResourceIds: []string{*instance.InstanceId},
	}
	for {
		output, err := svc.ListComplianceItems(ctx, &input, optsFn)
		if err != nil {
			return err
		}
		res <- output.ComplianceItems
		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	return nil
}

func resolveSSMInstanceARN(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	instance, ok := resource.Item.(types.InstanceInformation)
	if !ok {
		return fmt.Errorf("not a %T instance: %T", instance, resource.Item)
	}
	cl := meta.(*client.Client)
	return resource.Set(c.Name, client.GenerateResourceARN("ssm", "managed-instance", *instance.InstanceId, cl.Region, cl.AccountID))
}
