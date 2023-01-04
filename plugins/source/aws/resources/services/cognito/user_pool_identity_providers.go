package cognito

import (
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func UserPoolIdentityProviders() *schema.Table {
	return &schema.Table{
		Name:                "aws_cognito_user_pool_identity_providers",
		Description:         `https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_IdentityProviderType.html`,
		Resolver:            fetchCognitoUserPoolIdentityProviders,
		PreResourceResolver: getUserPoolIdentityProvider,
		Multiplex:           client.ServiceAccountRegionMultiplexer("cognito-identity"),
		Transform: transformers.TransformWithStruct(&types.IdentityProviderType{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "user_pool_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
