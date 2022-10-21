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

First we need to [download](https://github.com/cloudquery/cloudquery/releases) our pre-compiled CloudQuery binary, we can also download it with the following brew command:

```bash
brew install cloudquery/tap/cloudquery
# After initial install you can upgrade the version via:
brew upgrade cloudquery
```

Before running CloudQuery we need to generate a `config.hcl`. We won't dive into all the options available and let CloudQuery generate one for us.

```bash
cloudquery init digitalocean
```

Now that we have our `config.hcl` ready we need to add our DigitalOcean Access Key and Spaces Access Key. We can generate our tokens in DigitalOcean's [console](https://cloud.digitalocean.com/settings/api/tokens).

```
# DO Access keys
export DIGITALOCEAN_TOKEN=XXXX

# DO Spaces Access Keys
export SPACES_ACCESS_KEY_ID=XXXX
export SPACES_SECRET_ACCESS_KEY=XXXX
```

## Running

CloudQuery requires us to connect to a database. Either use an existing one (we change the `dsn` in the `config.hcl`), or simply create a database with the following command

```bash
docker run -p 5432:5432 -e POSTGRES_PASSWORD=pass -d postgres
```

Finally, we can execute CloudQuery to fetch all our DigitalOcean resources, simply run the fetch command and let the magic happen.

```bash
cloudquery fetch
```

## Use Cases

After we finish fetching our config data we can make queries for security, compliance, cost management and other purposes.

### Find public facing spaces

```sql
--  Public facing spaces are accessible by anyone: Easily query which space is public facing in your account
SELECT name, location, public, creation_date FROM digitalocean_spaces WHERE public = true;
```

### List Droplets with public facing ipv4 or ipv6

```sql
-- Find any droplets that have a public ipv6 or ipv4 IP
SELECT d.id AS droplet_id, dnv4.ip_address AS ip, dnv4.netmask, dnv4.gateway, dnv6.ip_address AS ipv6, dnv6.netmask AS ipv6_netmask, dnv6.gateway AS ipv6_gateway
	FROM digitalocean_droplets d
LEFT JOIN digitalocean_droplet_networks_v4 dnv4 ON d.cq_id = dnv4.droplet_cq_id
LEFT JOIN digitalocean_droplet_networks_v6 dnv6 ON d.cq_id = dnv6.droplet_cq_id WHERE dnv4.type = 'public' OR dnv6.type = 'public';
```

### Billing History including current month balance

```sql
-- Get you current monthly balance and previous billing histories in one table
SELECT invoice_id as id, description, amount, "date" FROM digitalocean_billing_history
UNION
SELECT'current' AS id, 'current month balance' AS description, month_to_date_usage AS amount , generated_at AS "date" FROM digitalocean_balance;
```

## What's next

We are going to continue to expand and maintain the DigitalOcean Provider, adding support for more current and future resources. Interested in seeing another provider? Check out [Developing New Provider](/docs/developers/creating-new-plugin) and/or open an issue on our [GitHub](https://github.com/cloudquery/cloudquery).
