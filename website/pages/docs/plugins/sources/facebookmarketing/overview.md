---
name: Facebook Marketing
stage: GA
title: Facebook Marketing Source Plugin
description: CloudQuery Facebook Marketing Plugin documentation
---

# Facebook Marketing Source Plugin

:badge

The Facebook Marketing source plugin for CloudQuery extracts information from the [Facebook marketing API](https://developers.facebook.com/docs/marketing-api/reference/v16.0).

## Configuration

This following configuration example connects a Facebook Marketing source to a Postgres destination.

:configuration

## Authentication

:authentication

## Facebook Marketing Spec

- `concurrency` (int, optional, default: 10000):
  A best effort maximum number of Go routines to use. Lower this number to reduce memory usage.
