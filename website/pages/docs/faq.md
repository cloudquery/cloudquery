---
title: FAQ
---

# FAQ

## Does CloudQuery access any application data in my cloud?

No. CloudQuery only accesses metadata and configuration data. It never pulls data from your application databases or cloud storage files.

## What happens when I run two (or more) syncs? Will the second sync remove resources that no longer exist from the database?

There are currently three types of `write` modes in destination plugins: `overwrite-delete-stale`, `overwrite`, and `append`. The default is `overwrite-delete-stale`.
- In `overwrite-delete-stale` data will be upserted based on primary keys and stale data will be deleted by deleting any data fetched by the same previous source config.
- In `overwrite` data will be upserted based on primary keys and it will be up to the user to setup recurring task to delete stale data. 
- In `append`, old rows are never deleted or updated - every sync adds new rows.

In `overwrite` and `append` mode, you can distinguish between rows from different syncs by inspecting the `_cq_sync_time` column.