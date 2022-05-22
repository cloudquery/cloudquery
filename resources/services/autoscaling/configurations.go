package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func AutoscalingLaunchConfigurations() *schema.Table {
	return &schema.Table{
		Name:         "aws_autoscaling_launch_configurations",
		Description:  "Describes a launch configuration.",
		Resolver:     fetchAutoscalingLaunchConfigurations,
		Multiplex:    client.ServiceAccountRegionMultiplexer("autoscaling"),
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
				Name:        "created_time",
				Description: "The creation date and time for the launch configuration.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "image_id",
				Description: "The ID of the Amazon Machine Image (AMI) to use to launch your EC2 instances. For more information, see Finding an AMI (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/finding-an-ami.html) in the Amazon EC2 User Guide for Linux Instances.",
				Type:        schema.TypeString,
			},
			{
				Name:        "instance_type",
				Description: "The instance type for the instances. For information about available instance types, see Available Instance Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#AvailableInstanceTypes) in the Amazon EC2 User Guide for Linux Instances.",
				Type:        schema.TypeString,
			},
			{
				Name:        "launch_configuration_name",
				Description: "The name of the launch configuration.",
				Type:        schema.TypeString,
			},
			{
				Name:          "associate_public_ip_address",
				Description:   "For Auto Scaling groups that are running in a VPC, specifies whether to assign a public IP address to the group's instances. For more information, see Launching Auto Scaling instances in a VPC (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html) in the Amazon EC2 Auto Scaling User Guide.",
				Type:          schema.TypeBool,
				IgnoreInTests: true,
			},
			{
				Name:          "classic_link_vpc_id",
				Description:   "The ID of a ClassicLink-enabled VPC to link your EC2-Classic instances to. For more information, see ClassicLink (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) in the Amazon EC2 User Guide for Linux Instances and Linking EC2-Classic instances to a VPC (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html#as-ClassicLink) in the Amazon EC2 Auto Scaling User Guide.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ClassicLinkVPCId"),
				IgnoreInTests: true,
			},
			{
				Name:        "classic_link_vpc_security_groups",
				Description: "The IDs of one or more security groups for the VPC specified in ClassicLinkVPCId. For more information, see ClassicLink (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) in the Amazon EC2 User Guide for Linux Instances and Linking EC2-Classic instances to a VPC (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html#as-ClassicLink) in the Amazon EC2 Auto Scaling User Guide.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("ClassicLinkVPCSecurityGroups"),
			},
			{
				Name:        "ebs_optimized",
				Description: "Specifies whether the launch configuration is optimized for EBS I/O (true) or not (false). For more information, see Amazon EBS-Optimized Instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSOptimized.html) in the Amazon EC2 User Guide for Linux Instances.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "iam_instance_profile",
				Description: "The name or the Amazon Resource Name (ARN) of the instance profile associated with the IAM role for the instance. The instance profile contains the IAM role. For more information, see IAM role for applications that run on Amazon EC2 instances (https://docs.aws.amazon.com/autoscaling/ec2/userguide/us-iam-role.html) in the Amazon EC2 Auto Scaling User Guide.",
				Type:        schema.TypeString,
			},
			{
				Name:        "instance_monitoring_enabled",
				Description: "If true, detailed monitoring is enabled. Otherwise, basic monitoring is enabled.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("InstanceMonitoring.Enabled"),
			},
			{
				Name:        "kernel_id",
				Description: "The ID of the kernel associated with the AMI.",
				Type:        schema.TypeString,
			},
			{
				Name:        "key_name",
				Description: "The name of the key pair. For more information, see Amazon EC2 Key Pairs (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html) in the Amazon EC2 User Guide for Linux Instances.",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the launch configuration.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LaunchConfigurationARN"),
			},
			{
				Name:        "metadata_options_http_endpoint",
				Description: "This parameter enables or disables the HTTP metadata endpoint on your instances. If the parameter is not specified, the default state is enabled. If you specify a value of disabled, you will not be able to access your instance metadata.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MetadataOptions.HttpEndpoint"),
			},
			{
				Name:        "metadata_options_http_put_response_hop_limit",
				Description: "The desired HTTP PUT response hop limit for instance metadata requests. The larger the number, the further instance metadata requests can travel. Default: 1",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("MetadataOptions.HttpPutResponseHopLimit"),
			},
			{
				Name:        "metadata_options_http_tokens",
				Description: "The state of token usage for your instance metadata requests.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MetadataOptions.HttpTokens"),
			},
			{
				Name:          "placement_tenancy",
				Description:   "The tenancy of the instance, either default or dedicated. An instance with dedicated tenancy runs on isolated, single-tenant hardware and can only be launched into a VPC. For more information, see Configuring instance tenancy with Amazon EC2 Auto Scaling (https://docs.aws.amazon.com/autoscaling/ec2/userguide/auto-scaling-dedicated-instances.html) in the Amazon EC2 Auto Scaling User Guide.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "ramdisk_id",
				Description: "The ID of the RAM disk associated with the AMI.",
				Type:        schema.TypeString,
			},
			{
				Name:        "security_groups",
				Description: "A list that contains the security groups to assign to the instances in the Auto Scaling group. For more information, see Security Groups for Your VPC (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_SecurityGroups.html) in the Amazon Virtual Private Cloud User Guide.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:          "spot_price",
				Description:   "The maximum hourly price to be paid for any Spot Instance launched to fulfill the request. Spot Instances are launched when the price you specify exceeds the current Spot price. For more information, see Requesting Spot Instances (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-launch-spot-instances.html) in the Amazon EC2 Auto Scaling User Guide.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "user_data",
				Description: "The user data to make available to the launched EC2 instances. For more information, see Instance metadata and user data (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html) (Linux) and Instance metadata and user data (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ec2-instance-metadata.html) (Windows). If you are using a command line tool, base64-encoding is performed for you, and you can load the text from a file. Otherwise, you must provide base64-encoded text. User data is limited to 16 KB.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_autoscaling_launch_configuration_block_device_mappings",
				Description: "Describes a block device mapping.",
				Resolver:    fetchAutoscalingLaunchConfigurationBlockDeviceMappings,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"launch_configuration_cq_id", "device_name"}},
				Columns: []schema.Column{
					{
						Name:        "launch_configuration_cq_id",
						Description: "Unique CloudQuery ID of aws_autoscaling_launch_configurations table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "device_name",
						Description: "The device name exposed to the EC2 instance (for example, /dev/sdh or xvdh). For more information, see Device Naming on Linux Instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/device_naming.html) in the Amazon EC2 User Guide for Linux Instances.",
						Type:        schema.TypeString,
					},
					{
						Name:        "ebs_delete_on_termination",
						Description: "Indicates whether the volume is deleted on instance termination. For Amazon EC2 Auto Scaling, the default value is true.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Ebs.DeleteOnTermination"),
					},
					{
						Name:          "ebs_encrypted",
						Description:   "Specifies whether the volume should be encrypted.",
						Type:          schema.TypeBool,
						Resolver:      schema.PathResolver("Ebs.Encrypted"),
						IgnoreInTests: true,
					},
					{
						Name:          "ebs_iops",
						Description:   "The number of I/O operations per second (IOPS) to provision for the volume.",
						Type:          schema.TypeInt,
						Resolver:      schema.PathResolver("Ebs.Iops"),
						IgnoreInTests: true,
					},
					{
						Name:          "ebs_snapshot_id",
						Description:   "The snapshot ID of the volume to use. You must specify either a VolumeSize or a SnapshotId.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("Ebs.SnapshotId"),
						IgnoreInTests: true,
					},
					{
						Name:        "ebs_volume_size",
						Description: "The volume size, in Gibibytes (GiB).",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Ebs.VolumeSize"),
					},
					{
						Name:        "ebs_volume_type",
						Description: "The volume type, which can be standard for Magnetic, io1 for Provisioned IOPS SSD, gp2 for General Purpose SSD, st1 for Throughput Optimized HDD, or sc1 for Cold HDD. For more information, see Amazon EBS Volume Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html) in the Amazon EC2 User Guide for Linux Instances. Valid Values: standard | io1 | gp2 | st1 | sc1",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Ebs.VolumeType"),
					},
					{
						Name:          "no_device",
						Description:   "Setting this value to true suppresses the specified device included in the block device mapping of the AMI.",
						Type:          schema.TypeBool,
						IgnoreInTests: true,
					},
					{
						Name:          "virtual_name",
						Description:   "The name of the virtual device (for example, ephemeral0). You can specify either VirtualName or Ebs, but not both.",
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
func fetchAutoscalingLaunchConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Autoscaling
	config := autoscaling.DescribeLaunchConfigurationsInput{}
	for {
		output, err := svc.DescribeLaunchConfigurations(ctx, &config, func(o *autoscaling.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.LaunchConfigurations

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func fetchAutoscalingLaunchConfigurationBlockDeviceMappings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	res <- parent.Item.(types.LaunchConfiguration).BlockDeviceMappings
	return nil
}
