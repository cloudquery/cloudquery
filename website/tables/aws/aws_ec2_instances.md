# Table: aws_ec2_instances

This table shows data for Amazon Elastic Compute Cloud (EC2) Instances.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Instance.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|state_transition_reason_time|`timestamp[us, tz=UTC]`|
|tags|`json`|
|ami_launch_index|`int64`|
|architecture|`utf8`|
|block_device_mappings|`json`|
|boot_mode|`utf8`|
|capacity_reservation_id|`utf8`|
|capacity_reservation_specification|`json`|
|client_token|`utf8`|
|cpu_options|`json`|
|current_instance_boot_mode|`utf8`|
|ebs_optimized|`bool`|
|elastic_gpu_associations|`json`|
|elastic_inference_accelerator_associations|`json`|
|ena_support|`bool`|
|enclave_options|`json`|
|hibernation_options|`json`|
|hypervisor|`utf8`|
|iam_instance_profile|`json`|
|image_id|`utf8`|
|instance_id|`utf8`|
|instance_lifecycle|`utf8`|
|instance_type|`utf8`|
|ipv6_address|`utf8`|
|kernel_id|`utf8`|
|key_name|`utf8`|
|launch_time|`timestamp[us, tz=UTC]`|
|licenses|`json`|
|maintenance_options|`json`|
|metadata_options|`json`|
|monitoring|`json`|
|network_interfaces|`json`|
|outpost_arn|`utf8`|
|placement|`json`|
|platform|`utf8`|
|platform_details|`utf8`|
|private_dns_name|`utf8`|
|private_dns_name_options|`json`|
|private_ip_address|`utf8`|
|product_codes|`json`|
|public_dns_name|`utf8`|
|public_ip_address|`utf8`|
|ramdisk_id|`utf8`|
|root_device_name|`utf8`|
|root_device_type|`utf8`|
|security_groups|`json`|
|source_dest_check|`bool`|
|spot_instance_request_id|`utf8`|
|sriov_net_support|`utf8`|
|state|`json`|
|state_reason|`json`|
|state_transition_reason|`utf8`|
|subnet_id|`utf8`|
|tpm_support|`utf8`|
|usage_operation|`utf8`|
|usage_operation_update_time|`timestamp[us, tz=UTC]`|
|virtualization_type|`utf8`|
|vpc_id|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### EC2 instances should not use multiple ENIs

```sql
WITH
  data
    AS (
      SELECT
        account_id, instance_id, count(nics->>'Status') AS cnt
      FROM
        aws_ec2_instances
        LEFT JOIN jsonb_array_elements(aws_ec2_instances.network_interfaces)
            AS nics ON true
      GROUP BY
        account_id, instance_id
    )
SELECT
  'EC2 instances should not use multiple ENIs' AS title,
  account_id,
  instance_id AS resource_id,
  CASE WHEN cnt > 1 THEN 'fail' ELSE 'pass' END AS status
FROM
  data;
```

### EC2 instances should not have a public IP address

```sql
SELECT
  'EC2 instances should not have a public IP address' AS title,
  account_id,
  instance_id AS resource_id,
  CASE WHEN public_ip_address IS NOT NULL THEN 'fail' ELSE 'pass' END AS status
FROM
  aws_ec2_instances;
```

### Security group is not currently in use so it should be deleted

```sql
WITH
  interface_groups
    AS (
      SELECT
        DISTINCT g->>'GroupId' AS id
      FROM
        aws_ec2_instances AS i,
        jsonb_array_elements(network_interfaces) AS a,
        jsonb_array_elements(a->'Groups') AS g
    )
SELECT
  'security group is not currently in use so it should be deleted' AS title,
  account_id,
  arn AS resource_id,
  CASE WHEN interface_groups.id IS NULL THEN 'fail' ELSE 'pass' END AS status
FROM
  aws_ec2_security_groups
  LEFT JOIN interface_groups ON
      aws_ec2_security_groups.group_id = interface_groups.id;
```

