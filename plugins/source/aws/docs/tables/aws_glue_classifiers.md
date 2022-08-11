
# Table: aws_glue_classifiers
Classifiers are triggered during a crawl task
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource|
|region|text|The AWS Region of the resource|
|name|text|Name of the classifier|
|csv_classifier_name|text|The name of the classifier|
|csv_classifier_allow_single_column|boolean|Enables the processing of files that contain only one column|
|csv_classifier_contains_header|text|Indicates whether the CSV file contains a header|
|csv_classifier_creation_time|timestamp without time zone|The time that this classifier was registered|
|csv_classifier_delimiter|text|A custom symbol to denote what separates each column entry in the row|
|csv_classifier_disable_value_trimming|boolean|Specifies not to trim values before identifying the type of column values|
|csv_classifier_header|text[]|A list of strings representing column names|
|csv_classifier_last_updated|timestamp without time zone|The time that this classifier was last updated|
|csv_classifier_quote_symbol|text|A custom symbol to denote what combines content into a single column value|
|csv_classifier_version|bigint|The version of this classifier|
|grok_classifier_classification|text|An identifier of the data format that the classifier matches, such as Twitter, JSON, Omniture logs, and so on|
|grok_classifier_grok_pattern|text|The grok pattern applied to a data store by this classifier|
|grok_classifier_name|text|The name of the classifier|
|grok_classifier_creation_time|timestamp without time zone|The time that this classifier was registered|
|grok_classifier_custom_patterns|text|Optional custom grok patterns defined by this classifier|
|grok_classifier_last_updated|timestamp without time zone|The time that this classifier was last updated|
|grok_classifier_version|bigint|The version of this classifier|
|json_classifier_json_path|text|A JsonPath string defining the JSON data for the classifier to classify|
|json_classifier_name|text|The name of the classifier|
|json_classifier_creation_time|timestamp without time zone|The time that this classifier was registered|
|json_classifier_last_updated|timestamp without time zone|The time that this classifier was last updated|
|json_classifier_version|bigint|The version of this classifier|
|xml_classifier_classification|text|An identifier of the data format that the classifier matches|
|xml_classifier_name|text|The name of the classifier|
|xml_classifier_creation_time|timestamp without time zone|The time that this classifier was registered|
|xml_classifier_last_updated|timestamp without time zone|The time that this classifier was last updated|
|xml_classifier_row_tag|text|The XML tag designating the element that contains each record in an XML document being parsed|
|xml_classifier_version|bigint|The version of this classifier|
