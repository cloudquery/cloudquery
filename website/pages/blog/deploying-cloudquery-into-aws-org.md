---
title: Deploying CloudQuery into an AWS Organization
tag: tutorial
date: 2022/07/13
description: >-
  Tutorial for setting up AWS IAM permissions in an AWS Organization so that
  CloudQuery can fetch all resources in all accounts
author: benjamin
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

CloudQuery makes fetching resources from your entire organization simple as long as you have the proper IAM trust relationships and permissions setup. In this blog post we will walk through one way of setting up these permissions so that you can have access to all of your configuration data in a single queryable database.

## General Architecture

We will be deploying a single CloudFormation template in an admin account and then relying on CloudFormation StackSets to propagate the configurations to all of the member accounts. In the admin account we will create an IAM role that is able to list all of the accounts in the organization as well as assume a role in the member accounts. In the member accounts we will be deploying a single IAM role that has a trust policy to only allow the role in the admin account to assume it. The permissions for the IAM role in the member account are locked down so that the role can only access metadata about the configuration and never has access to your code or data.

![](/images/blog/deploying-cloudquery-into-aws-org/image0.png)

## Prerequisites:

- CloudQuery binary installed and setup with a Postgresql database

- AWS CLI V2 installed and configured

- Admin access to the Root Account of your AWS Organization or an Account that is a Delegated Admin

## Walkthrough

### Step 1: Clone Solution Repository:

Clone the [CloudQuery IAM-For-Orgs](https://github.com/cloudquery/iam-for-aws-orgs) repository.

```bash
git clone https://github.com/cloudquery/iam-for-aws-orgs.git
cd iam-for-aws-orgs
```

### Step 2: Deploy IAM Resources

Prior to using the AWS CLI to execute the following command but make sure to replace `<ROOT_ORG_ID>` with your Organizational Unit (OU) of the root (if you want to deploy to your entire organization). Or a comma separated list of OUs if you want to deploy it only to a specific set of accounts:

```bash
aws cloudformation create-stack \
  --stack-name CloudQueryOrg-Deploy \
  --template-body file://./template.yml \
  --capabilities CAPABILITY_NAMED_IAM \
  --parameters ParameterKey=OrganizationUnitList,ParameterValue=<ROOT_ORG_ID>
```

You can monitor the state of CloudFormation deployment by running this command:

```bash
aws cloudformation describe-stacks \
  --stack-name CloudQueryOrg-Deploy \
  --query 'Stacks[].StackStatus | [0]'
```

### Step 3: Ensure you can Assume role

Once the CloudFormation template has finished you will want to test out to make sure you can assume the role that was created in the Admin account. First you will get the ARN of that IAM role by running this command:

```bash
aws cloudformation describe-stacks \
  --stack-name CloudQueryOrg-Deploy \
  --query "Stacks[0].Outputs[?OutputKey=='AdminRoleArn'].OutputValue | [0]"
```

where the output should be a string like this:

```bash
"arn:aws:iam::<REDACTED>:role/cloudquery-ro"
```

Once you have the ARN of the role that was created you can test out assuming it by running this CLI command:

```
aws sts assume-role \
  --role-arn arn:aws:iam::<REDACTED>:role/cloudquery-ro \
  --role-session-name cloudquery-test
```

The output should look something like this:

```bash
{
    "Credentials": {
        "AccessKeyId": "ASIA...",
        "SecretAccessKey": "...",
        "SessionToken": "...",
        "Expiration": "2022-07-12T17:38:11+00:00"
    },
    "AssumedRoleUser": {
        "AssumedRoleId": "...",
        "Arn": "arn:aws:sts::<REDACTED>:assumed-role/cloudquery-ro/cloudquery-test"
    }
}
```

### Step 4: Update `cloudquery.yml`

Now that you have properly configured the IAM role in the Admin account you can take the 2 outputs from the CloudFormation stack and plug them into your CloudQuery configuration file

```bash
aws cloudformation describe-stacks \
  --stack-name CloudQueryOrg-Deploy \
  --query "Stacks[].Outputs"
```

Where the output should be in this form:

```bash
[
    [
        {
            "OutputKey": "MemberRoleName",
            "OutputValue": "cloudquery-ro"
        },
        {
            "OutputKey": "AdminRoleArn",
            "OutputValue": "arn:aws:iam::<REDACTED>:role/cloudquery-ro"
        }
    ]
]
```

Here is an example configuration that makes use of the resources you just created

```yaml
kind: source
spec:
  name: aws-0
  registry: github
  path: cloudquery/aws
  version: <LatestVersion>
  tables: ['*']
  destinations: ["postgresql"]
  spec:
    aws_debug: false
    org:
      admin_account:
        role_arn: <AdminRoleArn>
      member_role_name: <MemberRoleName>
    regions:
      - '*'      
```

### Step 5: Run Fetch

Now you are all set to be able to execute a fetch and grab resources from your whole AWS Organization

```bash
cloudquery fetch
```

## Summary

In this walkthrough we deployed a CloudFormation template that created the appropriate AWS IAM roles that CloudQuery needs in order to fetch resources from all accounts in an AWS Organization. Then you used the outputs from that template to configure CloudQuery to fetch resources from the member account.

After you have this working on your local machine you can check out our [Terraform module](https://github.com/cloudquery/terraform-aws-cloudquery) for simplifying the deployment of CloudQuery into a production ready environment!

If you have any questions, comments or feedback about this walkthrough feel free to reach out on [discord](https://www.cloudquery.io/discord) or [twitter](https://twitter.com/cloudqueryio)
