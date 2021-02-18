package elbv2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/cloudquery/cq-provider-aws/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type LoadBalancer struct {
	_                     interface{} `neo:"raw:MERGE (a:AWSAccount {account_id: $account_id}) MERGE (a) - [:Resource] -> (n)"`
	ID                    uint        `gorm:"primarykey"`
	AccountID             string
	Region                string
	AvailabilityZones     []*LoadBalancerAvailabilityZone `gorm:"constraint:OnDelete:CASCADE;"`
	CanonicalHostedZoneId *string
	CreatedTime           *time.Time
	CustomerOwnedIpv4Pool *string
	DNSName               *string
	IpAddressType         *string
	LoadBalancerArn       *string `neo:"unique"`
	LoadBalancerName      *string
	Scheme                *string
	SecurityGroups        *string

	StateCode   *string
	StateReason *string

	Type  *string
	VpcId *string
}

func (LoadBalancer) TableName() string {
	return "aws_elbv2_load_balancers"
}

type LoadBalancerAvailabilityZone struct {
	ID             uint   `gorm:"primarykey"`
	LoadBalancerID uint   `neo:"ignore"`
	AccountID      string `gorm:"-"`
	Region         string `gorm:"-"`

	LoadBalancerAddresses []*LoadBalancerAddress `gorm:"constraint:OnDelete:CASCADE;"`
	OutpostId             *string
	SubnetId              *string
	ZoneName              *string
}

func (LoadBalancerAvailabilityZone) TableName() string {
	return "aws_elbv2_load_balancer_av_zones"
}

type LoadBalancerAddress struct {
	ID                             uint   `gorm:"primarykey"`
	LoadBalancerAvailabilityZoneID uint   `neo:"ignore"`
	AccountID                      string `gorm:"-"`
	Region                         string `gorm:"-"`

	AllocationId       *string
	IpAddress          *string
	PrivateIPv4Address *string
}

func (LoadBalancerAddress) TableName() string {
	return "aws_elbv2_load_balancer_addresses"
}

func (c *Client) transformLoadBalancerAddresss(values *[]types.LoadBalancerAddress) []*LoadBalancerAddress {
	var tValues []*LoadBalancerAddress
	for _, value := range *values {
		tValues = append(tValues, &LoadBalancerAddress{
			AccountID:          c.accountID,
			Region:             c.region,
			AllocationId:       value.AllocationId,
			IpAddress:          value.IpAddress,
			PrivateIPv4Address: value.PrivateIPv4Address,
		})
	}
	return tValues
}

func (c *Client) transformLoadBalancerAvailabilityZones(values *[]types.AvailabilityZone) []*LoadBalancerAvailabilityZone {
	var tValues []*LoadBalancerAvailabilityZone
	for _, value := range *values {
		tValues = append(tValues, &LoadBalancerAvailabilityZone{
			AccountID:             c.accountID,
			Region:                c.region,
			LoadBalancerAddresses: c.transformLoadBalancerAddresss(&value.LoadBalancerAddresses),
			OutpostId:             value.OutpostId,
			SubnetId:              value.SubnetId,
			ZoneName:              value.ZoneName,
		})
	}
	return tValues
}

func (c *Client) transformLoadBalancers(values *[]types.LoadBalancer) []*LoadBalancer {
	var tValues []*LoadBalancer
	for _, value := range *values {
		res := LoadBalancer{
			Region:                c.region,
			AccountID:             c.accountID,
			AvailabilityZones:     c.transformLoadBalancerAvailabilityZones(&value.AvailabilityZones),
			CanonicalHostedZoneId: value.CanonicalHostedZoneId,
			CreatedTime:           value.CreatedTime,
			CustomerOwnedIpv4Pool: value.CustomerOwnedIpv4Pool,
			DNSName:               value.DNSName,
			IpAddressType:         aws.String(string(value.IpAddressType)),
			LoadBalancerArn:       value.LoadBalancerArn,
			LoadBalancerName:      value.LoadBalancerName,
			Scheme:                aws.String(string(value.Scheme)),
			SecurityGroups:        common.StringListToString(&value.SecurityGroups),
			Type:                  aws.String(string(value.Type)),
			VpcId:                 value.VpcId,
		}

		if value.State != nil {
			res.StateReason = value.State.Reason
			res.StateCode = aws.String(string(value.State.Code))
		}
		tValues = append(tValues, &res)
	}
	return tValues
}

var LoadBalancerTables = []interface{}{
	&LoadBalancer{},
	&LoadBalancerAvailabilityZone{},
	&LoadBalancerAddress{},
}

func (c *Client) loadBalancers(gConfig interface{}) error {
	ctx := context.Background()
	var config elbv2.DescribeLoadBalancersInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(LoadBalancerTables...)

	for {
		output, err := c.svc.DescribeLoadBalancers(ctx, &config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformLoadBalancers(&output.LoadBalancers))
		c.log.Info("Fetched resources", zap.String("resource", "elbv2.load_balancers"), zap.Int("count", len(output.LoadBalancers)))
		if aws.ToString(output.NextMarker) == "" {
			break
		}
		config.Marker = output.NextMarker
	}
	return nil
}
