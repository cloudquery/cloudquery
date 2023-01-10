---
title: Announcing the Vercel Source Plugin
tag: announcement
date: 2022/12/14
description: >-
  CloudQuery is an open-source, extensible framework that gives you a
  single-pane-of-glass to your cloud-infrastructure using SQL. Today, we are
  happy to announce the release of the Vercel source plugin.
author: kemal
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

CloudQuery is an open-source, extendable framework that gives you a single-pane-of-glass to your cloud-infrastructure. Today, we are happy to announce the release of the Vercel source plugin for CloudQuery.

If you're not familiar with the Vercel platform, we recommend you visit their [homepage](https://vercel.com/).

CloudQuery now supports pulling resources from [Vercel](https://vercel.com/) using Vercel's API.

Take a look at our [Configuration](/docs/plugins/sources/vercel/configuration) section to configure required credentials for the plugin.

Below are some query examples to get you started.

## Detect domain registrations that will expire soon

If you have more than just one domain, keeping track of which one will expire soon can be a problem. Here's a query to check how many days there are left on each registration:

```sql
select name, expires_at, date_trunc('day', expires_at - current_timestamp) as days_left from vercel_domains where (expires_at - interval '90 day') < current_timestamp order by 1;
```

This query would output a table of domain names that are going to expire within the next 90 days:

```bash
      name      |   expires_at         | days_left
----------------+----------------------+-----------
 yourdomain.co  | 2023-01-28T08:09:41Z | 44 days
 yourdomain.com | 2023-01-31T15:58:08Z | 47 days
(2 rows)
```

Vercel also provides an [auto-renewal facility](https://vercel.com/docs/concepts/projects/domains/renew-a-domain) which will automatically renew your domain registration when it's about to expire.


## Get all Vercel team members

Using the query below you can list all Vercel team members across multiple teams.

```sql
select t.name AS team, u.username, u.name, u.role from vercel_teams t join vercel_team_members u on u.team_id=t.id order by 1, 2;
```

This will return list of users per team and their role:

```bash
   name   |   username   | name         |  role
------------+--------------+------------+--------
 yourteam | user1        | User Name    | MEMBER
 yourteam | user2        |              | MEMBER
 yourteam | user3        | Another User | OWNER
(3 rows)
```


## Get a list of name servers

```sql
select name, intended_nameservers, custom_nameservers, nameservers from vercel_domains order by 1;
```

This query will return a list of domains and their name servers.

## What's next

We are going to continue expanding the Vercel source plugin, adding support for more resources. Interested in seeing another plugin? Check out [Creating a New Plugin](/docs/developers/creating-new-plugin) and/or open an issue on our [GitHub](https://github.com/cloudquery/cloudquery).
