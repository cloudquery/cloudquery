package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

type RouteTable struct {
	ID              uint                     `gorm:"primarykey"`
	AccountID       string                   `neo:"unique"`
	Region          string                   `neo:"unique"`
	Associations    []*RouteTableAssociation `gorm:"constraint:OnDelete:CASCADE;"`
	OwnerId         *string
	PropagatingVgws []*RouteTablePropagatingVgw `gorm:"constraint:OnDelete:CASCADE;"`
	RouteTableId    *string                     `neo:"unique"`
	Routes          []*RouteTableRoute          `gorm:"constraint:OnDelete:CASCADE;"`
	Tags            []*RouteTableTag            `gorm:"constraint:OnDelete:CASCADE;"`
	VpcId           *string
}

func (RouteTable) TableName() string {
	return "aws_ec2_route_tables"
}

type RouteTableAssociation struct {
	ID           uint   `gorm:"primarykey"`
	RouteTableID uint   `neo:"ignore"`
	AccountID    string `gorm:"-" neo:"unique"`
	Region       string `gorm:"-" neo:"unique"`

	State         *string
	StatusMessage *string

	GatewayId               *string
	Main                    *bool
	RouteTableAssociationId *string `neo:"unique"`
	RouteTableId            *string
	SubnetId                *string
}

func (RouteTableAssociation) TableName() string {
	return "aws_ec2_route_table_associations"
}

type RouteTablePropagatingVgw struct {
	ID           uint   `gorm:"primarykey"`
	RouteTableID uint   `neo:"ignore"`
	AccountID    string `gorm:"-"`
	Region       string `gorm:"-"`

	GatewayId *string
}

func (RouteTablePropagatingVgw) TableName() string {
	return "aws_ec2_route_table_propagation_vgws"
}

type RouteTableRoute struct {
	ID           uint   `gorm:"primarykey"`
	RouteTableID uint   `neo:"ignore"`
	AccountID    string `gorm:"-"`
	Region       string `gorm:"-"`

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

func (RouteTableRoute) TableName() string {
	return "aws_ec2_route_table_routes"
}

type RouteTableTag struct {
	ID           uint `gorm:"primarykey"`
	RouteTableID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Key   *string
	Value *string
}

func (RouteTableTag) TableName() string {
	return "aws_ec2_route_table_tags"
}

func (c *Client) transformRouteTableAssociation(value *ec2.RouteTableAssociation) *RouteTableAssociation {
	res := RouteTableAssociation{
		AccountID:               c.accountID,
		Region:                  c.region,
		GatewayId:               value.GatewayId,
		Main:                    value.Main,
		RouteTableAssociationId: value.RouteTableAssociationId,
		RouteTableId:            value.RouteTableId,
		SubnetId:                value.SubnetId,
	}

	if value.AssociationState != nil {
		res.State = value.AssociationState.State
		res.StatusMessage = value.AssociationState.StatusMessage
	}

	return &res
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
		AccountID: c.accountID,
		Region:    c.region,
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
		AccountID:                   c.accountID,
		Region:                      c.region,
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
		AccountID: c.accountID,
		Region:    c.region,
		Key:       value.Key,
		Value:     value.Value,
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

var RouteTableTables = []interface{}{
	&RouteTable{},
	&RouteTableAssociation{},
	&RouteTablePropagatingVgw{},
	&RouteTableRoute{},
	&RouteTableTag{},
}

func (c *Client) routeTables(gConfig interface{}) error {
	var config ec2.DescribeRouteTablesInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(RouteTableTables...)
	for {
		output, err := c.svc.DescribeRouteTables(&config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformRouteTables(output.RouteTables))
		c.log.Info("Fetched resources", zap.String("resource", "ec2.route_tables"), zap.Int("count", len(output.RouteTables)))
		if aws.StringValue(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
