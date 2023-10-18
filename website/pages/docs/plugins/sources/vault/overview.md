---
name: Vault
stage: GA
title: Vault Source Plugin
description: CloudQuery Vault source plugin documentation
---
# Vault Source Plugin

:badge

The CloudQuery Vault plugin pulls data from Vault and loads it into any supported CloudQuery destination (e.g. PostgreSQL, BigQuery, Snowflake, and [more](/docs/plugins/destinations/overview)).

## Authentication

:authentication

## Example

This example syncs from Vault to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

:configuration

# Configuration Reference

This is the (nested) spec used by the Vault source plugin:

- `vault_address` (string, required\*):
  The is the address of the Vault server. This should be a complete URL (including the port) such as `"http://localhost:8200"`

- `concurrency` (integer, optional. Default: 10000):
  Maximum number of concurrent requests to the Vault server.

# Query Examples

## List the current auth methods with lease times

```sql copy
SELECT 
  path, type, config->'default_lease_ttl' as default_lease_ttl, config->'max_lease_ttl' as max_lease_ttl 
FROM 
  vault_sys_auths;
```
