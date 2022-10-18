---
title: Troubleshooting
---

# Troubleshooting

## Help Channels

### Discord

First things first - feel free to join our [Discord](https://www.cloudquery.io/discord)!

### GitHub Issues

There are a couple of ways to get help for any CloudQuery-related issues or questions.

1. Check out previous issues at [https://github.com/cloudquery/cloudquery](https://github.com/cloudquery/cloudquery) and open a new one if no previous one has been opened or resolved.
2. Reach out on the #help-and-support channel on [Discord](https://www.cloudquery.io/discord)

## Debugging

### Verbose Logging

Usually the first step that will be needed to debug/resolve an issue is to run `cloudquery` with `--log-level debug` to enable verbose logging.

### Error: "failed to migrate source"...

If you see an error such as `failed to migrate source`, it means that, while upgrading a plugin, the migration of the SQL schema failed.
CloudQuery makes a best-effort attempt to automatically and transparently manage the schemas of plugins, but 
this can sometimes fail during version upgrades.

The solution is to drop-and recreate the database or schema. (or all the plugins tables, such as `aws_*`).
Some queries that may help you (handle with care!):

```sql
DROP SCHEMA public CASCADE;
CREATE SCHEMA public;
```

```sql
\dt aws_*
```