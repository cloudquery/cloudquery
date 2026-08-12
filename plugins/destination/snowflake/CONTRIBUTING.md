# Contribution Guide to CloudQuery Snowflake Destination Plugin

## Running in Debug mode

Similar to all other CQ plugins you can run the plugin in debug mode via:

```bash
go run main.go serve
```

## Testing

The Snowflake tests hit a live Snowflake account, so they cost money and need
credentials. They are disabled by default and only run when `CQ_SNOWFLAKE_TEST`
is set to `1`.

When enabled, the tests authenticate with a keypair (JWT). The DSN is assembled
at runtime from the following `SNOW_*` environment variables:

| Variable           | Required | Description                                                                                     |
| ------------------ | -------- | ----------------------------------------------------------------------------------------------- |
| `CQ_SNOWFLAKE_TEST`| yes      | Gate. Set to `1` to enable the tests. Left unset, all Snowflake tests are skipped.              |
| `SNOW_ACCOUNT`     | yes      | Account identifier (e.g. `<org>-<account>`).                                                     |
| `SNOW_USER`        | yes      | Username the keypair is registered against.                                                     |
| `SNOW_PRIVATE_KEY` | yes      | PEM-encoded PKCS#8 RSA private key (the public half is registered on the user).                 |
| `SNOW_DATABASE`    | yes      | Database to run against. Must already exist and be accessible to the user's role.               |
| `SNOW_SCHEMA`      | yes      | Schema to run against. Must already exist (the plugin does not create it).                      |
| `SNOW_WAREHOUSE`   | yes      | Warehouse used to run the queries.                                                              |
| `SNOW_REGION`      | no       | Region. Leave **unset** when `SNOW_ACCOUNT` already embeds the region, otherwise the driver errors on a region conflict. |

Example invocation (placeholders only — substitute your own values or export the
variables beforehand):

```bash
export CQ_SNOWFLAKE_TEST=1
export SNOW_ACCOUNT=<your-account-identifier>
export SNOW_USER=<your-user>
export SNOW_PRIVATE_KEY="$(cat /path/to/rsa_key.p8)"
export SNOW_DATABASE=<your-database>
export SNOW_SCHEMA=<your-schema>
export SNOW_WAREHOUSE=<your-warehouse>
# export SNOW_REGION=<your-region>   # only if SNOW_ACCOUNT does not embed the region
go test ./client/...
```

Or inline:

```bash
CQ_SNOWFLAKE_TEST=1 SNOW_ACCOUNT=... SNOW_USER=... SNOW_PRIVATE_KEY="$(cat rsa_key.p8)" \
  SNOW_DATABASE=... SNOW_SCHEMA=... SNOW_WAREHOUSE=... go test ./client/...
```

### Required Snowflake privileges

The JWT DSN does not pass a role, so Snowflake uses the user's `DEFAULT_ROLE`
(or `PUBLIC`). Grant the following to that role so the tests can create and use
the file format, stage, and tables the plugin needs. The database, schema, and
warehouse must already exist; the plugin creates tables, a stage, and a file
format inside the schema but never creates the database or schema itself.

```sql
GRANT USAGE ON WAREHOUSE <warehouse> TO ROLE <role>;
GRANT USAGE ON DATABASE <database> TO ROLE <role>;
GRANT USAGE ON SCHEMA <database>.<schema> TO ROLE <role>;
GRANT CREATE TABLE ON SCHEMA <database>.<schema> TO ROLE <role>;
GRANT CREATE STAGE ON SCHEMA <database>.<schema> TO ROLE <role>;
GRANT CREATE FILE FORMAT ON SCHEMA <database>.<schema> TO ROLE <role>;

GRANT ROLE <role> TO USER <user>;
ALTER USER <user> SET DEFAULT_ROLE = <role>;
```

## Lint

```bash
make lint
```
