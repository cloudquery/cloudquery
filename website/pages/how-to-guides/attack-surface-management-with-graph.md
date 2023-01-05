---
title: How to use CloudQuery for Attack Surface Management and Graph Visualization
tag: tutorial
description: >-
  How to use CloudQuery for Attack Surface Management (ASM) and Graph Visualization with Neo4j
author: jsonkao
---

import { HowToGuideHeader } from "../../components/HowToGuideHeader"

<HowToGuideHeader/>

In this guide, we will demonstrate how to set up CloudQuery for customizable Attack Surface Management (ASM) and how to get started with utilizing graph visualization to execute security queries.

## Walkthrough

### Prerequisites

In this guide, we will use Neo4j as a destination and AWS as a source.  For more information on how to set those up, see our documentation on [Neo4j](https://www.cloudquery.io/docs/plugins/destinations/neo4j/overview) and [AWS](https://www.cloudquery.io/docs/plugins/sources/aws/overview).

Refer to Neo4j's installation documentation (https://neo4j.com/docs/operations-manual/current/installation/) for help setting up Neo4j. For this walkthrough, make sure a local instance of Neo4j is up and running. 

### Step 1: Install or Deploy CloudQuery

To get started with CloudQuery, check out our [quickstart guide](/docs/quickstart) and [AWS source plugin](/docs/plugins/sources/aws/overview) for how to configure CloudQuery and run it locally on your machine.

### Step 2: Sync Data from CloudQuery to Neo4j

The following command will sync data from AWS as a source to Neo4J as a destination:

`cloudquery sync aws-neo.yml neo4j.yml`

For more information on configuration files, see [AWS source configuration](https://www.cloudquery.io/docs/plugins/sources/aws/configuration) and [Neo4j destination configuration](https://www.cloudquery.io/docs/plugins/destinations/neo4j/overview)

You should see a `sync completed successfully` message.  CloudQuery has now synced your AWS data to Neo4j.

### Step 3: Test Out Neo4j

If running Neo4j locally, navigate to `http://localhost:7474/browser/` for the Neo4j browser. 

Let's start with a simple query to find all our IAM Roles.  

`MATCH (n:aws_iam_roles) RETURN n`

![IAM Roles](/images/how-to-guides/attack-surface-management-with-graph/iam-roles.png)

### Step 4: Create Relationships in Neo4j

### Step 5: Run Custom ASM Queries

We will now go through 3 example queries:

* IAM

* Networking

* Public Resources


## Summary

We have demonstrated how to get started with Attack Surface Management (ASM) and graph visualization with Neo4j along with some starter queries.  Now you should be able to customize and create more queries, relationships, and utilize CloudQuery to help improve the security posture of your organization!