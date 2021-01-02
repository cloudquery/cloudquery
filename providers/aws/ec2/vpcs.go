package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

type Vpc struct {
	_                         interface{} `neo:"raw:MERGE (a:AWSAccount {account_id: $account_id}) MERGE (a) - [:Resource] -> (n)"`
	ID                        uint        `gorm:"primarykey"`
	AccountID                 string      `neo:"unique"`
	Region                    string      `neo:"unique"`
	CidrBlock                 *string
	CidrBlockAssociations     []*VpcCidrBlockAssociation `gorm:"constraint:OnDelete:CASCADE;"`
	DhcpOptionsId             *string
	InstanceTenancy           *string
	Ipv6CidrBlockAssociations []*VpcIpv6CidrBlockAssociation `gorm:"constraint:OnDelete:CASCADE;"`
	IsDefault                 *bool
	OwnerId                   *string
	State                     *string
	Tags                      []*VpcTag `gorm:"constraint:OnDelete:CASCADE;"`
	VpcId                     *string   `neo:"unique"`
}

func (Vpc) TableName() string {
	return "aws_ec2_vpcs"
}

type VpcCidrBlockAssociation struct {
	ID    uint `gorm:"primarykey"`
	VpcID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	AssociationId *string
	CidrBlock     *string

	State         *string
	StatusMessage *string
}

func (VpcCidrBlockAssociation) TableName() string {
	return "aws_ec2_vpc_cidr_block_associations"
}

type VpcIpv6CidrBlockAssociation struct {
	ID    uint `gorm:"primarykey"`
	VpcID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	AssociationId *string
	Ipv6CidrBlock *string

	State         *string
	StatusMessage *string

	Ipv6Pool           *string
	NetworkBorderGroup *string
}

func (VpcIpv6CidrBlockAssociation) TableName() string {
	return "aws_ec2_vpc_ipv6_cidr_block_associations"
}

type VpcTag struct {
	ID    uint `gorm:"primarykey"`
	VpcID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Key   *string
	Value *string
}

func (VpcTag) TableName() string {
	return "aws_ec2_vpc_tags"
}

func (c *Client) transformVpcCidrBlockAssociation(value *ec2.VpcCidrBlockAssociation) *VpcCidrBlockAssociation {
	res := VpcCidrBlockAssociation{
		AccountID:     c.accountID,
		Region:        c.region,
		AssociationId: value.AssociationId,
		CidrBlock:     value.CidrBlock,
	}
	if value.CidrBlockState != nil {
		res.State = value.CidrBlockState.State
		res.StatusMessage = value.CidrBlockState.State
	}

	return &res
}

func (c *Client) transformVpcCidrBlockAssociations(values []*ec2.VpcCidrBlockAssociation) []*VpcCidrBlockAssociation {
	var tValues []*VpcCidrBlockAssociation
	for _, v := range values {
		tValues = append(tValues, c.transformVpcCidrBlockAssociation(v))
	}
	return tValues
}

func (c *Client) transformVpcIpv6CidrBlockAssociation(value *ec2.VpcIpv6CidrBlockAssociation) *VpcIpv6CidrBlockAssociation {
	res := VpcIpv6CidrBlockAssociation{
		AccountID:          c.accountID,
		Region:             c.region,
		AssociationId:      value.AssociationId,
		Ipv6CidrBlock:      value.Ipv6CidrBlock,
		Ipv6Pool:           value.Ipv6Pool,
		NetworkBorderGroup: value.NetworkBorderGroup,
	}

	if value.Ipv6CidrBlockState != nil {
		res.State = value.Ipv6CidrBlockState.State
		res.StatusMessage = value.Ipv6CidrBlockState.StatusMessage
	}

	return &res
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
		AccountID: c.accountID,
		Region:    c.region,
		Key:       value.Key,
		Value:     value.Value,
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
		Region:                    c.region,
		AccountID:                 c.accountID,
		CidrBlock:                 value.CidrBlock,
		CidrBlockAssociations:     c.transformVpcCidrBlockAssociations(value.CidrBlockAssociationSet),
		DhcpOptionsId:             value.DhcpOptionsId,
		InstanceTenancy:           value.InstanceTenancy,
		Ipv6CidrBlockAssociations: c.transformVpcIpv6CidrBlockAssociations(value.Ipv6CidrBlockAssociationSet),
		IsDefault:                 value.IsDefault,
		OwnerId:                   value.OwnerId,
		State:                     value.State,
		Tags:                      c.transformVpcTags(value.Tags),
		VpcId:                     value.VpcId,
	}
}

func (c *Client) transformVpcs(values []*ec2.Vpc) []*Vpc {
	var tValues []*Vpc
	for _, v := range values {
		tValues = append(tValues, c.transformVpc(v))
	}
	return tValues
}

var VPCTables = []interface{}{
	&Vpc{},
	&VpcCidrBlockAssociation{},
	&VpcIpv6CidrBlockAssociation{},
	&VpcTag{},
}

func (c *Client) vpcs(gConfig interface{}) error {
	var config ec2.DescribeVpcsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(VPCTables...)
	for {
		output, err := c.svc.DescribeVpcs(&config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformVpcs(output.Vpcs))
		c.log.Info("Fetched resources", zap.String("resource", "ec2.vpcs"), zap.Int("count", len(output.Vpcs)))
		if aws.StringValue(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
