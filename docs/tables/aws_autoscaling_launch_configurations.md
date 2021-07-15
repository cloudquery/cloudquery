
# Table: aws_autoscaling_launch_configurations
Describes a launch configuration.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|created_time|timestamp without time zone|The creation date and time for the launch configuration.|
|image_id|text|The ID of the Amazon Machine Image (AMI) to use to launch your EC2 instances. For more information, see Finding an AMI (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/finding-an-ami.html) in the Amazon EC2 User Guide for Linux Instances.|
|instance_type|text|The instance type for the instances. For information about available instance types, see Available Instance Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#AvailableInstanceTypes) in the Amazon EC2 User Guide for Linux Instances.|
|launch_configuration_name|text|The name of the launch configuration.|
|associate_public_ip_address|boolean|For Auto Scaling groups that are running in a VPC, specifies whether to assign a public IP address to the group's instances. For more information, see Launching Auto Scaling instances in a VPC (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html) in the Amazon EC2 Auto Scaling User Guide.|
|classic_link_vpc_id|text|The ID of a ClassicLink-enabled VPC to link your EC2-Classic instances to. For more information, see ClassicLink (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) in the Amazon EC2 User Guide for Linux Instances and Linking EC2-Classic instances to a VPC (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html#as-ClassicLink) in the Amazon EC2 Auto Scaling User Guide.|
|classic_link_vpc_security_groups|text[]|The IDs of one or more security groups for the VPC specified in ClassicLinkVPCId. For more information, see ClassicLink (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) in the Amazon EC2 User Guide for Linux Instances and Linking EC2-Classic instances to a VPC (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html#as-ClassicLink) in the Amazon EC2 Auto Scaling User Guide.|
|ebs_optimized|boolean|Specifies whether the launch configuration is optimized for EBS I/O (true) or not (false). For more information, see Amazon EBS-Optimized Instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSOptimized.html) in the Amazon EC2 User Guide for Linux Instances.|
|iam_instance_profile|text|The name or the Amazon Resource Name (ARN) of the instance profile associated with the IAM role for the instance. The instance profile contains the IAM role. For more information, see IAM role for applications that run on Amazon EC2 instances (https://docs.aws.amazon.com/autoscaling/ec2/userguide/us-iam-role.html) in the Amazon EC2 Auto Scaling User Guide.|
|instance_monitoring_enabled|boolean|If true, detailed monitoring is enabled. Otherwise, basic monitoring is enabled.|
|kernel_id|text|The ID of the kernel associated with the AMI.|
|key_name|text|The name of the key pair. For more information, see Amazon EC2 Key Pairs (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html) in the Amazon EC2 User Guide for Linux Instances.|
|arn|text|The Amazon Resource Name (ARN) of the launch configuration.|
|metadata_options_http_endpoint|text|This parameter enables or disables the HTTP metadata endpoint on your instances. If the parameter is not specified, the default state is enabled. If you specify a value of disabled, you will not be able to access your instance metadata.|
|metadata_options_http_put_response_hop_limit|integer|The desired HTTP PUT response hop limit for instance metadata requests. The larger the number, the further instance metadata requests can travel. Default: 1|
|metadata_options_http_tokens|text|The state of token usage for your instance metadata requests.|
|placement_tenancy|text|The tenancy of the instance, either default or dedicated. An instance with dedicated tenancy runs on isolated, single-tenant hardware and can only be launched into a VPC. For more information, see Configuring instance tenancy with Amazon EC2 Auto Scaling (https://docs.aws.amazon.com/autoscaling/ec2/userguide/auto-scaling-dedicated-instances.html) in the Amazon EC2 Auto Scaling User Guide.|
|ramdisk_id|text|The ID of the RAM disk associated with the AMI.|
|security_groups|text[]|A list that contains the security groups to assign to the instances in the Auto Scaling group. For more information, see Security Groups for Your VPC (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_SecurityGroups.html) in the Amazon Virtual Private Cloud User Guide.|
|spot_price|text|The maximum hourly price to be paid for any Spot Instance launched to fulfill the request. Spot Instances are launched when the price you specify exceeds the current Spot price. For more information, see Requesting Spot Instances (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-launch-spot-instances.html) in the Amazon EC2 Auto Scaling User Guide.|
|user_data|text|The user data to make available to the launched EC2 instances. For more information, see Instance metadata and user data (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html) (Linux) and Instance metadata and user data (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ec2-instance-metadata.html) (Windows). If you are using a command line tool, base64-encoding is performed for you, and you can load the text from a file. Otherwise, you must provide base64-encoded text. User data is limited to 16 KB.|
