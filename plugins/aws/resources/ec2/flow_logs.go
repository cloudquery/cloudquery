package ec2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/mitchellh/mapstructure"
	"time"
)

type FlowLog struct {
	_                        interface{} `neo:"raw:MERGE (a:AWSAccount {account_id: $account_id}) MERGE (a) - [:Resource] -> (n)"`
	ID                       uint        `gorm:"primarykey"`
	AccountID                string      `neo:"unique"`
	Region                   string      `neo:"unique"`
	CreationTime             *time.Time
	DeliverLogsErrorMessage  *string
	DeliverLogsPermissionArn *string
	DeliverLogsStatus        *string
	FlowLogId                *string `neo:"unique"`
	FlowLogStatus            *string
	LogDestination           *string
	LogDestinationType       *string
	LogFormat                *string
	LogGroupName             *string
	MaxAggregationInterval   int32
	ResourceId               *string
	Tags                     []*FlowLogTag `gorm:"constraint:OnDelete:CASCADE;"`
	TrafficType              *string
}

func (FlowLog) TableName() string {
	return "aws_ec2_flow_logs"
}

type FlowLogTag struct {
	ID        uint `gorm:"primarykey"`
	FlowLogID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Key   *string
	Value *string
}

func (FlowLogTag) TableName() string {
	return "aws_ec2_flow_log_tags"
}

func (c *Client) transformFlowLogTags(values *[]types.Tag) []*FlowLogTag {
	var tValues []*FlowLogTag
	for _, v := range *values {
		tValues = append(tValues, &FlowLogTag{
			AccountID: c.accountID,
			Region:    c.region,
			Key:       v.Key,
			Value:     v.Value,
		})
	}
	return tValues
}

func (c *Client) transformFlowLogs(values *[]types.FlowLog) []*FlowLog {
	var tValues []*FlowLog
	for _, v := range *values {
		tValues = append(tValues, &FlowLog{
			Region:                   c.region,
			AccountID:                c.accountID,
			CreationTime:             v.CreationTime,
			DeliverLogsErrorMessage:  v.DeliverLogsErrorMessage,
			DeliverLogsPermissionArn: v.DeliverLogsPermissionArn,
			DeliverLogsStatus:        v.DeliverLogsStatus,
			FlowLogId:                v.FlowLogId,
			FlowLogStatus:            v.FlowLogStatus,
			LogDestination:           v.LogDestination,
			LogDestinationType:       aws.String(string(v.LogDestinationType)),
			LogFormat:                v.LogFormat,
			LogGroupName:             v.LogGroupName,
			MaxAggregationInterval:   v.MaxAggregationInterval,
			ResourceId:               v.ResourceId,
			Tags:                     c.transformFlowLogTags(&v.Tags),
			TrafficType:              aws.String(string(v.TrafficType)),
		})
	}
	return tValues
}

var FlowLogsTables = []interface{}{
	&FlowLog{},
	&FlowLogTag{},
}

func (c *Client) FlowLogs(gConfig interface{}) error {
	var config ec2.DescribeFlowLogsInput
	ctx := context.Background()
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(FlowLogsTables...)
	for {
		output, err := c.svc.DescribeFlowLogs(ctx, &config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformFlowLogs(&output.FlowLogs))
		c.log.Info("Fetched resources", "resource", "ec2.flow_logs", "count", len(output.FlowLogs))
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
