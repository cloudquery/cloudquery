package cognito

import (
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func IdentityPools() *schema.Table {
	return &schema.Table{
		Name:                "aws_cognito_identity_pools",
		Resolver:            fetchCognitoIdentityPools,
		PreResourceResolver: getIdentityPool,
		Multiplex:           client.ServiceAccountRegionMultiplexer("cognito-identity"),
		Transform:           transformers.TransformWithStruct(&cognitoidentity.DescribeIdentityPoolOutput{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
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
