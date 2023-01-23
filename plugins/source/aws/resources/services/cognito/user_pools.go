package cognito

import (
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func UserPools() *schema.Table {
	return &schema.Table{
		Name:                "aws_cognito_user_pools",
		Description:         `https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UserPoolType.html`,
		Resolver:            fetchCognitoUserPools,
		PreResourceResolver: getUserPool,
		Multiplex:           client.ServiceAccountRegionMultiplexer("cognito-identity"),
		Transform:           transformers.TransformWithStruct(&types.UserPoolType{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			UserPoolIdentityProviders(),
		},
	}
}
