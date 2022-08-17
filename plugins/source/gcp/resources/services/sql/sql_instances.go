package sql

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	sql "google.golang.org/api/sqladmin/v1beta4"
)

func SQLInstances() *schema.Table {
	return &schema.Table{
		Name:        "gcp_sql_instances",
		Description: "A Cloud SQL instance resource",
		Resolver:    fetchSqlInstances,
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
				Name:        "backend_type",
				Description: "*SECOND_GEN*: Cloud SQL database instance *EXTERNAL*: A database server that is not managed by Google This property is read-only; use the *tier* property in the *settings* object to determine the database type  Possible values:   \"SQL_BACKEND_TYPE_UNSPECIFIED\" - This is an unknown backend type for instance   \"FIRST_GEN\" - V1 speckle instance   \"SECOND_GEN\" - V2 speckle instance   \"EXTERNAL\" - On premises instance",
				Type:        schema.TypeString,
			},
			{
				Name:        "connection_name",
				Description: "Connection name of the Cloud SQL instance used in connection strings",
				Type:        schema.TypeString,
			},
			{
				Name:        "current_disk_size",
				Description: "The current disk usage of the instance in bytes This property has been deprecated by google API and might be null",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "database_version",
				Description: "The database engine type and version The *databaseVersion* field cannot be changed after instance creation MySQL instances: *MYSQL_8_0*, *MYSQL_5_7* (default), or *MYSQL_5_6* PostgreSQL instances: *POSTGRES_9_6*, *POSTGRES_10*, *POSTGRES_11*, *POSTGRES_12*, or *POSTGRES_13* (default) SQL Server instances: *SQLSERVER_2017_STANDARD* (default), *SQLSERVER_2017_ENTERPRISE*, *SQLSERVER_2017_EXPRESS*, or *SQLSERVER_2017_WEB*  Possible values:   \"SQL_DATABASE_VERSION_UNSPECIFIED\" - This is an unknown database version   \"MYSQL_5_1\" - The database version is MySQL 51   \"MYSQL_5_5\" - The database version is MySQL 55   \"MYSQL_5_6\" - The database version is MySQL 56   \"MYSQL_5_7\" - The database version is MySQL 57   \"POSTGRES_9_6\" - The database version is PostgreSQL 96   \"POSTGRES_11\" - The database version is PostgreSQL 11   \"SQLSERVER_2017_STANDARD\" - The database version is SQL Server 2017 Standard   \"SQLSERVER_2017_ENTERPRISE\" - The database version is SQL Server 2017 Enterprise   \"SQLSERVER_2017_EXPRESS\" - The database version is SQL Server 2017 Express   \"SQLSERVER_2017_WEB\" - The database version is SQL Server 2017 Web   \"POSTGRES_10\" - The database version is PostgreSQL 10   \"POSTGRES_12\" - The database version is PostgreSQL 12   \"MYSQL_8_0\" - The database version is MySQL 8   \"POSTGRES_13\" - The database version is PostgreSQL 13",
				Type:        schema.TypeString,
			},
			{
				Name:        "disk_encryption_configuration_kind",
				Description: "This is always *sql#diskEncryptionConfiguration*",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DiskEncryptionConfiguration.Kind"),
			},
			{
				Name:     "disk_encryption_configuration_kms_key_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskEncryptionConfiguration.KmsKeyName"),
			},
			{
				Name:        "disk_encryption_status_kind",
				Description: "This is always *sql#diskEncryptionStatus*",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DiskEncryptionStatus.Kind"),
			},
			{
				Name:     "disk_encryption_status_kms_key_version_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskEncryptionStatus.KmsKeyVersionName"),
			},
			{
				Name:        "etag",
				Description: "This field is deprecated and will be removed from a future version of the API Use the *settingssettingsVersion* field instead",
				Type:        schema.TypeString,
			},
			{
				Name:        "failover_replica_available",
				Description: "The availability status of the failover replica A false status indicates that the failover replica is out of sync The primary instance can only failover to the failover replica when the status is true",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("FailoverReplica.Available"),
			},
			{
				Name:        "failover_replica_name",
				Description: "The name of the failover replica If specified at instance creation, a failover replica is created for the instance The name doesn't include the project ID This property is applicable only to Second Generation instances",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("FailoverReplica.Name"),
			},
			{
				Name:        "gce_zone",
				Description: "The Compute Engine zone that the instance is currently serving from This value could be different from the zone that was specified when the instance was created if the instance has failed over to its secondary zone",
				Type:        schema.TypeString,
			},
			{
				Name:        "instance_type",
				Description: "The instance type This can be one of the following *CLOUD_SQL_INSTANCE*: A Cloud SQL instance that is not replicating from a primary instance *ON_PREMISES_INSTANCE*: An instance running on the customer's premises *READ_REPLICA_INSTANCE*: A Cloud SQL instance configured as a read-replica  Possible values:   \"SQL_INSTANCE_TYPE_UNSPECIFIED\" - This is an unknown Cloud SQL instance type   \"CLOUD_SQL_INSTANCE\" - A regular Cloud SQL instance   \"ON_PREMISES_INSTANCE\" - An instance running on the customer's premises that is not managed by Cloud SQL   \"READ_REPLICA_INSTANCE\" - A Cloud SQL instance acting as a read-replica",
				Type:        schema.TypeString,
			},
			{
				Name:        "ipv6_address",
				Description: "The IPv6 address assigned to the instance (Deprecated) This property was applicable only to First Generation instances",
				Type:        schema.TypeString,
			},
			{
				Name:        "kind",
				Description: "This is always *sql#instance*",
				Type:        schema.TypeString,
			},
			{
				Name:        "master_instance_name",
				Description: "The name of the instance which will act as primary in the replication setup",
				Type:        schema.TypeString,
			},
			{
				Name:        "max_disk_size",
				Description: "The maximum disk size of the instance in bytes",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "name",
				Description: "Name of the Cloud SQL instance This does not include the project ID",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "Alias of a name.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Name"),
			},
			{
				Name:        "on_premises_configuration_ca_certificate",
				Description: "PEM representation of the trusted CA's x509 certificate",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("OnPremisesConfiguration.CaCertificate"),
			},
			{
				Name:        "on_premises_configuration_client_certificate",
				Description: "PEM representation of the replica's x509 certificate",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("OnPremisesConfiguration.ClientCertificate"),
			},
			{
				Name:     "on_premises_configuration_client_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesConfiguration.ClientKey"),
			},
			{
				Name:        "on_premises_configuration_dump_file_path",
				Description: "The dump file to create the Cloud SQL replica",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("OnPremisesConfiguration.DumpFilePath"),
			},
			{
				Name:     "on_premises_configuration_host_port",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesConfiguration.HostPort"),
			},
			{
				Name:        "on_premises_configuration_kind",
				Description: "This is always *sql#onPremisesConfiguration*",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("OnPremisesConfiguration.Kind"),
			},
			{
				Name:        "on_premises_configuration_password",
				Description: "The password for connecting to on-premises instance",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("OnPremisesConfiguration.Password"),
			},
			{
				Name:        "on_premises_configuration_username",
				Description: "The username for connecting to on-premises instance",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("OnPremisesConfiguration.Username"),
			},
			{
				Name:        "project",
				Description: "The project ID of the project containing the Cloud SQL instance The Google apps domain is prefixed if applicable",
				Type:        schema.TypeString,
			},
			{
				Name:        "region",
				Description: "The geographical region Can be *us-central* (*FIRST_GEN* instances only) *us-central1* (*SECOND_GEN* instances only) *asia-east1* or *europe-west1* Defaults to *us-central* or *us-central1* depending on the instance type The region cannot be changed after instance creation",
				Type:        schema.TypeString,
			},
			{
				Name:        "failover_target",
				Description: "Specifies if the replica is the failover target If the field is set to *true* the replica will be designated as a failover replica In case the primary instance fails, the replica instance will be promoted as the new primary instance Only one replica can be specified as failover target, and the replica has to be in different zone with the primary instance",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ReplicaConfiguration.FailoverTarget"),
			},
			{
				Name:     "configuration_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicaConfiguration.Kind"),
			},
			{
				Name:        "mysql_replica_configuration_ca_certificate",
				Description: "PEM representation of the trusted CA's x509 certificate",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ReplicaConfiguration.MysqlReplicaConfiguration.CaCertificate"),
			},
			{
				Name:        "mysql_replica_configuration_client_certificate",
				Description: "PEM representation of the replica's x509 certificate",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ReplicaConfiguration.MysqlReplicaConfiguration.ClientCertificate"),
			},
			{
				Name:     "mysql_replica_configuration_client_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicaConfiguration.MysqlReplicaConfiguration.ClientKey"),
			},
			{
				Name:        "mysql_replica_configuration_connect_retry_interval",
				Description: "Seconds to wait between connect retries MySQL's default is 60 seconds",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ReplicaConfiguration.MysqlReplicaConfiguration.ConnectRetryInterval"),
			},
			{
				Name:        "mysql_replica_configuration_dump_file_path",
				Description: "Path to a SQL dump file in Google Cloud Storage from which the replica instance is to be created The URI is in the form gs://bucketName/fileName Compressed gzip files (gz) are also supported Dumps have the binlog co-ordinates from which replication begins This can be accomplished by setting --master-data to 1 when using mysqldump",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ReplicaConfiguration.MysqlReplicaConfiguration.DumpFilePath"),
			},
			{
				Name:        "mysql_replica_configuration_kind",
				Description: "This is always *sql#mysqlReplicaConfiguration*",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ReplicaConfiguration.MysqlReplicaConfiguration.Kind"),
			},
			{
				Name:        "mysql_replica_configuration_master_heartbeat_period",
				Description: "Interval in milliseconds between replication heartbeats",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ReplicaConfiguration.MysqlReplicaConfiguration.MasterHeartbeatPeriod"),
			},
			{
				Name:        "mysql_replica_configuration_password",
				Description: "The password for the replication connection",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ReplicaConfiguration.MysqlReplicaConfiguration.Password"),
			},
			{
				Name:        "mysql_replica_configuration_ssl_cipher",
				Description: "A list of permissible ciphers to use for SSL encryption",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ReplicaConfiguration.MysqlReplicaConfiguration.SslCipher"),
			},
			{
				Name:        "mysql_replica_configuration_username",
				Description: "The username for the replication connection",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ReplicaConfiguration.MysqlReplicaConfiguration.Username"),
			},
			{
				Name:        "mysql_replica_configuration_verify_server_certificate",
				Description: "Whether or not to check the primary instance's Common Name value in the certificate that it sends during the SSL handshake",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ReplicaConfiguration.MysqlReplicaConfiguration.VerifyServerCertificate"),
			},
			{
				Name:          "replica_names",
				Description:   "The replicas of the instance",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "root_password",
				Description: "Initial root password Use only on creation",
				Type:        schema.TypeString,
			},
			{
				Name:        "satisfies_pzs",
				Description: "The status indicating if instance satisfiesPzs Reserved for future use",
				Type:        schema.TypeBool,
			},
			{
				Name:     "scheduled_maintenance_can_defer",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ScheduledMaintenance.CanDefer"),
			},
			{
				Name:        "scheduled_maintenance_can_reschedule",
				Description: "If the scheduled maintenance can be rescheduled",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ScheduledMaintenance.CanReschedule"),
			},
			{
				Name:        "scheduled_maintenance_start_time",
				Description: "The start time of any upcoming scheduled maintenance for this instance",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ScheduledMaintenance.StartTime"),
			},
			{
				Name:        "secondary_gce_zone",
				Description: "The Compute Engine zone that the failover instance is currently serving from for a regional instance This value could be different from the zone that was specified when the instance was created if the instance has failed over to its secondary/failover zone Reserved for future use",
				Type:        schema.TypeString,
			},
			{
				Name:        "self_link",
				Description: "The URI of this resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "server_ca_cert",
				Description: "PEM representation",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerCaCert.Cert"),
			},
			{
				Name:        "server_ca_cert_cert_serial_number",
				Description: "Serial number, as extracted from the certificate",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerCaCert.CertSerialNumber"),
			},
			{
				Name:        "server_ca_cert_common_name",
				Description: "User supplied name Constrained to [a-zA-Z-_ ]+",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerCaCert.CommonName"),
			},
			{
				Name:        "server_ca_cert_create_time",
				Description: "The time when the certificate was created in RFC 3339 format, for example *2012-11-15T16:19:00",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerCaCert.CreateTime"),
			},
			{
				Name:        "server_ca_cert_expiration_time",
				Description: "The time when the certificate expires in RFC 3339 format, for example *2012-11-15T16:19:00094Z*",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerCaCert.ExpirationTime"),
			},
			{
				Name:        "server_ca_cert_instance",
				Description: "Name of the database instance",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerCaCert.Instance"),
			},
			{
				Name:        "server_ca_cert_kind",
				Description: "This is always *sql#sslCert*",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerCaCert.Kind"),
			},
			{
				Name:        "server_ca_cert_self_link",
				Description: "The URI of this resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerCaCert.SelfLink"),
			},
			{
				Name:        "server_ca_cert_sha1_fingerprint",
				Description: "Sha1 Fingerprint",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerCaCert.Sha1Fingerprint"),
			},
			{
				Name:        "service_account_email_address",
				Description: "The service account email address assigned to the instance This property is applicable only to Second Generation instances",
				Type:        schema.TypeString,
			},
			{
				Name:        "settings_activation_policy",
				Description: "The activation policy specifies when the instance is activated; it is applicable only when the instance state is RUNNABLE Valid values: *ALWAYS*: The instance is on, and remains so even in the absence of connection requests *NEVER*: The instance is off; it is not activated, even if a connection request arrives  Possible values:   \"SQL_ACTIVATION_POLICY_UNSPECIFIED\" - Unknown activation plan   \"ALWAYS\" - The instance is always up and running   \"NEVER\" - The instance never starts   \"ON_DEMAND\" - The instance starts upon receiving requests",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Settings.ActivationPolicy"),
			},
			{
				Name:        "settings_active_directory_config_domain",
				Description: "The name of the domain (eg, mydomaincom)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Settings.ActiveDirectoryConfig.Domain"),
			},
			{
				Name:        "settings_active_directory_config_kind",
				Description: "This is always sql#activeDirectoryConfig",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Settings.ActiveDirectoryConfig.Kind"),
			},
			{
				Name:        "settings_authorized_gae_applications",
				Description: "The App Engine app IDs that can access this instance (Deprecated) Applied to First Generation instances only",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Settings.AuthorizedGaeApplications"),
			},
			{
				Name:        "settings_availability_type",
				Description: "Availability type Potential values: *ZONAL*: The instance serves data from only one zone Outages in that zone affect data accessibility *REGIONAL*: The instance can serve data from more than one zone in a region (it is highly available) For more information, see Overview of the High Availability Configuration  Possible values:   \"SQL_AVAILABILITY_TYPE_UNSPECIFIED\" - This is an unknown Availability type   \"ZONAL\" - Zonal available instance   \"REGIONAL\" - Regional available instance",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Settings.AvailabilityType"),
			},
			{
				Name:        "settings_backup_retention_settings_retained_backups",
				Description: "Depending on the value of retention_unit, this is used to determine if a backup needs to be deleted If retention_unit is 'COUNT', we will retain this many backups",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Settings.BackupConfiguration.BackupRetentionSettings.RetainedBackups"),
			},
			{
				Name:        "settings_backup_retention_settings_retention_unit",
				Description: "The unit that 'retained_backups' represents  Possible values:   \"RETENTION_UNIT_UNSPECIFIED\" - Backup retention unit is unspecified, will be treated as COUNT   \"COUNT\" - Retention will be by count, eg \"retain the most recent 7 backups\"",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Settings.BackupConfiguration.BackupRetentionSettings.RetentionUnit"),
			},
			{
				Name:     "settings_backup_binary_log_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Settings.BackupConfiguration.BinaryLogEnabled"),
			},
			{
				Name:     "settings_backup_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Settings.BackupConfiguration.Enabled"),
			},
			{
				Name:     "settings_backup_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Settings.BackupConfiguration.Kind"),
			},
			{
				Name:     "settings_backup_location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Settings.BackupConfiguration.Location"),
			},
			{
				Name:     "settings_backup_point_in_time_recovery_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Settings.BackupConfiguration.PointInTimeRecoveryEnabled"),
			},
			{
				Name:     "settings_backup_replication_log_archiving_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Settings.BackupConfiguration.ReplicationLogArchivingEnabled"),
			},
			{
				Name:     "settings_backup_start_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Settings.BackupConfiguration.StartTime"),
			},
			{
				Name:     "settings_backup_transaction_log_retention_days",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Settings.BackupConfiguration.TransactionLogRetentionDays"),
			},
			{
				Name:        "settings_collation",
				Description: "The name of server Instance collation",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Settings.Collation"),
			},
			{
				Name:        "settings_crash_safe_replication_enabled",
				Description: "Configuration specific to read replica instances Indicates whether database flags for crash-safe replication are enabled This property was only applicable to First Generation instances",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Settings.CrashSafeReplicationEnabled"),
			},
			{
				Name:        "settings_data_disk_size_gb",
				Description: "The size of data disk, in GB The data disk size minimum is 10GB",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Settings.DataDiskSizeGb"),
			},
			{
				Name:        "settings_data_disk_type",
				Description: "The type of data disk: PD_SSD (default) or PD_HDD Not used for First Generation instances  Possible values:   \"SQL_DATA_DISK_TYPE_UNSPECIFIED\" - This is an unknown data disk type   \"PD_SSD\" - An SSD data disk   \"PD_HDD\" - An HDD data disk   \"OBSOLETE_LOCAL_SSD\" - This field is deprecated and will be removed from a future version of the API",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Settings.DataDiskType"),
			},
			{
				Name:        "settings_database_flags",
				Description: "The database flags passed to the instance at startup",
				Type:        schema.TypeJSON,
				Resolver:    resolveSQLInstanceSettingsDatabaseFlags,
			},
			{
				Name:        "settings_database_replication_enabled",
				Description: "Configuration specific to read replica instances Indicates whether replication is enabled or not",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Settings.DatabaseReplicationEnabled"),
			},
			{
				Name:        "settings_insights_config_query_insights_enabled",
				Description: "Whether Query Insights feature is enabled",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Settings.InsightsConfig.QueryInsightsEnabled"),
			},
			{
				Name:        "settings_insights_config_query_string_length",
				Description: "Maximum query length stored in bytes Default value: 1024 bytes Range: 256-4500 bytes Query length more than this field value will be truncated to this value When unset, query length will be the default value Changing query length will restart the database",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Settings.InsightsConfig.QueryStringLength"),
			},
			{
				Name:        "settings_insights_config_record_application_tags",
				Description: "Whether Query Insights will record application tags from query when enabled",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Settings.InsightsConfig.RecordApplicationTags"),
			},
			{
				Name:        "settings_insights_config_record_client_address",
				Description: "Whether Query Insights will record client address when enabled",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Settings.InsightsConfig.RecordClientAddress"),
			},
			{
				Name:        "settings_ip_configuration_ipv4_enabled",
				Description: "Whether the instance is assigned a public IP address or not",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Settings.IpConfiguration.Ipv4Enabled"),
			},
			{
				Name:        "settings_ip_configuration_private_network",
				Description: "The resource link for the VPC network from which the Cloud SQL instance is accessible for private IP For example, */projects/myProject/global/networks/default* This setting can be updated, but it cannot be removed after it is set",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Settings.IpConfiguration.PrivateNetwork"),
			},
			{
				Name:        "settings_ip_configuration_require_ssl",
				Description: "Whether SSL connections over IP are enforced or not",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Settings.IpConfiguration.RequireSsl"),
			},
			{
				Name:        "settings_kind",
				Description: "This is always *sql#settings*",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Settings.Kind"),
			},
			{
				Name:        "settings_location_preference_follow_gae_application",
				Description: "The App Engine application to follow, it must be in the same region as the Cloud SQL instance",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Settings.LocationPreference.FollowGaeApplication"),
			},
			{
				Name:        "settings_location_preference_kind",
				Description: "This is always *sql#locationPreference*",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Settings.LocationPreference.Kind"),
			},
			{
				Name:        "settings_location_preference_secondary_zone",
				Description: "The preferred Compute Engine zone for the secondary/failover (for example: us-central1-a, us-central1-b, etc) Reserved for future use",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Settings.LocationPreference.SecondaryZone"),
			},
			{
				Name:        "settings_location_preference_zone",
				Description: "The preferred Compute Engine zone (for example: us-central1-a, us-central1-b, etc)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Settings.LocationPreference.Zone"),
			},
			{
				Name:        "settings_maintenance_window_day",
				Description: "day of week (1-7), starting on Monday",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Settings.MaintenanceWindow.Day"),
			},
			{
				Name:        "settings_maintenance_window_hour",
				Description: "hour of day - 0 to 23",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Settings.MaintenanceWindow.Hour"),
			},
			{
				Name:        "settings_maintenance_window_kind",
				Description: "This is always *sql#maintenanceWindow*",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Settings.MaintenanceWindow.Kind"),
			},
			{
				Name:        "settings_maintenance_window_update_track",
				Description: "Maintenance timing setting: *canary* (Earlier) or *stable* (Later) Learn more  Possible values:   \"SQL_UPDATE_TRACK_UNSPECIFIED\" - This is an unknown maintenance timing preference   \"canary\" - For instance update that requires a restart, this update track indicates your instance prefer to restart for new version early in maintenance window   \"stable\" - For instance update that requires a restart, this update track indicates your instance prefer to let Cloud SQL choose the timing of restart (within its Maintenance window, if applicable)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Settings.MaintenanceWindow.UpdateTrack"),
			},
			{
				Name:        "settings_pricing_plan",
				Description: "The pricing plan for this instance This can be either *PER_USE* or *PACKAGE* Only *PER_USE* is supported for Second Generation instances  Possible values:   \"SQL_PRICING_PLAN_UNSPECIFIED\" - This is an unknown pricing plan for this instance   \"PACKAGE\" - The instance is billed at a monthly flat rate   \"PER_USE\" - The instance is billed per usage",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Settings.PricingPlan"),
			},
			{
				Name:        "settings_replication_type",
				Description: "The type of replication this instance uses This can be either *ASYNCHRONOUS* or *SYNCHRONOUS*",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Settings.ReplicationType"),
			},
			{
				Name:        "settings_version",
				Description: "The version of instance settings This is a required field for update method to make sure concurrent updates are handled properly During update, use the most recent settingsVersion value for this instance and do not try to update this value",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Settings.SettingsVersion"),
			},
			{
				Name:        "settings_storage_auto_resize",
				Description: "Configuration to increase storage size automatically The default value is true",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Settings.StorageAutoResize"),
			},
			{
				Name:        "settings_storage_auto_resize_limit",
				Description: "The maximum size to which storage capacity can be automatically increased The default value is 0, which specifies that there is no limit",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Settings.StorageAutoResizeLimit"),
			},
			{
				Name:        "settings_tier",
				Description: "The tier (or machine type) for this instance, for example *db-custom-1-3840*",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Settings.Tier"),
			},
			{
				Name:          "settings_user_labels",
				Description:   "User-provided labels, represented as a dictionary where each label is a single key value pair",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("Settings.UserLabels"),
				IgnoreInTests: true,
			},
			{
				Name:        "state",
				Description: "The current serving state of the Cloud SQL instance This can be one of the following *SQL_INSTANCE_STATE_UNSPECIFIED*: The state of the instance is unknown *RUNNABLE*: The instance is running, or has been stopped by owner *SUSPENDED*: The instance is not available, for example due to problems with billing *PENDING_DELETE*: The instance is being deleted *PENDING_CREATE*: The instance is being created *MAINTENANCE*: The instance is down for maintenance *FAILED*: The instance creation failed  Possible values:   \"SQL_INSTANCE_STATE_UNSPECIFIED\" - The state of the instance is unknown   \"RUNNABLE\" - The instance is running, or has been stopped by owner   \"SUSPENDED\" - The instance is not available, for example due to problems with billing   \"PENDING_DELETE\" - The instance is being deleted   \"PENDING_CREATE\" - The instance is being created   \"MAINTENANCE\" - The instance is down for maintenance   \"FAILED\" - The creation of the instance failed or a fatal error occurred during maintenance",
				Type:        schema.TypeString,
			},
			{
				Name:          "suspension_reason",
				Description:   "If the instance state is SUSPENDED, the reason for the suspension  Possible values:   \"SQL_SUSPENSION_REASON_UNSPECIFIED\" - This is an unknown suspension reason   \"BILLING_ISSUE\" - The instance is suspended due to billing issues (for example:, GCP account issue)   \"LEGAL_ISSUE\" - The instance is suspended due to illegal content (for example:, child pornography, copyrighted material, etc)   \"OPERATIONAL_ISSUE\" - The instance is causing operational issues (for example:, causing the database to crash)",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "gcp_sql_instance_ip_addresses",
				Description: "Database instance IP Mapping",
				Resolver:    fetchSqlInstanceIpAddresses,
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique ID of gcp_sql_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "instance_name",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("name"),
					},
					{
						Name:        "ip_address",
						Description: "The IP address assigned",
						Type:        schema.TypeString,
					},
					{
						Name:        "time_to_retire",
						Description: "The due time for this IP to be retired in RFC 3339 format, for example *2012-11-15T16:19:00094Z* This field is only available when the IP is scheduled to be retired",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The type of this IP address A *PRIMARY* address is a public address that can accept incoming connections A *PRIVATE* address is a private address that can accept incoming connections An *OUTGOING* address is the source address of connections originating from the instance, if supported  Possible values:   \"SQL_IP_ADDRESS_TYPE_UNSPECIFIED\" - This is an unknown IP address type   \"PRIMARY\" - IP address the customer is supposed to connect to Usually this is the load balancer's IP address   \"OUTGOING\" - Source IP address of the connection a read replica establishes to its external primary instance This IP address can be allowlisted by the customer in case it has a firewall that filters incoming connection to its on premises primary instance   \"PRIVATE\" - Private IP used when using private IPs and network peering   \"MIGRATED_1ST_GEN\" - V1 IP of a migrated instance We want the user to decommission this IP as soon as the migration is complete Note: V1 instances with V1 ip addresses will be counted as PRIMARY",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "gcp_sql_instance_settings_deny_maintenance_periods",
				Description:   "Deny Maintenance Periods This specifies a date range during when all CSA rollout will be denied",
				Resolver:      fetchSqlInstanceSettingsDenyMaintenancePeriods,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique ID of gcp_sql_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "instance_name",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("name"),
					},
					{
						Name:        "end_date",
						Description: "\"deny maintenance period\" end date If the year of the end date is empty, the year of the start date also must be empty In this case, it means the deny maintenance period recurs every year The date is in format yyyy-mm-dd ie, 2020-11-01, or mm-dd, ie",
						Type:        schema.TypeString,
					},
					{
						Name:        "start_date",
						Description: "\"deny maintenance period\" start date If the year of the start date is empty, the year of the end date also must be empty In this case, it means the deny maintenance period recurs every year The date is in format yyyy-mm-dd ie, 2020-11-01, or mm-dd, ie",
						Type:        schema.TypeString,
					},
					{
						Name:        "time",
						Description: "Time in UTC when the \"deny maintenance period\" starts on start_date and ends on end_date The time is in format: HH:mm:SS, ie",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "gcp_sql_instance_settings_ip_config_authorized_networks",
				Description:   "An entry for an Access Control list",
				Resolver:      fetchSqlInstanceSettingsIpConfigurationAuthorizedNetworks,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique ID of gcp_sql_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "instance_name",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("name"),
					},
					{
						Name:        "expiration_time",
						Description: "The time when this access control entry expires in RFC 3339 format, for example *2012-11-15T16:19:00094Z*",
						Type:        schema.TypeString,
					},
					{
						Name:        "kind",
						Description: "This is always *sql#aclEntry*",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "A label to identify this entry",
						Type:        schema.TypeString,
					},
					{
						Name:        "value",
						Description: "The allowlisted value for the access control list",
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
func fetchSqlInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Sql.Instances.
			List(c.ProjectId).
			PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		res <- output.Items
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func resolveSQLInstanceSettingsDatabaseFlags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	db := resource.Item.(*sql.DatabaseInstance)
	flags := make(map[string]string)
	for _, f := range db.Settings.DatabaseFlags {
		flags[f.Name] = f.Value
	}
	return errors.WithStack(resource.Set("settings_database_flags", flags))
}
func fetchSqlInstanceIpAddresses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	db := parent.Item.(*sql.DatabaseInstance)
	res <- db.IpAddresses
	return nil
}
func fetchSqlInstanceSettingsDenyMaintenancePeriods(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	db := parent.Item.(*sql.DatabaseInstance)
	if db.Settings != nil {
		res <- db.Settings.DenyMaintenancePeriods
	}
	return nil
}
func fetchSqlInstanceSettingsIpConfigurationAuthorizedNetworks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	db := parent.Item.(*sql.DatabaseInstance)
	if db.Settings != nil && db.Settings.IpConfiguration != nil {
		res <- db.Settings.IpConfiguration.AuthorizedNetworks
	}
	return nil
}
