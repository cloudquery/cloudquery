package cognito

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func IdentityPools() *schema.Table {
	tableName := "aws_cognito_identity_pools"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_DescribeIdentityPool.html`,
		Resolver:            fetchCognitoIdentityPools,
		PreResourceResolver: getIdentityPool,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "cognito-identity"),
		Transform: transformers.TransformWithStruct(
			&cognitoidentity.DescribeIdentityPoolOutput{},
			transformers.WithNameTransformer(client.CreateReplaceTransformer(map[string]string{"ar_ns": "arns"})),
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: resolveIdentityPoolARN(),
			},
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("IdentityPoolId"),
				PrimaryKey: true,
			},
			{
				Name:     "saml_provider_arns",
				Type:     arrow.ListOf(arrow.BinaryTypes.String),
				Resolver: schema.PathResolver("SamlProviderARNs"),
			},
		},
	}
}

func fetchCognitoIdentityPools(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Cognitoidentity
	params := cognitoidentity.ListIdentityPoolsInput{
		// we want max results to reduce List calls as much as possible, services limited to less than or equal to 60"
		MaxResults: 60,
	}
	paginator := cognitoidentity.NewListIdentityPoolsPaginator(svc, &params)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *cognitoidentity.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.IdentityPools
	}
	return nil
}

func getIdentityPool(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Cognitoidentity
	item := resource.Item.(types.IdentityPoolShortDescription)

	ipo, err := svc.DescribeIdentityPool(ctx, &cognitoidentity.DescribeIdentityPoolInput{IdentityPoolId: item.IdentityPoolId}, func(options *cognitoidentity.Options) {
		options.Region = cl.Region
	})
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
