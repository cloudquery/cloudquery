---
title: Introducing The DigitalOcean CloudQuery Provider
tag: announcement
date: 2021/09/17
description: >-
  CloudQuery is an open-source, extensible framework that gives you a
  single-pane-of-glass to your cloud-infrastructure using SQL. Today, we are
  happy to announce the release of the DigitalOcean Provider for CloudQuery.
author: roneliahu
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

CloudQuery is an open-source, extendable framework that gives you a single-pane-of-glass to your cloud-infrastructure using SQL. Today, we are happy to announce the release of the DigitalOcean Provider for CloudQuery.

DigitalOcean is a cloud-provider that is popular both with small and larger companies. Some of DigitalOceans' advantages are its ease-of-use, flat pricing, and cheaper pricing for some services. This release brings the power of CloudQuery to DigitalOcean users, DevOps engineers and SREs, helping solve visibility, security and compliance challenges with SQL.

In this short tutorial, we will install CloudQuery and use it to fetch a DigitalOcean resources. Then, we will use SQL to get visibility into security, compliance and cost-management in DigitalOcean.

## Setup

- Follow our [quickstart guide](/docs/quickstart) to setup cloudquery.
- [authenticate with DigitalOcean](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/digitalocean#authentication).
- Run `cloudquery sync`.

## Use Cases

After we finish fetching our config data we can make queries for security, compliance, cost management and other purposes.

### Find public facing spaces

```sql
--  Public facing spaces are accessible by anyone: Easily query which space is public facing in your account
SELECT bucket->>'Name',location,public FROM digitalocean_spaces WHERE public = true;
```

### List Droplets with public facing ipv4 or ipv6

```sql
-- Find any droplets that have a public ipv6 or ipv4 IP
SELECT id, name, v4->>'ip_address' AS address_v4, v4->>'netmask' AS netmask_v4, v4->>'gateway' AS gateway_v4,
       v6->>'ip_address' AS address_v6, v6->>'netmask' AS netmask_v6, v6->>'gateway' AS gateway_v6
FROM 
  (SELECT id,name,v4,NULL as v6 FROM digitalocean_droplets CROSS JOIN JSONB_ARRAY_ELEMENTS(digitalocean_droplets.networks->'v4') AS v4 
  UNION
  SELECT id,name,NULL as v4,v6 FROM digitalocean_droplets CROSS JOIN JSONB_ARRAY_ELEMENTS(digitalocean_droplets.networks->'v6') AS v6) AS union_v46
WHERE v4->>'type' = 'public' OR v6->>'type' = 'public';
```

## What's next

We are going to continue to expand and maintain the DigitalOcean Provider, adding support for more current and future resources. Interested in seeing another provider? Check out [Developing New Provider](/docs/developers/creating-new-plugin) and/or open an issue on our [GitHub](https://github.com/cloudquery/cloudquery).
