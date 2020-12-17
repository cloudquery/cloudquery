package iam

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/gocarina/gocsv"
	"go.uber.org/zap"
	"log"
	"time"
)

type User struct {
	ID                   uint `gorm:"primarykey"`
	AccountID            string
	Arn                  *string
	PasswordEnabled      *bool
	PasswordLastUsed     *time.Time
	PasswordLastChanged  *time.Time
	PasswordNextRotation *time.Time
	MFAActive            *bool
	CreateDate           *time.Time
	Path                 *string
	PermissionsBoundary  *iam.AttachedPermissionsBoundary `gorm:"embedded;embeddedPrefix:permissions_boundary_"`
	Tags                 []*UserTag                       `gorm:"constraint:OnDelete:CASCADE;"`
	UserId               *string
	UserName             *string          `csv:"user"`
	AccessKeys           []*UserAccessKey `gorm:"constraint:OnDelete:CASCADE;"`
}

func (User) TableName() string {
	return "aws_iam_users"
}

type UserAccessKey struct {
	ID                  uint `gorm:"primarykey"`
	UserID              uint
	AccessKeyId         *string
	CreateDate          *time.Time
	Status              *string
	LastUsed            *time.Time
	LastRotated         *time.Time
	LastUsedServiceName *string
}

func (UserAccessKey) TableName() string {
	return "aws_iam_user_access_keys"
}

type UserTag struct {
	ID     uint `gorm:"primarykey"`
	UserID uint
	Key    *string
	Value  *string
}

func (UserTag) TableName() string {
	return "aws_iam_user_tags"
}

func (c *Client) transformAccessKey(value *iam.AccessKeyMetadata) *UserAccessKey {
	output, err := c.svc.GetAccessKeyLastUsed(&iam.GetAccessKeyLastUsedInput{AccessKeyId: value.AccessKeyId})
	if err != nil {
		log.Fatal(err)
	}

	res := UserAccessKey{
		AccessKeyId:         value.AccessKeyId,
		CreateDate:          value.CreateDate,
		Status:              value.Status,
		LastUsed:            output.AccessKeyLastUsed.LastUsedDate,
		LastUsedServiceName: output.AccessKeyLastUsed.ServiceName,
	}
	if output.AccessKeyLastUsed != nil {
		res.LastUsed = output.AccessKeyLastUsed.LastUsedDate
		res.LastUsedServiceName = output.AccessKeyLastUsed.ServiceName
	}
	return &res
}

func (c *Client) transformAccessKeys(values []*iam.AccessKeyMetadata) []*UserAccessKey {
	var tValues []*UserAccessKey
	for _, v := range values {
		tValues = append(tValues, c.transformAccessKey(v))
	}
	return tValues
}

func (c *Client) transformUserTag(value *iam.Tag) *UserTag {
	return &UserTag{
		Key:   value.Key,
		Value: value.Value,
	}
}

func (c *Client) transformUserTags(values []*iam.Tag) []*UserTag {
	var tValues []*UserTag
	for _, v := range values {
		tValues = append(tValues, c.transformUserTag(v))
	}
	return tValues
}

type ReportUser struct {
	User                  string    `csv:"user"`
	ARN                   string    `csv:"arn"`
	UserCreationTime      time.Time `csv:"user_creation_time"`
	PasswordEnabled       string    `csv:"password_enabled"`
	PasswordLastUsed      string    `csv:"password_last_used"`
	PasswordLastChanged   string    `csv:"password_last_changed"`
	PasswordNextRotation  string    `csv:"password_next_rotation"`
	MFAActive             bool      `csv:"mfa_active"`
	AccessKey1LastRotated string    `csv:"access_key_1_last_rotated"`
	AccessKey2LastRotated string    `csv:"access_key_2_last_rotated"`
}

