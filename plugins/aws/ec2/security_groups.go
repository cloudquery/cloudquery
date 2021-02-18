package ec2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

type SecurityGroup struct {
	ID            uint   `gorm:"primarykey"`
	AccountID     string `neo:"unique"`
	Region        string `neo:"unique"`
	Description   *string
	GroupId       *string `neo:"unique"`
	GroupName     *string
	IpPermissions []*SecurityGroupIpPermission `gorm:"constraint:OnDelete:CASCADE;"`
	OwnerId       *string
	Tags          []*SecurityGroupTag `gorm:"constraint:OnDelete:CASCADE;"`
	VpcId         *string
}

func (SecurityGroup) TableName() string {
	return "aws_ec2_security_groups"
}

type SecurityGroupIpPermission struct {
	ID               uint   `gorm:"primarykey"`
	SecurityGroupID  uint   `neo:"ignore"`
	AccountID        string `gorm:"-"`
	Region           string `gorm:"-"`
	FromPort         *int32
	IpProtocol       *string
	IpRanges         []*SecurityGroupIpRange      `gorm:"constraint:OnDelete:CASCADE;"`
	Ipv6Ranges       []*SecurityGroupIpv6Range    `gorm:"constraint:OnDelete:CASCADE;"`
	PrefixListIds    []*SecurityGroupPrefixListId `gorm:"constraint:OnDelete:CASCADE;"`
	ToPort           *int32
	UserIdGroupPairs []*SecurityGroupUserIdGroupPair `gorm:"constraint:OnDelete:CASCADE;"`
}

func (SecurityGroupIpPermission) TableName() string {
	return "aws_ec2_security_group_ip_permissions"
}

type SecurityGroupIpRange struct {
	ID                          uint   `gorm:"primarykey"`
	SecurityGroupIpPermissionID uint   `neo:"ignore"`
	AccountID                   string `gorm:"-"`
	Region                      string `gorm:"-"`
	CidrIp                      *string
	Description                 *string
}

func (SecurityGroupIpRange) TableName() string {
	return "aws_ec2_security_group_ip_ranges"
}

type SecurityGroupIpv6Range struct {
	ID                          uint   `gorm:"primarykey"`
	SecurityGroupIpPermissionID uint   `neo:"ignore"`
	AccountID                   string `gorm:"-"`
	Region                      string `gorm:"-"`
	CidrIpv6                    *string
	Description                 *string
}

func (SecurityGroupIpv6Range) TableName() string {
	return "aws_ec2_security_group_ipv6_ranges"
}

type SecurityGroupPrefixListId struct {
	ID                          uint   `gorm:"primarykey"`
	SecurityGroupIpPermissionID uint   `neo:"ignore"`
	AccountID                   string `gorm:"-"`
	Region                      string `gorm:"-"`
	Description                 *string
	PrefixListId                *string
}

func (SecurityGroupPrefixListId) TableName() string {
	return "aws_ec2_security_group_prefix_list_ids"
}

type SecurityGroupUserIdGroupPair struct {
	ID                          uint   `gorm:"primarykey"`
	SecurityGroupIpPermissionID uint   `neo:"ignore"`
	AccountID                   string `gorm:"-"`
	Region                      string `gorm:"-"`
	Description                 *string
	GroupId                     *string
	GroupName                   *string
	PeeringStatus               *string
	UserId                      *string
	VpcId                       *string
	VpcPeeringConnectionId      *string
}

func (SecurityGroupUserIdGroupPair) TableName() string {
	return "aws_ec2_security_group_user_id_group_paris"
}

type SecurityGroupTag struct {
	ID              uint   `gorm:"primarykey"`
	SecurityGroupID uint   `neo:"ignore"`
	AccountID       string `gorm:"-"`
	Region          string `gorm:"-"`
	Key             *string
	Value           *string
}

func (SecurityGroupTag) TableName() string {
	return "aws_ec2_security_group_tags"
}

func (c *Client) transformSecurityGroupIpRanges(values *[]types.IpRange) []*SecurityGroupIpRange {
	var tValues []*SecurityGroupIpRange
	for _, value := range *values {
		tValues = append(tValues, &SecurityGroupIpRange{
			AccountID:   c.accountID,
			Region:      c.region,
			CidrIp:      value.CidrIp,
			Description: value.Description,
		})
	}
	return tValues
}

