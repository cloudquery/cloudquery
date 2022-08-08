
# Table: aws_sagemaker_notebook_instances
Provides summary information for an Amazon SageMaker notebook instance.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|network_interface_id|text|The network interface IDs that Amazon SageMaker created at the time of creating the instance.|
|kms_key_id|text|The Amazon Web Services KMS key ID Amazon SageMaker uses to encrypt data when storing it on the ML storage volume attached to the instance.|
|subnet_id|text|The ID of the VPC subnet.|
|volume_size_in_gb|integer|The size, in GB, of the ML storage volume attached to the notebook instance.|
|accelerator_types|text[]|A list of the Elastic Inference (EI) instance types associated with this notebook instance.|
|security_groups|jsonb|The IDs of the VPC security groups.|
|direct_internet_access|boolean|Describes whether Amazon SageMaker provides internet access to the notebook instance.|
|tags|jsonb|The tags associated with the notebook instance.|
|arn|text|The Amazon Resource Name (ARN) of the notebook instance. |
|name|text|The name of the notebook instance that you want a summary for. |
|additional_code_repositories|text[]|An array of up to three Git repositories associated with the notebook instance. These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in Amazon Web Services CodeCommit (https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository|
|creation_time|timestamp without time zone|A timestamp that shows when the notebook instance was created.|
|default_code_repository|text|The Git repository associated with the notebook instance as its default code repository|
|instance_type|text|The type of ML compute instance that the notebook instance is running on.|
|last_modified_time|timestamp without time zone|A timestamp that shows when the notebook instance was last modified.|
|notebook_instance_lifecycle_config_name|text|The name of a notebook instance lifecycle configuration associated with this notebook instance|
|notebook_instance_status|text|The status of the notebook instance.|
|url|text|The URL that you use to connect to the Jupyter instance running in your notebook instance.|
