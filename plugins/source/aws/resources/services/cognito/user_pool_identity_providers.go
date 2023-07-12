package cognito

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
	cl := meta.(*client.Client)
	svc := cl.Services().Cognitoidentityprovider

	params := cognitoidentityprovider.ListIdentityProvidersInput{UserPoolId: pool.Id}
	paginator := cognitoidentityprovider.NewListIdentityProvidersPaginator(svc, &params)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *cognitoidentityprovider.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Providers
	}
	return nil
}

func getUserPoolIdentityProvider(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Cognitoidentityprovider
	item := resource.Item.(types.ProviderDescription)
	pool := resource.Parent.Item.(*types.UserPoolType)

	pd, err := svc.DescribeIdentityProvider(ctx, &cognitoidentityprovider.DescribeIdentityProviderInput{
		ProviderName: item.ProviderName,
		UserPoolId:   pool.Id,
	}, func(options *cognitoidentityprovider.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = pd.IdentityProvider
	return nil
}
