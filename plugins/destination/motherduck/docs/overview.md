---
name: MotherDuck
stage: Preview
title: MotherDuck Destination Plugin
description: CloudQuery MotherDuck destination plugin documentation
---
# MotherDuck Destination Plugin

:badge

This destination plugin lets you sync data from a CloudQuery source to [MotherDuck](https://motherduck.com/).

## Example Config

:configuration

## MotherDuck Spec

This is the top level spec used by the MotherDuck destination Plugin.

- `connection_string` (`string`) (required)

  Name of the database and extra connection options, such as `my_db`.

- `token` (`string`) (optional)

  MotherDuck API token. If empty, the plugin will open a web browser to authenticate.

- `batch_size` (`integer`) (optional) (default: `1000`)

  Maximum number of items that may be grouped together to be written in a single write.

- `batch_size_bytes` (`integer`) (optional) (default: `4194304` (4 MiB))

  Maximum size of items that may be grouped together to be written in a single write.

- `debug` (`boolean`) (optional) (default: `false`)

  Enables debug logging.
