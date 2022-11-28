---
title: Introducing The GitHub CloudQuery Provider
tag: announcement
date: 2022/08/08
description: >-
  CloudQuery is an open-source, extensible framework that gives you a
  single-pane-of-glass to your cloud-infrastructure using SQL. Today, we are
  happy to announce the release of the GitHub Provider for CloudQuery. 
author: roneliahu
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

CloudQuery is an open-source, extendable framework that gives you a single-pane-of-glass to your cloud-infrastructure using SQL. Today, we are happy to announce the release of the GitHub Provider for CloudQuery.

GitHub is a source-control provider that helps developers store and manage their code, as well as track and control changes to their code.

In this tutorial, we will install CloudQuery and use it to fetch GitHub resources. Then, we will use SQL to get visibility into security, compliance and cost-management in GitHub.

## Setup

- [Download CloudQuery](https://www.cloudquery.io/docs/quickstart).
- Acquire a GitHub (personal access token)[https://docs.github.com/en/enterprise-server@3.4/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token]
  with the scopes:
  ```
    read:org
    read:project
    public_repo
  ```.
- Create a `cloudquery.yml` file, similar to the following:
  ```yaml
  kind: source
  spec:
    name: github
    path: cloudquery/github
    version: "VERSION_SOURCE_GITHUB"
    tables: ["*"]
    destinations: ["postgresql"]
    
    spec:
      access_token: <YOUR_ACCESS_TOKEN>
      orgs: ["cloudquery"]
  ---
  kind: destination
  spec:
    name: "postgresql"
    path: cloudquery/postgresql
    version: "VERSION_DESTINATION_POSTGRESQL"
    write_mode: "overwrite-delete-stale"

    spec:
      connection_string: "postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable"
  ```
- Run
  ```bash
  cloudquery sync cloudquery.yml
  ```

## Use Cases

After we finish fetching our config data we can make queries for security, compliance, cost management and other purposes.

### Find all Public Repositories

```sql
-- Query all repositories in the organizations that are public
SELECT id, org, name, description FROM github_repositories
WHERE private = false
```

### Organization Billing

```sql
-- Get billing information state for actions, packages and storage
SELECT * FROM github_billing_action

SELECT * FROM github_billing_package

SELECT * FROM github_billing_storage
```

### Long Standing Open Issues

```sql
-- Find all open issues that have been open for more than 7 days
SELECT gr.org, gr.name, gi.id, gi.state, gi.created_at, gi.updated_at
FROM github_issues gi
INNER JOIN github_repositories gr ON gr.id = (gi.repository->'id')::bigint
WHERE state = 'open' AND NOW() - interval '7 days' > gi.created_at
```

## What's next

We are going to continue expanding the GitHub Provider, adding support for more resources. Interested in seeing another provider? Check out [Developing New Provider](/docs/developers/creating-new-plugin) and/or open an issue on our [GitHub](https://github.com/cloudquery/cloudquery).
