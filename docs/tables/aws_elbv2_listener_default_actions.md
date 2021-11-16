
# Table: aws_elbv2_listener_default_actions
Information about an action
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|listener_cq_id|uuid|Unique CloudQuery ID of aws_elbv2_listeners table (FK)|
|type|text|The type of action.  This member is required.|
|auth_cognito_user_pool_arn|text|The Amazon Resource Name (ARN) of the Amazon Cognito user pool.  This member is required.|
|auth_cognito_user_pool_client_id|text|The ID of the Amazon Cognito user pool client.  This member is required.|
|auth_cognito_user_pool_domain|text|The domain prefix or fully-qualified domain name of the Amazon Cognito user pool.  This member is required.|
|auth_cognito_authentication_request_extra_params|jsonb|The query parameters (up to 10) to include in the redirect request to the authorization endpoint.|
|auth_cognito_on_unauthenticated_request|text|The behavior if the user is not authenticated|
|auth_cognito_scope|text|The set of user claims to be requested from the IdP|
|auth_cognito_session_cookie_name|text|The name of the cookie used to maintain session information|
|auth_cognito_session_timeout|bigint|The maximum duration of the authentication session, in seconds|
|auth_oidc_authorization_endpoint|text|The authorization endpoint of the IdP|
|auth_oidc_client_id|text|The OAuth 2.0 client identifier.  This member is required.|
|auth_oidc_issuer|text|The OIDC issuer identifier of the IdP|
|auth_oidc_token_endpoint|text|The token endpoint of the IdP|
|auth_oidc_user_info_endpoint|text|The user info endpoint of the IdP|
|auth_oidc_authentication_request_extra_params|jsonb|The query parameters (up to 10) to include in the redirect request to the authorization endpoint.|
|auth_oidc_client_secret|text|The OAuth 2.0 client secret|
|auth_oidc_on_unauthenticated_request|text|The behavior if the user is not authenticated|
|auth_oidc_scope|text|The set of user claims to be requested from the IdP|
|auth_oidc_session_cookie_name|text|The name of the cookie used to maintain session information|
|auth_oidc_session_timeout|bigint|The maximum duration of the authentication session, in seconds|
|auth_oidc_use_existing_client_secret|boolean|Indicates whether to use the existing client secret when modifying a rule|
|fixed_response_config_status_code|text|The HTTP response code (2XX, 4XX, or 5XX).  This member is required.|
|fixed_response_config_content_type|text|The content type|
|fixed_response_config_message_body|text|The message.|
|forward_config_target_group_stickiness_config_duration_seconds|integer|The time period, in seconds, during which requests from a client should be routed to the same target group|
|forward_config_target_group_stickiness_config_enabled|boolean|Indicates whether target group stickiness is enabled.|
|order|integer|The order for the action|
|redirect_config_status_code|text|The HTTP redirect code|
|redirect_config_host|text|The hostname|
|redirect_config_path|text|The absolute path, starting with the leading "/"|
|redirect_config_port|text|The port|
|redirect_config_protocol|text|The protocol|
|redirect_config_query|text|The query parameters, URL-encoded when necessary, but not percent-encoded|
|target_group_arn|text|The Amazon Resource Name (ARN) of the target group|
