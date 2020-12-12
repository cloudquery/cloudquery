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
	UserName             *string `csv:"user"`
}

func (User)TableName() string {
	return "aws_iam_users"
}

type UserAccessKey struct {
	ID     uint `gorm:"primarykey"`
	UserID uint
}

type UserTag struct {
	ID     uint `gorm:"primarykey"`
	UserID uint
	Key    *string
	Value  *string
}

func (UserTag)TableName() string {
	return "aws_iam_user_tags"
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
	User                 string    `csv:"user"`
	ARN                  string    `csv:"arn"`
	UserCreationTime     time.Time `csv:"user_creation_time"`
	PasswordEnabled      string    `csv:"password_enabled"`
	PasswordLastUsed     string    `csv:"password_last_used"`
	PasswordLastChanged  string    `csv:"password_last_changed"`
	PasswordNextRotation string    `csv:"password_next_rotation"`
	MFAActive            bool      `csv:"mfa_active"`
}

func (c *Client) transformReportUser(reportUser *ReportUser) *User {
	var output *iam.GetUserOutput
	var err error
	if reportUser.User != "<root_account>" {
		output, err = c.svc.GetUser(&iam.GetUserInput{UserName: &reportUser.User})
		if err != nil {
			log.Fatal(err)
		}
	}

	res := User{
		AccountID:  c.accountID,
		Arn:        &reportUser.ARN,
		MFAActive:  &reportUser.MFAActive,
		CreateDate: &reportUser.UserCreationTime,
		UserName:   &reportUser.User,
	}

	if output != nil {
		res.Path = output.User.Path
		res.PermissionsBoundary = output.User.PermissionsBoundary
		res.Tags = c.transformUserTags(output.User.Tags)
		res.UserId = output.User.UserId
	}

	switch reportUser.PasswordEnabled {
	case "FALSE", "false":
		passwordEnabled := false
		res.PasswordEnabled = &passwordEnabled
	case "TRUE", "true":
		passwordEnabled := true
		res.PasswordEnabled = &passwordEnabled
	}

	layout := "2020-10-21T12:15:57+00:00"
	tm, err := time.Parse(layout, reportUser.PasswordLastUsed)
	if err == nil {
		res.PasswordLastUsed = &tm
	}

	tm, err = time.Parse(layout, reportUser.PasswordLastChanged)
	if err == nil {
		res.PasswordLastChanged = &tm
	}

	tm, err = time.Parse(layout, reportUser.PasswordNextRotation)
	if err == nil {
		res.PasswordNextRotation = &tm
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
