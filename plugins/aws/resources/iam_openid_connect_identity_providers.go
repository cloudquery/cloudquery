package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IamOpenidConnectIdentityProviders() *schema.Table {
	return &schema.Table{
		Name:         "aws_iam_openid_connect_identity_providers",
		Resolver:     fetchIamOpenidConnectIdentityProviders,
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
				Name:     "client_id_list",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ClientIDList"),
			},
			{
				Name: "create_date",
				Type: schema.TypeTimestamp,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveIamOpenidConnectIdentityProviderTags,
			},
			{
				Name: "thumbprint_list",
				Type: schema.TypeStringArray,
			},
			{
				Name: "url",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIamOpenidConnectIdentityProviders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().IAM
	response, err := svc.ListOpenIDConnectProviders(ctx, &iam.ListOpenIDConnectProvidersInput{})
	if err != nil {
		return err
	}

	for _, p := range response.OpenIDConnectProviderList {
		providerResponse, err := svc.GetOpenIDConnectProvider(ctx, &iam.GetOpenIDConnectProviderInput{OpenIDConnectProviderArn: p.Arn})
		if err != nil {
			return err
		}
		res <- providerResponse
	}
	return nil
}
func resolveIamOpenidConnectIdentityProviderTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(*iam.GetOpenIDConnectProviderOutput)
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
