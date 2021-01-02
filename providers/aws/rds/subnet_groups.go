package rds

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
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

func (c *Client) transformDBSubnetGroupSubnet(value *rds.Subnet) *DBSubnetGroupSubnet {
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

	return &res
}

func (c *Client) transformDBSubnetGroupSubnets(values []*rds.Subnet) []*DBSubnetGroupSubnet {
	var tValues []*DBSubnetGroupSubnet
	for _, v := range values {
		tValues = append(tValues, c.transformDBSubnetGroupSubnet(v))
	}
	return tValues
}

func (c *Client) transformDBSubnetGroup(value *rds.DBSubnetGroup) *DBSubnetGroup {
	return &DBSubnetGroup{
		Region:                   c.region,
		AccountID:                c.accountID,
		DBSubnetGroupArn:         value.DBSubnetGroupArn,
		DBSubnetGroupDescription: value.DBSubnetGroupDescription,
		DBSubnetGroupName:        value.DBSubnetGroupName,
		SubnetGroupStatus:        value.SubnetGroupStatus,
		Subnets:                  c.transformDBSubnetGroupSubnets(value.Subnets),
		VpcId:                    value.VpcId,
	}
}

func (c *Client) transformDBSubnetGroups(values []*rds.DBSubnetGroup) []*DBSubnetGroup {
	var tValues []*DBSubnetGroup
	for _, v := range values {
		tValues = append(tValues, c.transformDBSubnetGroup(v))
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
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(DBSubnetGroupTables...)

	for {
		output, err := c.svc.DescribeDBSubnetGroups(&config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformDBSubnetGroups(output.DBSubnetGroups))
		c.log.Info("Fetched resources", zap.String("resource", "rds.subnet_groups"), zap.Int("count", len(output.DBSubnetGroups)))
		if aws.StringValue(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
