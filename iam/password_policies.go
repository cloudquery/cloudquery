package iam

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

type PasswordPolicy struct {
	ID                         uint   `gorm:"primarykey"`
	AccountID                  string `neo:"unique"`
	AllowUsersToChangePassword *bool
	ExpirePasswords            *bool
	HardExpiry                 *bool
	MaxPasswordAge             *int32
	MinimumPasswordLength      *int32
	PasswordReusePrevention    *int32
	RequireLowercaseCharacters *bool
	RequireNumbers             *bool
	RequireSymbols             *bool
	RequireUppercaseCharacters *bool
}

func (PasswordPolicy) TableName() string {
	return "aws_iam_password_policies"
}

func (c *Client) transformPasswordPolicy(value *types.PasswordPolicy) *PasswordPolicy {
	return &PasswordPolicy{
		AccountID:                  c.accountID,
		AllowUsersToChangePassword: &value.AllowUsersToChangePassword,
		ExpirePasswords:            &value.ExpirePasswords,
		HardExpiry:                 value.HardExpiry,
		MaxPasswordAge:             value.MaxPasswordAge,
		MinimumPasswordLength:      value.MinimumPasswordLength,
		PasswordReusePrevention:    value.PasswordReusePrevention,
		RequireLowercaseCharacters: &value.RequireLowercaseCharacters,
		RequireNumbers:             &value.RequireNumbers,
		RequireSymbols:             &value.RequireSymbols,
		RequireUppercaseCharacters: &value.RequireUppercaseCharacters,
	}
}

var PasswordPolicyTables = []interface{}{
	&PasswordPolicy{},
}

func (c *Client) passwordPolicies(gConfig interface{}) error {
	var config iam.GetAccountPasswordPolicyInput
	ctx := context.Background()
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	output, err := c.svc.GetAccountPasswordPolicy(ctx, &config)
	if err != nil {
		var ae smithy.APIError
		if errors.As(err, &ae) && ae.ErrorCode() == "NoSuchEntity" {
			c.log.Info("Fetched PasswordPolicy", zap.Int("count", 0))
			return nil
		}
		return err
	}
	c.db.Where("account_id", c.accountID).Delete(PasswordPolicyTables...)
	c.db.InsertOne(c.transformPasswordPolicy(output.PasswordPolicy))
	c.log.Info("Fetched resources", zap.String("resource", "iam.password_policies"), zap.Int("count", 1))

	return nil
}
