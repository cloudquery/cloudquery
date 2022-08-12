package logging

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/logging/v2"
)

func LoggingMetrics() *schema.Table {
	return &schema.Table{
		Name:        "gcp_logging_metrics",
		Description: "Describes a logs-based metric The value of the metric is the number of log entries that match a logs filter in a given time intervalLogs-based metrics can also be used to extract values from logs and create a distribution of the values The distribution records the statistics of the extracted values along with an optional histogram of the values as specified by the bucket options",
		Resolver:    fetchLoggingMetrics,
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
				Description: "The units in which the metric value is reported It is only applicable if the value_type is INT64, DOUBLE, or DISTRIBUTION The unit defines the representation of the stored metric valuesDifferent systems might scale the values to be more easily displayed (so a value of 002kBy might be displayed as 20By, and a value of 3523kBy might be displayed as 35MBy) However, if the unit is kBy, then the value of the metric is always in thousands of bytes, no matter how it might be displayedIf you want a custom metric to record the exact number of CPU-seconds used by a job, you can create an INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently 1s{CPU} or just s) If the job uses 12,005 CPU-seconds, then the value is written as 12005Alternatively, if you want a custom metric to record data in a more granular way, you can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value 12005 (which is 12005/1000), or use Kis{CPU} and write 11723 (which is 12005/1024)The supported units are a subset of The Unified Code for Units of Measure (https://unitsofmeasureorg/ucumhtml) standard:Basic units (UNIT) bit bit By byte s second min minute h hour d day 1 dimensionlessPrefixes (PREFIX) k kilo (10^3) M mega (10^6) G giga (10^9) T tera (10^12) P peta (10^15) E exa (10^18) Z zetta (10^21) Y yotta (10^24) m milli (10^-3) u micro (10^-6) n nano (10^-9) p pico (10^-12) f femto (10^-15) a atto (10^-18) z zepto (10^-21) y yocto (10^-24) Ki kibi (2^10) Mi mebi (2^20) Gi gibi (2^30) Ti tebi (2^40) Pi pebi (2^50)GrammarThe grammar also includes these connectors: / division or ratio (as an infix operator) For examples, kBy/{email} or MiBy/10ms (although you should almost never have /s in a metric unit; rates should always be computed at query time from the underlying cumulative or delta value)  multiplication or composition (as an infix operator) For examples, GByd or k{watt}hThe grammar for a unit is as follows: Expression = Component { \"\" Component } { \"/\" Component } ; Component = ( [ PREFIX ] UNIT | \"%\" ) [ Annotation ] | Annotation | \"1\" ; Annotation = \"{\" NAME \"}\" ; Notes: Annotation is just a comment if it follows a UNIT If the annotation is used alone, then the unit is equivalent to 1 For examples, {request}/s == 1/s, By{transmitted}/s == By/s NAME is a sequence of non-blank printable ASCII characters not containing { or } 1 represents a unitary dimensionless unit (https://enwikipediaorg/wiki/Dimensionless_quantity) of 1, such as in 1/s It is typically used when none of the basic units are appropriate For example, \"new users per day\" can be represented as 1/d or {new-users}/d (and a metric value 5 would mean \"5 new users) Alternatively, \"thousands of page views per day\" would be represented as 1000/d or k1/d or k{page_views}/d (and a metric value of 53 would mean \"5300 page views per day\") % represents dimensionless value of 1/100, and annotates values giving a percentage (so the metric values are typically in the range of 0100, and a metric value 3 means \"3 percent\") 10^2% indicates a metric contains a ratio, typically in the range 01, that will be multiplied by 100 and displayed as a percentage (so a metric value 003 means \"3 percent\")",
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
		output, err := c.Services.Logging.Projects.Metrics.
			List(fmt.Sprintf("projects/%s", c.ProjectId)).
			PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

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
