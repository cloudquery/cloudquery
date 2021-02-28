package ec2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/mitchellh/mapstructure"
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

func (c *Client) transformVpcCidrBlockAssociations(values *[]types.VpcCidrBlockAssociation) []*VpcCidrBlockAssociation {
	var tValues []*VpcCidrBlockAssociation
	for _, value := range *values {
		res := VpcCidrBlockAssociation{
			AccountID:     c.accountID,
			Region:        c.region,
			AssociationId: value.AssociationId,
			CidrBlock:     value.CidrBlock,
		}
		if value.CidrBlockState != nil {
			res.State = aws.String(string(value.CidrBlockState.State))
			res.StatusMessage = aws.String(string(value.CidrBlockState.State))
		}
		tValues = append(tValues, &res)
	}
	return tValues
}

func (c *Client) transformVpcIpv6CidrBlockAssociations(values *[]types.VpcIpv6CidrBlockAssociation) []*VpcIpv6CidrBlockAssociation {
	var tValues []*VpcIpv6CidrBlockAssociation
	for _, value := range *values {
		res := VpcIpv6CidrBlockAssociation{
			AccountID:          c.accountID,
			Region:             c.region,
			AssociationId:      value.AssociationId,
			Ipv6CidrBlock:      value.Ipv6CidrBlock,
			Ipv6Pool:           value.Ipv6Pool,
			NetworkBorderGroup: value.NetworkBorderGroup,
		}

		if value.Ipv6CidrBlockState != nil {
			res.State = aws.String(string(value.Ipv6CidrBlockState.State))
			res.StatusMessage = value.Ipv6CidrBlockState.StatusMessage
		}
		tValues = append(tValues, &res)
	}
	return tValues
}

func (c *Client) transformVpcTags(values *[]types.Tag) []*VpcTag {
	var tValues []*VpcTag
	for _, value := range *values {
		tValues = append(tValues, &VpcTag{
			AccountID: c.accountID,
			Region:    c.region,
			Key:       value.Key,
			Value:     value.Value,
		})
	}
	return tValues
}

func (c *Client) transformVpcs(values *[]types.Vpc) []*Vpc {
	var tValues []*Vpc
	for _, value := range *values {
		tValues = append(tValues, &Vpc{
			Region:                    c.region,
			AccountID:                 c.accountID,
			CidrBlock:                 value.CidrBlock,
			CidrBlockAssociations:     c.transformVpcCidrBlockAssociations(&value.CidrBlockAssociationSet),
			DhcpOptionsId:             value.DhcpOptionsId,
			InstanceTenancy:           aws.String(string(value.InstanceTenancy)),
			Ipv6CidrBlockAssociations: c.transformVpcIpv6CidrBlockAssociations(&value.Ipv6CidrBlockAssociationSet),
			IsDefault:                 &value.IsDefault,
			OwnerId:                   value.OwnerId,
			State:                     aws.String(string(value.State)),
			Tags:                      c.transformVpcTags(&value.Tags),
			VpcId:                     value.VpcId,
		})
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
	ctx := context.Background()
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(VPCTables...)
	for {
		output, err := c.svc.DescribeVpcs(ctx, &config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformVpcs(&output.Vpcs))
		c.log.Info("Fetched resources","resource", "ec2.vpcs", "count", len(output.Vpcs))
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
