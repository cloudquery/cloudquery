---
name: GCP
stage: GA
title: GCP Source Plugin
description: CloudQuery GCP source plugin documentation
---
# GCP Source Plugin

:badge

The GCP Source plugin for CloudQuery extracts configuration from a variety of GCP APIs and loads it into any supported CloudQuery destination (e.g. PostgreSQL, BigQuery, Snowflake, and [more](/docs/plugins/destinations/overview)).

## Libraries in Use

- https://cloud.google.com/go/docs/reference
- https://github.com/googleapis/google-api-go-client

## Authentication

:authentication

## Query Examples:

### Find all buckets without uniform bucket-level access

```sql copy
select project_id, name from gcp_storage_buckets where uniform_bucket_level_access->>'Enabled' = 'true';
```
