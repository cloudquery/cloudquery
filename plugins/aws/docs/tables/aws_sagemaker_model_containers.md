
# Table: aws_sagemaker_model_containers
Describes the container, as part of model definition.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|model_cq_id|uuid|Unique CloudQuery ID of aws_sagemaker_model table (FK)|
|container_hostname|text|This parameter is ignored for models that contain only a PrimaryContainer|
|environment|jsonb|The environment variables to set in the Docker container|
|image|text|The path where inference code is stored|
|image_config_repository_access_mode|text|Set this to one of the following values:  * Platform - The model image is hosted in Amazon ECR.  * Vpc - The model image is hosted in a private Docker registry in your VPC.  This member is required.|
|image_config_repository_auth_config_repo_creds_provider_arn|text|The Amazon Resource Name (ARN) of an Amazon Web Services Lambda function that provides credentials to authenticate to the private Docker registry where your model image is hosted|
|mode|text|Whether the container hosts a single model or multiple models.|
|model_data_url|text|The S3 path where the model artifacts, which result from model training, are stored|
|model_package_name|text|The name or Amazon Resource Name (ARN) of the model package to use to create the model.|
|multi_model_config_model_cache_setting|text|Whether to cache models for a multi-model endpoint|
