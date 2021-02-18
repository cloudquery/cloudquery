package ec2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

type NetworkAcl struct {
	ID           uint                     `gorm:"primarykey"`
	AccountID    string                   `neo:"unique"`
	Region       string                   `neo:"unique"`
	Associations []*NetworkAclAssociation `gorm:"constraint:OnDelete:CASCADE;"`
	Entries      []*NetworkAclEntry       `gorm:"constraint:OnDelete:CASCADE;"`
	IsDefault    *bool
	NetworkAclId *string `neo:"unique"`
	OwnerId      *string
	Tags         []*NetworkAclTag `gorm:"constraint:OnDelete:CASCADE;"`
	VpcId        *string
}

func (NetworkAcl) TableName() string {
	return "aws_ec2_network_acls"
}

type NetworkAclAssociation struct {
	ID           uint `gorm:"primarykey"`
	NetworkAclID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	NetworkAclAssociationId *string
	NetworkAclId            *string
	SubnetId                *string
}

func (NetworkAclAssociation) TableName() string {
	return "aws_ec2_network_acl_associations"
}

type NetworkAclEntry struct {
	ID           uint `gorm:"primarykey"`
	NetworkAclID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	CidrBlock *string
	Egress    *bool

	IcmpTypeCode *int32
	IcmpTypeType *int32

	Ipv6CidrBlock *string

	PortRangeFrom *int32
	PortRangeTo   *int32

	Protocol   *string
	RuleAction *string
	RuleNumber *int32
}

func (NetworkAclEntry) TableName() string {
	return "aws_ec2_network_acl_entries"
}

type NetworkAclTag struct {
	ID           uint `gorm:"primarykey"`
	NetworkAclID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Key   *string
	Value *string
}

func (NetworkAclTag) TableName() string {
	return "aws_ec2_network_acl_tags"
}

func (c *Client) transformNetworkAclAssociations(values *[]types.NetworkAclAssociation) []*NetworkAclAssociation {
	var tValues []*NetworkAclAssociation
	for _, value := range *values {
		tValues = append(tValues, &NetworkAclAssociation{
			AccountID:               c.accountID,
			Region:                  c.region,
			NetworkAclAssociationId: value.NetworkAclAssociationId,
			NetworkAclId:            value.NetworkAclId,
			SubnetId:                value.SubnetId,
		})
	}
	return tValues
}

func (c *Client) transformNetworkAclEntrys(values *[]types.NetworkAclEntry) []*NetworkAclEntry {
	var tValues []*NetworkAclEntry
	for _, value := range *values {
		res := NetworkAclEntry{
			AccountID: c.accountID,
			Region:    c.region,

			CidrBlock: value.CidrBlock,
			Egress:    &value.Egress,

			Ipv6CidrBlock: value.Ipv6CidrBlock,

			Protocol:   value.Protocol,
			RuleAction: aws.String(string(value.RuleAction)),
			RuleNumber: &value.RuleNumber,
		}

		if value.IcmpTypeCode != nil {
			res.IcmpTypeCode = &value.IcmpTypeCode.Code
			res.IcmpTypeType = &value.IcmpTypeCode.Type
		}

		if value.PortRange != nil {
			res.PortRangeFrom = &value.PortRange.From
			res.PortRangeTo = &value.PortRange.To
		}
		tValues = append(tValues, &res)
	}
	return tValues
}

func (c *Client) transformNetworkAclTags(values *[]types.Tag) []*NetworkAclTag {
	var tValues []*NetworkAclTag
	for _, value := range *values {
		tValues = append(tValues, &NetworkAclTag{
			AccountID: c.accountID,
			Region:    c.region,
			Key:       value.Key,
			Value:     value.Value,
		})
	}
	return tValues
}

func (c *Client) transformNetworkAcls(values *[]types.NetworkAcl) []*NetworkAcl {
	var tValues []*NetworkAcl
	for _, value := range *values {
		tValues = append(tValues, &NetworkAcl{
			Region:       c.region,
			AccountID:    c.accountID,
			Associations: c.transformNetworkAclAssociations(&value.Associations),
			Entries:      c.transformNetworkAclEntrys(&value.Entries),
			IsDefault:    &value.IsDefault,
			NetworkAclId: value.NetworkAclId,
			OwnerId:      value.OwnerId,
			Tags:         c.transformNetworkAclTags(&value.Tags),
			VpcId:        value.VpcId,
		})
	}
	return tValues
}

var NetworkAclTables = []interface{}{
	&NetworkAcl{},
	&NetworkAclAssociation{},
	&NetworkAclEntry{},
	&NetworkAclTag{},
}

func (c *Client) networkAcls(gConfig interface{}) error {
	ctx := context.Background()
	var config ec2.DescribeNetworkAclsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(NetworkAclTables...)
	for {
		output, err := c.svc.DescribeNetworkAcls(ctx, &config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformNetworkAcls(&output.NetworkAcls))
		c.log.Info("Fetched resources", zap.String("resource", "ec2.network_acls"), zap.Int("count", len(output.NetworkAcls)))
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
