---
name: SQLite
stage: GA
title: SQLite Destination Plugin
description: CloudQuery SQLite destination plugin documentation
---
# SQLite Destination Plugin

:badge

This destination plugin lets you sync data from a CloudQuery source to a SQLite database. This can be useful for local data exploration as no other database or service is required.

## Example Config

:configuration

## SQLite Spec

This is the top level spec used by the SQLite destination Plugin.

- `connection_string` (`string`) (required)

  Path to a file, such as `./mydb.sql`.