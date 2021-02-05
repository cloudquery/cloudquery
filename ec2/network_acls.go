package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
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

	IcmpTypeCode *int64
	IcmpTypeType *int64

	Ipv6CidrBlock *string

	PortRangeFrom *int64
	PortRangeTo   *int64

	Protocol   *string
	RuleAction *string
	RuleNumber *int64
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

func (c *Client) transformNetworkAclAssociation(value *ec2.NetworkAclAssociation) *NetworkAclAssociation {
	return &NetworkAclAssociation{
		AccountID:               c.accountID,
		Region:                  c.region,
		NetworkAclAssociationId: value.NetworkAclAssociationId,
		NetworkAclId:            value.NetworkAclId,
		SubnetId:                value.SubnetId,
	}
}

func (c *Client) transformNetworkAclAssociations(values []*ec2.NetworkAclAssociation) []*NetworkAclAssociation {
	var tValues []*NetworkAclAssociation
	for _, v := range values {
		tValues = append(tValues, c.transformNetworkAclAssociation(v))
	}
	return tValues
}

func (c *Client) transformNetworkAclEntry(value *ec2.NetworkAclEntry) *NetworkAclEntry {
	res := NetworkAclEntry{
		AccountID: c.accountID,
		Region:    c.region,

		CidrBlock: value.CidrBlock,
		Egress:    value.Egress,

		Ipv6CidrBlock: value.Ipv6CidrBlock,

		Protocol:   value.Protocol,
		RuleAction: value.RuleAction,
		RuleNumber: value.RuleNumber,
	}

	if value.IcmpTypeCode != nil {
		res.IcmpTypeCode = value.IcmpTypeCode.Code
		res.IcmpTypeType = value.IcmpTypeCode.Type
	}

	if value.PortRange != nil {
		res.PortRangeFrom = value.PortRange.From
		res.PortRangeTo = value.PortRange.To
	}

	return &res
}

func (c *Client) transformNetworkAclEntrys(values []*ec2.NetworkAclEntry) []*NetworkAclEntry {
	var tValues []*NetworkAclEntry
	for _, v := range values {
		tValues = append(tValues, c.transformNetworkAclEntry(v))
	}
	return tValues
}

func (c *Client) transformNetworkAclTag(value *ec2.Tag) *NetworkAclTag {
	return &NetworkAclTag{
		AccountID: c.accountID,
		Region:    c.region,
		Key:       value.Key,
		Value:     value.Value,
	}
}

func (c *Client) transformNetworkAclTags(values []*ec2.Tag) []*NetworkAclTag {
	var tValues []*NetworkAclTag
	for _, v := range values {
		tValues = append(tValues, c.transformNetworkAclTag(v))
	}
	return tValues
}

func (c *Client) transformNetworkAcl(value *ec2.NetworkAcl) *NetworkAcl {
	return &NetworkAcl{
		Region:       c.region,
		AccountID:    c.accountID,
		Associations: c.transformNetworkAclAssociations(value.Associations),
		Entries:      c.transformNetworkAclEntrys(value.Entries),
		IsDefault:    value.IsDefault,
		NetworkAclId: value.NetworkAclId,
		OwnerId:      value.OwnerId,
		Tags:         c.transformNetworkAclTags(value.Tags),
		VpcId:        value.VpcId,
	}
}

func (c *Client) transformNetworkAcls(values []*ec2.NetworkAcl) []*NetworkAcl {
	var tValues []*NetworkAcl
	for _, v := range values {
		tValues = append(tValues, c.transformNetworkAcl(v))
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
	var config ec2.DescribeNetworkAclsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(NetworkAclTables...)
	for {
		output, err := c.svc.DescribeNetworkAcls(&config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformNetworkAcls(output.NetworkAcls))
		c.log.Info("Fetched resources", zap.String("resource", "ec2.network_acls"), zap.Int("count", len(output.NetworkAcls)))
		if aws.StringValue(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
