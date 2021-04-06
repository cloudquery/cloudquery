package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/plugin/schema"
)

func AutoscalingLaunchConfigurations() *schema.Table {
	return &schema.Table{
		Name:         "aws_autoscaling_launch_configurations",
		Resolver:     fetchAutoscalingLaunchConfigurations,
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
				Name: "created_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "image_id",
				Type: schema.TypeString,
			},
			{
				Name: "instance_type",
				Type: schema.TypeString,
			},
			{
				Name: "launch_configuration_name",
				Type: schema.TypeString,
			},
			{
				Name: "associate_public_ip_address",
				Type: schema.TypeBool,
			},
			{
				Name:     "classic_link_vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClassicLinkVPCId"),
			},
			{
				Name:     "classic_link_vpc_security_groups",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ClassicLinkVPCSecurityGroups"),
			},
			{
				Name: "ebs_optimized",
				Type: schema.TypeBool,
			},
			{
				Name: "iam_instance_profile",
				Type: schema.TypeString,
			},
			{
				Name:     "instance_monitoring_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("InstanceMonitoring.Enabled"),
			},
			{
				Name: "kernel_id",
				Type: schema.TypeString,
			},
			{
				Name: "key_name",
				Type: schema.TypeString,
			},
			{
				Name:     "launch_configuration_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LaunchConfigurationARN"),
			},
			{
				Name:     "metadata_options_http_endpoint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MetadataOptions.HttpEndpoint"),
			},
			{
				Name:     "metadata_options_http_put_response_hop_limit",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MetadataOptions.HttpPutResponseHopLimit"),
			},
			{
				Name:     "metadata_options_http_tokens",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MetadataOptions.HttpTokens"),
			},
			{
				Name: "placement_tenancy",
				Type: schema.TypeString,
			},
			{
				Name: "ramdisk_id",
				Type: schema.TypeString,
			},
			{
				Name: "security_groups",
				Type: schema.TypeStringArray,
			},
			{
				Name: "spot_price",
				Type: schema.TypeString,
			},
			{
				Name: "user_data",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_autoscaling_launch_configuration_block_device_mappings",
				Resolver: fetchAutoscalingLaunchConfigurationBlockDeviceMappings,
				Columns: []schema.Column{
					{
						Name:     "launch_configuration_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "device_name",
						Type: schema.TypeString,
					},
					{
						Name:     "ebs_delete_on_termination",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("Ebs.DeleteOnTermination"),
					},
					{
						Name:     "ebs_encrypted",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("Ebs.Encrypted"),
					},
					{
						Name:     "ebs_iops",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("Ebs.Iops"),
					},
					{
						Name:     "ebs_snapshot_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Ebs.SnapshotId"),
					},
					{
						Name:     "ebs_volume_size",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("Ebs.VolumeSize"),
					},
					{
						Name:     "ebs_volume_type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Ebs.VolumeType"),
					},
					{
						Name: "no_device",
						Type: schema.TypeBool,
					},
					{
						Name: "virtual_name",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchAutoscalingLaunchConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Autoscaling
	config := autoscaling.DescribeLaunchConfigurationsInput{}
	for {
		output, err := svc.DescribeLaunchConfigurations(ctx, &config, func(o *autoscaling.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- output.LaunchConfigurations

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func fetchAutoscalingLaunchConfigurationBlockDeviceMappings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	res <- parent.Item.(types.LaunchConfiguration).BlockDeviceMappings
	return nil
}
