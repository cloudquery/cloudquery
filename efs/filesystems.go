package efs

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type FileSystemDescription struct {
	_                            interface{} `neo:"raw:MERGE (a:AWSAccount {account_id: $account_id}) MERGE (a) - [:Resource] -> (n)"`
	ID                           uint        `gorm:"primarykey"`
	AccountID                    string
	Region                       string
	CreationTime                 *time.Time
	CreationToken                *string
	Encrypted                    *bool
	FileSystemArn                *string `neo:"unique"`
	FileSystemId                 *string
	KmsKeyId                     *string
	LifeCycleState               *string
	Name                         *string
	NumberOfMountTargets         *int32
	OwnerId                      *string
	PerformanceMode              *string
	ProvisionedThroughputInMibps *float64

	SizeInBytesTimestamp       *time.Time
	SizeInBytesValue           *int64
	SizeInBytesValueInIA       *int64
	SizeInBytesValueInStandard *int64

	Tags           []*FileSystemDescriptionTag `gorm:"constraint:OnDelete:CASCADE;"`
	ThroughputMode *string
}

func (FileSystemDescription) TableName() string {
	return "aws_efs_file_system_descriptions"
}

type FileSystemDescriptionTag struct {
	ID                      uint   `gorm:"primarykey"`
	FileSystemDescriptionID uint   `neo:"ignore"`
	AccountID               string `gorm:"-"`
	Region                  string `gorm:"-"`

	Key   *string
	Value *string
}

func (FileSystemDescriptionTag) TableName() string {
	return "aws_efs_file_system_description_tags"
}

func (c *Client) transformFileSystemDescriptionTags(values *[]types.Tag) []*FileSystemDescriptionTag {
	var tValues []*FileSystemDescriptionTag
	for _, value := range *values {
		tValues = append(tValues, &FileSystemDescriptionTag{
			Region:    c.region,
			AccountID: c.accountID,
			Key:       value.Key,
			Value:     value.Value,
		})
	}
	return tValues
}

func (c *Client) transformFileSystemDescriptions(values *[]types.FileSystemDescription) []*FileSystemDescription {
	var tValues []*FileSystemDescription
	for _, value := range *values {
		res := FileSystemDescription{
			Region:                       c.region,
			AccountID:                    c.accountID,
			CreationTime:                 value.CreationTime,
			CreationToken:                value.CreationToken,
			Encrypted:                    value.Encrypted,
			FileSystemArn:                value.FileSystemArn,
			FileSystemId:                 value.FileSystemId,
			KmsKeyId:                     value.KmsKeyId,
			LifeCycleState:               aws.String(string(value.LifeCycleState)),
			Name:                         value.Name,
			NumberOfMountTargets:         &value.NumberOfMountTargets,
			OwnerId:                      value.OwnerId,
			PerformanceMode:              aws.String(string(value.PerformanceMode)),
			ProvisionedThroughputInMibps: value.ProvisionedThroughputInMibps,
			Tags:                         c.transformFileSystemDescriptionTags(&value.Tags),
			ThroughputMode:               aws.String(string(value.ThroughputMode)),
		}

		if value.SizeInBytes != nil {
			res.SizeInBytesTimestamp = value.SizeInBytes.Timestamp
			res.SizeInBytesValue = &value.SizeInBytes.Value
			res.SizeInBytesValueInIA = value.SizeInBytes.ValueInIA
			res.SizeInBytesValueInStandard = value.SizeInBytes.ValueInStandard
		}
		tValues = append(tValues, &res)
	}
	return tValues
}

var FileSystemTables = []interface{}{
	&FileSystemDescription{},
	&FileSystemDescriptionTag{},
}

func (c *Client) fileSystems(gConfig interface{}) error {
	ctx := context.Background()
	var config efs.DescribeFileSystemsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(FileSystemTables...)
	for {
		output, err := c.svc.DescribeFileSystems(ctx, &config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformFileSystemDescriptions(&output.FileSystems))
		c.log.Info("Fetched resources", zap.String("resource", "efs.filesystems"), zap.Int("count", len(output.FileSystems)))
		if aws.ToString(output.NextMarker) == "" {
			break
		}
		config.Marker = output.NextMarker
	}
	return nil
}
