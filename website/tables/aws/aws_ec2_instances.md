# Table: aws_ec2_instances

This table shows data for Amazon Elastic Compute Cloud (EC2) Instances.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Instance.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id|utf8|
|region|utf8|
|arn (PK)|utf8|
|state_transition_reason_time|timestamp[us, tz=UTC]|
|tags|json|
|ami_launch_index|int64|
|architecture|utf8|
|block_device_mappings|json|
|boot_mode|utf8|
|capacity_reservation_id|utf8|
|capacity_reservation_specification|json|
|client_token|utf8|
|cpu_options|json|
|ebs_optimized|bool|
|elastic_gpu_associations|json|
|elastic_inference_accelerator_associations|json|
|ena_support|bool|
|enclave_options|json|
|hibernation_options|json|
|hypervisor|utf8|
|iam_instance_profile|json|
|image_id|utf8|
|instance_id|utf8|
|instance_lifecycle|utf8|
|instance_type|utf8|
|ipv6_address|utf8|
|kernel_id|utf8|
|key_name|utf8|
|launch_time|timestamp[us, tz=UTC]|
|licenses|json|
|maintenance_options|json|
|metadata_options|json|
|monitoring|json|
|network_interfaces|json|
|outpost_arn|utf8|
|placement|json|
|platform|utf8|
|platform_details|utf8|
|private_dns_name|utf8|
|private_dns_name_options|json|
|private_ip_address|utf8|
|product_codes|json|
|public_dns_name|utf8|
|public_ip_address|utf8|
|ramdisk_id|utf8|
|root_device_name|utf8|
|root_device_type|utf8|
|security_groups|json|
|source_dest_check|bool|
|spot_instance_request_id|utf8|
|sriov_net_support|utf8|
|state|json|
|state_reason|json|
|state_transition_reason|utf8|
|subnet_id|utf8|
|tpm_support|utf8|
|usage_operation|utf8|
|usage_operation_update_time|timestamp[us, tz=UTC]|
|virtualization_type|utf8|
|vpc_id|utf8|