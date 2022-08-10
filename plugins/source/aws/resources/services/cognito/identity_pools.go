package cognito

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func CognitoIdentityPools() *schema.Table {
	return &schema.Table{
		Name:          "aws_cognito_identity_pools",
		Description:   "An object representing an Amazon Cognito identity pool.",
		Resolver:      fetchCognitoIdentityPools,
		Multiplex:     client.ServiceAccountRegionMultiplexer("cognito-identity"),
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARN(client.CognitoIdentityService, func(resource *schema.Resource) ([]string, error) {
					return []string{"identitypool", *resource.Item.(*cognitoidentity.DescribeIdentityPoolOutput).IdentityPoolId}, nil
				}),
			},
			{
				Name:        "allow_unauthenticated_identities",
				Description: "TRUE if the identity pool supports unauthenticated logins.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "id",
				Description: "An identity pool ID in the format REGION:GUID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("IdentityPoolId"),
			},
			{
				Name:        "identity_pool_name",
				Description: "A string that you provide.",
				Type:        schema.TypeString,
			},
			{
				Name:        "allow_classic_flow",
				Description: "Enables or disables the Basic (Classic) authentication flow",
				Type:        schema.TypeBool,
			},
			{
				Name:        "developer_provider_name",
				Description: "The \"domain\" by which Cognito will refer to your users.",
				Type:        schema.TypeString,
			},
			{
				Name:        "identity_pool_tags",
				Description: "The tags that are assigned to the identity pool",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "open_id_connect_provider_arns",
				Description: "The ARNs of the OpenID Connect providers.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("OpenIdConnectProviderARNs"),
			},
			{
				Name:        "saml_provider_arns",
				Description: "An array of Amazon Resource Names (ARNs) of the SAML provider for your identity pool.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("SamlProviderARNs"),
			},
			{
				Name:        "supported_login_providers",
				Description: "Optional key:value pairs mapping provider names to provider app IDs.",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_cognito_identity_pool_cognito_identity_providers",
				Description:   "A provider representing an Amazon Cognito user pool and its client ID.",
				Resolver:      schema.PathTableResolver("CognitoIdentityProviders"),
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "identity_pool_cq_id",
						Description: "Unique CloudQuery ID of aws_cognito_identity_pools table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "identity_pool_id",
						Description: "An identity pool ID in the format REGION:GUID.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "client_id",
						Description: "The client ID for the Amazon Cognito user pool.",
						Type:        schema.TypeString,
					},
					{
						Name:        "provider_name",
						Description: "The provider name for an Amazon Cognito user pool",
						Type:        schema.TypeString,
					},
					{
						Name:        "server_side_token_check",
						Description: "TRUE if server-side token validation is enabled for the identity provider’s token",
						Type:        schema.TypeBool,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchCognitoIdentityPools(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().CognitoIdentityPools
	optsFunc := func(options *cognitoidentity.Options) { options.Region = c.Region }
	params := cognitoidentity.ListIdentityPoolsInput{
		// we want max results to reduce List calls as much as possible, services limited to less than or equal to 60"
		MaxResults: 60,
	}
	for {
		out, err := svc.ListIdentityPools(ctx, &params, optsFunc)
		if err != nil {
			return diag.WrapError(err)
		}
		for _, item := range out.IdentityPools {
			ipo, err := svc.DescribeIdentityPool(ctx, &cognitoidentity.DescribeIdentityPoolInput{IdentityPoolId: item.IdentityPoolId}, optsFunc)
			if err != nil {
				return diag.WrapError(err)
			}
			res <- ipo
		}
		if aws.ToString(out.NextToken) == "" {
			break
		}
		params.NextToken = out.NextToken
	}
	return nil
}