func (c *Client) transformSecurityGroupIpv6Ranges(values *[]types.Ipv6Range) []*SecurityGroupIpv6Range {
	var tValues []*SecurityGroupIpv6Range
	for _, value := range *values {
		tValues = append(tValues, &SecurityGroupIpv6Range{
			AccountID:   c.accountID,
			Region:      c.region,
			CidrIpv6:    value.CidrIpv6,
			Description: value.Description,
		})
	}
	return tValues
}


func (c *Client) transformSecurityGroupPrefixListIds(values *[]types.PrefixListId) []*SecurityGroupPrefixListId {
	var tValues []*SecurityGroupPrefixListId
	for _, value := range *values {
		tValues = append(tValues, &SecurityGroupPrefixListId{
			AccountID:    c.accountID,
			Region:       c.region,
			Description:  value.Description,
			PrefixListId: value.PrefixListId,
		})
	}
	return tValues
}

func (c *Client) transformSecurityGroupUserIdGroupPairs(values *[]types.UserIdGroupPair) []*SecurityGroupUserIdGroupPair {
	var tValues []*SecurityGroupUserIdGroupPair
	for _, value := range *values {
		tValues = append(tValues,  &SecurityGroupUserIdGroupPair{
			AccountID:              c.accountID,
			Region:                 c.region,
			Description:            value.Description,
			GroupId:                value.GroupId,
			GroupName:              value.GroupName,
			PeeringStatus:          value.PeeringStatus,
			UserId:                 value.UserId,
			VpcId:                  value.VpcId,
			VpcPeeringConnectionId: value.VpcPeeringConnectionId,
		})
	}
	return tValues
}

func (c *Client) transformSecurityGroupIpPermissions(values *[]types.IpPermission) []*SecurityGroupIpPermission {
	var tValues []*SecurityGroupIpPermission
	for _, value := range *values {
		tValues = append(tValues, &SecurityGroupIpPermission{
			AccountID:        c.accountID,
			Region:           c.region,
			FromPort:         &value.FromPort,
			IpProtocol:       value.IpProtocol,
			IpRanges:         c.transformSecurityGroupIpRanges(&value.IpRanges),
			Ipv6Ranges:       c.transformSecurityGroupIpv6Ranges(&value.Ipv6Ranges),
			PrefixListIds:    c.transformSecurityGroupPrefixListIds(&value.PrefixListIds),
			ToPort:           &value.ToPort,
			UserIdGroupPairs: c.transformSecurityGroupUserIdGroupPairs(&value.UserIdGroupPairs),
		})
	}
	return tValues
}

func (c *Client) transformSecurityGroupTags(values *[]types.Tag) []*SecurityGroupTag {
	var tValues []*SecurityGroupTag
	for _, value := range *values {
		tValues = append(tValues, &SecurityGroupTag{
			AccountID: c.accountID,
			Region:    c.region,
			Key:       value.Key,
			Value:     value.Value,
		})
	}
	return tValues
}

func (c *Client) transformSecurityGroups(values *[]types.SecurityGroup) []*SecurityGroup {
	var tValues []*SecurityGroup
	for _, value := range *values {
		tValues = append(tValues, &SecurityGroup{
			Region:        c.region,
			AccountID:     c.accountID,
			Description:   value.Description,
			GroupId:       value.GroupId,
			GroupName:     value.GroupName,
			IpPermissions: c.transformSecurityGroupIpPermissions(&value.IpPermissions),
			OwnerId:       value.OwnerId,
			Tags:          c.transformSecurityGroupTags(&value.Tags),
			VpcId:         value.VpcId,
		})
	}
	return tValues
}

var SecurityGroupTables = []interface{}{
	&SecurityGroup{},
	&SecurityGroupIpPermission{},
	&SecurityGroupIpRange{},
	&SecurityGroupIpv6Range{},
	&SecurityGroupPrefixListId{},
	&SecurityGroupUserIdGroupPair{},
	&SecurityGroupTag{},
}

func (c *Client) securityGroups(gConfig interface{}) error {
	ctx := context.Background()
	var config ec2.DescribeSecurityGroupsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(SecurityGroupTables...)
	for {
		output, err := c.svc.DescribeSecurityGroups(ctx, &config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformSecurityGroups(&output.SecurityGroups))
		c.log.Info("Fetched resources", zap.String("resource", "ec2.security_groups"), zap.Int("count", len(output.SecurityGroups)))
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
