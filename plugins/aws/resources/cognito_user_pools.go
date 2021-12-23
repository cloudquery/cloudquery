package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func CognitoUserPools() *schema.Table {
	return &schema.Table{
		Name:         "aws_cognito_user_pools",
		Description:  "A container for information about the user pool.",
		Resolver:     fetchCognitoUserPools,
		Multiplex:    client.ServiceAccountRegionMultiplexer("cognito-idp"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
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
				Name:        "account_recovery_setting",
				Description: "Use this setting to define which verified available method a user can use to recover their password when they call ForgotPassword",
				Type:        schema.TypeJSON,
				Resolver:    resolveCognitoUserPoolAccountRecoverySetting,
			},
			{
				Name:        "admin_create_user_admin_only",
				Description: "Set to True if only the administrator is allowed to create user profiles",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("AdminCreateUserConfig.AllowAdminCreateUserOnly"),
			},
			{
				Name:        "admin_create_user_invite_email_message",
				Description: "The message template for email messages",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AdminCreateUserConfig.InviteMessageTemplate.EmailMessage"),
			},
			{
				Name:        "admin_create_user_invite_email_subject",
				Description: "The subject line for email messages",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AdminCreateUserConfig.InviteMessageTemplate.EmailSubject"),
			},
			{
				Name:        "admin_create_user_invite_sms",
				Description: "The message template for SMS messages.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AdminCreateUserConfig.InviteMessageTemplate.SMSMessage"),
			},
			{
				Name:        "admin_create_user_config_unused_account_validity_days",
				Description: "The user account expiration limit, in days, after which the account is no longer usable",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("AdminCreateUserConfig.UnusedAccountValidityDays"),
			},
			{
				Name:        "alias_attributes",
				Description: "Specifies the attributes that are aliased in a user pool.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the user pool.",
				Type:        schema.TypeString,
			},
			{
				Name:        "auto_verified_attributes",
				Description: "Specifies the attributes that are auto-verified in a user pool.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "creation_date",
				Description: "The date the user pool was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "custom_domain",
				Description: "A custom domain name that you provide to Amazon Cognito",
				Type:        schema.TypeString,
			},
			{
				Name:        "challenge_required_on_new_device",
				Description: "Indicates whether a challenge is required on a new device",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DeviceConfiguration.ChallengeRequiredOnNewDevice"),
			},
			{
				Name:        "device_only_remembered_on_user_prompt",
				Description: "If true, a device is only remembered on user prompt.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DeviceConfiguration.DeviceOnlyRememberedOnUserPrompt"),
			},
			{
				Name:        "domain",
				Description: "Holds the domain prefix if the user pool has a domain associated with it.",
				Type:        schema.TypeString,
			},
			{
				Name:        "email_configuration_set",
				Description: "The set of configuration rules that can be applied to emails sent using Amazon SES",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EmailConfiguration.ConfigurationSet"),
			},
			{
				Name:        "email_configuration_sending_account",
				Description: "Specifies whether Amazon Cognito emails your users by using its built-in email functionality or your Amazon SES email configuration",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EmailConfiguration.EmailSendingAccount"),
			},
			{
				Name:        "email_configuration_from",
				Description: "Identifies either the sender’s email address or the sender’s name with their email address",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EmailConfiguration.From"),
			},
			{
				Name:        "email_configuration_reply_to_address",
				Description: "The destination to which the receiver of the email should reply to.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EmailConfiguration.ReplyToEmailAddress"),
			},
			{
				Name:        "email_configuration_source_arn",
				Description: "The Amazon Resource Name (ARN) of a verified email address in Amazon SES",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EmailConfiguration.SourceArn"),
			},
			{
				Name:        "email_configuration_failure",
				Description: "The reason why the email configuration cannot send the messages to your users.",
				Type:        schema.TypeString,
			},
			{
				Name:        "email_verification_message",
				Description: "The contents of the email verification message.",
				Type:        schema.TypeString,
			},
			{
				Name:        "email_verification_subject",
				Description: "The subject of the email verification message.",
				Type:        schema.TypeString,
			},
			{
				Name:        "estimated_number_of_users",
				Description: "A number estimating the size of the user pool.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "id",
				Description: "The ID of the user pool.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Id"),
			},
			{
				Name:        "lambda_config_create_auth_challenge",
				Description: "Creates an authentication challenge.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LambdaConfig.CreateAuthChallenge"),
			},
			{
				Name:        "lambda_config_custom_email_sender_lambda_arn",
				Description: "The Lambda Amazon Resource Name of the Lambda function that Amazon Cognito triggers to send email notifications to users.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LambdaConfig.CustomEmailSender.LambdaArn"),
			},
			{
				Name:        "lambda_config_custom_email_sender_lambda_version",
				Description: "The Lambda version represents the signature of the \"request\" attribute in the \"event\" information Amazon Cognito passes to your custom email Lambda function. The only supported value is V1_0.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LambdaConfig.CustomEmailSender.LambdaVersion"),
			},
			{
				Name:        "lambda_config_custom_message",
				Description: "A custom Message AWS Lambda trigger.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LambdaConfig.CustomMessage"),
			},
			{
				Name:        "lambda_config_custom_sms_sender_lambda_arn",
				Description: "The Lambda Amazon Resource Name of the Lambda function that Amazon Cognito triggers to send SMS notifications to users.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LambdaConfig.CustomSMSSender.LambdaArn"),
			},
			{
				Name:        "lambda_config_custom_sms_sender_lambda_version",
				Description: "The Lambda version represents the signature of the \"request\" attribute in the \"event\" information Amazon Cognito passes to your custom SMS Lambda function. The only supported value is V1_0.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LambdaConfig.CustomSMSSender.LambdaVersion"),
			},
			{
				Name:        "lambda_config_define_auth_challenge",
				Description: "Defines the authentication challenge.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LambdaConfig.DefineAuthChallenge"),
			},
			{
				Name:        "lambda_config_kms_key_id",
				Description: "The Amazon Resource Name of Key Management Service Customer master keys",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LambdaConfig.KMSKeyID"),
			},
			{
				Name:        "lambda_config_post_authentication",
				Description: "A post-authentication AWS Lambda trigger.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LambdaConfig.PostAuthentication"),
			},
			{
				Name:        "lambda_config_post_confirmation",
				Description: "A post-confirmation AWS Lambda trigger.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LambdaConfig.PostConfirmation"),
			},
			{
				Name:        "lambda_config_pre_authentication",
				Description: "A pre-authentication AWS Lambda trigger.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LambdaConfig.PreAuthentication"),
			},
			{
				Name:        "lambda_config_pre_sign_up",
				Description: "A pre-registration AWS Lambda trigger.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LambdaConfig.PreSignUp"),
			},
			{
				Name:        "lambda_config_pre_token_generation",
				Description: "A Lambda trigger that is invoked before token generation.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LambdaConfig.PreTokenGeneration"),
			},
			{
				Name:        "lambda_config_user_migration",
				Description: "The user migration Lambda config type.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LambdaConfig.UserMigration"),
			},
			{
				Name:        "lambda_config_verify_auth_challenge_response",
				Description: "Verifies the authentication challenge response.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LambdaConfig.VerifyAuthChallengeResponse"),
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
				Name:        "policies_password_policy_minimum_length",
				Description: "The minimum length of the password policy that you have set",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Policies.PasswordPolicy.MinimumLength"),
			},
			{
				Name:        "policies_password_policy_require_lowercase",
				Description: "In the password policy that you have set, refers to whether you have required users to use at least one lowercase letter in their password.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Policies.PasswordPolicy.RequireLowercase"),
			},
			{
				Name:        "policies_password_policy_require_numbers",
				Description: "In the password policy that you have set, refers to whether you have required users to use at least one number in their password.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Policies.PasswordPolicy.RequireNumbers"),
			},
			{
				Name:        "policies_password_policy_require_symbols",
				Description: "In the password policy that you have set, refers to whether you have required users to use at least one symbol in their password.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Policies.PasswordPolicy.RequireSymbols"),
			},
			{
				Name:        "policies_password_policy_require_uppercase",
				Description: "In the password policy that you have set, refers to whether you have required users to use at least one uppercase letter in their password.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Policies.PasswordPolicy.RequireUppercase"),
			},
			{
				Name:        "policies_password_policy_temporary_password_validity_days",
				Description: "In the password policy you have set, refers to the number of days a temporary password is valid",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Policies.PasswordPolicy.TemporaryPasswordValidityDays"),
			},
			{
				Name:        "sms_authentication_message",
				Description: "The contents of the SMS authentication message.",
				Type:        schema.TypeString,
			},
			{
				Name:        "sms_configuration_sns_caller_arn",
				Description: "The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) caller",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SmsConfiguration.SnsCallerArn"),
			},
			{
				Name:        "sms_configuration_external_id",
				Description: "The external ID is a value that we recommend you use to add security to your IAM role which is used to call Amazon SNS to send SMS messages for your user pool. If you provide an ExternalId, the Cognito User Pool will include it when attempting to assume your IAM role, so that you can set your roles trust policy to require the ExternalID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SmsConfiguration.ExternalId"),
			},
			{
				Name:        "sms_configuration_failure",
				Description: "The reason why the SMS configuration cannot send the messages to your users. This message might include comma-separated values to describe why your SMS configuration can't send messages to user pool end users.  * InvalidSmsRoleAccessPolicyException - The IAM role which Cognito uses to send SMS messages is not properly configured",
				Type:        schema.TypeString,
			},
			{
				Name:        "sms_verification_message",
				Description: "The contents of the SMS verification message.",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The status of a user pool.",
				Type:        schema.TypeString,
			},
			{
				Name:        "user_pool_add_ons_advanced_security_mode",
				Description: "The advanced security mode.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("UserPoolAddOns.AdvancedSecurityMode"),
			},
			{
				Name:        "user_pool_tags",
				Description: "The tags that are assigned to the user pool",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "username_attributes",
				Description: "Specifies whether email addresses or phone numbers can be specified as usernames when a user signs up.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "username_configuration_case_sensitive",
				Description: "Specifies whether username case sensitivity will be applied for all users in the user pool through Cognito APIs",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("UsernameConfiguration.CaseSensitive"),
			},
			{
				Name:        "verification_message_template_default_email_option",
				Description: "The default email option.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VerificationMessageTemplate.DefaultEmailOption"),
			},
			{
				Name:        "verification_message_template_email_message",
				Description: "The email message template",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VerificationMessageTemplate.EmailMessage"),
			},
			{
				Name:        "verification_message_template_email_message_by_link",
				Description: "The email message template for sending a confirmation link to the user. EmailMessageByLink is allowed only if  EmailSendingAccount (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is DEVELOPER.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VerificationMessageTemplate.EmailMessageByLink"),
			},
			{
				Name:        "verification_message_template_email_subject",
				Description: "The subject line for the email message template",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VerificationMessageTemplate.EmailSubject"),
			},
			{
				Name:        "verification_message_template_email_subject_by_link",
				Description: "The subject line for the email message template for sending a confirmation link to the user",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VerificationMessageTemplate.EmailSubjectByLink"),
			},
			{
				Name:        "verification_message_template_sms_message",
				Description: "The SMS message template.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VerificationMessageTemplate.SmsMessage"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_cognito_user_pool_schema_attributes",
				Description: "Contains information about the schema attribute.",
				Resolver:    fetchCognitoUserPoolSchemaAttributes,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"user_pool_cq_id", "name"}},
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
						Name:        "attribute_data_type",
						Description: "The attribute data type.",
						Type:        schema.TypeString,
					},
					{
						Name:        "developer_only_attribute",
						Description: "We recommend that you use WriteAttributes (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UserPoolClientType.html#CognitoUserPools-Type-UserPoolClientType-WriteAttributes) in the user pool client to control how attributes can be mutated for new use cases instead of using DeveloperOnlyAttribute",
						Type:        schema.TypeBool,
					},
					{
						Name:        "mutable",
						Description: "Specifies whether the value of the attribute can be changed",
						Type:        schema.TypeBool,
					},
					{
						Name:        "name",
						Description: "A schema attribute of the name type.",
						Type:        schema.TypeString,
					},
					{
						Name:        "number_attribute_constraints_max_value",
						Description: "The maximum value of an attribute that is of the number data type.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("NumberAttributeConstraints.MaxValue"),
					},
					{
						Name:        "number_attribute_constraints_min_value",
						Description: "The minimum value of an attribute that is of the number data type.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("NumberAttributeConstraints.MinValue"),
					},
					{
						Name:        "required",
						Description: "Specifies whether a user pool attribute is required",
						Type:        schema.TypeBool,
					},
					{
						Name:        "string_attribute_constraints_max_length",
						Description: "The maximum length.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("StringAttributeConstraints.MaxLength"),
					},
					{
						Name:        "string_attribute_constraints_min_length",
						Description: "The minimum length.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("StringAttributeConstraints.MinLength"),
					},
				},
			},
			{
				Name:        "aws_cognito_user_pool_identity_providers",
				Description: "A container for information about an identity provider.",
				Resolver:    fetchCognitoUserPoolIdentityProviders,
				IgnoreError: client.IgnoreAccessDeniedServiceDisabled,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"user_pool_cq_id", "provider_name"}},
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
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchCognitoUserPools(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().CognitoUserPools
	optsFunc := func(options *cognitoidentityprovider.Options) { options.Region = c.Region }
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

func resolveCognitoUserPoolAccountRecoverySetting(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	pool, ok := resource.Item.(*types.UserPoolType)
	if !ok {
		return fmt.Errorf("not a UserPoolType instance: %#v", resource.Item)
	}
	data, err := json.Marshal(pool.AccountRecoverySetting)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}

func fetchCognitoUserPoolSchemaAttributes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	pool, ok := parent.Item.(*types.UserPoolType)
	if !ok {
		return fmt.Errorf("not a UserPoolType instance: %#v", parent.Item)
	}
	res <- pool.SchemaAttributes
	return nil
}

func fetchCognitoUserPoolIdentityProviders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	pool, ok := parent.Item.(*types.UserPoolType)
	if !ok {
		return fmt.Errorf("not a UserPoolType instance: %#v", parent.Item)
	}
	c := meta.(*client.Client)
	svc := c.Services().CognitoUserPools
	optsFunc := func(options *cognitoidentityprovider.Options) { options.Region = c.Region }
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
