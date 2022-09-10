package cognito

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func CognitoUserPools() *schema.Table {
	return &schema.Table{
		Name:          "aws_cognito_user_pools",
		Description:   "A container for information about the user pool.",
		Resolver:      fetchCognitoUserPools,
		Multiplex:     client.ServiceAccountRegionMultiplexer("cognito-idp"),
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:            "account_id",
				Description:     "The AWS Account ID of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:          "account_recovery_setting",
				Description:   "Use this setting to define which verified available method a user can use to recover their password when they call ForgotPassword",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("AccountRecoverySetting"),
				IgnoreInTests: true,
			},
			{
				Name:     "admin_create_user_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AdminCreateUserConfig"),
			},
			{
				Name:          "alias_attributes",
				Description:   "Specifies the attributes that are aliased in a user pool.",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the user pool.",
				Type:        schema.TypeString,
			},
			{
				Name:          "auto_verified_attributes",
				Description:   "Specifies the attributes that are auto-verified in a user pool.",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "creation_date",
				Description: "The date the user pool was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:          "custom_domain",
				Description:   "A custom domain name that you provide to Amazon Cognito",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:     "device_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DeviceConfiguration"),
			},
			{
				Name:          "domain",
				Description:   "Holds the domain prefix if the user pool has a domain associated with it.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "email_configuration",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("EmailConfiguration"),
				IgnoreInTests: true,
			},
			{
				Name:          "email_configuration_failure",
				Description:   "The reason why the email configuration cannot send the messages to your users.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "email_verification_message",
				Description:   "The contents of the email verification message.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "email_verification_subject",
				Description:   "The subject of the email verification message.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "estimated_number_of_users",
				Description: "A number estimating the size of the user pool.",
				Type:        schema.TypeInt,
			},
			{
				Name:            "id",
				Description:     "The ID of the user pool.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:          "lambda_config",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("LambdaConfig"),
				IgnoreInTests: true,
			},
			{
				Name:        "last_modified_date",
				Description: "The date the user pool was last modified.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "mfa_configuration",
				Description: "Can be one of the following values:  * OFF - MFA tokens are not required and cannot be specified during user registration.  * ON - MFA tokens are required for all user registrations",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The name of the user pool.",
				Type:        schema.TypeString,
			},
			{
				Name:     "policies",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Policies"),
			},
			{
				Name:          "sms_authentication_message",
				Description:   "The contents of the SMS authentication message.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "sms_configuration",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("SmsConfiguration"),
				IgnoreInTests: true,
			},
			{
				Name:          "sms_configuration_failure",
				Description:   "The reason why the SMS configuration cannot send the messages to your users. This message might include comma-separated values to describe why your SMS configuration can't send messages to user pool end users.  * InvalidSmsRoleAccessPolicyException - The IAM role which Cognito uses to send SMS messages is not properly configured",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "sms_verification_message",
				Description:   "The contents of the SMS verification message.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "status",
				Description: "The status of a user pool.",
				Type:        schema.TypeString,
			},
			{
				Name:     "user_pool_add_ons",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("UserPoolAddOns"),
			},
			{
				Name:        "user_pool_tags",
				Description: "The tags that are assigned to the user pool",
				Type:        schema.TypeJSON,
			},
			{
				Name:          "username_attributes",
				Description:   "Specifies whether email addresses or phone numbers can be specified as usernames when a user signs up.",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:          "username_configuration",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("UsernameConfiguration"),
				IgnoreInTests: true,
			},
			{
				Name:     "verification_message_template",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VerificationMessageTemplate"),
			},
			{
				Name:          "schema_attributes",
				Description:   "Contains information about the schema attribute.",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("SchemaAttributes"),
				IgnoreInTests: true,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_cognito_user_pool_identity_providers",
				Description:   "A container for information about an identity provider.",
				Resolver:      fetchCognitoUserPoolIdentityProviders,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "user_pool_cq_id",
						Description: "Unique CloudQuery ID of aws_cognito_user_pools table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "user_pool_id",
						Description: "The ID of the user pool.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
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
						Name:        "attribute_mapping",
						Description: "A mapping of identity provider attributes to standard and custom user pool attributes.",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "creation_date",
						Description: "The date the identity provider was created.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "idp_identifiers",
						Description: "A list of identity provider identifiers.",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "last_modified_date",
						Description: "The date the identity provider was last modified.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "provider_details",
						Description: "The identity provider details",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "provider_name",
						Description: "The identity provider name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "provider_type",
						Description: "The identity provider type.",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
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
