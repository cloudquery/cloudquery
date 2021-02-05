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
	_  interface{} `neo:"raw:MERGE (a:AWSAccount {account_id: $account_id}) MERGE (a) - [:Resource] -> (n)"`
	ID uint        `gorm:"primarykey"`

	// AWS account id
	AccountID string

	Region                       string
	AssociatePublicIpAddress     *bool
	BlockDeviceMappings          []*LaunchConfigurationBlockDeviceMapping `gorm:"constraint:OnDelete:CASCADE;"`
	ClassicLinkVPCId             *string
	ClassicLinkVPCSecurityGroups *string
	CreatedTime                  *time.Time
	EbsOptimized                 *bool
	IamInstanceProfile           *string
	ImageId                      *string
	InstanceMonitoringEnabled    *bool
	InstanceType                 *string
	KernelId                     *string
	KeyName                      *string

	LaunchConfigurationARN  *string `neo:"unique"`
	LaunchConfigurationName *string

	MetadataHttpEndpoint            *string
	MetadataHttpPutResponseHopLimit *int64
	MetadataHttpTokens              *string

	PlacementTenancy *string
	RamdiskId        *string
	SecurityGroups   *string
	SpotPrice        *string
	UserData         *string
}

func (LaunchConfiguration) TableName() string {
	return "aws_autoscaling_launch_configurations"
}

type LaunchConfigurationBlockDeviceMapping struct {
	ID                    uint `gorm:"primarykey"`
	LaunchConfigurationID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	// The device name exposed to the EC2 instance (for example, /dev/sdh or xvdh).
	// For more information, see Device Naming on Linux Instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/device_naming.html)
	// in the Amazon EC2 User Guide for Linux Instances.
	DeviceName *string

	EbsDeleteOnTermination *bool
	EbsEncrypted           *bool
	EbsIops                *int64
	EbsSnapshotId          *string
	EbsVolumeSize          *int64
	EbsVolumeType          *string

	// If NoDevice is true for the root device, instances might fail the EC2 health
	// check. In that case, Amazon EC2 Auto Scaling launches replacement instances.
	NoDevice *bool

	// The name of the virtual device (for example, ephemeral0).
	VirtualName *string
}

func (LaunchConfigurationBlockDeviceMapping) TableName() string {
	return "aws_autoscaling_launch_configuration_block_device_mapping"
}

func (c *Client) transformLaunchConfigurationBlockDeviceMapping(value *autoscaling.BlockDeviceMapping) *LaunchConfigurationBlockDeviceMapping {
	res := LaunchConfigurationBlockDeviceMapping{
		AccountID:   c.accountID,
		Region:      c.region,
		DeviceName:  value.DeviceName,
		NoDevice:    value.NoDevice,
		VirtualName: value.VirtualName,
	}
	if value.Ebs != nil {
		res.EbsDeleteOnTermination = value.Ebs.DeleteOnTermination
		res.EbsEncrypted = value.Ebs.Encrypted
		res.EbsIops = value.Ebs.Iops
		res.EbsSnapshotId = value.Ebs.SnapshotId
		res.EbsVolumeSize = value.Ebs.VolumeSize
		res.EbsVolumeType = value.Ebs.VolumeType
	}
	return &res
}

func (c *Client) transformLaunchConfigurationBlockDeviceMappings(values []*autoscaling.BlockDeviceMapping) []*LaunchConfigurationBlockDeviceMapping {
	var tValues []*LaunchConfigurationBlockDeviceMapping
	for _, v := range values {
		tValues = append(tValues, c.transformLaunchConfigurationBlockDeviceMapping(v))
	}
	return tValues
}

func (c *Client) transformLaunchConfiguration(value *autoscaling.LaunchConfiguration) *LaunchConfiguration {
	res := LaunchConfiguration{
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
		InstanceType:                 value.InstanceType,
		KernelId:                     value.KernelId,
		KeyName:                      value.KeyName,
		LaunchConfigurationARN:       value.LaunchConfigurationARN,
		LaunchConfigurationName:      value.LaunchConfigurationName,
		PlacementTenancy:             value.PlacementTenancy,
		RamdiskId:                    value.RamdiskId,
		SecurityGroups:               common.StringListToString(value.SecurityGroups),
		SpotPrice:                    value.SpotPrice,
		UserData:                     value.UserData,
	}

	if value.MetadataOptions != nil {
		res.MetadataHttpEndpoint = value.MetadataOptions.HttpEndpoint
		res.MetadataHttpPutResponseHopLimit = value.MetadataOptions.HttpPutResponseHopLimit
		res.MetadataHttpTokens = value.MetadataOptions.HttpTokens
	}

	if value.InstanceMonitoring != nil {
		res.InstanceMonitoringEnabled = value.InstanceMonitoring.Enabled
	}

	return &res
}

func (c *Client) transformLaunchConfigurations(values []*autoscaling.LaunchConfiguration) []*LaunchConfiguration {
	var tValues []*LaunchConfiguration
	for _, v := range values {
		tValues = append(tValues, c.transformLaunchConfiguration(v))
	}
	return tValues
}

var LaunchConfigurationTables = []interface{}{
	&LaunchConfiguration{},
	&LaunchConfigurationBlockDeviceMapping{},
}

func (c *Client) launchConfigurations(gConfig interface{}) error {
	var config autoscaling.DescribeLaunchConfigurationsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(LaunchConfigurationTables...)
	for {
		output, err := c.svc.DescribeLaunchConfigurations(&config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformLaunchConfigurations(output.LaunchConfigurations))
		c.log.Info("Fetched resources", zap.String("resource", "auto_scaling.launch_configurations"), zap.Int("count", len(output.LaunchConfigurations)))
		if aws.StringValue(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
