---
name: MongoDB Atlas
stage: GA (Premium)
title: MongoDB Atlas Source Plugin
description: CloudQuery MongoDB Atlas source plugin documentation
---
# MongoDB Atlas Source Plugin

:badge{text="Premium"}

This is a premium plugin that you can buy [here](/integrations/mongodb-atlas).

The CloudQuery MongoDB Atlas plugin extracts information from your [MongoDB Atlas API](https://www.mongodb.com/docs/atlas/api/) and loads it into any supported CloudQuery destination (e.g. PostgreSQL, BigQuery, Snowflake, and [more](/docs/plugins/destinations/overview)).

## Authentication

:authentication

## Configuration

:configuration

:::callout{type="info"}
Make sure you use environment variable expansion in production instead of committing the credentials to the configuration file directly.
:::

## MongoDB Atlas Spec

This is the (nested) spec used by this plugin:

- `api_key` (string, optional):
   
   If empty, extracted from MONGODB_ATLAS_PUBLIC_KEY

- `api_secret` (string, optional):

  If empty, extracted from MONGODB_ATLAS_PRIVATE_KEY

- `base_url` (string, optional):

  If empty, extracted from MONGODB_ATLAS_URL. Default: https://cloud.mongodb.com
