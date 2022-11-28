---
title: A Deep Dive on AWS KMS Key Access and AWS Key Grants
tag: security
date: 2022/11/17
description: >-
  A Technical Deep Dive on AWS KMS Key Access and AWS Key Grants.
author: jsonkao
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>
We recently published a [blog post](/blog/aws-encryption-and-multi-account-access) about Encryption in AWS and Multi-Account Access.  As a follow-up to that post, we’re focusing on encryption access in AWS and information on [how CloudQuery can help](/docs/plugins/sources/overview) including how to audit access to those encryption keys and the underlying data to help assist with your security and compliance needs.

CloudQuery has released the following updates to encryption resources to aid with Encryption discovery in AWS in addition to our [existing coverage of AWS](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/aws/README.md) and [AWS resources](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/aws/docs/tables):

- Support for [KMS Key Aliases](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/aws/docs/tables/aws_kms_aliases.md)
- Support for [KMS Key Grants](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/aws/docs/tables/aws_kms_key_grants.md)
- Support for [CloudHSMv2](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/aws/docs/tables)

In this post, we will deep dive on KMS Key Access via KMS Key Grants and best practices with KMS Key Grants.  Access via KMS Key Grants can be a forgotten means of allowing unauthorized applications, users, and other undesired access to use and manage KMS Keys.

## KMS Key Access

