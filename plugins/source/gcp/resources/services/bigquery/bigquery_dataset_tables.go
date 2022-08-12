package bigquery

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/bigquery/v2"
)

const maxGoroutines = 1

func BigqueryDatasetTables() *schema.Table {
	return &schema.Table{
		Name:        "gcp_bigquery_dataset_tables",
		Description: "Model options used for the first training run These options are immutable for subsequent training runs Default values are used for any options not specified in the input query",

		Resolver: listBigqueryDatasetTables,
		Columns: []schema.Column{
			{
				Name:     "dataset_cq_id",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIdResolver,
			},
			{
				Name:     "dataset_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("id"),
			},
			{
				Name:        "clustering_fields",
				Description: "One or more fields on which data should be clustered Only top-level, non-repeated, simple-type fields are supported When you cluster a table using multiple columns, the order of columns you specify is important The order of the specified columns determines the sort order of the data",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Clustering.Fields"),
			},
			{
				Name:        "creation_time",
				Description: "The time when this table was created, in milliseconds since the epoch",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "description",
				Description: "A user-friendly description of this table",
				Type:        schema.TypeString,
			},
			{
				Name:        "encryption_configuration_kms_key_name",
				Description: "Describes the Cloud KMS encryption key that will be used to protect destination BigQuery table The BigQuery Service Account associated with your project requires access to this encryption key",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EncryptionConfiguration.KmsKeyName"),
			},
			{
				Name:        "etag",
				Description: "A hash of the table metadata Used to ensure there were no concurrent modifications to the resource when attempting an update Not guaranteed to change when the table contents or the fields numRows, numBytes, numLongTermBytes or lastModifiedTime change",
				Type:        schema.TypeString,
			},
			{
				Name:        "expiration_time",
				Description: "The time when this table expires, in milliseconds since the epoch If not present, the table will persist indefinitely Expired tables will be deleted and their storage reclaimed The defaultTableExpirationMs property of the encapsulating dataset can be used to set a default expirationTime on newly created tables",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "external_data_configuration_autodetect",
				Description: "Try to detect schema and format options automatically Any option specified explicitly will be honored",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ExternalDataConfiguration.Autodetect"),
			},
			{
				Name:        "external_data_configuration_compression",
				Description: "The compression type of the data source Possible values include GZIP and NONE The default value is NONE This setting is ignored for Google Cloud Bigtable, Google Cloud Datastore backups and Avro formats",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExternalDataConfiguration.Compression"),
			},
			{
				Name:        "external_data_configuration_connection_id",
				Description: "Connection for external data source",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExternalDataConfiguration.ConnectionId"),
			},
			{
				Name:        "external_data_configuration_ignore_unknown_values",
				Description: "Indicates if BigQuery should allow extra values that are not represented in the table schema",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ExternalDataConfiguration.IgnoreUnknownValues"),
			},
			{
				Name:        "external_data_configuration_max_bad_records",
				Description: "The maximum number of bad records that BigQuery can ignore when reading data If the number of bad records exceeds this value, an invalid error is returned in the job result This is only valid for CSV, JSON, and Google Sheets The default value is 0, which requires that all records are valid This setting is ignored for Google Cloud Bigtable, Google Cloud Datastore backups and Avro formats",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ExternalDataConfiguration.MaxBadRecords"),
			},
			{
				Name:          "external_data_configuration_schema",
				Description:   "The schema for the data Schema is required for CSV and JSON formats Schema is disallowed for Google Cloud Bigtable, Cloud Datastore backups, and Avro formats",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
				Resolver:      resolveBigqueryDatasetTableExternalDataConfigurationSchema,
			},
			{
				Name:        "external_data_configuration_source_format",
				Description: "The data format For CSV files, specify \"CSV\" For Google sheets, specify \"GOOGLE_SHEETS\" For newline-delimited JSON, specify \"NEWLINE_DELIMITED_JSON\" For Avro files, specify \"AVRO\" For Google Cloud Datastore backups, specify \"DATASTORE_BACKUP\" [Beta] For Google Cloud Bigtable, specify \"BIGTABLE\"",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExternalDataConfiguration.SourceFormat"),
			},
			{
				Name:          "external_data_configuration_source_uris",
				Description:   "The fully-qualified URIs that point to your data in Google Cloud For Google Cloud Storage URIs: Each URI can contain one '*' wildcard character and it must come after the 'bucket' name Size limits related to load jobs apply to external data sources For Google Cloud Bigtable URIs: Exactly one URI can be specified and it has be a fully specified and valid HTTPS URL for a Google Cloud Bigtable table For Google Cloud Datastore backups, exactly one URI can be specified Also, the '*' wildcard character is not allowed",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("ExternalDataConfiguration.SourceUris"),
			},
			{
				Name:        "friendly_name",
				Description: "A descriptive name for this table",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "An opaque ID uniquely identifying the table",
				Type:        schema.TypeString,
			},
			{
				Name:        "kind",
				Description: "The type of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "labels",
				Description: "The labels associated with this table You can use these to organize and group your tables Label keys and values can be no longer than 63 characters, can only contain lowercase letters, numeric characters, underscores and dashes International characters are allowed Label values are optional Label keys must start with a letter and each label in the list must have a different key",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "last_modified_time",
				Description: "The time when this table was last modified, in milliseconds since the epoch",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "location",
				Description: "The geographic location where the table resides This value is inherited from the dataset",
				Type:        schema.TypeString,
			},
			{
				Name:        "materialized_view_enable_refresh",
				Description: "Enable automatic refresh of the materialized view when the base table is updated The default value is \"true\"",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("MaterializedView.EnableRefresh"),
			},
			{
				Name:        "materialized_view_last_refresh_time",
				Description: "The time when this materialized view was last modified, in milliseconds since the epoch",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("MaterializedView.LastRefreshTime"),
			},
			{
				Name:        "materialized_view_query",
				Description: "A query whose result is persisted",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MaterializedView.Query"),
			},
			{
				Name:        "materialized_view_refresh_interval_ms",
				Description: "The maximum frequency at which this materialized view will be refreshed The default value is \"1800000\" (30 minutes)",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("MaterializedView.RefreshIntervalMs"),
			},
			{
				Name:          "model_options_labels",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("Model.ModelOptions.Labels"),
			},
			{
				Name:     "model_options_loss_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Model.ModelOptions.LossType"),
			},
			{
				Name:     "model_options_model_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Model.ModelOptions.ModelType"),
			},
			{
				Name:        "num_bytes",
				Description: "The size of this table in bytes, excluding any data in the streaming buffer",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "num_long_term_bytes",
				Description: "The number of bytes in the table that are considered \"long-term storage\"",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "num_physical_bytes",
				Description: "The physical size of this table in bytes, excluding any data in the streaming buffer This includes compression and storage used for time travel",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "num_rows",
				Description: "The number of rows of data in this table, excluding any data in the streaming buffer",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "range_partitioning_field",
				Description: "The table is partitioned by this field The field must be a top-level NULLABLE/REQUIRED field The only supported type is INTEGER/INT64",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RangePartitioning.Field"),
			},
			{
				Name:        "range_partitioning_range_end",
				Description: "The end of range partitioning, exclusive",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("RangePartitioning.Range.End"),
			},
			{
				Name:        "range_partitioning_range_interval",
				Description: "The width of each interval",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("RangePartitioning.Range.Interval"),
			},
			{
				Name:        "range_partitioning_range_start",
				Description: "The start of range partitioning, inclusive",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("RangePartitioning.Range.Start"),
			},
			{
				Name:        "require_partition_filter",
				Description: "If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified",
				Type:        schema.TypeBool,
			},
			{
				Name:        "schema",
				Description: "Describes the schema of this table",
				Type:        schema.TypeJSON,
				Resolver:    resolveBigqueryDatasetTableSchema,
			},
			{
				Name:        "self_link",
				Description: "A URL that can be used to access this resource again",
				Type:        schema.TypeString,
			},
			{
				Name:        "streaming_buffer_estimated_bytes",
				Description: "A lower-bound estimate of the number of bytes currently in the streaming buffer",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("StreamingBuffer.EstimatedBytes"),
			},
			{
				Name:        "streaming_buffer_estimated_rows",
				Description: "A lower-bound estimate of the number of rows currently in the streaming buffer",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("StreamingBuffer.EstimatedRows"),
			},
			{
				Name:        "streaming_buffer_oldest_entry_time",
				Description: "Contains the timestamp of the oldest entry in the streaming buffer, in milliseconds since the epoch, if the streaming buffer is available",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("StreamingBuffer.OldestEntryTime"),
			},
			{
				Name:        "time_partitioning_expiration_ms",
				Description: "Number of milliseconds for which to keep the storage for partitions in the table The storage in a partition will have an expiration time of its partition time plus this value",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("TimePartitioning.ExpirationMs"),
			},
			{
				Name:        "time_partitioning_field",
				Description: "If not set, the table is partitioned by pseudo column, referenced via either '_PARTITIONTIME' as TIMESTAMP type, or '_PARTITIONDATE' as DATE type If field is specified, the table is instead partitioned by this field The field must be a top-level TIMESTAMP or DATE field Its mode must be NULLABLE or REQUIRED",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TimePartitioning.Field"),
			},
			{
				Name:     "time_partitioning_require_partition_filter",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("TimePartitioning.RequirePartitionFilter"),
			},
			{
				Name:        "time_partitioning_type",
				Description: "The supported types are DAY, HOUR, MONTH, and YEAR, which will generate one partition per day, hour, month, and year, respectively When the type is not specified, the default behavior is DAY",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TimePartitioning.Type"),
			},
			{
				Name:        "type",
				Description: "Describes the table type The following values are supported: TABLE: A normal BigQuery table VIEW: A virtual table defined by a SQL query SNAPSHOT: An immutable, read-only table that is a copy of another table MATERIALIZED_VIEW: SQL query whose result is persisted EXTERNAL: A table that references data stored in an external storage system, such as Google Cloud Storage The default value is TABLE",
				Type:        schema.TypeString,
			},
			{
				Name:        "view_query",
				Description: "A query that BigQuery executes when the view is referenced",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("View.Query"),
			},
			{
				Name:        "view_use_legacy_sql",
				Description: "Specifies whether to use BigQuery's legacy SQL for this view The default value is true If set to false, the view will use BigQuery's standard SQL: https://cloudgooglecom/bigquery/sql-reference/ Queries and views that reference this view must use the same flag value",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("View.UseLegacySql"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "gcp_bigquery_dataset_table_dataset_model_training_runs",
				Description:   "Training options used by this training run These options are mutable for subsequent training runs Default values are explicitly stored for options not specified in the input query of the first training run For subsequent training runs, any option not explicitly specified in the input query will be copied from the previous training run",
				Resolver:      fetchBigqueryDatasetTableDatasetModelTrainingRuns,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "dataset_table_cq_id",
						Description: "Unique ID of gcp_bigquery_dataset_tables table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "dataset_table_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "start_time",
						Description: "Training run start time in milliseconds since the epoch",
						Type:        schema.TypeString,
					},
					{
						Name:        "state",
						Description: "Different state applicable for a training run IN PROGRESS: Training run is in progress FAILED: Training run ended due to a non-retryable failure SUCCEEDED: Training run successfully completed CANCELLED: Training run cancelled by the user",
						Type:        schema.TypeString,
					},
					{
						Name:     "training_options_early_stop",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("TrainingOptions.EarlyStop"),
					},
					{
						Name:     "training_options_l1_reg",
						Type:     schema.TypeFloat,
						Resolver: schema.PathResolver("TrainingOptions.L1Reg"),
					},
					{
						Name:     "training_options_l2_reg",
						Type:     schema.TypeFloat,
						Resolver: schema.PathResolver("TrainingOptions.L2Reg"),
					},
					{
						Name:     "training_options_learn_rate",
						Type:     schema.TypeFloat,
						Resolver: schema.PathResolver("TrainingOptions.LearnRate"),
					},
					{
						Name:     "training_options_learn_rate_strategy",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("TrainingOptions.LearnRateStrategy"),
					},
					{
						Name:     "training_options_line_search_init_learn_rate",
						Type:     schema.TypeFloat,
						Resolver: schema.PathResolver("TrainingOptions.LineSearchInitLearnRate"),
					},
					{
						Name:     "training_options_max_iteration",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("TrainingOptions.MaxIteration"),
					},
					{
						Name:     "training_options_min_rel_progress",
						Type:     schema.TypeFloat,
						Resolver: schema.PathResolver("TrainingOptions.MinRelProgress"),
					},
					{
						Name:     "training_options_warm_start",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("TrainingOptions.WarmStart"),
					},
				},
			},
			{
				Name:          "gcp_bigquery_dataset_table_user_defined_functions",
				Description:   "This is used for defining User Defined Function (UDF) resources only when using legacy SQL",
				Resolver:      fetchBigqueryDatasetTableUserDefinedFunctions,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "dataset_table_cq_id",
						Description: "Unique ID of gcp_bigquery_dataset_tables table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "dataset_table_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "inline_code",
						Description: "An inline resource that contains code for a user-defined function (UDF) Providing a inline code resource is equivalent to providing a URI for a file containing the same code",
						Type:        schema.TypeString,
					},
					{
						Name:        "resource_uri",
						Description: "A code resource to load from a Google Cloud Storage URI (gs://bucket/path)",
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
func listBigqueryDatasetTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*bigquery.Dataset)
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.BigQuery.Tables.List(c.ProjectId, p.DatasetReference.DatasetId).Context(ctx).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}
		errs, ctx := errgroup.WithContext(ctx)
		errs.SetLimit(maxGoroutines)
		for _, t := range output.Tables {
			func(t *bigquery.TableListTables) {
				errs.Go(func() error {
					return fetchBigqueryDatasetTables(ctx, c, p, t, res)
				})
			}(t)
		}
		err = errs.Wait()
		if err != nil {
			return errors.WithStack(err)
		}

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}

