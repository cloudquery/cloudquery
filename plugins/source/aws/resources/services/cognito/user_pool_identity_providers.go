package cognito

import (
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func UserPoolIdentityProviders() *schema.Table {
	tableName := "aws_cognito_user_pool_identity_providers"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_IdentityProviderType.html`,
		Resolver:            fetchCognitoUserPoolIdentityProviders,
		PreResourceResolver: getUserPoolIdentityProvider,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "cognito-identity"),
		Transform:           client.TransformWithStruct(&types.IdentityProviderType{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "user_pool_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
