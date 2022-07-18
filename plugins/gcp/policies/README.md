# CloudQuery Policies
CloudQuery SQL Policies for GCP

## Policies and Compliance Frameworks Available

- [GCP CIS v1.2.0](./cis_v1.2.0/policy.sql)

## Running

You can execute policies with `psql`. For example:

```bash
# Execute the whole CIS Policy
psql -U postgres -f  ./cis_v1.2.0/policy.sql
```

This will create all the results in `gcp_policy_results` table which you can query directly, connect to any BI system (Grafana, Preset, AWS QuickSight, PowerBI, ...).

You can also output it into CSV or HTML with the following built-in psql commands:

```
# default tabular output
psql -U postgres -c "select * from gcp_policy_results"
# CSV output
psql -U postgres -c "select * from gcp_policy_results" --csv
# HTML output
psql -U postgres -c "select * from gcp_policy_results" --html
```
