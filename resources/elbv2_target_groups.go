package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/plugin/schema"
)

func Elbv2TargetGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_elbv2_target_groups",
		Resolver:     fetchElbv2TargetGroups,
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
				Name: "health_check_enabled",
				Type: schema.TypeBool,
			},
			{
				Name: "health_check_interval_seconds",
				Type: schema.TypeInt,
			},
			{
				Name: "health_check_path",
				Type: schema.TypeString,
			},
			{
				Name: "health_check_port",
				Type: schema.TypeString,
			},
			{
				Name: "health_check_protocol",
				Type: schema.TypeString,
			},
			{
				Name: "health_check_timeout_seconds",
				Type: schema.TypeInt,
			},
			{
				Name: "healthy_threshold_count",
				Type: schema.TypeInt,
			},
			{
				Name: "load_balancer_arns",
				Type: schema.TypeStringArray,
			},
			{
				Name:     "matcher_grpc_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Matcher.GrpcCode"),
			},
			{
				Name:     "matcher_http_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Matcher.HttpCode"),
			},
			{
				Name: "port",
				Type: schema.TypeInt,
			},
			{
				Name: "protocol",
				Type: schema.TypeString,
			},
			{
				Name: "protocol_version",
				Type: schema.TypeString,
			},
			{
				Name: "target_group_arn",
				Type: schema.TypeString,
			},
			{
				Name: "target_group_name",
				Type: schema.TypeString,
			},
			{
				Name: "target_type",
				Type: schema.TypeString,
			},
			{
				Name: "unhealthy_threshold_count",
				Type: schema.TypeInt,
			},
			{
				Name: "vpc_id",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchElbv2TargetGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config elbv2.DescribeTargetGroupsInput
	c := meta.(*client.Client)
	svc := c.Services().ELBv2
	for {
		response, err := svc.DescribeTargetGroups(ctx, &config, func(options *elbv2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.TargetGroups
		if aws.ToString(response.NextMarker) == "" {
			break
		}
		config.Marker = response.NextMarker
	}
	return nil
}
