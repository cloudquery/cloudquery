package iam

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type Role struct {
	ID                       uint    `gorm:"primarykey"`
	AccountID                string  `neo:"unique"`
	Arn                      *string `neo:"unique"`
	AssumeRolePolicyDocument *string
	CreateDate               *time.Time
	Description              *string
	MaxSessionDuration       *int64
	Path                     *string

	PermissionsBoundaryArn  *string
	PermissionsBoundaryType *string

	RoleId         *string `neo:"unique"`
	LastUsedDate   *time.Time
	LastUsedRegion *string

	RoleName *string
	Tags     []*RoleTag `gorm:"constraint:OnDelete:CASCADE;"`
}

func (Role) TableName() string {
	return "aws_iam_roles"
}

type RoleTag struct {
	ID        uint   `gorm:"primarykey"`
	RoleID    uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Key       *string
	Value     *string
}

func (RoleTag) TableName() string {
	return "aws_iam_role_tags"
}

func (c *Client) transformRoleTag(value *iam.Tag) *RoleTag {
	return &RoleTag{
		AccountID: c.accountID,
		Key:       value.Key,
		Value:     value.Value,
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
	res := Role{
		AccountID:                c.accountID,
		Arn:                      value.Arn,
		AssumeRolePolicyDocument: value.AssumeRolePolicyDocument,
		CreateDate:               value.CreateDate,
		Description:              value.Description,
		MaxSessionDuration:       value.MaxSessionDuration,
		Path:                     value.Path,
		RoleId:                   value.RoleId,
		RoleName:                 value.RoleName,
		Tags:                     c.transformRoleTags(value.Tags),
	}

	if value.PermissionsBoundary != nil {
		res.PermissionsBoundaryArn = value.PermissionsBoundary.PermissionsBoundaryArn
		res.PermissionsBoundaryType = value.PermissionsBoundary.PermissionsBoundaryType
	}

	if value.RoleLastUsed != nil {
		res.LastUsedDate = value.RoleLastUsed.LastUsedDate
		res.LastUsedRegion = value.RoleLastUsed.Region
	}

	return &res
}

func (c *Client) transformRoles(values []*iam.Role) []*Role {
	var tValues []*Role
	for _, v := range values {
		tValues = append(tValues, c.transformRole(v))
	}
	return tValues
}

var RoleTables = []interface{}{
	&Role{},
	&RoleTag{},
}

func (c *Client) roles(gConfig interface{}) error {
	var config iam.ListRolesInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("account_id", c.accountID).Delete(RoleTables...)

	for {
		output, err := c.svc.ListRoles(&config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformRoles(output.Roles))
		c.log.Info("Fetched resources", zap.String("resource", "iam.roles"), zap.Int("count", len(output.Roles)))
		if aws.StringValue(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
