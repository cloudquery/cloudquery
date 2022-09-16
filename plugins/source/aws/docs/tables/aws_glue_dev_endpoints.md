
# Table: aws_glue_dev_endpoints
A development endpoint where a developer can remotely debug extract, transform, and load (ETL) scripts
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) of the workflow.|
|tags|jsonb|Resource tags.|
|arguments|jsonb|A map of arguments used to configure the DevEndpoint|
|availability_zone|text|The AWS Availability Zone where this DevEndpoint is located|
|created_timestamp|timestamp without time zone|The point in time at which this DevEndpoint was created|
|name|text|The name of the DevEndpoint|
|extra_jars_s3_path|text|The path to one or more Java jar files in an S3 bucket that should be loaded in your DevEndpoint|
|extra_python_libs_s3_path|text|The paths to one or more Python libraries in an Amazon S3 bucket that should be loaded in your DevEndpoint|
|failure_reason|text|The reason for a current failure in this DevEndpoint|
|glue_version|text|Glue version determines the versions of Apache Spark and Python that Glue supports|
|last_modified_timestamp|timestamp without time zone|The point in time at which this DevEndpoint was last modified|
|last_update_status|text|The status of the last update|
|number_of_nodes|bigint|The number of Glue Data Processing Units (DPUs) allocated to this DevEndpoint|
|number_of_workers|bigint|The number of workers of a defined workerType that are allocated to the development endpoint|
|private_address|text|A private IP address to access the DevEndpoint within a VPC if the DevEndpoint is created within one|
|public_address|text|The public IP address used by this DevEndpoint|
|public_key|text|The public key to be used by this DevEndpoint for authentication|
|public_keys|text[]|A list of public keys to be used by the DevEndpoints for authentication|
|role_arn|text|The Amazon Resource Name (ARN) of the IAM role used in this DevEndpoint|
|security_configuration|text|The name of the SecurityConfiguration structure to be used with this DevEndpoint|
|security_group_ids|text[]|A list of security group identifiers used in this DevEndpoint|
|status|text|The current status of this DevEndpoint|
|subnet_id|text|The subnet ID for this DevEndpoint|
|vpc_id|text|The ID of the virtual private cloud (VPC) used by this DevEndpoint|
|worker_type|text|The type of predefined worker that is allocated to the development endpoint Accepts a value of Standard, G1X, or G2X|
|yarn_endpoint_address|text|The YARN endpoint address used by this DevEndpoint|
|zeppelin_remote_spark_interpreter_port|bigint|The Apache Zeppelin port for the remote Apache Spark interpreter|
