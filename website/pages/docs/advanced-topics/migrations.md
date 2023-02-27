# Overview

This section describes how CloudQuery is dealing with schema changes in plugins.
The overall idea is to not have breaking changes, but rather always add columns, because it is common for users to build views on top which we don't want to break. Those migration tactics are usually implemented in the destination plugins as source plugins are database agnostic and just send back JSON objects.

CloudQuery has [two modes](/docs/reference/destination-spec#migrate_mode) of migrating to a new schema, `safe` which is supported by all destinations, and `forced` which is only supported by ClickHouse, MySQL, PostgreSQL, and SQLite at the moment.

The `safe` mode is the default and will not run migrations that would result in data loss, and will print an error instead. The `forced` mode will run migrations that may result in data loss and the migration should always succeed without errors.

The following table describes when each mode is used:

| Existing table schema change | Migration mode needed | Reasoning |
| --- | --- | --- |
| Adding a new column that is neither a **primary key** nor a **not null** column | `safe` | New syncs **can** succeed by adding the new column to an existing table |
| Removing a column that is neither a **primary key** nor a **not null** column | `safe` | New syncs **can** succeed by ignoring the column removal |
| Adding a new column that is a **primary key** or a **not null** column | `forced` | New syncs **can't** succeed without back-filling the data, or dropping and re-adding the table |
| Removing a column that is a **primary key** or a **not null** column | `forced` | New syncs **can't** succeed as the column will not be populated with data, so dropping and re-adding the table is required |
| Changing a column type | `forced` | New syncs **can't** succeed without casting existing data into the new type, which is not always possible and can have performance implications in production environments |
