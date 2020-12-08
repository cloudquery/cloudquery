package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

type Subnet struct {
	ID                                             uint `gorm:"primarykey"`
	AccountID                                    string
	Region                                       string
	AssignIpv6AddressOnCreation                   *bool
	AvailabilityZone                            *string
	AvailabilityZoneId                          *string
	AvailableIpAddressCount                      *int64
	CidrBlock                                   *string
	CustomerOwnedIpv4Pool                       *string
	DefaultForAz                                  *bool
	Ipv6CidrBlockAssociationSet []*SubnetIpv6CidrBlockAssociation `gorm:"constraint:OnDelete:CASCADE;"`
	MapCustomerOwnedIpOnLaunch                    *bool
	MapPublicIpOnLaunch                           *bool
	OutpostArn                                  *string
	OwnerId                                     *string
	State                                       *string
	SubnetArn                                   *string
	SubnetId                                    *string
	Tags                                   []*SubnetTag `gorm:"constraint:OnDelete:CASCADE;"`
	VpcId                                       *string
}

type SubnetIpv6CidrBlockAssociation struct {
	ID                                             uint `gorm:"primarykey"`
	SubnetID                                       uint
	AssociationId                               *string
	Ipv6CidrBlock                               *string
	Ipv6CidrBlockState        *ec2.SubnetCidrBlockState `gorm:"embedded;embeddedPrefix:ipv_6_cidr_block_state_"`
}

type SubnetTag struct {
	ID                                             uint `gorm:"primarykey"`
	SubnetID                                       uint
	Key                                         *string
	Value                                       *string
}

func (c *Client) transformSubnetIpv6CidrBlockAssociation(value *ec2.SubnetIpv6CidrBlockAssociation) *SubnetIpv6CidrBlockAssociation {
	return &SubnetIpv6CidrBlockAssociation{
		AssociationId: value.AssociationId,
		Ipv6CidrBlock: value.Ipv6CidrBlock,
		Ipv6CidrBlockState: value.Ipv6CidrBlockState,
	}
}

func (c *Client) transformSubnetIpv6CidrBlockAssociations(values []*ec2.SubnetIpv6CidrBlockAssociation) []*SubnetIpv6CidrBlockAssociation {
	var tValues []*SubnetIpv6CidrBlockAssociation
	for _, v := range values {
		tValues = append(tValues, c.transformSubnetIpv6CidrBlockAssociation(v))
	}
	return tValues
}

func (c *Client) transformSubnetTag(value *ec2.Tag) *SubnetTag {
	return &SubnetTag{
		Key: value.Key,
		Value: value.Value,
	}
}

func (c *Client) transformSubnetTags(values []*ec2.Tag) []*SubnetTag {
	var tValues []*SubnetTag
	for _, v := range values {
		tValues = append(tValues, c.transformSubnetTag(v))
	}
	return tValues
}

func (c *Client) transformSubnet(value *ec2.Subnet) *Subnet {
	return &Subnet{
		Region: c.region,
		AccountID: c.accountID,
		AssignIpv6AddressOnCreation: value.AssignIpv6AddressOnCreation,
		AvailabilityZone: value.AvailabilityZone,
		AvailabilityZoneId: value.AvailabilityZoneId,
		AvailableIpAddressCount: value.AvailableIpAddressCount,
		CidrBlock: value.CidrBlock,
		CustomerOwnedIpv4Pool: value.CustomerOwnedIpv4Pool,
		DefaultForAz: value.DefaultForAz,
		Ipv6CidrBlockAssociationSet: c.transformSubnetIpv6CidrBlockAssociations(value.Ipv6CidrBlockAssociationSet),
		MapCustomerOwnedIpOnLaunch: value.MapCustomerOwnedIpOnLaunch,
		MapPublicIpOnLaunch: value.MapPublicIpOnLaunch,
		OutpostArn: value.OutpostArn,
		OwnerId: value.OwnerId,
		State: value.State,
		SubnetArn: value.SubnetArn,
		SubnetId: value.SubnetId,
		Tags: c.transformSubnetTags(value.Tags),
		VpcId: value.VpcId,
	}
}

func (c *Client) transformSubnets(values []*ec2.Subnet) []*Subnet {
	var tValues []*Subnet
	for _, v := range values {
		tValues = append(tValues, c.transformSubnet(v))
	}
	return tValues
}

func (c *Client)subnets(gConfig interface{}) error {
	var config ec2.DescribeSubnetsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["ec2Subnet"] {
		err := c.db.AutoMigrate(
			&Subnet{},
			&SubnetIpv6CidrBlockAssociation{},
			&SubnetTag{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["ec2Subnet"] = true
	}
	for {
		output, err := c.svc.DescribeSubnets(&config)
		if err != nil {
			return err
		}
		c.log.Debug("deleting previous Subnets", zap.String("region", c.region), zap.String("account_id", c.accountID))
		c.db.Where("region = ?", c.region).Where("account_id = ?", c.accountID).Delete(&Subnet{})
		common.ChunkedCreate(c.db, c.transformSubnets(output.Subnets))
		c.log.Info("populating Subnets", zap.Int("count", len(output.Subnets)))
		if aws.StringValue(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}


