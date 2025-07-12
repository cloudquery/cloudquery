The Hacker News Source plugin for CloudQuery extracts configuration from the [Hacker News API](https://github.com/HackerNews/API) and loads it into any supported CloudQuery destination (e.g. PostgreSQL, BigQuery, Snowflake, and [more](https://hub.cloudquery.io/plugins/destination)).

It can be used for real applications, but is mainly intended to serve as an example of CloudQuery Source plugin with an incremental table.

## Configuration

The following configuration syncs from Hacker News to a Postgres destination, using a special table (`cq_hackernews_state`) to store the state of the last sync. It is also possible to any other CloudQuery destination as a state backend. For more on this, see [Managing Incremental Tables](/docs/advanced-topics/managing-incremental-tables).

The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec). The config for the `postgresql` destination is not shown here. See our [Quickstart](/docs/quickstart) if you need help setting up the destination.

:configuration

:::callout{type="info"}
Note that if `backend_options` is omitted, by default no backend will be used.
This will result in all items being fetched on every sync.

For more information about managing state for incremental tables, see [Managing Incremental Tables](/docs/advanced-topics/managing-incremental-tables).
:::

- `item_concurrency` (`integer`) (optional):

  The number of items to fetch concurrently. Defaults to `100`.

- `start_time` (`string`) (optional):

  A date-time string in `RFC3339` format.
  For example, `"2023-01-01T00:00:00Z"` will sync all items created on or after January 1, 2023.
  If not specified, the plugin will fetch all items.

  Note that because this is an incremental table, a previous cursor position will take precedence over this setting, unless the given start time is after the last cursor position.

## Example Queries

### Compare the number of mentions for two terms 

```sql copy
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

```text copy
+-------------------+----------+
| name              | mentions |
|-------------------+----------|
| data engineer     | 1415     |
| software engineer | 14411    |
+-------------------+----------+
```

### List the top stories for a given domain in 2022 

```sql copy
SELECT   h.url,
         h.score
FROM     hackernews_items h
WHERE    h.url ilike '%xkcd.com%'
AND      h.time BETWEEN date '2022-01-01' AND date '2023-01-01'
ORDER BY h.score DESC limit 5
```

```text copy
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

```sql copy
SELECT   h.by     AS USER,
         count(*) AS comments
FROM     hackernews_items h
WHERE    h.by != ''
AND      h.type = 'comment'
AND      h.time BETWEEN date '2022-01-01' AND      date '2023-01-01'
GROUP BY h.by
ORDER BY comments DESC limit 3;
```

```text copy
+---------+----------+
| user    | comments |
|---------+----------|
| bombcar | 7307     |
| dang    | 6688     |
| pjmlp   | 6450     |
+---------+----------+
```

### List recently posted remote-friendly YC startup jobs

```sql copy
SELECT   h.time,
         h.title
FROM     hackernews_items h
WHERE    h.type ='job'
AND      h.title ilike '%remote%'
ORDER BY h.time DESC limit 3;
```

```text copy
+---------------------+---------------------------------------------------------------------+
| time                | title                                                               |
|---------------------+---------------------------------------------------------------------|
| 2023-01-09 17:00:05 | Kable (YC W22) Is Hiring Lead Engineer (Remote/US)                  |
| 2023-01-07 12:04:59 | Svix (YC W21) Is Hiring (Remote) – Enterprise-Ready Webhook Service |
| 2022-12-29 21:01:08 | Hive (YC S14) is hiring devs #3-10 in 2023 (Canada remote)          |
+---------------------+---------------------------------------------------------------------+
```
