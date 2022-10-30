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

First we need to [download](https://github.com/cloudquery/cloudquery/releases) the pre-compiled CloudQuery binary. We can also install it with the following [brew](https://brew.sh/) command:

```bash
brew install cloudquery/tap/cloudquery
# After initial install you can upgrade the version via:
brew upgrade cloudquery
```

Before running CloudQuery we need to generate a `config.yaml` via the `init` command:

```bash
cloudquery init github
```

Now that we have our `config.yaml` ready we need to add our GitHub Personal Access Token.
The required scopes:

- `read:org`
- `read:project`
- `public_repo`

To create one, refer to [this guide](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) by GitHub.
This token should be considered a secret and the config file should not be committed to source control.

Add the token we created, and the organization(s) you want to fetch, to the configuration.

```yaml
  - name: github
    configuration:
      access_token: "<YOUR ACCESS TOKEN HERE>"
      orgs: ["<YOUR ORG NAME>"]
    resources:
      - organizations
      - repositories
      - teams
      - billing.actions
      - billing.packages
      - billing.storage
      - issues
      - hooks
      - installations
      - external_groups
```

## Running

CloudQuery requires us to connect to a database. Either use an existing one (we change the `dsn` in the `config.yaml`), or simply create a database with the following command

```bash
docker run -p 5432:5432 -e POSTGRES_PASSWORD=pass -d postgres
```

Finally, we can execute CloudQuery to fetch all our GitHub resources. Simply run the fetch command and let the magic happen.

```bash
cloudquery fetch
```

## Use Cases

After we finish fetching our config data we can make queries for security, compliance, cost management and other purposes.

### Find all Public Repositories

```sql
-- Query all repositories in the organizations that are public
SELECT id, org, name, description, created_at_time, pushed_at_time,
updated_at_time FROM github_repositories
WHERE private = false
```

### List All Organization Owners

```sql
-- Find all users in organization that have the admin role
SELECT id, org, login, role, state FROM github_organization_members AS gom
    INNER JOIN github_organization_member_membership gomm ON gom.cq_id = gomm.organization_member_cq_id
WHERE role = 'admin'
```

### Organization Billing

```sql
-- Get billing information state for actions, packages and storage
SELECT * FROM github_action_billing

SELECT * FROM github_package_billing

SELECT * FROM github_storage_billing
```

### Long Standing Open Issues

```sql
-- Find all open issues that have been open for more than 7 days
SELECT gr.org, gr.name, gi.id, gi.state, gi.user_login AS opened_by, gi.created_at, gi.updated_at
FROM github_issues gi
INNER JOIN github_repositories gr ON gr.id = gi.repository_id
WHERE state = 'open' AND NOW() - interval '7 days' > created_at
```

## What's next

We are going to continue expanding the GitHub Provider, adding support for more resources. Interested in seeing another provider? Check out [Developing New Provider](/docs/developers/creating-new-plugin) and/or open an issue on our [GitHub](https://github.com/cloudquery/cloudquery).
