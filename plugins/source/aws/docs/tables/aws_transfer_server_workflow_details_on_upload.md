
# Table: aws_transfer_server_workflow_details_on_upload
Specifies the workflow ID for the workflow to assign and the execution role that's used for executing the workflow
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|server_cq_id|uuid|Unique CloudQuery ID of aws_transfer_servers table (FK)|
|execution_role|text|Includes the necessary permissions for S3, EFS, and Lambda operations that Transfer can assume, so that all workflow steps can operate on the required resources|
|workflow_id|text|A unique identifier for the workflow|
