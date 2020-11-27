package autoscaling

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type LaunchConfiguration struct {
	ID uint `gorm:"primarykey"`

	// AWS account id
	AccountID string

	// AWS region of the launch configuration
	Region string

	// For Auto Scaling groups that are running in a VPC, specifies whether to assign
	// a public IP address to the group's instances.
	AssociatePublicIpAddress *bool

	// A block device mapping, which specifies the block devices for the instance.
	BlockDeviceMappings []*LaunchConfigurationBlockDeviceMapping `gorm:"constraint:OnDelete:CASCADE;"`

	// The ID of a ClassicLink-enabled VPC to link your EC2-Classic instances to.
	ClassicLinkVPCId *string

	// The IDs of one or more security groups for the VPC specified in ClassicLinkVPCId.
	ClassicLinkVPCSecurityGroups *string

	// The creation date and time for the launch configuration.
	CreatedTime *time.Time

	// Specifies whether the launch configuration is optimized for EBS I/O (true)
	// or not (false).
	EbsOptimized *bool

	// The name or the Amazon Resource Name (ARN) of the instance profile associated
	// with the IAM role for the instance. The instance profile contains the IAM
	// role.
	IamInstanceProfile *string

	// The ID of the Amazon Machine Image (AMI) to use to launch your EC2 instances.
	ImageId *string

	// Controls whether instances in this group are launched with detailed (true)
	// or basic (false) monitoring.
	InstanceMonitoring *autoscaling.InstanceMonitoring `gorm:"embedded;embeddedPrefix:instance_monitoring_"`

	// The instance type for the instances.
	InstanceType *string

	// The ID of the kernel associated with the AMI.
	KernelId *string

	// The name of the key pair.
	KeyName *string

	// The Amazon Resource Name (ARN) of the launch configuration.
	LaunchConfigurationARN *string

	// The name of the launch configuration.
	LaunchConfigurationName *string

	// The metadata options for the instances. For more information, see Instance
	MetadataOptions *autoscaling.InstanceMetadataOptions `gorm:"embedded;embeddedPrefix:metadata_options_"`

	// The tenancy of the instance, either default or dedicated. An instance with
	// dedicated tenancy runs on isolated, single-tenant hardware and can only be
	// launched into a VPC.
	PlacementTenancy *string

	// The ID of the RAM disk associated with the AMI.
	RamdiskId *string

	// A list that contains the security groups to assign to the instances in the
	// Auto Scaling group.
	SecurityGroups *string

	// The maximum hourly price to be paid for any Spot Instance launched to fulfill
	// the request. Spot Instances are launched when the price you specify exceeds
	// the current Spot price.
	SpotPrice *string

	// The Base64-encoded user data to make available to the launched EC2 instances.
	UserData *string
}

type LaunchConfigurationBlockDeviceMapping struct {
	ID                    uint `gorm:"primarykey"`
	LaunchConfigurationID uint

	// The device name exposed to the EC2 instance (for example, /dev/sdh or xvdh).
	// For more information, see Device Naming on Linux Instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/device_naming.html)
	// in the Amazon EC2 User Guide for Linux Instances.
	DeviceName *string

	// Parameters used to automatically set up EBS volumes when an instance is launched.
	Ebs *autoscaling.Ebs `gorm:"embedded;embeddedPrefix:ebs_"`

	// If NoDevice is true for the root device, instances might fail the EC2 health
	// check. In that case, Amazon EC2 Auto Scaling launches replacement instances.
	NoDevice *bool

	// The name of the virtual device (for example, ephemeral0).
	VirtualName *string
}

func (c *Client) transformLaunchConfigurationBlockDeviceMapping(value *autoscaling.BlockDeviceMapping) *LaunchConfigurationBlockDeviceMapping {
	return &LaunchConfigurationBlockDeviceMapping{
		DeviceName:  value.DeviceName,
		Ebs:         value.Ebs,
		NoDevice:    value.NoDevice,
		VirtualName: value.VirtualName,
	}
}

func (c *Client) transformLaunchConfigurationBlockDeviceMappings(values []*autoscaling.BlockDeviceMapping) []*LaunchConfigurationBlockDeviceMapping {
	var tValues []*LaunchConfigurationBlockDeviceMapping
	for _, v := range values {
		tValues = append(tValues, c.transformLaunchConfigurationBlockDeviceMapping(v))
	}
	return tValues
}

func (c *Client) transformLaunchConfiguration(value *autoscaling.LaunchConfiguration) *LaunchConfiguration {
	return &LaunchConfiguration{
		Region:                       c.region,
		AccountID:                    c.accountID,
		AssociatePublicIpAddress:     value.AssociatePublicIpAddress,
		BlockDeviceMappings:          c.transformLaunchConfigurationBlockDeviceMappings(value.BlockDeviceMappings),
		ClassicLinkVPCId:             value.ClassicLinkVPCId,
		ClassicLinkVPCSecurityGroups: common.StringListToString(value.ClassicLinkVPCSecurityGroups),
		CreatedTime:                  value.CreatedTime,
		EbsOptimized:                 value.EbsOptimized,
		IamInstanceProfile:           value.IamInstanceProfile,
		ImageId:                      value.ImageId,
		InstanceMonitoring:           value.InstanceMonitoring,
		InstanceType:                 value.InstanceType,
		KernelId:                     value.KernelId,
		KeyName:                      value.KeyName,
		LaunchConfigurationARN:       value.LaunchConfigurationARN,
		LaunchConfigurationName:      value.LaunchConfigurationName,
		MetadataOptions:              value.MetadataOptions,
		PlacementTenancy:             value.PlacementTenancy,
		RamdiskId:                    value.RamdiskId,
		SecurityGroups:               common.StringListToString(value.SecurityGroups),
		SpotPrice:                    value.SpotPrice,
		UserData:                     value.UserData,
	}
}

func (c *Client) transformLaunchConfigurations(values []*autoscaling.LaunchConfiguration) []*LaunchConfiguration {
	var tValues []*LaunchConfiguration
	for _, v := range values {
		tValues = append(tValues, c.transformLaunchConfiguration(v))
	}
	return tValues
}

func (c *Client) LaunchConfigurations(gConfig interface{}) error {
	var config autoscaling.DescribeLaunchConfigurationsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["autoscalingLaunchConfiguration"] {
		err := c.db.AutoMigrate(
			&LaunchConfiguration{},
			&LaunchConfigurationBlockDeviceMapping{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["autoscalingLaunchConfiguration"] = true
	}
	for {
		output, err := c.svc.DescribeLaunchConfigurations(&config)
		if err != nil {
			return err
		}
		c.log.Debug("deleting previous LaunchConfigurations", zap.String("region", c.region), zap.String("account_id", c.accountID))
		c.db.Where("region = ?", c.region).Where("account_id = ?", c.accountID).Delete(&LaunchConfiguration{})
		common.ChunkedCreate(c.db, c.transformLaunchConfigurations(output.LaunchConfigurations))
		c.log.Info("populating LaunchConfigurations", zap.Int("count", len(output.LaunchConfigurations)))
		if aws.StringValue(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
