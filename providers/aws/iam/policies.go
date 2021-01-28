package iam

import (
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
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
	PolicyVersions                []*PolicyVersion `gorm:"constraint:OnDelete:CASCADE;"`
}

func (Policy) TableName() string {
	return "aws_iam_policies"
}

type PolicyVersion struct {
	ID               uint   `gorm:"primarykey`
	PolicyID         uint   `neo:"ignore"`
	AccountID        string `gorm:"-"`
	VersionID        *string
	CreateDate       *time.Time
	Document         *string
	IsDefaultVersion *bool
}

func (PolicyVersion) TableName() string {
	return "aws_iam_policy_versions"
}

func (c *Client) transformPolicyVersionList(values []*iam.PolicyVersion) ([]*PolicyVersion, error) {
	var tValues []*PolicyVersion

	for _, value := range values {
		decodedDocument, err := url.QueryUnescape(*value.Document)
		if err != nil {
			return nil, err
		}

		tValues = append(tValues, &PolicyVersion{
			AccountID:        c.accountID,
			VersionID:        value.VersionId,
			CreateDate:       value.CreateDate,
			Document:         &decodedDocument,
			IsDefaultVersion: value.IsDefaultVersion,
		})
	}
	return tValues, nil
}

func (c *Client) transformPolicy(value *iam.ManagedPolicyDetail) (*Policy, error) {
	tValue := &Policy{
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

	policyVersions, err := c.transformPolicyVersionList(value.PolicyVersionList)
	if err != nil {
		return nil, err
	}
	tValue.PolicyVersions = append(tValue.PolicyVersions, policyVersions...)

	return tValue, nil
}

func (c *Client) transformPolicies(values []*iam.ManagedPolicyDetail) ([]*Policy, error) {
	var tValues []*Policy
	for _, v := range values {
		policy, err := c.transformPolicy(v)
		if err != nil {
			return nil, err
		}
		tValues = append(tValues, policy)
	}
	return tValues, nil
}

var PolicyTables = []interface{}{
	&Policy{},
	&PolicyVersion{},
}

func (c *Client) policies(gConfig interface{}) error {
	var config iam.GetAccountAuthorizationDetailsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	c.db.Where("account_id", c.accountID).Delete(PolicyTables...)
	for {
		output, err := c.svc.GetAccountAuthorizationDetails(&config)
		if err != nil {
			return err
		}

		tValues, err := c.transformPolicies(output.Policies)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(tValues)

		c.log.Info("Fetched resources", zap.String("resource", "iam.policies"), zap.Int("count", len(output.Policies)))
		if aws.StringValue(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
