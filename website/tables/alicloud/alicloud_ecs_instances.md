# Table: alicloud_ecs_instances

This table shows data for Alibaba Cloud Elastic Compute Service (ECS) Instances.

https://www.alibabacloud.com/help/en/elastic-compute-service/latest/describeinstances#t9865.html

The composite primary key for this table is (**account_id**, **region_id**, **instance_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|hostname|`utf8`|
|image_id|`utf8`|
|instance_type|`utf8`|
|auto_release_time|`timestamp[us, tz=UTC]`|
|last_invoked_time|`timestamp[us, tz=UTC]`|
|os_type|`utf8`|
|device_available|`bool`|
|instance_network_type|`utf8`|
|registration_time|`timestamp[us, tz=UTC]`|
|local_storage_amount|`int64`|
|network_type|`utf8`|
|intranet_ip|`utf8`|
|is_spot|`bool`|
|instance_charge_type|`utf8`|
|machine_id|`utf8`|
|private_pool_options_id|`utf8`|
|cluster_id|`utf8`|
|socket_id|`utf8`|
|instance_name|`utf8`|
|private_pool_options_match_criteria|`utf8`|
|deployment_set_group_no|`int64`|
|credit_specification|`utf8`|
|gpu_amount|`int64`|
|connected|`bool`|
|invocation_count|`int64`|
|start_time|`timestamp[us, tz=UTC]`|
|zone_id|`utf8`|
|internet_max_bandwidth_in|`int64`|
|internet_charge_type|`utf8`|
|host_name|`utf8`|
|status|`utf8`|
|cpu|`int64`|
|isp|`utf8`|
|os_version|`utf8`|
|spot_price_limit|`float64`|
|os_name|`utf8`|
|instance_owner_id|`int64`|
|os_name_en|`utf8`|
|serial_number|`utf8`|
|region_id (PK)|`utf8`|
|io_optimized|`bool`|
|internet_max_bandwidth_out|`int64`|
|resource_group_id|`utf8`|
|activation_id|`utf8`|
|instance_type_family|`utf8`|
|instance_id (PK)|`utf8`|
|deployment_set_id|`utf8`|
|gpu_spec|`utf8`|
|description|`utf8`|
|recyclable|`bool`|
|sale_cycle|`utf8`|
|expired_time|`timestamp[us, tz=UTC]`|
|internet_ip|`utf8`|
|memory|`int64`|
|creation_time|`timestamp[us, tz=UTC]`|
|agent_version|`utf8`|
|key_pair_name|`utf8`|
|hpc_cluster_id|`utf8`|
|local_storage_capacity|`int64`|
|vlan_id|`utf8`|
|stopped_mode|`utf8`|
|spot_strategy|`utf8`|
|spot_duration|`int64`|
|deletion_protection|`bool`|
|security_group_ids|`json`|
|inner_ip_address|`json`|
|public_ip_address|`json`|
|rdma_ip_address|`json`|
|image_options|`json`|
|dedicated_host_attribute|`json`|
|ecs_capacity_reservation_attr|`json`|
|hibernation_options|`json`|
|dedicated_instance_attribute|`json`|
|eip_address|`json`|
|metadata_options|`json`|
|cpu_options|`json`|
|vpc_attributes|`json`|
|network_interfaces|`json`|
|tags|`json`|
|operation_locks|`json`|