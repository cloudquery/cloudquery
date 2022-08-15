
# Table: aws_fsx_filesystems
A description of a specific Amazon FSx file system.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|creation_time|timestamp without time zone|The time that the file system was created, in seconds (since 1970-01-01T00:00:00Z), also known as Unix time.|
|dns_name|text|The Domain Name System (DNS) name for the file system.|
|failure_details_message|text|A message describing any failures that occurred during file system creation.|
|id|text|The system-generated, unique 17-digit ID of the file system.|
|type|text|The type of Amazon FSx file system, which can be LUSTRE, WINDOWS, ONTAP, or OPENZFS.|
|version|text|The Lustre version of the Amazon FSx for Lustre file system, either 2.10 or 2.12.|
|kms_key_id|text|The ID of the Key Management Service (KMS) key used to encrypt Amazon FSx file system data|
|lifecycle|text|The lifecycle status of the file system|
|network_interface_ids|text[]|The IDs of the elastic network interfaces from which a specific file system is accessible|
|owner_id|text|The Amazon Web Services account that created the file system|
|arn|text|The Amazon Resource Name (ARN) of the file system resource.|
|storage_capacity|bigint|The storage capacity of the file system in gibibytes (GiB).|
|storage_type|text|The type of storage the file system is using|
|subnet_ids|text[]|Specifies the IDs of the subnets that the file system is accessible from|
|tags|jsonb|The tags to associate with the file system|
|vpc_id|text|The ID of the primary virtual private cloud (VPC) for the file system.|
