package ec2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

type Subnet struct {
	ID                          uint   `gorm:"primarykey"`
	AccountID                   string `neo:"unique"`
	Region                      string `neo:"unique"`
	AssignIpv6AddressOnCreation *bool
	AvailabilityZone            *string
	AvailabilityZoneId          *string
	AvailableIpAddressCount     *int32
	CidrBlock                   *string
	CustomerOwnedIpv4Pool       *string
	DefaultForAz                *bool
	Ipv6CidrBlockAssociationSet []*SubnetIpv6CidrBlockAssociation `gorm:"constraint:OnDelete:CASCADE;"`
	MapCustomerOwnedIpOnLaunch  *bool
	MapPublicIpOnLaunch         *bool
	OutpostArn                  *string
	OwnerId                     *string
	State                       *string
	SubnetArn                   *string
	SubnetId                    *string      `neo:"unique"`
	Tags                        []*SubnetTag `gorm:"constraint:OnDelete:CASCADE;"`
	VpcId                       *string
}

func (Subnet) TableName() string {
	return "aws_ec2_subnets"
}

type SubnetIpv6CidrBlockAssociation struct {
	ID       uint `gorm:"primarykey"`
	SubnetID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	AssociationId *string
	Ipv6CidrBlock *string
	State         *string
	StatusMessage *string
}

func (SubnetIpv6CidrBlockAssociation) TableName() string {
	return "aws_ec2_subnet_ipv6_cidr_block_associations"
}

type SubnetTag struct {
	ID        uint   `gorm:"primarykey"`
	SubnetID  uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`
	Key       *string
	Value     *string
}

func (SubnetTag) TableName() string {
	return "aws_ec2_subnet_tags"
}

func (c *Client) transformSubnetIpv6CidrBlockAssociations(values *[]types.SubnetIpv6CidrBlockAssociation) []*SubnetIpv6CidrBlockAssociation {
	var tValues []*SubnetIpv6CidrBlockAssociation
	for _, value := range *values {
		res := SubnetIpv6CidrBlockAssociation{
			AccountID:     c.accountID,
			Region:        c.region,
			AssociationId: value.AssociationId,
			Ipv6CidrBlock: value.Ipv6CidrBlock,
		}
		if value.Ipv6CidrBlock != nil {
			res.State = aws.String(string(value.Ipv6CidrBlockState.State))
			res.StatusMessage = value.Ipv6CidrBlockState.StatusMessage
		}

		tValues = append(tValues, &res)
	}
	return tValues
}

func (c *Client) transformSubnetTags(values *[]types.Tag) []*SubnetTag {
	var tValues []*SubnetTag
	for _, value := range *values {
		tValues = append(tValues, &SubnetTag{
			AccountID: c.accountID,
			Region:    c.region,
			Key:       value.Key,
			Value:     value.Value,
		})
	}
	return tValues
}

func (c *Client) transformSubnets(values *[]types.Subnet) []*Subnet {
	var tValues []*Subnet
	for _, value := range *values {
		tValues = append(tValues, &Subnet{
			Region:                      c.region,
			AccountID:                   c.accountID,
			AssignIpv6AddressOnCreation: &value.AssignIpv6AddressOnCreation,
			AvailabilityZone:            value.AvailabilityZone,
			AvailabilityZoneId:          value.AvailabilityZoneId,
			AvailableIpAddressCount:     &value.AvailableIpAddressCount,
			CidrBlock:                   value.CidrBlock,
			CustomerOwnedIpv4Pool:       value.CustomerOwnedIpv4Pool,
			DefaultForAz:                &value.DefaultForAz,
			Ipv6CidrBlockAssociationSet: c.transformSubnetIpv6CidrBlockAssociations(&value.Ipv6CidrBlockAssociationSet),
			MapCustomerOwnedIpOnLaunch:  &value.MapCustomerOwnedIpOnLaunch,
			MapPublicIpOnLaunch:         &value.MapPublicIpOnLaunch,
			OutpostArn:                  value.OutpostArn,
			OwnerId:                     value.OwnerId,
			State:                       aws.String(string(value.State)),
			SubnetArn:                   value.SubnetArn,
			SubnetId:                    value.SubnetId,
			Tags:                        c.transformSubnetTags(&value.Tags),
			VpcId:                       value.VpcId,
		})
	}
	return tValues
}

var SubnetTables = []interface{}{
	&Subnet{},
	&SubnetIpv6CidrBlockAssociation{},
	&SubnetTag{},
}

func (c *Client) subnets(gConfig interface{}) error {
	ctx := context.Background()
	var config ec2.DescribeSubnetsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(SubnetTables...)
	for {
		output, err := c.svc.DescribeSubnets(ctx, &config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformSubnets(&output.Subnets))
		c.log.Info("Fetched resources", zap.String("resource", "ec2.subnets"), zap.Int("count", len(output.Subnets)))
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
