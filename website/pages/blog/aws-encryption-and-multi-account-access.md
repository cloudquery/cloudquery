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

In this blog post, we’ll focus on how to encrypt everything in multi-account AWS environments and how to make encryption decisions with your unique environment and data security needs in mind.

In follow-up blog posts of this series on encryption, we’ll follow up with information on how CloudQuery can help shortly after the release of CloudQuery v1 and other blog posts deep-diving into encryption and data security in cloud.

## Multi Account Access in AWS

For symmetric encryption, AWS offers [2 primary services](https://docs.aws.amazon.com/crypto/latest/userguide/awscryp-service-toplevel.html): **[AWS Key Management Service (KMS)](https://docs.aws.amazon.com/crypto/latest/userguide/awscryp-service-toplevel.html) and [CloudHSM](https://aws.amazon.com/cloudhsm/)**. When AWS KMS was [first announced in 2014](https://aws.amazon.com/blogs/aws/new-key-management-service/), it was launched to support encrypting data at rest for S3, EBS, and Redshift. Now, KMS supports multiple different types of keys including symmetric encryption keys, asymmetric keys for encryption or signing, and HMAC keys to generate and verify HMAC tags. [KMS now supports many more services](https://aws.amazon.com/kms/features/#AWS_Service_Integration).  Some of those services are available for direct access from other AWS accounts, such as S3, SQS, Secrets Manager, and more.

AWS provides a table to describe AWS KMS keys and information about how they can be used and their features as shown below. 