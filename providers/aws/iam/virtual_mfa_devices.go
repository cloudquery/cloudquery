package iam

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type VirtualMFADevice struct {
	ID                       uint    `gorm:"primarykey"`
	AccountID                string  `neo:"unique"`
	SerialNumber *string	`neo:"unique"`
	UserARN *string
	EnableDate               *time.Time
}

func (VirtualMFADevice) TableName() string {
	return "aws_iam_virtual_mfa_devices"
}

var VirtualMFADeviceTables = []interface{}{
	&VirtualMFADevice{},
}

func (c *Client) transformMFADevices(values []*iam.VirtualMFADevice) []*VirtualMFADevice {
	var tValues []*VirtualMFADevice
	for _, v := range values {
		tValues = append(tValues, &VirtualMFADevice{
			SerialNumber: v.SerialNumber,
			UserARN: 		v.User.Arn,
			EnableDate:   v.EnableDate,
		})
	}
	return tValues
}

func (c *Client) virtualMFADevices(gConfig interface{}) error {
	var config iam.ListVirtualMFADevicesInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("account_id", c.accountID).Delete(VirtualMFADeviceTables...)

	for {
		output, err := c.svc.ListVirtualMFADevices(&config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformMFADevices(output.VirtualMFADevices))
		c.log.Info("Fetched resources", zap.String("resource", "iam.virtual_mfa_devices"), zap.Int("count", len(output.VirtualMFADevices)))
		if aws.StringValue(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
