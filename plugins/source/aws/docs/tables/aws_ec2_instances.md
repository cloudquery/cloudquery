# Table: aws_ec2_instances

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Instance.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|state_transition_reason_time|Timestamp|
|ami_launch_index|Int|
|architecture|String|
|block_device_mappings|JSON|
|boot_mode|String|
|capacity_reservation_id|String|
|capacity_reservation_specification|JSON|
|client_token|String|
|cpu_options|JSON|
|ebs_optimized|Bool|
|elastic_gpu_associations|JSON|
|elastic_inference_accelerator_associations|JSON|
|ena_support|Bool|
|enclave_options|JSON|
|hibernation_options|JSON|
|hypervisor|String|
|iam_instance_profile|JSON|
|image_id|String|
|instance_id|String|
|instance_lifecycle|String|
|instance_type|String|
|ipv6_address|String|
|kernel_id|String|
|key_name|String|
|launch_time|Timestamp|
|licenses|JSON|
|maintenance_options|JSON|
|metadata_options|JSON|
|monitoring|JSON|
|network_interfaces|JSON|
|outpost_arn|String|
|placement|JSON|
|platform|String|
|platform_details|String|
|private_dns_name|String|
|private_dns_name_options|JSON|
|private_ip_address|String|
|product_codes|JSON|
|public_dns_name|String|
|public_ip_address|String|
|ramdisk_id|String|
|root_device_name|String|
|root_device_type|String|
|security_groups|JSON|
|source_dest_check|Bool|
|spot_instance_request_id|String|
|sriov_net_support|String|
|state|JSON|
|state_reason|JSON|
|state_transition_reason|String|
|subnet_id|String|
|tags|JSON|
|tpm_support|String|
|usage_operation|String|
|usage_operation_update_time|Timestamp|
|virtualization_type|String|
|vpc_id|String|