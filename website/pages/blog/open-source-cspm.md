---
title: 'Building Open Source CSPM with CloudQuery, PostgreSQL and Grafana'
tag: security
date: 2022/07/24
description: Unbundling the cloud security stack with a data platform
author: yevgenypats
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

In this blog post, we will walk you through how to set up [CloudQuery](https://github.com/cloudquery/cloudquery) to build your own customizable compliance, CSPM (Cloud Security Posture Management) dashboard with PostgreSQL and Grafana.

CSPMs are probably the biggest offenders of yet another dashboard and here in CloudQuery we believe it’s time to unbundle those and apply the best practices in data engineering and the modern data stack to cloud security.

## General Architecture

- **ETL (Extract-Transform-Load) ingestion layer:** [CloudQuery](https://github.com/cloudquery/cloudquery)
- **Datastore:** PostgreSQL
- **Policies:** Standard [SQL Policies](/docs/core-concepts/policies) to be executed via `psql`
- **Data Visualization and Exploration Platform:** Grafana

## What you will get

- **Raw SQL access** to all your cloud asset inventory, open source SQL based policies.
- **Multi-Cloud Asset Inventory:** Ingest configuration from all your clouds to a single datastore with a unified structure.
- **Avoid yet-another-dashboard fatigue:** Reuse your existing BI/Visualization stack (Grafana in this example) to build an open source CSPM.

### Step 1: **Install or Deploy CloudQuery**

If it’s your first time using CloudQuery we suggest you first run it locally to get familiar with the tool. Take a look at our [quick start guide](/docs/quickstart).

If you are already familiar with CloudQuery, take a look at how to deploy it to AWS on RDS Aurora and EKS at [github.com/cludquery/terraform-aws-cloudquery](https://github.com/cloudquery/terraform-aws-cloudquery) , or GCP and Cloud SQL at [https://github.com/cloudquery/terraform-gcp-cloudquery](https://github.com/cloudquery/terraform-gcp-cloudquery)

### Step 2: **Install Grafana**

Grafana is a well-known open source observability and visualization tool. It is open source, so there are a number of ways to deploy it:

- **Self-hosted (local, docker, k8s):** [Official guide.](https://grafana.com/docs/grafana/latest/setup-grafana/installation/)
- **SaaS/managed:** [Grafana.com](https://grafana.com/)
- AWS Managed Grafana: [https://aws.amazon.com/grafana/](https://aws.amazon.com/grafana/)

### Step 3: Run Policies (CSPM - Cloud Security Posture Management)

CloudQuery policies and rules are implemented in pure SQL and they store results in a single table that you can easily query and visualize. Here is a [link](/docs/core-concepts/policies) to all available policies and compliance frameworks. In this section we will go quickly through how to run multiple benchmarks for AWS.

```bash
git clone https://github.com/cloudquery/cloudquery.git
cd cloudquery/plugins/source/aws/policies
# change the DSN to your PostgreSQL instance populated by CloudQuery
psql postgres://postgres:pass@localhost:5432/postgres -f policy.sql
```

This should run all [available](https://github.com/cloudquery/cq-provider-aws/tree/main/policies#policies-and-compliance-frameworks-available) compliance framework and store the results in [aws_policy_results](https://github.com/cloudquery/cq-provider-aws/tree/main/policies#policies-and-compliance-frameworks-available). Now you can query the table directly and export in various formats such as CSV or HTML, all with standard `psql` , and of course visualize them in your favorite BI tool. We prepared a pre-built dashboard for Grafana that you can check out [here](https://github.com/cloudquery/cq-provider-aws/tree/main/dashboards#aws-compliance-and-cspm-cloud-security-posture-management-dashboard):

![](/images/blog/open-source-cspm/image0.png)

## Summary

That’s it! Now you have fully functional CSPM (KSPM, or any other SPM) with those nice bonuses:

1. Access to raw data available and stored in your PostgreSQL.

2. Policies easily customizable and defined in pure SQL.
