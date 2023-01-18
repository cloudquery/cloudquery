# Hacker News Source Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("source", `hackernews`)}/>

The Hacker News Source plugin for CloudQuery extracts configuration from the [Hacker News API](https://github.com/HackerNews/API) and loads it into any supported CloudQuery destination (e.g. PostgreSQL).

It can be used for real applications, but is mainly intended to serve as an example of an incremental CloudQuery Source plugin. 

## Configuration

The following configuration syncs from Hacker News to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](https://www.cloudquery.io/docs/reference/source-spec). The config for the `postgresql` destination is not shown here. See our [Quickstart](/docs/quickstart) if you need help setting up the destination.

```yaml
kind: source
spec:
  name: "hackernews"
  path: "cloudquery/hackernews"
  version: "VERSION_SOURCE_HACKERNEWS"
  tables: ["*"]
  backend: local
  destinations: 
    - "postgresql"
  spec:
    item_concurrency: 100
```

import { Callout } from 'nextra-theme-docs'

<Callout type="info">

Note that if `backend: local` is not specified, the default will be no backend. This will result in all items being fetched on every sync, instead of incremental syncs.

</Callout>

- `item_concurrency` (int, optional):
    The number of items to fetch concurrently. Defaults to 100.

- `start_time` (string, optional):
    A date-time string in RFC3339 format. For example, `"2023-01-01T00:00:00Z"` will sync all items created on or after January 1, 2023. If not specified, the plugin will fetch all items. Note that because this is an incremental table, a previous cursor position will take precedence over this setting, unless the given start time is after the last cursor position.

## Example Queries

### Compare the number of mentions for two terms 

```sql
SELECT 'data engineer' AS NAME,
       count(*)        AS mentions
FROM   hackernews_items
WHERE  title ilike '%data engineer%'
OR     text ilike '%data engineer%'
UNION
SELECT 'software engineer' AS NAME,
       count(*)            AS mentions
FROM   hackernews_items
WHERE  title ilike '%software engineer%'
OR     text ilike '%software engineer%';
```

```text
+-------------------+----------+
| name              | mentions |
|-------------------+----------|
| data engineer     | 1415     |
| software engineer | 14411    |
+-------------------+----------+
```

### List the top stories for a given domain in 2022 

```sql
SELECT   h.url,
         h.score
FROM     hackernews_items h
WHERE    h.url ilike '%xkcd.com%'
AND      h.time BETWEEN date '2022-01-01' AND date '2023-01-01'
ORDER BY h.score DESC limit 5
```

```text
+-------------------------------+-------+
| url                           | score |
|-------------------------------+-------|
| https://what-if.xkcd.com/158/ | 387   |
| https://xkcd.com/2617/        | 361   |
| https://xkcd.com/             | 100   |
| https://xkcd.com/2682/        | 77    |
| https://what-if.xkcd.com/161/ | 54    |
+-------------------------------+-------+
```

### List the top 3 users by number of comments in 2022

```sql
SELECT   h.by     AS USER,
         count(*) AS comments
FROM     hackernews_items h
WHERE    h.by != ''
AND      h.type = 'comment'
AND      h.time BETWEEN date '2022-01-01' AND      date '2023-01-01'
GROUP BY h.by
ORDER BY comments DESC limit 3;
```

```text
+---------+----------+
| user    | comments |
|---------+----------|
| bombcar | 7307     |
| dang    | 6688     |
| pjmlp   | 6450     |
+---------+----------+
```

### List recently posted remote-friendly YC startup jobs

```sql
SELECT   h.time,
         h.title
FROM     hackernews_items h
WHERE    h.type ='job'
AND      h.title ilike '%remote%'
ORDER BY h.time DESC limit 3;
```

```text
+---------------------+---------------------------------------------------------------------+
| time                | title                                                               |
|---------------------+---------------------------------------------------------------------|
| 2023-01-09 17:00:05 | Kable (YC W22) Is Hiring Lead Engineer (Remote/US)                  |
| 2023-01-07 12:04:59 | Svix (YC W21) Is Hiring (Remote) – Enterprise-Ready Webhook Service |
| 2022-12-29 21:01:08 | Hive (YC S14) is hiring devs #3-10 in 2023 (Canada remote)          |
+---------------------+---------------------------------------------------------------------+
```