package iam

import (
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
		Transform:           client.TransformWithStruct(&models.IAMSAMLIdentityProviderWrapper{}, transformers.WithSkipFields("ResultMetadata")),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "iam"),
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
