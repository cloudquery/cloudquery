package cognito

import (
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func IdentityPools() *schema.Table {
	tableName := "aws_cognito_identity_pools"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_DescribeIdentityPool.html`,
		Resolver:            fetchCognitoIdentityPools,
		PreResourceResolver: getIdentityPool,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "cognito-identity"),
		Transform:           client.TransformWithStruct(&cognitoidentity.DescribeIdentityPoolOutput{}),
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
