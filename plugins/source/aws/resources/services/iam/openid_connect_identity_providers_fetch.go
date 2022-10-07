package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam/models"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchIamOpenidConnectIdentityProviders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
		res <- models.IamOpenIdIdentityProviderWrapper{GetOpenIDConnectProviderOutput: providerResponse, Arn: *p.Arn}
	}
	return nil
}
