package lightsail

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"golang.org/x/sync/errgroup"
)

//go:generate cq-gen --resource databases --config gen.hcl --output .
func Databases() *schema.Table {
	return &schema.Table{
		Name:         "aws_lightsail_databases",
		Description:  "Describes a database",
		Resolver:     fetchLightsailDatabases,
		Multiplex:    client.ServiceAccountRegionMultiplexer("lightsail"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the database",
				Type:        schema.TypeString,
			},
			{
				Name:        "backup_retention_enabled",
				Description: "A Boolean value indicating whether automated backup retention is enabled for the database",
				Type:        schema.TypeBool,
			},
			{
				Name:        "ca_certificate_identifier",
				Description: "The certificate associated with the database",
				Type:        schema.TypeString,
			},
			{
				Name:        "created_at",
				Description: "The timestamp when the database was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "engine",
				Description: "The database software (for example, MySQL)",
				Type:        schema.TypeString,
			},
			{
				Name:        "engine_version",
				Description: "The database engine version (for example, 5723)",
				Type:        schema.TypeString,
			},
			{
				Name:        "hardware_cpu_count",
				Description: "The number of vCPUs for the database",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Hardware.CpuCount"),
			},
			{
				Name:        "hardware_disk_size_in_gb",
				Description: "The size of the disk for the database",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Hardware.DiskSizeInGb"),
			},
			{
				Name:        "hardware_ram_size_in_gb",
				Description: "The amount of RAM in GB for the database",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("Hardware.RamSizeInGb"),
			},
			{
				Name:        "latest_restorable_time",
				Description: "The latest point in time to which the database can be restored",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "availability_zone",
				Description: "The Availability Zone",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Location.AvailabilityZone"),
			},
			{
				Name:        "master_database_name",
				Description: "The name of the master database created when the Lightsail database resource is created",
				Type:        schema.TypeString,
			},
			{
				Name:        "master_endpoint_address",
				Description: "Specifies the DNS address of the database",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MasterEndpoint.Address"),
			},
			{
				Name:        "master_endpoint_port",
				Description: "Specifies the port that the database is listening on",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("MasterEndpoint.Port"),
			},
			{
				Name:        "master_username",
				Description: "The master user name of the database",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The unique name of the database resource in Lightsail",
				Type:        schema.TypeString,
			},
			{
				Name:        "parameter_apply_status",
				Description: "The status of parameter updates for the database",
				Type:        schema.TypeString,
			},
			{
				Name:          "pending_modified_values",
				Description:   "Describes pending database value modifications",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:        "preferred_backup_window",
				Description: "The daily time range during which automated backups are created for the database (for example, 16:00-16:30)",
				Type:        schema.TypeString,
			},
			{
				Name:        "preferred_maintenance_window",
				Description: "The weekly time range during which system maintenance can occur on the database In the format ddd:hh24:mi-ddd:hh24:mi",
				Type:        schema.TypeString,
			},
			{
				Name:        "publicly_accessible",
				Description: "A Boolean value indicating whether the database is publicly accessible",
				Type:        schema.TypeBool,
			},
			{
				Name:        "relational_database_blueprint_id",
				Description: "The blueprint ID for the database",
				Type:        schema.TypeString,
			},
			{
				Name:        "relational_database_bundle_id",
				Description: "The bundle ID for the database",
				Type:        schema.TypeString,
			},
			{
				Name:        "resource_type",
				Description: "The Lightsail resource type for the database (for example, RelationalDatabase)",
				Type:        schema.TypeString,
			},
			{
				Name:          "secondary_availability_zone",
				Description:   "Describes the secondary Availability Zone of a high availability database",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "state",
				Description: "Describes the current state of the database",
				Type:        schema.TypeString,
			},
			{
				Name:        "support_code",
				Description: "The support code for the database",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The tag keys and optional values for the resource",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_lightsail_database_pending_maintenance_actions",
				Description: "Describes a pending database maintenance action",
				Resolver:    fetchLightsailDatabasePendingMaintenanceActions,
				Columns: []schema.Column{
					{
						Name:        "database_cq_id",
						Description: "Unique CloudQuery ID of aws_lightsail_databases table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "action",
						Description: "The type of pending database maintenance action",
						Type:        schema.TypeString,
					},
					{
						Name:          "current_apply_date",
						Description:   "The effective date of the pending database maintenance action",
						Type:          schema.TypeTimestamp,
						IgnoreInTests: true,
					},
					{
						Name:        "description",
						Description: "Additional detail about the pending database maintenance action",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_lightsail_database_parameters",
				Description: "Describes the parameters of a database",
				Resolver:    fetchLightsailDatabaseParameters,
				Columns: []schema.Column{
					{
						Name:        "database_cq_id",
						Description: "Unique CloudQuery ID of aws_lightsail_databases table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "allowed_values",
						Description: "Specifies the valid range of values for the parameter",
						Type:        schema.TypeString,
					},
					{
						Name:        "apply_method",
						Description: "Indicates when parameter updates are applied",
						Type:        schema.TypeString,
					},
					{
						Name:        "apply_type",
						Description: "Specifies the engine-specific parameter type",
						Type:        schema.TypeString,
					},
					{
						Name:        "data_type",
						Description: "Specifies the valid data type for the parameter",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "Provides a description of the parameter",
						Type:        schema.TypeString,
					},
					{
						Name:        "is_modifiable",
						Description: "A Boolean value indicating whether the parameter can be modified",
						Type:        schema.TypeBool,
					},
					{
						Name:        "name",
						Description: "Specifies the name of the parameter",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ParameterName"),
					},
					{
						Name:        "value",
						Description: "Specifies the value of the parameter",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ParameterValue"),
					},
				},
			},
			{
				Name:        "aws_lightsail_database_events",
				Description: "Describes an event for a database",
				Resolver:    fetchLightsailDatabaseEvents,
				Columns: []schema.Column{
					{
						Name:        "database_cq_id",
						Description: "Unique CloudQuery ID of aws_lightsail_databases table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "created_at",
						Description: "The timestamp when the database event was created",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "event_categories",
						Description: "The category that the database event belongs to",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "message",
						Description: "The message of the database event",
						Type:        schema.TypeString,
					},
					{
						Name:        "resource",
						Description: "The database that the database event relates to",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_lightsail_database_log_events",
				Resolver:      fetchLightsailDatabaseLogEvents,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "database_cq_id",
						Description: "Unique CloudQuery ID of aws_lightsail_databases table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "created_at",
						Description: "The timestamp when the database log event was created",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("LogEvent.CreatedAt"),
					},
					{
						Name:        "message",
						Description: "The message of the database log event",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LogEvent.Message"),
					},
					{
						Name:        "log_stream_name",
						Description: "An object describing the result of your get relational database log streams request",
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

func fetchLightsailDatabases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input lightsail.GetRelationalDatabasesInput
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	for {
		response, err := svc.GetRelationalDatabases(ctx, &input, func(options *lightsail.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.RelationalDatabases
		if aws.ToString(response.NextPageToken) == "" {
			break
		}
		input.PageToken = response.NextPageToken
	}
	return nil
}
func fetchLightsailDatabasePendingMaintenanceActions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RelationalDatabase)
	res <- r.PendingMaintenanceActions
	return nil
}
func fetchLightsailDatabaseParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RelationalDatabase)
	input := lightsail.GetRelationalDatabaseParametersInput{
		RelationalDatabaseName: r.Name,
	}
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	for {
		response, err := svc.GetRelationalDatabaseParameters(ctx, &input, func(options *lightsail.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Parameters
		if aws.ToString(response.NextPageToken) == "" {
			break
		}
		input.PageToken = response.NextPageToken
	}
	return nil
}
func fetchLightsailDatabaseEvents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RelationalDatabase)
	input := lightsail.GetRelationalDatabaseEventsInput{
		RelationalDatabaseName: r.Name,
		DurationInMinutes:      aws.Int32(20160), //two weeks
	}
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	for {
		response, err := svc.GetRelationalDatabaseEvents(ctx, &input, func(options *lightsail.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.RelationalDatabaseEvents
		if aws.ToString(response.NextPageToken) == "" {
			break
		}
		input.PageToken = response.NextPageToken
	}
	return nil
}
func fetchLightsailDatabaseLogEvents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RelationalDatabase)
	input := lightsail.GetRelationalDatabaseLogStreamsInput{
		RelationalDatabaseName: r.Name,
	}
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	streams, err := svc.GetRelationalDatabaseLogStreams(ctx, &input, func(options *lightsail.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	endTime := time.Now()
	startTime := endTime.Add(-time.Hour * 24 * 14) //two weeks
	errs, ctx := errgroup.WithContext(ctx)
	errs.SetLimit(MaxGoroutines)
	for _, s := range streams.LogStreams {
		func(database, stream string, startTime, endTime time.Time) {
			errs.Go(func() error {
				return fetchLogEvents(ctx, res, c, database, stream, startTime, endTime)
			})
		}(*r.Name, s, startTime, endTime)
	}
	err = errs.Wait()
	if err != nil {
		return diag.WrapError(err)
	}
	return nil
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func fetchLogEvents(ctx context.Context, res chan<- interface{}, c *client.Client, database, stream string, startTime, endTime time.Time) error {
	svc := c.Services().Lightsail
	input := lightsail.GetRelationalDatabaseLogEventsInput{
		RelationalDatabaseName: &database,
		LogStreamName:          &stream,
		StartTime:              &startTime,
		EndTime:                &endTime,
	}
	for {
		response, err := svc.GetRelationalDatabaseLogEvents(ctx, &input, func(options *lightsail.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.ResourceLogEvents
		for _, e := range response.ResourceLogEvents {
			res <- LogEventWrapper{e, stream}
		}
		if aws.ToString(response.NextForwardToken) == "" || len(response.ResourceLogEvents) == 0 {
			break
		}
		input.PageToken = response.NextForwardToken
	}
	return nil
}
