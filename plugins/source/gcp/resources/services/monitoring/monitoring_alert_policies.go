package monitoring

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/monitoring/v3"
)

func MonitoringAlertPolicies() *schema.Table {
	return &schema.Table{
		Name:        "gcp_monitoring_alert_policies",
		Description: "A description of the conditions under which some aspect of your system is considered to be \"unhealthy\" and the ways to notify people or services about this state For an overview of alert policies, see Introduction to Alerting (https://cloudgooglecom/monitoring/alerts/)",
		Resolver:    fetchMonitoringAlertPolicies,
		Multiplex:   client.ProjectMultiplex,

		Options: schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "name"}},

		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "combiner",
				Description: "How to combine the results of multiple conditions to determine if an incident should be opened If condition_time_series_query_language is present, this must be COMBINE_UNSPECIFIED  Possible values:   \"COMBINE_UNSPECIFIED\" - An unspecified combiner   \"AND\" - Combine conditions using the logical AND operator An incident is created only if all the conditions are met simultaneously This combiner is satisfied if all conditions are met, even if they are met on completely different resources   \"OR\" - Combine conditions using the logical OR operator An incident is created if any of the listed conditions is met   \"AND_WITH_MATCHING_RESOURCE\" - Combine conditions using logical AND operator, but unlike the regular AND option, an incident is created only if all conditions are met simultaneously on at least one resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "creation_record_mutate_time",
				Description: "When the change occurred",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CreationRecord.MutateTime"),
			},
			{
				Name:        "creation_record_mutated_by",
				Description: "The email address of the user making the change",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CreationRecord.MutatedBy"),
			},
			{
				Name:        "display_name",
				Description: "A short name or phrase used to identify the policy in dashboards, notifications, and incidents To avoid confusion, don't use the same display name for multiple policies in the same project The name is limited to 512 Unicode characters",
				Type:        schema.TypeString,
			},
			{
				Name:        "documentation_content",
				Description: "The text of the documentation, interpreted according to mime_type The content may not exceed 8,192 Unicode characters and may not exceed more than 10,240 bytes when encoded in UTF-8 format, whichever is smaller",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Documentation.Content"),
			},
			{
				Name:        "documentation_mime_type",
				Description: "The format of the content field Presently, only the value \"text/markdown\" is supported See Markdown (https://enwikipediaorg/wiki/Markdown) for more information",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Documentation.MimeType"),
			},
			{
				Name:        "enabled",
				Description: "Whether or not the policy is enabled On write, the default interpretation if unset is that the policy is enabled On read, clients should not make any assumption about the state if it has not been populated The field should always be populated on List and Get operations, unless a field projection has been specified that strips it out",
				Type:        schema.TypeBool,
			},
			{
				Name:        "mutate_time",
				Description: "When the change occurred",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MutationRecord.MutateTime"),
			},
			{
				Name:        "mutated_by",
				Description: "The email address of the user making the change",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MutationRecord.MutatedBy"),
			},
			{
				Name:        "name",
				Description: "The resource name for this policy",
				Type:        schema.TypeString,
			},
			{
				Name:        "notification_channels",
				Description: "Identifies the notification channels to which notifications should be sent when incidents are opened or closed or when new violations occur on an already opened incident Each element of this array corresponds to the name field in each of the NotificationChannel objects that are returned from the ListNotificationChannels method",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "labels",
				Description: "Labels for this resource",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("UserLabels"),
			},
			{
				Name:        "validity_code",
				Description: "The status code, which should be an enum value of googlerpcCode",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Validity.Code"),
			},
			{
				Name:        "validity_message",
				Description: "A developer-facing error message, which should be in English Any user-facing error message should be localized and sent in the googlerpcStatusdetails field, or localized by the client",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Validity.Message"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "gcp_monitoring_alert_policy_conditions",
				Description: "A condition is a true/false test that determines when an alerting policy should open an incident If a condition evaluates to true, it signifies that something is wrong",
				Resolver:    fetchMonitoringAlertPolicyConditions,
				Columns: []schema.Column{
					{
						Name:        "alert_policy_cq_id",
						Description: "Unique ID of gcp_monitoring_alert_policies table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "absent_duration",
						Description: "The amount of time that a time series must fail to report new data to be considered failing The minimum value of this field is 120 seconds Larger values that are a multiple of a minute--for example, 240 or 300 seconds--are supported If an invalid value is given, an error will be returned The Durationnanos field is ignored",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ConditionAbsent.Duration"),
					},
					{
						Name:        "absent_filter",
						Description: "A filter (https://cloudgooglecom/monitoring/api/v3/filters) that identifies which time series should be compared with the thresholdThe filter is similar to the one that is specified in the ListTimeSeries request (https://cloudgooglecom/monitoring/api/ref_v3/rest/v3/projectstimeSeries/list) (that call is useful to verify the time series that will be retrieved / processed) The filter must specify the metric type and the resource type Optionally, it can specify resource labels and metric labels This field must not exceed 2048 Unicode characters in length",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ConditionAbsent.Filter"),
					},
					{
						Name:        "absent_trigger_count",
						Description: "The absolute number of time series that must fail the predicate for the condition to be triggered",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ConditionAbsent.Trigger.Count"),
					},
					{
						Name:        "absent_trigger_percent",
						Description: "The percentage of time series that must fail the predicate for the condition to be triggered",
						Type:        schema.TypeFloat,
						Resolver:    schema.PathResolver("ConditionAbsent.Trigger.Percent"),
					},
					{
						Name:        "monitoring_query_language_duration",
						Description: "The amount of time that a time series must violate the threshold to be considered failing Currently, only values that are a multiple of a minute--eg, 0, 60, 120, or 300 seconds--are supported If an invalid value is given, an error will be returned When choosing a duration, it is useful to keep in mind the frequency of the underlying time series data (which may also be affected by any alignments specified in the aggregations field); a good duration is long enough so that a single outlier does not generate spurious alerts, but short enough that unhealthy states are detected and alerted on quickly",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ConditionMonitoringQueryLanguage.Duration"),
					},
					{
						Name:        "monitoring_query_language_query",
						Description: "Monitoring Query Language (https://cloudgooglecom/monitoring/mql) query that outputs a boolean stream",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ConditionMonitoringQueryLanguage.Query"),
					},
					{
						Name:        "monitoring_query_language_trigger_count",
						Description: "The absolute number of time series that must fail the predicate for the condition to be triggered",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ConditionMonitoringQueryLanguage.Trigger.Count"),
					},
					{
						Name:        "monitoring_query_language_trigger_percent",
						Description: "The percentage of time series that must fail the predicate for the condition to be triggered",
						Type:        schema.TypeFloat,
						Resolver:    schema.PathResolver("ConditionMonitoringQueryLanguage.Trigger.Percent"),
					},
					{
						Name:        "threshold_comparison",
						Description: "The comparison to apply between the time series (indicated by filter and aggregation) and the threshold (indicated by threshold_value) The comparison is applied on each time series, with the time series on the left-hand side and the threshold on the right-hand sideOnly COMPARISON_LT and COMPARISON_GT are supported currently  Possible values:   \"COMPARISON_UNSPECIFIED\" - No ordering relationship is specified   \"COMPARISON_GT\" - True if the left argument is greater than the right argument   \"COMPARISON_GE\" - True if the left argument is greater than or equal to the right argument   \"COMPARISON_LT\" - True if the left argument is less than the right argument   \"COMPARISON_LE\" - True if the left argument is less than or equal to the right argument   \"COMPARISON_EQ\" - True if the left argument is equal to the right argument   \"COMPARISON_NE\" - True if the left argument is not equal to the right argument",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ConditionThreshold.Comparison"),
					},
					{
						Name:        "threshold_denominator_filter",
						Description: "A filter (https://cloudgooglecom/monitoring/api/v3/filters) that identifies a time series that should be used as the denominator of a ratio that will be compared with the threshold If a denominator_filter is specified, the time series specified by the filter field will be used as the numeratorThe filter must specify the metric type and optionally may contain restrictions on resource type, resource labels, and metric labels This field may not exceed 2048 Unicode characters in length",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ConditionThreshold.DenominatorFilter"),
					},
					{
						Name:        "threshold_duration",
						Description: "The amount of time that a time series must violate the threshold to be considered failing Currently, only values that are a multiple of a minute--eg, 0, 60, 120, or 300 seconds--are supported If an invalid value is given, an error will be returned When choosing a duration, it is useful to keep in mind the frequency of the underlying time series data (which may also be affected by any alignments specified in the aggregations field); a good duration is long enough so that a single outlier does not generate spurious alerts, but short enough that unhealthy states are detected and alerted on quickly",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ConditionThreshold.Duration"),
					},
					{
						Name:        "threshold_filter",
						Description: "A filter (https://cloudgooglecom/monitoring/api/v3/filters) that identifies which time series should be compared with the thresholdThe filter is similar to the one that is specified in the ListTimeSeries request (https://cloudgooglecom/monitoring/api/ref_v3/rest/v3/projectstimeSeries/list) (that call is useful to verify the time series that will be retrieved / processed) The filter must specify the metric type and the resource type Optionally, it can specify resource labels and metric labels This field must not exceed 2048 Unicode characters in length",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ConditionThreshold.Filter"),
					},
					{
						Name:        "threshold_value",
						Description: "A value against which to compare the time series",
						Type:        schema.TypeFloat,
						Resolver:    schema.PathResolver("ConditionThreshold.ThresholdValue"),
					},
					{
						Name:        "threshold_trigger_count",
						Description: "The absolute number of time series that must fail the predicate for the condition to be triggered",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ConditionThreshold.Trigger.Count"),
					},
					{
						Name:        "threshold_trigger_percent",
						Description: "The percentage of time series that must fail the predicate for the condition to be triggered",
						Type:        schema.TypeFloat,
						Resolver:    schema.PathResolver("ConditionThreshold.Trigger.Percent"),
					},
					{
						Name:        "display_name",
						Description: "A short name or phrase used to identify the condition in dashboards, notifications, and incidents To avoid confusion, don't use the same display name for multiple conditions in the same policy",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "if the condition exists The unique resource name for this condition Its format is: projects/[PROJECT_ID_OR_NUMBER]/alertPolicies/[POLICY_ID]/conditions/[ CONDITION_ID] [CONDITION_ID] is assigned by Stackdriver Monitoring when the condition is created as part of a new or updated alerting policyWhen calling the alertPoliciescreate method, do not include the name field in the conditions of the requested alerting policy Stackdriver Monitoring creates the condition identifiers and includes them in the new policyWhen calling the alertPoliciesupdate method to update a policy, including a condition name causes the existing condition to be updated Conditions without names are added to the updated policy Existing conditions are deleted if they are not updatedBest practice is to preserve [CONDITION_ID] if you make only small changes, such as those to condition thresholds, durations, or trigger values Otherwise, treat the change as a new condition and let the existing condition be deleted",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "gcp_monitoring_alert_policy_condition_absent_aggregations",
						Description: "Describes how to combine multiple time series to provide a different view of the data Aggregation of time series is done in two steps First, each time series in the set is aligned to the same time interval boundaries, then the set of time series is optionally reduced in numberAlignment consists of applying the per_series_aligner operation to each time series after its data has been divided into regular alignment_period time intervals This process takes all of the data points in an alignment period, applies a mathematical transformation such as averaging, minimum, maximum, delta, etc, and converts them into a single data point per periodReduction is when the aligned and transformed time series can optionally be combined, reducing the number of time series through similar mathematical transformations Reduction involves applying a cross_series_reducer to all the time series, optionally sorting the time series into subsets with group_by_fields, and applying the reducer to each subsetThe raw time series data can contain a huge amount of information from multiple sources Alignment and reduction transforms this mass of data into a more manageable and representative collection of data, for example \"the 95% latency across the average of all tasks in a cluster\" This representative data can be more easily graphed and comprehended, and the individual time series data is still available for later drilldown For more details, see Filtering and aggregation (https://cloudgooglecom/monitoring/api/v3/aggregation)",
						Resolver:    fetchMonitoringAlertPolicyConditionAbsentAggregations,
						Columns: []schema.Column{
							{
								Name:        "alert_policy_condition_cq_id",
								Description: "Unique ID of gcp_monitoring_alert_policy_conditions table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "alignment_period",
								Description: "The alignment_period specifies a time interval, in seconds, that is used to divide the data in all the time series into consistent blocks of time This will be done before the per-series aligner can be applied to the dataThe value must be at least 60 seconds If a per-series aligner other than ALIGN_NONE is specified, this field is required or an error is returned If no per-series aligner is specified, or the aligner ALIGN_NONE is specified, then this field is ignoredThe maximum value of the alignment_period is 104 weeks (2 years) for charts, and 90,000 seconds (25 hours) for alerting policies",
								Type:        schema.TypeString,
							},
							{
								Name:        "cross_series_reducer",
								Description: "The reduction operation to be used to combine time series into a single time series, where the value of each data point in the resulting series is a function of all the already aligned values in the input time seriesNot all reducer operations can be applied to all time series The valid choices depend on the metric_kind and the value_type of the original time series Reduction can yield a time series with a different metric_kind or value_type than the input time seriesTime series data must first be aligned (see per_series_aligner) in order to perform cross-time series reduction If cross_series_reducer is specified, then per_series_aligner must be specified, and must not be ALIGN_NONE An alignment_period must also be specified; otherwise, an error is returned  Possible values:   \"REDUCE_NONE\" - No cross-time series reduction The output of the Aligner is returned   \"REDUCE_MEAN\" - Reduce by computing the mean value across time series for each alignment period This reducer is valid for DELTA and GAUGE metrics with numeric or distribution values The value_type of the output is DOUBLE   \"REDUCE_MIN\" - Reduce by computing the minimum value across time series for each alignment period This reducer is valid for DELTA and GAUGE metrics with numeric values The value_type of the output is the same as the value_type of the input   \"REDUCE_MAX\" - Reduce by computing the maximum value across time series for each alignment period This reducer is valid for DELTA and GAUGE metrics with numeric values The value_type of the output is the same as the value_type of the input   \"REDUCE_SUM\" - Reduce by computing the sum across time series for each alignment period This reducer is valid for DELTA and GAUGE metrics with numeric and distribution values The value_type of the output is the same as the value_type of the input   \"REDUCE_STDDEV\" - Reduce by computing the standard deviation across time series for each alignment period This reducer is valid for DELTA and GAUGE metrics with numeric or distribution values The value_type of the output is DOUBLE   \"REDUCE_COUNT\" - Reduce by computing the number of data points across time series for each alignment period This reducer is valid for DELTA and GAUGE metrics of numeric, Boolean, distribution, and string value_type The value_type of the output is INT64   \"REDUCE_COUNT_TRUE\" - Reduce by computing the number of True-valued data points across time series for each alignment period This reducer is valid for DELTA and GAUGE metrics of Boolean value_type The value_type of the output is INT64   \"REDUCE_COUNT_FALSE\" - Reduce by computing the number of False-valued data points across time series for each alignment period This reducer is valid for DELTA and GAUGE metrics of Boolean value_type The value_type of the output is INT64   \"REDUCE_FRACTION_TRUE\" - Reduce by computing the ratio of the number of True-valued data points to the total number of data points for each alignment period This reducer is valid for DELTA and GAUGE metrics of Boolean value_type The output value is in the range 00, 10 and has value_type DOUBLE   \"REDUCE_PERCENTILE_99\" - Reduce by computing the 99th percentile (https://enwikipediaorg/wiki/Percentile) of data points across time series for each alignment period This reducer is valid for GAUGE and DELTA metrics of numeric and distribution type The value of the output is DOUBLE   \"REDUCE_PERCENTILE_95\" - Reduce by computing the 95th percentile (https://enwikipediaorg/wiki/Percentile) of data points across time series for each alignment period This reducer is valid for GAUGE and DELTA metrics of numeric and distribution type The value of the output is DOUBLE   \"REDUCE_PERCENTILE_50\" - Reduce by computing the 50th percentile (https://enwikipediaorg/wiki/Percentile) of data points across time series for each alignment period This reducer is valid for GAUGE and DELTA metrics of numeric and distribution type The value of the output is DOUBLE   \"REDUCE_PERCENTILE_05\" - Reduce by computing the 5th percentile (https://enwikipediaorg/wiki/Percentile) of data points across time series for each alignment period This reducer is valid for GAUGE and DELTA metrics of numeric and distribution type The value of the output is DOUBLE",
								Type:        schema.TypeString,
							},
							{
								Name:        "group_by_fields",
								Description: "The set of fields to preserve when cross_series_reducer is specified The group_by_fields determine how the time series are partitioned into subsets prior to applying the aggregation operation Each subset contains time series that have the same value for each of the grouping fields Each individual time series is a member of exactly one subset The cross_series_reducer is applied to each subset of time series It is not possible to reduce across different resource types, so this field implicitly contains resourcetype Fields not specified in group_by_fields are aggregated away If group_by_fields is not specified and all the time series have the same resource type, then the time series are aggregated into a single output time series If cross_series_reducer is not defined, this field is ignored",
								Type:        schema.TypeStringArray,
							},
							{
								Name:        "per_series_aligner",
								Description: "An Aligner describes how to bring the data points in a single time series into temporal alignment Except for ALIGN_NONE, all alignments cause all the data points in an alignment_period to be mathematically grouped together, resulting in a single data point for each alignment_period with end timestamp at the end of the periodNot all alignment operations may be applied to all time series The valid choices depend on the metric_kind and value_type of the original time series Alignment can change the metric_kind or the value_type of the time seriesTime series data must be aligned in order to perform cross-time series reduction If cross_series_reducer is specified, then per_series_aligner must be specified and not equal to ALIGN_NONE and alignment_period must be specified; otherwise, an error is returned  Possible values:   \"ALIGN_NONE\" - No alignment Raw data is returned Not valid if cross-series reduction is requested The value_type of the result is the same as the value_type of the input   \"ALIGN_DELTA\" - Align and convert to DELTA The output is delta = y1 - y0This alignment is valid for CUMULATIVE and DELTA metrics If the selected alignment period results in periods with no data, then the aligned value for such a period is created by interpolation The value_type of the aligned result is the same as the value_type of the input   \"ALIGN_RATE\" - Align and convert to a rate The result is computed as rate = (y1 - y0)/(t1 - t0), or \"delta over time\" Think of this aligner as providing the slope of the line that passes through the value at the start and at the end of the alignment_periodThis aligner is valid for CUMULATIVE and DELTA metrics with numeric values If the selected alignment period results in periods with no data, then the aligned value for such a period is created by interpolation The output is a GAUGE metric with value_type DOUBLEIf, by \"rate\", you mean \"percentage change\", see the ALIGN_PERCENT_CHANGE aligner instead   \"ALIGN_INTERPOLATE\" - Align by interpolating between adjacent points around the alignment period boundary This aligner is valid for GAUGE metrics with numeric values The value_type of the aligned result is the same as the value_type of the input   \"ALIGN_NEXT_OLDER\" - Align by moving the most recent data point before the end of the alignment period to the boundary at the end of the alignment period This aligner is valid for GAUGE metrics The value_type of the aligned result is the same as the value_type of the input   \"ALIGN_MIN\" - Align the time series by returning the minimum value in each alignment period This aligner is valid for GAUGE and DELTA metrics with numeric values The value_type of the aligned result is the same as the value_type of the input   \"ALIGN_MAX\" - Align the time series by returning the maximum value in each alignment period This aligner is valid for GAUGE and DELTA metrics with numeric values The value_type of the aligned result is the same as the value_type of the input   \"ALIGN_MEAN\" - Align the time series by returning the mean value in each alignment period This aligner is valid for GAUGE and DELTA metrics with numeric values The value_type of the aligned result is DOUBLE   \"ALIGN_COUNT\" - Align the time series by returning the number of values in each alignment period This aligner is valid for GAUGE and DELTA metrics with numeric or Boolean values The value_type of the aligned result is INT64   \"ALIGN_SUM\" - Align the time series by returning the sum of the values in each alignment period This aligner is valid for GAUGE and DELTA metrics with numeric and distribution values The value_type of the aligned result is the same as the value_type of the input   \"ALIGN_STDDEV\" - Align the time series by returning the standard deviation of the values in each alignment period This aligner is valid for GAUGE and DELTA metrics with numeric values The value_type of the output is DOUBLE   \"ALIGN_COUNT_TRUE\" - Align the time series by returning the number of True values in each alignment period This aligner is valid for GAUGE metrics with Boolean values The value_type of the output is INT64   \"ALIGN_COUNT_FALSE\" - Align the time series by returning the number of False values in each alignment period This aligner is valid for GAUGE metrics with Boolean values The value_type of the output is INT64   \"ALIGN_FRACTION_TRUE\" - Align the time series by returning the ratio of the number of True values to the total number of values in each alignment period This aligner is valid for GAUGE metrics with Boolean values The output value is in the range 00, 10 and has value_type DOUBLE   \"ALIGN_PERCENTILE_99\" - Align the time series by using percentile aggregation (https://enwikipediaorg/wiki/Percentile) The resulting data point in each alignment period is the 99th percentile of all data points in the period This aligner is valid for GAUGE and DELTA metrics with distribution values The output is a GAUGE metric with value_type DOUBLE   \"ALIGN_PERCENTILE_95\" - Align the time series by using percentile aggregation (https://enwikipediaorg/wiki/Percentile) The resulting data point in each alignment period is the 95th percentile of all data points in the period This aligner is valid for GAUGE and DELTA metrics with distribution values The output is a GAUGE metric with value_type DOUBLE   \"ALIGN_PERCENTILE_50\" - Align the time series by using percentile aggregation (https://enwikipediaorg/wiki/Percentile) The resulting data point in each alignment period is the 50th percentile of all data points in the period This aligner is valid for GAUGE and DELTA metrics with distribution values The output is a GAUGE metric with value_type DOUBLE   \"ALIGN_PERCENTILE_05\" - Align the time series by using percentile aggregation (https://enwikipediaorg/wiki/Percentile) The resulting data point in each alignment period is the 5th percentile of all data points in the period This aligner is valid for GAUGE and DELTA metrics with distribution values The output is a GAUGE metric with value_type DOUBLE   \"ALIGN_PERCENT_CHANGE\" - Align and convert to a percentage change This aligner is valid for GAUGE and DELTA metrics with numeric values This alignment returns ((current - previous)/previous) * 100, where the value of previous is determined based on the alignment_periodIf the values of current and previous are both 0, then the returned value is 0 If only previous is 0, the returned value is infinityA 10-minute moving mean is computed at each point of the alignment period prior to the above calculation to smooth the metric and prevent false positives from very short-lived spikes The moving mean is only applicable for data whose values are >= 0 Any values < 0 are treated as a missing datapoint, and are ignored While DELTA metrics are accepted by this alignment, special care should be taken that the values for the metric will always be positive The output is a GAUGE metric with value_type DOUBLE",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "gcp_monitoring_alert_policy_condition_threshold_aggregations",
						Description: "Describes how to combine multiple time series to provide a different view of the data Aggregation of time series is done in two steps First, each time series in the set is aligned to the same time interval boundaries, then the set of time series is optionally reduced in numberAlignment consists of applying the per_series_aligner operation to each time series after its data has been divided into regular alignment_period time intervals This process takes all of the data points in an alignment period, applies a mathematical transformation such as averaging, minimum, maximum, delta, etc, and converts them into a single data point per periodReduction is when the aligned and transformed time series can optionally be combined, reducing the number of time series through similar mathematical transformations Reduction involves applying a cross_series_reducer to all the time series, optionally sorting the time series into subsets with group_by_fields, and applying the reducer to each subsetThe raw time series data can contain a huge amount of information from multiple sources Alignment and reduction transforms this mass of data into a more manageable and representative collection of data, for example \"the 95% latency across the average of all tasks in a cluster\" This representative data can be more easily graphed and comprehended, and the individual time series data is still available for later drilldown For more details, see Filtering and aggregation (https://cloudgooglecom/monitoring/api/v3/aggregation)",
						Resolver:    fetchMonitoringAlertPolicyConditionThresholdAggregations,
						Columns: []schema.Column{
							{
								Name:        "alert_policy_condition_cq_id",
								Description: "Unique ID of gcp_monitoring_alert_policy_conditions table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "alignment_period",
								Description: "The alignment_period specifies a time interval, in seconds, that is used to divide the data in all the time series into consistent blocks of time This will be done before the per-series aligner can be applied to the dataThe value must be at least 60 seconds If a per-series aligner other than ALIGN_NONE is specified, this field is required or an error is returned If no per-series aligner is specified, or the aligner ALIGN_NONE is specified, then this field is ignoredThe maximum value of the alignment_period is 104 weeks (2 years) for charts, and 90,000 seconds (25 hours) for alerting policies",
								Type:        schema.TypeString,
							},
							{
								Name:        "cross_series_reducer",
								Description: "The reduction operation to be used to combine time series into a single time series, where the value of each data point in the resulting series is a function of all the already aligned values in the input time seriesNot all reducer operations can be applied to all time series The valid choices depend on the metric_kind and the value_type of the original time series Reduction can yield a time series with a different metric_kind or value_type than the input time seriesTime series data must first be aligned (see per_series_aligner) in order to perform cross-time series reduction If cross_series_reducer is specified, then per_series_aligner must be specified, and must not be ALIGN_NONE An alignment_period must also be specified; otherwise, an error is returned  Possible values:   \"REDUCE_NONE\" - No cross-time series reduction The output of the Aligner is returned   \"REDUCE_MEAN\" - Reduce by computing the mean value across time series for each alignment period This reducer is valid for DELTA and GAUGE metrics with numeric or distribution values The value_type of the output is DOUBLE   \"REDUCE_MIN\" - Reduce by computing the minimum value across time series for each alignment period This reducer is valid for DELTA and GAUGE metrics with numeric values The value_type of the output is the same as the value_type of the input   \"REDUCE_MAX\" - Reduce by computing the maximum value across time series for each alignment period This reducer is valid for DELTA and GAUGE metrics with numeric values The value_type of the output is the same as the value_type of the input   \"REDUCE_SUM\" - Reduce by computing the sum across time series for each alignment period This reducer is valid for DELTA and GAUGE metrics with numeric and distribution values The value_type of the output is the same as the value_type of the input   \"REDUCE_STDDEV\" - Reduce by computing the standard deviation across time series for each alignment period This reducer is valid for DELTA and GAUGE metrics with numeric or distribution values The value_type of the output is DOUBLE   \"REDUCE_COUNT\" - Reduce by computing the number of data points across time series for each alignment period This reducer is valid for DELTA and GAUGE metrics of numeric, Boolean, distribution, and string value_type The value_type of the output is INT64   \"REDUCE_COUNT_TRUE\" - Reduce by computing the number of True-valued data points across time series for each alignment period This reducer is valid for DELTA and GAUGE metrics of Boolean value_type The value_type of the output is INT64   \"REDUCE_COUNT_FALSE\" - Reduce by computing the number of False-valued data points across time series for each alignment period This reducer is valid for DELTA and GAUGE metrics of Boolean value_type The value_type of the output is INT64   \"REDUCE_FRACTION_TRUE\" - Reduce by computing the ratio of the number of True-valued data points to the total number of data points for each alignment period This reducer is valid for DELTA and GAUGE metrics of Boolean value_type The output value is in the range 00, 10 and has value_type DOUBLE   \"REDUCE_PERCENTILE_99\" - Reduce by computing the 99th percentile (https://enwikipediaorg/wiki/Percentile) of data points across time series for each alignment period This reducer is valid for GAUGE and DELTA metrics of numeric and distribution type The value of the output is DOUBLE   \"REDUCE_PERCENTILE_95\" - Reduce by computing the 95th percentile (https://enwikipediaorg/wiki/Percentile) of data points across time series for each alignment period This reducer is valid for GAUGE and DELTA metrics of numeric and distribution type The value of the output is DOUBLE   \"REDUCE_PERCENTILE_50\" - Reduce by computing the 50th percentile (https://enwikipediaorg/wiki/Percentile) of data points across time series for each alignment period This reducer is valid for GAUGE and DELTA metrics of numeric and distribution type The value of the output is DOUBLE   \"REDUCE_PERCENTILE_05\" - Reduce by computing the 5th percentile (https://enwikipediaorg/wiki/Percentile) of data points across time series for each alignment period This reducer is valid for GAUGE and DELTA metrics of numeric and distribution type The value of the output is DOUBLE",
								Type:        schema.TypeString,
							},
							{
								Name:        "group_by_fields",
								Description: "The set of fields to preserve when cross_series_reducer is specified The group_by_fields determine how the time series are partitioned into subsets prior to applying the aggregation operation Each subset contains time series that have the same value for each of the grouping fields Each individual time series is a member of exactly one subset The cross_series_reducer is applied to each subset of time series It is not possible to reduce across different resource types, so this field implicitly contains resourcetype Fields not specified in group_by_fields are aggregated away If group_by_fields is not specified and all the time series have the same resource type, then the time series are aggregated into a single output time series If cross_series_reducer is not defined, this field is ignored",
								Type:        schema.TypeStringArray,
							},
							{
								Name:        "per_series_aligner",
								Description: "An Aligner describes how to bring the data points in a single time series into temporal alignment Except for ALIGN_NONE, all alignments cause all the data points in an alignment_period to be mathematically grouped together, resulting in a single data point for each alignment_period with end timestamp at the end of the periodNot all alignment operations may be applied to all time series The valid choices depend on the metric_kind and value_type of the original time series Alignment can change the metric_kind or the value_type of the time seriesTime series data must be aligned in order to perform cross-time series reduction If cross_series_reducer is specified, then per_series_aligner must be specified and not equal to ALIGN_NONE and alignment_period must be specified; otherwise, an error is returned  Possible values:   \"ALIGN_NONE\" - No alignment Raw data is returned Not valid if cross-series reduction is requested The value_type of the result is the same as the value_type of the input   \"ALIGN_DELTA\" - Align and convert to DELTA The output is delta = y1 - y0This alignment is valid for CUMULATIVE and DELTA metrics If the selected alignment period results in periods with no data, then the aligned value for such a period is created by interpolation The value_type of the aligned result is the same as the value_type of the input   \"ALIGN_RATE\" - Align and convert to a rate The result is computed as rate = (y1 - y0)/(t1 - t0), or \"delta over time\" Think of this aligner as providing the slope of the line that passes through the value at the start and at the end of the alignment_periodThis aligner is valid for CUMULATIVE and DELTA metrics with numeric values If the selected alignment period results in periods with no data, then the aligned value for such a period is created by interpolation The output is a GAUGE metric with value_type DOUBLEIf, by \"rate\", you mean \"percentage change\", see the ALIGN_PERCENT_CHANGE aligner instead   \"ALIGN_INTERPOLATE\" - Align by interpolating between adjacent points around the alignment period boundary This aligner is valid for GAUGE metrics with numeric values The value_type of the aligned result is the same as the value_type of the input   \"ALIGN_NEXT_OLDER\" - Align by moving the most recent data point before the end of the alignment period to the boundary at the end of the alignment period This aligner is valid for GAUGE metrics The value_type of the aligned result is the same as the value_type of the input   \"ALIGN_MIN\" - Align the time series by returning the minimum value in each alignment period This aligner is valid for GAUGE and DELTA metrics with numeric values The value_type of the aligned result is the same as the value_type of the input   \"ALIGN_MAX\" - Align the time series by returning the maximum value in each alignment period This aligner is valid for GAUGE and DELTA metrics with numeric values The value_type of the aligned result is the same as the value_type of the input   \"ALIGN_MEAN\" - Align the time series by returning the mean value in each alignment period This aligner is valid for GAUGE and DELTA metrics with numeric values The value_type of the aligned result is DOUBLE   \"ALIGN_COUNT\" - Align the time series by returning the number of values in each alignment period This aligner is valid for GAUGE and DELTA metrics with numeric or Boolean values The value_type of the aligned result is INT64   \"ALIGN_SUM\" - Align the time series by returning the sum of the values in each alignment period This aligner is valid for GAUGE and DELTA metrics with numeric and distribution values The value_type of the aligned result is the same as the value_type of the input   \"ALIGN_STDDEV\" - Align the time series by returning the standard deviation of the values in each alignment period This aligner is valid for GAUGE and DELTA metrics with numeric values The value_type of the output is DOUBLE   \"ALIGN_COUNT_TRUE\" - Align the time series by returning the number of True values in each alignment period This aligner is valid for GAUGE metrics with Boolean values The value_type of the output is INT64   \"ALIGN_COUNT_FALSE\" - Align the time series by returning the number of False values in each alignment period This aligner is valid for GAUGE metrics with Boolean values The value_type of the output is INT64   \"ALIGN_FRACTION_TRUE\" - Align the time series by returning the ratio of the number of True values to the total number of values in each alignment period This aligner is valid for GAUGE metrics with Boolean values The output value is in the range 00, 10 and has value_type DOUBLE   \"ALIGN_PERCENTILE_99\" - Align the time series by using percentile aggregation (https://enwikipediaorg/wiki/Percentile) The resulting data point in each alignment period is the 99th percentile of all data points in the period This aligner is valid for GAUGE and DELTA metrics with distribution values The output is a GAUGE metric with value_type DOUBLE   \"ALIGN_PERCENTILE_95\" - Align the time series by using percentile aggregation (https://enwikipediaorg/wiki/Percentile) The resulting data point in each alignment period is the 95th percentile of all data points in the period This aligner is valid for GAUGE and DELTA metrics with distribution values The output is a GAUGE metric with value_type DOUBLE   \"ALIGN_PERCENTILE_50\" - Align the time series by using percentile aggregation (https://enwikipediaorg/wiki/Percentile) The resulting data point in each alignment period is the 50th percentile of all data points in the period This aligner is valid for GAUGE and DELTA metrics with distribution values The output is a GAUGE metric with value_type DOUBLE   \"ALIGN_PERCENTILE_05\" - Align the time series by using percentile aggregation (https://enwikipediaorg/wiki/Percentile) The resulting data point in each alignment period is the 5th percentile of all data points in the period This aligner is valid for GAUGE and DELTA metrics with distribution values The output is a GAUGE metric with value_type DOUBLE   \"ALIGN_PERCENT_CHANGE\" - Align and convert to a percentage change This aligner is valid for GAUGE and DELTA metrics with numeric values This alignment returns ((current - previous)/previous) * 100, where the value of previous is determined based on the alignment_periodIf the values of current and previous are both 0, then the returned value is 0 If only previous is 0, the returned value is infinityA 10-minute moving mean is computed at each point of the alignment period prior to the above calculation to smooth the metric and prevent false positives from very short-lived spikes The moving mean is only applicable for data whose values are >= 0 Any values < 0 are treated as a missing datapoint, and are ignored While DELTA metrics are accepted by this alignment, special care should be taken that the values for the metric will always be positive The output is a GAUGE metric with value_type DOUBLE",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "gcp_monitoring_alert_policy_condition_denominator_aggs",
						Description: "Describes how to combine multiple time series to provide a different view of the data Aggregation of time series is done in two steps First, each time series in the set is aligned to the same time interval boundaries, then the set of time series is optionally reduced in numberAlignment consists of applying the per_series_aligner operation to each time series after its data has been divided into regular alignment_period time intervals This process takes all of the data points in an alignment period, applies a mathematical transformation such as averaging, minimum, maximum, delta, etc, and converts them into a single data point per periodReduction is when the aligned and transformed time series can optionally be combined, reducing the number of time series through similar mathematical transformations Reduction involves applying a cross_series_reducer to all the time series, optionally sorting the time series into subsets with group_by_fields, and applying the reducer to each subsetThe raw time series data can contain a huge amount of information from multiple sources Alignment and reduction transforms this mass of data into a more manageable and representative collection of data, for example \"the 95% latency across the average of all tasks in a cluster\" This representative data can be more easily graphed and comprehended, and the individual time series data is still available for later drilldown For more details, see Filtering and aggregation (https://cloudgooglecom/monitoring/api/v3/aggregation)",
						Resolver:    fetchMonitoringAlertPolicyConditionDenominatorAggregations,
						Columns: []schema.Column{
							{
								Name:        "alert_policy_condition_cq_id",
								Description: "Unique ID of gcp_monitoring_alert_policy_conditions table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "alignment_period",
								Description: "The alignment_period specifies a time interval, in seconds, that is used to divide the data in all the time series into consistent blocks of time This will be done before the per-series aligner can be applied to the dataThe value must be at least 60 seconds If a per-series aligner other than ALIGN_NONE is specified, this field is required or an error is returned If no per-series aligner is specified, or the aligner ALIGN_NONE is specified, then this field is ignoredThe maximum value of the alignment_period is 104 weeks (2 years) for charts, and 90,000 seconds (25 hours) for alerting policies",
								Type:        schema.TypeString,
							},
							{
								Name:        "cross_series_reducer",
								Description: "The reduction operation to be used to combine time series into a single time series, where the value of each data point in the resulting series is a function of all the already aligned values in the input time seriesNot all reducer operations can be applied to all time series The valid choices depend on the metric_kind and the value_type of the original time series Reduction can yield a time series with a different metric_kind or value_type than the input time seriesTime series data must first be aligned (see per_series_aligner) in order to perform cross-time series reduction If cross_series_reducer is specified, then per_series_aligner must be specified, and must not be ALIGN_NONE An alignment_period must also be specified; otherwise, an error is returned  Possible values:   \"REDUCE_NONE\" - No cross-time series reduction The output of the Aligner is returned   \"REDUCE_MEAN\" - Reduce by computing the mean value across time series for each alignment period This reducer is valid for DELTA and GAUGE metrics with numeric or distribution values The value_type of the output is DOUBLE   \"REDUCE_MIN\" - Reduce by computing the minimum value across time series for each alignment period This reducer is valid for DELTA and GAUGE metrics with numeric values The value_type of the output is the same as the value_type of the input   \"REDUCE_MAX\" - Reduce by computing the maximum value across time series for each alignment period This reducer is valid for DELTA and GAUGE metrics with numeric values The value_type of the output is the same as the value_type of the input   \"REDUCE_SUM\" - Reduce by computing the sum across time series for each alignment period This reducer is valid for DELTA and GAUGE metrics with numeric and distribution values The value_type of the output is the same as the value_type of the input   \"REDUCE_STDDEV\" - Reduce by computing the standard deviation across time series for each alignment period This reducer is valid for DELTA and GAUGE metrics with numeric or distribution values The value_type of the output is DOUBLE   \"REDUCE_COUNT\" - Reduce by computing the number of data points across time series for each alignment period This reducer is valid for DELTA and GAUGE metrics of numeric, Boolean, distribution, and string value_type The value_type of the output is INT64   \"REDUCE_COUNT_TRUE\" - Reduce by computing the number of True-valued data points across time series for each alignment period This reducer is valid for DELTA and GAUGE metrics of Boolean value_type The value_type of the output is INT64   \"REDUCE_COUNT_FALSE\" - Reduce by computing the number of False-valued data points across time series for each alignment period This reducer is valid for DELTA and GAUGE metrics of Boolean value_type The value_type of the output is INT64   \"REDUCE_FRACTION_TRUE\" - Reduce by computing the ratio of the number of True-valued data points to the total number of data points for each alignment period This reducer is valid for DELTA and GAUGE metrics of Boolean value_type The output value is in the range 00, 10 and has value_type DOUBLE   \"REDUCE_PERCENTILE_99\" - Reduce by computing the 99th percentile (https://enwikipediaorg/wiki/Percentile) of data points across time series for each alignment period This reducer is valid for GAUGE and DELTA metrics of numeric and distribution type The value of the output is DOUBLE   \"REDUCE_PERCENTILE_95\" - Reduce by computing the 95th percentile (https://enwikipediaorg/wiki/Percentile) of data points across time series for each alignment period This reducer is valid for GAUGE and DELTA metrics of numeric and distribution type The value of the output is DOUBLE   \"REDUCE_PERCENTILE_50\" - Reduce by computing the 50th percentile (https://enwikipediaorg/wiki/Percentile) of data points across time series for each alignment period This reducer is valid for GAUGE and DELTA metrics of numeric and distribution type The value of the output is DOUBLE   \"REDUCE_PERCENTILE_05\" - Reduce by computing the 5th percentile (https://enwikipediaorg/wiki/Percentile) of data points across time series for each alignment period This reducer is valid for GAUGE and DELTA metrics of numeric and distribution type The value of the output is DOUBLE",
								Type:        schema.TypeString,
							},
							{
								Name:        "group_by_fields",
								Description: "The set of fields to preserve when cross_series_reducer is specified The group_by_fields determine how the time series are partitioned into subsets prior to applying the aggregation operation Each subset contains time series that have the same value for each of the grouping fields Each individual time series is a member of exactly one subset The cross_series_reducer is applied to each subset of time series It is not possible to reduce across different resource types, so this field implicitly contains resourcetype Fields not specified in group_by_fields are aggregated away If group_by_fields is not specified and all the time series have the same resource type, then the time series are aggregated into a single output time series If cross_series_reducer is not defined, this field is ignored",
								Type:        schema.TypeStringArray,
							},
							{
								Name:        "per_series_aligner",
								Description: "An Aligner describes how to bring the data points in a single time series into temporal alignment Except for ALIGN_NONE, all alignments cause all the data points in an alignment_period to be mathematically grouped together, resulting in a single data point for each alignment_period with end timestamp at the end of the periodNot all alignment operations may be applied to all time series The valid choices depend on the metric_kind and value_type of the original time series Alignment can change the metric_kind or the value_type of the time seriesTime series data must be aligned in order to perform cross-time series reduction If cross_series_reducer is specified, then per_series_aligner must be specified and not equal to ALIGN_NONE and alignment_period must be specified; otherwise, an error is returned  Possible values:   \"ALIGN_NONE\" - No alignment Raw data is returned Not valid if cross-series reduction is requested The value_type of the result is the same as the value_type of the input   \"ALIGN_DELTA\" - Align and convert to DELTA The output is delta = y1 - y0This alignment is valid for CUMULATIVE and DELTA metrics If the selected alignment period results in periods with no data, then the aligned value for such a period is created by interpolation The value_type of the aligned result is the same as the value_type of the input   \"ALIGN_RATE\" - Align and convert to a rate The result is computed as rate = (y1 - y0)/(t1 - t0), or \"delta over time\" Think of this aligner as providing the slope of the line that passes through the value at the start and at the end of the alignment_periodThis aligner is valid for CUMULATIVE and DELTA metrics with numeric values If the selected alignment period results in periods with no data, then the aligned value for such a period is created by interpolation The output is a GAUGE metric with value_type DOUBLEIf, by \"rate\", you mean \"percentage change\", see the ALIGN_PERCENT_CHANGE aligner instead   \"ALIGN_INTERPOLATE\" - Align by interpolating between adjacent points around the alignment period boundary This aligner is valid for GAUGE metrics with numeric values The value_type of the aligned result is the same as the value_type of the input   \"ALIGN_NEXT_OLDER\" - Align by moving the most recent data point before the end of the alignment period to the boundary at the end of the alignment period This aligner is valid for GAUGE metrics The value_type of the aligned result is the same as the value_type of the input   \"ALIGN_MIN\" - Align the time series by returning the minimum value in each alignment period This aligner is valid for GAUGE and DELTA metrics with numeric values The value_type of the aligned result is the same as the value_type of the input   \"ALIGN_MAX\" - Align the time series by returning the maximum value in each alignment period This aligner is valid for GAUGE and DELTA metrics with numeric values The value_type of the aligned result is the same as the value_type of the input   \"ALIGN_MEAN\" - Align the time series by returning the mean value in each alignment period This aligner is valid for GAUGE and DELTA metrics with numeric values The value_type of the aligned result is DOUBLE   \"ALIGN_COUNT\" - Align the time series by returning the number of values in each alignment period This aligner is valid for GAUGE and DELTA metrics with numeric or Boolean values The value_type of the aligned result is INT64   \"ALIGN_SUM\" - Align the time series by returning the sum of the values in each alignment period This aligner is valid for GAUGE and DELTA metrics with numeric and distribution values The value_type of the aligned result is the same as the value_type of the input   \"ALIGN_STDDEV\" - Align the time series by returning the standard deviation of the values in each alignment period This aligner is valid for GAUGE and DELTA metrics with numeric values The value_type of the output is DOUBLE   \"ALIGN_COUNT_TRUE\" - Align the time series by returning the number of True values in each alignment period This aligner is valid for GAUGE metrics with Boolean values The value_type of the output is INT64   \"ALIGN_COUNT_FALSE\" - Align the time series by returning the number of False values in each alignment period This aligner is valid for GAUGE metrics with Boolean values The value_type of the output is INT64   \"ALIGN_FRACTION_TRUE\" - Align the time series by returning the ratio of the number of True values to the total number of values in each alignment period This aligner is valid for GAUGE metrics with Boolean values The output value is in the range 00, 10 and has value_type DOUBLE   \"ALIGN_PERCENTILE_99\" - Align the time series by using percentile aggregation (https://enwikipediaorg/wiki/Percentile) The resulting data point in each alignment period is the 99th percentile of all data points in the period This aligner is valid for GAUGE and DELTA metrics with distribution values The output is a GAUGE metric with value_type DOUBLE   \"ALIGN_PERCENTILE_95\" - Align the time series by using percentile aggregation (https://enwikipediaorg/wiki/Percentile) The resulting data point in each alignment period is the 95th percentile of all data points in the period This aligner is valid for GAUGE and DELTA metrics with distribution values The output is a GAUGE metric with value_type DOUBLE   \"ALIGN_PERCENTILE_50\" - Align the time series by using percentile aggregation (https://enwikipediaorg/wiki/Percentile) The resulting data point in each alignment period is the 50th percentile of all data points in the period This aligner is valid for GAUGE and DELTA metrics with distribution values The output is a GAUGE metric with value_type DOUBLE   \"ALIGN_PERCENTILE_05\" - Align the time series by using percentile aggregation (https://enwikipediaorg/wiki/Percentile) The resulting data point in each alignment period is the 5th percentile of all data points in the period This aligner is valid for GAUGE and DELTA metrics with distribution values The output is a GAUGE metric with value_type DOUBLE   \"ALIGN_PERCENT_CHANGE\" - Align and convert to a percentage change This aligner is valid for GAUGE and DELTA metrics with numeric values This alignment returns ((current - previous)/previous) * 100, where the value of previous is determined based on the alignment_periodIf the values of current and previous are both 0, then the returned value is 0 If only previous is 0, the returned value is infinityA 10-minute moving mean is computed at each point of the alignment period prior to the above calculation to smooth the metric and prevent false positives from very short-lived spikes The moving mean is only applicable for data whose values are >= 0 Any values < 0 are treated as a missing datapoint, and are ignored While DELTA metrics are accepted by this alignment, special care should be taken that the values for the metric will always be positive The output is a GAUGE metric with value_type DOUBLE",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchMonitoringAlertPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Monitoring.Projects.AlertPolicies.
			List(fmt.Sprintf("projects/%s", c.ProjectId)).
			PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		res <- output.AlertPolicies

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func fetchMonitoringAlertPolicyConditions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*monitoring.AlertPolicy)
	res <- p.Conditions
	return nil
}
func fetchMonitoringAlertPolicyConditionAbsentAggregations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*monitoring.Condition)
	if p.ConditionAbsent == nil {
		return nil
	}
	res <- p.ConditionAbsent.Aggregations
	return nil
}
func fetchMonitoringAlertPolicyConditionThresholdAggregations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*monitoring.Condition)
	if p.ConditionThreshold == nil {
		return nil
	}
	res <- p.ConditionThreshold.Aggregations
	return nil
}
func fetchMonitoringAlertPolicyConditionDenominatorAggregations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*monitoring.Condition)
	if p.ConditionThreshold == nil {
		return nil
	}
	res <- p.ConditionThreshold.DenominatorAggregations
	return nil
}
