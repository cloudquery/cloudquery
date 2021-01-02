package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type NatGateway struct {
	_                   interface{} `neo:"raw:MERGE (a:AWSAccount {account_id: $account_id}) MERGE (a) - [:Resource] -> (n)"`
	ID                  uint        `gorm:"primarykey"`
	AccountID           string      `neo:"unique"`
	Region              string      `neo:"unique"`
	CreateTime          *time.Time
	DeleteTime          *time.Time
	FailureCode         *string
	FailureMessage      *string
	NatGatewayAddresses []*NatGatewayAddress `gorm:"constraint:OnDelete:CASCADE;"`
	NatGatewayId        *string              `neo:"unique"`

	ProvisionTime          *time.Time
	Provisioned            *string
	ProvisionedRequestTime *time.Time
	ProvisionedRequested   *string
	ProvisionedStatus      *string

	State    *string
	SubnetId *string
	Tags     []*NatGatewayTag `gorm:"constraint:OnDelete:CASCADE;"`
	VpcId    *string
}

func (NatGateway) TableName() string {
	return "aws_ec2_nat_gateways"
}

type NatGatewayAddress struct {
	ID           uint `gorm:"primarykey"`
	NatGatewayID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	AllocationId       *string
	NetworkInterfaceId *string
	PrivateIp          *string
	PublicIp           *string
}

func (NatGatewayAddress) TableName() string {
	return "aws_ec2_nat_gateway_addresses"
}

type NatGatewayTag struct {
	ID           uint `gorm:"primarykey"`
	NatGatewayID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Key   *string
	Value *string
}

func (NatGatewayTag) TableName() string {
	return "aws_ec2_nat_gateway_tags"
}

func (c *Client) transformNatGatewayAddress(value *ec2.NatGatewayAddress) *NatGatewayAddress {
	return &NatGatewayAddress{
		AccountID:          c.accountID,
		Region:             c.region,
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
		AccountID: c.accountID,
		Region:    c.region,
		Key:       value.Key,
		Value:     value.Value,
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
	res := NatGateway{
		Region:              c.region,
		AccountID:           c.accountID,
		CreateTime:          value.CreateTime,
		DeleteTime:          value.DeleteTime,
		FailureCode:         value.FailureCode,
		FailureMessage:      value.FailureMessage,
		NatGatewayAddresses: c.transformNatGatewayAddresss(value.NatGatewayAddresses),
		NatGatewayId:        value.NatGatewayId,
		State:               value.State,
		SubnetId:            value.SubnetId,
		Tags:                c.transformNatGatewayTags(value.Tags),
		VpcId:               value.VpcId,
	}

	if value.ProvisionedBandwidth != nil {
		res.ProvisionTime = value.ProvisionedBandwidth.ProvisionTime
		res.Provisioned = value.ProvisionedBandwidth.Provisioned
		res.ProvisionedRequestTime = value.ProvisionedBandwidth.RequestTime
		res.ProvisionedRequested = value.ProvisionedBandwidth.Requested
		res.ProvisionedStatus = value.ProvisionedBandwidth.Status
	}

	return &res
}

func (c *Client) transformNatGateways(values []*ec2.NatGateway) []*NatGateway {
	var tValues []*NatGateway
	for _, v := range values {
		tValues = append(tValues, c.transformNatGateway(v))
	}
	return tValues
}

var NatGatewayTables = []interface{}{
	&NatGateway{},
	&NatGatewayAddress{},
	&NatGatewayTag{},
}

func (c *Client) natGateways(gConfig interface{}) error {
	var config ec2.DescribeNatGatewaysInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(NatGatewayTables...)
	for {
		output, err := c.svc.DescribeNatGateways(&config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformNatGateways(output.NatGateways))
		c.log.Info("Fetched resources", zap.String("resource", "ec2.nat_gateways"), zap.Int("count", len(output.NatGateways)))
		if aws.StringValue(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
