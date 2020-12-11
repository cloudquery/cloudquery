package iam

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type User struct {
	ID                  uint `gorm:"primarykey"`
	AccountID           string
	Arn                 *string
	CreateDate          *time.Time
	PasswordLastUsed    *time.Time
	Path                *string
	PermissionsBoundary *iam.AttachedPermissionsBoundary `gorm:"embedded;embeddedPrefix:permissions_boundary_"`
	Tags                []*UserTag                       `gorm:"constraint:OnDelete:CASCADE;"`
	UserId              *string
	UserName            *string
}

type UserTag struct {
	ID     uint `gorm:"primarykey"`
	UserID uint
	Key    *string
	Value  *string
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

func (c *Client) transformUser(value *iam.User) *User {
	return &User{
		AccountID:           c.accountID,
		Arn:                 value.Arn,
		CreateDate:          value.CreateDate,
		PasswordLastUsed:    value.PasswordLastUsed,
		Path:                value.Path,
		PermissionsBoundary: value.PermissionsBoundary,
		Tags:                c.transformUserTags(value.Tags),
		UserId:              value.UserId,
		UserName:            value.UserName,
	}
}

func (c *Client) transformUsers(values []*iam.User) []*User {
	var tValues []*User
	for _, v := range values {
		tValues = append(tValues, c.transformUser(v))
	}
	return tValues
}

func (c *Client) users(gConfig interface{}) error {
	var config iam.ListUsersInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
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
	for {
		output, err := c.svc.ListUsers(&config)
		if err != nil {
			return err
		}
		c.db.Where("account_id = ?", c.accountID).Delete(&User{})
		common.ChunkedCreate(c.db, c.transformUsers(output.Users))
		c.log.Info("populating Users", zap.Int("count", len(output.Users)))
		if aws.StringValue(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
