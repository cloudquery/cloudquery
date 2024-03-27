---
name: GitHub
stage: GA
title: GitHub Source Plugin
description: CloudQuery GitHub source plugin documentation
---

# GitHub Source Plugin

:badge

The CloudQuery GitHub plugin extracts your GitHub API and loads it into any supported CloudQuery destination (e.g. PostgreSQL, BigQuery, Snowflake, and [more](/docs/plugins/destinations/overview)).

## Authentication

:authentication

## Configuration

:configuration

## GitHub Spec

This is the (nested) spec used by GitHub Source Plugin

- `repos` (`[]string`, optional. Default: empty):
  List of repositories to sync from. The format is `owner/repo` (e.g. `cloudquery/cloudquery`). You must specify either `orgs` or `repos` in the configuration.

- `orgs` (`[]string`, optional. Default: empty):
  List of organizations to sync from. You must specify either `orgs` or `repos` in the configuration.

- `concurrency` (int, optional, default: 1500):
  The best effort maximum number of Go routines to use. Lower this number to reduce memory usage or to avoid hitting GitHub API rate limits.

- `discovery_concurrency` (`int`) (default: `1`)

  During initialization the GitHub source plugin discovers all repositories under the organizations configured in `orgs`, to be used later on during the sync process.
  By default the plugin discovers repositories one organization at a time. You can increase `discovery_concurrency` to discover multiple organizations in parallel, or use a negative value to discover all organizations in parallel.
  Please note that it's possible to hit GitHub API rate limits when using a high value for `discovery_concurrency`.

- `include_archived_repos` (`bool`) (default: `false`)

  By default archived repositories are not included in the sync. To include archived repositories set `include_archived_repos` to `true`.

- `local_cache_path` (`string`, optional, default: empty):
  Path to a local cache directory. If set, the plugin will cache the GitHub API responses in this directory. Defaults to an empty string (no cache).
  By using a cache, the plugin can use [conditional requests when appropriate](https://docs.github.com/en/rest/using-the-rest-api/best-practices-for-using-the-rest-api?#use-conditional-requests-if-appropriate), and help avoid hitting GitHub API rate limits.
