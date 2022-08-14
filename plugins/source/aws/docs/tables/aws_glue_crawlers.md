
# Table: aws_glue_crawlers
Specifies a crawler program that examines a data source and uses classifiers to try to determine its schema
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|arn|text|ARN of the resource.|
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|tags|jsonb||
|classifiers|text[]|A list of UTF-8 strings that specify the custom classifiers that are associated with the crawler|
|configuration|text|Crawler configuration information|
|crawl_elapsed_time|bigint|If the crawler is running, contains the total time elapsed since the last crawl began|
|crawler_security_configuration|text|The name of the SecurityConfiguration structure to be used by this crawler|
|creation_time|timestamp without time zone|The time that the crawler was created|
|database_name|text|The name of the database in which the crawler's output is stored|
|description|text|A description of the crawler|
|lake_formation_configuration_account_id|text|Required for cross account crawls|
|lake_formation_configuration_use_lake_formation_credentials|boolean|Specifies whether to use Lake Formation credentials for the crawler instead of the IAM role credentials|
|last_crawl_error_message|text|If an error occurred, the error information about the last crawl|
|last_crawl_log_group|text|The log group for the last crawl|
|last_crawl_log_stream|text|The log stream for the last crawl|
|last_crawl_message_prefix|text|The prefix for a message about this crawl|
|last_crawl_start_time|timestamp without time zone|The time at which the crawl started|
|last_crawl_status|text|Status of the last crawl|
|last_updated|timestamp without time zone|The time that the crawler was last updated|
|lineage_configuration_crawler_lineage_settings|text|Specifies whether data lineage is enabled for the crawler|
|name|text|The name of the crawler|
|recrawl_behavior|text|Specifies whether to crawl the entire dataset again or to crawl only folders that were added since the last crawler run|
|role|text|The Amazon Resource Name (ARN) of an IAM role that's used to access customer resources, such as Amazon Simple Storage Service (Amazon S3) data|
|schedule_expression|text|A cron expression used to specify the schedule (see Time-Based Schedules for Jobs and Crawlers (https://docsawsamazoncom/glue/latest/dg/monitor-data-warehouse-schedulehtml) For example, to run something every day at 12:15 UTC, you would specify: cron(15 12 * * ? *)|
|schedule_state|text|The state of the schedule|
|schema_change_policy_delete_behavior|text|The deletion behavior when the crawler finds a deleted object|
|schema_change_policy_update_behavior|text|The update behavior when the crawler finds a changed schema|
|state|text|Indicates whether the crawler is running, or whether a run is pending|
|table_prefix|text|The prefix added to the names of tables that are created|
|version|bigint|The version of the crawler|
