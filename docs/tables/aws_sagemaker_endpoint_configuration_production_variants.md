
# Table: aws_sagemaker_endpoint_configuration_production_variants
Identifies a model that you want to host and the resources chosen to deploy for hosting it
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|endpoint_configuration_cq_id|uuid|Unique CloudQuery ID of aws_sagemaker_endpoint_configurations table (FK)|
|initial_instance_count|integer|Number of instances to launch initially.  This member is required.|
|instance_type|text|The ML compute instance type.  This member is required.|
|model_name|text|The name of the model that you want to host|
|variant_name|text|The name of the production variant.  This member is required.|
|accelerator_type|text|The size of the Elastic Inference (EI) instance to use for the production variant|
|core_dump_config_destination_s3_uri|text|The Amazon S3 bucket to send the core dump to.  This member is required.|
|core_dump_config_kms_key_id|text|The Amazon Web Services Key Management Service (Amazon Web Services KMS) key that Amazon SageMaker uses to encrypt the core dump data at rest using Amazon S3 server-side encryption|
|initial_variant_weight|float|Determines initial traffic distribution among all of the models that you specify in the endpoint configuration|
