package sql

import (
	"google.golang.org/api/sql/v1beta4"
)

type DBInstance struct {
	ID              uint `gorm:"primarykey"`
	ProjectID       string
	BackendType     string
	ConnectionName  string
	CurrentDiskSize int64
	DatabaseVersion string

	DiskEncryptionConfigurationKind       string
	DiskEncryptionConfigurationKmsKeyName string

	DiskEncryptionStatusKind              string
	DiskEncryptionStatusKmsKeyVersionName string

	Etag string

	FailoverReplicaAvailable bool
	FailoverReplicaName      string

	GceZone            string
	InstanceType       string
	IpAddresses        []*DBInstanceIpMapping `gorm:"constraint:OnDelete:CASCADE;"`
	Ipv6Address        string
	Kind               string
	MasterInstanceName string
	MaxDiskSize        int64
	Name               string

	OnPremisesConfigurationCaCertificate     string
	OnPremisesConfigurationClientCertificate string
	OnPremisesConfigurationClientKey         string
	OnPremisesConfigurationDumpFilePath      string
	OnPremisesConfigurationHostPort          string
	OnPremisesConfigurationKind              string
	OnPremisesConfigurationPassword          string
	OnPremisesConfigurationUsername          string

	Project string
	Region  string

	ReplicaConfigurationFailoverTarget bool
	ReplicaConfigurationKind           string

	MysqlReplicaConfigurationCaCertificate           string
	MysqlReplicaConfigurationClientCertificate       string
	MysqlReplicaConfigurationClientKey               string
	MysqlReplicaConfigurationConnectRetryInterval    int64
	MysqlReplicaConfigurationDumpFilePath            string
	MysqlReplicaConfigurationKind                    string
	MysqlReplicaConfigurationMasterHeartbeatPeriod   int64
	MysqlReplicaConfigurationPassword                string
	MysqlReplicaConfigurationSslCipher               string
	MysqlReplicaConfigurationUsername                string
	MysqlReplicaConfigurationVerifyServerCertificate bool

	ReplicaNames []*DBInstanceReplicaName `gorm:"constraint:OnDelete:CASCADE;"`
	RootPassword string

	ScheduledMaintenanceCanDefer      bool
	ScheduledMaintenanceCanReschedule bool
	ScheduledMaintenanceStartTime     string

	SelfLink string

	//ServerCaCertCert string
	ServerCaCertSerialNumber    string
	ServerCaCertCommonName      string
	ServerCaCertCreateTime      string
	ServerCaCertExpirationTime  string
	ServerCaCertInstance        string
	ServerCaCertKind            string
	ServerCaCertSelfLink        string
	ServerCaCertSha1Fingerprint string

	ServiceAccountEmailAddress string

	SettingsActivationPolicy          string
	SettingsAuthorizedGaeApplications []*DBInstanceSettingsAuthorizedGaeApplication `gorm:"constraint:OnDelete:CASCADE;"`
	SettingsAvailabilityType          string

	SettingsBackupConfigurationBinaryLogEnabled               bool
	SettingsBackupConfigurationEnabled                        bool
	SettingsBackupConfigurationKind                           string
	SettingsBackupConfigurationLocation                       string
	SettingsBackupConfigurationPointInTimeRecoveryEnabled     bool
	SettingsBackupConfigurationReplicationLogArchivingEnabled bool
	SettingsBackupConfigurationStartTime                      string

	SettingsCrashSafeReplicationEnabled bool
	SettingsDataDiskSizeGb              int64
	SettingsDataDiskType                string
	SettingsDatabaseFlags               []*DBInstanceDatabaseFlag `gorm:"constraint:OnDelete:CASCADE;"`
	SettingsDatabaseReplicationEnabled  bool

	SettingsIpConfigAuthorizedNetworks []*DBInstanceAclEntry `gorm:"constraint:OnDelete:CASCADE;"`
	SettingsIpConfigIpv4Enabled        bool
	SettingsIpConfigPrivateNetwork     string
	SettingsIpConfigRequireSsl         bool

	SettingsKind string

	SettingsLocationPreferenceFollowGaeApplication string
	SettingsLocationPreferenceKind                 string
	SettingsLocationPreferenceZone                 string

	SettingsMaintenanceWindowDay         int64
	SettingsMaintenanceWindowHour        int64
	SettingsMaintenanceWindowKind        string
	SettingsMaintenanceWindowUpdateTrack string

	SettingsPricingPlan            string
	SettingsReplicationType        string
	SettingsSettingsVersion        int64
	SettingsStorageAutoResize      bool
	SettingsStorageAutoResizeLimit int64
	SettingsTier                   string

	State            string
	SuspensionReason []*DBInstanceSuspensionReason `gorm:"constraint:OnDelete:CASCADE;"`
}

func (DBInstance) TableName() string {
	return "gcp_sql_db_instances"
}

type DBInstanceIpMapping struct {
	ID           uint   `gorm:"primarykey"`
	ProjectID    string `gorm:"-"`
	DBInstanceID uint   `neo:"ignore"`
	IpAddress    string
	TimeToRetire string
	Type         string
}

