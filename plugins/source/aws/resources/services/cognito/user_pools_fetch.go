package cognito

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchCognitoUserPools(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Cognitoidentityprovider
	params := cognitoidentityprovider.ListUserPoolsInput{
		// we want max results to reduce List calls as much as possible, services limited to less than or equal to 60"
		MaxResults: 60,
	}
	for {
		out, err := svc.ListUserPools(ctx, &params)
		if err != nil {
			return err
		}
		res <- out.UserPools

		if aws.ToString(out.NextToken) == "" {
			break
		}
		params.NextToken = out.NextToken
	}
	return nil
}

func getUserPool(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Cognitoidentityprovider
	item := resource.Item.(types.UserPoolDescriptionType)

	upo, err := svc.DescribeUserPool(ctx, &cognitoidentityprovider.DescribeUserPoolInput{UserPoolId: item.Id})
	if err != nil {
		return err
	}

	resource.Item = upo.UserPool
	return nil
}

func fetchCognitoUserPoolIdentityProviders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	pool := parent.Item.(*types.UserPoolType)
	c := meta.(*client.Client)
	svc := c.Services().Cognitoidentityprovider

	params := cognitoidentityprovider.ListIdentityProvidersInput{UserPoolId: pool.Id}
	for {
		out, err := svc.ListIdentityProviders(ctx, &params)
		if err != nil {
			return err
		}
		res <- out.Providers

		if aws.ToString(out.NextToken) == "" {
			break
		}
		params.NextToken = out.NextToken
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
	})
	if err != nil {
		return err
	}

	resource.Item = pd.IdentityProvider
	return nil
}
