package cognito

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchCognitoIdentityPools(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Cognitoidentity
	params := cognitoidentity.ListIdentityPoolsInput{
		// we want max results to reduce List calls as much as possible, services limited to less than or equal to 60"
		MaxResults: 60,
	}
	for {
		out, err := svc.ListIdentityPools(ctx, &params)
		if err != nil {
			return err
		}
		res <- out.IdentityPools

		if aws.ToString(out.NextToken) == "" {
			break
		}
		params.NextToken = out.NextToken
	}
	return nil
}

func getIdentityPool(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Cognitoidentity
	item := resource.Item.(types.IdentityPoolShortDescription)

	ipo, err := svc.DescribeIdentityPool(ctx, &cognitoidentity.DescribeIdentityPoolInput{IdentityPoolId: item.IdentityPoolId})
	if err != nil {
		return err
	}

	resource.Item = ipo
	return nil
}

func resolveIdentityPoolARN() schema.ColumnResolver {
	return client.ResolveARN(client.CognitoIdentityService, func(resource *schema.Resource) ([]string, error) {
		return []string{"identitypool", *resource.Item.(*cognitoidentity.DescribeIdentityPoolOutput).IdentityPoolId}, nil
	})
}
