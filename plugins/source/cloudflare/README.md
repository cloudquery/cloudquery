## Cloudflare Plugin

The CloudQuery Cloudflare plugin pulls configuration out of Cloudflare resources, normalizes them and stores them in a PostgreSQL database.

### Links

- [Configuration](./docs/configuration.md)
- [Tables](./docs/tables/README.md)

### Authentication

In order to fetch information from Cloudflare, `cloudquery` needs to be authenticated. There are a few options for authentication:

- Export the `CLOUDFLARE_API_TOKEN` environment variable before running `cloudquery sync`.
- Export the `CLOUDFLARE_EMAIL` and `CLOUDFLARE_API_KEY` environment variables before running cloudquery
- Specifying either the `api_token` or `api_email, api_key` parameters in the YAML configuration (See [Configuration](./docs/configuration.md) for more details).

## Query Examples

### Find all zones with `dev_mode` enabled

```sql
SELECT id, account_id, host_name, name, original_ns FROM cloudflare_zones WHERE dev_mode = true;
```

### Find all DNS records

```sql
SELECT id, account_id, zone_id, name, type FROM cloudflare_dns_records;
```
