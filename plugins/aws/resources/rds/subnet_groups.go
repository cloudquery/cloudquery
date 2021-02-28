package rds

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/mitchellh/mapstructure"
)

type DBSubnetGroup struct {
	ID                       uint `gorm:"primarykey"`
	AccountID                string
	Region                   string
	DBSubnetGroupArn         *string `neo:"unique"`
	DBSubnetGroupDescription *string
	DBSubnetGroupName        *string
	SubnetGroupStatus        *string
	Subnets                  []*DBSubnetGroupSubnet `gorm:"constraint:OnDelete:CASCADE;"`
	VpcId                    *string
}

func (DBSubnetGroup) TableName() string {
	return "aws_rds_db_subnet_groups"
}

type DBSubnetGroupSubnet struct {
	ID              uint   `gorm:"primarykey"`
	DBSubnetGroupID uint   `neo:"ignore"`
	AccountID       string `gorm:"-"`
	Region          string `gorm:"-"`

	AvailabilityZoneName *string
	Identifier           *string
	OutpostArn           *string
	Status               *string
}

func (DBSubnetGroupSubnet) TableName() string {
	return "aws_rds_db_subnet_group_subnets"
}

func (c *Client) transformDBSubnetGroupSubnets(values *[]types.Subnet) []*DBSubnetGroupSubnet {
	var tValues []*DBSubnetGroupSubnet
	for _, value := range *values {
		res := DBSubnetGroupSubnet{
			Identifier: value.SubnetIdentifier,
			Status:     value.SubnetStatus,
		}

		if value.SubnetAvailabilityZone != nil {
			res.AvailabilityZoneName = value.SubnetAvailabilityZone.Name
		}

		if value.SubnetOutpost != nil {
			res.OutpostArn = value.SubnetOutpost.Arn
		}
		tValues = append(tValues, &res)
	}
	return tValues
}

func (c *Client) transformDBSubnetGroups(values *[]types.DBSubnetGroup) []*DBSubnetGroup {
	var tValues []*DBSubnetGroup
	for _, value := range *values {
		tValues = append(tValues, &DBSubnetGroup{
			Region:                   c.region,
			AccountID:                c.accountID,
			DBSubnetGroupArn:         value.DBSubnetGroupArn,
			DBSubnetGroupDescription: value.DBSubnetGroupDescription,
			DBSubnetGroupName:        value.DBSubnetGroupName,
			SubnetGroupStatus:        value.SubnetGroupStatus,
			Subnets:                  c.transformDBSubnetGroupSubnets(&value.Subnets),
			VpcId:                    value.VpcId,
		})
	}
	return tValues
}

var DBSubnetGroupTables = []interface{}{
	&DBSubnetGroup{},
	&DBSubnetGroupSubnet{},
}

func (c *Client) dbSubnetGroups(gConfig interface{}) error {
	var config rds.DescribeDBSubnetGroupsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	ctx := context.Background()
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(DBSubnetGroupTables...)

	for {
		output, err := c.svc.DescribeDBSubnetGroups(ctx, &config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformDBSubnetGroups(&output.DBSubnetGroups))
		c.log.Info("Fetched resources", "resource", "rds.subnet_groups", "count", len(output.DBSubnetGroups))
		if aws.ToString(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
