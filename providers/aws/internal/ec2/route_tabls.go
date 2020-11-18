package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/mitchellh/mapstructure"
	"github.com/cloudquery/cloudquery/providers/common"
	"go.uber.org/zap"
)

type RouteTable struct {
	ID              uint `gorm:"primarykey"`
	AccountID       string
	Region          string
	Associations    []*RouteTableAssociation `gorm:"constraint:OnDelete:CASCADE;"`
	OwnerId         *string
	PropagatingVgws []*RouteTablePropagatingVgw `gorm:"constraint:OnDelete:CASCADE;"`
	RouteTableId    *string
	Routes          []*RouteTableRoute `gorm:"constraint:OnDelete:CASCADE;"`
	Tags            []*RouteTableTag   `gorm:"constraint:OnDelete:CASCADE;"`
	VpcId           *string
}

type RouteTableAssociation struct {
	ID                      uint `gorm:"primarykey"`
	RouteTableID            uint
	AssociationState        *ec2.RouteTableAssociationState `gorm:"embedded;embeddedPrefix:association_state_"`
	GatewayId               *string
	Main                    *bool
	RouteTableAssociationId *string
	RouteTableId            *string
	SubnetId                *string
}

type RouteTablePropagatingVgw struct {
	ID           uint `gorm:"primarykey"`
	RouteTableID uint
	GatewayId    *string
}

type RouteTableRoute struct {
	ID                          uint `gorm:"primarykey"`
	RouteTableID                uint
	CarrierGatewayId            *string
	DestinationCidrBlock        *string
	DestinationIpv6CidrBlock    *string
	DestinationPrefixListId     *string
	EgressOnlyInternetGatewayId *string
	GatewayId                   *string
	InstanceId                  *string
	InstanceOwnerId             *string
	LocalGatewayId              *string
	NatGatewayId                *string
	NetworkInterfaceId          *string
	Origin                      *string
	State                       *string
	TransitGatewayId            *string
	VpcPeeringConnectionId      *string
}

type RouteTableTag struct {
	ID           uint `gorm:"primarykey"`
	RouteTableID uint
	Key          *string
	Value        *string
}

func (c *Client) transformRouteTableAssociation(value *ec2.RouteTableAssociation) *RouteTableAssociation {
	return &RouteTableAssociation{
		AssociationState:        value.AssociationState,
		GatewayId:               value.GatewayId,
		Main:                    value.Main,
		RouteTableAssociationId: value.RouteTableAssociationId,
		RouteTableId:            value.RouteTableId,
		SubnetId:                value.SubnetId,
	}
}

func (c *Client) transformRouteTableAssociations(values []*ec2.RouteTableAssociation) []*RouteTableAssociation {
	var tValues []*RouteTableAssociation
	for _, v := range values {
		tValues = append(tValues, c.transformRouteTableAssociation(v))
	}
	return tValues
}

func (c *Client) transformRouteTablePropagatingVgw(value *ec2.PropagatingVgw) *RouteTablePropagatingVgw {
	return &RouteTablePropagatingVgw{
		GatewayId: value.GatewayId,
	}
}

func (c *Client) transformRouteTablePropagatingVgws(values []*ec2.PropagatingVgw) []*RouteTablePropagatingVgw {
	var tValues []*RouteTablePropagatingVgw
	for _, v := range values {
		tValues = append(tValues, c.transformRouteTablePropagatingVgw(v))
	}
	return tValues
}

func (c *Client) transformRouteTableRoute(value *ec2.Route) *RouteTableRoute {
	return &RouteTableRoute{
		CarrierGatewayId:            value.CarrierGatewayId,
		DestinationCidrBlock:        value.DestinationCidrBlock,
		DestinationIpv6CidrBlock:    value.DestinationIpv6CidrBlock,
		DestinationPrefixListId:     value.DestinationPrefixListId,
		EgressOnlyInternetGatewayId: value.EgressOnlyInternetGatewayId,
		GatewayId:                   value.GatewayId,
		InstanceId:                  value.InstanceId,
		InstanceOwnerId:             value.InstanceOwnerId,
		LocalGatewayId:              value.LocalGatewayId,
		NatGatewayId:                value.NatGatewayId,
		NetworkInterfaceId:          value.NetworkInterfaceId,
		Origin:                      value.Origin,
		State:                       value.State,
		TransitGatewayId:            value.TransitGatewayId,
		VpcPeeringConnectionId:      value.VpcPeeringConnectionId,
	}
}

func (c *Client) transformRouteTableRoutes(values []*ec2.Route) []*RouteTableRoute {
	var tValues []*RouteTableRoute
	for _, v := range values {
		tValues = append(tValues, c.transformRouteTableRoute(v))
	}
	return tValues
}

func (c *Client) transformRouteTableTag(value *ec2.Tag) *RouteTableTag {
	return &RouteTableTag{
		Key:   value.Key,
		Value: value.Value,
	}
}

func (c *Client) transformRouteTableTags(values []*ec2.Tag) []*RouteTableTag {
	var tValues []*RouteTableTag
	for _, v := range values {
		tValues = append(tValues, c.transformRouteTableTag(v))
	}
	return tValues
}

func (c *Client) transformRouteTable(value *ec2.RouteTable) *RouteTable {
	return &RouteTable{
		Region:          c.region,
		AccountID:       c.accountID,
		Associations:    c.transformRouteTableAssociations(value.Associations),
		OwnerId:         value.OwnerId,
		PropagatingVgws: c.transformRouteTablePropagatingVgws(value.PropagatingVgws),
		RouteTableId:    value.RouteTableId,
		Routes:          c.transformRouteTableRoutes(value.Routes),
		Tags:            c.transformRouteTableTags(value.Tags),
		VpcId:           value.VpcId,
	}
}

func (c *Client) transformRouteTables(values []*ec2.RouteTable) []*RouteTable {
	var tValues []*RouteTable
	for _, v := range values {
		tValues = append(tValues, c.transformRouteTable(v))
	}
	return tValues
}

func (c *Client) RouteTables(gConfig interface{}) error {
	var config ec2.DescribeRouteTablesInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["ec2RouteTable"] {
		err := c.db.AutoMigrate(
			&RouteTable{},
			&RouteTableAssociation{},
			&RouteTablePropagatingVgw{},
			&RouteTableRoute{},
			&RouteTableTag{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["ec2RouteTable"] = true
	}
	for {
		output, err := c.svc.DescribeRouteTables(&config)
		if err != nil {
			return err
		}
		c.log.Debug("deleting previous RouteTables", zap.String("region", c.region), zap.String("account_id", c.accountID))
		c.db.Where("region = ?", c.region).Where("account_id = ?", c.accountID).Delete(&RouteTable{})
		common.ChunkedCreate(c.db, c.transformRouteTables(output.RouteTables))
		c.log.Info("populating RouteTables", zap.Int("count", len(output.RouteTables)))
		if aws.StringValue(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
