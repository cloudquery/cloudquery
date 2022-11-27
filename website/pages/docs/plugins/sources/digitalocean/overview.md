# DigitalOcean Plugin

The CloudQuery DigitalOcean plugin pulls configuration from DigitalOcean.

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
  version: "VERSION_SOURCE_DIGITALOCEAN"
  tables: ["*"]
  destinations: ["postgresql"]
```

## Query Examples

### Find public facing spaces

```sql
--  public facing spaces are accessible by anyone, easily query which space is public facing in your account
SELECT bucket->>'Name',location,public FROM digitalocean_spaces WHERE public = true;
```

### List Droplets with public facing ipv4 or ipv6

```sql
-- Find any droplets that have a public ipv6 or ipv4 IP
SELECT id, name, v4->>'ip_address' AS address_v4, v4->>'netmask' AS netmask_v4, v4->>'gateway' AS gateway_v4,
       v6->>'ip_address' AS address_v6, v6->>'netmask' AS netmask_v6, v6->>'gateway' AS gateway_v6
FROM 
  (SELECT id,name,v4,NULL as v6 FROM digitalocean_droplets CROSS JOIN JSONB_ARRAY_ELEMENTS(digitalocean_droplets.networks->'v4') AS v4 
  UNION
  SELECT id,name,NULL as v4,v6 FROM digitalocean_droplets CROSS JOIN JSONB_ARRAY_ELEMENTS(digitalocean_droplets.networks->'v6') AS v6) AS union_v46
WHERE v4->>'type' = 'public' OR v6->>'type' = 'public';
```