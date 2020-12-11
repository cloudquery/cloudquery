package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type VpcPeeringConnection struct {
	ID        uint `gorm:"primarykey"`
	AccountID string
	Region    string

	AccepterCidrBlock        *string
	AccepterCidrBlockSet     []*VpcPeeringConnectionAccepterCidrBlock     `gorm:"constraint:OnDelete:CASCADE;"`
	AccepterIpv6CidrBlockSet []*VpcPeeringConnectionAccepterIpv6CidrBlock `gorm:"constraint:OnDelete:CASCADE;"`
	AccepterOwnerId          *string
	AccepterOptions          *ec2.VpcPeeringConnectionOptionsDescription `gorm:"embedded;embeddedPrefix:accepter_option_"`
	AccepterRegion           *string
	AccepterVpcId            *string

	ExpirationTime *time.Time

	RequesterCidrBlock        *string
	RequesterCidrBlockSet     []*VpcPeeringConnectionRequesterCidrBlock     `gorm:"constraint:OnDelete:CASCADE;"`
	RequesterIpv6CidrBlockSet []*VpcPeeringConnectionRequesterIpv6CidrBlock `gorm:"constraint:OnDelete:CASCADE;"`
	RequesterOwnerId          *string
	RequesterOptions          *ec2.VpcPeeringConnectionOptionsDescription `gorm:"embedded;embeddedPrefix:requester_option_"`
	RequesterRegion           *string
	RequesterVpcId            *string

	Status                 *ec2.VpcPeeringConnectionStateReason `gorm:"embedded;embeddedPrefix:status_"`
	Tags                   []*VpcPeeringConnectionTag           `gorm:"constraint:OnDelete:CASCADE;"`
	VpcPeeringConnectionId *string
}

type VpcPeeringConnectionAccepterCidrBlock struct {
	ID                     uint `gorm:"primarykey"`
	VpcPeeringConnectionID uint
	CidrBlock              *string
}

type VpcPeeringConnectionAccepterIpv6CidrBlock struct {
	ID                     uint `gorm:"primarykey"`
	VpcPeeringConnectionID uint
	Ipv6CidrBlock          *string
}

type VpcPeeringConnectionRequesterCidrBlock struct {
	ID                     uint `gorm:"primarykey"`
	VpcPeeringConnectionID uint
	CidrBlock              *string
}

type VpcPeeringConnectionRequesterIpv6CidrBlock struct {
	ID                     uint `gorm:"primarykey"`
	VpcPeeringConnectionID uint
	Ipv6CidrBlock          *string
}

type VpcPeeringConnectionTag struct {
	ID                     uint `gorm:"primarykey"`
	VpcPeeringConnectionID uint
	Key                    *string
	Value                  *string
}

func (c *Client) transformVpcPeeringConnectionAccepterCidrBlocks(values []*ec2.CidrBlock) []*VpcPeeringConnectionAccepterCidrBlock {
	var tValues []*VpcPeeringConnectionAccepterCidrBlock
	for _, v := range values {
		tValues = append(tValues, &VpcPeeringConnectionAccepterCidrBlock{
			CidrBlock: v.CidrBlock,
		})
	}
	return tValues
}

func (c *Client) transformVpcPeeringConnectionAccepterIpv6CidrBlocks(values []*ec2.Ipv6CidrBlock) []*VpcPeeringConnectionAccepterIpv6CidrBlock {
	var tValues []*VpcPeeringConnectionAccepterIpv6CidrBlock
	for _, v := range values {
		tValues = append(tValues, &VpcPeeringConnectionAccepterIpv6CidrBlock{
			Ipv6CidrBlock: v.Ipv6CidrBlock,
		})
	}
	return tValues
}

func (c *Client) transformVpcPeeringConnectionRequesterCidrBlocks(values []*ec2.CidrBlock) []*VpcPeeringConnectionRequesterCidrBlock {
	var tValues []*VpcPeeringConnectionRequesterCidrBlock
	for _, v := range values {
		tValues = append(tValues, &VpcPeeringConnectionRequesterCidrBlock{
			CidrBlock: v.CidrBlock,
		})
	}
	return tValues
}

func (c *Client) transformVpcPeeringConnectionRequesterIpv6CidrBlocks(values []*ec2.Ipv6CidrBlock) []*VpcPeeringConnectionRequesterIpv6CidrBlock {
	var tValues []*VpcPeeringConnectionRequesterIpv6CidrBlock
	for _, v := range values {
		tValues = append(tValues, &VpcPeeringConnectionRequesterIpv6CidrBlock{
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
		res.AccepterCidrBlock = value.AccepterVpcInfo.CidrBlock
		res.AccepterCidrBlockSet = c.transformVpcPeeringConnectionAccepterCidrBlocks(value.AccepterVpcInfo.CidrBlockSet)
		res.AccepterIpv6CidrBlockSet = c.transformVpcPeeringConnectionAccepterIpv6CidrBlocks(value.AccepterVpcInfo.Ipv6CidrBlockSet)
		res.AccepterOwnerId = value.AccepterVpcInfo.OwnerId
		res.AccepterOptions = value.AccepterVpcInfo.PeeringOptions
		res.AccepterRegion = value.AccepterVpcInfo.Region
		res.AccepterVpcId = value.AccepterVpcInfo.VpcId
	}

	if value.RequesterVpcInfo != nil {
		res.RequesterCidrBlock = value.RequesterVpcInfo.CidrBlock
		res.RequesterCidrBlockSet = c.transformVpcPeeringConnectionRequesterCidrBlocks(value.RequesterVpcInfo.CidrBlockSet)
		res.RequesterIpv6CidrBlockSet = c.transformVpcPeeringConnectionRequesterIpv6CidrBlocks(value.RequesterVpcInfo.Ipv6CidrBlockSet)
		res.RequesterOwnerId = value.RequesterVpcInfo.OwnerId
		res.RequesterOptions = value.RequesterVpcInfo.PeeringOptions
		res.RequesterRegion = value.RequesterVpcInfo.Region
		res.RequesterVpcId = value.RequesterVpcInfo.VpcId
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

func (c *Client) vpcPeeringConnections(gConfig interface{}) error {
	var config ec2.DescribeVpcPeeringConnectionsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["ec2VpcPeeringConnection"] {
		err := c.db.AutoMigrate(
			&VpcPeeringConnection{},
			&VpcPeeringConnectionAccepterCidrBlock{},
			&VpcPeeringConnectionAccepterIpv6CidrBlock{},
			&VpcPeeringConnectionRequesterCidrBlock{},
			&VpcPeeringConnectionRequesterIpv6CidrBlock{},
			&VpcPeeringConnectionTag{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["ec2VpcPeeringConnection"] = true
	}
	for {
		output, err := c.svc.DescribeVpcPeeringConnections(&config)
		if err != nil {
			return err
		}
		c.db.Where("region = ?", c.region).Where("account_id = ?", c.accountID).Delete(&VpcPeeringConnection{})
		common.ChunkedCreate(c.db, c.transformVpcPeeringConnections(output.VpcPeeringConnections))
		c.log.Info("Fetched resources", zap.Int("count", len(output.VpcPeeringConnections)))
		if aws.StringValue(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
