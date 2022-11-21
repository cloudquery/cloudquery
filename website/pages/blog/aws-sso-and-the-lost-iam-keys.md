---
title: AWS SSO And The Lost IAM Keys
tag: security
date: 2021/08/20
description: AWS SSO and IAM Security Best Practices
author: yevgenypats
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>


[AWS SSO](https://docs.aws.amazon.com/singlesignon/latest/userguide/what-is.html) is one of the best and most popular ways to centrally manage access of users/developers to AWS Accounts, especially when combined with [AWS Organizations](https://aws.amazon.com/organizations/) for multi-account access and management.

AWS SSO Usually integrates with an IDP (like, [Okta](https://aws.amazon.com/blogs/aws/single-sign-on-between-okta-universal-directory-and-aws/), [G Suite](https://aws.amazon.com/blogs/security/how-to-use-g-suite-as-external-identity-provider-aws-sso/), [AzureAD](https://docs.microsoft.com/en-us/azure/active-directory/saas-apps/amazon-web-service-tutorial)). This approach has many benefits, for instance: Only users that are in your Okta, G Suite, AzureAD directory can access the AWS accounts. Authentication and MFA is managed centrally at IDP level. Any user that leaves the organisation is also automatically revoked access from AWS.


## The lost IAM Keys

One issue that remains from the access and security perspective is the handling of IAM Keys. IAM Keys can be created by users in the IAM Console for `dev/test/prod` purposes.

When managing IAM Keys, we have to account for couple of scenarios:

- IAM Key rotation - It's a best practice to rotate IAM keys every ~90 days.
- When a user leaves the organisation - the main user account is automatically deactivated at the IDP level, but the IAM Keys created by this user should be rotated or deleted (if they're not being used anywhere else).

## Tagging Policy

To solve these issues, we first need to be able to locate/correlate IAM keys and their creators. The best way to do this in AWS (and usually in other cloud providers) is via a tagging policy. You can add a tag for each IAM key `creator=user@your_domain.com` . You can also enforce this policy via [AWS Policy](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_tag-policies.html).

## The Code Approach

Currently the standard way to solve the "Lost IAM Keys" issue would be to write a small/medium script that will do the following:

- Extract all the IAM Keys, tags and last rotated timestamp.
- Filter all keys by last-rotated-timestamp older than 90 days to find keys that need to be rotated, and, via tag, their creator (via tag) who needs to rotate them.
- Extract all users currently active in the IDP Directory.
- Find all keys with emails that are not active in the IDP, or, for every deactivated user, find all related keys.

This approach will work, but will require writing quite a bit of code, testing and maintaining. It will also require integrating with at least two APIs/SDKs (AWS/Okta/G Suite).

## The CloudQuery Approach

The other approach (and the reason why we started [CloudQuery](https://github.com/cloudquery/cloudquery)) is the belief the following issue should be solved in two simple steps:

- Cloud Native ETL - That takes care of connecting to the various APIs, extracting the configuration and meta-data, transforming/normalising the data and loading it into a relational-database.
- Once you have all the up-to-date data. This issue can be solved with two SQL queries.

The first step in our case is taken care of by CloudQuery ([GitHub](https://github.com/cloudquery/cloudquery)) - all you have to do is follow our [quickstart guide](/docs/quickstart).

Once CloudQuery loads the data into PostgreSQL, you can run the following queries to answer the above questions:

```sql
/* All keys with not used in the last 90 days */
SELECT DISTINCT aws_iam_users.account_id, aws_iam_users.arn, MAX(password_last_used), aws_iam_users.user_name, access_key_id, MAX(last_used) FROM aws_iam_users
    JOIN aws_iam_user_access_keys on aws_iam_users.arn = aws_iam_user_access_keys.user_arn
WHERE (password_last_used < (now() - '90 days'::interval) OR
        (last_used < (now() - '90 days'::interval))) 
GROUP BY (aws_iam_users.account_id, aws_iam_users.arn, aws_iam_users.user_name, access_key_id);

/* This will return all IAM keys with emails that don't exist in the Okta directory. */
select tags, arn from aws_iam_users
	LEFT JOIN okta_users on aws_iam_users.tags->>'email' = okta_users.profile->>'email'
where okta_users.profile->>'email' is NULL;
```

## Continuous Monitoring & Alerting

You can run CQ periodically, either from a local machine or a server/lambda, and create alerts using the above queries and [CQ Policies](https://www.cloudquery.io/blog/announcing-cloudquery-policies).
