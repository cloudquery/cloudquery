package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam/models"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchIamOpenidConnectIdentityProviders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Iam
	response, err := svc.ListOpenIDConnectProviders(ctx, &iam.ListOpenIDConnectProvidersInput{})
	if err != nil {
		return err
	}

	res <- response.OpenIDConnectProviderList
	return nil
}

func getOpenIdConnectIdentityProvider(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	svc := meta.(*client.Client).Services().Iam

	p := resource.Item.(types.OpenIDConnectProviderListEntry)
	providerResponse, err := svc.GetOpenIDConnectProvider(ctx, &iam.GetOpenIDConnectProviderInput{OpenIDConnectProviderArn: p.Arn})
	if err != nil {
		return err
	}

	resource.Item = &models.IamOpenIdIdentityProviderWrapper{GetOpenIDConnectProviderOutput: providerResponse, Arn: *p.Arn}
	return nil
}
