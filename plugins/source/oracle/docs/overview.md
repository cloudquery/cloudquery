---
name: Oracle
stage: GA
title: Oracle Source Plugin
description: CloudQuery Oracle source plugin documentation
---

# Oracle Source Plugin

:badge

The CloudQuery Oracle plugin extracts Oracle Cloud Infrastructure data (`oci`) and loads it into any supported CloudQuery destination (e.g. PostgreSQL, BigQuery, Snowflake, and [more](/docs/plugins/destinations/overview)). It is based on [the OCI Go SDK](https://github.com/oracle/oci-go-sdk) and the [Oracle Cloud REST API](https://docs.oracle.com/en-us/iaas/api/#/).

## Authentication

:authentication

## Configuration

In order to get started with the Oracle plugin, you need to create a YAML file in your CloudQuery configuration directory (e.g. named `oracle.yml`).

The following example sets up the Oracle plugin, and connects it to a postgresql destination:

:configuration

See [tables](/docs/plugins/sources/oracle/tables) for a full list of available tables.

### Oracle Spec

This is the (nested) spec used by Oracle Source Plugin

- `concurrency` (`integer`) (optional) (default: `10000`)

  The best effort maximum number of Go routines to use.
  Lower this number to reduce memory usage.

## Dedicated regions

[OCI Dedicated regions](https://docs.oracle.com/en-us/iaas/Content/General/Concepts/dedicatedregions.htm) can be accessed via the following procedure:

1. Specify the region to be used for the discovery as the dedicated region.
   The following options are available:
   * Set `OCI_CLI_region` environment variable
   * Set `region` value in the configuration file
 
2. Specify the dedicated region information:
   The following options are available:
   * Set `OCI_REGION_METADATA` [environment variable](https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdk_adding_new_region_endpoints.htm#SDK_Adding_Regions_Environment_Variable)
   * Add information to the [regions config file](https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdk_adding_new_region_endpoints.htm#SDK_Adding_Regions_Config_File)
