---
title: How to Set Up Cross Account Access in AWS with AssumeRole
tag: security
description: >-
  How to set up cross account access in AWS for variety of use-cases such as
  audit, security and compliance at scale.
author: benjamin
---

import { HowToGuideHeader } from "../../components/HowToGuideHeader"

<HowToGuideHeader/>

In this how-to guide we will walk through how to create a role in an external account that we want to AssumeRole into. In our example we will provide the new role in the external account with broad ReadOnly permissions (but you are free to change to whatever you want).

![AWS Schema](/images/blog/cross-account-access-aws-assumerole/scheme.png)


## Introduction to AWS IAM

If you are familiar AWS IAM feel free to skip this section. But here is a quick recap: AWS IAM is the underlying service that AWS services rely on for authorization. While there are many ways of getting valid AWS credentials (setting up AWS SSO) this blog is going to focus on how you can use those credentials once you have them. Before we dive in a few high level concepts that we need to clarify:

- Credentials used to sign API calls are always associated with a single **user** or **role** entity.
- A user is a single identity that is able to use long lived credentials while a role is a short lived identity that can be temporarily assumed by other entities.
- Credentials are authorized to perform only the actions that are defined in the policies that are associated with the role or user, and the scope of those permissions can never grant access to a resource in another account.

This means that a single policy cannot grant User in Account A access to a resource in Account B. A common scenario is for one user or role to access resources in Account B for use cases such as auditing and monitoring. In order for a user in Account A to access a resource in a different Account they must get access to credentials that are from Account B. The most common and secure way to enable User A in Account A to access resource B in Account B is done via cross account role assumption. This means that Account B defines a role, Role-B, that has a Trust Policy that says that entities from Account A are allowed to assume it, this effectively means that entities from Account A are able to get credentials that have the access of Role-B. So once that is in place User-A is able to assume Role-B and use the new credentials to interact with Resource-B.

## Walkthrough

In this tutorial we will show you how to do in the console (ClickOps) but feel free to automate it via your favorite IaC.

### Step 1 

Go to `iam→roles` and [click Create Role](https://us-east-1.console.aws.amazon.com/iamv2/home?region=us-east-1#/roles/create?step=selectEntities)

![](/images/blog/cross-account-access-aws-assumerole/step1.png)

### Step 2

For **_Trusted Entity,_** Choose **_AWS account_**, click **_Another AWS account_** and fill-in the account id you want to access **from** and click next

![](/images/blog/cross-account-access-aws-assumerole/step2.png)

### Step 3

Attach the **_Permission policy_** you want this role to have. In our case we will attach `ReadOnlyAccess` - mark the checkbox and click next.

![](/images/blog/cross-account-access-aws-assumerole/step3.png)

### Step 4

Last screen should look like the following and you should name the role the same name in all accounts you want to AssumeRole into. In this case we marked it as `CrossAccountReadOnlyRole`

![](/images/blog/cross-account-access-aws-assumerole/step4.png)

### Step 5

Repeat 1-4 for each account you want to access from.

### Step 6

Login into your account where you want to assume role from and go to iam→policies and click Create Policy. Paste the following policy in the `JSON` tab

```sql copy
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "VisualEditor0",
            "Effect": "Allow",
            "Action": "sts:AssumeRole",
            "Resource": "arn:aws:iam::*:role/CrossAccountReadOnlyRole"
        }
    ]
}
```

So it should look something like the following:

![](/images/blog/cross-account-access-aws-assumerole/step6.png)

### Step 7

Click Next, give it tags and then pick a name. We will use `CloudQueryCrossAccountReadOnlyAccess`

![](/images/blog/cross-account-access-aws-assumerole/step7.png)

### Step 8

Now you can go to user or role that you use in your main account and attach the policy:

![](/images/blog/cross-account-access-aws-assumerole/step8.png)

### Step 9

That’s it! You are done now you can use assume role via this role.
