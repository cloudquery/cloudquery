---
title: Auditing Service Account Keys with CloudQuery
tag: security
date: 2022/11/04
description: >-
  A case against Service Account Keys.
author: SCKelemen, Alex
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>


# Intro
There are several common methods of accessing cloud 
environments, each with their own risks and uses. Service 
Account Keys, however, are one of the more dangerous methods.
Before using service account keys, there are several risks 
and shortfalls you should consider:

Service Account Keys offer a simple and easy mechanism for 
granting access to cloud resources. Service Account Keys are 
easy to understand, and are compatible with most existing 
architectures and systems. 

Initially, you might be appealed by the ease of using 
account keys. All you need to do is simply put the key
somewhere your systems can access. This could be a file,
environment, variable, or bucket.

Unfortunately, with this highly flexible global bearer strategy,
tracking the location of the key becomes unmanageable. Multiple 
users and services can use the same key, defeating the 
non-repudiation principle. It also creates perverse incentives,
where reusing an existing key becomes easier, than creating a new,
purpose built key.

This makes understanding your system and access policy difficult,
and amplifies the impact of lost or stolen keys. Additionally,
rotating these keys because a complex and difficult challenge,
as it's difficult to find every use of the key, and even more difficult
to find a window where all services and uses can be upgraded together.

As the needs of your services evolve, these Service Accounts tend
to gain more permissions, which leaves overly permissed keys. All
services should have their own accounts with the minimum permissions
required.

# Service Accounts vs Service Account Keys

Service Accounts are accounts created for use by a service.
These are generally not interactive, often programmatic accounts.
Service Accounts allow you to associate an identity with a workload,
and allow permissions to be attached to that identity.

Service Account Keys, are cryptographic keys or shared secrets that
act as a form a proof to authenticate a call to a Service Account.
Service Account Keys are just another way to use Service Accounts, albeit
one with significant risks.


Typically, the account keys would look something like this:

GCP Service Account Key:
```yaml
id: "key ID"
private_key: >-
  -----BEGIN PRIVATE KEY-----
  PRIVATE_KEY
  -----END PRIVATE KEY-----
client_id: "client ID"
...
```

AWS Account Key:
```yaml
AccessKeyID: "AKIAIOSFODNN7EXAMPLE"
SecretKey:   "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
```

## Service Account Keys vs Secrets

It is probably tempting to store your Service Account Keys in
a Key or Secret Management system such as AWS Secret Manager, GCP Secrets Manager,
Azure KeyVault, or Hashicorp Vault. These systems allow secure
dissemination of secret material to services or users.

However, this approach is also fraught with peril.

Although secret managers allow secure storage, once the secret material leaves
the storage system, the security of the secret material is reliant on the client,
and is vulnerable to leaks.
Additionally, accessing this secret material requires authentication and authorization,
implying some preexisting privilege and identity. Regardless of the secure storage,
the identity problem still persists.

Since you are already using Identity and Access Management to access your secrets, 
why not lean into the IAM platform and use it for authorizing the end service directly?

## AWS

AWS has a feature called [Access Keys](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html). 
Access Keys can be disabled by marking them inactive, so at the very least you can quickly pull the authorization without 
completely getting rid of the keys. 

## GCP

GCP has two types of Service Account Keys: `SYSTEM_MANAGED` and `USER_MANAGED` keys.

`SYSTEM_MANAGED` keys are keys used internally by GCP for creating short-lived service account credentials,
and for signing blobs and JSON Web Tokens. Users don’t have control over these keys.

`USER_MANAGED` keys are keys created by end users. These service account keys should
be avoided.


# Implementing checks

## Queries

### AWS

Find all Access Keys:
```sql
select * from aws_iam_user_access_keys;
```

Permissions: 

- `iam:CreateAccessKey`
- `iam:DeleteAccessKey`
- `iam:GetAccessKeyLastUsed`
- `iam:GetUser`
- `iam:ListAccessKeys`
- `iam:UpdateAccessKey`

### GCP

Find all user-created Service Account Keys:
```sql
select * from gcp_iam_service_account_keys where key_type != 'SYSTEM_MANAGED';
```


# Ensuring compliance

If you still prefer to stick with account keys, you can at least make these checks regular. 
Setup a CQ sync and run it every day/week and use the queries above to ensure everything is okay.

The better option would be to drop the account keys usage altogether.

## Enforcing 0 Account Key Policy

AWS and GCP allow you to restrict the creation of service account keys with Organizational Policies and Service Control Policies. 

AWS Service Control Policy:
```json
{
  "Version": "2012-10-17",
  "Statement": {
    "Effect": "Deny",
    "Action": "iam:CreateAccessKey",
    "Resource": "arn:aws:iam::/*",
}
```

GCP Organizational Policy:
```jsx
constraints/iam.disableServiceAccountKeyCreation is:True
```


# Conclusion

There are many benefits that are offered by cloud infrastructure. Perhaps one 
of the greatest being high quality Identity and Access Management. By using a
cloud provider such as AWS, Azure, or GCP, you are implicitly trusting the provider. 
Leaning into the cloud and leveraging cloud native Identity and Access Management
is a powerful, and strongly beneficial tactic. 