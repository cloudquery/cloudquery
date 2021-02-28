package redshift

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/mitchellh/mapstructure"
)

type ClusterSubnetGroup struct {
	ID        uint   `gorm:"primarykey"`
	AccountID string `neo:"unique"`
	Region    string `neo:"unique"`

	Name        *string `neo:"unique"`
	Description *string
	Status      *string
	Subnets     []*ClusterSubnetGroupSubnet `gorm:"constraint:OnDelete:CASCADE;"`
	Tags        []*ClusterSubnetGroupTag    `gorm:"constraint:OnDelete:CASCADE;"`
	VpcId       *string
}

func (ClusterSubnetGroup) TableName() string {
	return "aws_redshift_cluster_subnet_groups"
}

type ClusterSubnetGroupTag struct {
	ID                   uint   `gorm:"primarykey"`
	ClusterSubnetGroupID uint   `neo:"ignore"`
	AccountID            string `gorm:"-"`
	Region               string `gorm:"-"`

	Key   *string
	Value *string
}

func (ClusterSubnetGroupTag) TableName() string {
	return "aws_redshift_cluster_subnet_group_tags"
}

type ClusterSubnetGroupSubnet struct {
	ID                   uint   `gorm:"primarykey"`
	ClusterSubnetGroupID uint   `neo:"ignore"`
	AccountID            string `gorm:"-"`
	Region               string `gorm:"-"`

	AZName      *string
	AZPlatforms []*ClusterSubnetGroupSupportedPlatform `gorm:"constraint:OnDelete:CASCADE;"`
	Identifier  *string
	Status      *string
}

func (ClusterSubnetGroupSubnet) TableName() string {
	return "aws_redshift_cluster_subnet_group_subnets"
}

type ClusterSubnetGroupSupportedPlatform struct {
	ID                         uint   `gorm:"primarykey"`
	ClusterSubnetGroupSubnetID uint   `neo:"ignore"`
	AccountID                  string `gorm:"-"`
	Region                     string `gorm:"-"`

	Name *string
}

func (ClusterSubnetGroupSupportedPlatform) TableName() string {
	return "aws_redshift_cluster_subnet_group_supported_platforms"
}

func (c *Client) transformClusterSubnetGroups(values *[]types.ClusterSubnetGroup) []*ClusterSubnetGroup {
	var tValues []*ClusterSubnetGroup
	for _, value := range *values {
		tValue := ClusterSubnetGroup{
			AccountID:   c.accountID,
			Region:      c.region,
			Name:        value.ClusterSubnetGroupName,
			Description: value.Description,
			Status:      value.SubnetGroupStatus,
			Subnets:     c.transformClusterSubnetGroupSubnets(&value.Subnets),
			Tags:        c.transformClusterSubnetGroupTags(&value.Tags),
			VpcId:       value.VpcId,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformClusterSubnetGroupSupportedPlatforms(values *[]types.SupportedPlatform) []*ClusterSubnetGroupSupportedPlatform {
	var tValues []*ClusterSubnetGroupSupportedPlatform
	for _, value := range *values {
		tValue := ClusterSubnetGroupSupportedPlatform{
			AccountID: c.accountID,
			Region:    c.region,
			Name:      value.Name,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformClusterSubnetGroupSubnets(values *[]types.Subnet) []*ClusterSubnetGroupSubnet {
	var tValues []*ClusterSubnetGroupSubnet
	for _, value := range *values {
		tValue := ClusterSubnetGroupSubnet{
			AccountID:  c.accountID,
			Region:     c.region,
			Identifier: value.SubnetIdentifier,
			Status:     value.SubnetStatus,
		}
		if value.SubnetAvailabilityZone != nil {
			tValue.AZName = value.SubnetAvailabilityZone.Name
			tValue.AZPlatforms = c.transformClusterSubnetGroupSupportedPlatforms(&value.SubnetAvailabilityZone.SupportedPlatforms)
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformClusterSubnetGroupTags(values *[]types.Tag) []*ClusterSubnetGroupTag {
	var tValues []*ClusterSubnetGroupTag
	for _, value := range *values {
		tValue := ClusterSubnetGroupTag{
			AccountID: c.accountID,
			Region:    c.region,
			Key:       value.Key,
			Value:     value.Value,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

var ClusterSubnetGroupTables = []interface{}{
	&ClusterSubnetGroup{},
	&ClusterSubnetGroupSubnet{},
	&ClusterSubnetGroupSupportedPlatform{},
	&ClusterSubnetGroupTag{},
}

func (c *Client) clusterSubnetGroups(gConfig interface{}) error {
	var config redshift.DescribeClusterSubnetGroupsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	ctx := context.Background()
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(ClusterSubnetGroupTables...)

	for {
		output, err := c.svc.DescribeClusterSubnetGroups(ctx, &config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformClusterSubnetGroups(&output.ClusterSubnetGroups))
		c.log.Info("Fetched resources", "resource", "redshift.cluster_subnet_groups", "count", len(output.ClusterSubnetGroups))
		if aws.ToString(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}

	return nil
}
