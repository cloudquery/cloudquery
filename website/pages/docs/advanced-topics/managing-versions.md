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

CloudQuery plugins are versioned independently of the CLI. Releases happen on a weekly schedule, using semantic versioning to indicate breaking schema changes (as described in [Source Plugin Release Stages](/docs/plugins/sources/overview#source-plugin-release-stages)). We recommend pinning plugin versions to avoid unexpected changes to your data model, and only upgrading to new versions when you need to take advantage of new features or bug fixes. That said, if you are okay with the risk of breaking changes (or able to use `migrate_mode: forced`), [this how-to guide](/how-to-guides/update-plugins-using-renovate) describes how to keep plugin versions up-to-date automatically using Renovate. In all cases, we recommend performing upgrades in a staging environment first before applying them to production.
