---
title: Introducing The Tailscale Source Plugin
tag: announcement
date: 2022/12/14
description: >-
  CloudQuery is an open-source, extensible framework that gives you a
  single-pane-of-glass to your cloud-infrastructure using SQL. Today, we are
  happy to announce the release of the Tailscale source plugin.
author: alex
--- 

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

CloudQuery is an open-source, extendable framework that gives you a single-pane-of-glass to your cloud-infrastructure using SQL.
Today, we are happy to announce the release of the Tailscale source plugin for CloudQuery.

[Tailscale](https://tailscale.com/) is a company that provides open-source software defined mesh virtual private network (VPN) software and a web-based management service.

CloudQuery now supports pulling ACL, Device and DNS resources from [Tailscale](https://tailscale.com/) using Tailscale API.

Take a look at our [Configuration](/docs/plugins/sources/tailscale/configuration) section to configure required credentials for the plugin.

Below are some real world query examples to get you started.

## Detect not yet authorized devices

If you have added some devices but didn't authorize them you will not be able to connect them to your [tailnet](https://tailscale.com/kb/1136/tailnet/).
It's a good idea to keep your inventory clean, as Tailscale limits the amount of devices.
Here's a query to check for unauthorized devices:

```sql
select dev.id, dev.user
from tailscale_devices as dev
where not dev.authorized
order by id;
```

This query would output a table of device IDs and corresponding users for unauthorized devices:

```
        id         |       user             
-------------------+------------------
 12345678901234567 | user@example.com
(1 row)
```

You can now put this in a dashboard or set up an alert on it.

## Make sure all your devices have enabled key expiry

To secure your infrastructure it's a good idea to make device keys expire.
Here's a query tp check for devices that have disabled key expiry.

```sql
select dev.id, dev.user
from tailscale_devices as dev
where dev.key_expiry_disabled
order by id;
```

This query would output a table of device IDs and corresponding users for devices that have key expiry disabled:

```
        id         |       user             
-------------------+------------------
 12345678901234567 | user@example.com
(1 row)

```

## What's next

We are going to continue expanding the Tailscale source plugin, adding support for more resources.
Interested in seeing another plugin?
Check out [Creating a New Plugin](/docs/developers/creating-new-plugin) and/or open an issue on our [GitHub](https://github.com/cloudquery/cloudquery).
