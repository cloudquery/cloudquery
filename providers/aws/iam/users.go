package iam

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/gocarina/gocsv"
	"go.uber.org/zap"
	"time"
)

type User struct {
	ID                   uint    `gorm:"primarykey"`
	AccountID            string  `neo:"unique"`
	Arn                  *string `neo:"unique"`
	PasswordEnabled      *bool
	PasswordLastUsed     *time.Time
	PasswordLastChanged  *time.Time
	PasswordNextRotation *time.Time
	MFAActive            *bool
	CreateDate           *time.Time
	Path                 *string

	PermissionsBoundaryArn  *string
	PermissionsBoundaryType *string

	Tags       []*UserTag `gorm:"constraint:OnDelete:CASCADE;"`
	UserId     *string
	UserName   *string          `csv:"user"`
	AccessKeys []*UserAccessKey `gorm:"constraint:OnDelete:CASCADE;"`
}

func (User) TableName() string {
	return "aws_iam_users"
}

type UserAccessKey struct {
	ID        uint   `gorm:"primarykey"`
	UserID    uint   `neo:"ignore"`
	AccountID string `gorm:"-"`

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
	ID        uint   `gorm:"primarykey"`
	UserID    uint   `neo:"ignore"`
	AccountID string `gorm:"-"`

	Key   *string
	Value *string
}

func (UserTag) TableName() string {
	return "aws_iam_user_tags"
}

func (c *Client) transformAccessKey(value *iam.AccessKeyMetadata) (*UserAccessKey, error) {
	output, err := c.svc.GetAccessKeyLastUsed(&iam.GetAccessKeyLastUsedInput{AccessKeyId: value.AccessKeyId})
	if err != nil {
		return nil, err
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
	return &res, nil
}

func (c *Client) transformAccessKeys(values []*iam.AccessKeyMetadata) ([]*UserAccessKey, error) {
	var tValues []*UserAccessKey
	for _, v := range values {
		tValue, err := c.transformAccessKey(v)
		if err != nil {
			return nil, err
		}
		tValues = append(tValues, tValue)
	}
	return tValues, nil
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

func (c *Client) transformReportUser(reportUser *ReportUser) (*User, error) {
	//var err error
	location, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}

	createDate := reportUser.UserCreationTime.In(location)
	res := User{
		AccountID:  c.accountID,
		Arn:        &reportUser.ARN,
		MFAActive:  &reportUser.MFAActive,
		CreateDate: &createDate,
		UserName:   &reportUser.User,
	}

	if reportUser.User != "<root_account>" {
		output, err := c.svc.GetUser(&iam.GetUserInput{UserName: &reportUser.User})
		if err != nil {
			return nil, err
		}
		res.Path = output.User.Path
		if output.User.PermissionsBoundary != nil {
			res.PermissionsBoundaryType = output.User.PermissionsBoundary.PermissionsBoundaryType
			res.PermissionsBoundaryArn = output.User.PermissionsBoundary.PermissionsBoundaryArn
		}
		res.Tags = c.transformUserTags(output.User.Tags)
		res.UserId = output.User.UserId

		outputAccessKeys, err := c.svc.ListAccessKeys(&iam.ListAccessKeysInput{
			UserName: &reportUser.User,
		})
		if err != nil {
			return nil, err
		}
		res.AccessKeys, err = c.transformAccessKeys(outputAccessKeys.AccessKeyMetadata)
		if err != nil {
			return nil, err
		}
	}

	switch reportUser.PasswordEnabled {
	case "FALSE", "false":
		passwordEnabled := false
		res.PasswordEnabled = &passwordEnabled
	case "TRUE", "true":
		passwordEnabled := true
		res.PasswordEnabled = &passwordEnabled
	}

	passwordLastUsed, err := time.ParseInLocation(time.RFC3339, reportUser.PasswordLastUsed, location)
	if err == nil {
		res.PasswordLastUsed = &passwordLastUsed
	}
	passwordLastChanged, err := time.ParseInLocation(time.RFC3339, reportUser.PasswordLastChanged, location)
	if err == nil {
		res.PasswordLastChanged = &passwordLastChanged
	}

	passwordNextRotation, err := time.ParseInLocation(time.RFC3339, reportUser.PasswordNextRotation, location)
	if err == nil {
		res.PasswordNextRotation = &passwordNextRotation
	}

	lastRotated1, err := time.ParseInLocation(time.RFC3339, reportUser.AccessKey1LastRotated, location)
	if err == nil && len(res.AccessKeys) > 0 {
		res.AccessKeys[0].LastRotated = &lastRotated1
	}

	lastRotated2, err := time.ParseInLocation(time.RFC3339, reportUser.AccessKey2LastRotated, location)
	if err == nil && len(res.AccessKeys) > 1 {
		res.AccessKeys[1].LastRotated = &lastRotated2
	}

	return &res, nil
}

func (c *Client) transformReportUsers(values []*ReportUser) ([]*User, error) {
	var tValues []*User
	for _, v := range values {
		tValue, err := c.transformReportUser(v)
		if err != nil {
			return nil, err
		}
		tValues = append(tValues, tValue)
	}
	return tValues, nil
}

var UserTables = []interface{}{
	&User{},
	&UserAccessKey{},
	&UserTag{},
}

func (c *Client) users(_ interface{}) error {
	var err error
	var reportOutput *iam.GetCredentialReportOutput

	c.db.Where("account_id", c.accountID).Delete(UserTables...)
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

	tValues, err := c.transformReportUsers(users)
	if err != nil {
		return err
	}
	c.db.ChunkedCreate(tValues)
	c.log.Info("Fetched resources", zap.String("resource", "iam.users"), zap.Int("count", len(users)))
	return nil
}
