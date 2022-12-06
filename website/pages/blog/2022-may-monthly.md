---
title: May 2022 Monthly Updates
tag: recap
date: 2022/06/06
description: >-
  CloudQuery Monthly Newsletter, May 2022. Updates to features, providers,
  deployment and new blog posts.
author: benjamin
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

Welcome to CloudQuery’s latest product news! This month we’ve been heads-down working on some new features but most
importantly making CloudQuery easier to use no matter if you are a developer/contributor or an operator i.e -
DevOps/SRE/Security team.

## 📈 High Level Metrics

- Number of Releases: 26
- New Resources: 12
  - AWS: 9
  - GCP: 3
- Bug fixes: 100+

## GCP Provider Improvements

**Project Auto Discovery**: The [GCP Provider](/docs/plugins/sources/overview)
can now auto-discover all projects under your organization and recursively fetch all the configuration from each project
(and project sprawl in GCP is common!). This should significantly simplify CloudQuery configuration and maintenance as
well as open-up opportunity to build new workflows and alert on things like project creation.

## History and Migration Removal

Sometimes we also remove features instead of adding to ensure we are able to move fast and with quality. This should
increase developer velocity, both internally and externally, remove 🐛 as well as open the door to support more
databases! Checkout our [blog](https://www.cloudquery.io/blog/migration-and-history-deprecation) for full explanation.

## Store Policy Data in The Database

CloudQuery supports six open-source security & compliance [policies](/docs/core-concepts/policies) implemented
in SQL so you can codify your security and compliance posture. With this release you can also store the result in
PostgreSQL and enable more workflows downstream like monitoring security results in your favorite BI tool,
alerting and much more! To see more checkout the [documentation](/docs/core-concepts/policies).

## Deployments

Running CloudQuery locally is cool but what is even better is running it remotely with non-ephermal database so
everyone in your team can use it, build custom workflows on top, access it directly or via your favorite BI tool.
To easy the deployment we release:

- [Terraform Module](https://github.com/cloudquery/terraform-aws-cloudquery) to deploy CloudQuery on AWS
- [Terraform Module](https://github.com/cloudquery/terraform-gcp-cloudquery) to deploy CloudQuery on GCP
- Or deploy it directly on k8s with a [Helm Chart](https://github.com/cloudquery/helm-charts)

## 📚 Blogs

- [Recreating a Deleted SSO IDP](https://www.cloudquery.io/blog/aws-sso-if-deleted-sso-identity-provider): A tutorial
  on how to recover when you accidentally delete your AWS SSO Identity Provider from an important account
- [AWS Resource View](https://www.cloudquery.io/blog/aws-resources-view): Utilizing CloudQuery data create a view
  that allows you to see data from all of your AWS Accounts

- Connecting CloudQuery Database to a BI Tool:
  - [Apache Superset](https://www.cloudquery.io/blog/cloud-asset-inventory-cloudquery-apache-superset)
  - [AWS QuickSight](https://www.cloudquery.io/blog/cloud-asset-inventory-cloudquery-aws-quicksight)
  - [Metabase](https://www.cloudquery.io/blog/cloud-asset-inventory-cloudquery-metabase)
