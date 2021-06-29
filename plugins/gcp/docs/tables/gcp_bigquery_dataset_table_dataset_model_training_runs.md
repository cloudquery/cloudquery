
# Table: gcp_bigquery_dataset_table_dataset_model_training_runs
Training options used by this training run These options are mutable for subsequent training runs Default values are explicitly stored for options not specified in the input query of the first training run For subsequent training runs, any option not explicitly specified in the input query will be copied from the previous training run
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|dataset_table_id|uuid|Unique ID of gcp_bigquery_dataset_tables table (FK)|
|start_time|text|Training run start time in milliseconds since the epoch|
|state|text|Different state applicable for a training run IN PROGRESS: Training run is in progress FAILED: Training run ended due to a non-retryable failure SUCCEEDED: Training run successfully completed CANCELLED: Training run cancelled by the user|
|training_options_early_stop|boolean||
|training_options_l1_reg|float||
|training_options_l2_reg|float||
|training_options_learn_rate|float||
|training_options_learn_rate_strategy|text||
|training_options_line_search_init_learn_rate|float||
|training_options_max_iteration|bigint||
|training_options_min_rel_progress|float||
|training_options_warm_start|boolean||
