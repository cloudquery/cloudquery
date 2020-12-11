package efs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/efs"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type FileSystemDescription struct {
	ID                           uint `gorm:"primarykey"`
	AccountID                    string
	Region                       string
	CreationTime                 *time.Time
	CreationToken                *string
	Encrypted                    *bool
	FileSystemArn                *string
	FileSystemId                 *string
	KmsKeyId                     *string
	LifeCycleState               *string
	Name                         *string
	NumberOfMountTargets         *int64
	OwnerId                      *string
	PerformanceMode              *string
	ProvisionedThroughputInMibps *float64
	SizeInBytes                  *efs.FileSystemSize         `gorm:"embedded;embeddedPrefix:size_in_bytes_"`
	Tags                         []*FileSystemDescriptionTag `gorm:"constraint:OnDelete:CASCADE;"`
	ThroughputMode               *string
}

type FileSystemDescriptionTag struct {
	ID                      uint `gorm:"primarykey"`
	FileSystemDescriptionID uint
	Key                     *string
	Value                   *string
}

func (c *Client) transformFileSystemDescriptionTag(value *efs.Tag) *FileSystemDescriptionTag {
	return &FileSystemDescriptionTag{
		Key:   value.Key,
		Value: value.Value,
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
	return &FileSystemDescription{
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
		SizeInBytes:                  value.SizeInBytes,
		Tags:                         c.transformFileSystemDescriptionTags(value.Tags),
		ThroughputMode:               value.ThroughputMode,
	}
}

func (c *Client) transformFileSystemDescriptions(values []*efs.FileSystemDescription) []*FileSystemDescription {
	var tValues []*FileSystemDescription
	for _, v := range values {
		tValues = append(tValues, c.transformFileSystemDescription(v))
	}
	return tValues
}

func (c *Client) fileSystems(gConfig interface{}) error {
	var config efs.DescribeFileSystemsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["efsFileSystemDescription"] {
		err := c.db.AutoMigrate(
			&FileSystemDescription{},
			&FileSystemDescriptionTag{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["efsFileSystemDescription"] = true
	}
	for {
		output, err := c.svc.DescribeFileSystems(&config)
		if err != nil {
			return err
		}
		c.db.Where("region = ?", c.region).Where("account_id = ?", c.accountID).Delete(&FileSystemDescription{})
		common.ChunkedCreate(c.db, c.transformFileSystemDescriptions(output.FileSystems))
		c.log.Info("populating FileSystemDescriptions", zap.Int("count", len(output.FileSystems)))
		if aws.StringValue(output.NextMarker) == "" {
			break
		}
		config.Marker = output.NextMarker
	}
	return nil
}
