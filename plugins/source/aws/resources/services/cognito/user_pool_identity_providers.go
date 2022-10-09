// Code generated by codegen; DO NOT EDIT.

package cognito

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func UserPoolIdentityProviders() *schema.Table {
	return &schema.Table{
		Name:                "aws_cognito_user_pool_identity_providers",
		Resolver:            fetchCognitoUserPoolIdentityProviders,
		PreResourceResolver: getUserPoolIdentityProvider,
		Multiplex:           client.ServiceAccountRegionMultiplexer("cognito-identity"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "user_pool_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "attribute_mapping",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AttributeMapping"),
			},
			{
				Name:     "creation_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationDate"),
			},
			{
				Name:     "idp_identifiers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("IdpIdentifiers"),
			},
			{
				Name:     "last_modified_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastModifiedDate"),
			},
			{
				Name:     "provider_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ProviderDetails"),
			},
			{
				Name:     "provider_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProviderName"),
			},
			{
				Name:     "provider_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProviderType"),
			},
			{
				Name:     "user_pool_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UserPoolId"),
			},
		},
	}
}
