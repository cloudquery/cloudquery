package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IamSamlIdentityProviders() *schema.Table {
	return &schema.Table{
		Name:         "aws_iam_saml_identity_providers",
		Resolver:     fetchIamSamlIdentityProviders,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name: "create_date",
				Type: schema.TypeTimestamp,
			},
			{
				Name:     "saml_metadata_document",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SAMLMetadataDocument"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveIamSamlIdentityProviderTags,
			},
			{
				Name: "valid_until",
				Type: schema.TypeTimestamp,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIamSamlIdentityProviders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().IAM
	response, err := svc.ListSAMLProviders(ctx, &iam.ListSAMLProvidersInput{})
	if err != nil {
		return err
	}

	for _, p := range response.SAMLProviderList {
		providerResponse, err := svc.GetSAMLProvider(ctx, &iam.GetSAMLProviderInput{SAMLProviderArn: p.Arn})
		if err != nil {
			return err
		}
		res <- providerResponse
	}
	return nil
}
func resolveIamSamlIdentityProviderTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(*iam.GetSAMLProviderOutput)
	if !ok {
		return fmt.Errorf("not iam identity provider")
	}
	response := make(map[string]interface{}, len(r.Tags))
	for _, t := range r.Tags {
		response[*t.Key] = t.Value
	}

	resource.Set(c.Name, response)
	return nil
}
