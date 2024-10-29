---
title: Performance Tuning
description: Tips and tricks for improving the performance of `cloudquery sync` for large cloud estates.
---

# Performance Tuning

This page contains a number of tips and tricks for improving the performance of `cloudquery sync` for large cloud estates.

## Identifying Slow Tables

The first step in improving the performance of a sync is to identify which tables are taking the longest to sync. This can be done by running the `cloudquery sync` command with the `--tables-metrics-location` flag.
This flag takes a path to a file where the metrics will be written. The metrics are displayed as a human-readable table and refreshed as the sync progresses.
To use this feature run `cloudquery sync --tables-metrics-location metrics.txt` and open the file `metrics.txt` in a text editor that supports live updates. Alternatively, you can use `watch cat metrics.txt` to monitor the file in real-time.
You can see an example output below. Tables that are still in progress will be shown first with a `N/A` in the `END TIME` column, and the rest sorted by table name.
The `CLIENT ID` column differs between integrations and can be used to identify which account, region, project, location, etc. is being synced (the terms differ between integrations).

```text
+-----------------------------------------------------+-----------------------+---------------------------------------+---------------------------------------+------------+-----------+--------+--------+
| TABLE                                               | CLIENT ID             | START TIME                            | END TIME                              |   DURATION | RESOURCES | ERRORS | PANICS |
+-----------------------------------------------------+-----------------------+---------------------------------------+---------------------------------------+------------+-----------+--------+--------+
| gcp_compute_interconnect_locations                  | project:cq-playground | 2024-07-25 13:25:45.349231 +0100 WEST | N/A                                   | 13.398995s |         3 |      0 |      0 |
| gcp_compute_zones                                   | project:cq-playground | 2024-07-25 13:25:45.34586 +0100 WEST  | N/A                                   | 13.402371s |         9 |      0 |      0 |
| gcp_compute_addresses                               | project:cq-playground | 2024-07-25 13:25:45.345863 +0100 WEST | 2024-07-25 13:25:46.189411 +0100 WEST |  843.548ms |         0 |      0 |      0 |
| gcp_compute_autoscalers                             | project:cq-playground | 2024-07-25 13:25:45.345865 +0100 WEST | 2024-07-25 13:25:46.004631 +0100 WEST |  658.766ms |         0 |      0 |      0 |
| gcp_compute_backend_buckets                         | project:cq-playground | 2024-07-25 13:25:45.34589 +0100 WEST  | 2024-07-25 13:25:45.855523 +0100 WEST |  509.633ms |         1 |      0 |      0 |
| gcp_compute_backend_services                        | project:cq-playground | 2024-07-25 13:25:45.345819 +0100 WEST | 2024-07-25 13:25:45.996089 +0100 WEST |   650.27ms |         6 |      0 |      0 |
| gcp_compute_disk_types                              | project:cq-playground | 2024-07-25 13:25:45.348244 +0100 WEST | 2024-07-25 13:25:47.491563 +0100 WEST |  2.143319s |      1074 |      0 |      0 |
| gcp_compute_disks                                   | project:cq-playground | 2024-07-25 13:25:45.348809 +0100 WEST | 2024-07-25 13:25:46.058995 +0100 WEST |  710.186ms |         0 |      0 |      0 |
| gcp_compute_external_vpn_gateways                   | project:cq-playground | 2024-07-25 13:25:45.348643 +0100 WEST | 2024-07-25 13:25:45.800374 +0100 WEST |  451.731ms |         0 |      0 |      0 |
| gcp_compute_firewalls                               | project:cq-playground | 2024-07-25 13:25:45.347932 +0100 WEST | 2024-07-25 13:25:45.800827 +0100 WEST |  452.895ms |         0 |      0 |      0 |
| gcp_compute_forwarding_rules                        | project:cq-playground | 2024-07-25 13:25:45.348903 +0100 WEST | 2024-07-25 13:25:46.012584 +0100 WEST |  663.681ms |         2 |      0 |      0 |
| gcp_compute_global_addresses                        | project:cq-playground | 2024-07-25 13:25:45.348121 +0100 WEST | 2024-07-25 13:25:45.799297 +0100 WEST |  451.176ms |         0 |      0 |      0 |
| gcp_compute_health_checks                           | project:cq-playground | 2024-07-25 13:25:45.348224 +0100 WEST | 2024-07-25 13:25:46.036481 +0100 WEST |  688.257ms |         1 |      0 |      0 |
| gcp_compute_images                                  | project:cq-playground | 2024-07-25 13:25:45.345884 +0100 WEST | 2024-07-25 13:25:45.832301 +0100 WEST |  486.417ms |         0 |      0 |      0 |
| gcp_compute_instance_groups                         | project:cq-playground | 2024-07-25 13:25:45.347338 +0100 WEST | 2024-07-25 13:25:46.062304 +0100 WEST |  714.966ms |         0 |      0 |      0 |
| gcp_compute_instances                               | project:cq-playground | 2024-07-25 13:25:45.348076 +0100 WEST | 2024-07-25 13:25:46.152832 +0100 WEST |  804.756ms |         0 |      0 |      0 |
| gcp_compute_interconnect_attachments                | project:cq-playground | 2024-07-25 13:25:46.238178 +0100 WEST | 2024-07-25 13:25:54.035826 +0100 WEST |  7.797648s |         0 |      0 |      0 |
| gcp_compute_interconnect_remote_locations           | project:cq-playground | 2024-07-25 13:25:45.349069 +0100 WEST | 2024-07-25 13:25:45.869177 +0100 WEST |  520.108ms |        69 |      0 |      0 |
| gcp_compute_interconnects                           | project:cq-playground | 2024-07-25 13:25:45.349127 +0100 WEST | 2024-07-25 13:25:45.799541 +0100 WEST |  450.414ms |         0 |      0 |      0 |
| gcp_compute_machine_types                           | project:cq-playground | 2024-07-25 13:25:45.857729 +0100 WEST | 2024-07-25 13:25:54.233986 +0100 WEST |  8.376257s |      2121 |      0 |      0 |
| gcp_compute_network_endpoint_groups                 | project:cq-playground | 2024-07-25 13:25:45.34896 +0100 WEST  | 2024-07-25 13:25:46.042662 +0100 WEST |  693.702ms |         4 |      0 |      0 |
| gcp_compute_networks                                | project:cq-playground | 2024-07-25 13:25:45.349043 +0100 WEST | 2024-07-25 13:25:45.865695 +0100 WEST |  516.652ms |         1 |      0 |      0 |
| gcp_compute_osconfig_inventories                    | project:cq-playground | 2024-07-25 13:25:45.857659 +0100 WEST | 2024-07-25 13:25:54.049883 +0100 WEST |  8.192224s |         0 |      0 |      0 |
| gcp_compute_osconfig_os_patch_deployments           | project:cq-playground | 2024-07-25 13:25:45.345874 +0100 WEST | 2024-07-25 13:25:46.012176 +0100 WEST |  666.302ms |         1 |      0 |      0 |
| gcp_compute_osconfig_os_patch_jobs_instance_details | project:cq-playground | 2024-07-25 13:25:45.995674 +0100 WEST | 2024-07-25 13:25:48.640675 +0100 WEST |  2.645001s |         3 |      0 |      0 |
| gcp_compute_osconfig_os_patch_jobs                  | project:cq-playground | 2024-07-25 13:25:45.348153 +0100 WEST | 2024-07-25 13:25:48.640696 +0100 WEST |  3.292543s |         3 |      0 |      0 |
| gcp_compute_osconfig_os_policy_assignment_reports   | project:cq-playground | 2024-07-25 13:25:45.857781 +0100 WEST | 2024-07-25 13:25:54.189788 +0100 WEST |  8.332007s |         0 |      0 |      0 |
| gcp_compute_osconfig_os_policy_assignments          | project:cq-playground | 2024-07-25 13:25:45.857757 +0100 WEST | 2024-07-25 13:25:54.159643 +0100 WEST |  8.301886s |         2 |      0 |      0 |
| gcp_compute_osconfig_os_vulnerability_reports       | project:cq-playground | 2024-07-25 13:25:45.857441 +0100 WEST | 2024-07-25 13:25:54.145726 +0100 WEST |  8.288285s |         0 |      0 |      0 |
| gcp_compute_projects                                | project:cq-playground | 2024-07-25 13:25:45.34828 +0100 WEST  | 2024-07-25 13:25:45.875861 +0100 WEST |  527.581ms |         1 |      0 |      0 |
| gcp_compute_routers                                 | project:cq-playground | 2024-07-25 13:25:45.348312 +0100 WEST | 2024-07-25 13:25:46.082812 +0100 WEST |    734.5ms |         0 |      0 |      0 |
| gcp_compute_routes                                  | project:cq-playground | 2024-07-25 13:25:45.345797 +0100 WEST | 2024-07-25 13:25:45.779476 +0100 WEST |  433.679ms |        43 |      0 |      0 |
| gcp_compute_security_policies                       | project:cq-playground | 2024-07-25 13:25:45.347514 +0100 WEST | 2024-07-25 13:25:45.932167 +0100 WEST |  584.653ms |         0 |      0 |      0 |
| gcp_compute_ssl_certificates                        | project:cq-playground | 2024-07-25 13:25:45.347976 +0100 WEST | 2024-07-25 13:25:46.045754 +0100 WEST |  697.778ms |         2 |      0 |      0 |
| gcp_compute_ssl_policies                            | project:cq-playground | 2024-07-25 13:25:45.347981 +0100 WEST | 2024-07-25 13:25:45.858365 +0100 WEST |  510.384ms |         1 |      0 |      0 |
| gcp_compute_subnetworks                             | project:cq-playground | 2024-07-25 13:25:45.345729 +0100 WEST | 2024-07-25 13:25:46.182778 +0100 WEST |  837.049ms |        42 |      0 |      0 |
| gcp_compute_target_http_proxies                     | project:cq-playground | 2024-07-25 13:25:45.34819 +0100 WEST  | 2024-07-25 13:25:46.060256 +0100 WEST |  712.066ms |         2 |      0 |      0 |
| gcp_compute_target_instances                        | project:cq-playground | 2024-07-25 13:25:45.347976 +0100 WEST | 2024-07-25 13:25:46.035427 +0100 WEST |  687.451ms |         0 |      0 |      0 |
| gcp_compute_target_pools                            | project:cq-playground | 2024-07-25 13:25:45.348173 +0100 WEST | 2024-07-25 13:25:46.055082 +0100 WEST |  706.909ms |         0 |      0 |      0 |
| gcp_compute_target_ssl_proxies                      | project:cq-playground | 2024-07-25 13:25:45.345808 +0100 WEST | 2024-07-25 13:25:45.774172 +0100 WEST |  428.364ms |         0 |      0 |      0 |
| gcp_compute_target_vpn_gateways                     | project:cq-playground | 2024-07-25 13:25:45.347995 +0100 WEST | 2024-07-25 13:25:46.033584 +0100 WEST |  685.589ms |         0 |      0 |      0 |
| gcp_compute_url_maps                                | project:cq-playground | 2024-07-25 13:25:45.345904 +0100 WEST | 2024-07-25 13:25:45.996447 +0100 WEST |  650.543ms |         3 |      0 |      0 |
| gcp_compute_vpn_gateways                            | project:cq-playground | 2024-07-25 13:25:45.347537 +0100 WEST | 2024-07-25 13:25:46.003861 +0100 WEST |  656.324ms |         0 |      0 |      0 |
| gcp_compute_vpn_tunnels                             | project:cq-playground | 2024-07-25 13:25:45.348019 +0100 WEST | 2024-07-25 13:25:47.007435 +0100 WEST |  1.659416s |         0 |      0 |      0 |
+-----------------------------------------------------+-----------------------+---------------------------------------+---------------------------------------+------------+-----------+--------+--------+
```

