package cloudtrail

import (
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Trail struct {
	ID                         uint `gorm:"primarykey"`
	AccountID                  string
	Region                     string
	CloudWatchLogsLogGroupArn  *string
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
	TrailARN                   *string
}

func (Trail) TableName() string {
	return "aws_cloudtrail_trails"
}

func (c *Client) transformTrail(value *cloudtrail.Trail) *Trail {
	res := Trail{
		Region:                     c.region,
		AccountID:                  c.accountID,
		CloudWatchLogsLogGroupArn:  value.CloudWatchLogsLogGroupArn,
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
	}

	return &res
}

func (c *Client) transformTrails(values []*cloudtrail.Trail) []*Trail {
	var tValues []*Trail
	for _, v := range values {
		tValues = append(tValues, c.transformTrail(v))
	}
	return tValues
}

func MigrateTrails(db *gorm.DB) error {
	return db.AutoMigrate(
		&Trail{},
	)
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
	c.db.Where("region = ?", c.region).Where("account_id = ?", c.accountID).Delete(&Trail{})
	common.ChunkedCreate(c.db, c.transformTrails(output.TrailList))
	c.log.Info("Fetched resources", zap.String("resource", "cloudtrail.trails"), zap.Int("count", len(output.TrailList)))

	return nil
}
