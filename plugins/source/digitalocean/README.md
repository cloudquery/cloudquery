# DigitalOcean Plugin

The CloudQuery DigitalOcean plugin pulls configuration from DigitalOcean resources and loads it into any supported CloudQuery destination (e.g. PostgreSQL).

## Links

- [Tables](./docs/tables/README.md)

## Authentication

- [Obtain a DigitalOcean API authentication token](https://docs.digitalocean.com/reference/api/api-reference/#section/Authentication)
- [Obtain a Spaces API access key](https://cloud.digitalocean.com/settings/api/tokens?i=d6d4a6)

Set the following environment variables in your shell (with values from the previous steps):

- `DIGITALOCEAN_TOKEN`
- `SPACES_ACCESS_KEY_ID` and `SPACES_SECRET_ACCESS_KEY`

## Configuration

In order to get started with the DigitalOcean plugin, you need to create a YAML file in your CloudQuery configuration directory (e.g. named `digitalocean.yml`).

The following example sets up the DigitalOcean plugin, and connects it to a postgresql destination:

```yaml
kind: source
spec:
  # Source spec section
  name: digitalocean
  path: cloudquery/digitalocean
  version: "v2.2.21" # latest version of digitalocean plugin
  tables: ["*"]
  destinations: ["postgresql"]
```

## Query Examples

### Find public facing spaces

```sql
--  public facing spaces are accessible by anyone, easily query which space is public facing in your account
SELECT name, location, public, creation_date FROM digitalocean_spaces WHERE public = true;
```

### List Droplets with public facing ipv4 or ipv6

```sql
-- Find any droplets that have a public ipv6 or ipv4 IP
SELECT d.id as droplet_id, dnv4.ip_address as ip, dnv4.netmask, dnv4.gateway,  dnv6.ip_address as ipv6, dnv6.netmask as ipv6_netmask, dnv6.gateway as ipv6_gateway
	from digitalocean_droplets d 
LEFT JOIN digitalocean_droplet_networks_v4 dnv4 ON d.cq_id = dnv4.droplet_cq_id 
LEFT JOIN digitalocean_droplet_networks_v6 dnv6 ON d.cq_id = dnv6.droplet_cq_id where dnv4.type = 'public' OR dnv6.type = 'public';
```

### Billing History including current month balance

```sql
-- Get you current monthly balance and previous billing histories in one table
SELECT invoice_id as id, description, amount, "date" FROM digitalocean_billing_history
UNION 
SELECT'current' as id, 'current month balance' as description, month_to_date_usage as amount , generated_at as "date" FROM digitalocean_balance;
```
