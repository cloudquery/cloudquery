# Overview

CloudQuery is an open-source cloud asset inventory powered by SQL, and as such, when providers change their schema (change/remove columns) some migrations from the previous run are required. CloudQuery automatically drops and recreates those tables automatically as needed.

## Running

### Sync (Upgrade/Downgrade) providers

The following command will upgrade or downgrade the provider to the version defined in our `cloudquery.yml`, if the version is defined as `latest` the latest version will be downloaded and the provider migrated to the latest schema version.

```bash
cloudquery provider sync aws
```

### Drop providers schema

The following command will drop a providers tables. Running CloudQuery fetch after a drop command will result in a recreation of all tables.

```bash
cloudquery provider drop aws
```

## Fetch Auto upgrade

CloudQuery automatically attempts to upgrade providers when fetch is executed, this action can be disabled by passing the `--skip-schema-upgrade` flag.
