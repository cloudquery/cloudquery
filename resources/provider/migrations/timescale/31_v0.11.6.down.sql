-- Resource: aws.regions
ALTER TABLE IF EXISTS "aws_regions" DROP COLUMN IF EXISTS "partition";

-- Resource: athena.data_catalogs
DROP TABLE IF EXISTS aws_athena_data_catalog_database_table_columns;
DROP TABLE IF EXISTS aws_athena_data_catalog_database_table_partition_keys;
DROP TABLE IF EXISTS aws_athena_data_catalog_database_tables;
DROP TABLE IF EXISTS aws_athena_data_catalog_databases;
DROP TABLE IF EXISTS aws_athena_data_catalogs;

-- Resource: athena.work_groups
DROP TABLE IF EXISTS aws_athena_work_group_prepared_statements;
DROP TABLE IF EXISTS aws_athena_work_group_query_executions;
DROP TABLE IF EXISTS aws_athena_work_group_named_queries;
DROP TABLE IF EXISTS aws_athena_work_groups;
