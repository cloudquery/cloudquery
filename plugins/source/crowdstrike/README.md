# CrowdStrike Plugin

This plugin pulls information from CrowdStrike and loads it into any supported CloudQuery destination (e.g. PostgreSQL).

## Links

- [Tables](./docs/tables/README.md)

## Authentication

In order to fetch information from CrowdStrike, `cloudquery` needs to be authenticated. A client id and secret is required for authentication. Follow [these steps](https://www.crowdstrike.com/blog/tech-center/get-access-falcon-apis/) to set these up. Note that you will also need to enlist the client to have the appropriate scope for what you want to query.

## Configuration

To configure CloudQuery to extract from CrowdStrike, create a `.yml` file in your CloudQuery configuration directory.
For example, the following configuration will extract information from CrowdStrike, and connect it to a `postgresql` destination plugin

```yaml
kind: source
spec:
  # Source spec section
  name: crowdstrike
  path: cloudquery/crowdstrike
  version: "0.0.1" # latest version of crowdstrike plugin
  tables: ["*"]
  destinations: ["postgresql"]
  spec:
    client_id: <CLIENT_ID>
    client_secret: <CLIENT_SECRET>
```

## Example

You can reduce alert fatigue by narrowing alerts down from CrowdStrike using fuzzy matching.

```sql
select * from crowdstrike_alerts_query where resources like ('%filter_here%');
```