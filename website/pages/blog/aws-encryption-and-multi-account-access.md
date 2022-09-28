---
title: Encryption in AWS and Multi-Account Access
tag: security
date: 2022/09/29
description: >-
  How to encrypt in AWS given a multi-account environment.
author: jsonkao
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

## *Dance like nobody is watching. Encrypt like everyone is.*

 - *Werner Vogels, Amazon CTO*

As AWS [outlined at their 2022 Re:Inforce Security Conference](https://www.youtube.com/watch?v=PPunA7tPMyk&t=3062s) and [mentioned by Werner Vogels at an AWS Summit in 2019](https://youtu.be/vWfkbGF6fiA?t=4339), *encrypt everything* and *encrypt like everyone is [watching]*.  

![**AWS Summit Series 2019 - Santa Clara: Keynote featuring Werner Vogels**](/images/blog/aws-encryption-and-multi-account-access/encrypt-like-watching.png)
*AWS Summit Series 2019 - Santa Clara: Keynote featuring Werner Vogels*

In this blog post, we’ll focus on how to encrypt everything in multi-account AWS environments and how to make encryption decisions with your unique environment and data security needs in mind.

In follow-up blog posts of this series on encryption, we’ll follow up with information on how CloudQuery can help shortly after the release of CloudQuery v1 and other blog posts deep-diving into encryption and data security in cloud.

## Multi Account Access in AWS

For symmetric encryption, AWS offers [2 primary services](https://docs.aws.amazon.com/crypto/latest/userguide/awscryp-service-toplevel.html): **[AWS Key Management Service (KMS)](https://docs.aws.amazon.com/crypto/latest/userguide/awscryp-service-toplevel.html) and [CloudHSM](https://aws.amazon.com/cloudhsm/)**. When AWS KMS was [first announced in 2014](https://aws.amazon.com/blogs/aws/new-key-management-service/), it was launched to support encrypting data at rest for S3, EBS, and Redshift. Now, KMS supports multiple different types of keys including symmetric encryption keys, asymmetric keys for encryption or signing, and HMAC keys to generate and verify HMAC tags. [KMS now supports many more services](https://aws.amazon.com/kms/features/#AWS_Service_Integration).  Some of those services are available for direct access from other AWS accounts, such as S3, SQS, Secrets Manager, and more.

AWS provides a table to describe AWS KMS keys and information about how they can be used and their features as shown below. 

![[AWS Table for Customer Keys and AWS Keys](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html)](/images/blog/aws-encryption-and-multi-account-access/aws-kms-table.png)
*[AWS Table for Customer Keys and AWS Keys](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html)*

We’re going to enrich that table with more detail regarding multi-account access and management while expanding the table with more information about different types of customer managed keys with different key material origins.  We’ve differentiated between the 3 types of Customer Managed Keys: Keys with External Key material, Keys backed by a Custom Key Store (CloudHSM), and lastly, keys with AWS-Provided Key material.  

**Expanded Table:**

| Key Type | Multi-account Access | Can view metadata | Can manage KMS Key | Used Only for my AWS account | Automatic Rotation | Pricing |
| --- | --- | --- | --- | --- | --- | --- |
| Customer Managed Key: External Key Material | Yes | Yes | Yes | Yes (1)  | No | Monthly & Per-Use Fee |
| Customer Managed Key: Custom Key Store | Yes | Yes | Yes | Yes (1) | No | Monthly & Per-Use Fee |
| Customer Managed Key: AWS-Provided Key Material | Yes | Yes | Yes | Yes | Optional | Monthly & Per-Use Fee |
| AWS Managed Key | No | Yes | No | Yes | Required | Per-use fee |
| AWS Owned Key | Varies | No | No | No | Varies | Varies |

### Multi-Account Access in AWS

In advanced use cases, enterprise cloud workloads may be split up by infrastructure or by project.  For example, one account may host data such as a data lake account and another account may host compute resources.  In these multi-account AWS environments, cross-account access to resources can be necessary to reduce complexity and to reduce the need for data and resource duplication.

With cross-account access to resources, encryption also plays a role in how cross-account access can be granted to users and applications originating from a different account.  The type of KMS Key chosen can affect how cross-account setup can be done and in some cases, make it more complex to manage.

We’ll walk through a setup where cross-account access may be desired.

![Cross-Account Access in AWS to an Encrypted S3 Bucket](/images/blog/aws-encryption-and-multi-account-access/cross-account-diagrampng)

Cross-Account Access in AWS to an Encrypted S3 Bucket

In this example, we’ll use the example where an IAM role in a Compute AWS Account needs access to an encrypted S3 bucket in the Data AWS Account.  One method to grant access is to configure the following components if we’re using a customer managed KMS Key:

- IAM Role in the Compute AWS Account with corresponding policies that grant access to the S3 Bucket, Objects, and KMS Key in the Data AWS Account.
- S3 Bucket with a bucket policy in the Data AWS Account that grants access to the IAM Role from the Compute AWS Account.
- KMS Key with a Key Policy in the Data AWS Account that grants access to the IAM Role from the Compute AWS Account.

Let’s revisit the encryption key table that we expanded upon earlier in this post with the additional column for multi-account access.  What happens when we try to use an AWS-managed KMS Key for this cross-account use case?  Note - this is different than the Amazon S3-managed key (SSE-S3) option, which if there’s interest, we’ll do a deep dive into the encryption and security settings available for AWS S3 including S3-managed keys, bucket keys, and more.