func (c *Client) transformReportUser(reportUser *ReportUser) *User {
	var err error
	res := User{
		AccountID:  c.accountID,
		Arn:        &reportUser.ARN,
		MFAActive:  &reportUser.MFAActive,
		CreateDate: &reportUser.UserCreationTime,
		UserName:   &reportUser.User,
	}

	if reportUser.User != "<root_account>" {
		output, err := c.svc.GetUser(&iam.GetUserInput{UserName: &reportUser.User})
		if err != nil {
			log.Fatal(err)
		}
		res.Path = output.User.Path
		res.PermissionsBoundary = output.User.PermissionsBoundary
		res.Tags = c.transformUserTags(output.User.Tags)
		res.UserId = output.User.UserId

		outputAccessKeys, err := c.svc.ListAccessKeys(&iam.ListAccessKeysInput{
			UserName: &reportUser.User,
		})
		if err != nil {
			log.Fatal(err)
		}
		res.AccessKeys = c.transformAccessKeys(outputAccessKeys.AccessKeyMetadata)
	}

	switch reportUser.PasswordEnabled {
	case "FALSE", "false":
		passwordEnabled := false
		res.PasswordEnabled = &passwordEnabled
	case "TRUE", "true":
		passwordEnabled := true
		res.PasswordEnabled = &passwordEnabled
	}

	passwordLastUsed, err := time.Parse(time.RFC3339, reportUser.PasswordLastUsed)
	if err == nil {
		res.PasswordLastUsed = &passwordLastUsed
	}
	passwordLastChanged, err := time.Parse(time.RFC3339, reportUser.PasswordLastChanged)
	if err == nil {
		res.PasswordLastChanged = &passwordLastChanged
	}

	passwordNextRotation, err := time.Parse(time.RFC3339, reportUser.PasswordNextRotation)
	if err == nil {
		res.PasswordNextRotation = &passwordNextRotation
	}

	lastRotated1, err := time.Parse(time.RFC3339, reportUser.AccessKey1LastRotated)
	if err == nil && len(res.AccessKeys) > 0 {
		res.AccessKeys[0].LastRotated = &lastRotated1
	}

	lastRotated2, err := time.Parse(time.RFC3339, reportUser.AccessKey2LastRotated)
	if err == nil && len(res.AccessKeys) > 1 {
		res.AccessKeys[1].LastRotated = &lastRotated2
	}

	return &res
}

func (c *Client) transformReportUsers(values []*ReportUser) []*User {
	var tValues []*User
	for _, v := range values {
		tValues = append(tValues, c.transformReportUser(v))
	}
	return tValues
}

func (c *Client) users(_ interface{}) error {
	if !c.resourceMigrated["iamUser"] {
		err := c.db.AutoMigrate(
			&User{},
			&UserAccessKey{},
			&UserTag{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["iamUser"] = true
	}

	var err error
	var reportOutput *iam.GetCredentialReportOutput
	for {
		reportOutput, err = c.svc.GetCredentialReport(&iam.GetCredentialReportInput{})
		if err != nil {
			if err.(awserr.Error).Code() != iam.ErrCodeCredentialReportNotPresentException ||
				err.(awserr.Error).Code() != iam.ErrCodeCredentialReportExpiredException {
				_, err := c.svc.GenerateCredentialReport(&iam.GenerateCredentialReportInput{})
				if err != nil {
					return err
				}
			} else if err.(awserr.Error).Code() != iam.ErrCodeCredentialReportNotReadyException {
				c.log.Info("Waiting for credential report to be generated", zap.String("resource", "iam.users"))
				time.Sleep(2 * time.Second)
			} else {
				return err
			}
		} else {
			break
		}
	}
	var users []*ReportUser
	err = gocsv.UnmarshalBytes(reportOutput.Content, &users)
	if err != nil {
		return err
	}

	c.db.Where("account_id = ?", c.accountID).Delete(&User{})
	common.ChunkedCreate(c.db, c.transformReportUsers(users))
	c.log.Info("Fetched resources", zap.String("resource", "iam.users"), zap.Int("count", len(users)))
	return nil
}
