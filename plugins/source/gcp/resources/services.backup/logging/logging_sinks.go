package logging

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/logging/v2"
)

func LoggingSinks() *schema.Table {
	return &schema.Table{
		Name:        "gcp_logging_sinks",
		Description: "Describes a sink used to export log entries to one of the following destinations in any project: a Cloud Storage bucket, a BigQuery dataset, a Cloud Pub/Sub topic or a Cloud Logging Bucket A logs filter controls which log entries are exported The sink must be created within a project, organization, billing account, or folder",
		Resolver:    fetchLoggingSinks,
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
				Name:        "bigquery_options_use_partitioned_tables",
				Description: "Whether to use BigQuery's partition tables (https://cloudgooglecom/bigquery/docs/partitioned-tables) By default, Logging creates dated tables based on the log entries' timestamps, eg syslog_20170523 With partitioned tables the date suffix is no longer present and special query syntax (https://cloudgooglecom/bigquery/docs/querying-partitioned-tables) has to be used instead In both cases, tables are sharded based on UTC timezone",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("BigqueryOptions.UsePartitionedTables"),
			},
			{
				Name:        "bigquery_options_uses_timestamp_column_partitioning",
				Description: "True if new timestamp column based partitioning is in use, false if legacy ingestion-time partitioning is in use All new sinks will have this field set true and will use timestamp column based partitioning If use_partitioned_tables is false, this value has no meaning and will be false Legacy sinks using partitioned tables will have this field set to false",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("BigqueryOptions.UsesTimestampColumnPartitioning"),
			},
			{
				Name:        "create_time",
				Description: "The creation timestamp of the sink, This field may not be present for older sinks",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "A description of this sink The maximum length of the description is 8000 characters",
				Type:        schema.TypeString,
			},
			{
				Name:        "destination",
				Description: "The export destination: \"storagegoogleapiscom/[GCS_BUCKET]\" \"bigquerygoogleapiscom/projects/[PROJECT_ID]/datasets/[DATASET]\" \"pubsubgoogleapiscom/projects/[PROJECT_ID]/topics/[TOPIC_ID]\" The sink's writer_identity, set when the sink is created, must have permission to write to the destination or else the log entries are not exported For more information, see Exporting Logs with Sinks (https://cloudgooglecom/logging/docs/api/tasks/exporting-logs)",
				Type:        schema.TypeString,
			},
			{
				Name:        "disabled",
				Description: "If set to True, then this sink is disabled and it does not export any log entries",
				Type:        schema.TypeBool,
			},
			{
				Name:        "filter",
				Description: "An advanced logs filter (https://cloudgooglecom/logging/docs/view/advanced-queries) The only exported log entries are those that are in the resource owning the sink and that match the filter",
				Type:        schema.TypeString,
			},
			{
				Name:        "include_children",
				Description: "This field applies only to sinks owned by organizations and folders If the field is false, the default, only the logs owned by the sink's parent resource are available for export If the field is true, then logs from all the projects, folders, and billing accounts contained in the sink's parent resource are also available for export Whether a particular log entry from the children is exported depends on the sink's filter expression For example, if this field is true, then the filter resourcetype=gce_instance would export all Compute Engine VM instance log entries from all projects in the sink's parent To only export entries from certain child projects, filter on the project part of the log name: logName:(\"projects/test-project1/\" OR \"projects/test-project2/\") AND resource",
				Type:        schema.TypeBool,
			},
			{
				Name:        "name",
				Description: "The client-assigned sink identifier, unique within the project Example: \"my-syslog-errors-to-pubsub\" Sink identifiers are limited to 100 characters and can include only the following characters: upper and lower-case alphanumeric characters, underscores, hyphens, and periods First character has to be alphanumeric",
				Type:        schema.TypeString,
			},
			{
				Name:        "output_version_format",
				Description: "Deprecated This field is unused  Possible values:   \"VERSION_FORMAT_UNSPECIFIED\" - An unspecified format version that will default to V2   \"V2\" - LogEntry version 2 format   \"V1\" - LogEntry version 1 format",
				Type:        schema.TypeString,
			},
			{
				Name:        "update_time",
				Description: "The last update timestamp of the sinkThis field may not be present for older sinks",
				Type:        schema.TypeString,
			},
			{
				Name:        "writer_identity",
				Description: "An IAM identity—a service account or group—under which Logging writes the exported log entries to the sink's destination This field is set by sinkscreate and sinksupdate based on the value of unique_writer_identity in those methodsUntil you grant this identity write-access to the destination, log entry exports from this sink will fail For more information, see Granting Access for a Resource (https://cloudgooglecom/iam/docs/granting-roles-to-service-accounts#granting_access_to_a_service_account_for_a_resource) Consult the destination service's documentation to determine the appropriate IAM roles to assign to the identity",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "gcp_logging_sink_exclusions",
				Description: "Specifies a set of log entries that are not to be stored in Logging If your GCP resource receives a large volume of logs, you can use exclusions to reduce your chargeable logs Exclusions are processed after log sinks, so you can export log entries before they are excluded Note that organization-level and folder-level exclusions don't apply to child resources, and that you can't exclude audit log entries",
				Resolver:    fetchLoggingSinkExclusions,
				Columns: []schema.Column{
					{
						Name:        "sink_cq_id",
						Description: "Unique ID of gcp_logging_sinks table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "create_time",
						Description: "The creation timestamp of the exclusionThis field may not be present for older exclusions",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "A description of this exclusion",
						Type:        schema.TypeString,
					},
					{
						Name:        "disabled",
						Description: "If set to True, then this exclusion is disabled and it does not exclude any log entries You can update an exclusion to change the value of this field",
						Type:        schema.TypeBool,
					},
					{
						Name:        "filter",
						Description: "An advanced logs filter (https://cloudgooglecom/logging/docs/view/advanced-queries) that matches the log entries to be excluded By using the sample function (https://cloudgooglecom/logging/docs/view/advanced-queries#sample), you can exclude less than 100% of the matching log entries For example, the following query matches 99% of low-severity log entries from Google Cloud Storage buckets:\"resourcetype=gcs_bucket severity<ERROR sample(insertId, 0",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "A client-assigned identifier, such as \"load-balancer-exclusion\" Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods First character has to be alphanumeric",
						Type:        schema.TypeString,
					},
					{
						Name:        "update_time",
						Description: "The last update timestamp of the exclusionThis field may not be present for older exclusions",
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
func fetchLoggingSinks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Logging.Sinks.
			List(fmt.Sprintf("projects/%s", c.ProjectId)).
			PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		res <- output.Sinks
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func fetchLoggingSinkExclusions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*logging.LogSink)
	res <- p.Exclusions
	return nil
}
