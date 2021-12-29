package iam

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IamPasswordPolicies() *schema.Table {
	return &schema.Table{
		Name:         "aws_iam_password_policies",
		Description:  "Contains information about the account password policy.",
		Resolver:     fetchIamPasswordPolicies,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "allow_users_to_change_password",
				Description: "Specifies whether IAM users are allowed to change their own password. ",
				Type:        schema.TypeBool,
			},
			{
				Name:        "expire_passwords",
				Description: "Indicates whether passwords in the account expire. Returns true if MaxPasswordAge contains a value greater than 0. Returns false if MaxPasswordAge is 0 or not present. ",
				Type:        schema.TypeBool,
			},
			{
				Name:        "hard_expiry",
				Description: "Specifies whether IAM users are prevented from setting a new password after their password has expired. ",
				Type:        schema.TypeBool,
			},
			{
				Name:        "max_password_age",
				Description: "The number of days that an IAM user password is valid. ",
				Type:        schema.TypeInt,
			},
			{
				Name:        "minimum_password_length",
				Description: "Minimum length to require for IAM user passwords. ",
				Type:        schema.TypeInt,
			},
			{
				Name:        "password_reuse_prevention",
				Description: "Specifies the number of previous passwords that IAM users are prevented from reusing. ",
				Type:        schema.TypeInt,
			},
			{
				Name:        "require_lowercase_characters",
				Description: "Specifies whether IAM user passwords must contain at least one lowercase character (a to z). ",
				Type:        schema.TypeBool,
			},
			{
				Name:        "require_numbers",
				Description: "Specifies whether IAM user passwords must contain at least one numeric character (0 to 9). ",
				Type:        schema.TypeBool,
			},
			{
				Name:        "require_symbols",
				Description: "Specifies whether IAM user passwords must contain at least one of the following symbols: ! @ # $ % ^ & * ( ) _ + - = [ ] { } | ' ",
				Type:        schema.TypeBool,
			},
			{
				Name:        "require_uppercase_characters",
				Description: "Specifies whether IAM user passwords must contain at least one uppercase character (A to Z). ",
				Type:        schema.TypeBool,
			},
			{
				Name:        "policy_exists",
				Description: "Specifies whether IAM user passwords configuration exists",
				Type:        schema.TypeBool,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIamPasswordPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config iam.GetAccountPasswordPolicyInput
	svc := meta.(*client.Client).Services().IAM
	response, err := svc.GetAccountPasswordPolicy(ctx, &config)
	if err != nil {
		var ae smithy.APIError
		if !errors.As(err, &ae) && ae.ErrorCode() != "NoSuchEntity" {
			return err
		}
		res <- PasswordPolicy{types.PasswordPolicy{}, false}
	} else {
		res <- PasswordPolicy{*response.PasswordPolicy, true}
	}

	return nil
}

type PasswordPolicy struct {
	types.PasswordPolicy
	PolicyExists bool
}
