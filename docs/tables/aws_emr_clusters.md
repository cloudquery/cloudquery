
# Table: aws_emr_clusters
The summary description of the cluster.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|cluster_arn|text|The Amazon Resource Name of the cluster.|
|resource_id|text|The unique identifier for the cluster.|
|name|text|The name of the cluster.|
|normalized_instance_hours|integer|An approximation of the cost of the cluster, represented in m1.|
|outpost_arn|text|The Amazon Resource Name (ARN) of the Outpost where the cluster is launched.|
|status_state|text|The current state of the cluster.|
|status_state_change_reason_code|text|The programmatic code for the state change reason.|
|status_state_change_reason_message|text|The descriptive message for the state change reason.|
|status_timeline_creation_date_time|timestamp without time zone|The creation date and time of the cluster.|
|status_timeline_end_date_time|timestamp without time zone|The date and time when the cluster was terminated.|
|status_timeline_ready_date_time|timestamp without time zone|The date and time when the cluster was ready to run steps.|
