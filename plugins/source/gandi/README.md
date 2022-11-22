# Gandi Plugin

The CloudQuery Gandi plugin pulls configuration out of Gandi resources and loads it into any supported CloudQuery destination (e.g. PostgreSQL).

## Links

- [Configuration](./docs/configuration.md)
- [Tables](./docs/tables/README.md)

## Authentication

In order to fetch information from Gandi, `cloudquery` needs to be authenticated. An API key is required for authentication. Get your API key from [Gandi's Account Settings Page](https://account.gandi.net/en/).

## Query Examples

### Get all domains

```sql
select * from gandi_domains;
```

### Inspect glue records for a specific domain

```sql
select * from gandi_domain_glue_records where fqdn = 'yourdomain.com';
```

### Inspect LiveDNS snapshots for a given domain

```sql
select count(1) as number_of_snapshots, max(created_at) as last_snapshot_at from gandi_livedns_snapshots where fqdn = 'yourdomain.com';
```
