---
title: How to run AWS CIS Benchmark with CloudQuery
tag: security
date: 2021/02/15
description: >-
  Learn how to run AWS CIS benchmark with CloudQuery using customizable SQL
  statements.
author: yevgenypats
---

import { BlogHeader } from "../../components/BlogHeader"
import { Callout } from 'nextra-theme-docs'

<BlogHeader/>

<Callout type="warning">
HCL policies were deprecated - see up-to-date policy documentation [here](https://www.cloudquery.io/docs/core-concepts/policies).
</Callout>

## Intro

**[The AWS CIS benchmarks](https://www.cisecurity.org/benchmark/amazon_web_services/)** are a set of well-known configuration best-practices that helps companies improve their AWS security posture and comply with various compliance framework like (SOC2, etc.).

The guide is in PDF, some instructions require going through the UI and others require running long set of commands. Doing this manually is a tedious and error prone work not talking about doing this continuously.

There are some open-source tools like [prowler](https://github.com/toniblyx/prowler) that runs all those commands in one large shell script. The dis-advantage with this kind of approach is that it is hard to customized and for example exclude resource that you know they are not compliant to get a clean report.

This blog will show you how to run AWS CIS benchmark with **[CloudQuery](https://github.com/cloudquery/cloudquery)** using out-of-the-box SQL statements that you can customize to your environment.

## Running CloudQuery

### Downloading

CloudQuery is an open-source framework that transforms your cloud infrastructure into SQL database for easy monitoring, governance and security. It's written in Go so it's just a single Binary!

You can download the pre-compiled binary from [releases](https://github.com/cloudquery/cloudquery/releases), or using CLI:

```powershell
export OS=darwin # Possible values: linux,windows,darwin
curl -L https://versions.cloudquery.io/latest/v2/cloudquery_${OS}_amd64 -o cloudquery
chmod a+x cloudquery
```

For mac you can use `homebrew`:

```powershell
brew install cloudquery/tap/cloudquery
# After initial install you can upgrade the version via:
brew upgrade cloudquery
```

### Choosing database

CloudQuery currently supports two types of databases: PostgreSQL & TimescaleDB (for historical snapshots). In this post we will use the default SQLite which is great for local development and testing (see [here](/docs) on how to use others).

### Authenticating with AWS

CloudQuery uses AWS SDK under-the-hood so authentication works the same [way](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html). To sum up you can use the following environment variables or files:

- AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY
- ~/.aws/credentials created via aws configure
- AWS_PROFILE

### Fetching the data

To fetch the data you first need to specify which resources you want to fetch. We will use the default that specifies
all the resources that are currently supported (you can customize it and comment out things you don't want).

To generate the default `config.yml` run `cloudquery init aws` which will output a `config.yml` ready to use in the current directory.

Now the money time! run **`cloudquery fetch`** . This will fetch all specified resources in all regions and in all accessible accounts (or specific accounts if specified).

### Running AWS CIS benchmark test

We've created CloudQuery [policy pack](https://github.com/cloudquery-policies/aws/blob/main/cis_v1.2.0/policy.hcl)
That runs all CIS checks with SQL statements so you know SQL you can easily customized it to your needs.

Now run `cloudquery policy run aws//cis_v1.2.0` to see results!

Here is a snippet of the report:

```powershell
✓ policy "cloudquery-policies-aws" -  evaluating -                0s   Finished Queries: 85/85

📋 cloudquery-policies-aws Results:

⚠️ Policy finished with warnings

	✓   1.1  AWS CIS 1.1 Avoid the use of 'root' account. Show used in last 30 days (Scored)                                               passed

	✓   1.2  AWS CIS 1.2 Ensure MFA is enabled for all IAM users that have a console password (Scored)                                     passed

	✓   1.3  AWS CIS 1.3 Ensure credentials unused for 90 days or greater are disabled (Scored)                                            passed

	✓   1.4  AWS CIS 1.4 Ensure access keys are rotated every 90 days or less                                                              passed

	✓   1.5  AWS CIS 1.5  Ensure IAM password policy requires at least one uppercase letter                                                passed

	✓   1.6  AWS CIS 1.6  Ensure IAM password policy requires at least one lowercase letter                                                passed

..............
Finished policies run...
```

You can also specify `--output` if you want also to store the results in JSON format so you can forward it to some logging system.

### Running CloudQuery continuously

In the next blog we will show how to set-up CloudQuery in a lambda function to run periodically so you can continuously monitor your rules.