func fetchBigqueryDatasetTables(ctx context.Context, c *client.Client, p *bigquery.Dataset, t *bigquery.TableListTables, res chan<- interface{}) error {
	item, err := c.Services.BigQuery.Tables.Get(c.ProjectId, p.DatasetReference.DatasetId, t.TableReference.TableId).Do()
	if err != nil {
		return errors.WithStack(err)
	}
	res <- item
	return nil
}

func resolveBigqueryDatasetTableExternalDataConfigurationSchema(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(*bigquery.Table)
	if p.ExternalDataConfiguration == nil || p.ExternalDataConfiguration.Schema == nil {
		return nil
	}

	s := make(map[string]interface{})
	for _, f := range p.ExternalDataConfiguration.Schema.Fields {
		s[f.Name] = f.Type
	}
	return errors.WithStack(resource.Set(c.Name, s))
}
func resolveBigqueryDatasetTableSchema(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(*bigquery.Table)
	if p.Schema == nil {
		return nil
	}

	s := make(map[string]interface{})
	for _, f := range p.Schema.Fields {
		s[f.Name] = f.Type
	}
	return errors.WithStack(resource.Set(c.Name, s))
}
func fetchBigqueryDatasetTableDatasetModelTrainingRuns(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*bigquery.Table)
	if p.Model == nil {
		return nil
	}

	res <- p.Model.TrainingRuns
	return nil
}
func fetchBigqueryDatasetTableUserDefinedFunctions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*bigquery.Table)
	if p.View == nil {
		return nil
	}

	res <- p.View.UserDefinedFunctionResources
	return nil
}
