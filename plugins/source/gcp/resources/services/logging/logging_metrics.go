package logging

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/logging/v2"
)

func LoggingMetrics() *schema.Table {
	return &schema.Table{
		Name:         "gcp_logging_metrics",
		Description:  "Describes a logs-based metric The value of the metric is the number of log entries that match a logs filter in a given time intervalLogs-based metrics can also be used to extract values from logs and create a distribution of the values The distribution records the statistics of the extracted values along with an optional histogram of the values as specified by the bucket options",
		Resolver:     fetchLoggingMetrics,
		Multiplex:    client.ProjectMultiplex,
		IgnoreError:  client.IgnoreErrorHandler,
		DeleteFilter: client.DeleteProjectFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "name"}},
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "exponential_buckets_options_growth_factor",
				Description: "Must be greater than 1",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("BucketOptions.ExponentialBuckets.GrowthFactor"),
			},
			{
				Name:        "exponential_buckets_options_num_finite_buckets",
				Description: "Must be greater than 0",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("BucketOptions.ExponentialBuckets.NumFiniteBuckets"),
			},
			{
				Name:        "exponential_buckets_options_scale",
				Description: "Must be greater than 0",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("BucketOptions.ExponentialBuckets.Scale"),
			},
			{
				Name:        "linear_buckets_options_num_finite_buckets",
				Description: "Must be greater than 0",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("BucketOptions.LinearBuckets.NumFiniteBuckets"),
			},
			{
				Name:        "linear_buckets_options_offset",
				Description: "Lower bound of the first bucket",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("BucketOptions.LinearBuckets.Offset"),
			},
			{
				Name:        "linear_buckets_options_width",
				Description: "Must be greater than 0",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("BucketOptions.LinearBuckets.Width"),
			},
			{
				Name:        "create_time",
				Description: "The creation timestamp of the metricThis field may not be present for older metrics",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "A description of this metric, which is used in documentation The maximum length of the description is 8000 characters",
				Type:        schema.TypeString,
			},
			{
				Name:        "filter",
				Description: "An advanced logs filter (https://cloudgooglecom/logging/docs/view/advanced_filters) which is used to match log entries Example: \"resourcetype=gae_app AND severity>=ERROR\" The maximum length of the filter is 20000 characters",
				Type:        schema.TypeString,
			},
			{
				Name:        "label_extractors",
				Description: "A map from a label key string to an extractor expression which is used to extract data from a log entry field and assign as the label value Each label key specified in the LabelDescriptor must have an associated extractor expression in this map The syntax of the extractor expression is the same as for the value_extractor fieldThe extracted value is converted to the type defined in the label descriptor If the either the extraction or the type conversion fails, the label will have a default value The default value for a string label is an empty string, for an integer label its 0, and for a boolean label its falseNote that there are upper bounds on the maximum number of labels and the number of active time series that are allowed in a project",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "metric_descriptor_description",
				Description: "A detailed description of the metric, which can be used in documentation",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MetricDescriptor.Description"),
			},
			{
				Name:        "metric_descriptor_display_name",
				Description: "A concise name for the metric, which can be displayed in user interfaces Use sentence case without an ending period, for example \"Request count\" This field is optional but it is recommended to be set for any metrics associated with user-visible concepts, such as Quota",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MetricDescriptor.DisplayName"),
			},
			{
				Name:        "metric_descriptor_launch_stage",
				Description: "The launch stage of the metric definition",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MetricDescriptor.LaunchStage"),
			},
			{
				Name:        "metric_descriptor_metadata_ingest_delay",
				Description: "The delay of data points caused by ingestion Data points older than this age are guaranteed to be ingested and available to be read, excluding data loss due to errors",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MetricDescriptor.Metadata.IngestDelay"),
			},
			{
				Name:        "metric_descriptor_metadata_sample_period",
				Description: "The sampling period of metric data points For metrics which are written periodically, consecutive data points are stored at this time interval, excluding data loss due to errors Metrics with a higher granularity have a smaller sampling period",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MetricDescriptor.Metadata.SamplePeriod"),
			},
			{
				Name:        "metric_descriptor_metric_kind",
				Description: "Whether the metric records instantaneous values, changes to a value, etc Some combinations of metric_kind and value_type might not be supported  Possible values:   \"METRIC_KIND_UNSPECIFIED\" - Do not use this default value   \"GAUGE\" - An instantaneous measurement of a value   \"DELTA\" - The change in a value during a time interval   \"CUMULATIVE\" - A value accumulated over a time interval Cumulative measurements in a time series should have the same start time and increasing end times, until an event resets the cumulative value to zero and sets a new start time for the following points",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MetricDescriptor.MetricKind"),
			},
			{
				Name:          "metric_descriptor_monitored_resource_types",
				Description:   "Read-only If present, then a time series, which is identified partially by a metric type and a MonitoredResourceDescriptor, that is associated with this metric type can only be associated with one of the monitored resource types listed here",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("MetricDescriptor.MonitoredResourceTypes"),
			},
			{
				Name:        "metric_descriptor_name",
				Description: "The resource name of the metric descriptor",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MetricDescriptor.Name"),
			},
			{
				Name:        "metric_descriptor_type",
				Description: "The metric type, including its DNS name prefix The type is not URL-encoded All user-defined metric types have the DNS name customgoogleapiscom or externalgoogleapiscom Metric types should use a natural hierarchical grouping For example: \"customgoogleapiscom/invoice/paid/amount\" \"externalgoogleapiscom/prometheus/up\" \"appenginegoogleapis",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MetricDescriptor.Type"),
			},
			{
				Name:        "metric_descriptor_unit",
				Description: "The units in which the metric value is reported",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MetricDescriptor.Unit"),
			},
			{
				Name:        "metric_descriptor_value_type",
				Description: "Whether the measurement is an integer, a floating-point number, etc Some combinations of metric_kind and value_type might not be supported  Possible values:   \"VALUE_TYPE_UNSPECIFIED\" - Do not use this default value   \"BOOL\" - The value is a boolean This value type can be used only if the metric kind is GAUGE   \"INT64\" - The value is a signed 64-bit integer   \"DOUBLE\" - The value is a double precision floating point number   \"STRING\" - The value is a text string This value type can be used only if the metric kind is GAUGE   \"DISTRIBUTION\" - The value is a Distribution   \"MONEY\" - The value is money",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MetricDescriptor.ValueType"),
			},
			{
				Name:        "name",
				Description: "The client-assigned metric identifier Examples: \"error_count\", \"nginx/requests\"Metric identifiers are limited to 100 characters and can include only the following characters: A-Z, a-z, 0-9, and the special characters _-,+!*',()%/ The forward-slash character (/) denotes a hierarchy of name pieces, and it cannot be the first character of the nameThe metric identifier in this field must not be URL-encoded (https://enwikipediaorg/wiki/Percent-encoding) However, when the metric identifier appears as the [METRIC_ID] part of a metric_name API parameter, then the metric identifier must be URL-encoded Example: \"projects/my-project/metrics/nginx%2Frequests\"",
				Type:        schema.TypeString,
			},
			{
				Name:        "update_time",
				Description: "The last update timestamp of the metricThis field may not be present for older metrics",
				Type:        schema.TypeString,
			},
			{
				Name:        "value_extractor",
				Description: "A value_extractor is required when using a distribution logs-based metric to extract the values to record from a log entry Two functions are supported for value extraction: EXTRACT(field) or REGEXP_EXTRACT(field, regex) The argument are: 1 field: The name of the log entry field from which the value is to be extracted 2 regex: A regular expression using the Google RE2 syntax (https://githubcom/google/re2/wiki/Syntax) with a single capture group to extract data from the specified log entry field The value of the field is converted to a string before applying the regex It is an error to specify a regex that does not include exactly one capture groupThe result of the extraction must be convertible to a double type, as the distribution always records double values If either the extraction or the conversion to double fails, then those values are not recorded in the distributionExample: REGEXP_EXTRACT(jsonPayloadrequest, \"*quantity=(\\d+)",
				Type:        schema.TypeString,
			},
			{
				Name:        "version",
				Description: "Deprecated The API version that created or updated this metric The v2 format is used by default and cannot be changed  Possible values:   \"V2\" - Logging API v2   \"V1\" - Logging API v1",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "gcp_logging_metric_descriptor_labels",
				Description: "A description of a label",
				Resolver:    fetchLoggingMetricDescriptorLabels,
				Columns: []schema.Column{
					{
						Name:        "metric_cq_id",
						Description: "Unique ID of gcp_logging_metrics table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "description",
						Description: "A human-readable description for the label",
						Type:        schema.TypeString,
					},
					{
						Name:        "key",
						Description: "The label key",
						Type:        schema.TypeString,
					},
					{
						Name:        "value_type",
						Description: "The type of data that can be assigned to the label  Possible values:   \"STRING\" - A variable-length string This is the default   \"BOOL\" - Boolean; true or false   \"INT64\" - A 64-bit signed integer",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchLoggingMetrics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Logging.Projects.Metrics.
			List(fmt.Sprintf("projects/%s", c.ProjectId)).
			PageToken(nextPageToken)
		list, err := c.RetryingDo(ctx, call)
		if err != nil {
			return diag.WrapError(err)
		}
		output := list.(*logging.ListLogMetricsResponse)

		res <- output.Metrics
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func fetchLoggingMetricDescriptorLabels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*logging.LogMetric)
	if p.MetricDescriptor == nil {
		return nil
	}

	res <- p.MetricDescriptor.Labels
	return nil
}
