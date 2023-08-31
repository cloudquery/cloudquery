---
title: Steampipe alternative | Comparison with CloudQuery
description: CloudQuery is an open source ELT framework that can be used as an alternative to Steampipe. CloudQuery is built for performance, and is easy to deploy and maintain.
---

# CloudQuery vs Steampipe

Steampipe is using [PostgreSQL Foreign data wrappers](https://wiki.postgresql.org/wiki/Foreign_data_wrappers) to create a PostgreSQL abstraction on top of APIs.

**Key Differences:**

- CloudQuery uses an EL (Extract-Load) approach, which means supporting more databases and giving the user the ability to use any standard transformation tools such as [dbt](https://www.getdbt.com/), and visualization tools on top of PostgreSQL (or any other supported database, for that matter).
- **Database Agnostic**: CloudQuery supports multiple databases such as PostgreSQL, BigQuery and others. This makes it play nicely with the whole SQL eco-system and gives you the ability to re-use other tools like Grafana/BI. Steampipe wraps PostgreSQL and while it allows you to expose it as a service, it limits your ability to use other BI tools.
- Steampipe can work better for small-scale live queries, but if you need to get all your cloud assets in one place in larger multi-account setups, CloudQuery is a more scalable approach. You can scale Ingestion, Storage, Transformation workloads independently for each step.
