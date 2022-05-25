package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Ec2InstanceStatuses() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_instance_statuses",
		Description:  "Describes the status of an instance.",
		Resolver:     fetchEc2InstanceStatuses,
		Multiplex:    client.ServiceAccountRegionMultiplexer("ec2"),
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
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARN(client.EC2Service, func(resource *schema.Resource) ([]string, error) {
					return []string{"instance", *resource.Item.(types.InstanceStatus).InstanceId}, nil
				}),
			},
			{
				Name:        "availability_zone",
				Description: "The Availability Zone of the instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "instance_id",
				Description: "The ID of the instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "instance_state_code",
				Description: "The state of the instance as a 16-bit unsigned integer.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("InstanceState.Code"),
			},
			{
				Name:        "instance_state_name",
				Description: "The current state of the instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InstanceState.Name"),
			},
			{
				Name:        "details",
				Description: "The system instance health or application instance health.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("InstanceStatus.Details"),
			},
			{
				Name:        "status",
				Description: "The instance status.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InstanceStatus.Status"),
			},
			{
				Name:          "outpost_arn",
				Description:   "The Amazon Resource Name (ARN) of the Outpost.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "system_status",
				Description: "The system status.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SystemStatus.Status"),
			},
			{
				Name:        "system_status_details",
				Description: "The system instance health or application instance health.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("SystemStatus.Details"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_ec2_instance_status_events",
				Description:   "Any scheduled events associated with the instance.",
				Resolver:      fetchEc2InstanceStatusEvents,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "instance_status_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_instance_statuses table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "code",
						Description: "The event code.",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "A description of the event.",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "The ID of the event.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InstanceEventId"),
					},
					{
						Name:        "not_after",
						Description: "The latest scheduled end time for the event.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "not_before",
						Description: "The earliest scheduled start time for the event.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "not_before_deadline",
						Description: "The deadline for starting the event.",
						Type:        schema.TypeTimestamp,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchEc2InstanceStatuses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeInstanceStatusInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeInstanceStatus(ctx, &config, func(options *ec2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.InstanceStatuses
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func fetchEc2InstanceStatusEvents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.InstanceStatus)
	res <- r.Events
	return nil
}
