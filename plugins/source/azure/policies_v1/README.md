# CloudQuery Policies

CloudQuery SQL Policies for Azure

## Policies and Compliance Frameworks Available

- [Azure CIS v1.3.0](./cis_v1.3.0/policy.sql)
- [Azure HIPAA HITRUST v9.2](./hipaa_hitrust_v9.2/policy.sql)

## Running

You can execute policies with `psql`. For example:

```bash
# Execute the whole CIS Policy
psql -U postgres -f  ./cis_v1.3.0/policy.sql
```

This will create all the results in `azure_policy_results` table which you can query directly, connect to any BI system (Grafana, Preset, AWS QuickSight, PowerBI, …).

You can also output it into CSV or HTML with the following built-in `psql` commands:

```bash
# default tabular output
psql -U postgres -c "select * from azure_policy_results"
# CSV output
psql -U postgres -c "select * from azure_policy_results" --csv
# HTML output
psql -U postgres -c "select * from azure_policy_results" --html
```
