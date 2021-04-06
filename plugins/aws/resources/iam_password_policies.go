package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/plugin/schema"
)

func IamPasswordPolicies() *schema.Table {
	return &schema.Table{
		Name:         "aws_iam_password_policies",
		Resolver:     fetchIamPasswordPolicies,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name: "allow_users_to_change_password",
				Type: schema.TypeBool,
			},
			{
				Name: "expire_passwords",
				Type: schema.TypeBool,
			},
			{
				Name: "hard_expiry",
				Type: schema.TypeBool,
			},
			{
				Name: "max_password_age",
				Type: schema.TypeInt,
			},
			{
				Name: "minimum_password_length",
				Type: schema.TypeInt,
			},
			{
				Name: "password_reuse_prevention",
				Type: schema.TypeInt,
			},
			{
				Name: "require_lowercase_characters",
				Type: schema.TypeBool,
			},
			{
				Name: "require_numbers",
				Type: schema.TypeBool,
			},
			{
				Name: "require_symbols",
				Type: schema.TypeBool,
			},
			{
				Name: "require_uppercase_characters",
				Type: schema.TypeBool,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIamPasswordPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var passwordPolicies []*types.PasswordPolicy
	var config iam.GetAccountPasswordPolicyInput
	svc := meta.(*client.Client).Services().IAM
	response, err := svc.GetAccountPasswordPolicy(ctx, &config)
	if err != nil {
		return err
	}
	passwordPolicies = append(passwordPolicies, response.PasswordPolicy)
	res <- passwordPolicies
	return nil
}
