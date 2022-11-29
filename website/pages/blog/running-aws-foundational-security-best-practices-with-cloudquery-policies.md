---
title: Running AWS Foundational Security Best Practices with CloudQuery Policies
tag: security
date: 2022/02/07
description: >-
  Automate, customize, codify and run AWS Foundational Security Best Practices
  with CloudQuery Policies.
author: mikeelsmore
---

import { BlogHeader } from "../../components/BlogHeader"
import { Callout } from 'nextra-theme-docs'

<BlogHeader/>

<Callout type="warning">
HCL policies were deprecated - see up-to-date policy documentation [here](https://www.cloudquery.io/docs/core-concepts/policies).
</Callout>

Back in mid-2020 AWS Security Hub released a new security standard called AWS Foundational Security Best Practices. This new standard sets security controls to detect when an AWS account or deployed resources don’t match up to the best practices set out by the AWS security experts. The complete standard can be found in the [AWS Security Hub documentation](https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards-fsbp.html).

As with any security guidelines, factors such as AWS environments, requirements, and capacity of your security team, will impact how you implement those guidelines.

The new AWS Foundational Security Best Practices CloudQuery policy gives you a powerful way to automate, customize, codify, and run your cloud security & compliance continuously with HCL and SQL.

The CloudQuery AWS Foundational Security Policy covers 200+ checks - you can review them on [GitHub](https://github.com/cloudquery-policies/aws/tree/main/foundational_security) or review them in the [GitHub](https://github.com/cloudquery/cq-provider-aws/tree/main/policies/foundational_security).

## Prerequisites

Please follow the [quickstart guide](/docs/quickstart) to install CloudQuery, and `fetch` your AWS configuration into a PostgreSQL database.

## Running

After fetching your AWS configuration into a PostgreSQL database, you can use SQL to check your cloud deployment for compliance!

For example, you can check for certificates that are going to expire soon and need to be renewed.

```sql
#https://github.com/cloudquery-policies/aws/blob/main/queries/acm/certificates_should_be_renewed.sql

SELECT arn

FROM aws_acm_certificates

WHERE not_after < NOW() AT TIME ZONE 'UTC' + INTERVAL '30' DAY;
```

You can also use the `cloudquery` command to run the entire AWS Foundational Security Best Practices policy pack. The policy is split into sections as sub-policies, so you can run either the entire policy, a sub-policy, or even one specific check.

```bash
# execute the AWS foundational-security-best-practices policy pack

cloudquery policy run aws//foundational_security

# execute the ACM section in AWS Foundational Security policy

cloudquery policy run aws//foundational_security/acm

# execute the S3 related section in AWS Foundational Security policy

cloudquery policy run aws//foundational_security/s3

# describe all available policies and sub-policies available for AWS on cloudquery

cloudquery policy describe aws

# execute the entire AWS policy pack, including other benchmarks.

cloudquery policy run aws
```

You can also output the results into a JSON and pass them to downstream processing for automated monitoring and alerting.

```bash
cloudquery policy run aws//foundational_security --output-dir=results
```

## Build your own and share!

Do you have a policy that you want to codify, or that you’ve been running with python or bash scripts? You are welcome to try codifying it with CloudQuery Policies (See our [github](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/aws/policies_v1) and [docs](/docs/core-concepts/policies) for how to develop one). Feel free to drop on [discord](https://www.cloudquery.io/discord) or [GitHub](https://github.com/cloudquery/cloudquery/issues) to get any help, and we will share your policy on [CloudQuery Hub](https://www.cloudquery.io/).
