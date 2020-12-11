package ec2

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

type Image struct {
	ID                  uint `gorm:"primarykey"`
	AccountID           string
	Region              string
	Architecture        *string
	BlockDeviceMappings []*ImageBlockDeviceMapping `gorm:"constraint:OnDelete:CASCADE;"`
	CreationDate        *string
	Description         *string
	EnaSupport          *bool
	Hypervisor          *string
	ImageId             *string
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
	StateReason         *ec2.StateReason `gorm:"embedded;embeddedPrefix:state_reason_"`
	Tags                []*ImageTag      `gorm:"constraint:OnDelete:CASCADE;"`
	UsageOperation      *string
	VirtualizationType  *string
}

type ImageBlockDeviceMapping struct {
	ID          uint `gorm:"primarykey"`
	ImageID     uint
	DeviceName  *string
	Ebs         *ec2.EbsBlockDevice `gorm:"embedded;embeddedPrefix:ebs_"`
	NoDevice    *string
	VirtualName *string
}

type ImageProductCode struct {
	ID              uint `gorm:"primarykey"`
	ImageID         uint
	ProductCodeId   *string
	ProductCodeType *string
}

type ImageTag struct {
	ID      uint `gorm:"primarykey"`
	ImageID uint
	Key     *string
	Value   *string
}

func (c *Client) transformImageBlockDeviceMapping(value *ec2.BlockDeviceMapping) *ImageBlockDeviceMapping {
	return &ImageBlockDeviceMapping{
		DeviceName:  value.DeviceName,
		Ebs:         value.Ebs,
		NoDevice:    value.NoDevice,
		VirtualName: value.VirtualName,
	}
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
		Key:   value.Key,
		Value: value.Value,
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
	return &Image{
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
		StateReason:         value.StateReason,
		Tags:                c.transformImageTags(value.Tags),
		UsageOperation:      value.UsageOperation,
		VirtualizationType:  value.VirtualizationType,
	}
}

func (c *Client) transformImages(values []*ec2.Image) []*Image {
	var tValues []*Image
	for _, v := range values {
		tValues = append(tValues, c.transformImage(v))
	}
	return tValues
}

func (c *Client) images(gConfig interface{}) error {
	var config ec2.DescribeImagesInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["ec2Image"] {
		err := c.db.AutoMigrate(
			&Image{},
			&ImageBlockDeviceMapping{},
			&ImageProductCode{},
			&ImageTag{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["ec2Image"] = true
	}

	output, err := c.svc.DescribeImages(&config)
	if err != nil {
		return err
	}
	c.db.Where("region = ?", c.region).Where("account_id = ?", c.accountID).Delete(&Image{})
	common.ChunkedCreate(c.db, c.transformImages(output.Images))
	c.log.Info("Fetched resources", zap.Int("count", len(output.Images)))
	return nil
}
