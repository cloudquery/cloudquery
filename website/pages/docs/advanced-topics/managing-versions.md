---
title: Managing Versions
description: CloudQuery plugins are versioned independently of the CLI. Releases happen on a weekly schedule, using semantic versioning to indicate breaking schema changes.
---

# Managing Versions

## Managing CLI Versions

### Downloading the CLI

All CloudQuery CLI versions are available for download on the [releases page](https://github.com/cloudquery/cloudquery/releases?q=cli-&expanded=true).

### Homebrew

To update the CLI via Homebrew, run:

```bash
brew upgrade cloudquery/tap/cloudquery
```

To install a specific version of the CLI, run:

```bash
brew install cloudquery/tap/cloudquery@<version>
``` 

(e.g. `brew install cloudquery/tap/cloudquery@2.3.10`)

## Managing Plugin Versions

CloudQuery plugins are versioned independently of the CLI. Releases happen on a weekly schedule, using semantic versioning to indicate breaking schema changes as described in [Source Plugin Release Stages](#source-plugin-release-stages). We recommend pinning plugin versions to avoid unexpected changes to your data model, and only upgrading to new versions when you need to take advantage of new features or bug fixes. That said, if you are okay with the risk of breaking changes (or able to use `migrate_mode: forced`), [this how-to guide](https://www.cloudquery.io/blog/update-plugins-using-renovate) describes how to keep plugin versions up-to-date automatically using Renovate. In all cases, we recommend performing upgrades in a staging environment first before applying them to production.

### Semantic Versioning

For version `Major.Minor.Patch`:

- `Major` is incremented when there are breaking changes (e.g. breaking configuration spec structure, column type changes).
- `Minor` is incremented when we add features in a backwards compatible way.
- `Patch` is incremented when we fix bugs in a backwards compatible way.

## Source Plugin Release Stages

[Official source plugins](https://hub.cloudquery.io/plugins/source?authors=official) go through two release stages: Preview and Generally Available (GA).

Both Preview and GA plugins follow [semantic versioning](#semantic-versioning).

The main differences between the two stages are:

1. Preview plugins are experimental and may have frequent breaking changes.
2. Preview plugins might get deprecated due to lack of usage.
3. Premium plugins in Preview are free to use.
4. Long Term Support and bug fixes are only guaranteed for plugins that are Generally Available.
