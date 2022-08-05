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

func Ec2FlowLogs() *schema.Table {
	return &schema.Table{
		Name:          "aws_ec2_flow_logs",
		Description:   "Describes a flow log.",
		Resolver:      fetchEc2FlowLogs,
		Multiplex:     client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreError:   client.IgnoreCommonErrors,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
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
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARN(client.EC2Service, func(resource *schema.Resource) ([]string, error) {
					return []string{"vpc-flow-log", *resource.Item.(types.FlowLog).FlowLogId}, nil
				}),
			},
			{
				Name:        "id",
				Description: "The flow log ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("FlowLogId"),
			},
			{
				Name:        "creation_time",
				Description: "The date and time the flow log was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "deliver_logs_error_message",
				Description: "Information about the error that occurred.",
				Type:        schema.TypeString,
			},
			{
				Name:        "deliver_logs_permission_arn",
				Description: "The ARN of the IAM role that posts logs to CloudWatch Logs.",
				Type:        schema.TypeString,
			},
			{
				Name:        "deliver_logs_status",
				Description: "The status of the logs delivery (SUCCESS | FAILED).",
				Type:        schema.TypeString,
			},
			{
				Name:        "flow_log_id",
				Description: "The flow log ID.",
				Type:        schema.TypeString,
			},
			{
				Name:        "flow_log_status",
				Description: "The status of the flow log (ACTIVE).",
				Type:        schema.TypeString,
			},
			{
				Name:        "log_destination",
				Description: "Specifies the destination to which the flow log data is published.",
				Type:        schema.TypeString,
			},
			{
				Name:        "log_destination_type",
				Description: "Specifies the type of destination to which the flow log data is published.",
				Type:        schema.TypeString,
			},
			{
				Name:        "log_format",
				Description: "The format of the flow log record.",
				Type:        schema.TypeString,
			},
			{
				Name:        "log_group_name",
				Description: "The name of the flow log group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "max_aggregation_interval",
				Description: "The maximum interval of time, in seconds, during which a flow of packets is captured and aggregated into a flow log record.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "resource_id",
				Description: "The ID of the resource on which the flow log was created.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The tags for the flow log.",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "traffic_type",
				Description: "The type of traffic captured for the flow log.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEc2FlowLogs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeFlowLogsInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeFlowLogs(ctx, &config, func(options *ec2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.FlowLogs
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
