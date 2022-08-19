
# Table: aws_fsx_storage_vms
Describes the Amazon FSx for NetApp ONTAP storage virtual machine (SVM) configuration
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource|
|region|text|The AWS Region of the resource|
|ad_cfg_net_bios_name|text|The NetBIOS name of the Active Directory computer object that is joined to your SVM|
|ad_cfg_dns_ips|text[]|A list of up to three IP addresses of DNS servers or domain controllers in the self-managed AD directory|
|ad_cfg_domain_name|text|The fully qualified domain name of the self-managed AD directory|
|ad_cfg_file_system_administrators_group|text|The name of the domain group whose members have administrative privileges for the FSx file system|
|ad_cfg_organizational_unit_distinguished_name|text|The fully qualified distinguished name of the organizational unit within the self-managed AD directory to which the Windows File Server or ONTAP storage virtual machine (SVM) instance is joined|
|ad_cfg_user_name|text|The user name for the service account on your self-managed AD domain that FSx uses to join to your AD domain|
|creation_time|timestamp without time zone|The time that the resource was created, in seconds (since 1970-01-01T00:00:00Z), also known as Unix time|
|endpoints_iscsi_dns_name|text|The Domain Name Service (DNS) name for the file system|
|endpoints_iscsi_ip_addresses|text[]|The SVM endpoint's IP addresses|
|endpoints_management_dns_name|text|The Domain Name Service (DNS) name for the file system|
|endpoints_management_ip_addresses|text[]|The SVM endpoint's IP addresses|
|endpoints_nfs_dns_name|text|The Domain Name Service (DNS) name for the file system|
|endpoints_nfs_ip_addresses|text[]|The SVM endpoint's IP addresses|
|endpoints_smb_dns_name|text|The Domain Name Service (DNS) name for the file system|
|endpoints_smb_ip_addresses|text[]|The SVM endpoint's IP addresses|
|file_system_id|text|The globally unique ID of the file system, assigned by Amazon FSx|
|lifecycle|text|Describes the SVM's lifecycle status|
|lifecycle_transition_reason_message|text|A detailed error message|
|name|text|The name of the SVM, if provisioned|
|arn|text|The Amazon Resource Name (ARN) for a given resource|
|root_volume_security_style|text|The security style of the root volume of the SVM|
|id|text|The SVM's system generated unique ID|
|subtype|text|Describes the SVM's subtype|
|tags|jsonb|A list of Tag values, with a maximum of 50 elements|
|uuid|text|The SVM's UUID (universally unique identifier)|
