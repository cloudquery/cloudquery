package cognito

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchCognitoUserPools(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().CognitoUserPools
	optsFunc := func(options *cognitoidentityprovider.Options) {
		options.Region = c.Region
	}
	params := cognitoidentityprovider.ListUserPoolsInput{
		// we want max results to reduce List calls as much as possible, services limited to less than or equal to 60"
		MaxResults: 60,
	}
	for {
		out, err := svc.ListUserPools(ctx, &params, optsFunc)
		if err != nil {
			return err
		}
		for _, item := range out.UserPools {
			upo, err := svc.DescribeUserPool(ctx, &cognitoidentityprovider.DescribeUserPoolInput{UserPoolId: item.Id}, optsFunc)
			if err != nil {
				return err
			}
			res <- upo.UserPool
		}
		if aws.ToString(out.NextToken) == "" {
			break
		}
		params.NextToken = out.NextToken
	}
	return nil
}

func fetchCognitoUserPoolIdentityProviders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	pool := parent.Item.(*types.UserPoolType)
	c := meta.(*client.Client)
	svc := c.Services().CognitoUserPools
	optsFunc := func(options *cognitoidentityprovider.Options) {
		options.Region = c.Region
	}
	params := cognitoidentityprovider.ListIdentityProvidersInput{UserPoolId: pool.Id}
	for {
		out, err := svc.ListIdentityProviders(ctx, &params, optsFunc)
		if err != nil {
			return err
		}
		for _, item := range out.Providers {
			pd, err := svc.DescribeIdentityProvider(ctx, &cognitoidentityprovider.DescribeIdentityProviderInput{
				ProviderName: item.ProviderName,
				UserPoolId:   pool.Id,
			}, optsFunc)
			if err != nil {
				return err
			}
			res <- pd.IdentityProvider
		}

		if aws.ToString(out.NextToken) == "" {
			break
		}
		params.NextToken = out.NextToken
	}
	return nil
}
