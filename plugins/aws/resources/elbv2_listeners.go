package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Elbv2Listeners() *schema.Table {
	return &schema.Table{
		Name:        "aws_elbv2_listeners",
		Description: "Information about a listener.",
		Resolver:    fetchElbv2Listeners,
		IgnoreError: client.IgnoreAccessDeniedServiceDisabled,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Name:        "load_balancer_cq_id",
				Description: "Unique CloudQuery ID of aws_elbv2_load_balancers table (FK)",
				Type:        schema.TypeUUID,
				Resolver:    schema.ParentIdResolver,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveElbv2listenerTags,
			},
			{
				Name:        "alpn_policy",
				Description: "[TLS listener] The name of the Application-Layer Protocol Negotiation (ALPN) policy.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the listener.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ListenerArn"),
			},
			{
				Name:        "load_balancer_arn",
				Description: "The Amazon Resource Name (ARN) of the load balancer.",
				Type:        schema.TypeString,
			},
			{
				Name:        "port",
				Description: "The port on which the load balancer is listening.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "protocol",
				Description: "The protocol for connections from clients to the load balancer.",
				Type:        schema.TypeString,
			},
			{
				Name:        "ssl_policy",
				Description: "[HTTPS or TLS listener] The security policy that defines which protocols and ciphers are supported.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_elbv2_listener_certificates",
				Description: "Information about an SSL server certificate.",
				Resolver:    fetchElbv2ListenerCertificates,
				Columns: []schema.Column{
					{
						Name:        "listener_cq_id",
						Description: "Unique CloudQuery ID of aws_elbv2_listeners table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "certificate_arn",
						Description: "The Amazon Resource Name (ARN) of the certificate.",
						Type:        schema.TypeString,
					},
					{
						Name:        "is_default",
						Description: "Indicates whether the certificate is the default certificate",
						Type:        schema.TypeBool,
					},
				},
			},
			{
				Name:        "aws_elbv2_listener_default_actions",
				Description: "Information about an action",
				Resolver:    fetchElbv2ListenerDefaultActions,
				Columns: []schema.Column{
					{
						Name:        "listener_cq_id",
						Description: "Unique CloudQuery ID of aws_elbv2_listeners table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "type",
						Description: "The type of action.  This member is required.",
						Type:        schema.TypeString,
					},
					{
						Name:        "auth_cognito_user_pool_arn",
						Description: "The Amazon Resource Name (ARN) of the Amazon Cognito user pool.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AuthenticateCognitoConfig.UserPoolArn"),
					},
					{
						Name:        "auth_cognito_user_pool_client_id",
						Description: "The ID of the Amazon Cognito user pool client.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AuthenticateCognitoConfig.UserPoolClientId"),
					},
					{
						Name:        "auth_cognito_user_pool_domain",
						Description: "The domain prefix or fully-qualified domain name of the Amazon Cognito user pool.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AuthenticateCognitoConfig.UserPoolDomain"),
					},
					{
						Name:        "auth_cognito_authentication_request_extra_params",
						Description: "The query parameters (up to 10) to include in the redirect request to the authorization endpoint.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("AuthenticateCognitoConfig.AuthenticationRequestExtraParams"),
					},
					{
						Name:        "auth_cognito_on_unauthenticated_request",
						Description: "The behavior if the user is not authenticated",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AuthenticateCognitoConfig.OnUnauthenticatedRequest"),
					},
					{
						Name:        "auth_cognito_scope",
						Description: "The set of user claims to be requested from the IdP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AuthenticateCognitoConfig.Scope"),
					},
					{
						Name:        "auth_cognito_session_cookie_name",
						Description: "The name of the cookie used to maintain session information",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AuthenticateCognitoConfig.SessionCookieName"),
					},
					{
						Name:        "auth_cognito_session_timeout",
						Description: "The maximum duration of the authentication session, in seconds",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("AuthenticateCognitoConfig.SessionTimeout"),
					},
					{
						Name:        "auth_oidc_authorization_endpoint",
						Description: "The authorization endpoint of the IdP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AuthenticateOidcConfig.AuthorizationEndpoint"),
					},
					{
						Name:        "auth_oidc_client_id",
						Description: "The OAuth 2.0 client identifier.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AuthenticateOidcConfig.ClientId"),
					},
					{
						Name:        "auth_oidc_issuer",
						Description: "The OIDC issuer identifier of the IdP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AuthenticateOidcConfig.Issuer"),
					},
					{
						Name:        "auth_oidc_token_endpoint",
						Description: "The token endpoint of the IdP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AuthenticateOidcConfig.TokenEndpoint"),
					},
					{
						Name:        "auth_oidc_user_info_endpoint",
						Description: "The user info endpoint of the IdP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AuthenticateOidcConfig.UserInfoEndpoint"),
					},
					{
						Name:        "auth_oidc_authentication_request_extra_params",
						Description: "The query parameters (up to 10) to include in the redirect request to the authorization endpoint.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("AuthenticateOidcConfig.AuthenticationRequestExtraParams"),
					},
					{
						Name:        "auth_oidc_client_secret",
						Description: "The OAuth 2.0 client secret",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AuthenticateOidcConfig.ClientSecret"),
					},
					{
						Name:        "auth_oidc_on_unauthenticated_request",
						Description: "The behavior if the user is not authenticated",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AuthenticateOidcConfig.OnUnauthenticatedRequest"),
					},
					{
						Name:        "auth_oidc_scope",
						Description: "The set of user claims to be requested from the IdP",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AuthenticateOidcConfig.Scope"),
					},
					{
						Name:        "auth_oidc_session_cookie_name",
						Description: "The name of the cookie used to maintain session information",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AuthenticateOidcConfig.SessionCookieName"),
					},
					{
						Name:        "auth_oidc_session_timeout",
						Description: "The maximum duration of the authentication session, in seconds",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("AuthenticateOidcConfig.SessionTimeout"),
					},
					{
						Name:        "auth_oidc_use_existing_client_secret",
						Description: "Indicates whether to use the existing client secret when modifying a rule",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("AuthenticateOidcConfig.UseExistingClientSecret"),
					},
					{
						Name:        "fixed_response_config_status_code",
						Description: "The HTTP response code (2XX, 4XX, or 5XX).  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FixedResponseConfig.StatusCode"),
					},
					{
						Name:        "fixed_response_config_content_type",
						Description: "The content type",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FixedResponseConfig.ContentType"),
					},
					{
						Name:        "fixed_response_config_message_body",
						Description: "The message.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FixedResponseConfig.MessageBody"),
					},
					{
						Name:        "forward_config_target_group_stickiness_config_duration_seconds",
						Description: "The time period, in seconds, during which requests from a client should be routed to the same target group",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("ForwardConfig.TargetGroupStickinessConfig.DurationSeconds"),
					},
					{
						Name:        "forward_config_target_group_stickiness_config_enabled",
						Description: "Indicates whether target group stickiness is enabled.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("ForwardConfig.TargetGroupStickinessConfig.Enabled"),
					},
					{
						Name:        "order",
						Description: "The order for the action",
						Type:        schema.TypeInt,
					},
					{
						Name:        "redirect_config_status_code",
						Description: "The HTTP redirect code",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RedirectConfig.StatusCode"),
					},
					{
						Name:        "redirect_config_host",
						Description: "The hostname",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RedirectConfig.Host"),
					},
					{
						Name:        "redirect_config_path",
						Description: "The absolute path, starting with the leading \"/\"",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RedirectConfig.Path"),
					},
					{
						Name:        "redirect_config_port",
						Description: "The port",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RedirectConfig.Port"),
					},
					{
						Name:        "redirect_config_protocol",
						Description: "The protocol",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RedirectConfig.Protocol"),
					},
					{
						Name:        "redirect_config_query",
						Description: "The query parameters, URL-encoded when necessary, but not percent-encoded",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RedirectConfig.Query"),
					},
					{
						Name:        "target_group_arn",
						Description: "The Amazon Resource Name (ARN) of the target group",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_elbv2_listener_default_action_forward_config_target_groups",
						Description: "Information about how traffic will be distributed between multiple target groups in a forward rule.",
						Resolver:    fetchElbv2ListenerDefaultActionForwardConfigTargetGroups,
						Columns: []schema.Column{
							{
								Name:        "listener_default_action_cq_id",
								Description: "Unique CloudQuery ID of aws_elbv2_listener_default_actions table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "target_group_arn",
								Description: "The Amazon Resource Name (ARN) of the target group.",
								Type:        schema.TypeString,
							},
							{
								Name:        "weight",
								Description: "The weight",
								Type:        schema.TypeInt,
							},
						},
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchElbv2Listeners(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	lb, ok := parent.Item.(types.LoadBalancer)
	if !ok {
		return fmt.Errorf("expected to have types.LoadBalancer but got %T", parent.Item)
	}
	config := elbv2.DescribeListenersInput{
		LoadBalancerArn: lb.LoadBalancerArn,
	}
	c := meta.(*client.Client)
	svc := c.Services().ELBv2
	for {
		response, err := svc.DescribeListeners(ctx, &config, func(options *elbv2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.Listeners
		if aws.ToString(response.NextMarker) == "" {
			break
		}
		config.Marker = response.NextMarker
	}
	return nil
}
func resolveElbv2listenerTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().ELBv2
	listener, ok := resource.Item.(types.Listener)
	if !ok {
		return fmt.Errorf("expected to have types.Listener but got %T", resource.Item)
	}
	tagsOutput, err := svc.DescribeTags(ctx, &elbv2.DescribeTagsInput{
		ResourceArns: []string{
			*listener.ListenerArn,
		},
	}, func(o *elbv2.Options) {
		o.Region = region
	})
	if err != nil {
		return err
	}
	if len(tagsOutput.TagDescriptions) == 0 {
		return nil
	}
	tags := make(map[string]*string)
	for _, td := range tagsOutput.TagDescriptions {
		for _, s := range td.Tags {
			tags[*s.Key] = s.Value
		}
	}

	return resource.Set(c.Name, tags)
}
func fetchElbv2ListenerCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().ELBv2
	listener, ok := parent.Item.(types.Listener)
	if !ok {
		return fmt.Errorf("expected to have types.Listener but got %T", parent.Item)
	}
	config := elbv2.DescribeListenerCertificatesInput{ListenerArn: listener.ListenerArn}
	for {
		response, err := svc.DescribeListenerCertificates(ctx, &config, func(options *elbv2.Options) {
			options.Region = region
		})
		if err != nil {
			return err
		}
		res <- response.Certificates
		if aws.ToString(response.NextMarker) == "" {
			break
		}
		config.Marker = response.NextMarker
	}
	return nil
}

func fetchElbv2ListenerDefaultActions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	listener, ok := parent.Item.(types.Listener)
	if !ok {
		return fmt.Errorf("expected to have types.Listener but got %T", parent.Item)
	}
	res <- listener.DefaultActions
	return nil
}
func fetchElbv2ListenerDefaultActionForwardConfigTargetGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	action, ok := parent.Item.(types.Action)
	if !ok {
		return fmt.Errorf("expected to have types.Action but got %T", parent.Item)
	}
	res <- action.ForwardConfig.TargetGroups
	return nil
}
