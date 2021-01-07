package cloudtrail

import (
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"regexp"
	"time"
)

type Trail struct {
	_                          interface{} `neo:"raw:MERGE (a:AWSAccount {account_id: $account_id}) MERGE (a) - [:Resource] -> (n)"`
	ID                         uint        `gorm:"primarykey"`
	AccountID                  string
	Region                     string
	CloudWatchLogsLogGroupArn  *string
	CloudWatchLogsLogGroupName *string
	CloudWatchLogsRoleArn      *string
	HasCustomEventSelectors    *bool
	HasInsightSelectors        *bool
	HomeRegion                 *string
	IncludeGlobalServiceEvents *bool
	IsMultiRegionTrail         *bool
	IsOrganizationTrail        *bool
	KmsKeyId                   *string
	LogFileValidationEnabled   *bool
	Name                       *string
	S3BucketName               *string
	S3KeyPrefix                *string
	SnsTopicARN                *string
	SnsTopicName               *string
	TrailARN                   *string `neo:"unique"`
	EventSelectors 			   []*EventSelector `gorm:"constraint:OnDelete:CASCADE;"`
	IsLogging *bool

	// Status
	LatestCloudWatchLogsDeliveryError *string
	LatestCloudWatchLogsDeliveryTime *time.Time
	LatestDeliveryAttemptSucceeded *string
	LatestDeliveryAttemptTime *string
	LatestDeliveryError *string
	LatestDeliveryTime *time.Time
	LatestDigestDeliveryError *string
	LatestDigestDeliveryTime *time.Time
	LatestNotificationAttemptSucceeded *string
	LatestNotificationAttemptTime *string
	LatestNotificationError *string
	LatestNotificationTime *time.Time
	StartLoggingTime *time.Time
	StopLoggingTime *time.Time
	TimeLoggingStarted *string
	TimeLoggingStopped *string
}

func (Trail) TableName() string {
	return "aws_cloudtrail_trails"
}
//log-group:([a-zA-Z0-9/]+):
var groupNameRegex = regexp.MustCompile("arn:aws:logs:[a-z0-9-]+:[0-9]+:log-group:([a-zA-Z0-9-/]+):")

func (c *Client) transformTrails(values []*cloudtrail.Trail) ([]*Trail, error) {
	var tValues []*Trail
	for _, value := range values {
		groupName := groupNameRegex.FindStringSubmatch(*value.CloudWatchLogsLogGroupArn)[1]
		statusOutput, err := c.svc.GetTrailStatus(&cloudtrail.GetTrailStatusInput{Name: value.TrailARN})
		if err != nil {
			return nil, err
		}
		res := Trail{
			Region:                     c.region,
			AccountID:                  c.accountID,
			CloudWatchLogsLogGroupArn:  value.CloudWatchLogsLogGroupArn,
			CloudWatchLogsLogGroupName: &groupName,
			CloudWatchLogsRoleArn:      value.CloudWatchLogsRoleArn,
			HasCustomEventSelectors:    value.HasCustomEventSelectors,
			HasInsightSelectors:        value.HasInsightSelectors,
			HomeRegion:                 value.HomeRegion,
			IncludeGlobalServiceEvents: value.IncludeGlobalServiceEvents,
			IsMultiRegionTrail:         value.IsMultiRegionTrail,
			IsOrganizationTrail:        value.IsOrganizationTrail,
			KmsKeyId:                   value.KmsKeyId,
			LogFileValidationEnabled:   value.LogFileValidationEnabled,
			Name:                       value.Name,
			S3BucketName:               value.S3BucketName,
			S3KeyPrefix:                value.S3KeyPrefix,
			SnsTopicARN:                value.SnsTopicARN,
			SnsTopicName:               value.SnsTopicName,
			TrailARN:                   value.TrailARN,
			IsLogging: statusOutput.IsLogging,
			LatestCloudWatchLogsDeliveryError: statusOutput.LatestCloudWatchLogsDeliveryError,
			LatestCloudWatchLogsDeliveryTime: statusOutput.LatestCloudWatchLogsDeliveryTime,
			LatestDeliveryAttemptSucceeded: statusOutput.LatestDeliveryAttemptSucceeded,
			LatestDeliveryAttemptTime: statusOutput.LatestDeliveryAttemptTime,
			LatestDeliveryError: statusOutput.LatestDeliveryError,
			LatestDeliveryTime: statusOutput.LatestDeliveryTime,
			LatestDigestDeliveryError: statusOutput.LatestDigestDeliveryError,
			LatestDigestDeliveryTime: statusOutput.LatestDigestDeliveryTime,
			LatestNotificationAttemptSucceeded: statusOutput.LatestNotificationAttemptSucceeded,
			LatestNotificationAttemptTime: statusOutput.LatestNotificationAttemptTime,
			LatestNotificationError: statusOutput.LatestNotificationError,
			LatestNotificationTime: statusOutput.LatestNotificationTime,
			StartLoggingTime: statusOutput.StartLoggingTime,
			StopLoggingTime: statusOutput.StopLoggingTime,
			TimeLoggingStarted: statusOutput.TimeLoggingStarted,
			TimeLoggingStopped: statusOutput.TimeLoggingStopped,
		}

		output, err := c.svc.GetEventSelectors(&cloudtrail.GetEventSelectorsInput{TrailName: value.TrailARN})
		if err != nil {
			return nil, err
		}
		res.EventSelectors = c.transformEventSelectors(output.EventSelectors)

		tValues = append(tValues, &res)
	}

	return tValues, nil
}

var TrailTables = []interface{}{
	&Trail{},

	&EventSelector{},
}

func (c *Client) trails(gConfig interface{}) error {
	var config cloudtrail.DescribeTrailsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	output, err := c.svc.DescribeTrails(&config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(TrailTables...)
	tValues, err := c.transformTrails(output.TrailList)
	if err != nil {
		return err
	}
	c.db.ChunkedCreate(tValues)
	c.log.Info("Fetched resources", zap.String("resource", "cloudtrail.trails"), zap.Int("count", len(output.TrailList)))

	return nil
}
