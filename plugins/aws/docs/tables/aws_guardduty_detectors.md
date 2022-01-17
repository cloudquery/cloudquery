
# Table: aws_guardduty_detectors

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) for the resource.|
|id|text|The Unique Identifier of the Detector.|
|service_role|text|The GuardDuty service role.|
|status|text|The detector status.|
|created_at|timestamp without time zone|The timestamp of when the detector was created.|
|data_sources_cloud_trail_status|text|Describes whether CloudTrail is enabled as a data source for the detector.|
|data_sources_dns_logs_status|text|Denotes whether DNS logs is enabled as a data source.|
|data_sources_flow_logs_status|text|Denotes whether VPC flow logs is enabled as a data source.|
|data_sources_s3_logs_status|text|A value that describes whether S3 data event logs are automatically enabled for new members of the organization.|
|finding_publishing_frequency|text|The publishing frequency of the finding.|
|tags|jsonb|The tags of the detector resource.|
|updated_at|timestamp without time zone|The last-updated timestamp for the detector.|
