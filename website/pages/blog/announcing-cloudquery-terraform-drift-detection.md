---
title: Announcing CloudQuery Terraform Drift Detection
tag: announcement
date: 2021/11/16
description: Use CloudQuery to detect IaC drift locally and in Continuous Integration (CI)
author: yevgenypats
---

import { BlogHeader } from "../../components/BlogHeader"
import { Callout } from 'nextra-theme-docs'

<BlogHeader/>

<Callout type="warning">
This feature was deprecated, [see blog post](https://www.cloudquery.io/blog/terraform-drift-deprecation).
</Callout>

---

We are excited to announce the first release of CloudQuery Terraform drift detection!

In this blog we’ll talk a bit about what is infrastructure drift and why we decided to release this feature on top of our open-source cloud asset inventory platform.

## What is (infrastructure) drift?

Infrastructure drift is when your infrastructure state is not as described in your IaC (Infrastructure-as-code), i.e Terraform, Pulumi, CloudFormation.

Infrastructure drift can be split into two types of drifts.

1. **Resource Managed by IaC:** This kind of drift can be detected by the IaC tool. For example running “terraform plan” will show you the current drift of resources that are described in the terraform files.
2. **Resource not Managed by IaC:** This kind of drift **cannot** be detected by IaC tools because IaC tools are not aware of resources that are created outside of their workflows and are not defined in their files. Here is a [good](https://www.hashicorp.com/blog/detecting-and-managing-drift-with-terraform) article about “detecting and managing drift with Terraform” from HashiCorp, which explains the same thing.

In an ideal world you want everything to be managed by your IaC but unfortunately it’s not always the case and even if it is, drift can still be created without the IaC tools being able to detect them and you will have hanging resources forever (I’ll give an example).

- **Manual changes:** An obvious reason that will cause drifts and should be avoided. But mistakes happen, emergencies happen when you need to quickly fix it up and you remind yourself later to fix it also in terraform. So of-course it shouldn’t be a best-practice but also it’s unavoidable and this is the nature of our work.
- **Bugs/Mistake in IaC workflows:** For example you deleted the resource from terraform and you either forgot to run terraform again or your CI failed, the state got corrupted. Now you will have a drift which cannot be detected by your IaC as it’s not managed already by your IaC.

## Why it is important?

Drifting from the second type (**Resource not Managed by IaC**) which eventually will happen can cause a number of issues, cost is an obvious one but it can also cause security issues, outages and so on.

## Why CloudQuery

Just a quick recap for those who are unfamiliar with [CloudQuery](https://github.com/cloudquery/cloudquery): it is the open-source cloud asset inventory powered by SQL.

To solve the drift-detection problem we described in an efficient way you either need a queryable asset inventory or you need to write a lot of code to extract all that information. As CloudQuery is exactly that, we decided to implement this use-case as a module on top of CloudQuery (we are excited to see what you can build on top of CloudQuery :) ).

CloudQuery is not the first open-source tool to try and solve this problem, credit where credit is due as [`driftctl`](https://github.com/cloudskiff/driftctl) has some good prior art in the space that did a good job. When operating it we found limitations in how we needed to run drift detection, at this point we decided to run the drift detection on our central cloud asset inventory.

### How `driftctl` works

For every CI job that runs your IAC (we will use terraform for the sake of this example), you create an additional step which runs `driftctl scan`. Following is what happening when you run `driftctl` scan

- **Fetch Step**: Fetch all configuration from your cloud account into memory. Two main issues that we stumbled upon, specifically on medium-big cloud accounts:
  - **memory consumption:** as **all** the cloud state is stored in memory it can get pretty big and in some Continuous Integration (CI) environments which will just kill the job as you can’t always control or increase the memory limit.
  - **Execution Time:** The execution time doesn’t correlate to the terraform plan (which can be short, especially if nothing changed) step but to the size of the account. This can increase the execution time by a magnitude (x10000) of times which can be unacceptable for some teams/companies.
  - **Throttling Errors:** One of the challenges in Cloud ETL that we know first hand is Cloud Providers API Throttling errors. In larger teams where you might have terraform plan/apply running in different CI jobs in parallel, adding a tool similar to `driftctl` and fetching all the information in parallel in big accounts will start causing throttling errors, impacting production workflows and failing CI jobs.
- **Drift Detection:** `driftctl` compares the terraform state JSON with the in-memory cloud state.
  - These steps don’t have any serious limitations apart from the need to write Go code to solve a data problem, but that will work and potentially more of a subjective thing.

### How CloudQuery works

CloudQuery's basic capability is extracting all your asset configurations from all your cloud accounts, transforming, normalizing them and loading them into PostgreSQL.

Once you have all your cloud asset configuration loaded into PostgreSQL, the drift detection problem is turned into a data problem and this is how we solved it, using SQL.

The **Fetch** step is not running in the CI:

- **Fetch Step:** CloudQuery fetch is not running in the CI but running centrally where you deployed in your cloud.
  - **Memory:** Doesn’t have memory limitations as it’s uses an external database
  - **Execution Time:** Doesn’t impact the CI step as it happens centrally on your schedule.
  - **Throttling Errors:** This is dealt with at the core of cloudquery and it doesn’t hit throttle errors even if you have 1000 CI jobs running in parallel as they just talk to the PostgreSQL.

Following is what will happen if you run cloudquery scan in the CI:

- **Drift Detection:** CloudQuery will talk to your centrally deployed CloudQuery database and will run the drift detection using SQL.
  - This is fast (if you believe in PostgreSQL of-course)!
  - If there are any drifts that are caused because the data in the cloud asset inventory is not fresh, CloudQuery will know to fetch those specifically in your central deployment.

## Getting Started

CloudQuery drift detection is currently in alpha (experimental) version, so we welcome any feedback but it’s not yet ready for production use.

Currently the easiest way is actually to run the **fetch** step also in the CI and use something like postgres inside GitHub action, it is not ideal for production deployments but for testing and trying it out this is the fastest way to go.

Check out our [documentation](https://www.cloudquery.io/docs) to get started on how to run it locally and in the CI.

## What do you think?

We would love to hear your feedback (support terraform drift detection for more providers, i.e GCP, Azure etc.) on GitHub, Discord as this will help us get faster to our beta/GA release that will make terraform drift detection ready for production use.
