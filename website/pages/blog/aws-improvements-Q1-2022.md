---
title: Q1 Improvements for the AWS Provider
tag: recap
date: 2022/03/30
description: >-
  Highlight the improvements to the CloudQuery AWS Provider that have landed in
  the first quarter of 2022.
author: benjamin
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

## Q1 2022

We want to highlight some of the features that we have released that will improve the usability for CloudQuery users looking to catalog their AWS assets.

### TL;DR 📕

- AWS Org Support
- Better Support for the credentials in the location you choose
- Added ARN to nearly all tables
- Added 7 more resources and released 50+ bug fixes

### What is CloudQuery ?

CloudQuery is the open-source cloud asset inventory powered by SQL, enabling you to [catalog, audit, and evaluate the configurations](/docs/core-concepts/policies).

CloudQuery key use-cases and features:

- **Search**: Use standard SQL to find any asset based on any configuration or relation to other assets.
- **Visualize**: Connect CloudQuery standard PostgreSQL database to your favorite BI/Visualization tool such as Grafana, QuickSight, etc.
- **Policy-as-Code**: Codify your security & compliance rules with SQL as the query engine.

### What’s New in the AWS Provider?

- **Org Support**:
  You used to have to manually create (and maintain) an `account` block for each account in your entire organization. This was difficult for larger organizations where accounts are constantly being added and removed. We now integrate directly with AWS Organizations to find and configure all accounts in your Organization or in specific Organizational Units. Here is an example of a configuration for using the new organizations feature:

  ```yaml
  kind: source
  spec:
    name: aws-0
    registry: github
    path: cloudquery/aws
    version: <LatestVersion>
    tables: ['*']
    # skip_tables: []
    destinations: ["postgresql"]
    spec:
      aws_debug: false
      org:
        admin_account:
          local_profile: "<NAMED_PROFILE>"
        member_role_name: OrganizationAccountAccessRole
      regions:
        - '*'
  ```
  
  for more information feel free to check out the documentation [here](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/aws/docs/configuration.md)

- **Credentials**:

  - On a per account basis you can reference local credentials in your `~/.aws/config` or `~/.aws/credentials` files. Prior to this all accounts sourced their credentials from default credential chain.
    In the example below, `account1` utilizes the default credential provider chain while `account2` sources its credentials from the shared credentials file

  ```yaml
  kind: source
  spec:
    name: aws-0
    registry: github
    path: cloudquery/aws
    version: <LatestVersion>
    tables: ['*']
    # skip_tables: []
    destinations: ["postgresql"]
    spec:
      accounts:
        - id: "account1"
          role_arn: "<ARN_OF_ROLE_IN_account1>"
        - id: "account2"
          local_profile: "<NAMED_PROFILE>"
          role_arn: "<ARN_OF_ROLE_IN_account2>"
          session_name: "NAMED_OF_SESSION"
  ```

  - On any account users can now specify a session name that CloudQuery will use when assuming a role in an account. This is important because some organizations enforce naming restrictions of their IAM Role sessions for better audibility.

- **ARN Field in All Tables**:

  - We heard customer feedback where you wanted to be able to identify resources by their full ARN so went through all of the resources that we support and made sure that we included a column containing the ARN of the resource! (more fun to come with this in the near future)

- **New Resources**:

  - Our team has been hard at work adding support for new resources including:
    - Access Analyzers
    - CloudFormation Stacks
    - EC2 Instance Status (added by the community!)
    - EC2 Security Groups
    - EFS filesystem backup policy
    - S3 Bucket Ownership
    - AWS Workspaces

  If we are not supporting a resource that you need please reach out to us on [GitHub](https://github.com/cloudquery/cq-provider-aws), [Discord](https://www.cloudquery.io/discord) or [Twitter](https://twitter.com/cloudqueryio)

What's next?
As always more resources, providers, policies, improved stability, and support for storing policy results in your database. If you would like to influence our roadmap feel free to open an issue on our [GitHub](https://github.com/cloudquery/cloudquery) or [Discord](https://www.cloudquery.io/discord)!
