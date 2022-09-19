package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

type PasswordPolicyWrapper struct {
	types.PasswordPolicy
	PolicyExists bool
}

func fetchIamPasswordPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.GetAccountPasswordPolicyInput
	c := meta.(*client.Client)
	svc := c.Services().IAM
	response, err := svc.GetAccountPasswordPolicy(ctx, &config)
	if err != nil {
		if c.IsNotFoundError(err) {
			res <- PasswordPolicyWrapper{types.PasswordPolicy{}, false}
			return nil
		}
		return err
	}
	res <- PasswordPolicyWrapper{*response.PasswordPolicy, true}
	return nil
}
