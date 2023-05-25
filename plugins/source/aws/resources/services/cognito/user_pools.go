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

func UserPools() *schema.Table {
	tableName := "aws_cognito_user_pools"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UserPoolType.html`,
		Resolver:            fetchCognitoUserPools,
		PreResourceResolver: getUserPool,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "cognito-identity"),
		Transform:           transformers.TransformWithStruct(&types.UserPoolType{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Id"),
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			userPoolIdentityProviders(),
		},
	}
}

func fetchCognitoUserPools(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Cognitoidentityprovider
	params := cognitoidentityprovider.ListUserPoolsInput{
		// we want max results to reduce List calls as much as possible, services limited to less than or equal to 60"
		MaxResults: 60,
	}
	paginator := cognitoidentityprovider.NewListUserPoolsPaginator(svc, &params)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *cognitoidentityprovider.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- page.UserPools
	}
	return nil
}

func getUserPool(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Cognitoidentityprovider
	item := resource.Item.(types.UserPoolDescriptionType)

	upo, err := svc.DescribeUserPool(ctx, &cognitoidentityprovider.DescribeUserPoolInput{UserPoolId: item.Id}, func(options *cognitoidentityprovider.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}

	resource.Item = upo.UserPool
	return nil
}
