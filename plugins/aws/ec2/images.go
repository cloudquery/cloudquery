package ec2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

type Image struct {
	_                   interface{} `neo:"raw:MERGE (a:AWSAccount {account_id: $account_id}) MERGE (a) - [:Resource] -> (n)"`
	ID                  uint        `gorm:"primarykey"`
	AccountID           string      `neo:"unique"`
	Region              string      `neo:"unique"`
	Architecture        *string
	BlockDeviceMappings []*ImageBlockDeviceMapping `gorm:"constraint:OnDelete:CASCADE;"`
	CreationDate        *string
	Description         *string
	EnaSupport          *bool
	Hypervisor          *string
	ImageId             *string `neo:"unique"`
	ImageLocation       *string
	ImageOwnerAlias     *string
	ImageType           *string
	KernelId            *string
	Name                *string
	OwnerId             *string
	Platform            *string
	PlatformDetails     *string
	ProductCodes        []*ImageProductCode `gorm:"constraint:OnDelete:CASCADE;"`
	Public              *bool
	RamdiskId           *string
	RootDeviceName      *string
	RootDeviceType      *string
	SriovNetSupport     *string
	State               *string

	StateReasonCode    *string
	StateReasonMessage *string

	Tags               []*ImageTag `gorm:"constraint:OnDelete:CASCADE;"`
	UsageOperation     *string
	VirtualizationType *string
}

func (Image) TableName() string {
	return "aws_ec2_images"
}

type ImageBlockDeviceMapping struct {
	ID      uint `gorm:"primarykey"`
	ImageID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	DeviceName             *string
	EbsDeleteOnTermination *bool
	EbsEncrypted           *bool
	EbsIops                *int32
	EbsKmsKeyId            *string
	EbsSnapshotId          *string
	EbsVolumeSize          *int32
	EbsVolumeType          *string

	NoDevice    *string
	VirtualName *string
}

func (ImageBlockDeviceMapping) TableName() string {
	return "aws_ec2_image_block_device_mappings"
}

type ImageProductCode struct {
	ID      uint `gorm:"primarykey"`
	ImageID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	ProductCodeId   *string
	ProductCodeType *string
}

func (ImageProductCode) TableName() string {
	return "aws_ec2_image_product_codes"
}

type ImageTag struct {
	ID      uint `gorm:"primarykey"`
	ImageID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Key   *string
	Value *string
}

func (ImageTag) TableName() string {
	return "aws_ec2_image_tags"
}

func (c *Client) transformImageBlockDeviceMappings(values *[]types.BlockDeviceMapping) []*ImageBlockDeviceMapping {
	var tValues []*ImageBlockDeviceMapping
	for _, value := range *values {
		res := ImageBlockDeviceMapping{
			DeviceName:  value.DeviceName,
			AccountID:   c.accountID,
			Region:      c.region,
			NoDevice:    value.NoDevice,
			VirtualName: value.VirtualName,
		}

		if value.Ebs != nil {
			res.EbsDeleteOnTermination = &value.Ebs.DeleteOnTermination
			res.EbsEncrypted = &value.Ebs.Encrypted
			res.EbsIops = &value.Ebs.Iops
			res.EbsKmsKeyId = value.Ebs.KmsKeyId
			res.EbsSnapshotId = value.Ebs.SnapshotId
			res.EbsVolumeSize = &value.Ebs.VolumeSize
			res.EbsVolumeType = aws.String(string(value.Ebs.VolumeType))
		}
		tValues = append(tValues, &res)
	}
	return tValues
}

func (c *Client) transformImageProductCodes(values *[]types.ProductCode) []*ImageProductCode {
	var tValues []*ImageProductCode
	for _, v := range *values {
		tValues = append(tValues, &ImageProductCode{
			AccountID:       c.accountID,
			Region:          c.region,
			ProductCodeId:   v.ProductCodeId,
			ProductCodeType: aws.String(string(v.ProductCodeType)),
		})
	}
	return tValues
}


func (c *Client) transformImageTags(values *[]types.Tag) []*ImageTag {
	var tValues []*ImageTag
	for _, v := range *values {
		tValues = append(tValues, &ImageTag{
			AccountID: c.accountID,
			Region:    c.region,
			Key:       v.Key,
			Value:     v.Value,
		})
	}
	return tValues
}

func (c *Client) transformImages(values *[]types.Image) []*Image {
	var tValues []*Image
	for _, value := range *values {
		res := Image{
			Region:              c.region,
			AccountID:           c.accountID,
			Architecture:        aws.String(string(value.Architecture)),
			BlockDeviceMappings: c.transformImageBlockDeviceMappings(&value.BlockDeviceMappings),
			CreationDate:        value.CreationDate,
			Description:         value.Description,
			EnaSupport:          &value.EnaSupport,
			Hypervisor:          aws.String(string(value.Hypervisor)),
			ImageId:             value.ImageId,
			ImageLocation:       value.ImageLocation,
			ImageOwnerAlias:     value.ImageOwnerAlias,
			ImageType:           aws.String(string(value.ImageType)),
			KernelId:            value.KernelId,
			Name:                value.Name,
			OwnerId:             value.OwnerId,
			Platform:            aws.String(string(value.Platform)),
			PlatformDetails:     value.PlatformDetails,
			ProductCodes:        c.transformImageProductCodes(&value.ProductCodes),
			Public:              &value.Public,
			RamdiskId:           value.RamdiskId,
			RootDeviceName:      value.RootDeviceName,
			RootDeviceType:      aws.String(string(value.RootDeviceType)),
			SriovNetSupport:     value.SriovNetSupport,
			State:               aws.String(string(value.State)),
			Tags:                c.transformImageTags(&value.Tags),
			UsageOperation:      value.UsageOperation,
			VirtualizationType:  aws.String(string(value.VirtualizationType)),
		}

		if value.StateReason != nil {
			res.StateReasonCode = value.StateReason.Code
			res.StateReasonMessage = value.StateReason.Message
		}
		tValues = append(tValues, &res)
	}
	return tValues
}

var ImageTables = []interface{}{
	&Image{},
	&ImageBlockDeviceMapping{},
	&ImageProductCode{},
	&ImageTag{},
}

func (c *Client) images(gConfig interface{}) error {
	ctx := context.Background()
	var config ec2.DescribeImagesInput
	err := mapstructure.Decode(gConfig, &config)
	if config.Owners == nil {
		config.Owners = append(config.Owners, "self")
	}
	if err != nil {
		return err
	}

	output, err := c.svc.DescribeImages(ctx, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(ImageTables...)
	c.db.ChunkedCreate(c.transformImages(&output.Images))
	c.log.Info("Fetched resources", zap.String("resource", "ec2.images"), zap.Int("count", len(output.Images)))
	return nil
}
