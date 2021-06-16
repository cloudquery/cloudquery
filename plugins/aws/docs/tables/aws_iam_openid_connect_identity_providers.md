
# Table: aws_iam_openid_connect_identity_providers
IAM OIDC identity providers are entities in IAM that describe an external identity provider (IdP) service that supports the OpenID Connect (OIDC) standard, such as Google or Salesforce.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|client_id_list|text[]|A list of client IDs (also known as audiences) that are associated with the specified IAM OIDC provider resource object. For more information, see CreateOpenIDConnectProvider. |
|create_date|timestamp without time zone|The date and time when the IAM OIDC provider resource object was created in the AWS account. |
|tags|jsonb|A list of tags that are attached to the specified IAM OIDC provider. The returned list of tags is sorted by tag key. For more information about tagging, see Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the IAM User Guide. |
|thumbprint_list|text[]|A list of certificate thumbprints that are associated with the specified IAM OIDC provider resource object. For more information, see CreateOpenIDConnectProvider. |
|url|text|The URL that the IAM OIDC provider resource object is associated with. For more information, see CreateOpenIDConnectProvider. |
