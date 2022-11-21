---
title: Announcing CloudQuery History
tag: announcement
date: 2021/12/05
description: Announcing CloudQuery History
author: roneliahu
---

import { BlogHeader } from "../../components/BlogHeader"
import { Callout } from 'nextra-theme-docs'

<BlogHeader/>

<Callout type="warning">
This feature was deprecated, [see blog post](https://www.cloudquery.io/blog/migration-and-history-deprecation).
</Callout>

---

Today we are excited to announce the release of CloudQuery History in alpha! CloudQuery History adds [TimescaleDB](https://github.com/timescale/timescaledb) support to give users the ability to track their complete cloud asset inventory snapshots over time!

Achieving better visibility into your cloud infrastructure is key in maintaining security, compliance, cost and operational efficiency, and this is why we started CloudQuery in the [first](https://www.cloudquery.io/blog/announcing-cloudquery-seed-funding) place. Maintaining a historical record of your cloud infrastructure configuration is an integral part of your cloud environment lifecycle.

## Why TimescaleDB?

[TimescaleDB](https://www.timescale.com/) is an open-source relational database with support for time-series data. It uses SQL and works as an extension to PostgreSQL. This ensures all current CloudQuery features such as policies and providers continue to work seamlessly and get out-of-the-box support for historical snapshots!

## Audit logs vs complete snapshots

Current native solutions like AWS CloudTrail, GCP Cloud Audit Log , Azure Activity/Resource Logs and other SaaS services that have audit logs, records only API calls and changes. Obviously, it is advised to enable them (if not already enabled by default) as they can help with investigation & compliance.

Audit logs are great, although they only focus on what **changed** and not on what was the **state** of your whole cloud account at a certain point in time. CloudQuery History provides full historical snapshots of your cloud asset inventory, unlocking the following benefits:

- **Visualize Historical State:** Enhance your current visualization workflows such as Grafana and re-use the [dashboards](https://www.cloudquery.io/blog/open-source-cloud-asset-inventory-with-cloudquery-and-grafana) to view current and historical state.
- **Alert on change using standard SQL:** Use TimescaleDB's [hyperfunctions](https://docs.timescale.com/api/latest/hyperfunctions/) and [continuous aggregates](https://docs.timescale.com/api/latest/continuous-aggregates/) to aggregate at predefined intervals and materialize results and find changes that occurred between fetches.
- **Compliance:** To ensure you were compliant not only in point in time but also over time, you can re-use pre-made and custom [CloudQuery Policies](/docs/core-concepts/policies) to prove compliance over-time.
- **Visibility:** Find resources that might already be deleted. Inspect what was created and understand what happened.
- **Postmortems and incident response:** Full historical snapshots of your cloud assets allow you to gain better insights into what happened in your environment and determine the blast radius. Re-use any standard analytics or BI tools.

## Getting Started

History is currently in alpha version, so we welcome any feedback as it's not yet ready for production use.

Setting up History is fairly simple, you are required to either install the TimescaleDB extension on your existing PostgreSQL instance or setup a self hosted TimescaleDB instance. See [here](https://docs.timescale.com/timescaledb/latest/how-to-guides/install-timescaledb/self-hosted/) for more details.

See CloudQuery [quickstart guide](/docs/quickstart) and [history configuration](/docs) for more details on how to configure CloudQuery to run with History enabled.

## What next?

CloudQuery history opens up endless possibilities for managing compliance, security, visualization and much more! We would love to hear your feedback, either on [GitHub](https://github.com/cloudquery/cloudquery) or [Discord](https://www.cloudquery.io/discord).
