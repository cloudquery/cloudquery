package ec2

import (
	"github.com/aws/aws-sdk-go/service/ec2"
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
	EbsIops                *int64
	EbsKmsKeyId            *string
	EbsSnapshotId          *string
	EbsVolumeSize          *int64
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

func (c *Client) transformImageBlockDeviceMapping(value *ec2.BlockDeviceMapping) *ImageBlockDeviceMapping {
	res := ImageBlockDeviceMapping{
		DeviceName:  value.DeviceName,
		AccountID:   c.accountID,
		Region:      c.region,
		NoDevice:    value.NoDevice,
		VirtualName: value.VirtualName,
	}

	if value.Ebs != nil {
		res.EbsDeleteOnTermination = value.Ebs.DeleteOnTermination
		res.EbsEncrypted = value.Ebs.Encrypted
		res.EbsIops = value.Ebs.Iops
		res.EbsKmsKeyId = value.Ebs.KmsKeyId
		res.EbsSnapshotId = value.Ebs.SnapshotId
		res.EbsVolumeSize = value.Ebs.VolumeSize
		res.EbsVolumeType = value.Ebs.VolumeType
	}

	return &res
}

func (c *Client) transformImageBlockDeviceMappings(values []*ec2.BlockDeviceMapping) []*ImageBlockDeviceMapping {
	var tValues []*ImageBlockDeviceMapping
	for _, v := range values {
		tValues = append(tValues, c.transformImageBlockDeviceMapping(v))
	}
	return tValues
}

func (c *Client) transformImageProductCode(value *ec2.ProductCode) *ImageProductCode {
	return &ImageProductCode{
		AccountID:       c.accountID,
		Region:          c.region,
		ProductCodeId:   value.ProductCodeId,
		ProductCodeType: value.ProductCodeType,
	}
}

func (c *Client) transformImageProductCodes(values []*ec2.ProductCode) []*ImageProductCode {
	var tValues []*ImageProductCode
	for _, v := range values {
		tValues = append(tValues, c.transformImageProductCode(v))
	}
	return tValues
}

func (c *Client) transformImageTag(value *ec2.Tag) *ImageTag {
	return &ImageTag{
		AccountID: c.accountID,
		Region:    c.region,
		Key:       value.Key,
		Value:     value.Value,
	}
}

func (c *Client) transformImageTags(values []*ec2.Tag) []*ImageTag {
	var tValues []*ImageTag
	for _, v := range values {
		tValues = append(tValues, c.transformImageTag(v))
	}
	return tValues
}

func (c *Client) transformImage(value *ec2.Image) *Image {
	res := Image{
		Region:              c.region,
		AccountID:           c.accountID,
		Architecture:        value.Architecture,
		BlockDeviceMappings: c.transformImageBlockDeviceMappings(value.BlockDeviceMappings),
		CreationDate:        value.CreationDate,
		Description:         value.Description,
		EnaSupport:          value.EnaSupport,
		Hypervisor:          value.Hypervisor,
		ImageId:             value.ImageId,
		ImageLocation:       value.ImageLocation,
		ImageOwnerAlias:     value.ImageOwnerAlias,
		ImageType:           value.ImageType,
		KernelId:            value.KernelId,
		Name:                value.Name,
		OwnerId:             value.OwnerId,
		Platform:            value.Platform,
		PlatformDetails:     value.PlatformDetails,
		ProductCodes:        c.transformImageProductCodes(value.ProductCodes),
		Public:              value.Public,
		RamdiskId:           value.RamdiskId,
		RootDeviceName:      value.RootDeviceName,
		RootDeviceType:      value.RootDeviceType,
		SriovNetSupport:     value.SriovNetSupport,
		State:               value.State,
		Tags:                c.transformImageTags(value.Tags),
		UsageOperation:      value.UsageOperation,
		VirtualizationType:  value.VirtualizationType,
	}

	if value.StateReason != nil {
		res.StateReasonCode = value.StateReason.Code
		res.StateReasonMessage = value.StateReason.Message
	}
	return &res
}

func (c *Client) transformImages(values []*ec2.Image) []*Image {
	var tValues []*Image
	for _, v := range values {
		tValues = append(tValues, c.transformImage(v))
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
	var config ec2.DescribeImagesInput
	err := mapstructure.Decode(gConfig, &config)
	if config.Owners == nil {
		self := "self"
		config.Owners = append(config.Owners, &self)
	}
	if err != nil {
		return err
	}

	output, err := c.svc.DescribeImages(&config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(ImageTables...)
	c.db.ChunkedCreate(c.transformImages(output.Images))
	c.log.Info("Fetched resources", zap.String("resource", "ec2.images"), zap.Int("count", len(output.Images)))
	return nil
}
