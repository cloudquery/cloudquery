---
name: Notion
stage: Preview
title: Notion Source Plugin
description: CloudQuery Notion source plugin documentation
---

# Notion Source Plugin

:badge

A Notion source plugin for CloudQuery that loads data from notion to any database, data warehouse or data lake supported by [CloudQuery](https://www.cloudquery.io/), such as PostgreSQL, BigQuery, Athena, and many more.

## Example Configuration

This example syncs from Notion to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

:configuration

## Authentication

:authentication

## Configuration Reference

This is the (nested) spec used by the Notion source plugin:

- `bearer_token` (string, required):

  The bearer token to use for authentication.
