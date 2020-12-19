package iam

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type Role struct {
	ID                       uint `gorm:"primarykey"`
	AccountID                string
	Arn                      *string
	AssumeRolePolicyDocument *string
	CreateDate               *time.Time
	Description              *string
	MaxSessionDuration       *int64
	Path                     *string
	PermissionsBoundary      *iam.AttachedPermissionsBoundary `gorm:"embedded;embeddedPrefix:permissions_boundary_"`
	RoleId                   *string
	RoleLastUsed             *iam.RoleLastUsed `gorm:"embedded;embeddedPrefix:role_last_used_"`
	RoleName                 *string
	Tags                     []*RoleTag `gorm:"constraint:OnDelete:CASCADE;"`
}

func (Role) TableName() string {
	return "aws_iam_roles"
}

type RoleTag struct {
	ID     uint `gorm:"primarykey"`
	RoleID uint
	Key    *string
	Value  *string
}

func (RoleTag) TableName() string {
	return "aws_iam_role_tags"
}

func (c *Client) transformRoleTag(value *iam.Tag) *RoleTag {
	return &RoleTag{
		Key:   value.Key,
		Value: value.Value,
	}
}

func (c *Client) transformRoleTags(values []*iam.Tag) []*RoleTag {
	var tValues []*RoleTag
	for _, v := range values {
		tValues = append(tValues, c.transformRoleTag(v))
	}
	return tValues
}

func (c *Client) transformRole(value *iam.Role) *Role {
	return &Role{
		AccountID:                c.accountID,
		Arn:                      value.Arn,
		AssumeRolePolicyDocument: value.AssumeRolePolicyDocument,
		CreateDate:               value.CreateDate,
		Description:              value.Description,
		MaxSessionDuration:       value.MaxSessionDuration,
		Path:                     value.Path,
		PermissionsBoundary:      value.PermissionsBoundary,
		RoleId:                   value.RoleId,
		RoleLastUsed:             value.RoleLastUsed,
		RoleName:                 value.RoleName,
		Tags:                     c.transformRoleTags(value.Tags),
	}
}

func (c *Client) transformRoles(values []*iam.Role) []*Role {
	var tValues []*Role
	for _, v := range values {
		tValues = append(tValues, c.transformRole(v))
	}
	return tValues
}

func MigrateRoles(db *gorm.DB) error {
	return db.AutoMigrate(
		&Role{},
		&RoleTag{},
	)
}

func (c *Client) roles(gConfig interface{}) error {
	var config iam.ListRolesInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	for {
		output, err := c.svc.ListRoles(&config)
		if err != nil {
			return err
		}
		c.db.Where("account_id = ?", c.accountID).Delete(&Role{})
		common.ChunkedCreate(c.db, c.transformRoles(output.Roles))
		c.log.Info("Fetched resources", zap.String("resource", "iam.roles"), zap.Int("count", len(output.Roles)))
		if aws.StringValue(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
