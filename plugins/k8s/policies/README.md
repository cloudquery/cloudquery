# CloudQuery Policies
CloudQuery SQL Policies for Kubernetes

## Policies and Compliance Frameworks Available

- [Kubernetes NSA CISA v1](./nsa_cisa_v1/policy.sql)

## Running

You can execute policies with `psql`. For example:

```bash
# Execute the whole CISA Policy
psql -U postgres -f  ./nsa_cisa_v1/policy.sql
```

This will create all the results in `k8s_policy_results` table which you can query directly, connect to any BI system (Grafana, Preset, AWS QuickSight, PowerBI, ...).

You can also output it into CSV or HTML with the following built-in psql commands:

```
# default tabular output
psql -U postgres -c "select * from k8s_policy_results"
# CSV output
psql -U postgres -c "select * from k8s_policy_results" --csv
# HTML output
psql -U postgres -c "select * from k8s_policy_results" --html
```
