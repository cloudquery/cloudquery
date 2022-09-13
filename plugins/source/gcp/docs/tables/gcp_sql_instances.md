
# Table: gcp_sql_instances
A Cloud SQL instance resource
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|backend_type|text|*SECOND_GEN*: Cloud SQL database instance *EXTERNAL*: A database server that is not managed by Google This property is read-only; use the *tier* property in the *settings* object to determine the database type  Possible values:   "SQL_BACKEND_TYPE_UNSPECIFIED" - This is an unknown backend type for instance   "FIRST_GEN" - V1 speckle instance   "SECOND_GEN" - V2 speckle instance   "EXTERNAL" - On premises instance|
|connection_name|text|Connection name of the Cloud SQL instance used in connection strings|
|current_disk_size|bigint|The current disk usage of the instance in bytes This property has been deprecated by google API and might be null|
|database_version|text|The database engine type and version The *databaseVersion* field cannot be changed after instance creation MySQL instances: *MYSQL_8_0*, *MYSQL_5_7* (default), or *MYSQL_5_6* PostgreSQL instances: *POSTGRES_9_6*, *POSTGRES_10*, *POSTGRES_11*, *POSTGRES_12*, or *POSTGRES_13* (default) SQL Server instances: *SQLSERVER_2017_STANDARD* (default), *SQLSERVER_2017_ENTERPRISE*, *SQLSERVER_2017_EXPRESS*, or *SQLSERVER_2017_WEB*  Possible values:   "SQL_DATABASE_VERSION_UNSPECIFIED" - This is an unknown database version   "MYSQL_5_1" - The database version is MySQL 51   "MYSQL_5_5" - The database version is MySQL 55   "MYSQL_5_6" - The database version is MySQL 56   "MYSQL_5_7" - The database version is MySQL 57   "POSTGRES_9_6" - The database version is PostgreSQL 96   "POSTGRES_11" - The database version is PostgreSQL 11   "SQLSERVER_2017_STANDARD" - The database version is SQL Server 2017 Standard   "SQLSERVER_2017_ENTERPRISE" - The database version is SQL Server 2017 Enterprise   "SQLSERVER_2017_EXPRESS" - The database version is SQL Server 2017 Express   "SQLSERVER_2017_WEB" - The database version is SQL Server 2017 Web   "POSTGRES_10" - The database version is PostgreSQL 10   "POSTGRES_12" - The database version is PostgreSQL 12   "MYSQL_8_0" - The database version is MySQL 8   "POSTGRES_13" - The database version is PostgreSQL 13|
|disk_encryption_configuration_kind|text|This is always *sql#diskEncryptionConfiguration*|
|disk_encryption_configuration_kms_key_name|text||
|disk_encryption_status_kind|text|This is always *sql#diskEncryptionStatus*|
|disk_encryption_status_kms_key_version_name|text||
|etag|text|This field is deprecated and will be removed from a future version of the API Use the *settingssettingsVersion* field instead|
|failover_replica_available|boolean|The availability status of the failover replica A false status indicates that the failover replica is out of sync The primary instance can only failover to the failover replica when the status is true|
|failover_replica_name|text|The name of the failover replica If specified at instance creation, a failover replica is created for the instance The name doesn't include the project ID This property is applicable only to Second Generation instances|
|gce_zone|text|The Compute Engine zone that the instance is currently serving from This value could be different from the zone that was specified when the instance was created if the instance has failed over to its secondary zone|
|instance_type|text|The instance type This can be one of the following *CLOUD_SQL_INSTANCE*: A Cloud SQL instance that is not replicating from a primary instance *ON_PREMISES_INSTANCE*: An instance running on the customer's premises *READ_REPLICA_INSTANCE*: A Cloud SQL instance configured as a read-replica  Possible values:   "SQL_INSTANCE_TYPE_UNSPECIFIED" - This is an unknown Cloud SQL instance type   "CLOUD_SQL_INSTANCE" - A regular Cloud SQL instance   "ON_PREMISES_INSTANCE" - An instance running on the customer's premises that is not managed by Cloud SQL   "READ_REPLICA_INSTANCE" - A Cloud SQL instance acting as a read-replica|
|ipv6_address|text|The IPv6 address assigned to the instance (Deprecated) This property was applicable only to First Generation instances|
|kind|text|This is always *sql#instance*|
|master_instance_name|text|The name of the instance which will act as primary in the replication setup|
|max_disk_size|bigint|The maximum disk size of the instance in bytes|
|name|text|Name of the Cloud SQL instance This does not include the project ID|
|id|text|Alias of a name.|
|on_premises_configuration_ca_certificate|text|PEM representation of the trusted CA's x509 certificate|
|on_premises_configuration_client_certificate|text|PEM representation of the replica's x509 certificate|
|on_premises_configuration_client_key|text||
|on_premises_configuration_dump_file_path|text|The dump file to create the Cloud SQL replica|
|on_premises_configuration_host_port|text||
|on_premises_configuration_kind|text|This is always *sql#onPremisesConfiguration*|
|on_premises_configuration_password|text|The password for connecting to on-premises instance|
|on_premises_configuration_username|text|The username for connecting to on-premises instance|
|project|text|The project ID of the project containing the Cloud SQL instance The Google apps domain is prefixed if applicable|
|region|text|The geographical region Can be *us-central* (*FIRST_GEN* instances only) *us-central1* (*SECOND_GEN* instances only) *asia-east1* or *europe-west1* Defaults to *us-central* or *us-central1* depending on the instance type The region cannot be changed after instance creation|
|failover_target|boolean|Specifies if the replica is the failover target If the field is set to *true* the replica will be designated as a failover replica In case the primary instance fails, the replica instance will be promoted as the new primary instance Only one replica can be specified as failover target, and the replica has to be in different zone with the primary instance|
|configuration_kind|text||
|mysql_replica_configuration_ca_certificate|text|PEM representation of the trusted CA's x509 certificate|
|mysql_replica_configuration_client_certificate|text|PEM representation of the replica's x509 certificate|
|mysql_replica_configuration_client_key|text||
|mysql_replica_configuration_connect_retry_interval|bigint|Seconds to wait between connect retries MySQL's default is 60 seconds|
|mysql_replica_configuration_dump_file_path|text|Path to a SQL dump file in Google Cloud Storage from which the replica instance is to be created The URI is in the form gs://bucketName/fileName Compressed gzip files (gz) are also supported Dumps have the binlog co-ordinates from which replication begins This can be accomplished by setting --master-data to 1 when using mysqldump|
|mysql_replica_configuration_kind|text|This is always *sql#mysqlReplicaConfiguration*|
|mysql_replica_configuration_master_heartbeat_period|bigint|Interval in milliseconds between replication heartbeats|
|mysql_replica_configuration_password|text|The password for the replication connection|
|mysql_replica_configuration_ssl_cipher|text|A list of permissible ciphers to use for SSL encryption|
|mysql_replica_configuration_username|text|The username for the replication connection|
|mysql_replica_configuration_verify_server_certificate|boolean|Whether or not to check the primary instance's Common Name value in the certificate that it sends during the SSL handshake|
|replica_names|text[]|The replicas of the instance|
|root_password|text|Initial root password Use only on creation|
|satisfies_pzs|boolean|The status indicating if instance satisfiesPzs Reserved for future use|
|scheduled_maintenance_can_defer|boolean||
|scheduled_maintenance_can_reschedule|boolean|If the scheduled maintenance can be rescheduled|
|scheduled_maintenance_start_time|text|The start time of any upcoming scheduled maintenance for this instance|
|secondary_gce_zone|text|The Compute Engine zone that the failover instance is currently serving from for a regional instance This value could be different from the zone that was specified when the instance was created if the instance has failed over to its secondary/failover zone Reserved for future use|
|self_link|text|The URI of this resource|
|server_ca_cert|text|PEM representation|
|server_ca_cert_cert_serial_number|text|Serial number, as extracted from the certificate|
|server_ca_cert_common_name|text|User supplied name Constrained to [a-zA-Z-_ ]+|
|server_ca_cert_create_time|text|The time when the certificate was created in RFC 3339 format, for example *2012-11-15T16:19:00|
|server_ca_cert_expiration_time|text|The time when the certificate expires in RFC 3339 format, for example *2012-11-15T16:19:00094Z*|
|server_ca_cert_instance|text|Name of the database instance|
|server_ca_cert_kind|text|This is always *sql#sslCert*|
|server_ca_cert_self_link|text|The URI of this resource|
|server_ca_cert_sha1_fingerprint|text|Sha1 Fingerprint|
|service_account_email_address|text|The service account email address assigned to the instance This property is applicable only to Second Generation instances|
|settings_activation_policy|text|The activation policy specifies when the instance is activated; it is applicable only when the instance state is RUNNABLE Valid values: *ALWAYS*: The instance is on, and remains so even in the absence of connection requests *NEVER*: The instance is off; it is not activated, even if a connection request arrives  Possible values:   "SQL_ACTIVATION_POLICY_UNSPECIFIED" - Unknown activation plan   "ALWAYS" - The instance is always up and running   "NEVER" - The instance never starts   "ON_DEMAND" - The instance starts upon receiving requests|
|settings_active_directory_config_domain|text|The name of the domain (eg, mydomaincom)|
|settings_active_directory_config_kind|text|This is always sql#activeDirectoryConfig|
|settings_authorized_gae_applications|text[]|The App Engine app IDs that can access this instance (Deprecated) Applied to First Generation instances only|
|settings_availability_type|text|Availability type Potential values: *ZONAL*: The instance serves data from only one zone Outages in that zone affect data accessibility *REGIONAL*: The instance can serve data from more than one zone in a region (it is highly available) For more information, see Overview of the High Availability Configuration  Possible values:   "SQL_AVAILABILITY_TYPE_UNSPECIFIED" - This is an unknown Availability type   "ZONAL" - Zonal available instance   "REGIONAL" - Regional available instance|
|settings_backup_retention_settings_retained_backups|bigint|Depending on the value of retention_unit, this is used to determine if a backup needs to be deleted If retention_unit is 'COUNT', we will retain this many backups|
|settings_backup_retention_settings_retention_unit|text|The unit that 'retained_backups' represents  Possible values:   "RETENTION_UNIT_UNSPECIFIED" - Backup retention unit is unspecified, will be treated as COUNT   "COUNT" - Retention will be by count, eg "retain the most recent 7 backups"|
|settings_backup_binary_log_enabled|boolean||
|settings_backup_enabled|boolean||
|settings_backup_kind|text||
|settings_backup_location|text||
|settings_backup_point_in_time_recovery_enabled|boolean||
|settings_backup_replication_log_archiving_enabled|boolean||
|settings_backup_start_time|text||
|settings_backup_transaction_log_retention_days|bigint||
|settings_collation|text|The name of server Instance collation|
|settings_crash_safe_replication_enabled|boolean|Configuration specific to read replica instances Indicates whether database flags for crash-safe replication are enabled This property was only applicable to First Generation instances|
|settings_data_disk_size_gb|bigint|The size of data disk, in GB The data disk size minimum is 10GB|
|settings_data_disk_type|text|The type of data disk: PD_SSD (default) or PD_HDD Not used for First Generation instances  Possible values:   "SQL_DATA_DISK_TYPE_UNSPECIFIED" - This is an unknown data disk type   "PD_SSD" - An SSD data disk   "PD_HDD" - An HDD data disk   "OBSOLETE_LOCAL_SSD" - This field is deprecated and will be removed from a future version of the API|
|settings_database_flags|jsonb|The database flags passed to the instance at startup|
|settings_database_replication_enabled|boolean|Configuration specific to read replica instances Indicates whether replication is enabled or not|
|settings_insights_config_query_insights_enabled|boolean|Whether Query Insights feature is enabled|
|settings_insights_config_query_string_length|bigint|Maximum query length stored in bytes Default value: 1024 bytes Range: 256-4500 bytes Query length more than this field value will be truncated to this value When unset, query length will be the default value Changing query length will restart the database|
|settings_insights_config_record_application_tags|boolean|Whether Query Insights will record application tags from query when enabled|
|settings_insights_config_record_client_address|boolean|Whether Query Insights will record client address when enabled|
|settings_ip_configuration_ipv4_enabled|boolean|Whether the instance is assigned a public IP address or not|
|settings_ip_configuration_private_network|text|The resource link for the VPC network from which the Cloud SQL instance is accessible for private IP For example, */projects/myProject/global/networks/default* This setting can be updated, but it cannot be removed after it is set|
|settings_ip_configuration_require_ssl|boolean|Whether SSL connections over IP are enforced or not|
|settings_kind|text|This is always *sql#settings*|
|settings_location_preference_follow_gae_application|text|The App Engine application to follow, it must be in the same region as the Cloud SQL instance|
|settings_location_preference_kind|text|This is always *sql#locationPreference*|
|settings_location_preference_secondary_zone|text|The preferred Compute Engine zone for the secondary/failover (for example: us-central1-a, us-central1-b, etc) Reserved for future use|
|settings_location_preference_zone|text|The preferred Compute Engine zone (for example: us-central1-a, us-central1-b, etc)|
|settings_maintenance_window_day|bigint|day of week (1-7), starting on Monday|
|settings_maintenance_window_hour|bigint|hour of day - 0 to 23|
|settings_maintenance_window_kind|text|This is always *sql#maintenanceWindow*|
|settings_maintenance_window_update_track|text|Maintenance timing setting: *canary* (Earlier) or *stable* (Later) Learn more  Possible values:   "SQL_UPDATE_TRACK_UNSPECIFIED" - This is an unknown maintenance timing preference   "canary" - For instance update that requires a restart, this update track indicates your instance prefer to restart for new version early in maintenance window   "stable" - For instance update that requires a restart, this update track indicates your instance prefer to let Cloud SQL choose the timing of restart (within its Maintenance window, if applicable)|
|settings_pricing_plan|text|The pricing plan for this instance This can be either *PER_USE* or *PACKAGE* Only *PER_USE* is supported for Second Generation instances  Possible values:   "SQL_PRICING_PLAN_UNSPECIFIED" - This is an unknown pricing plan for this instance   "PACKAGE" - The instance is billed at a monthly flat rate   "PER_USE" - The instance is billed per usage|
|settings_replication_type|text|The type of replication this instance uses This can be either *ASYNCHRONOUS* or *SYNCHRONOUS*|
|settings_version|bigint|The version of instance settings This is a required field for update method to make sure concurrent updates are handled properly During update, use the most recent settingsVersion value for this instance and do not try to update this value|
|settings_storage_auto_resize|boolean|Configuration to increase storage size automatically The default value is true|
|settings_storage_auto_resize_limit|bigint|The maximum size to which storage capacity can be automatically increased The default value is 0, which specifies that there is no limit|
|settings_tier|text|The tier (or machine type) for this instance, for example *db-custom-1-3840*|
|settings_user_labels|jsonb|User-provided labels, represented as a dictionary where each label is a single key value pair|
|state|text|The current serving state of the Cloud SQL instance This can be one of the following *SQL_INSTANCE_STATE_UNSPECIFIED*: The state of the instance is unknown *RUNNABLE*: The instance is running, or has been stopped by owner *SUSPENDED*: The instance is not available, for example due to problems with billing *PENDING_DELETE*: The instance is being deleted *PENDING_CREATE*: The instance is being created *MAINTENANCE*: The instance is down for maintenance *FAILED*: The instance creation failed  Possible values:   "SQL_INSTANCE_STATE_UNSPECIFIED" - The state of the instance is unknown   "RUNNABLE" - The instance is running, or has been stopped by owner   "SUSPENDED" - The instance is not available, for example due to problems with billing   "PENDING_DELETE" - The instance is being deleted   "PENDING_CREATE" - The instance is being created   "MAINTENANCE" - The instance is down for maintenance   "FAILED" - The creation of the instance failed or a fatal error occurred during maintenance|
|suspension_reason|text[]|If the instance state is SUSPENDED, the reason for the suspension  Possible values:   "SQL_SUSPENSION_REASON_UNSPECIFIED" - This is an unknown suspension reason   "BILLING_ISSUE" - The instance is suspended due to billing issues (for example:, GCP account issue)   "LEGAL_ISSUE" - The instance is suspended due to illegal content (for example:, child pornography, copyrighted material, etc)   "OPERATIONAL_ISSUE" - The instance is causing operational issues (for example:, causing the database to crash)|
