package ec2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
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

func (c *Client) transformNatGatewayAddresss(values *[]types.NatGatewayAddress) []*NatGatewayAddress {
	var tValues []*NatGatewayAddress
	for _, value := range *values {
		tValues = append(tValues, &NatGatewayAddress{
			AccountID:          c.accountID,
			Region:             c.region,
			AllocationId:       value.AllocationId,
			NetworkInterfaceId: value.NetworkInterfaceId,
			PrivateIp:          value.PrivateIp,
			PublicIp:           value.PublicIp,
		})
	}
	return tValues
}

func (c *Client) transformNatGatewayTags(values *[]types.Tag) []*NatGatewayTag {
	var tValues []*NatGatewayTag
	for _, value := range *values {
		tValues = append(tValues, &NatGatewayTag{
			AccountID: c.accountID,
			Region:    c.region,
			Key:       value.Key,
			Value:     value.Value,
		})
	}
	return tValues
}

func (c *Client) transformNatGateways(values *[]types.NatGateway) []*NatGateway {
	var tValues []*NatGateway
	for _, value := range *values {
		res := NatGateway{
			Region:              c.region,
			AccountID:           c.accountID,
			CreateTime:          value.CreateTime,
			DeleteTime:          value.DeleteTime,
			FailureCode:         value.FailureCode,
			FailureMessage:      value.FailureMessage,
			NatGatewayAddresses: c.transformNatGatewayAddresss(&value.NatGatewayAddresses),
			NatGatewayId:        value.NatGatewayId,
			State:               aws.String(string(value.State)),
			SubnetId:            value.SubnetId,
			Tags:                c.transformNatGatewayTags(&value.Tags),
			VpcId:               value.VpcId,
		}

		if value.ProvisionedBandwidth != nil {
			res.ProvisionTime = value.ProvisionedBandwidth.ProvisionTime
			res.Provisioned = value.ProvisionedBandwidth.Provisioned
			res.ProvisionedRequestTime = value.ProvisionedBandwidth.RequestTime
			res.ProvisionedRequested = value.ProvisionedBandwidth.Requested
			res.ProvisionedStatus = value.ProvisionedBandwidth.Status
		}
		tValues = append(tValues, &res)
	}
	return tValues
}

var NatGatewayTables = []interface{}{
	&NatGateway{},
	&NatGatewayAddress{},
	&NatGatewayTag{},
}

func (c *Client) natGateways(gConfig interface{}) error {
	ctx := context.Background()
	var config ec2.DescribeNatGatewaysInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(NatGatewayTables...)
	for {
		output, err := c.svc.DescribeNatGateways(ctx, &config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformNatGateways(&output.NatGateways))
		c.log.Info("Fetched resources", zap.String("resource", "ec2.nat_gateways"), zap.Int("count", len(output.NatGateways)))
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
