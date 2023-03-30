package iam

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func OpenidConnectIdentityProviders() *schema.Table {
	tableName := "aws_iam_openid_connect_identity_providers"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetOpenIDConnectProvider.html`,
		Resolver:            fetchIamOpenidConnectIdentityProviders,
		PreResourceResolver: getOpenIdConnectIdentityProvider,
		Transform:           transformers.TransformWithStruct(&models.IamOpenIdIdentityProviderWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
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
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}
