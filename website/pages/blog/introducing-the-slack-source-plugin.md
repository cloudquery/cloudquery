---
title: Introducing the Slack Source Plugin
tag: announcement
date: 2022/12/20
description: >-
  CloudQuery is an open source high performance data integration 
  platform designed for security and infrastructure teams. Today, we are
  happy to announce the release of the Slack source plugin.
author: hermanschaaf
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

CloudQuery is an open source high performance data integration platform designed for security and infrastructure teams. Today, we are happy to announce the release of the Slack source plugin for CloudQuery.

Slack is a popular communication tool used by many organizations. With the CloudQuery Slack source plugin, you can now load Slack workspace data into Postgres, Snowflake, BigQuery, SQLite, or any other destination supported by CloudQuery. 

As part of its initial release, the Slack plugin supports pulling data for the following Slack resources:
 - users into the [`slack_users`](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/slack/docs/tables/slack_users.md) and [`slack_user_presences`](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/slack/docs/tables/slack_user_presences.md) tables,
 - teams into the [`slack_teams`](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/slack/docs/tables/slack_teams.md) table,
 - channels into the [`slack_conversations`](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/slack/docs/tables/slack_conversations.md) table,
 - messages into the [`slack_conversation_histories`](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/slack/docs/tables/slack_conversation_histories.md) and [`slack_conversation_replies`](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/slack/docs/tables/slack_conversation_replies.md) tables,
 - files into the [`slack_files`](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/slack/docs/tables/slack_teams.md) table

Let's look at a two use cases for the Slack plugin: one serious, one fun. 

## Use cases

### Slack Security

It's important to make sure that your Slack workspace is secure and that you have visibility into who has access to your workspace. The Slack source plugin can help you with this task. With our data loaded into Postgres, we can start by writing a query to find out who has access to our workspace:

```sql
select name, profile->>'email' as email from slack_users
```

Now, let's filter this list down to users who don't have a company email address (in our case `@cloudquery.io`), and check whether these users are properly restricted:

```sql
select 
    name, 
    profile->>'email' as email, 
    is_restricted, 
    is_ultra_restricted, 
    is_stranger
 from slack_users
 where
     not deleted
     and not is_bot
     and not name = 'slackbot'
     and profile->>'email' not like '%cloudquery.io' 
```

We can also cross-reference user accounts against data from the [Okta plugin](/docs/plugins/sources/okta/overview) to see if they should still have Slack access:

```sql
select 
    su.name as name, 
    su.profile->>'email' as email 
from slack_users su 
    left join okta_users ou 
        on su.profile->>'email' = ou.profile->>'email' 
where 
    ou.id is null 
    and not su.is_bot 
    and not name = 'slackbot'
    and not su.deleted
```

The output of the above query will show us all Slack users who don't have an active Okta account, and therefore shouldn't have access to our Slack workspace.

### Custom Analytics

The Slack plugin can also be used to build custom analytics on top of your Slack workspace data. For example, let's say we want to find out which channels are the most active in our workspace. We can do this by writing a query that counts the number of messages in each channel:

```sql
select
     c.name as channel,
     count(r.*) as messages
 from slack_conversations c
     join slack_conversation_replies r on c.id = r.channel_id
 where
     c.is_channel
     and not c.is_archived 
 group by c.name 
 order by messages desc
```

(Note: the CloudQuery plugin only collects data for channels that the bot has been added to.)

Or, perhaps we are curious about the most active users in our workspace. We can write a query that counts the number of messages sent in public channels by each user:

```sql
select 
    u.name, count(h.user) 
from slack_conversation_histories h 
    join slack_conversation_replies r on h.ts = r.conversation_history_ts 
    join slack_users u on u.id = h.user 
group by u.name order by count desc
```

We can even break this down by day, week, or month to graph how active users are over time:

```sql
select
     u.name,
     date_trunc('day', to_timestamp(round(h.ts::float))) as day,
     count(h.user)
from slack_conversation_histories h
     join slack_conversation_replies r on h.ts = r.conversation_history_ts
     join slack_users u on u.id = h.user 
group by u.name, day 
order by day, count desc
```

## Getting Started

To get started syncing Slack data, see the [Slack source plugin documentation](/docs/plugins/sources/slack/overview) for instructions.

## What's next

The Slack API comes with some strict [rate limits](https://api.slack.com/docs/rate-limits) that mean syncing messages from channels with a long history can take a long time. We are [thinking about ways to address this](https://github.com/cloudquery/cloudquery/issues/5809), please 👍 or comment on the GitHub issue if you are interested!

We are also going to continue expanding the Slack source plugin, adding support for more resources as they become available in the Slack API. If you are interested in a specific Slack resource, feel free to raise an [issue on GitHub](https://github.com/cloudquery/cloudquery/issues). Or if you need some help to get started, join us over on [Discord](https://www.cloudquery.io/discord), we'd love to help.
