package rds

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DBSubnetGroup struct {
	ID                       uint `gorm:"primarykey"`
	AccountID                string
	Region                   string
	DBSubnetGroupArn         *string
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
	ID                     uint `gorm:"primarykey"`
	DBSubnetGroupID        uint
	SubnetAvailabilityZone *rds.AvailabilityZone `gorm:"embedded;embeddedPrefix:availability_zone_"`
	SubnetIdentifier       *string
	SubnetOutpost          *rds.Outpost `gorm:"embedded;embeddedPrefix:outpost_"`
	SubnetStatus           *string
}

func (DBSubnetGroupSubnet) TableName() string {
	return "aws_rds_db_subnet_group_subnets"
}

func (c *Client) transformDBSubnetGroupSubnet(value *rds.Subnet) *DBSubnetGroupSubnet {
	return &DBSubnetGroupSubnet{
		SubnetAvailabilityZone: value.SubnetAvailabilityZone,
		SubnetIdentifier:       value.SubnetIdentifier,
		SubnetOutpost:          value.SubnetOutpost,
		SubnetStatus:           value.SubnetStatus,
	}
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

func MigrateDBSubnetGroups(db *gorm.DB) error {
	return db.AutoMigrate(
		&DBSubnetGroup{},
		&DBSubnetGroupSubnet{},
	)
}

func (c *Client) dbSubnetGroups(gConfig interface{}) error {
	var config rds.DescribeDBSubnetGroupsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	for {
		output, err := c.svc.DescribeDBSubnetGroups(&config)
		if err != nil {
			return err
		}
		c.db.Where("region = ?", c.region).Where("account_id = ?", c.accountID).Delete(&DBSubnetGroup{})
		common.ChunkedCreate(c.db, c.transformDBSubnetGroups(output.DBSubnetGroups))
		c.log.Info("Fetched resources", zap.String("resource", "rds.subnet_groups"), zap.Int("count", len(output.DBSubnetGroups)))
		if aws.StringValue(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
