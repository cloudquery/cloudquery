---
title: CloudSploit alternative | Comparison with CloudQuery
---

# CloudQuery vs CloudSploit

CloudSploit by Aqua is an open-source project designed to allow detection of security risks in cloud infrastructure accounts, including: Amazon Web Services (AWS), Microsoft Azure, Google Cloud Platform (GCP), Oracle Cloud Infrastructure (OCI), and GitHub. The scripts are designed to return a series of potential misconfigurations and security risks.

## Similarities

CloudQuery can be used as an alternative to CloudSploit. Here are some of the similarities:

- Both CloudQuery and CloudSploit are open-source projects.
- Both CloudQuery and CloudSploit can be used to detect security risks in cloud infrastructure accounts and support the big cloud providers (AWS, GCP, Azure, OCI, GitHub).
- Both CloudQuery and CloudSploit can be used to return a series of potential misconfigurations and security risks. In CloudQuery, this is done via Policies: plain SQL that should be executed after a sync completes. See [AWS Policies](/docs/plugins/sources/aws/policies), [GCP Policies](/docs/plugins/sources/gcp/policies) and [Azure Policies](/docs/plugins/sources/azure/policies), for example.

## Key Differences

There are some key differences between CloudSploit and CloudQuery:

- While CloudSploit is available in self-hosted and SaaS versions, CloudQuery is open-source and currently only supports a self-hosted version.
- CloudQuery supports more cloud providers. Apart from AWS, Azure and GCP, Oracle Cloud Infrastructure (OCI) and GitHub, CloudQuery also supports Azure DevOps, Alibaba Cloud, Datadog, Gandi, Heroku, Kubernetes, Tailscale, Vercel, and [many more](/docs/plugins/sources/overview).
- CloudQuery table and column names are generally more consistent and predictable, as they are taken directly from the cloud provider's API and/or SDK.
- Like CloudSploit, CloudQuery can store output as JSON, but it primarily supports collection into databases, data warehouses or data lakes like PostgreSQL, BigQuery, Snowflake [and more](/docs/plugins/destinations/overview). Postgres is recommended for use with built-in Policies. CloudQuery does not currently enable the storage of raw responses from the cloud provider's API: only normalized results can be stored.
- CloudQuery does not support remediations out of the box, but remediation-like functionality can be built on top of the policies provided by CloudQuery. Once a remediation has been applied, CloudQuery can be synced again to validate the remediation.
