---
title: Announcing CloudQuery Policies
tag: announcement
date: 2021/06/25
description: Announcing First Release of CloudQuery Policies
author: michelvocks
---

import { BlogHeader } from "../../components/BlogHeader"
import { Callout } from 'nextra-theme-docs'

<BlogHeader/>

<Callout type="warning">
HCL policies were deprecated - see up-to-date policy documentation [here](https://www.cloudquery.io/docs/core-concepts/policies).
</Callout>

We are excited to announce the release of the CloudQuery Policies!

CloudQuery brings the power of SQL to your cloud infrastructure, providing easy monitoring, governance and security.

The new CloudQuery Policies feature brings policy-as-code to the CloudQuery ecosystem. CQ Policies enable users to codify, version and run security, governance, cost and compliance rules, using SQL.

## What’s inside?

- **HCL support**: Policies can be written in HCL and/or JSON as the logical layer and SQL as the query layer.
- **Native GitHub support**: Policies can be downloaded and run directly from GitHub via the CQ CLI (`cloudquery policy download --help`)
- **Views**: Users can create complex SQL views once and reference them in subsequent queries.
- **Query and sub-policy execution**: Users can execute specific queries or sub-policies via CQ CLI (`cloudquery policy run my-policy --sub-path=my-sub-policy`).
- **Inline policy functions**: New inline policy functions are now available that allow users to define more dynamic policies. One example is the new file function that allows users to out-source query definitions to other files and to dynamically insert them during execution time.

You can see all of the available policy packs, as well as their documentation and detailed specification of all checks, on [hub.cloudquery.io/policies](/docs/core-concepts/policies).

For detailed documentation please check-out our [docs](/docs/core-concepts/policies).

## Example Policies

### Basic

At the basic level, each policy contains the minimum version of cq-provider that it requires, and a list of SQL queries with their respective description.

```hcl
policy "test-policy" {
  description = "This is a test policy"
  configuration {
    provider "aws" {
      version = ">= 1.0"
    }
  }

  query "top-level-query" {
    description = "Top Level Query"
    query = "SELECT * FROM test_policy_table WHERE name LIKE 'peter'"
  }
}

```

### Views

Sometimes there is the need to join between multiple tables. Instead of creating complex queries multiple times, views gives you the ability to define the view once and then reference the created view in other queries in your policies.

```hcl
policy "test-policy" {
  description = "This is a test policy"
  configuration {
    provider "aws" {
      version = ">= 1.0"
    }
  }

  view "myview" {
    description = "My awesome view"
    query "complex-query" {
      query = "SELECT * FROM test_policy_table WHERE name LIKE 'john'"
    }
  }

  query "top-level-query" {
    description = "Top Level Query"
      query = "SELECT * FROM myview"
  }
}
```

## Policy-In-Policy

It is possible to nest policies inside policies to build a policy hierarchy that will give you the ability to reference and execute only a subset of the main policy file.

```hcl
policy "test-policy" {
  description = "Test Policy"
  configuration {
    provider "aws" {
      version = ">= 1.0"
    }
  }

  view "testview" {
    description = "Test View"
    query "testviewquery" {
      query = "SELECT * FROM test_policy_table WHERE name LIKE 'john'"
    }
  }

  query "top-level-query" {
    description = "Top Level Query"
    query = "SELECT * FROM test_policy_table WHERE name LIKE 'peter'"
  }

  policy "sub-policy-1" {
    description = "Sub Policy 1"
    query "sub-level-query" {
      query = "SELECT * from testview"
      expect_output = true
    }

    policy "sub-sub-policy-1" {
      description = "Sub Sub Policy 1"
      query "sub-sub-level-query" {
        query = "SELECT * from test_policy_table WHERE name LIKE 'peter'"
      }
    }
  }

  policy "sub-policy-2" {
    description = "Sub Policy 2"
    query "sub-level-query" {
      query = "SELECT * from test_policy_table WHERE name LIKE 'peter'"
    }
  }
}
```

## Running

We extended the CloudQuery CLI to support downloading and running policies directly from GitHub.

Download a policy repository:

```bash
cloudquery policy download cq-policy-core
```

Run specific policy (latest version):

```bash
cloudquery policy run cq-policy-core aws/cis-v1.20
```

Run specific policy version:

```bash
cloudquery policy run cq-policy-core aws/cis-v1.20@v0.0.1
```

You can also run a specific query inside a policy:

```bash
cloudquery policy run cq-policy-core aws/cis-v1.20 --sub-path="aws-cis-section-1/1.1"
```

## What’s Next?

We are always eager to hear feedback so feel free to file feature-requests/bugs/issues at [github.com/cloudquery/cloudquery](https://github.com/cloudquery/cloudquery/issues).

Also, we have more exciting features coming up to enhance the new policy feature so subscribe to our [twitter](https://twitter.com/cloudqueryio) and/or mailing list.
