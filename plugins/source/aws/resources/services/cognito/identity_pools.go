package cognito

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func IdentityPools() *schema.Table {
	tableName := "aws_cognito_identity_pools"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_DescribeIdentityPool.html`,
		Resolver:            fetchCognitoIdentityPools,
		PreResourceResolver: getIdentityPool,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "cognito-identity"),
		Transform:           transformers.TransformWithStruct(&cognitoidentity.DescribeIdentityPoolOutput{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveIdentityPoolARN(),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IdentityPoolId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "saml_provider_ar_ns",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SamlProviderARNs"),
			},
		},
	}
}

func fetchCognitoIdentityPools(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Cognitoidentity
	params := cognitoidentity.ListIdentityPoolsInput{
		// we want max results to reduce List calls as much as possible, services limited to less than or equal to 60"
		MaxResults: 60,
	}
	paginator := cognitoidentity.NewListIdentityPoolsPaginator(svc, &params)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.IdentityPools
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
