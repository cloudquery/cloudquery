---
title: FAQ
description: Frequently asked questions about CloudQuery.
---

# FAQ

## Does CloudQuery access any application data in my cloud?

No. CloudQuery cloud provider integrations like AWS, GCP and Azure generally only access metadata and configuration data. Some tables like `aws_cloudwatch_metrics` and `aws_cloudwatch_logs` can be used to sync log and metric data, but only if you select these tables.

## What happens when I run two (or more) syncs? Will the second sync remove resources that no longer exist from the database?

There are currently three types of `write` modes in destination integrations: `overwrite-delete-stale`, `overwrite`, and `append`. The default is `overwrite-delete-stale`.
- In `overwrite-delete-stale` data will be upserted based on primary keys and stale data will be deleted by deleting any data fetched by the same previous source configuration.
- In `overwrite` data will be upserted based on primary keys and it will be up to the user to setup recurring task to delete stale data. 
- In `append`, old rows are never deleted or updated - every sync adds new rows.

In `overwrite` and `append` mode, you can distinguish between rows from different syncs by inspecting the `_cq_sync_time` column.