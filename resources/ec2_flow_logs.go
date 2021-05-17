package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Ec2FlowLogs() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_flow_logs",
		Resolver:     fetchEc2FlowLogs,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
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
				Name: "creation_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "deliver_logs_error_message",
				Type: schema.TypeString,
			},
			{
				Name: "deliver_logs_permission_arn",
				Type: schema.TypeString,
			},
			{
				Name: "deliver_logs_status",
				Type: schema.TypeString,
			},
			{
				Name: "flow_log_id",
				Type: schema.TypeString,
			},
			{
				Name: "flow_log_status",
				Type: schema.TypeString,
			},
			{
				Name: "log_destination",
				Type: schema.TypeString,
			},
			{
				Name: "log_destination_type",
				Type: schema.TypeString,
			},
			{
				Name: "log_format",
				Type: schema.TypeString,
			},
			{
				Name: "log_group_name",
				Type: schema.TypeString,
			},
			{
				Name: "max_aggregation_interval",
				Type: schema.TypeInt,
			},
			{
				Name: "resource_id",
				Type: schema.TypeString,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveEc2flowLogTags,
			},
			{
				Name: "traffic_type",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEc2FlowLogs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config ec2.DescribeFlowLogsInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeFlowLogs(ctx, &config, func(options *ec2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- output.FlowLogs
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func resolveEc2flowLogTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.FlowLog)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
