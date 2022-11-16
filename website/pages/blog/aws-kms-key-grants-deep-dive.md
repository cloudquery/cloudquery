---
title: A Deep Dive on AWS KMS Key Access and AWS Key Grants
tag: security
date: 2022/11/16
description: >-
  A Technical Deep Dive on AWS KMS Key Access and AWS Key Grants, one mechanism for granting access to AWS KMS Keys.
author: jsonkao
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>
We recently published a [blog post](https://www.cloudquery.io/blog/aws-encryption-and-multi-account-access) about Encryption in AWS and Multi-Account Access.  As a follow-up to that post, we’re focusing on encryption access in AWS and information on how CloudQuery can help including how to audit access to those encryption keys and the underlying data to help assist with your security and compliance needs.

CloudQuery has released the following updates to encryption resources to aid with Encryption discovery in AWS:

- Support for [KMS Key Aliases](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/aws/docs/tables/aws_kms_aliases.md)
- Support for [KMS Key Grants](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/aws/docs/tables/aws_kms_key_grants.md)
- Support for [CloudHSMv2](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/aws/docs/tables)

In this post, we will deep dive on KMS Key Access via KMS Key Grants and best practices with KMS Key Grants.  Access via KMS Key Grants can be a forgotten means of allowing unauthorized applications, users, and other undesired access to use and manage KMS Keys.

## KMS Key Access

[KMS Key Access](https://docs.aws.amazon.com/kms/latest/developerguide/control-access.html) is similar to other resource access in AWS and the underlying access system in AWS: AWS Identity and Access Management.  [Typical AWS evaluation of access](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html) to a resource is done via AWS’s policy evaluation logic that evaluates the request context, evaluates whether the actions are within a single account or [cross-account](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic-cross-account.html) (between 2 distinct AWS accounts), and evaluating identity-based policies with resource-based policies and other advanced policies such as permission boundaries, Organizationals Service-Control Policies, Session Policies, and more.

From AWS: 

![Policy Evaluation for evaluating identity-based policies with resource-based policies](/images/blog/aws-kms-key-grants-deep-dive/policy-eval-resource.png)

Policy Evaluation for evaluating identity-based policies with resource-based policies

Typical policy evaluation within a single account is done as a logical ****or****, where an explicit permission on either the resource-based policy or the identity-based policy evaluations allows for access to the resource for [resources that support resource-based policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-services-that-work-with-iam.html).  More information from AWS can be found [here](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_identity-vs-resource.html).  An example is an S3 Bucket Policy.  If an S3 Bucket Policy allows an IAM entity within the account explicit access to the S3 Bucket, the IAM entity does not need explicit allows on the identity-based policies for access to the S3 Bucket.

Unlike most other resource-based policies, [KMS Key access within an AWS Account](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html#policy-eval-basics) ********must******** be granted on the KMS resource-based policy: the KMS Key policy.  If the KMS Key policy doesn’t explicitly allow access, an identity cannot access the KMS key via identity-based policies.  We’ll deep dive on KMS Key policies and recommendations in a later post.

However, KMS Keys are a unique resource in AWS, where there’s a third access mechanism for KMS Keys: KMS Key Grants. 

There are [3 access mechanisms for KMS Keys](https://docs.aws.amazon.com/kms/latest/developerguide/control-access.html):

- Key Policy (Resource-based policy)
- IAM Policies
- Grants

Thus, we have the following important differences about access to KMS Keys:

- Permissions must be granted explicitly on the KMS Key Policy (Resource-based policy).
- Access to KMS Keys can also be managed by KMS Key Grants.

## Access via KMS Key Grants

Grants allow for AWS principals to use KMS Keys.  One such example is when [sharing encrypted EBS Snapshots](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-modifying-snapshot-permissions.html#share-kms-key).  Grants can also be used for temporary permissions because they can be used without modifying key policies or IAM policies.  

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

The permission to create a grant differs slightly depending on how the grantee principal gets permissions to call CreateGrant:

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

[https://docs.aws.amazon.com/kms/latest/developerguide/create-grant-overview.html#grant-creategrant](https://docs.aws.amazon.com/kms/latest/developerguide/create-grant-overview.html#grant-creategrant)

## Logging of KMS Key Grants

CloudTrail keeps a record of events and actions [taken via the AWS CLI, SDKs & APIs, and Management Console](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-user-guide.html).   By default, CloudTrail logs management events and not data events.  

CloudTrail records some activity made outside of your AWS Account.  Activity outside of the AWS Account that can be recorded include:

- KMS Key Grants on a KMS Key in your AWS Account.
- EBS on SharedSnapshotVolumeCreated and SharedSnapshotCopyInitiated for Shared EBS Snapshots.

Even if the IAM principal creating a key Grants made on KMS resources within an account (Account A) lives outside of your account (Account B), the `CreateGrant` API call will still show up in account A's CloudTrail logs.

![CloudTrail Event Record for KMS CreateGrant](/images/blog/aws-kms-key-grants-deep-dive/create-grant-cloudtrail.png)

- [https://docs.aws.amazon.com/kms/latest/developerguide/security-logging-monitoring.html](https://docs.aws.amazon.com/kms/latest/developerguide/security-logging-monitoring.html)

## Managing KMS Key Grant Lifecycle

Retire vs Revoke

## Recommendations

Along with [AWS Best Practices](https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant-best-practices), we recommend the additional following practices:

- Limit Grant Ability to usage of the KMS Key and not management of the KMS Key.

Grant Operations [(from AWS)](https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant-concepts):

- Cryptographic operations
    - [Decrypt](https://docs.aws.amazon.com/kms/latest/APIReference/API_Decrypt.html)
    - [Encrypt](https://docs.aws.amazon.com/kms/latest/APIReference/API_Encrypt.html)
    - [GenerateDataKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_GenerateDataKey.html)
    - [GenerateDataKeyPair](https://docs.aws.amazon.com/kms/latest/APIReference/API_GenerateDataKeyPair.html)
    - [GenerateDataKeyPairWithoutPlaintext](https://docs.aws.amazon.com/kms/latest/APIReference/API_GenerateDataKeyPairWithoutPlaintext.html)
    - [GenerateDataKeyWithoutPlaintext](https://docs.aws.amazon.com/kms/latest/APIReference/API_GenerateDataKeyWithoutPlaintext.html)
    - [GenerateMac](https://docs.aws.amazon.com/kms/latest/APIReference/API_GenerateMac.html)
    - [ReEncryptFrom](https://docs.aws.amazon.com/kms/latest/APIReference/API_ReEncrypt.html)
    - [ReEncryptTo](https://docs.aws.amazon.com/kms/latest/APIReference/API_ReEncrypt.html)
    - [Sign](https://docs.aws.amazon.com/kms/latest/APIReference/API_Sign.html)
    - [Verify](https://docs.aws.amazon.com/kms/latest/APIReference/API_Verify.html)
    - [VerifyMac](https://docs.aws.amazon.com/kms/latest/APIReference/API_VerifyMac.html)
- Other operations
    - [CreateGrant](https://docs.aws.amazon.com/kms/latest/APIReference/API_CreateGrant.html)
    - [DescribeKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html)
    - [GetPublicKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_GetPublicKey.html)
    - [RetireGrant](https://docs.aws.amazon.com/kms/latest/APIReference/API_RetireGrant.html)
- Limit Delegation abilities to only the Key Administrators and Services

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

- Limit Grant Delegation to only AWS Services usage.

We’ll use the `kms:GrantIsForAWSResource` Condition Key here.  An example below is for EC2 Autoscaling with encrypted volumes.

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

- Logging and Monitoring KMS Key Grants
    
    
- Control usage and lifecycle of KMS Key Grants.
    - Limit usage of KMS Key Grants to AWS Services.
        
        ```sql
        SELECT * from aws_kms_key_grants where (grantee_principal NOT LIKE '%.amazonaws.com' AND grantee_principal NOT LIKE 'AWS Internal');
        ```
        
    - Ensure KMS Key Grant Lifecycles are managed properly.  AWS Services retires grants as soon as the task is complete.
    - Use Condition Keys
    - However, this can be configured to only allow usage to AWS Services with the condition key `kms:GrantIsForAWSResource.`

Your organization’s use cases may be slightly different. If you have comments, feedback on this post, follow-up topics you’d like to see, or would like to talk about your KMS and encryption experiences  - email us at security@cloudquery.io or come chat with us on [Discord](https://www.cloudquery.io/discord)!