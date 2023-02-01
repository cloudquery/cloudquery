---
title: Running AWS PCI DSS with CloudQuery Policies
tag: announcement
date: 2021/12/10
description: >-
  Automate, customize, codify and run PCI DSS Compliance with CloudQuery
  Policies.
author: yevgenypats
---

import { BlogHeader } from "../../components/BlogHeader"
import { Callout } from 'nextra-theme-docs'

<BlogHeader/>

<Callout type="warning">
HCL policies were deprecated - see up-to-date policy documentation [here](https://www.cloudquery.io/docs/core-concepts/policies).
</Callout>

CloudQuery policies gives you a powerful way to automate, customize, codify, and run your cloud security & compliance continuously with HCL and SQL. In this blog we will show you how to run our open-source AWS PCI DSS (Payment Card Industry Data Security Standard) compliance policy. Official PCI DSS Guide is available [here](https://www.pcisecuritystandards.org/).

## Overview

CloudQuery AWS PCI DSS Policy contains more than 40 checks and is available on GitHub.

```hcl
// policy.hcl

policy "pci-dss-v3.2.1" {
    description = "PCI DSS V3.2.1"
    configuration {
        provider "aws" {
            version = ">= v0.5.0"
        }
    }
  .......
    policy "autoscaling" {
        description = "checks for autoscaling"
        query "autoscaling_groups_elb_check" {
            description = "Auto Scaling groups associated with a load balancer should use health checks"
            query = file("queries/autoscaling/autoscaling_groups_elb_check.sql")
        }
    }
}

// queries/autoscaling/autoscaling_groups_elb_check.sql
SELECT "account_id", "region", "arn", "name"
FROM aws_autoscaling_groups
WHERE array_length("load_balancer_names", 1) > 0
AND "health_check_type" IS DISTINCT FROM 'ELB'
```

The policy is split into sections (services) as sub-policies so you can run either the whole policy, sub-policy or even a one specific check. The query itself is defined in a separate file so we can re-use it in other policies (such as CIS or other custom ones).

## Running

Running this is as simple as ensuring your database has the latest cloud asset configuration with the fetch command and then executing each of pre-made queries with the policy run command.

### Quick Start

Following is a quick start to run the policy. Otherwise checkout full details on our [docs](/docs/core-concepts/policies).

### Prerequisite

```bash
# install with brew
brew install cloudquery/tap/cloudquery
# or download precompiled binaries from https://github.com/cloudquery/cloudquery/releases

# Download & Configure AWS Provider
cloudquery init aws

# Connect or run  a local PostgreSQL
docker run -p 5432:5432 -e POSTGRES_PASSWORD=pass -d  postgres

# fetch you cloud assets configuration
cloudquery fetch
```

### Running

```bash
# describe all available policies and sub-policies in the AWS security & compliance pack
cloudquery policy describe aws

# execute the whole policy pack (cis + pci_dss)
cloudquery policy run aws

# execute specific policy pack
cloudquery policy run aws//pci_dss_v3.2.1

# execute specific section in PCI DSS
cloudquery policy run aws//pci_dss_v3.2.1/autoscaling/1
```

You can also output the results into a JSON and pass them to downstream processing for automated monitoring and alerting.

```bash
cloudquery policy run aws//pci_dss_v3.2.1 --output-dir=results
```

## Build your own and share!

Do you have a policy that you want to codify or you’ve been running it with python or bash scripts? You are welcome to try out codifying it with CloudQuery Policies. Feel free to drop on discord or github to get any help and we will share your policy on CloudQuery Hub.
