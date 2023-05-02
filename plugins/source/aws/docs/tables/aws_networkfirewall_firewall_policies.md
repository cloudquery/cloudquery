# Table: aws_networkfirewall_firewall_policies

https://docs.aws.amazon.com/network-firewall/latest/APIReference/API_FirewallPolicy.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|tags|JSON|
|firewall_policy_response|JSON|
|update_token|String|
|firewall_policy|JSON|