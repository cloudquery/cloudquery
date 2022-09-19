package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

type IAMSAMLIdentityProviderWrapper struct {
	*iam.GetSAMLProviderOutput
	Arn string
}

func fetchIamSamlIdentityProviders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
		res <- IAMSAMLIdentityProviderWrapper{GetSAMLProviderOutput: providerResponse, Arn: *p.Arn}
	}
	return nil
}
