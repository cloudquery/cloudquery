package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cq-provider-aws/client"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IamOpenidConnectIdentityProviders() *schema.Table {
	return &schema.Table{
		Name:          "aws_iam_openid_connect_identity_providers",
		Description:   "IAM OIDC identity providers are entities in IAM that describe an external identity provider (IdP) service that supports the OpenID Connect (OIDC) standard, such as Google or Salesforce.",
		Resolver:      fetchIamOpenidConnectIdentityProviders,
		Multiplex:     client.AccountMultiplex,
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "arn",
				Description: "Amazon Resource Name (ARN) of the openid connect identity provider.",
				Type:        schema.TypeString,
			},
			{
				Name:        "client_id_list",
				Description: "A list of client IDs (also known as audiences) that are associated with the specified IAM OIDC provider resource object. For more information, see CreateOpenIDConnectProvider. ",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("ClientIDList"),
			},
			{
				Name:        "create_date",
				Description: "The date and time when the IAM OIDC provider resource object was created in the AWS account. ",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "tags",
				Description: "A list of tags that are attached to the specified IAM OIDC provider. The returned list of tags is sorted by tag key. For more information about tagging, see Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the IAM User Guide. ",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "thumbprint_list",
				Description: "A list of certificate thumbprints that are associated with the specified IAM OIDC provider resource object. For more information, see CreateOpenIDConnectProvider. ",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "url",
				Description: "The URL that the IAM OIDC provider resource object is associated with. For more information, see CreateOpenIDConnectProvider. ",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIamOpenidConnectIdentityProviders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().IAM
	response, err := svc.ListOpenIDConnectProviders(ctx, &iam.ListOpenIDConnectProvidersInput{})
	if err != nil {
		return diag.WrapError(err)
	}

	for _, p := range response.OpenIDConnectProviderList {
		providerResponse, err := svc.GetOpenIDConnectProvider(ctx, &iam.GetOpenIDConnectProviderInput{OpenIDConnectProviderArn: p.Arn})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- IamOpenIdIdentityProviderWrapper{providerResponse, *p.Arn}
	}
	return nil
}

type IamOpenIdIdentityProviderWrapper struct {
	*iam.GetOpenIDConnectProviderOutput
	Arn string
}
