
# Table: aws_fsx_volume_ontap_configuration
The configuration of an Amazon FSx for NetApp ONTAP volume
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|volume_cq_id|uuid|Unique CloudQuery ID of aws_fsx_volumes table (FK)|
|flex_cache_endpoint_type|text|Specifies the FlexCache endpoint type of the volume|
|junction_path|text|Specifies the directory that network-attached storage (NAS) clients use to mount the volume, along with the storage virtual machine (SVM) Domain Name System (DNS) name or IP address|
|volume_type|text|Specifies the type of volume|
|security_style|text|The security style for the volume, which can be UNIX, NTFS, or MIXED|
|size_in_megabytes|bigint|The configured size of the volume, in megabytes (MBs)|
|storage_efficiency_enabled|boolean|The volume's storage efficiency setting|
|storage_virtual_machine_id|text|The ID of the volume's storage virtual machine|
|storage_virtual_machine_root|boolean|A Boolean flag indicating whether this volume is the root volume for its storage virtual machine (SVM)|
|tiering_policy_cooling_period|bigint|Specifies the number of days that user data in a volume must remain inactive before it is considered "cold" and moved to the capacity pool|
|tiering_policy_name|text|Specifies the tiering policy used to transition data|
|uuid|text|The volume's universally unique identifier (UUID)|
