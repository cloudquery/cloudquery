package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SamlIdentityProviders() *schema.Table {
	tableName := "aws_iam_saml_identity_providers"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/IAM/latest/APIReference/API_SAMLProviderListEntry.html`,
		Resolver:            fetchIamSamlIdentityProviders,
		PreResourceResolver: getSamlIdentityProvider,
		Transform: transformers.TransformWithStruct(
			&models.IAMSAMLIdentityProviderWrapper{},
			transformers.WithUnwrapAllEmbeddedStructs(),
			transformers.WithSkipFields("ResultMetadata"),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name: "arn",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchIamSamlIdentityProviders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Iam
	response, err := svc.ListSAMLProviders(ctx, &iam.ListSAMLProvidersInput{})
	if err != nil {
		return err
	}

	res <- response.SAMLProviderList
	return nil
}

func getSamlIdentityProvider(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	svc := meta.(*client.Client).Services().Iam
	p := resource.Item.(types.SAMLProviderListEntry)

	providerResponse, err := svc.GetSAMLProvider(ctx, &iam.GetSAMLProviderInput{SAMLProviderArn: p.Arn})
	if err != nil {
		return err
	}

	resource.Item = models.IAMSAMLIdentityProviderWrapper{
		GetSAMLProviderOutput: providerResponse,
		Arn:                   *p.Arn,
		Tags:                  client.TagsToMap(providerResponse.Tags),
	}
	return nil
}
