package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func PasswordPolicies() *schema.Table {
	tableName := "aws_iam_password_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_PasswordPolicy.html`,
		Resolver:    fetchIamPasswordPolicies,
		Transform:   transformers.TransformWithStruct(&models.PasswordPolicyWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
		},
	}
}

func fetchIamPasswordPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config iam.GetAccountPasswordPolicyInput
	c := meta.(*client.Client)
	svc := c.Services().Iam
	response, err := svc.GetAccountPasswordPolicy(ctx, &config)
	if err != nil {
		if c.IsNotFoundError(err) {
			res <- models.PasswordPolicyWrapper{PolicyExists: false}
			return nil
		}
		return err
	}
	res <- models.PasswordPolicyWrapper{PasswordPolicy: *response.PasswordPolicy, PolicyExists: true}
	return nil
}
