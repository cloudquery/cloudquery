---
name: Terraform
stage: GA
title: Terraform Source Plugin
description: CloudQuery Terraform source plugin documentation
---

# Terraform Source Plugin

:badge

The CloudQuery Terraform plugin extracts terraform state and loads it into any supported CloudQuery destination (e.g. PostgreSQL, BigQuery, Snowflake, and [more](/docs/plugins/destinations/overview)).

## Configuration

:configuration

## Authentication for S3 backends

:authentication

## Terraform Spec

This is the (nested) spec used by the Terraform source plugin:

- `concurrency` (int, optional, default: 10000):
  A best effort maximum number of Go routines to use. Lower this number to reduce memory usage.
