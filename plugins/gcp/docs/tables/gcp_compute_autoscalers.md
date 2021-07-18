
# Table: gcp_compute_autoscalers
Represents an Autoscaler resource.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|cool_down_period_sec|bigint|The number of seconds that the autoscaler waits before it starts collecting information from a new instance This prevents the autoscaler from collecting information when the instance is initializing, during which the collected usage would not be reliable The default time autoscaler waits is 60 seconds  Virtual machine initialization times might vary because of numerous factors We recommend that you test how long an instance may take to initialize To do this, create an instance and time the startup process|
|cpu_utilization_predictive_method|text|Indicates whether predictive autoscaling based on CPU metric is enabled Valid values are:  * NONE (default) No predictive method is used The autoscaler scales the group to meet current demand based on real-time metrics * OPTIMIZE_AVAILABILITY Predictive autoscaling improves availability by monitoring daily and weekly load patterns and scaling out ahead of anticipated demand|
|cpu_utilization_utilization_target|float|The target CPU utilization that the autoscaler maintains Must be a float value in the range (0, 1] If not specified, the default is 06  If the CPU level is below the target utilization, the autoscaler scales in the number of instances until it reaches the minimum number of instances you specified or until the average CPU of your instances reaches the target utilization  If the average CPU is above the target utilization, the autoscaler scales out until it reaches the maximum number of instances you specified or until the average utilization reaches the target utilization|
|load_balancing_utilization_utilization_target|float|Fraction of backend capacity utilization (set in HTTP(S) load balancing configuration) that the autoscaler maintains Must be a positive float value If not defined, the default is 08|
|max_num_replicas|bigint|The maximum number of instances that the autoscaler can scale out to This is required when creating or updating an autoscaler The maximum number of replicas must not be lower than minimal number of replicas|
|min_num_replicas|bigint|The minimum number of replicas that the autoscaler can scale in to This cannot be less than 0 If not provided, autoscaler chooses a default value depending on maximum number of instances allowed|
|mode|text|Defines operating mode for this policy|
|scale_in_control_max_scaled_in_replicas_calculated|bigint|Absolute value of VM instances calculated based on the specific mode |
|scale_in_control_max_scaled_in_replicas_fixed|bigint|Specifies a fixed number of VM instances This must be a positive integer|
|scale_in_control_max_scaled_in_replicas_percent|bigint|Specifies a percentage of instances between 0 to 100%, inclusive For example, specify 80 for 80%|
|scale_in_control_time_window_sec|bigint|How far back autoscaling looks when computing recommendations to include directives regarding slower scale in, as described above|
|scaling_schedules|jsonb|Scaling schedules defined for an autoscaler Multiple schedules can be set on an autoscaler, and they can overlap During overlapping periods the greatest min_required_replicas of all scaling schedules is applied Up to 128 scaling schedules are allowed|
|creation_timestamp|text|Creation timestamp in RFC3339 text format|
|description|text|An optional description of this resource Provide this property when you create the resource|
|id|text|The unique identifier for the resource This identifier is defined by the server|
|kind|text|Type of the resource Always compute#autoscaler for autoscalers|
|name|text|Name of the resource Provided by the client when the resource is created The name must be 1-63 characters long, and comply with RFC1035 Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash|
|recommended_size|bigint|Target recommended MIG size (number of instances) computed by autoscaler Autoscaler calculates the recommended MIG size even when the autoscaling policy mode is different from ON This field is empty when autoscaler is not connected to an existing managed instance group or autoscaler did not generate its prediction|
|region|text|URL of the region where the instance group resides (for autoscalers living in regional scope)|
|scaling_schedule_status|jsonb|Status information of existing scaling schedules|
|self_link|text|Server-defined URL for the resource|
|status|text|The status of the autoscaler configuration Current set of possible values: - PENDING: Autoscaler backend hasn't read new/updated configuration  - DELETING: Configuration is being deleted - ACTIVE: Configuration is acknowledged to be effective Some warnings might be present in the statusDetails field - ERROR: Configuration has errors Actionable for users Details are present in the statusDetails field  New values might be added in the future|
|status_details|jsonb|Human-readable details about the current state of the autoscaler Read the documentation for Commonly returned status messages for examples of status messages you might encounter|
|target|text|URL of the managed instance group that this autoscaler will scale This field is required when creating an autoscaler|
|zone|text|URL of the zone where the instance group resides (for autoscalers living in zonal scope)|
