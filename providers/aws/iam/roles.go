package iam

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type Role struct {
	ID                       uint `gorm:"primarykey"`
	AccountID                string
	Region                   string
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

type RoleTag struct {
	ID     uint `gorm:"primarykey"`
	RoleID uint
	Key    *string
	Value  *string
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
		Region:                   c.region,
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

func (c *Client) roles(gConfig interface{}) error {
	var config iam.ListRolesInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["iamRole"] {
		err := c.db.AutoMigrate(
			&Role{},
			&RoleTag{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["iamRole"] = true
	}
	for {
		output, err := c.svc.ListRoles(&config)
		if err != nil {
			return err
		}
		c.log.Debug("deleting previous Roles", zap.String("region", c.region), zap.String("account_id", c.accountID))
		c.db.Where("region = ?", c.region).Where("account_id = ?", c.accountID).Delete(&Role{})
		common.ChunkedCreate(c.db, c.transformRoles(output.Roles))
		c.log.Info("populating Roles", zap.Int("count", len(output.Roles)))
		if aws.StringValue(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
