---
title: CloudQuery and LocalStack
tag: integration
date: 2022/12/23
description: >-
  How to setup CloudQuery to work with LocalStack
author: benjamin
---

import { HowToGuideHeader } from "../../components/HowToGuideHeader"
import { Callout } from 'nextra-theme-docs'

<HowToGuideHeader/>


In this guide we will walk through how to configure CloudQuery to sync from a LocalStack instance.



## Introduction to LocalStack

LocalStack is a fully functional local cloud stack that enables developers to Develop and test their cloud and serverless applications offline.  It allows developers to test their cloud applications locally, without the need for an internet connection or an AWS account. This can be particularly useful for development and testing, as it allows developers to work on their applications offline and avoid incurring any charges for using AWS services.



## Walkthrough

Before beginning this tutorial make sure you have the following tools installed:
- [Docker](https://www.docker.com/products/docker-desktop/)
- [CloudQuery](https://www.cloudquery.io/docs/quickstart)


### Step 1

Start LocalStack running on your machine:

```bash
docker run --rm -it \
    -p 4566:4566 \
    -p 4510-4559:4510-4559 \
    -e DEBUG=1 \
    localstack/localstack
```
## Step 2

Configure CloudQuery to use the LocalStack endpoint
```yaml
kind: source
spec:
  # Source spec section
  name: "aws"
  registry: "github"
  path: "cloudquery/aws"
  version: "VERSION_SOURCE_AWS"
  destinations: ["postgresql"]
  skip_tables:
    - aws_route53_delegation_sets
    - aws_iam_policies
  tables:
    - "*"
  spec:
    regions: 
      - "us-east-1"
    
    # Configure the AWS SDK to use the localstack endpoint
    custom_endpoint_url: http://localhost:4566
    custom_endpoint_hostname_immutable: true
    custom_endpoint_partition_id: "aws"
    custom_endpoint_signing_region: "us-east-1"
    # There is no reason to retry failed requests to localstack
    max_retries: 0
```
<Callout>
Note that it is important to skip `aws_route53_delegation_sets` and `aws_iam_policies` as bugs in LocalStack force CloudQuery into an infinite loop
</Callout>

### Step 3

Run CloudQuery


``` bash
cloudquery sync config.yml
```
