package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type VpcPeeringConnection struct {
	ID        uint `gorm:"primarykey"`
	AccountID string
	Region    string

	AccCidrBlock        *string
	AccCidrBlockSet     []*VpcPeeringConnectionAccCidrBlock     `gorm:"constraint:OnDelete:CASCADE;"`
	AccIpv6CidrBlockSet []*VpcPeeringConnectionAccIpv6CidrBlock `gorm:"constraint:OnDelete:CASCADE;"`
	AccOwnerId          *string
	AccOptions          *ec2.VpcPeeringConnectionOptionsDescription `gorm:"embedded;embeddedPrefix:acc_option_"`
	AccRegion           *string
	AccVpcId            *string

	ExpirationTime *time.Time

	ReqCidrBlock        *string
	ReqCidrBlockSet     []*VpcPeeringConnectionReqCidrBlock     `gorm:"constraint:OnDelete:CASCADE;"`
	ReqIpv6CidrBlockSet []*VpcPeeringConnectionReqIpv6CidrBlock `gorm:"constraint:OnDelete:CASCADE;"`
	ReqOwnerId          *string
	ReqOptions          *ec2.VpcPeeringConnectionOptionsDescription `gorm:"embedded;embeddedPrefix:req_option_"`
	ReqRegion           *string
	ReqVpcId            *string

	Status                 *ec2.VpcPeeringConnectionStateReason `gorm:"embedded;embeddedPrefix:status_"`
	Tags                   []*VpcPeeringConnectionTag           `gorm:"constraint:OnDelete:CASCADE;"`
	VpcPeeringConnectionId *string
}

func (VpcPeeringConnection) TableName() string {
	return "aws_ec2_vpc_peering_connections"
}

type VpcPeeringConnectionAccCidrBlock struct {
	ID                     uint `gorm:"primarykey"`
	VpcPeeringConnectionID uint
	CidrBlock              *string
}

func (VpcPeeringConnectionAccCidrBlock) TableName() string {
	return "aws_ec2_vpc_peering_connection_acc_cidr_blocks"
}

type VpcPeeringConnectionAccIpv6CidrBlock struct {
	ID                     uint `gorm:"primarykey"`
	VpcPeeringConnectionID uint
	Ipv6CidrBlock          *string
}

func (VpcPeeringConnectionAccIpv6CidrBlock) TableName() string {
	return "aws_ec2_vpc_peering_connection_acc_ipv6_cidr_blocks"
}

type VpcPeeringConnectionReqCidrBlock struct {
	ID                     uint `gorm:"primarykey"`
	VpcPeeringConnectionID uint
	CidrBlock              *string
}

func (VpcPeeringConnectionReqCidrBlock) TableName() string {
	return "aws_ec2_vpc_peering_connection_req_cidr_blocks"
}

type VpcPeeringConnectionReqIpv6CidrBlock struct {
	ID                     uint `gorm:"primarykey"`
	VpcPeeringConnectionID uint
	Ipv6CidrBlock          *string
}

func (VpcPeeringConnectionReqIpv6CidrBlock) TableName() string {
	return "aws_ec2_vpc_peering_connection_req_ipv6_cidr_blocks"
}

type VpcPeeringConnectionTag struct {
	ID                     uint `gorm:"primarykey"`
	VpcPeeringConnectionID uint
	Key                    *string
	Value                  *string
}

func (VpcPeeringConnectionTag) TableName() string {
	return "aws_ec2_vpc_peering_connection_tags"
}

func (c *Client) transformVpcPeeringConnectionAccepterCidrBlocks(values []*ec2.CidrBlock) []*VpcPeeringConnectionAccCidrBlock {
	var tValues []*VpcPeeringConnectionAccCidrBlock
	for _, v := range values {
		tValues = append(tValues, &VpcPeeringConnectionAccCidrBlock{
			CidrBlock: v.CidrBlock,
		})
	}
	return tValues
}

func (c *Client) transformVpcPeeringConnectionAccepterIpv6CidrBlocks(values []*ec2.Ipv6CidrBlock) []*VpcPeeringConnectionAccIpv6CidrBlock {
	var tValues []*VpcPeeringConnectionAccIpv6CidrBlock
	for _, v := range values {
		tValues = append(tValues, &VpcPeeringConnectionAccIpv6CidrBlock{
			Ipv6CidrBlock: v.Ipv6CidrBlock,
		})
	}
	return tValues
}

func (c *Client) transformVpcPeeringConnectionRequesterCidrBlocks(values []*ec2.CidrBlock) []*VpcPeeringConnectionReqCidrBlock {
	var tValues []*VpcPeeringConnectionReqCidrBlock
	for _, v := range values {
		tValues = append(tValues, &VpcPeeringConnectionReqCidrBlock{
			CidrBlock: v.CidrBlock,
		})
	}
	return tValues
}

