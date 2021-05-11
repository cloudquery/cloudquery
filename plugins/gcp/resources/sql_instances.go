package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	sql "google.golang.org/api/sqladmin/v1beta4"
)

func SQLInstances() *schema.Table {
	return &schema.Table{
		Name:        "gcp_sql_instances",
		Resolver:    fetchSqlInstances,
		Multiplex:   client.ProjectMultiplex,
		IgnoreError: client.IgnoreErrorHandler,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name: "backend_type",
				Type: schema.TypeString,
			},
			{
				Name: "connection_name",
				Type: schema.TypeString,
			},
			{
				Name: "current_disk_size",
				Type: schema.TypeBigInt,
			},
			{
				Name: "database_version",
				Type: schema.TypeString,
			},
			{
				Name:     "disk_encryption_configuration_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskEncryptionConfiguration.Kind"),
			},
			{
				Name:     "disk_encryption_configuration_kms_key_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskEncryptionConfiguration.KmsKeyName"),
			},
			{
				Name:     "disk_encryption_status_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskEncryptionStatus.Kind"),
			},
			{
				Name:     "disk_encryption_status_kms_key_version_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskEncryptionStatus.KmsKeyVersionName"),
			},
			{
				Name: "etag",
				Type: schema.TypeString,
			},
			{
				Name:     "failover_replica_available",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("FailoverReplica.Available"),
			},
			{
				Name:     "failover_replica_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FailoverReplica.Name"),
			},
			{
				Name: "gce_zone",
				Type: schema.TypeString,
			},
			{
				Name: "instance_type",
				Type: schema.TypeString,
			},
			{
				Name: "ipv6_address",
				Type: schema.TypeString,
			},
			{
				Name: "kind",
				Type: schema.TypeString,
			},
			{
				Name: "master_instance_name",
				Type: schema.TypeString,
			},
			{
				Name: "max_disk_size",
				Type: schema.TypeBigInt,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name:     "on_premises_configuration_ca_certificate",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesConfiguration.CaCertificate"),
			},
			{
				Name:     "on_premises_configuration_client_certificate",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesConfiguration.ClientCertificate"),
			},
			{
				Name:     "on_premises_configuration_client_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesConfiguration.ClientKey"),
			},
			{
				Name:     "on_premises_configuration_dump_file_path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesConfiguration.DumpFilePath"),
			},
			{
				Name:     "on_premises_configuration_host_port",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesConfiguration.HostPort"),
			},
			{
				Name:     "on_premises_configuration_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesConfiguration.Kind"),
			},
			{
				Name:     "on_premises_configuration_password",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesConfiguration.Password"),
			},
			{
				Name:     "on_premises_configuration_username",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesConfiguration.Username"),
			},
			{
				Name: "project",
				Type: schema.TypeString,
			},
			{
				Name: "region",
				Type: schema.TypeString,
			},
			{
				Name:     "failover_target",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ReplicaConfiguration.FailoverTarget"),
			},
			{
				Name:     "configuration_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicaConfiguration.Kind"),
			},
			{
				Name:     "mysql_replica_configuration_ca_certificate",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicaConfiguration.MysqlReplicaConfiguration.CaCertificate"),
			},
			{
				Name:     "mysql_replica_configuration_client_certificate",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicaConfiguration.MysqlReplicaConfiguration.ClientCertificate"),
			},
			{
				Name:     "mysql_replica_configuration_client_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicaConfiguration.MysqlReplicaConfiguration.ClientKey"),
			},
			{
				Name:     "mysql_replica_configuration_connect_retry_interval",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ReplicaConfiguration.MysqlReplicaConfiguration.ConnectRetryInterval"),
			},
			{
				Name:     "mysql_replica_configuration_dump_file_path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicaConfiguration.MysqlReplicaConfiguration.DumpFilePath"),
			},
			{
				Name:     "mysql_replica_configuration_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicaConfiguration.MysqlReplicaConfiguration.Kind"),
			},
			{
				Name:     "mysql_replica_configuration_master_heartbeat_period",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ReplicaConfiguration.MysqlReplicaConfiguration.MasterHeartbeatPeriod"),
			},
			{
				Name:     "mysql_replica_configuration_password",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicaConfiguration.MysqlReplicaConfiguration.Password"),
			},
			{
				Name:     "mysql_replica_configuration_ssl_cipher",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicaConfiguration.MysqlReplicaConfiguration.SslCipher"),
			},
			{
				Name:     "mysql_replica_configuration_username",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicaConfiguration.MysqlReplicaConfiguration.Username"),
			},
			{
				Name:     "mysql_replica_configuration_verify_server_certificate",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ReplicaConfiguration.MysqlReplicaConfiguration.VerifyServerCertificate"),
			},
			{
				Name: "replica_names",
				Type: schema.TypeStringArray,
			},
			{
				Name: "root_password",
				Type: schema.TypeString,
			},
			{
				Name: "satisfies_pzs",
				Type: schema.TypeBool,
			},
			{
				Name:     "scheduled_maintenance_can_defer",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ScheduledMaintenance.CanDefer"),
			},
			{
				Name:     "scheduled_maintenance_can_reschedule",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ScheduledMaintenance.CanReschedule"),
			},
			{
				Name:     "scheduled_maintenance_start_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ScheduledMaintenance.StartTime"),
			},
			{
				Name: "secondary_gce_zone",
				Type: schema.TypeString,
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
			},
			{
				Name:     "server_ca_cert",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerCaCert.Cert"),
			},
			{
				Name:     "server_ca_cert_cert_serial_number",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerCaCert.CertSerialNumber"),
			},
			{
				Name:     "server_ca_cert_common_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerCaCert.CommonName"),
			},
			{
				Name:     "server_ca_cert_create_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerCaCert.CreateTime"),
			},
			{
				Name:     "server_ca_cert_expiration_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerCaCert.ExpirationTime"),
			},
			{
				Name:     "server_ca_cert_instance",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerCaCert.Instance"),
			},
			{
				Name:     "server_ca_cert_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerCaCert.Kind"),
			},
			{
				Name:     "server_ca_cert_self_link",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerCaCert.SelfLink"),
			},
			{
				Name:     "server_ca_cert_sha1_fingerprint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerCaCert.Sha1Fingerprint"),
			},
			{
				Name: "service_account_email_address",
				Type: schema.TypeString,
			},
			{
				Name:     "settings_activation_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Settings.ActivationPolicy"),
			},
			{
				Name:     "settings_active_directory_config_domain",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Settings.ActiveDirectoryConfig.Domain"),
			},
			{
				Name:     "settings_active_directory_config_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Settings.ActiveDirectoryConfig.Kind"),
			},
			{
				Name:     "settings_authorized_gae_applications",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Settings.AuthorizedGaeApplications"),
			},
			{
				Name:     "settings_availability_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Settings.AvailabilityType"),
			},
			{
				Name:     "settings_backup_retention_settings_retained_backups",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Settings.BackupConfiguration.BackupRetentionSettings.RetainedBackups"),
			},
			{
				Name:     "settings_backup_retention_settings_retention_unit",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Settings.BackupConfiguration.BackupRetentionSettings.RetentionUnit"),
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
				Name:     "settings_collation",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Settings.Collation"),
			},
			{
				Name:     "settings_crash_safe_replication_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Settings.CrashSafeReplicationEnabled"),
			},
			{
				Name:     "settings_data_disk_size_gb",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Settings.DataDiskSizeGb"),
			},
			{
				Name:     "settings_data_disk_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Settings.DataDiskType"),
			},
			{
				Name:     "settings_database_flags",
				Type:     schema.TypeJSON,
				Resolver: resolveSQLInstanceSettingsDatabaseFlags,
			},
			{
				Name:     "settings_database_replication_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Settings.DatabaseReplicationEnabled"),
			},
			{
				Name:     "settings_insights_config_query_insights_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Settings.InsightsConfig.QueryInsightsEnabled"),
			},
			{
				Name:     "settings_insights_config_query_string_length",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Settings.InsightsConfig.QueryStringLength"),
			},
			{
				Name:     "settings_insights_config_record_application_tags",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Settings.InsightsConfig.RecordApplicationTags"),
			},
			{
				Name:     "settings_insights_config_record_client_address",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Settings.InsightsConfig.RecordClientAddress"),
			},
			{
				Name:     "settings_ip_configuration_ipv4_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Settings.IpConfiguration.Ipv4Enabled"),
			},
			{
				Name:     "settings_ip_configuration_private_network",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Settings.IpConfiguration.PrivateNetwork"),
			},
			{
				Name:     "settings_ip_configuration_require_ssl",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Settings.IpConfiguration.RequireSsl"),
			},
			{
				Name:     "settings_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Settings.Kind"),
			},
			{
				Name:     "settings_location_preference_follow_gae_application",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Settings.LocationPreference.FollowGaeApplication"),
			},
			{
				Name:     "settings_location_preference_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Settings.LocationPreference.Kind"),
			},
			{
				Name:     "settings_location_preference_secondary_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Settings.LocationPreference.SecondaryZone"),
			},
			{
				Name:     "settings_location_preference_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Settings.LocationPreference.Zone"),
			},
			{
				Name:     "settings_maintenance_window_day",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Settings.MaintenanceWindow.Day"),
			},
			{
				Name:     "settings_maintenance_window_hour",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Settings.MaintenanceWindow.Hour"),
			},
			{
				Name:     "settings_maintenance_window_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Settings.MaintenanceWindow.Kind"),
			},
			{
				Name:     "settings_maintenance_window_update_track",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Settings.MaintenanceWindow.UpdateTrack"),
			},
			{
				Name:     "settings_pricing_plan",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Settings.PricingPlan"),
			},
			{
				Name:     "settings_replication_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Settings.ReplicationType"),
			},
			{
				Name:     "settings_version",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Settings.SettingsVersion"),
			},
			{
				Name:     "settings_storage_auto_resize",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Settings.StorageAutoResize"),
			},
			{
				Name:     "settings_storage_auto_resize_limit",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Settings.StorageAutoResizeLimit"),
			},
			{
				Name:     "settings_tier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Settings.Tier"),
			},
			{
				Name:     "settings_user_labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Settings.UserLabels"),
			},
			{
				Name: "state",
				Type: schema.TypeString,
			},
			{
				Name: "suspension_reason",
				Type: schema.TypeStringArray,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "gcp_sql_instance_ip_addresses",
				Resolver: fetchSqlInstanceIpAddresses,
				Columns: []schema.Column{
					{
						Name:     "instance_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "ip_address",
						Type: schema.TypeString,
					},
					{
						Name: "time_to_retire",
						Type: schema.TypeString,
					},
					{
						Name: "type",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "gcp_sql_instance_settings_deny_maintenance_periods",
				Resolver: fetchSqlInstanceSettingsDenyMaintenancePeriods,
				Columns: []schema.Column{
					{
						Name:     "instance_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "end_date",
						Type: schema.TypeString,
					},
					{
						Name: "start_date",
						Type: schema.TypeString,
					},
					{
						Name: "time",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "gcp_sql_instance_settings_ip_configuration_authorized_networks",
				Resolver: fetchSqlInstanceSettingsIpConfigurationAuthorizedNetworks,
				Columns: []schema.Column{
					{
						Name:     "instance_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "expiration_time",
						Type: schema.TypeString,
					},
					{
						Name: "kind",
						Type: schema.TypeString,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "value",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchSqlInstances(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Sql.Instances.List(c.ProjectId).Context(ctx)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}
		res <- output.Items
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func resolveSQLInstanceSettingsDatabaseFlags(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	db := resource.Item.(*sql.DatabaseInstance)
	flags := make(map[string]string)
	for _, f := range db.Settings.DatabaseFlags {
		flags[f.Name] = f.Value
	}
	return resource.Set("settings_database_flags", flags)
}

func fetchSqlInstanceIpAddresses(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	db := parent.Item.(*sql.DatabaseInstance)
	res <- db.IpAddresses
	return nil
}
func fetchSqlInstanceSettingsDenyMaintenancePeriods(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	db := parent.Item.(*sql.DatabaseInstance)
	if db.Settings != nil {
		res <- db.Settings.DenyMaintenancePeriods
	}
	return nil
}
func fetchSqlInstanceSettingsIpConfigurationAuthorizedNetworks(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	db := parent.Item.(*sql.DatabaseInstance)
	if db.Settings != nil && db.Settings.IpConfiguration != nil {
		res <- db.Settings.IpConfiguration.AuthorizedNetworks
	}
	return nil
}