### EC2 instances should use IMDSv2

```sql
SELECT
  'EC2 instances should use IMDSv2' AS title,
  account_id,
  instance_id AS resource_id,
  CASE
  WHEN metadata_options->>'HttpTokens' IS DISTINCT FROM 'required' THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_ec2_instances;
```

### Find all ec2 instances that have unrestricted access to the internet with a wide open security group and routing

```sql
-- Find all AWS instances that are in a subnet that includes a catchall route
SELECT
  'Find all ec2 instances that have unrestricted access to the internet with a wide open security group and routing'
    AS title,
  account_id,
  instance_id AS resource_id,
  'fail' AS status
FROM
  aws_ec2_instances
WHERE
  subnet_id
  IN (
      SELECT
        a->>'SubnetId'
      FROM
        aws_ec2_route_tables AS t,
        jsonb_array_elements(t.associations) AS a,
        jsonb_array_elements(t.routes) AS r
      WHERE
        r->>'DestinationCidrBlock' = '0.0.0.0/0'
        OR r->>'DestinationIpv6CidrBlock' = '::/0'
    )
  AND instance_id
    IN (
        SELECT
          instance_id
        FROM
          aws_ec2_instances,
          jsonb_array_elements(security_groups) AS sg
          INNER JOIN view_aws_security_group_egress_rules ON id = sg->>'GroupId'
        WHERE
          ip = '0.0.0.0/0' OR ip6 = '::/0'
      );
```

### All ec2 instances that have unrestricted access to the internet via a security group

```sql
-- Find all AWS instances that have a security group that allows unrestricted egress
SELECT
  'All ec2 instances that have unrestricted access to the internet via a security group'
    AS title,
  aws_ec2_instances.account_id,
  instance_id AS resource_id,
  'fail' AS status
FROM
  aws_ec2_instances,
  jsonb_array_elements(security_groups) AS sg
  INNER JOIN view_aws_security_group_egress_rules ON id = sg->>'GroupId'
WHERE
  ip = '0.0.0.0/0' OR ip6 = '::/0';
```

### Find all instances with a public IP address

```sql
SELECT
  'Find all instances with a public IP address' AS title,
  account_id,
  arn AS resource_id,
  CASE WHEN public_ip_address IS NOT NULL THEN 'fail' ELSE 'pass' END AS status
FROM
  aws_ec2_instances;
```

### Unused EC2 security group

```sql
WITH
  interface_groups
    AS (
      SELECT
        DISTINCT a->>'GroupId' AS group_id
      FROM
        aws_ec2_instances, jsonb_array_elements(security_groups) AS a
    )
SELECT
  'Unused EC2 security group' AS title,
  sg.account_id,
  sg.arn AS resource_id,
  'fail' AS status
FROM
  aws_ec2_security_groups AS sg
  LEFT JOIN interface_groups ON interface_groups.group_id = sg.group_id
WHERE
  interface_groups.group_id IS NULL;
```

### Stopped EC2 instances should be removed after a specified time period

```sql
SELECT
  'Stopped EC2 instances should be removed after a specified time period'
    AS title,
  account_id,
  instance_id AS resource_id,
  CASE
  WHEN state->>'Name' = 'stopped'
  AND now() - state_transition_reason_time > '30'::INTERVAL DAY
  THEN 'fail'
  ELSE 'pass'
  END
FROM
  aws_ec2_instances;
```

### Amazon EC2 instances should be managed by AWS Systems Manager

```sql
SELECT
  'Amazon EC2 instances should be managed by AWS Systems Manager' AS title,
  aws_ec2_instances.account_id,
  aws_ec2_instances.arn AS resource_id,
  CASE
  WHEN aws_ssm_instances.instance_id IS NULL THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_ec2_instances
  LEFT JOIN aws_ssm_instances ON
      aws_ec2_instances.instance_id = aws_ssm_instances.instance_id;
```