func (c *Client) transformVpcPeeringConnectionRequesterIpv6CidrBlocks(values []*ec2.Ipv6CidrBlock) []*VpcPeeringConnectionReqIpv6CidrBlock {
	var tValues []*VpcPeeringConnectionReqIpv6CidrBlock
	for _, v := range values {
		tValues = append(tValues, &VpcPeeringConnectionReqIpv6CidrBlock{
			Ipv6CidrBlock: v.Ipv6CidrBlock,
		})
	}
	return tValues
}

func (c *Client) transformVpcPeeringConnectionTag(value *ec2.Tag) *VpcPeeringConnectionTag {
	return &VpcPeeringConnectionTag{
		Key:   value.Key,
		Value: value.Value,
	}
}

func (c *Client) transformVpcPeeringConnectionTags(values []*ec2.Tag) []*VpcPeeringConnectionTag {
	var tValues []*VpcPeeringConnectionTag
	for _, v := range values {
		tValues = append(tValues, c.transformVpcPeeringConnectionTag(v))
	}
	return tValues
}

func (c *Client) transformVpcPeeringConnection(value *ec2.VpcPeeringConnection) *VpcPeeringConnection {
	res := VpcPeeringConnection{
		Region:                 c.region,
		AccountID:              c.accountID,
		ExpirationTime:         value.ExpirationTime,
		Status:                 value.Status,
		Tags:                   c.transformVpcPeeringConnectionTags(value.Tags),
		VpcPeeringConnectionId: value.VpcPeeringConnectionId,
	}

	if value.AccepterVpcInfo != nil {
		res.AccCidrBlock = value.AccepterVpcInfo.CidrBlock
		res.AccCidrBlockSet = c.transformVpcPeeringConnectionAccepterCidrBlocks(value.AccepterVpcInfo.CidrBlockSet)
		res.AccIpv6CidrBlockSet = c.transformVpcPeeringConnectionAccepterIpv6CidrBlocks(value.AccepterVpcInfo.Ipv6CidrBlockSet)
		res.AccOwnerId = value.AccepterVpcInfo.OwnerId
		res.AccOptions = value.AccepterVpcInfo.PeeringOptions
		res.AccRegion = value.AccepterVpcInfo.Region
		res.AccVpcId = value.AccepterVpcInfo.VpcId
	}

	if value.RequesterVpcInfo != nil {
		res.ReqCidrBlock = value.RequesterVpcInfo.CidrBlock
		res.ReqCidrBlockSet = c.transformVpcPeeringConnectionRequesterCidrBlocks(value.RequesterVpcInfo.CidrBlockSet)
		res.ReqIpv6CidrBlockSet = c.transformVpcPeeringConnectionRequesterIpv6CidrBlocks(value.RequesterVpcInfo.Ipv6CidrBlockSet)
		res.ReqOwnerId = value.RequesterVpcInfo.OwnerId
		res.ReqOptions = value.RequesterVpcInfo.PeeringOptions
		res.ReqRegion = value.RequesterVpcInfo.Region
		res.ReqVpcId = value.RequesterVpcInfo.VpcId
	}

	return &res
}

func (c *Client) transformVpcPeeringConnections(values []*ec2.VpcPeeringConnection) []*VpcPeeringConnection {
	var tValues []*VpcPeeringConnection
	for _, v := range values {
		tValues = append(tValues, c.transformVpcPeeringConnection(v))
	}
	return tValues
}

func MigrateVPCPeeringConnections(db *gorm.DB) error {
	return db.AutoMigrate(
		&VpcPeeringConnection{},
		&VpcPeeringConnectionAccCidrBlock{},
		&VpcPeeringConnectionAccIpv6CidrBlock{},
		&VpcPeeringConnectionReqCidrBlock{},
		&VpcPeeringConnectionReqIpv6CidrBlock{},
		&VpcPeeringConnectionTag{},
	)
}

func (c *Client) vpcPeeringConnections(gConfig interface{}) error {
	var config ec2.DescribeVpcPeeringConnectionsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	for {
		output, err := c.svc.DescribeVpcPeeringConnections(&config)
		if err != nil {
			return err
		}
		c.db.Where("region = ?", c.region).Where("account_id = ?", c.accountID).Delete(&VpcPeeringConnection{})
		common.ChunkedCreate(c.db, c.transformVpcPeeringConnections(output.VpcPeeringConnections))
		c.log.Info("Fetched resources", zap.String("resource", "ec2.vpc_peering_connections"), zap.Int("count", len(output.VpcPeeringConnections)))
		if aws.StringValue(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
