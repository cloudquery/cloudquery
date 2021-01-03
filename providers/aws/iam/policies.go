package iam

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type Policy struct {
	ID                            uint    `gorm:"primarykey"`
	AccountID                     string  `neo:"unique"`
	Arn                           *string `neo:"unique"`
	AttachmentCount               *int64
	CreateDate                    *time.Time
	DefaultVersionId              *string
	Description                   *string
	IsAttachable                  *bool
	Path                          *string
	PermissionsBoundaryUsageCount *int64
	PolicyId                      *string `neo:"unique"`
	PolicyName                    *string
	UpdateDate                    *time.Time
}

func (Policy) TableName() string {
	return "aws_iam_policies"
}

func (c *Client) transformPolicy(value *iam.Policy) *Policy {
	return &Policy{
		AccountID:                     c.accountID,
		Arn:                           value.Arn,
		AttachmentCount:               value.AttachmentCount,
		CreateDate:                    value.CreateDate,
		DefaultVersionId:              value.DefaultVersionId,
		Description:                   value.Description,
		IsAttachable:                  value.IsAttachable,
		Path:                          value.Path,
		PermissionsBoundaryUsageCount: value.PermissionsBoundaryUsageCount,
		PolicyId:                      value.PolicyId,
		PolicyName:                    value.PolicyName,
		UpdateDate:                    value.UpdateDate,
	}
}

func (c *Client) transformPolicies(values []*iam.Policy) []*Policy {
	var tValues []*Policy
	for _, v := range values {
		tValues = append(tValues, c.transformPolicy(v))
	}
	return tValues
}

var PolicyTables = []interface{}{
	&Policy{},
}

func (c *Client) policies(gConfig interface{}) error {
	var config iam.ListPoliciesInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	for {
		output, err := c.svc.ListPolicies(&config)
		if err != nil {
			return err
		}
		c.db.Where("account_id", c.accountID).Delete(PolicyTables...)
		c.db.ChunkedCreate(c.transformPolicies(output.Policies))
		c.log.Info("Fetched resources", zap.String("resource", "iam.policies"), zap.Int("count", len(output.Policies)))
		if aws.StringValue(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
