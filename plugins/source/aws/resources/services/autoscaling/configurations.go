package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func AutoscalingLaunchConfigurations() *schema.Table {
	return &schema.Table{
		Name:        "aws_autoscaling_launch_configurations",
		Description: "Describes a launch configuration.",
		Resolver:    fetchAutoscalingLaunchConfigurations,
		Multiplex:   client.ServiceAccountRegionMultiplexer("autoscaling"),
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
				Name:     "instance_monitoring",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("InstanceMonitoring"),
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
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) of the launch configuration.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("LaunchConfigurationARN"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:     "metadata_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MetadataOptions"),
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
			{
				Name:        "block_device_mappings",
				Description: "Describes a block device mapping.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("BlockDeviceMappings"),
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchAutoscalingLaunchConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Autoscaling
	config := autoscaling.DescribeLaunchConfigurationsInput{}
	for {
		output, err := svc.DescribeLaunchConfigurations(ctx, &config)
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