:::callout{type="info"}
This feature is available starting from CLI version [v5.25.0](https://github.com/cloudquery/cloudquery/releases/tag/cli-v5.25.0) and integrations released on July 10th 2024 or later.
:::

## Use Wildcard Matching

Sometimes the easiest way to improve the performance of the `sync` command is to limit the number of tables that get synced. The `tables` and `skip_tables` source configuration options both support wildcard matching. This means that you can use `*` anywhere in a name to match multiple tables.

For example, when using the `aws` source integration, it is possible to use a wildcard pattern to match all tables related to AWS EC2:

```yaml copy
tables:
 - aws_ec2_*
```

This can also be combined with `skip_tables`. For example, let's say we want to include all EC2 tables, but not EBS-related ones:

```yaml copy
tables: 
- "aws_ec2_*"
skip_tables:
- "aws_ec2_ebs_*"
```

:::callout{type="info"}
The CloudQuery CLI will warn if a wildcard pattern does not match any known tables.
:::

## Tune Concurrency

The `concurrency` setting, available for all source integrations as part of the [source spec](/docs/reference/source-spec#concurrency), controls the approximate number of concurrent requests that will be made while performing a sync. Setting this to a low number will reduce the number of concurrent requests, reducing the memory used and making the sync less likely to hit rate limits. The trade-off is that syncs will take longer to complete.

## Adjust Batch Size

Most destination integrations have batching related settings that can be adjusted to improve performance. Tuning these can improve performance, but it can also increase the memory usage of the sync process. Here are the batching related settings you will come across:

- `batch_size`: The number of rows that are inserted into the destination at once. The default value for this setting is usually between 1000 to 10000 rows, depending on the destination integration.

- `batch_size_bytes`: Maximum size of items that may be grouped together to be written in a single write. This is useful for limiting the memory usage of the sync process. The default value for this varies between 4 MB to 100 MB, depending on the destination integration.

- `batch_timeout`: Maximum interval between batch writes. Even if data stops coming in, the batch will be written after this interval. The default value for this setting is usually between 10 seconds and 1 minute, depending on the destination integration.

Some destination integrations (such as file or S3 destinations) start a new object or file for every batch, and some simply buffer the data in memory to be written at once.

:::callout{type="info"}
You should check the documentation for the destination integration you are using to see what the default values are and consider how they can be adjusted to suit your use case.
:::

Here's a conservative example for the PostgreSQL destination integration that reduces the overall memory usage, but may also increase the time it takes to sync:

```yaml
kind: destination
spec:
  name: "postgresql"
  path: "cloudquery/postgresql"
  registry: "cloudquery"
  version: "VERSION_DESTINATION_POSTGRESQL"
  spec:
    connection_string: "postgres://user:pass@localhost:5432/mydb?sslmode=disable" # replace with your connection string
    batch_size: 10000 # 10000 rows, default
    batch_size_bytes: 4194304 # 4 MB, dramatically tuned down from the 100 MB default
    batch_timeout: "30s" # 30 seconds, tuned down from 60 seconds
```

With this configuration, the PostgreSQL destination integration will write 10,000 rows at a time, or 4 MB of data at a time, or every 30 seconds, whichever comes first.

## Use a Different Scheduler

By default, CloudQuery syncs will fetch all tables in parallel, writing data to the destination(s) as they come in. However, the `concurrency` setting, mentioned above, places a limit on how many **table-clients** can be synced at a time. What "table-client" means depends on the source integration and the table. In AWS, for example, a client is usually a combination of account and region. Get all the combinations of accounts and regions for all tables, and you have all the table-clients for a sync. For the GCP source integration, clients generally map to projects.

The default CloudQuery scheduler, known as `dfs`, will sync up to `concurrency / 100` table-clients at a time (we are ignoring child relations for the purposes of this discussion). Let's take an example GCP cloud estate with 5000 projects, syncing 100 tables. This makes for approximately 500,000 table-client pairs, and a concurrency of 10,000 will allow 100 table-client pairs to be synced at a time. The `dfs` scheduler will start with the first table and its first 100 projects, and then move on to finish all projects for that table before moving on to the next table. This means, in practice, only one table is really being synced at a time!

Usually this works out fine, as long as the cloud platform's rate limits are aligned with the clients. But if rate limits are applied per-table, rather than per-project, `dfs` can be suboptimal. A better strategy in this case would be to choose the first client for every table before moving on to the next client. This is what the `round-robin` scheduler does.

Only some integrations support this setting. The following example configuration enables `round-robin` scheduling for the GCP source integration:

```yaml
kind: source
spec:
  name: "gcp"
  path: "cloudquery/gcp"
  registry: "cloudquery"
  version: "VERSION_SOURCE_GCP"
  tables: ["gcp_storage_*", "gcp_compute_*"]
  destinations: ["postgresql"]
  spec:
    scheduler: "round-robin"
    project_ids: ...
```

Finally, the `shuffle` strategy aims to provide a balance between `dfs` and `round-robin` by randomizing the order in which table-client pairs are chosen. The following example enables `shuffle` for the GCP integration, which can help reduce the likelihood of hitting rate limits by randomly mixing the underlying services to which API calls that are made concurrently, rather than hitting a single API with all calls at once:

```yaml
kind: source
spec:
  name: "gcp"
  path: "cloudquery/gcp"
  registry: "cloudquery"
  version: "VERSION_SOURCE_GCP"
  tables: ["gcp_storage_*", "gcp_compute_*"]
  destinations: ["postgresql"]
  spec:
    project_ids: ...
    scheduler: "shuffle"
    # ...
```

:::callout{type="warn"}
The `shuffle` scheduler is the **default** for the AWS source integration.
:::

## Avoid `skip_dependent_tables: false`

Starting with version [v6.0.0](https://github.com/cloudquery/cloudquery/releases/tag/cli-v6.0.0) of the CloudQuery CLI `skip_dependent_tables` is set to `true` by default, to avoid new tables implicitly being synced when added to integrations. This can be overridden by setting `skip_dependent_tables: false` in the source configuration.

When setting `skip_dependent_tables: false`, all tables that depend on other tables will be synced by default.
When syncing dependent tables multiple API calls need to be made for every row in the parent table. This can lead to thousands of API calls, increasing the time it takes to sync.

Let's say we have three tables: `A`, `B` and `C`. `A` is the top-level table. `B` depends on it, and `C` depends on `B`:

```text copy
A
↳ B
  ↳ C
```

By default only `A` will be synced. If you set `skip_dependent_tables: false`, `B` and `C` will also be synced. This can be a problem if `A` has a large number of rows, as it will result in a large number of API calls to sync `B` and `C`.

To avoid setting `skip_dependent_tables: false` and still get dependent tables synced, you can explicitly list the dependent tables in the source configuration, or [use wildcard matching](#use-wildcard-matching).