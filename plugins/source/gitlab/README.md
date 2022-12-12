# Gitlab (Self-hosted) Plugin

This plugin pulls information from Gitlab (Self-hosted) instances and loads it into any supported CloudQuery destination (e.g. PostgreSQL).

## Links

- [Tables](./docs/tables/README.md)

## Authentication

In order to fetch information from Gitlab, `cloudquery` needs to be authenticated. A personal access token (PAT) is required for authentication. Follow [these steps](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html#create-a-personal-access-token/) to set generate one. Note that this token should only have read-only access to the resources you intend to use.

## Configuration

To configure CloudQuery to extract from Gitlab, create a `.yml` file in your CloudQuery configuration directory.
For example, the following configuration will extract information from Gitlab, and connect it to a `postgresql` destination plugin

```yaml
kind: source
spec:
  # Source spec section
  name: gitlab
  path: cloudquery/gitlab
  version: "0.0.1" # latest version of gitlab plugin
  tables: ["*"]
  destinations: ["postgresql"]
  spec:
    token: <Personal_Access_Token>
    base_url: <Instance_Url>
```

## Example

You can reduce alert fatigue by narrowing alerts down from gitlab using fuzzy matching.

```sql
-- select * from gitlab_____ where resources like ('%filter_here%');
```