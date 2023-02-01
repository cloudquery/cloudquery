# Table: alicloud_ecs_instances

https://www.alibabacloud.com/help/en/elastic-compute-service/latest/describeinstances#t9865.html

The composite primary key for this table is (**account_id**, **region_id**, **instance_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|hostname|String|
|image_id|String|
|instance_type|String|
|auto_release_time|Timestamp|
|last_invoked_time|Timestamp|
|os_type|String|
|device_available|Bool|
|instance_network_type|String|
|registration_time|Timestamp|
|local_storage_amount|Int|
|network_type|String|
|intranet_ip|String|
|is_spot|Bool|
|instance_charge_type|String|
|machine_id|String|
|private_pool_options_id|String|
|cluster_id|String|
|instance_name|String|
|private_pool_options_match_criteria|String|
|deployment_set_group_no|Int|
|credit_specification|String|
|gpu_amount|Int|
|connected|Bool|
|invocation_count|Int|
|start_time|Timestamp|
|zone_id|String|
|internet_max_bandwidth_in|Int|
|internet_charge_type|String|
|host_name|String|
|status|String|
|cpu|Int|
|isp|String|
|os_version|String|
|spot_price_limit|Float|
|os_name|String|
|os_name_en|String|
|serial_number|String|
|region_id (PK)|String|
|io_optimized|Bool|
|internet_max_bandwidth_out|Int|
|resource_group_id|String|
|activation_id|String|
|instance_type_family|String|
|instance_id (PK)|String|
|deployment_set_id|String|
|gpu_spec|String|
|description|String|
|recyclable|Bool|
|sale_cycle|String|
|expired_time|Timestamp|
|internet_ip|String|
|memory|Int|
|creation_time|Timestamp|
|agent_version|String|
|key_pair_name|String|
|hpc_cluster_id|String|
|local_storage_capacity|Int|
|vlan_id|String|
|stopped_mode|String|
|spot_strategy|String|
|spot_duration|Int|
|deletion_protection|Bool|
|security_group_ids|JSON|
|inner_ip_address|JSON|
|public_ip_address|JSON|
|rdma_ip_address|JSON|
|image_options|JSON|
|dedicated_host_attribute|JSON|
|ecs_capacity_reservation_attr|JSON|
|hibernation_options|JSON|
|dedicated_instance_attribute|JSON|
|eip_address|JSON|
|metadata_options|JSON|
|cpu_options|JSON|
|vpc_attributes|JSON|
|network_interfaces|JSON|
|tags|JSON|
|operation_locks|JSON|