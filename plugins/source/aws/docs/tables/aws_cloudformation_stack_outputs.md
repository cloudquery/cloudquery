
# Table: aws_cloudformation_stack_outputs
The Output data type.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|stack_cq_id|uuid|Unique CloudQuery ID of aws_cloudformation_stacks table (FK)|
|description|text|User defined description associated with the output.|
|export_name|text|The name of the export associated with the output.|
|output_key|text|The key associated with the output.|
|output_value|text|The value associated with the output.|
