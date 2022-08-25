---
title: Fetching From Multiple AWS Accounts With AssumeRole
---

import { Callout } from 'nextra-theme-docs'

# Fetching From Multiple AWS Accounts With AssumeRole

If you need to `fetch` from multiple accounts, you can configure CloudQuery to `AssumeRole` into them.
With this setup, you just need to supply a single set of credentials to CloudQuery (for the user
that can AssumeRole into the accounts you'd like to fetch from). This tutorial will walk you through all the
steps required to set this up (You can also take a look at the [AWS tutorial for AssumeRole](https://docs.aws.amazon.com/IAM/latest/UserGuide/tutorial_cross-account-with-roles.html)).

## Prerequisites

- Basic familiarity with CloudQuery (see our [Getting started guide](../getting-started/getting-started-with-aws)).
- Basic familiarity with AWS CLI authentication (See the [AWS guide](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html) or our own [authentication section](../getting-started/getting-started-with-aws#authenticate-with-aws)).
- One (or more) AWS accounts.
- A postgreSQL database.

## Definitions

For the example in this tutorial, we'll have 3 accounts:

- `SourceAccount`: The account CloudQuery will run from. CloudQuery won't `fetch` resources from this account.
- `TargetAccountA`: The first account CloudQuery should `fetch` from.
- `TargetAccountB`: The second account CloudQuery should `fetch` from.

<Callout type="warning">

AWS best practices [recommend](https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html#delegate-using-roles) using
IAM roles to delegate permissions when possible. For the sake of simplicity, in this tutorial we assume
CloudQuery runs as `SourceAccountUser` in the source-account - but we recommend you take a look at
[this guide](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-role.html) to learn about how to
setup CloudQuery to use a role instead of a user.

</Callout>

<Callout type="info">

The example in this tutorial will work just as well if `SourceAccount` and `TargetAccountA` are the same account.
i.e. you don't need a separate account for `SourceAccount`.

</Callout>

## Create roles in the target accounts

You will need to create roles in `TargetAccountA` account and `TargetAccountB` account that CloudQuery can `AssumeRole` into.
These roles will need to have a "trust policy" that allows `SourceAccount` to `AssumeRole` into them.

First, create a file named `trust_policy.json` (replace `<SourceAccountUserArn>` with the user-arn CloudQuery will run from):

```json title="trust_policy.json"
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "AWS": "<SourceAccountUserArn>"
            },
            "Action": "sts:AssumeRole",
            "Condition": {}
        }
    ]
}
```

You will have to run the `aws cli` commands here twice: first for `TargetAccountA` and then for `TargetAccountB`.
i.e., between running these commands, you will need to change your `AWS_ACCESS_KEY_*` environment variables,
`AWS_PROFILE` environment variable, or similar.

```bash title="AWS CLI commands for TargetAccountA"
aws iam create-role --role-name CloudQueryFetchRole --assume-role-policy-document file://trust_policy.json
```

```bash title="AWS CLI commands for TargetAccountB"
aws iam create-role --role-name CloudQueryFetchRole --assume-role-policy-document file://trust_policy.json
```

## Grant permissions to the roles in the target accounts

The roles in your target account will need to have permissions to read your cloud configuration.

<Callout type="info">

CloudQuery requires only _read_ permissions (we will never make any changes to your cloud setup).
Attaching the `ReadOnlyAccess` policy to the user/role CloudQuery is running as should work for the most part,
but you can fine-tune it even more to have read-only access for the specific set of resources that you want
CloudQuery to fetch.
[See also this blog post](https://alestic.com/2015/10/aws-iam-readonly-too-permissive/).

</Callout>

```bash title="AWS CLI commands for TargetAccountA"
aws iam attach-role-policy --role-name CloudQueryFetchRole --policy-arn arn:aws:iam::aws:policy/ReadOnlyAccess
```

```bash title="AWS CLI commands for TargetAccountB"
aws iam attach-role-policy --role-name CloudQueryFetchRole --policy-arn arn:aws:iam::aws:policy/ReadOnlyAccess
```

## Grant permissions for the Source User to assume-role into the target roles

The user cloudquery will run as in `SourceAccount` will need permissions to assume role into the roles we
created in the target account.

First, create an `allow_assume_role_into_target_accounts.json` file (replace `<TargetAccountAId` and `TargetAccountBId`):

```json title="allow_assume_role_into_target_accounts.json"
{
  "Version": "2012-10-17",
  "Statement": {
    "Effect": "Allow",
    "Action": "sts:AssumeRole",
    "Resource": [
      "arn:aws:iam::<TargetAccountAId>:role/CloudQueryFetchRole",
      "arn:aws:iam::<TargetAccountBId>:role/CloudQueryFetchRole"
    ]
  }
}
```

Authenticate your AWS-CLI to access `SourceAccount`, and then run:

```bash title="AWS CLI Commands for SourceAccount"
aws iam create-policy --policy-name AllowAssumeRoleIntoTargetAccounts --policy-document file://allow_assume_role_into_target_accounts.json
```

Next, attach the `AllowAssumeRoleIntoTargetAccounts` policy to the user CloudQuery will run as (replace `<SourceAccountUserName>` and `<SourceAccountId>`):

```bash title="AWS CLI Commands for SourceAccount"
aws iam attach-user-policy --user-name <SourceAccountUserName> --policy-arn "arn:aws:iam::<SourceAccountId>:policy/AllowAssumeRoleIntoTargetAccounts"
```

## Configure CloudQuery

Now that all AWS permissions are ready, you need to configure CloudQuery to assume-role into the target accounts.

Your `cloudquery.yml` file should look similar to the following (call `cloudquery init aws` if you don't have one yet).
Remember to replace `<TargetAccountAId>`, `<TargetAccountBId>`, and `<PostgreSQL_DSN>` with the appropriate values:

```yaml title="cloudquery.yml"
cloudquery:
  ...

  providers:
    - name: "aws"
      source: "cloudquery/aws"
      version: "latest"

  connection:
    type: postgres
    username: postgres
    password: pass
    host: localhost
    port: 5432
    database: postgres
    sslmode: disable

providers:
  - name: "aws"
    ...
    configuration:
      accounts:
        - account_name: "target_account_a"
          role_arn: "arn:aws:iam::<TargetAccountAId>:role/CloudQueryFetchRole"
        - account_name: "target_account_b"
          role_arn: "arn:aws:iam::<TargetAccountBId>:role/CloudQueryFetchRole"

```

<Callout type="warning">

Warning

`accounts` is a confusing name - the block should really be named `account`. Each `accounts` represents a **single** account
to assume-role into, and you can have as many `accounts` blocks as you'd like.

</Callout>

<Callout type="info">

When you run without any `accounts` blocks in your `configuration`, CloudQuery will use the default credential
tool-chain to locate the credentials that will be used in the `fetch`.
However, if there is at least one `accounts` block in your `configuration`, cloudquery **will not** fetch from the
account you are currently authenticated as - only from the accounts specified in
your `accounts`.

</Callout>

## Run `cloudquery fetch`

You are now almost ready to fetch. After you
[Authenticate CloudQuery with credentials for SourceAccountUser](../getting-started/getting-started-with-aws#authenticate-with-aws),

you can run:

```bash title="Authenticated as SourceAccountUser"
cloudquery fetch
```
