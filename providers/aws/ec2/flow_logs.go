package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type FlowLog struct {
	ID                       uint `gorm:"primarykey"`
	AccountID                string
	Region                   string
	CreationTime             *time.Time
	DeliverLogsErrorMessage  *string
	DeliverLogsPermissionArn *string
	DeliverLogsStatus        *string
	FlowLogId                *string
	FlowLogStatus            *string
	LogDestination           *string
	LogDestinationType       *string
	LogFormat                *string
	LogGroupName             *string
	MaxAggregationInterval   *int64
	ResourceId               *string
	Tags                     []*FlowLogTag `gorm:"constraint:OnDelete:CASCADE;"`
	TrafficType              *string
}

func (FlowLog) TableName() string {
	return "aws_ec2_flow_logs"
}

type FlowLogTag struct {
	ID        uint `gorm:"primarykey"`
	FlowLogID uint
	Key       *string
	Value     *string
}

func (FlowLogTag) TableName() string {
	return "aws_ec2_flow_log_tags"
}

func (c *Client) transformFlowLogTag(value *ec2.Tag) *FlowLogTag {
	return &FlowLogTag{
		Key:   value.Key,
		Value: value.Value,
	}
}

func (c *Client) transformFlowLogTags(values []*ec2.Tag) []*FlowLogTag {
	var tValues []*FlowLogTag
	for _, v := range values {
		tValues = append(tValues, c.transformFlowLogTag(v))
	}
	return tValues
}

func (c *Client) transformFlowLog(value *ec2.FlowLog) *FlowLog {
	return &FlowLog{
		Region:                   c.region,
		AccountID:                c.accountID,
		CreationTime:             value.CreationTime,
		DeliverLogsErrorMessage:  value.DeliverLogsErrorMessage,
		DeliverLogsPermissionArn: value.DeliverLogsPermissionArn,
		DeliverLogsStatus:        value.DeliverLogsStatus,
		FlowLogId:                value.FlowLogId,
		FlowLogStatus:            value.FlowLogStatus,
		LogDestination:           value.LogDestination,
		LogDestinationType:       value.LogDestinationType,
		LogFormat:                value.LogFormat,
		LogGroupName:             value.LogGroupName,
		MaxAggregationInterval:   value.MaxAggregationInterval,
		ResourceId:               value.ResourceId,
		Tags:                     c.transformFlowLogTags(value.Tags),
		TrafficType:              value.TrafficType,
	}
}

func (c *Client) transformFlowLogs(values []*ec2.FlowLog) []*FlowLog {
	var tValues []*FlowLog
	for _, v := range values {
		tValues = append(tValues, c.transformFlowLog(v))
	}
	return tValues
}

func (c *Client) FlowLogs(gConfig interface{}) error {
	var config ec2.DescribeFlowLogsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["ec2FlowLog"] {
		err := c.db.AutoMigrate(
			&FlowLog{},
			&FlowLogTag{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["ec2FlowLog"] = true
	}
	for {
		output, err := c.svc.DescribeFlowLogs(&config)
		if err != nil {
			return err
		}
		c.db.Where("region = ?", c.region).Where("account_id = ?", c.accountID).Delete(&FlowLog{})
		common.ChunkedCreate(c.db, c.transformFlowLogs(output.FlowLogs))
		c.log.Info("Fetched resources", zap.String("resource", "ec2.flow_logs"), zap.Int("count", len(output.FlowLogs)))
		if aws.StringValue(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
