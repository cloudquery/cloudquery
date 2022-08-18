---
title: Continuous AWS IAM Security Best Practices
tag: security
date: 2021/03/15
description: >-
  Walkthrough on how to automate, validate and monitor AWS IAM Security best
  practices with CloudQuery
author: yevgenypats
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

There are some great guides on the internet for AWS Security best practices (both official and unofficial).
However, one of the challenges we saw with those guides is that they tell you what the end goal is,
but they usually leave it up to the user on how to implement it (at scale),
let alone how to continuously monitor those best practices to make sure all your hard work doesn’t go through the window.

In this blog post we will go through the official security IAM best practices, and we'll show how to validate and monitor them using SQL statements with CloudQuery.

## Account Setup

You can run all the following commands on a single AWS account. You can also run them on multiple accounts in parallel, using an account that can assume-role into all your other relevant accounts.

## CloudQuery Setup

To be able to run the following tutorial you need to install and configure CloudQuery:

### Install

You can download the pre-compiled binary from releases, or using CLI:

```powershell
export OS=Darwin # Possible values: Linux,Windows,Darwin
curl -L https://github.com/cloudquery/cloudquery/releases/latest/download/cloudquery_${OS}_x86_64 -o cloudquery
chmod a+x cloudquery
```

For mac you can use `homebrew`:

```powershell
brew install cloudquery/tap/cloudquery
# After initial install you can upgrade the version via:
brew upgrade cloudquery
```

### Run

```powershell
docker run -p 5432:5432 -e POSTGRES_PASSWORD=pass -d postgres
cloudquery init aws
cloudquery fetch --dsn "host=localhost user=postgres password=pass DB.name=postgres port=5432"
```

To use this on multiple account see [README.md](https://github.com/cloudquery/cloudquery#aws)

## Security Best Practices in IAM

In this section, we will share snippets of SQL statements, most of which you can find in our [aws//cis_v1.2.0
policy](https://github.com/cloudquery-policies/aws/tree/main/cis_v1.2.0).
You can also use this as a reference to create your own policies, which can be customized to your needs and the security policies of your organization.

### Lock away your root account

One best practice in AWS is to enable MFA (on your root account) and create an IAM admin user to handle all your day-to-day work immediately after you open an account.
List root accounts that were accessed in the last 30 days:

```sql
SELECT account_id, arn, password_last_used, user_name FROM aws_iam_users
WHERE user_name = '<root_account>' AND password_last_used > (now() - '30 days'::interval)
```

Also, it is strongly advised not to create any access keys for the root account.
To list root accounts with access keys:

```sql
select * from aws_iam_users
   JOIN aws_iam_user_access_keys keys on aws_iam_users.id = keys.user_id
WHERE user_name = '<root>'
```

### Use groups to assign permissions to IAM

Instead of attaching IAM policies to users it is advised to attach IAM policies to groups and roles, and then to assign the groups and roles to users.
To list accounts with IAM policies attached directly:

```sql
SELECT aws_iam_users.account_id, arn, user_name FROM aws_iam_users
JOIN aws_iam_user_attached_policies policies on aws_iam_users.id = policies.user_id
```

### Grant least privilege

Creating and reviewing IAM policies with least privilege in mind is a big topic on its own.
We will cover this in a follow-up post!

### Configure a strong password policy for your users

If you allow users to change their own password, then creating a password policy is highly recommended.
This snippet is taken from AWS CIS benchmark. It provides a good example of a strong password policy and offers a way to check that it is enabled in all accounts.

```sql
- name: "AWS CIS 1.5  Ensure IAM password policy requires at least one uppercase letter"
 query: >
   SELECT account_id, require_uppercase_characters FROM aws_iam_password_policies
    WHERE require_uppercase_characters = FALSE
- name: "AWS CIS 1.6  Ensure IAM password policy requires at least one lowercase letter"
 query: >
   SELECT account_id, require_lowercase_characters FROM aws_iam_password_policies
    WHERE require_lowercase_characters = FALSE
- name: "AWS CIS 1.7  Ensure IAM password policy requires at least one symbol"
 query: >
   SELECT account_id, require_symbols FROM aws_iam_password_policies
    WHERE require_symbols = FALSE
- name: "AWS CIS 1.8  Ensure IAM password policy requires at least one number"
 query: >
   SELECT account_id, require_numbers FROM aws_iam_password_policies
    WHERE require_numbers = FALSE
- name: "AWS CIS 1.9 Ensure IAM password policy requires minimum length of 14 or greater"
 query: >
   SELECT account_id, minimum_password_length FROM aws_iam_password_policies
    WHERE minimum_password_length < 14
- name: "AWS CIS 1.10 Ensure IAM password policy prevents password reuse"
 query: >
   SELECT account_id, password_reuse_prevention FROM aws_iam_password_policies
    WHERE password_reuse_prevention is NULL or password_reuse_prevention > 24
- name: "AWS CIS 1.11 Ensure IAM password policy expires passwords within 90 days or less"
 query: >
   SELECT account_id, max_password_age FROM aws_iam_password_policies
    WHERE max_password_age is NULL or max_password_age < 90
```

### Enable MFA

AWS recommends that MFA be enabled for the root account, as well as any other privileged IAM users.
To list all accounts with MFA disabled:

```sql
SELECT account_id, arn, password_last_used, user_name, mfa_active FROM aws_iam_users WHERE NOT mfa_active
```

For example, if you only want to alert on specific accounts, you can add a list of accounts to the SQL statements.
Or you can go even further by joining on the permissions table and alerting on accounts with specific permissions that do not have MFA enabled.

### Rotate access keys

To list all access keys that weren’t rotated in the last 90 days you can run the following:

```sql
SELECT account_id, arn, password_last_used, user_name, access_key_id, last_used, last_rotated FROM aws_iam_users
 JOIN aws_iam_user_access_keys on aws_iam_users.id = aws_iam_user_access_keys.user_id
WHERE last_rotated < (now() - '90 days'::interval)
```

### Remove unnecessary credentials

To list all IAM access keys that haven’t been used for more than 90 days:

```sql
SELECT account_id, arn, password_last_used, user_name, access_key_id, last_used FROM aws_iam_users
 JOIN aws_iam_user_access_keys on aws_iam_users.id = aws_iam_user_access_keys.user_id
WHERE (password_enabled AND password_last_used < (now() - '90 days'::interval) OR
      (last_used < (now() - '90 days'::interval)))
```

### CloudQuery Policies

To be able to run all SQL checks together, you can use CloudQuery policies.
These are YAML files with a simple structure:

```sql
views:
name: “name_of_view”
query: >
   CREATE VIEW name_of_view AS
    SELECT  …… JOIN … WHERE …
queries:
name: “name of check”
query: >
	SELECT … FROM …
name: “name of check”
invert: true
query: >
	SELECT … FROM ...
```

There are two sections:

- **Views**: you can use this feature to create new views, which you can reuse later instead of having to duplicate joins across queries!
- **Queries**: Here, you can write all your queries and checks. By default, if a query returns results, the check will fail with the violating results printed in the UI.

You can run the policy with the following command

```powershell
cloudquery query --path policy.yml
```

It’s also possible to output the results in JSON format via --output flag, and to ingest it into a centralized logging system.
You can view our past guest blog on how to deploy CloudQuery on AWS lambda to periodically fetch and query data [here](https://www.cloudquery.io/docs/deployment/overview).

**Stay tuned for more blogs on cloud security, compliance and other cool stuff!**
