# CloudQuery Policies
CloudQuery SQL Policies for GCP

## Policies and Compliance Frameworks Available

- [GCP CIS v1.2.0](./cis_v1.2.0/policy.sql)

## Running

You can execute policies with `psql`. For example:

```bash
# Execute CIS V1.2.0 Policy
psql ${DSN} -f  ./cis_v1.2.0/policy.sql
```

This will create all the results in `gcp_policy_results` table which you can query directly, connect to any BI system (Grafana, Preset, AWS QuickSight, PowerBI, …).

You can also output it into CSV or HTML with the following built-in `psql` commands:

```bash
# Set DSN to your PostgreSQL populated by CloudQuery
export DSN=postgres://postgres:pass@localhost:5432/postgres

# default tabular output
psql ${DSN} -c "select * from gcp_policy_results"
# CSV output
psql ${DSN} -c "select * from gcp_policy_results" --csv
# HTML output
psql ${DSN} -c "select * from gcp_policy_results" --html

# see only failed checks
psql ${DSN} -c "select * from gcp_policy_results where status='fail'"
```

## Dashboards

Currently we have pre-built dashboards on top of `gcp_policy_results` table which is available [here](../dashboards/grafana/compliance.json).

<img alt="Azure Asset Inventory Grafana Dashboard" src="../dashboards/grafana/compliance.png" width=50% height=50%>

See [installation instructions](../dashboards/README.md#gcp-compliance-and-cspm-cloud-security-posture-management-dashboard)