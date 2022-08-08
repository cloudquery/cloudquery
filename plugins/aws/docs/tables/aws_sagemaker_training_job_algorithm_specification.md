
# Table: aws_sagemaker_training_job_algorithm_specification
Specifies the training algorithm to use in a CreateTrainingJob request
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|training_job_cq_id|uuid|Unique CloudQuery ID of aws_sagemaker_training_jobs table (FK)|
|training_input_mode|text|The training input mode that the algorithm supports|
|algorithm_name|text|The name of the algorithm resource to use for the training job|
|enable_sage_maker_metrics_time_series|boolean|To generate and save time-series metrics during training, set to true|
|metric_definitions|jsonb|A list of metric definition objects|
|training_image|text|The registry path of the Docker image that contains the training algorithm|
