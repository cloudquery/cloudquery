
# Table: k8s_batch_job_status_conditions
JobCondition describes current state of a job.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|job_cq_id|uuid|Unique CloudQuery ID of k8s_batch_jobs table (FK)|
|type|text|Type of job condition, Complete or Failed.|
|status|text|Status of the condition, one of True, False, Unknown.|
|reason|text|(brief) reason for the condition's last transition.|
|message|text|Human readable message indicating details about last transition.|
