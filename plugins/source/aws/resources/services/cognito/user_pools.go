package cognito

import (
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func UserPools() *schema.Table {
	tableName := "aws_cognito_user_pools"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UserPoolType.html`,
		Resolver:            fetchCognitoUserPools,
		PreResourceResolver: getUserPool,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "cognito-identity"),
		Transform:           client.TransformWithStruct(&types.UserPoolType{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
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
