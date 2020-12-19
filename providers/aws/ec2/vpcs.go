package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Vpc struct {
	ID                          uint `gorm:"primarykey"`
	AccountID                   string
	Region                      string
	CidrBlock                   *string
	CidrBlockAssociationSet     []*VpcCidrBlockAssociation `gorm:"constraint:OnDelete:CASCADE;"`
	DhcpOptionsId               *string
	InstanceTenancy             *string
	Ipv6CidrBlockAssociationSet []*VpcIpv6CidrBlockAssociation `gorm:"constraint:OnDelete:CASCADE;"`
	IsDefault                   *bool
	OwnerId                     *string
	State                       *string
	Tags                        []*VpcTag `gorm:"constraint:OnDelete:CASCADE;"`
	VpcId                       *string
}

func (Vpc) TableName() string {
	return "aws_ec2_vpcs"
}

type VpcCidrBlockAssociation struct {
	ID             uint `gorm:"primarykey"`
	VpcID          uint
	AssociationId  *string
	CidrBlock      *string
	CidrBlockState *ec2.VpcCidrBlockState `gorm:"embedded;embeddedPrefix:cidr_block_state_"`
}

func (VpcCidrBlockAssociation) TableName() string {
	return "aws_ec2_vpc_cidr_block_associations"
}

type VpcIpv6CidrBlockAssociation struct {
	ID                 uint `gorm:"primarykey"`
	VpcID              uint
	AssociationId      *string
	Ipv6CidrBlock      *string
	Ipv6CidrBlockState *ec2.VpcCidrBlockState `gorm:"embedded;embeddedPrefix:ipv_6_cidr_block_state_"`
	Ipv6Pool           *string
	NetworkBorderGroup *string
}

func (VpcIpv6CidrBlockAssociation) TableName() string {
	return "aws_ec2_vpc_ipv6_cidr_block_associations"
}

type VpcTag struct {
	ID    uint `gorm:"primarykey"`
	VpcID uint
	Key   *string
	Value *string
}

func (VpcTag) TableName() string {
	return "aws_ec2_vpc_tags"
}

func (c *Client) transformVpcCidrBlockAssociation(value *ec2.VpcCidrBlockAssociation) *VpcCidrBlockAssociation {
	return &VpcCidrBlockAssociation{
		AssociationId:  value.AssociationId,
		CidrBlock:      value.CidrBlock,
		CidrBlockState: value.CidrBlockState,
	}
}

func (c *Client) transformVpcCidrBlockAssociations(values []*ec2.VpcCidrBlockAssociation) []*VpcCidrBlockAssociation {
	var tValues []*VpcCidrBlockAssociation
	for _, v := range values {
		tValues = append(tValues, c.transformVpcCidrBlockAssociation(v))
	}
	return tValues
}

func (c *Client) transformVpcIpv6CidrBlockAssociation(value *ec2.VpcIpv6CidrBlockAssociation) *VpcIpv6CidrBlockAssociation {
	return &VpcIpv6CidrBlockAssociation{
		AssociationId:      value.AssociationId,
		Ipv6CidrBlock:      value.Ipv6CidrBlock,
		Ipv6CidrBlockState: value.Ipv6CidrBlockState,
		Ipv6Pool:           value.Ipv6Pool,
		NetworkBorderGroup: value.NetworkBorderGroup,
	}
}

func (c *Client) transformVpcIpv6CidrBlockAssociations(values []*ec2.VpcIpv6CidrBlockAssociation) []*VpcIpv6CidrBlockAssociation {
	var tValues []*VpcIpv6CidrBlockAssociation
	for _, v := range values {
		tValues = append(tValues, c.transformVpcIpv6CidrBlockAssociation(v))
	}
	return tValues
}

func (c *Client) transformVpcTag(value *ec2.Tag) *VpcTag {
	return &VpcTag{
		Key:   value.Key,
		Value: value.Value,
	}
}

func (c *Client) transformVpcTags(values []*ec2.Tag) []*VpcTag {
	var tValues []*VpcTag
	for _, v := range values {
		tValues = append(tValues, c.transformVpcTag(v))
	}
	return tValues
}

func (c *Client) transformVpc(value *ec2.Vpc) *Vpc {
	return &Vpc{
		Region:                      c.region,
		AccountID:                   c.accountID,
		CidrBlock:                   value.CidrBlock,
		CidrBlockAssociationSet:     c.transformVpcCidrBlockAssociations(value.CidrBlockAssociationSet),
		DhcpOptionsId:               value.DhcpOptionsId,
		InstanceTenancy:             value.InstanceTenancy,
		Ipv6CidrBlockAssociationSet: c.transformVpcIpv6CidrBlockAssociations(value.Ipv6CidrBlockAssociationSet),
		IsDefault:                   value.IsDefault,
		OwnerId:                     value.OwnerId,
		State:                       value.State,
		Tags:                        c.transformVpcTags(value.Tags),
		VpcId:                       value.VpcId,
	}
}

func (c *Client) transformVpcs(values []*ec2.Vpc) []*Vpc {
	var tValues []*Vpc
	for _, v := range values {
		tValues = append(tValues, c.transformVpc(v))
	}
	return tValues
}

func MigrateVPCs(db *gorm.DB) error {
	return db.AutoMigrate(
		&Vpc{},
		&VpcCidrBlockAssociation{},
		&VpcIpv6CidrBlockAssociation{},
		&VpcTag{},
	)
}

func (c *Client) vpcs(gConfig interface{}) error {
	var config ec2.DescribeVpcsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	for {
		output, err := c.svc.DescribeVpcs(&config)
		if err != nil {
			return err
		}
		c.db.Where("region = ?", c.region).Where("account_id = ?", c.accountID).Delete(&Vpc{})
		common.ChunkedCreate(c.db, c.transformVpcs(output.Vpcs))
		c.log.Info("Fetched resources", zap.String("resource", "ec2.vpcs"), zap.Int("count", len(output.Vpcs)))
		if aws.StringValue(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
