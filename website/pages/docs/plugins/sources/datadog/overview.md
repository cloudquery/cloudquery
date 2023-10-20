---
name: Datadog
stage: GA
title: Datadog Source Plugin
description: CloudQuery Datadog Plugin documentation
---
# Datadog Source Plugin

:badge

The CloudQuery Datadog plugin extracts your Datadog information and loads it into any supported CloudQuery destination (e.g. PostgreSQL, BigQuery, Snowflake, and [more](/docs/plugins/destinations/overview)).

## Authentication

:authentication

## Query Examples

### Find all not verified users

```sql copy
SELECT 
    "attributes" ->> 'name' AS username
FROM
    datadog_users
WHERE
    ("attributes" ->> 'verified')::boolean is distinct FROM TRUE
```