func (DBInstanceIpMapping) TableName() string {
	return "gcp_sql_database_instance_ip_mappings"
}

type DBInstanceReplicaName struct {
	ID           uint `gorm:"primarykey"`
	DBInstanceID uint
	Value        string
}

func (DBInstanceReplicaName) TableName() string {
	return "gcp_sql_db_instance_replica_names"
}

type DBInstanceSettingsAuthorizedGaeApplication struct {
	ID           uint `gorm:"primarykey"`
	DBInstanceID uint
	Value        string
}

func (DBInstanceSettingsAuthorizedGaeApplication) TableName() string {
	return "gcp_sql_db_instance_settings_authorized_gae_applications"
}

type DBInstanceDatabaseFlag struct {
	ID           uint   `gorm:"primarykey"`
	ProjectID    string `gorm:"-"`
	DBInstanceID uint   `neo:"ignore"`
	Name         string
	Value        string
}

func (DBInstanceDatabaseFlag) TableName() string {
	return "gcp_sql_db_instance_database_flags"
}

type DBInstanceAclEntry struct {
	ID             uint   `gorm:"primarykey"`
	ProjectID      string `gorm:"-"`
	DBInstanceID   uint   `neo:"ignore"`
	ExpirationTime string
	Kind           string
	Name           string
	Value          string
}

func (DBInstanceAclEntry) TableName() string {
	return "gcp_sql_db_instance_acl_entries"
}

type DBInstanceSuspensionReason struct {
	ID           uint `gorm:"primarykey"`
	DBInstanceID uint
	Value        string
}

func (DBInstanceSuspensionReason) TableName() string {
	return "gcp_sql_db_instance_suspension_reasons"
}

