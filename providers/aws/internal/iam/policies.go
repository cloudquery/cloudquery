package iam

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/mitchellh/mapstructure"
	"github.com/cloudquery/cloudquery/providers/common"
	"go.uber.org/zap"
	"time"
)

type Policy struct {
	ID                            uint `gorm:"primarykey"`
	AccountID                     string
	Region                        string
	Arn                           *string
	AttachmentCount               *int64
	CreateDate                    *time.Time
	DefaultVersionId              *string
	Description                   *string
	IsAttachable                  *bool
	Path                          *string
	PermissionsBoundaryUsageCount *int64
	PolicyId                      *string
	PolicyName                    *string
	UpdateDate                    *time.Time
}

func (c *Client) transformPolicy(value *iam.Policy) *Policy {
	return &Policy{
		Region:                        c.region,
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

func (c *Client) transformPolicys(values []*iam.Policy) []*Policy {
	var tValues []*Policy
	for _, v := range values {
		tValues = append(tValues, c.transformPolicy(v))
	}
	return tValues
}

func (c *Client) Policys(gConfig interface{}) error {
	var config iam.ListPoliciesInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["iamPolicy"] {
		err := c.db.AutoMigrate(
			&Policy{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["iamPolicy"] = true
	}
	for {
		output, err := c.svc.ListPolicies(&config)
		if err != nil {
			return err
		}
		c.log.Debug("deleting previous Policys", zap.String("region", c.region), zap.String("account_id", c.accountID))
		c.db.Where("region = ?", c.region).Where("account_id = ?", c.accountID).Delete(&Policy{})
		common.ChunkedCreate(c.db, c.transformPolicys(output.Policies))
		c.log.Info("populating Policys", zap.Int("count", len(output.Policies)))
		if aws.StringValue(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
