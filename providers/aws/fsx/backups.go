package fsx

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/fsx"
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
	ProgressPercent *int64
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

func (c *Client) transformBackupTag(value *fsx.Tag) *BackupTag {
	return &BackupTag{
		Region:    c.region,
		AccountID: c.accountID,
		Key:       value.Key,
		Value:     value.Value,
	}
}

func (c *Client) transformBackupTags(values []*fsx.Tag) []*BackupTag {
	var tValues []*BackupTag
	for _, v := range values {
		tValues = append(tValues, c.transformBackupTag(v))
	}
	return tValues
}

func (c *Client) transformBackup(value *fsx.Backup) *Backup {
	res := Backup{
		Region:          c.region,
		AccountID:       c.accountID,
		BackupId:        value.BackupId,
		CreationTime:    value.CreationTime,
		KmsKeyId:        value.KmsKeyId,
		Lifecycle:       value.Lifecycle,
		ProgressPercent: value.ProgressPercent,
		ResourceARN:     value.ResourceARN,
		Tags:            c.transformBackupTags(value.Tags),
		Type:            value.Type,
	}

	if value.DirectoryInformation != nil {
		res.ActiveDirectoryDomainName = value.DirectoryInformation.DomainName
		res.ActiveDirectoryId = value.DirectoryInformation.ActiveDirectoryId
	}

	if value.FailureDetails != nil {
		res.FailureDetailsMessage = value.FailureDetails.Message
	}

	return &res
}

func (c *Client) transformBackups(values []*fsx.Backup) []*Backup {
	var tValues []*Backup
	for _, v := range values {
		tValues = append(tValues, c.transformBackup(v))
	}
	return tValues
}

var BackupTables = []interface{}{
	&Backup{},
	&BackupTag{},
}

func (c *Client) backups(gConfig interface{}) error {
	var config fsx.DescribeBackupsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(BackupTables...)

	for {
		output, err := c.svc.DescribeBackups(&config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformBackups(output.Backups))
		c.log.Info("Fetched resources", zap.String("resource", "fsx.backups"), zap.Int("count", len(output.Backups)))
		if aws.StringValue(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
