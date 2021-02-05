package efs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/efs"
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
	NumberOfMountTargets         *int64
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

func (c *Client) transformFileSystemDescriptionTag(value *efs.Tag) *FileSystemDescriptionTag {
	return &FileSystemDescriptionTag{
		Region:    c.region,
		AccountID: c.accountID,
		Key:       value.Key,
		Value:     value.Value,
	}
}

func (c *Client) transformFileSystemDescriptionTags(values []*efs.Tag) []*FileSystemDescriptionTag {
	var tValues []*FileSystemDescriptionTag
	for _, v := range values {
		tValues = append(tValues, c.transformFileSystemDescriptionTag(v))
	}
	return tValues
}

func (c *Client) transformFileSystemDescription(value *efs.FileSystemDescription) *FileSystemDescription {
	res := FileSystemDescription{
		Region:                       c.region,
		AccountID:                    c.accountID,
		CreationTime:                 value.CreationTime,
		CreationToken:                value.CreationToken,
		Encrypted:                    value.Encrypted,
		FileSystemArn:                value.FileSystemArn,
		FileSystemId:                 value.FileSystemId,
		KmsKeyId:                     value.KmsKeyId,
		LifeCycleState:               value.LifeCycleState,
		Name:                         value.Name,
		NumberOfMountTargets:         value.NumberOfMountTargets,
		OwnerId:                      value.OwnerId,
		PerformanceMode:              value.PerformanceMode,
		ProvisionedThroughputInMibps: value.ProvisionedThroughputInMibps,
		Tags:                         c.transformFileSystemDescriptionTags(value.Tags),
		ThroughputMode:               value.ThroughputMode,
	}

	if value.SizeInBytes != nil {
		res.SizeInBytesTimestamp = value.SizeInBytes.Timestamp
		res.SizeInBytesValue = value.SizeInBytes.Value
		res.SizeInBytesValueInIA = value.SizeInBytes.ValueInIA
		res.SizeInBytesValueInStandard = value.SizeInBytes.ValueInStandard
	}

	return &res
}

func (c *Client) transformFileSystemDescriptions(values []*efs.FileSystemDescription) []*FileSystemDescription {
	var tValues []*FileSystemDescription
	for _, v := range values {
		tValues = append(tValues, c.transformFileSystemDescription(v))
	}
	return tValues
}

var FileSystemTables = []interface{}{
	&FileSystemDescription{},
	&FileSystemDescriptionTag{},
}

func (c *Client) fileSystems(gConfig interface{}) error {
	var config efs.DescribeFileSystemsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(FileSystemTables...)
	for {
		output, err := c.svc.DescribeFileSystems(&config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformFileSystemDescriptions(output.FileSystems))
		c.log.Info("Fetched resources", zap.String("resource", "efs.filesystems"), zap.Int("count", len(output.FileSystems)))
		if aws.StringValue(output.NextMarker) == "" {
			break
		}
		config.Marker = output.NextMarker
	}
	return nil
}
