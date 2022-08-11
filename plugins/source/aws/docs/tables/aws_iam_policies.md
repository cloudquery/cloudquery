
# Table: aws_iam_policies
Contains information about a managed policy, including the policy's ARN, versions, and the number of principal entities (users, groups, and roles) that the policy is attached to.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|arn|text|The Amazon Resource Name (ARN). ARNs are unique identifiers for AWS resources. For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the AWS General Reference. |
|attachment_count|integer|The number of principal entities (users, groups, and roles) that the policy is attached to. |
|create_date|timestamp without time zone|The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the policy was created. |
|default_version_id|text|The identifier for the version of the policy that is set as the default (operative) version. For more information about policy versions, see Versioning for managed policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-versions.html) in the IAM User Guide. |
|description|text|A friendly description of the policy. |
|is_attachable|boolean|Specifies whether the policy can be attached to an IAM user, group, or role. |
|path|text|The path to the policy. For more information about paths, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide. |
|permissions_boundary_usage_count|integer|The number of entities (users and roles) for which the policy is used as the permissions boundary. For more information about permissions boundaries, see Permissions boundaries for IAM identities (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html) in the IAM User Guide. |
|id|text|The stable and unique string identifying the policy. For more information about IDs, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide. |
|name|text|The friendly name (not ARN) identifying the policy. |
|update_date|timestamp without time zone|The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the policy was last updated. When a policy has only one version, this field contains the date and time when the policy was created. When a policy has more than one version, this field contains the date and time when the most recent policy version was created. |
