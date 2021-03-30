package iam

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go"
	"github.com/gocarina/gocsv"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/mitchellh/mapstructure"
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

func (c *Client) transformAccessKeys(ctx context.Context, values *[]types.AccessKeyMetadata) ([]*UserAccessKey, error) {
	var tValues []*UserAccessKey
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

var UserTables = []interface{}{
	&User{},
	&UserAccessKey{},
	&UserTag{},

	&UserAttachedPolicy{},
	&UserGroup{},
}

func (c *Client) transformUser(ctx context.Context, user *types.User, reportUser *ReportUser) (*User, error) {
	location, err := time.LoadLocation("UTC")
	if err != nil {
		return nil, err
	}
	res := User{
		AccountID:        c.accountID,
		Arn:              user.Arn,
		CreateDate:       user.CreateDate,
		UserName:         user.UserName,
		PasswordLastUsed: user.PasswordLastUsed,
	}

	if *user.UserName != "<root_account>" {
		output, err := c.svc.GetUser(ctx, &iam.GetUserInput{UserName: user.UserName})
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
			UserName: user.UserName,
		})
		if err != nil {
			return nil, err
		}
		res.AccessKeys, err = c.transformAccessKeys(ctx, &outputAccessKeys.AccessKeyMetadata)
		if err != nil {
			return nil, err
		}

		listAttachedUserPoliciesInput := iam.ListAttachedUserPoliciesInput{
			UserName: user.UserName,
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
			UserName: user.UserName,
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
			UserName: user.UserName,
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

	if reportUser != nil {
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
	}

	return &res, nil
}

func getMatchingReportUser(reportUsers []*ReportUser, user types.User) *ReportUser {
	for _, reportUser := range reportUsers {
		if *user.Arn == reportUser.ARN {
			return reportUser
		}
	}
	return nil
}

func (c *Client) transformUsers(ctx context.Context, values []types.User, reportUsers []*ReportUser) ([]*User, error) {
	var tValues []*User
	for _, v := range values {
		tValue, err := c.transformUser(ctx, &v, getMatchingReportUser(reportUsers, v))
		if err != nil {
			return nil, err
		}
		tValues = append(tValues, tValue)
	}
	return tValues, nil
}

func (c *Client) GetCredentialReport(ctx context.Context) ([]*ReportUser, error) {
	var err error
	var apiErr smithy.APIError
	var reportOutput *iam.GetCredentialReportOutput
	for {
		reportOutput, err = c.svc.GetCredentialReport(ctx, &iam.GetCredentialReportInput{})
		if err != nil {
			if errors.As(err, &apiErr) {
				switch apiErr.ErrorCode() {
				case "ReportNotPresent", "ReportExpired":
					_, err := c.svc.GenerateCredentialReport(ctx, &iam.GenerateCredentialReportInput{})
					if err != nil {
						return nil, err
					}
				case "ReportInProgress":
					c.log.Info("Waiting for credential report to be generated", "resource", "iam.users")
					time.Sleep(2 * time.Second)
				default:
					return nil, err
				}
			} else {
				return nil, err
			}
		} else {
			break
		}
	}
	var users []*ReportUser
	err = gocsv.UnmarshalBytes(reportOutput.Content, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (c *Client) users(ctx context.Context, gConfig interface{}) error {
	var config iam.ListUsersInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	reportUsers, err := c.GetCredentialReport(ctx)
	if err != nil {
		return err
	}
	c.db.Where("account_id", c.accountID).Delete(UserTables...)
	for {
		output, err := c.svc.ListUsers(ctx, &config)
		if err != nil {
			return err
		}
		// Do transform users
		tValues, err := c.transformUsers(ctx, output.Users, reportUsers)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(tValues)
		c.log.Info("Fetched resources", "resource", "iam.users", "count", len(output.Users))
		if aws.ToString(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
