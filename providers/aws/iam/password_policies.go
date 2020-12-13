package iam

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

type PasswordPolicy struct {
	ID                         uint `gorm:"primarykey"`
	AccountID                  string
	AllowUsersToChangePassword *bool
	ExpirePasswords            *bool
	HardExpiry                 *bool
	MaxPasswordAge             *int64
	MinimumPasswordLength      *int64
	PasswordReusePrevention    *int64
	RequireLowercaseCharacters *bool
	RequireNumbers             *bool
	RequireSymbols             *bool
	RequireUppercaseCharacters *bool
}

func (PasswordPolicy) TableName() string {
	return "aws_iam_password_policies"
}

func (c *Client) transformPasswordPolicy(value *iam.PasswordPolicy) *PasswordPolicy {
	return &PasswordPolicy{
		AccountID:                  c.accountID,
		AllowUsersToChangePassword: value.AllowUsersToChangePassword,
		ExpirePasswords:            value.ExpirePasswords,
		HardExpiry:                 value.HardExpiry,
		MaxPasswordAge:             value.MaxPasswordAge,
		MinimumPasswordLength:      value.MinimumPasswordLength,
		PasswordReusePrevention:    value.PasswordReusePrevention,
		RequireLowercaseCharacters: value.RequireLowercaseCharacters,
		RequireNumbers:             value.RequireNumbers,
		RequireSymbols:             value.RequireSymbols,
		RequireUppercaseCharacters: value.RequireUppercaseCharacters,
	}
}

func (c *Client) passwordPolicies(gConfig interface{}) error {
	var config iam.GetAccountPasswordPolicyInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["iamPasswordPolicy"] {
		err := c.db.AutoMigrate(
			&PasswordPolicy{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["iamPasswordPolicy"] = true
	}

	output, err := c.svc.GetAccountPasswordPolicy(&config)
	if err != nil {
		if err.(awserr.Error).Code() != "NoSuchEntity" {
			c.log.Info("Fetched PasswordPolicy", zap.Int("count", 0))
			return nil
		}
		return err
	}
	c.db.Where("account_id = ?", c.accountID).Delete(&PasswordPolicy{})
	c.db.Create(c.transformPasswordPolicy(output.PasswordPolicy))
	c.log.Info("Fetched resources", zap.String("resource", "iam.password_policies"), zap.Int("count", 1))

	return nil
}
