---
name: AWS
stage: GA
title: AWS Source Plugin
description: CloudQuery AWS Source Plugin documentation
---

# AWS Source Plugin

:badge

The AWS Source plugin extracts information from many of the supported services by Amazon Web Services (AWS) and loads it into any supported CloudQuery destination (e.g. PostgreSQL, BigQuery, Snowflake, and [more](/docs/plugins/destinations/overview)).

## Authentication

:authentication

## Query Examples

### Find all public-facing load balancers

```sql copy
SELECT * FROM aws_elbv2_load_balancers WHERE scheme = 'internet-facing';
```

### Find all unencrypted RDS instances

```sql copy
SELECT * FROM aws_rds_clusters WHERE storage_encrypted IS FALSE;
```

### Find all S3 buckets that are permitted to be public

```sql copy
SELECT arn, region
FROM aws_s3_buckets
WHERE block_public_acls IS NOT TRUE
    OR block_public_policy IS NOT TRUE
    OR ignore_public_acls IS NOT TRUE
    OR restrict_public_buckets IS NOT TRUE
```
