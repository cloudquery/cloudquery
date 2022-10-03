---
title: FAQ
---

# FAQ

## Does CloudQuery access any application data in my cloud?

No. CloudQuery only accesses metadata and configuration data. It never pulls data from your application databases or cloud storage files.

### What happens when I run two (or more) sync? Will the second sync remove resources that no longer exist from the database?

There are currently two types of `write` modes in destination plugins: `overwrite`, `overwrite-delete-stale` and `append`, if in append data will always be added and never deleted. If in `overwrite` data will be upserted based on primary keys and it will be up to the user to setup recurring task to delete stale data. If in `overwrite-delete-stale` data will be upserted based on primary keys and stale data will be deleted by deleting any data fetched by the same previous source config.
