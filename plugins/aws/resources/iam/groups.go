package iam

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/mitchellh/mapstructure"
	"time"
)

type Group struct {
	ID         uint `gorm:"primarykey"`
	AccountID  string
	Arn        *string `neo:"unique"`
	CreateDate *time.Time
	GroupId    *string
	GroupName  *string
	Path       *string
	Policies   []*GroupPolicy `gorm:"constraint:OnDelete:CASCADE;"`
}

func (Group) TableName() string {
	return "aws_iam_groups"
}

type GroupPolicy struct {
	ID         uint   `gorm:"primarykey"`
	GroupID    uint   `neo:"ignore"`
	AccountID  string `gorm:"-"`
	PolicyArn  *string
	PolicyName *string
}

func (GroupPolicy) TableName() string {
	return "aws_iam_group_policies"
}

func (c *Client) transformGroupPolicies(values *[]types.AttachedPolicy) []*GroupPolicy {
	var tValues []*GroupPolicy
	for _, value := range *values {
		tValue := GroupPolicy{
			AccountID:  c.accountID,
			PolicyArn:  value.PolicyArn,
			PolicyName: value.PolicyName,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformGroups(values *[]types.Group) ([]*Group, error) {
	var tValues []*Group
	ctx := context.Background()
	for _, value := range *values {
		tValue := &Group{
			AccountID:  c.accountID,
			Arn:        value.Arn,
			CreateDate: value.CreateDate,
			GroupId:    value.GroupId,
			GroupName:  value.GroupName,
			Path:       value.Path,
		}

		listAttachedUserPoliciesInput := iam.ListAttachedGroupPoliciesInput{
			GroupName: value.GroupName,
		}
		for {
			outputAttachedPolicies, err := c.svc.ListAttachedGroupPolicies(ctx, &listAttachedUserPoliciesInput)
			if err != nil {
				return nil, err
			}
			tValue.Policies = append(tValue.Policies, c.transformGroupPolicies(&outputAttachedPolicies.AttachedPolicies)...)
			if outputAttachedPolicies.Marker == nil {
				break
			}
			listAttachedUserPoliciesInput.Marker = outputAttachedPolicies.Marker
		}

		tValues = append(tValues, tValue)

	}
	return tValues, nil
}

var GroupTables = []interface{}{
	&Group{},
	&GroupPolicy{},
}

func (c *Client) groups(gConfig interface{}) error {
	var config iam.ListGroupsInput
	ctx := context.Background()
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("account_id", c.accountID).Delete(GroupTables...)

	for {
		output, err := c.svc.ListGroups(ctx, &config)
		if err != nil {
			return err
		}
		tValues, err := c.transformGroups(&output.Groups)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(tValues)
		c.log.Info("Fetched resources", "resource", "iam.groups", "count", len(output.Groups))
		if aws.ToString(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
