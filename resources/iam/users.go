package iam

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/gocarina/gocsv"
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

	Tags             []*UserTag `gorm:"constraint:OnDelete:CASCADE;"`
	UserId           *string
	UserName         *string               `csv:"user"`
	AccessKeys       []*UserAccessKey      `gorm:"constraint:OnDelete:CASCADE;"`
	AttachedPolicies []*UserAttachedPolicy `gorm:"constraint:OnDelete:CASCADE;"`
	Groups           []*UserGroup          `gorm:"constraint:OnDelete:CASCADE;"`
}

func (User) TableName() string {
	return "aws_iam_users"
}

type UserGroup struct {
	ID        uint   `gorm:"primarykey"`
	UserID    uint   `neo:"ignore"`
	AccountID string `gorm:"-"`

	Arn        *string
	CreateDate *time.Time
	GroupId    *string
	GroupName  *string
	Path       *string
}

func (UserGroup) TableName() string {
	return "aws_iam_user_groups"
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

func (c *Client) transformUserGroups(values *[]types.Group) []*UserGroup {
	var tValues []*UserGroup
	for _, value := range *values {
		tValues = append(tValues, &UserGroup{
			AccountID:  c.accountID,
			Arn:        value.Arn,
			CreateDate: value.CreateDate,
			GroupId:    value.GroupId,
			GroupName:  value.GroupName,
			Path:       value.Path,
		})
	}
	return tValues
}

func (c *Client) transformAccessKeys(values *[]types.AccessKeyMetadata) ([]*UserAccessKey, error) {
	var tValues []*UserAccessKey
	ctx := context.Background()
	for _, value := range *values {
		output, err := c.svc.GetAccessKeyLastUsed(ctx, &iam.GetAccessKeyLastUsedInput{AccessKeyId: value.AccessKeyId})
		if err != nil {
			return nil, err
		}

		res := UserAccessKey{
			AccessKeyId:         value.AccessKeyId,
			CreateDate:          value.CreateDate,
			Status:              aws.String(string(value.Status)),
			LastUsed:            output.AccessKeyLastUsed.LastUsedDate,
			LastUsedServiceName: output.AccessKeyLastUsed.ServiceName,
		}
		if output.AccessKeyLastUsed != nil {
			res.LastUsed = output.AccessKeyLastUsed.LastUsedDate
			res.LastUsedServiceName = output.AccessKeyLastUsed.ServiceName
		}
		tValues = append(tValues, &res)
	}
	return tValues, nil
}

func (c *Client) transformUserTags(values *[]types.Tag) []*UserTag {
	var tValues []*UserTag
	for _, value := range *values {
		tValues = append(tValues, &UserTag{
			Key:   value.Key,
			Value: value.Value,
		})
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
	AccessKey1Active      bool      `csv:"access_key_1_active"`
	AccessKey2Active      bool      `csv:"access_key_2_active"`
	AccessKey1LastRotated string    `csv:"access_key_1_last_rotated"`
	AccessKey2LastRotated string    `csv:"access_key_2_last_rotated"`
}

func (c *Client) transformReportUser(reportUser *ReportUser) (*User, error) {
	//var err error
	ctx := context.Background()
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
		output, err := c.svc.GetUser(ctx, &iam.GetUserInput{UserName: &reportUser.User})
		if err != nil {
			return nil, err
		}
		res.Path = output.User.Path
		if output.User.PermissionsBoundary != nil {
			res.PermissionsBoundaryType = aws.String(string(output.User.PermissionsBoundary.PermissionsBoundaryType))
			res.PermissionsBoundaryArn = output.User.PermissionsBoundary.PermissionsBoundaryArn
		}
		res.UserId = output.User.UserId

		outputAccessKeys, err := c.svc.ListAccessKeys(ctx, &iam.ListAccessKeysInput{
			UserName: &reportUser.User,
		})
		if err != nil {
			return nil, err
		}
		res.AccessKeys, err = c.transformAccessKeys(&outputAccessKeys.AccessKeyMetadata)
		if err != nil {
			return nil, err
		}

		listAttachedUserPoliciesInput := iam.ListAttachedUserPoliciesInput{
			UserName: &reportUser.User,
		}
		for {
			outputAttachedPolicies, err := c.svc.ListAttachedUserPolicies(ctx, &listAttachedUserPoliciesInput)
			if err != nil {
				return nil, err
			}
			res.AttachedPolicies = append(res.AttachedPolicies, c.transformAttachedPolicies(&outputAttachedPolicies.AttachedPolicies)...)
			if outputAttachedPolicies.Marker == nil {
				break
			}
			listAttachedUserPoliciesInput.Marker = outputAttachedPolicies.Marker
		}

		listGroupsForUserInput := iam.ListGroupsForUserInput{
			UserName: &reportUser.User,
		}
		for {
			outputListGroupsForUsers, err := c.svc.ListGroupsForUser(ctx, &listGroupsForUserInput)
			if err != nil {
				return nil, err
			}
			res.Groups = append(res.Groups, c.transformUserGroups(&outputListGroupsForUsers.Groups)...)
			if outputListGroupsForUsers.Marker == nil {
				break
			}
			listGroupsForUserInput.Marker = outputListGroupsForUsers.Marker
		}

		listUserTagsInput := iam.ListUserTagsInput{
			UserName: &reportUser.User,
		}
		for {
			outputUserTags, err := c.svc.ListUserTags(ctx, &listUserTagsInput)
			if err != nil {
				return nil, err
			}
			res.Tags = append(res.Tags, c.transformUserTags(&outputUserTags.Tags)...)
			if outputUserTags.Marker == nil {
				break
			}
			listUserTagsInput.Marker = outputUserTags.Marker
		}

	}

	switch strings.ToLower(reportUser.PasswordEnabled) {
	case "false":
		passwordEnabled := false
		res.PasswordEnabled = &passwordEnabled
	case "true":
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

	&UserAttachedPolicy{},
	&UserGroup{},
}

func (c *Client) users(_ interface{}) error {
	var err error
	var apiErr smithy.APIError
	var reportOutput *iam.GetCredentialReportOutput
	ctx := context.Background()
	c.db.Where("account_id", c.accountID).Delete(UserTables...)
	for {
		reportOutput, err = c.svc.GetCredentialReport(ctx, &iam.GetCredentialReportInput{})

		if err != nil {
			if errors.As(err, &apiErr) {
				switch apiErr.ErrorCode() {
				case "ReportNotPresent", "ReportExpired":
					_, err := c.svc.GenerateCredentialReport(ctx, &iam.GenerateCredentialReportInput{})
					if err != nil {
						return err
					}
				case "ReportInProgress":
					c.log.Info("Waiting for credential report to be generated", "resource", "iam.users")
					time.Sleep(2 * time.Second)
				default:
					return err
				}
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
	c.log.Info("Fetched resources", "resource", "iam.users", "count", len(users))
	return nil
}
