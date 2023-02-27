package iam

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func PasswordPolicies() *schema.Table {
	return &schema.Table{
		Name:        "aws_iam_password_policies",
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_PasswordPolicy.html`,
		Resolver:    fetchIamPasswordPolicies,
		Transform:   transformers.TransformWithStruct(&models.PasswordPolicyWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Multiplex:   client.ServiceAccountRegionMultiplexer("iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
		},
	}
}
