package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type NatGateway struct {
	ID                   uint `gorm:"primarykey"`
	AccountID            string
	Region               string
	CreateTime           *time.Time
	DeleteTime           *time.Time
	FailureCode          *string
	FailureMessage       *string
	NatGatewayAddresses  []*NatGatewayAddress `gorm:"constraint:OnDelete:CASCADE;"`
	NatGatewayId         *string
	ProvisionedBandwidth *ec2.ProvisionedBandwidth `gorm:"embedded;embeddedPrefix:provisioned_bandwidth_"`
	State                *string
	SubnetId             *string
	Tags                 []*NatGatewayTag `gorm:"constraint:OnDelete:CASCADE;"`
	VpcId                *string
}

type NatGatewayAddress struct {
	ID                 uint `gorm:"primarykey"`
	NatGatewayID       uint
	AllocationId       *string
	NetworkInterfaceId *string
	PrivateIp          *string
	PublicIp           *string
}

type NatGatewayTag struct {
	ID           uint `gorm:"primarykey"`
	NatGatewayID uint
	Key          *string
	Value        *string
}

func (c *Client) transformNatGatewayAddress(value *ec2.NatGatewayAddress) *NatGatewayAddress {
	return &NatGatewayAddress{
		AllocationId:       value.AllocationId,
		NetworkInterfaceId: value.NetworkInterfaceId,
		PrivateIp:          value.PrivateIp,
		PublicIp:           value.PublicIp,
	}
}

func (c *Client) transformNatGatewayAddresss(values []*ec2.NatGatewayAddress) []*NatGatewayAddress {
	var tValues []*NatGatewayAddress
	for _, v := range values {
		tValues = append(tValues, c.transformNatGatewayAddress(v))
	}
	return tValues
}

func (c *Client) transformNatGatewayTag(value *ec2.Tag) *NatGatewayTag {
	return &NatGatewayTag{
		Key:   value.Key,
		Value: value.Value,
	}
}

func (c *Client) transformNatGatewayTags(values []*ec2.Tag) []*NatGatewayTag {
	var tValues []*NatGatewayTag
	for _, v := range values {
		tValues = append(tValues, c.transformNatGatewayTag(v))
	}
	return tValues
}

func (c *Client) transformNatGateway(value *ec2.NatGateway) *NatGateway {
	return &NatGateway{
		Region:               c.region,
		AccountID:            c.accountID,
		CreateTime:           value.CreateTime,
		DeleteTime:           value.DeleteTime,
		FailureCode:          value.FailureCode,
		FailureMessage:       value.FailureMessage,
		NatGatewayAddresses:  c.transformNatGatewayAddresss(value.NatGatewayAddresses),
		NatGatewayId:         value.NatGatewayId,
		ProvisionedBandwidth: value.ProvisionedBandwidth,
		State:                value.State,
		SubnetId:             value.SubnetId,
		Tags:                 c.transformNatGatewayTags(value.Tags),
		VpcId:                value.VpcId,
	}
}

func (c *Client) transformNatGateways(values []*ec2.NatGateway) []*NatGateway {
	var tValues []*NatGateway
	for _, v := range values {
		tValues = append(tValues, c.transformNatGateway(v))
	}
	return tValues
}

func (c *Client) natGateways(gConfig interface{}) error {
	var config ec2.DescribeNatGatewaysInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["ec2NatGateway"] {
		err := c.db.AutoMigrate(
			&NatGateway{},
			&NatGatewayAddress{},
			&NatGatewayTag{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["ec2NatGateway"] = true
	}
	for {
		output, err := c.svc.DescribeNatGateways(&config)
		if err != nil {
			return err
		}
		c.log.Debug("deleting previous NatGateways", zap.String("region", c.region), zap.String("account_id", c.accountID))
		c.db.Where("region = ?", c.region).Where("account_id = ?", c.accountID).Delete(&NatGateway{})
		common.ChunkedCreate(c.db, c.transformNatGateways(output.NatGateways))
		c.log.Info("populating NatGateways", zap.Int("count", len(output.NatGateways)))
		if aws.StringValue(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
