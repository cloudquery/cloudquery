In order to store information in a PostgreSQL database, cloudquery needs to be authenticated. Insert the PostgreSQL connection string in the `connection_string` field in the sync config file.

Acceptable formats are:

- URI: `postgres://postgres:pass@localhost:5432/postgres?sslmode=disable`. Any special URI characters need to be percent-encoded.
- DSN: `"user=postgres password=pass+0-[word host=localhost port=5432 dbname=postgres sslmode=disable"`
