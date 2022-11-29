# AWS Source Plugin Recipes

Full spec options for the AWS source plugin are available [here](/docs/plugins/sources/aws/configuration#aws-spec).

```yaml copy
kind: source
spec:
  name: aws
  path: cloudquery/aws
  version: "VERSION_SOURCE_AWS"
  tables: ["*"]

  # Comment out any of the following tables if you want to sync them
  # unless otherwise indicated they are configuration parameters rather than configured resources
  skip_tables:
    - aws_ec2_vpc_endpoint_services # this resource includes services that are available from AWS as well as other AWS Accounts
    - aws_docdb_cluster_parameter_groups
    - aws_docdb_engine_versions
    - aws_ec2_instance_types
    - aws_elasticache_engine_versions
    - aws_elasticache_parameter_groups
    - aws_elasticache_reserved_cache_nodes_offerings
    - aws_elasticache_service_updates
    - aws_neptune_cluster_parameter_groups
    - aws_neptune_db_parameter_groups
    - aws_rds_cluster_parameter_groups
    - aws_rds_db_parameter_groups
    - aws_rds_engine_versions
    - aws_servicequotas_services
  destinations: ["<destination>"]
```
