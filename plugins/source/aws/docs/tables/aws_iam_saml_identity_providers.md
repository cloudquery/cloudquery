
# Table: aws_iam_saml_identity_providers
SAML provider resource objects defined in IAM for the AWS account.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|arn|text|Amazon Resource Name (ARN) of the saml identity provider.|
|create_date|timestamp without time zone|The date and time when the SAML provider was created. |
|saml_metadata_document|text|The XML metadata document that includes information about an identity provider. |
|tags|jsonb|A list of tags that are attached to the specified IAM SAML provider. The returned list of tags is sorted by tag key. For more information about tagging, see Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the IAM User Guide. |
|valid_until|timestamp without time zone|The expiration date and time for the SAML provider. |
