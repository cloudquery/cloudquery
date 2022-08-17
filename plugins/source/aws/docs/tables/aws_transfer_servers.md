
# Table: aws_transfer_servers
Describes the properties of a file transfer protocol-enabled server that was specified
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|Specifies the unique Amazon Resource Name (ARN) of the server|
|certificate|text|Specifies the ARN of the Amazon Web ServicesCertificate Manager (ACM) certificate|
|domain|text|Specifies the domain of the storage system that is used for file transfers|
|endpoint_details_address_allocation_ids|text[]|A list of address allocation IDs that are required to attach an Elastic IP address to your server's endpoint|
|endpoint_details_security_group_ids|text[]|A list of security groups IDs that are available to attach to your server's endpoint|
|endpoint_details_subnet_ids|text[]|A list of subnet IDs that are required to host your server endpoint in your VPC This property can only be set when EndpointType is set to VPC|
|endpoint_details_vpc_endpoint_id|text|The ID of the VPC endpoint|
|endpoint_details_vpc_id|text|The VPC ID of the VPC in which a server's endpoint will be hosted|
|endpoint_type|text|Defines the type of endpoint that your server is connected to|
|host_key_fingerprint|text|Specifies the Base64-encoded SHA256 fingerprint of the server's host key|
|identity_provider_details_directory_id|text|The identifier of the Directory Service directory that you want to stop sharing|
|identity_provider_details_function|text|The ARN for a lambda function to use for the Identity provider|
|identity_provider_details_invocation_role|text|Provides the type of InvocationRole used to authenticate the user account|
|identity_provider_details_url|text|Provides the location of the service endpoint used to authenticate users|
|identity_provider_type|text|The mode of authentication for a server|
|logging_role|text|The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role that allows a server to turn on Amazon CloudWatch logging for Amazon S3 or Amazon EFSevents|
|post_authentication_login_banner|text|Specifies a string to display when users connect to a server|
|pre_authentication_login_banner|text|Specifies a string to display when users connect to a server|
|protocol_details_as2_transports|text[]|Indicates the transport method for the AS2 messages|
|protocol_details_passive_ip|text|Indicates passive mode, for FTP and FTPS protocols|
|protocol_details_set_stat_option|text|Use the SetStatOption to ignore the error that is generated when the client attempts to use SETSTAT on a file you are uploading to an S3 bucket|
|protocol_details_tls_session_resumption_mode|text|A property used with Transfer Family servers that use the FTPS protocol|
|protocols|text[]|Specifies the file transfer protocol or protocols over which your file transfer protocol client can connect to your server's endpoint|
|security_policy_name|text|Specifies the name of the security policy that is attached to the server|
|server_id|text|Specifies the unique system-assigned identifier for a server that you instantiate|
|state|text|The condition of the server that was described|
|tags|jsonb|Specifies the key-value pairs that you can use to search for and group servers that were assigned to the server that was described|
|user_count|bigint|Specifies the number of users that are assigned to a server you specified with the ServerId|
