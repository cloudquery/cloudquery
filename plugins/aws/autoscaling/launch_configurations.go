package autoscaling

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	//"github.com/cloudquery/cq-provider-aws/common"
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

	MetadataHttpEndpoint            string
	MetadataHttpPutResponseHopLimit *int32
	MetadataHttpTokens              string

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
	EbsIops                *int32
	EbsSnapshotId          *string
	EbsVolumeSize          *int32
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


func (c *Client) transformLaunchConfigurationBlockDeviceMappings(values []types.BlockDeviceMapping) []*LaunchConfigurationBlockDeviceMapping {
	var tValues []*LaunchConfigurationBlockDeviceMapping
	for _, v := range values {
		tValue := LaunchConfigurationBlockDeviceMapping{
			AccountID:   c.accountID,
			Region:      c.region,
			DeviceName:  v.DeviceName,
			NoDevice:    v.NoDevice,
			VirtualName: v.VirtualName,
		}
		if v.Ebs != nil {
			tValue.EbsDeleteOnTermination = v.Ebs.DeleteOnTermination
			tValue.EbsEncrypted = v.Ebs.Encrypted
			tValue.EbsIops = v.Ebs.Iops
			tValue.EbsSnapshotId = v.Ebs.SnapshotId
			tValue.EbsVolumeSize = v.Ebs.VolumeSize
			tValue.EbsVolumeType = v.Ebs.VolumeType
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}


func (c *Client) transformLaunchConfigurations(values []types.LaunchConfiguration) []*LaunchConfiguration {
	var tValues []*LaunchConfiguration
	for _, v := range values {
		tValue := LaunchConfiguration{
			Region:                       c.region,
			AccountID:                    c.accountID,
			AssociatePublicIpAddress:     v.AssociatePublicIpAddress,
			BlockDeviceMappings:          c.transformLaunchConfigurationBlockDeviceMappings(v.BlockDeviceMappings),
			ClassicLinkVPCId:             v.ClassicLinkVPCId,
			//ClassicLinkVPCSecurityGroups: common.StringListToString(v.ClassicLinkVPCSecurityGroups),
			CreatedTime:                  v.CreatedTime,
			EbsOptimized:                 v.EbsOptimized,
			IamInstanceProfile:           v.IamInstanceProfile,
			ImageId:                      v.ImageId,
			InstanceType:                 v.InstanceType,
			KernelId:                     v.KernelId,
			KeyName:                      v.KeyName,
			LaunchConfigurationARN:       v.LaunchConfigurationARN,
			LaunchConfigurationName:      v.LaunchConfigurationName,
			PlacementTenancy:             v.PlacementTenancy,
			RamdiskId:                    v.RamdiskId,
			//SecurityGroups:               common.StringListToString(v.SecurityGroups),
			SpotPrice:                    v.SpotPrice,
			UserData:                     v.UserData,
		}

		if v.MetadataOptions != nil {
			tValue.MetadataHttpEndpoint = string(v.MetadataOptions.HttpEndpoint)
			tValue.MetadataHttpPutResponseHopLimit = v.MetadataOptions.HttpPutResponseHopLimit
			tValue.MetadataHttpTokens = string(v.MetadataOptions.HttpTokens)
		}

		if v.InstanceMonitoring != nil {
			tValue.InstanceMonitoringEnabled = v.InstanceMonitoring.Enabled
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

var LaunchConfigurationTables = []interface{}{
	&LaunchConfiguration{},
	&LaunchConfigurationBlockDeviceMapping{},
}

func (c *Client) launchConfigurations(gConfig interface{}) error {
	var config autoscaling.DescribeLaunchConfigurationsInput
	ctx := context.Background()
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(LaunchConfigurationTables...)
	for {
		output, err := c.svc.DescribeLaunchConfigurations(ctx, &config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformLaunchConfigurations(output.LaunchConfigurations))
		c.log.Info("Fetched resources", zap.String("resource", "auto_scaling.launch_configurations"), zap.Int("count", len(output.LaunchConfigurations)))
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