func (c *Client) transformDatabaseInstances(values []*sql.DatabaseInstance) []*DBInstance {
	var tValues []*DBInstance
	for _, value := range values {
		tValue := DBInstance{
			ProjectID:                  c.projectID,
			BackendType:                value.BackendType,
			ConnectionName:             value.ConnectionName,
			CurrentDiskSize:            value.CurrentDiskSize,
			DatabaseVersion:            value.DatabaseVersion,
			Etag:                       value.Etag,
			GceZone:                    value.GceZone,
			InstanceType:               value.InstanceType,
			IpAddresses:                c.transformDatabaseInstanceIpMappings(value.IpAddresses),
			Ipv6Address:                value.Ipv6Address,
			Kind:                       value.Kind,
			MasterInstanceName:         value.MasterInstanceName,
			MaxDiskSize:                value.MaxDiskSize,
			Name:                       value.Name,
			Project:                    value.Project,
			Region:                     value.Region,
			ReplicaNames:               c.transformDatabaseInstanceReplicaNames(value.ReplicaNames),
			RootPassword:               value.RootPassword,
			SelfLink:                   value.SelfLink,
			ServiceAccountEmailAddress: value.ServiceAccountEmailAddress,
			State:                      value.State,
			SuspensionReason:           c.transformDatabaseInstanceSuspensionReason(value.SuspensionReason),
		}
		if value.DiskEncryptionConfiguration != nil {

			tValue.DiskEncryptionConfigurationKind = value.DiskEncryptionConfiguration.Kind
			tValue.DiskEncryptionConfigurationKmsKeyName = value.DiskEncryptionConfiguration.KmsKeyName

		}
		if value.DiskEncryptionStatus != nil {

			tValue.DiskEncryptionStatusKind = value.DiskEncryptionStatus.Kind
			tValue.DiskEncryptionStatusKmsKeyVersionName = value.DiskEncryptionStatus.KmsKeyVersionName

		}
		if value.FailoverReplica != nil {

			tValue.FailoverReplicaAvailable = value.FailoverReplica.Available
			tValue.FailoverReplicaName = value.FailoverReplica.Name

		}
		if value.OnPremisesConfiguration != nil {

			tValue.OnPremisesConfigurationCaCertificate = value.OnPremisesConfiguration.CaCertificate
			tValue.OnPremisesConfigurationClientCertificate = value.OnPremisesConfiguration.ClientCertificate
			tValue.OnPremisesConfigurationClientKey = value.OnPremisesConfiguration.ClientKey
			tValue.OnPremisesConfigurationDumpFilePath = value.OnPremisesConfiguration.DumpFilePath
			tValue.OnPremisesConfigurationHostPort = value.OnPremisesConfiguration.HostPort
			tValue.OnPremisesConfigurationKind = value.OnPremisesConfiguration.Kind
			tValue.OnPremisesConfigurationPassword = value.OnPremisesConfiguration.Password
			tValue.OnPremisesConfigurationUsername = value.OnPremisesConfiguration.Username

		}
		if value.ReplicaConfiguration != nil {

			tValue.ReplicaConfigurationFailoverTarget = value.ReplicaConfiguration.FailoverTarget
			tValue.ReplicaConfigurationKind = value.ReplicaConfiguration.Kind

		}
		if value.ScheduledMaintenance != nil {

			tValue.ScheduledMaintenanceCanDefer = value.ScheduledMaintenance.CanDefer
			tValue.ScheduledMaintenanceCanReschedule = value.ScheduledMaintenance.CanReschedule
			tValue.ScheduledMaintenanceStartTime = value.ScheduledMaintenance.StartTime

		}
		if value.ServerCaCert != nil {

			//tValue.ServerCaCertCert = value.ServerCaCert.Cert
			tValue.ServerCaCertSerialNumber = value.ServerCaCert.CertSerialNumber
			tValue.ServerCaCertCommonName = value.ServerCaCert.CommonName
			tValue.ServerCaCertCreateTime = value.ServerCaCert.CreateTime
			tValue.ServerCaCertExpirationTime = value.ServerCaCert.ExpirationTime
			tValue.ServerCaCertInstance = value.ServerCaCert.Instance
			tValue.ServerCaCertKind = value.ServerCaCert.Kind
			tValue.ServerCaCertSelfLink = value.ServerCaCert.SelfLink
			tValue.ServerCaCertSha1Fingerprint = value.ServerCaCert.Sha1Fingerprint

		}
		if value.Settings != nil {

			tValue.SettingsActivationPolicy = value.Settings.ActivationPolicy
			tValue.SettingsAuthorizedGaeApplications = c.transformDatabaseInstanceSettingsAuthorizedGaeApplications(value.Settings.AuthorizedGaeApplications)
			tValue.SettingsAvailabilityType = value.Settings.AvailabilityType
			tValue.SettingsCrashSafeReplicationEnabled = value.Settings.CrashSafeReplicationEnabled
			tValue.SettingsDataDiskSizeGb = value.Settings.DataDiskSizeGb
			tValue.SettingsDataDiskType = value.Settings.DataDiskType
			tValue.SettingsDatabaseFlags = c.transformDatabaseInstanceDatabaseFlags(value.Settings.DatabaseFlags)
			tValue.SettingsDatabaseReplicationEnabled = value.Settings.DatabaseReplicationEnabled
			tValue.SettingsKind = value.Settings.Kind
			tValue.SettingsPricingPlan = value.Settings.PricingPlan
			tValue.SettingsReplicationType = value.Settings.ReplicationType
			tValue.SettingsSettingsVersion = value.Settings.SettingsVersion
			tValue.SettingsStorageAutoResize = value.Settings.StorageAutoResize
			tValue.SettingsStorageAutoResizeLimit = value.Settings.StorageAutoResizeLimit
			tValue.SettingsTier = value.Settings.Tier

		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformDatabaseInstanceIpMappings(values []*sql.IpMapping) []*DBInstanceIpMapping {
	var tValues []*DBInstanceIpMapping
	for _, value := range values {
		tValue := DBInstanceIpMapping{
			ProjectID:    c.projectID,
			IpAddress:    value.IpAddress,
			TimeToRetire: value.TimeToRetire,
			Type:         value.Type,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}
func (c *Client) transformDatabaseInstanceReplicaNames(values []string) []*DBInstanceReplicaName {
	var tValues []*DBInstanceReplicaName
	for _, v := range values {
		tValues = append(tValues, &DBInstanceReplicaName{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformDatabaseInstanceSettingsAuthorizedGaeApplications(values []string) []*DBInstanceSettingsAuthorizedGaeApplication {
	var tValues []*DBInstanceSettingsAuthorizedGaeApplication
	for _, v := range values {
		tValues = append(tValues, &DBInstanceSettingsAuthorizedGaeApplication{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformDatabaseInstanceDatabaseFlags(values []*sql.DatabaseFlags) []*DBInstanceDatabaseFlag {
	var tValues []*DBInstanceDatabaseFlag
	for _, value := range values {
		tValue := DBInstanceDatabaseFlag{
			ProjectID: c.projectID,
			Name:      value.Name,
			Value:     value.Value,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformDatabaseInstanceAclEntries(values []*sql.AclEntry) []*DBInstanceAclEntry {
	var tValues []*DBInstanceAclEntry
	for _, value := range values {
		tValue := DBInstanceAclEntry{
			ProjectID:      c.projectID,
			ExpirationTime: value.ExpirationTime,
			Kind:           value.Kind,
			Name:           value.Name,
			Value:          value.Value,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformDatabaseInstanceSuspensionReason(values []string) []*DBInstanceSuspensionReason {
	var tValues []*DBInstanceSuspensionReason
	for _, v := range values {
		tValues = append(tValues, &DBInstanceSuspensionReason{
			Value: v,
		})
	}
	return tValues
}

var DatabaseInstanceTables = []interface{}{
	&DBInstance{},
	&DBInstanceIpMapping{},
	&DBInstanceDatabaseFlag{},
	&DBInstanceAclEntry{},
}

func (c *Client) instances(_ interface{}) error {

	nextPageToken := ""
	c.db.Where("project_id", c.projectID).Delete(DatabaseInstanceTables...)
	for {
		call := c.svc.Instances.List(c.projectID)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		c.db.ChunkedCreate(c.transformDatabaseInstances(output.Items))
		c.log.Info("Fetched resources", "resource", "sql.instances", "count", len(output.Items))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