[KMS Key Access](https://docs.aws.amazon.com/kms/latest/developerguide/control-access.html) is similar to other resource access in AWS and the underlying access system in AWS: AWS Identity and Access Management.  [Typical AWS evaluation of access](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html) to a resource is done via AWS’s policy evaluation logic that evaluates the request context, evaluates whether the actions are within a single account or [cross-account](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic-cross-account.html) (between 2 distinct AWS accounts), and evaluating identity-based policies with resource-based policies and other advanced policies such as permission boundaries, Organizational Service-Control Policies, Session Policies, and more.

The below image is from AWS's [documentation regarding policy evaluation](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html) for evaluating identity-based policies with resource-based policies.

![From AWS, Policy Evaluation for evaluating identity-based policies with resource-based policies](/images/blog/aws-kms-key-grants-deep-dive/policy-eval-resource.png)

Typical policy evaluation within a single account is done as a logical ****or****, where an explicit permission on either the resource-based policy or the identity-based policy evaluations allows for access to the resource for [resources that support resource-based policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-services-that-work-with-iam.html).  More information from AWS can be found [here](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_identity-vs-resource.html).  An example is an S3 Bucket Policy.  If an S3 Bucket Policy allows an IAM entity within the account explicit access to the S3 Bucket, the IAM entity does not need explicit allows on the identity-based policies for access to the S3 Bucket.

Unlike most other resource-based policies, [KMS Key access within an AWS Account](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html#policy-eval-basics) ********must******** be granted on the KMS resource-based policy: the KMS Key policy.  If the KMS Key policy doesn’t explicitly allow access, an identity cannot access the KMS key via identity-based policies.  We’ll deep dive on KMS Key policies and recommendations in a later post.

However, KMS Keys are a unique resource in AWS, where there’s a third access mechanism for KMS Keys: KMS Key Grants. KMS Key Grants are not manageable or viewable via the AWS Management Console and can only be managed and viewed via AWS CLI and API/SDK.

There are [3 access mechanisms for KMS Keys](https://docs.aws.amazon.com/kms/latest/developerguide/control-access.html):

- Key Policy (Resource-based policy)
- IAM Policies
- Grants

Thus, we have the following important differences about access to KMS Keys vs access to other resources in AWS:

- Permissions must be granted explicitly on the KMS Key Policy (Resource-based policy).
- Access to KMS Keys can also be managed by KMS Key Grants.

## Access via KMS Key Grants

Grants allow for AWS principals to use KMS Keys.  One such example is when [sharing encrypted EBS Snapshots](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-modifying-snapshot-permissions.html#share-kms-key).  Grants can also be used for temporary permissions because they can be used without modifying key policies or IAM policies.  AWS also (recommends using KMS Key Grants)[https://docs.aws.amazon.com/kms/latest/developerguide/resource-limits.html#key-policy-limit] if the Key Policy size approaches the limit of 32 KB.

As such, grants are used by [AWS services that integrate with AWS KMS](https://docs.aws.amazon.com/kms/latest/developerguide/service-integration.html) such as RDS and ACM.

We’ve submitted requests to AWS to update their AWS Key Management Service Documentation (on 11/14/22) to better reflect service integration with KMS: [https://docs.aws.amazon.com/kms/latest/developerguide/service-integration.html](https://docs.aws.amazon.com/kms/latest/developerguide/service-integration.html)

### Grant Permissions

The following permissions and associated API/CLI actions are related to managing KMS Key Grants:

- `kms:CreateGrant`
- `kms:ListGrants`
- `kms:ListRetirableGrants`
- `kms:RetireGrant`
- `kms:RevokeGrant`

We’ll go into detail on some of these permissions later in this post.

A sample key policy that allows for a key administrator to Create Grants is as follows.  Note with this policy, we’ve modified the default key policy from AWS so that it does not [enable IAM policies](https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-default.html).  If you plan on using this key policy, ensure that any other permissions necessary are either explicitly allowed on the key policy or are enabled in IAM similar to the `Allow access to key metadata to the account` statement.

```json
{
    "Sid": "AllowRootAccountAccess",
    "Effect": "Allow",
    "Principal": {
        "AWS": "arn:aws:iam::123412341234:root"
    },
    "Action": "kms:*",
    "Resource": "*",
    "Condition": {
        "StringEquals": {
            "aws:PrincipalType": "Account"
        }
    }
},
{
    "Sid": "AllowKeyAdministrator",
    "Effect": "Allow",
    "Principal": {
        "AWS": "arn:aws:iam::123412341234:role/KeyAdminRole"
    },
    "Action": "kms:*",
    "Resource": "*"
},
{
     "Sid": "Allow access to key metadata to the account",
     "Effect": "Allow",
     "Principal": {
         "AWS": "arn:aws:iam::123412341234:root"
     },
     "Action": [
         "kms:Describe*",
         "kms:Get*",
         "kms:List*"
      ],
      "Resource": "*"
 }
```

### Creating Grants

```bash
aws kms create-grant --key-id arn:aws:kms:us-east-1:123412341234:key/aaaaaaaa-1234-1234-1234-123412341234 --grantee-principal arn:aws:iam::123412341234:user/sample-user --operations DescribeKey --profile yourprofile --region us-east-1
```

The ability and permissions to create a grant must either come from a policy or a grant.  If no grants exist, the permissions must come from the KMS Key Policy or if the Key Policy allows for management via identity policies, can then come from identity policies.

The permission to [create a grant](https://docs.aws.amazon.com/kms/latest/developerguide/create-grant-overview.html#grant-creategrant) differs slightly depending on how the grantee principal gets permissions to call CreateGrant:

|  | Policy as Permission Source | Grant as Permission Source |
| --- | --- | --- |
| Constraints | Refer to Policy for Constraints, kms:GrantConstraintType condition key. | Must be as strict as the parent grant |
| Operations | Refer to Policy if kms:GrantOperations is used. | Can only allow a subset (some or all) of the operations in the parent grant. |

### Nested Granting

A Grantee Principal that was granted access to `CreateGrant` via a Grant can then create child grants.  

```bash
aws kms create-grant --key-id arn:aws:kms:us-east-1:123412341234:key/aaaaaaaa-1234-1234-1234-123412341234 --grantee-principal arn:aws:iam::123412341234:user/hugh-grant --operations DescribeKey CreateGrant --profile myprofilewithoutpermpolicy --region us-east-1
```
In the above code snippet, we've granted our `hugh-grant` user the permissions to `DescribeKey` and `CreateGrant`.

Now, we’ll try the following 3 use cases:

- Creating a Grant with more operations than the parent grant.
```bash
aws kms create-grant --key-id arn:aws:kms:us-east-1:123412341234:key/aaaaaaaa-1234-1234-1234-123412341234 --grantee-principal arn:aws:iam::123412341234:user/sample-user --operations DescribeKey Decrypt CreateGrant --profile hughgrant --region us-east-1
```

- Creating a Grant with the same operations as the parent grant.
```bash
aws kms create-grant --key-id arn:aws:kms:us-east-1:123412341234:key/aaaaaaaa-1234-1234-1234-123412341234 --grantee-principal arn:aws:iam::123412341234:user/sample-user --operations DescribeKey CreateGrant --profile hughgrant --region us-east-1
```

- Creating a Grant with fewer operations as the parent grant.
```bash
aws kms create-grant --key-id arn:aws:kms:us-east-1:123412341234:key/aaaaaaaa-1234-1234-1234-123412341234 --grantee-principal arn:aws:iam::123412341234:user/sample-user --operations DescribeKey --profile hughgrant --region us-east-1
```

In this case, the following use cases succeed since the child grant is as strict or stricter than the parent grant:
- Creating a Grant with the same operations as the parent grant.
- Creating a Grant with fewer operations as the parent grant.

The following use case failed since the child grant is not as strict or stricter than the parent grant:
- Creating a Grant with more operations than the parent grant.  

With the failed use case, the error message on CLI will display a more generic `AccessDeniedException` error message that will reference `no identity-based policy allows the kms:CreateGrant action`.  

![Example Failure of KMS Create Grant](/images/blog/aws-kms-key-grants-deep-dive/grant-error.png)

## Logging of KMS Key Grants

CloudTrail keeps a record of events and actions [taken via the AWS CLI, SDKs & APIs, and Management Console](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-user-guide.html).   By default, CloudTrail logs management events and not data events.  

CloudTrail records some activity made outside of your AWS Account.  Activity outside of the AWS Account that can be recorded include:

- KMS Key Grants on a KMS Key in your AWS Account.
- EBS on SharedSnapshotVolumeCreated and SharedSnapshotCopyInitiated for Shared EBS Snapshots.

Even if the IAM principal creating a key Grants made on KMS resources within an account (Account A) lives outside of your account (Account B), the `CreateGrant` API call will still show up in account A's CloudTrail logs.

![CloudTrail Event Record for KMS CreateGrant](/images/blog/aws-kms-key-grants-deep-dive/create-grant-cloudtrail.png)

Note that the `userIdentity` [element of the CloudTrail event record](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-event-reference-user-identity.html) may contain unexpected results for the `type` field.  Typically, this will specify the type of identity such as `Root`, `IAMUser`, `AssumedRole`, `Role`, `FederatedUser`, `Directory`, `AWSService`, `Unknown`, or `AWSAccount`.  For cross-account access, `AWSAccount` may appear in the logs.  In the example above, `AWSAccount` shows up for cross-account access using an IAM user. We've pointed out this difference to AWS.  

## Managing KMS Key Grant Lifecycle

AWS provides 2 different operations for managing KMS Key Grant lifecycle:
- [RetireGrant](https://docs.aws.amazon.com/kms/latest/APIReference/API_RetireGrant.html)
- [RevokeGrant](https://docs.aws.amazon.com/kms/latest/APIReference/API_RevokeGrant.html)

While both of these actions provide the same result of deleting a KMS Key Grant, which eliminates the permissions the grant allows. This is one of a few use cases in AWS where multiple operations may provide the same result. However, these permissions are authorized differently.  

A few key differences include:
- RevokeGrant is similar to a typical AWS KMS operation and can be controlled via key policies and IAM policies via the `kms:RevokeGrant` permission.  The usage of `kms:RetireGrant` in policies does not behave similarly.

- RetireGrant can be granted by specifying a retiring principal when creating the grant.  RetireGrant can also be granted to the Grantee Principal similar to other available operations in CreateGrant.  Principals specified in the grant can then retire a grant without the `kms:RetireGrant` permission explicitly allowed in either a key policy or an identity policy.

```bash
aws kms create-grant --key-id arn:aws:kms:us-east-1:123412341234:key/aaaaaaaa-1234-1234-1234-123412341234 --grantee-principal arn:aws:iam::123412341234:user/sample-user --retiring-principal arn:aws:iam::123412341234:user/sample-principal --operations DescribeKey --profile testprofile --region us-east-1
```

```bash
aws kms create-grant --key-id arn:aws:kms:us-east-1:123412341234:key/aaaaaaaa-1234-1234-1234-123412341234 --grantee-principal arn:aws:iam::123412341234:user/sample-user --operations DescribeKey RetireGrant --profile testprofile --region us-east-1
```

## Recommendations

Along with [AWS Best Practices for KMS Key Grants](https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant-best-practices), we recommend the additional following practices:

1. Limit grant ability to usage of the KMS Key and not management of future grants.  

* The list of [supported grant operations provided by AWS](https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant-concepts) includes cryptographic operations such as `Decrypt` and other operations such as `CreateGrant`.  We recommend limiting the ability to `CreateGrant` to only Key Administrators and first-level grants.

2. Further limit delegation abilities to only the Key Administrators and Services

* We recommend limiting the ability to delegate access to the KMS key to only Key Administrators and Services.  This includes the `kms:CreateGrant` permission and the `kms:PutKeyPolicy`.  Below is an example key policy that limits key delegation abilities to the root account user and to the KeyAdminRole, but does not enable IAM policy access. 

```json
{
    "Sid": "AllowRootAccountAccess",
    "Effect": "Allow",
    "Principal": {
        "AWS": "arn:aws:iam::123412341234:root"
    },
    "Action": "kms:*",
    "Resource": "*",
    "Condition": {
        "StringEquals": {
            "aws:PrincipalType": "Account"
        }
    }
},
{
    "Sid": "AllowKeyAdministrator",
    "Effect": "Allow",
    "Principal": {
        "AWS": "arn:aws:iam::123412341234:role/KeyAdminRole"
    },
    "Action": "kms:*",
    "Resource": "*"
}
```

3. Limit grant delegation to AWS Services usage where possible.

* This can be done by the `kms:GrantIsForAWSResource` Condition Key here.  An example key policy for EC2 Autoscaling with encrypted volumes is below.

```json
{
   "Sid": "Allow service-linked role use of the customer managed key",
   "Effect": "Allow",
   "Principal": {
       "AWS": [
           "arn:aws:iam::123412341234:role/aws-service-role/autoscaling.amazonaws.com/AWSServiceRoleForAutoScaling"
       ]
   },
   "Action": [
       "kms:Encrypt",
       "kms:Decrypt",
       "kms:ReEncrypt*",
       "kms:GenerateDataKey*",
       "kms:DescribeKey"
   ],
   "Resource": "*"
},
{
   "Sid": "Allow attachment of persistent resources",
   "Effect": "Allow",
   "Principal": {
       "AWS": [
           "arn:aws:iam::123412341234:role/aws-service-role/autoscaling.amazonaws.com/AWSServiceRoleForAutoScaling"
       ]
   },
   "Action": [
       "kms:CreateGrant"
   ],
   "Resource": "*",
   "Condition": {
       "Bool": {
           "kms:GrantIsForAWSResource": true
       }
    }
}
```

4. Logging and Monitoring KMS Key Grants

* We recommend monitoring KMS Key Grants in CloudTrail as shown in the logging section above to monitor for external accounts and undesired usage and creation of KMS Key Grants. 

5. Take inventory of KMS Key Grants and validate usage of KMS Key Grants.

* The following CloudQuery query in postgresql finds grants where the grantee principal is not an AWS service.
```sql
SELECT * from aws_kms_key_grants where (grantee_principal NOT LIKE '%.amazonaws.com' AND grantee_principal NOT LIKE 'AWS Internal');
```

6. If retiring/revoking KMS Key Grants, ensure the Key Grant is not being used by production systems to avoid outages and adverse impact.

7. Determine appropriate organization and team strategy for different access mechanisms for KMS regarding KMS Key Policies, Identity Policies, and KMS Key Grants.  

If you have comments, feedback on this post, follow-up topics you’d like to see, or would like to talk about your KMS and encryption experiences  - email us at security@cloudquery.io or come chat with us on [Discord](https://www.cloudquery.io/discord)!  We'd love to hear your feedback and thoughts on encryption and your experiences with KMS Key Grants.

## References and Useful Links

[AWS: Grants in AWS KMS](https://docs.aws.amazon.com/kms/latest/developerguide/grants.html)

[AWS: Authentication and access control for AWS KMS](https://docs.aws.amazon.com/kms/latest/developerguide/control-access.html)

[AWS: Logging and monitoring in AWS Key Management Service](https://docs.aws.amazon.com/kms/latest/developerguide/security-logging-monitoring.html)

[CloudQuery: Source Plugins](/docs/plugins/sources/overview)

[CloudQuery: AWS Plugin](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/aws)