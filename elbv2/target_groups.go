package elbv2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

type TargetGroup struct {
	ID        uint `gorm:"primarykey"`
	AccountID string
	Region    string

	HealthCheckEnabled         *bool
	HealthCheckIntervalSeconds *int32
	HealthCheckPath            *string
	HealthCheckPort            *string
	HealthCheckProtocol        *string
	HealthCheckTimeoutSeconds  *int32
	HealthyThresholdCount      *int32
	LBArns                     []*TargetGroupLoadBalancerArns `gorm:"constraint:OnDelete:CASCADE;"`

	MatcherHttpCode         *string
	Port                    *int32
	Protocol                *string
	TargetGroupArn          *string `neo:"unique"`
	TargetGroupName         *string
	TargetType              *string
	UnhealthyThresholdCount *int32
	VpcId                   *string
}

func (TargetGroup) TableName() string {
	return "aws_elbv2_target_groups"
}

type TargetGroupLoadBalancerArns struct {
	ID            uint   `gorm:"primarykey"`
	TargetGroupID uint   `neo:"ignore"`
	AccountID     string `gorm:"-"`
	Region        string `gorm:"-"`
	Value         string
}

func (TargetGroupLoadBalancerArns) TableName() string {
	return "aws_elbv2_target_group_load_balancer_arns"
}

func (c *Client) transformTargetGroups(values *[]types.TargetGroup) []*TargetGroup {
	var tValues []*TargetGroup
	for _, value := range *values {
		tValue := TargetGroup{
			AccountID:                  c.accountID,
			Region:                     c.region,
			HealthCheckEnabled:         value.HealthCheckEnabled,
			HealthCheckIntervalSeconds: value.HealthCheckIntervalSeconds,
			HealthCheckPath:            value.HealthCheckPath,
			HealthCheckPort:            value.HealthCheckPort,
			HealthCheckProtocol:        aws.String(string(value.HealthCheckProtocol)),
			HealthCheckTimeoutSeconds:  value.HealthCheckTimeoutSeconds,
			HealthyThresholdCount:      value.HealthyThresholdCount,
			LBArns:                     c.transformTargetGroupLoadBalancerArns(&value.LoadBalancerArns),
			Port:                       value.Port,
			Protocol:                   aws.String(string(value.Protocol)),
			TargetGroupArn:             value.TargetGroupArn,
			TargetGroupName:            value.TargetGroupName,
			TargetType:                 aws.String(string(value.TargetType)),
			UnhealthyThresholdCount:    value.UnhealthyThresholdCount,
			VpcId:                      value.VpcId,
		}
		if value.Matcher != nil {

			tValue.MatcherHttpCode = value.Matcher.HttpCode
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}
func (c *Client) transformTargetGroupLoadBalancerArns(values *[]string) []*TargetGroupLoadBalancerArns {
	var tValues []*TargetGroupLoadBalancerArns
	for _, v := range *values {
		tValues = append(tValues, &TargetGroupLoadBalancerArns{
			AccountID: c.accountID,
			Region:    c.region,
			Value:     v,
		})
	}
	return tValues
}

var TargetGroupTables = []interface{}{
	&TargetGroup{},
	&TargetGroupLoadBalancerArns{},
}

func (c *Client) targetGroups(gConfig interface{}) error {
	ctx := context.Background()
	var config elbv2.DescribeTargetGroupsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(TargetGroupTables...)
	for {
		output, err := c.svc.DescribeTargetGroups(ctx, &config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformTargetGroups(&output.TargetGroups))
		c.log.Info("Fetched resources", zap.String("resource", "elbv2.target_groups"), zap.Int("count", len(output.TargetGroups)))
		if aws.ToString(output.NextMarker) == "" {
			break
		}
		config.Marker = output.NextMarker
	}

	return nil
}
