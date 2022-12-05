---
title: Introducing The Gandi Source Plugin
tag: announcement
date: 2022/12/01
description: >-
  CloudQuery is an open-source, extensible framework that gives you a
  single-pane-of-glass to your cloud-infrastructure using SQL. Today, we are
  happy to announce the release of the Gandi source plugin.
author: kemal
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

CloudQuery is an open-source, extendable framework that gives you a single-pane-of-glass to your cloud-infrastructure using SQL. Today, we are happy to announce the release of the Gandi source plugin for CloudQuery.

Gandi is a company providing domain name registration, web hosting and related services. Gandi also provides a cloud platform, which includes a virtual private server (VPS) offering.

CloudQuery now supports pulling domain and DNS resources from [Gandi](https://gandi.net/) using Gandi's API.

Take a look at our [Configuration](/docs/plugins/sources/gandi/configuration) section to configure required credentials for the plugin.

Below are some real world query examples to get you started.

## Detect domain registrations that will expire soon

If you have more than just one domain, keeping track of which one will expire soon can be a problem. Here's a query to check how many days there are left on each registration:

```sql
select fqdn, dates->>'registry_ends_at' as registry_ends_at, date_trunc('day', (dates->>'registry_ends_at')::timestamp - current_timestamp) as days_left from gandi_domains where ((dates->>'registry_ends_at')::timestamp - interval '90 day') < current_timestamp order by 1;
```

This query would output a table of domain names that are going to expire within the next 90 days:

```
      fqdn      |   registry_ends_at   | days_left
----------------+----------------------+-----------
 yourdomain.co  | 2023-01-28T08:09:41Z | 58 days
 yourdomain.com | 2023-01-31T15:58:08Z | 61 days
(2 rows)
```

You can now put this in a dashboard or set up an alert on it.

## Make sure all your LiveDNS domains have automatic snapshots

Automatic snapshots is a Gandi LiveDNS feature to ensure you always have a backup to a previous state of the domain configuration. With the following query it's possible to list any domains you manage, but which don't have automatic snapshots enabled for some reason.

```sql
select d.fqdn from gandi_livedns_domains d left join gandi_livedns_snapshots s on s.fqdn=d.fqdn and s.automatic where s.fqdn is null;
```

This will return a list of domain names that don't have the automatic snapshots feature enabled. You can set up an alert on this query and make sure your domain configurations are automatically backed up.

```
     fqdn
---------------
 yourdomain.co
(1 row)
```


## Get a list of glue records

Glue records are used to solve circular dependencies in DNS and are crucial to set up correctly and maintain. The following query will list all glue records in your Gandi account.

```sql
select fqdn, name, ips from gandi_domain_glue_records order by 1, 2;
```

This will return:

```
     fqdn       | name |       ips
----------------+------+------------------
yourdomain.com | ns1  | {8.8.8.8}
yourdomain.com | ns2  | {8.8.4.4}
(2 rows)
```

## Get domains in Gandi LiveDNS

Using the following query it's possible to check which domains 'live' LiveDNS (and which don't) and get their configured name servers.

```sql
select fqdn, current, nameservers from gandi_domain_live_dns order by 1;
```

The results look like:

```
       fqdn        | current |                        nameservers
-------------------+---------+------------------------------------------------------------
 yourcomain.com    | livedns | {ns-168-a.gandi.net,ns-16-b.gandi.net,ns-53-c.gandi.net}
 yourdomain.co     | other   | {ns5.gandi.net,ns6.gandi.net}
(2 rows)
```

## What's next

We are going to continue expanding the Gandi source plugin, adding support for more resources. Interested in seeing another plugin? Check out [Creating a New Plugin](/docs/developers/creating-new-plugin) and/or open an issue on our [GitHub](https://github.com/cloudquery/cloudquery).
