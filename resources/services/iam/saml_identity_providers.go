package iam

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IamSamlIdentityProviders() *schema.Table {
	return &schema.Table{
		Name:         "aws_iam_saml_identity_providers",
		Description:  "SAML provider resource objects defined in IAM for the AWS account.",
		Resolver:     fetchIamSamlIdentityProviders,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "arn",
				Description: "Amazon Resource Name (ARN) of the saml identity provider.",
				Type:        schema.TypeString,
			},
			{
				Name:        "create_date",
				Description: "The date and time when the SAML provider was created. ",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "saml_metadata_document",
				Description: "The XML metadata document that includes information about an identity provider. ",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SAMLMetadataDocument"),
			},
			{
				Name:        "tags",
				Description: "A list of tags that are attached to the specified IAM SAML provider. The returned list of tags is sorted by tag key. For more information about tagging, see Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the IAM User Guide. ",
				Type:        schema.TypeJSON,
				Resolver:    resolveIamSamlIdentityProviderTags,
			},
			{
				Name:        "valid_until",
				Description: "The expiration date and time for the SAML provider. ",
				Type:        schema.TypeTimestamp,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIamSamlIdentityProviders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().IAM
	response, err := svc.ListSAMLProviders(ctx, &iam.ListSAMLProvidersInput{})
	if err != nil {
		return diag.WrapError(err)
	}

	for _, p := range response.SAMLProviderList {
		providerResponse, err := svc.GetSAMLProvider(ctx, &iam.GetSAMLProviderInput{SAMLProviderArn: p.Arn})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- IamSamlIdentityProviderWrapper{GetSAMLProviderOutput: providerResponse, Arn: *p.Arn}
	}
	return nil
}
func resolveIamSamlIdentityProviderTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(IamSamlIdentityProviderWrapper)
	if !ok {
		return fmt.Errorf("not iam identity provider")
	}
	response := make(map[string]interface{}, len(r.Tags))
	for _, t := range r.Tags {
		response[*t.Key] = t.Value
	}

	return resource.Set(c.Name, response)
}

type IamSamlIdentityProviderWrapper struct {
	*iam.GetSAMLProviderOutput
	Arn string
}
