
# Table: aws_iam_policy_versions
Contains information about a version of a managed policy.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|policy_cq_id|uuid|Policy CloudQuery ID the policy versions belongs too.|
|policy_id|text|Policy ID the policy versions belongs too.|
|create_date|timestamp without time zone|The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the policy version was created. |
|document|jsonb|The policy document. The policy document is returned in the response to the GetPolicyVersion and GetAccountAuthorizationDetails operations. It is not returned in the response to the CreatePolicyVersion or ListPolicyVersions operations. The policy document returned in this structure is URL-encoded compliant with RFC 3986 (https://tools.ietf.org/html/rfc3986). You can use a URL decoding method to convert the policy back to plain JSON text. For example, if you use Java, you can use the decode method of the java.net.URLDecoder utility class in the Java SDK. Other languages and SDKs provide similar functionality. |
|is_default_version|boolean|Specifies whether the policy version is set as the policy's default version. |
|version_id|text|The identifier for the policy version. Policy version identifiers always begin with v (always lowercase). When a policy is created, the first policy version is v1. |
