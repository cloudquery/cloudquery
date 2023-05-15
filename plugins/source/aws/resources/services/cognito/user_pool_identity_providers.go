package cognito

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func userPoolIdentityProviders() *schema.Table {
	tableName := "aws_cognito_user_pool_identity_providers"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_IdentityProviderType.html`,
		Resolver:            fetchCognitoUserPoolIdentityProviders,
		PreResourceResolver: getUserPoolIdentityProvider,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "cognito-identity"),
		Transform:           transformers.TransformWithStruct(&types.IdentityProviderType{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "user_pool_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchCognitoUserPoolIdentityProviders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	pool := parent.Item.(*types.UserPoolType)
	c := meta.(*client.Client)
	svc := c.Services().Cognitoidentityprovider

	params := cognitoidentityprovider.ListIdentityProvidersInput{UserPoolId: pool.Id}
	paginator := cognitoidentityprovider.NewListIdentityProvidersPaginator(svc, &params)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *cognitoidentityprovider.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- page.Providers
	}
	return nil
}

func getUserPoolIdentityProvider(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Cognitoidentityprovider
	item := resource.Item.(types.ProviderDescription)
	pool := resource.Parent.Item.(*types.UserPoolType)

	pd, err := svc.DescribeIdentityProvider(ctx, &cognitoidentityprovider.DescribeIdentityProviderInput{
		ProviderName: item.ProviderName,
		UserPoolId:   pool.Id,
	}, func(options *cognitoidentityprovider.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}

	resource.Item = pd.IdentityProvider
	return nil
}
