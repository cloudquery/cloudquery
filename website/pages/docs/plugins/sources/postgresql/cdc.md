# PostgreSQL CDC (Logical Replication) Configuration

## Overview

[Change Data Capture](https://en.wikipedia.org/wiki/Change_data_capture) (CDC) is a set of software design patterns used to determine and track the data that has changed so that action can be taken using the changed data.

In PostgreSQL CDC can be implemented using [Logical Replication](https://www.postgresql.org/docs/current/logical-replication.html).

Logical Replication is a feature of PostgreSQL that allows you to stream changes from a database to another database, file or custom handler. It is also used to keep a copy of a database up to date with the original database in an efficient manner (Internally it is implemented by the so called [Write-Ahead Log](https://www.postgresql.org/docs/current/wal-intro.html)).

In this document we won't go into details of how Logical Replication works internally, but we will show you how to enable it in number of environments and how to configure CloudQuery PostgreSQL source plugin that can stream the changes to any of [CloudQuery supported destinations](../../destinations/overview).

Also, CloudQuery source plugin streams the changes directly to any of CQ destinations without any need for additional infrastructure (e.g. Kafka, RabbitMQ, etc). This means the setup is much easier.

## PSQL Test

To test that the current PostgreSQL instance supports Logical Replication we can run the following command in PSQL:

```sql
SHOW wal_level;
```

or the following SQL query:

```sql
SELECT setting FROM pg_settings WHERE name='wal_level'
```

The default is `replication` but for CDC to work we need to set it to `logical` (This can only be done on the database startup).

## Docker Setup

If you are running it locally for testing purposes you can use the following command to docker to enable logical replication:

```bash
docker run -p 5432:5432 -e POSTGRES_PASSWORD=pass -d postgres:11 -c "wal_level=logical"
```

## Self Hosted

Change the `wal_level` to `logical` in the `postgresql.conf` file and restart the database.

## RDS

You will need to create a custom RDS parameter group, associate it with your RDS instance, set `rds.logical_replication` to `1` and restart the database.

See full resolution on [AWS documentation](https://aws.amazon.com/premiumsupport/knowledge-center/rds-postgresql-use-logical-replication/)

## AWS Aurora serverless (V2)

AWS Aurora serverless V1 doesn't support Logical Replication, but V2 does.

Similar to RDS, you will need to create a custom parameter group, associate it with your Aurora serverless instance, set `rds.logical_replication` to `1` and restart the database.

See full resolution on [AWS documentation](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Replication.Logical.html)


## GCP Cloud SQL (PostgreSQL)

You will need to set the `cloudsql.logical_deconging` configuration to on. See full documentation on [GCP documentation](https://cloud.google.com/sql/docs/postgres/replication/configure-logical-replication#configuring-your-postgresql-instance)

## Azure DB (PostgreSQL)

Please follow the Official Azure documentation on [how to enable logical replication](https://learn.microsoft.com/en-us/azure/postgresql/single-server/concepts-logical) via the CLI or Console (UI).
