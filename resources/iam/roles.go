package iam

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/mitchellh/mapstructure"
)

type Role struct {
	ID                       uint    `gorm:"primarykey"`
	AccountID                string  `neo:"unique"`
	Arn                      *string `neo:"unique"`
	AssumeRolePolicyDocument *string
	CreateDate               *time.Time
	Description              *string
	MaxSessionDuration       *int32
	Path                     *string

	PermissionsBoundaryArn  *string
	PermissionsBoundaryType *string

	RoleId         *string `neo:"unique"`
	LastUsedDate   *time.Time
	LastUsedRegion *string

	RoleName *string
	Tags     []*RoleTag    `gorm:"constraint:OnDelete:CASCADE;"`
	Policies []*RolePolicy `gorm:"constraint:OnDelete:CASCADE;"`
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

type RolePolicy struct {
	ID         uint   `gorm:"primarykey"`
	RoleID     uint   `neo:"ignore"`
	AccountID  string `gorm:"-"`
	PolicyArn  *string
	PolicyName *string
}

func (RolePolicy) TableName() string {
	return "aws_iam_role_policies"
}

func (c *Client) transformRolesPolicies(values *[]types.AttachedPolicy) []*RolePolicy {
	var tValues []*RolePolicy
	for _, value := range *values {
		tValue := RolePolicy{
			AccountID:  c.accountID,
			PolicyArn:  value.PolicyArn,
			PolicyName: value.PolicyName,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformRoleTags(values *[]types.Tag) []*RoleTag {
	var tValues []*RoleTag
	for _, value := range *values {
		tValues = append(tValues, &RoleTag{
			AccountID: c.accountID,
			Key:       value.Key,
			Value:     value.Value,
		})
	}
	return tValues
}

func (c *Client) transformRoles(values *[]types.Role) ([]*Role, error) {
	var tValues []*Role
	ctx := context.Background()
	for _, value := range *values {
		var decodedDocument string = ""
		var err error = nil
		if value.AssumeRolePolicyDocument != nil {
			decodedDocument, err = url.QueryUnescape(*value.AssumeRolePolicyDocument)
			if err != nil {
				return nil, err
			}
		}

		tValue := Role{
			AccountID:                c.accountID,
			Arn:                      value.Arn,
			AssumeRolePolicyDocument: &decodedDocument,
			CreateDate:               value.CreateDate,
			Description:              value.Description,
			MaxSessionDuration:       value.MaxSessionDuration,
			Path:                     value.Path,
			RoleId:                   value.RoleId,
			RoleName:                 value.RoleName,
		}

		if value.PermissionsBoundary != nil {
			tValue.PermissionsBoundaryArn = value.PermissionsBoundary.PermissionsBoundaryArn
			tValue.PermissionsBoundaryType = aws.String(string(value.PermissionsBoundary.PermissionsBoundaryType))
		}

		if value.RoleLastUsed != nil {
			tValue.LastUsedDate = value.RoleLastUsed.LastUsedDate
			tValue.LastUsedRegion = value.RoleLastUsed.Region
		}

		listAttachedRolePoliciesInput := iam.ListAttachedRolePoliciesInput{
			RoleName: value.RoleName,
		}
		for {
			outputAttachedPolicies, err := c.svc.ListAttachedRolePolicies(ctx, &listAttachedRolePoliciesInput)
			if err != nil {
				return nil, err
			}
			tValue.Policies = append(tValue.Policies, c.transformRolesPolicies(&outputAttachedPolicies.AttachedPolicies)...)
			if outputAttachedPolicies.Marker == nil {
				break
			}
			listAttachedRolePoliciesInput.Marker = outputAttachedPolicies.Marker
		}

		listRoleTagsInput := iam.ListRoleTagsInput{
			RoleName: value.RoleName,
		}
		for {
			outputRoleTags, err := c.svc.ListRoleTags(ctx, &listRoleTagsInput)
			if err != nil {
				return nil, err
			}
			tValue.Tags = append(tValue.Tags, c.transformRoleTags(&outputRoleTags.Tags)...)
			if outputRoleTags.Marker == nil {
				break
			}
			listRoleTagsInput.Marker = outputRoleTags.Marker
		}

		tValues = append(tValues, &tValue)
	}
	return tValues, nil
}

var RoleTables = []interface{}{
	&Role{},
	&RoleTag{},
	&RolePolicy{},
}

func (c *Client) roles(gConfig interface{}) error {
	ctx := context.Background()
	var config iam.ListRolesInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("account_id", c.accountID).Delete(RoleTables...)

	for {
		output, err := c.svc.ListRoles(ctx, &config)
		if err != nil {
			return err
		}
		tValues, err := c.transformRoles(&output.Roles)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(tValues)
		c.log.Info("Fetched resources", "resource", "iam.roles", "count", len(output.Roles))
		if aws.ToString(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
