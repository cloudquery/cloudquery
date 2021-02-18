package fsx

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type Backup struct {
	_            interface{} `neo:"raw:MERGE (a:AWSAccount {account_id: $account_id}) MERGE (a) - [:Resource] -> (n)"`
	ID           uint        `gorm:"primarykey"`
	AccountID    string
	Region       string
	BackupId     *string
	CreationTime *time.Time

	ActiveDirectoryId         *string
	ActiveDirectoryDomainName *string

	FailureDetailsMessage *string

	KmsKeyId        *string
	Lifecycle       *string
	ProgressPercent *int32
	ResourceARN     *string      `neo:"unique"`
	Tags            []*BackupTag `gorm:"constraint:OnDelete:CASCADE;"`
	Type            *string
}

func (Backup) TableName() string {
	return "aws_fsx_backups"
}

type BackupTag struct {
	ID        uint   `gorm:"primarykey"`
	BackupID  uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`
	Key       *string
	Value     *string
}

func (BackupTag) TableName() string {
	return "aws_fsx_backup_tags"
}

func (c *Client) transformBackupTags(values *[]types.Tag) []*BackupTag {
	var tValues []*BackupTag
	for _, value := range *values {
		tValues = append(tValues, &BackupTag{
			Region:    c.region,
			AccountID: c.accountID,
			Key:       value.Key,
			Value:     value.Value,
		})
	}
	return tValues
}


func (c *Client) transformBackups(values *[]types.Backup) []*Backup {
	var tValues []*Backup
	for _, value := range *values {
		res := Backup{
			Region:          c.region,
			AccountID:       c.accountID,
			BackupId:        value.BackupId,
			CreationTime:    value.CreationTime,
			KmsKeyId:        value.KmsKeyId,
			Lifecycle:       aws.String(string(value.Lifecycle)),
			ProgressPercent: value.ProgressPercent,
			ResourceARN:     value.ResourceARN,
			Tags:            c.transformBackupTags(&value.Tags),
			Type:            aws.String(string(value.Type)),
		}

		if value.DirectoryInformation != nil {
			res.ActiveDirectoryDomainName = value.DirectoryInformation.DomainName
			res.ActiveDirectoryId = value.DirectoryInformation.ActiveDirectoryId
		}

		if value.FailureDetails != nil {
			res.FailureDetailsMessage = value.FailureDetails.Message
		}
		tValues = append(tValues, &res)
	}
	return tValues
}

var BackupTables = []interface{}{
	&Backup{},
	&BackupTag{},
}

func (c *Client) backups(gConfig interface{}) error {
	var config fsx.DescribeBackupsInput
	ctx := context.Background()
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(BackupTables...)

	for {
		output, err := c.svc.DescribeBackups(ctx, &config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformBackups(&output.Backups))
		c.log.Info("Fetched resources", zap.String("resource", "fsx.backups"), zap.Int("count", len(output.Backups)))
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
