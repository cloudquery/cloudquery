package storage

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-06-01/storage"
	"github.com/Azure/go-autorest/autorest"
	"github.com/cloudquery/cloudquery/database"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type Account struct {
	ID uint `gorm:"primarykey"`
	//AccountID string

	SkuName string
	SkuTier string
	Kind    string

	IdentityPrincipalID *string
	IdentityTenantID    *string
	IdentityType        *string

	AccountPropertiesProvisioningState string

	AccountPropertiesPrimaryEndpointsBlob  *string
	AccountPropertiesPrimaryEndpointsQueue *string
	AccountPropertiesPrimaryEndpointsTable *string
	AccountPropertiesPrimaryEndpointsFile  *string
	AccountPropertiesPrimaryEndpointsWeb   *string
	AccountPropertiesPrimaryEndpointsDfs   *string

	AccountPropertiesPrimaryEndpointsMicrosoftEndpointsBlob  *string
	AccountPropertiesPrimaryEndpointsMicrosoftEndpointsQueue *string
	AccountPropertiesPrimaryEndpointsMicrosoftEndpointsTable *string
	AccountPropertiesPrimaryEndpointsMicrosoftEndpointsFile  *string
	AccountPropertiesPrimaryEndpointsMicrosoftEndpointsWeb   *string
	AccountPropertiesPrimaryEndpointsMicrosoftEndpointsDfs   *string

	AccountPropertiesPrimaryEndpointsInternetEndpointsBlob *string
	AccountPropertiesPrimaryEndpointsInternetEndpointsFile *string
	AccountPropertiesPrimaryEndpointsInternetEndpointsWeb  *string
	AccountPropertiesPrimaryEndpointsInternetEndpointsDfs  *string
	AccountPropertiesPrimaryLocation                       *string
	AccountPropertiesStatusOfPrimary                       string
	AccountPropertiesLastGeoFailoverTime                   time.Time
	AccountPropertiesSecondaryLocation                     *string
	AccountPropertiesStatusOfSecondary                     string
	AccountPropertiesCreationTime                          time.Time

	AccountPropertiesCustomDomainName             *string
	AccountPropertiesCustomDomainUseSubDomainName *bool

	AccountPropertiesSecondaryEndpointsBlob  *string
	AccountPropertiesSecondaryEndpointsQueue *string
	AccountPropertiesSecondaryEndpointsTable *string
	AccountPropertiesSecondaryEndpointsFile  *string
	AccountPropertiesSecondaryEndpointsWeb   *string
	AccountPropertiesSecondaryEndpointsDfs   *string

	AccountPropertiesSecondaryEndpointsMicrosoftEndpointsBlob  *string
	AccountPropertiesSecondaryEndpointsMicrosoftEndpointsQueue *string
	AccountPropertiesSecondaryEndpointsMicrosoftEndpointsTable *string
	AccountPropertiesSecondaryEndpointsMicrosoftEndpointsFile  *string
	AccountPropertiesSecondaryEndpointsMicrosoftEndpointsWeb   *string
	AccountPropertiesSecondaryEndpointsMicrosoftEndpointsDfs   *string

	AccountPropertiesSecondaryEndpointsInternetEndpointsBlob *string
	AccountPropertiesSecondaryEndpointsInternetEndpointsFile *string
	AccountPropertiesSecondaryEndpointsInternetEndpointsWeb  *string
	AccountPropertiesSecondaryEndpointsInternetEndpointsDfs  *string

	AccountPropertiesEncryptionServicesBlobEnabled         *bool
	AccountPropertiesEncryptionServicesBlobLastEnabledTime time.Time
	AccountPropertiesEncryptionServicesBlobKeyType         string

	AccountPropertiesEncryptionServicesFileEnabled         *bool
	AccountPropertiesEncryptionServicesFileLastEnabledTime time.Time
	AccountPropertiesEncryptionServicesFileKeyType         string

	AccountPropertiesEncryptionServicesTableEnabled         *bool
	AccountPropertiesEncryptionServicesTableLastEnabledTime time.Time
	AccountPropertiesEncryptionServicesTableKeyType         string

	AccountPropertiesEncryptionServicesQueueEnabled            *bool
	AccountPropertiesEncryptionServicesQueueLastEnabledTime    time.Time
	AccountPropertiesEncryptionServicesQueueKeyType            string
	AccountPropertiesEncryptionKeySource                       string
	AccountPropertiesEncryptionRequireInfrastructureEncryption *bool

	AccountPropertiesEncryptionKeyVaultPropertiesKeyName                       *string
	AccountPropertiesEncryptionKeyVaultPropertiesKeyVersion                    *string
	AccountPropertiesEncryptionKeyVaultPropertiesKeyVaultURI                   *string
	AccountPropertiesEncryptionKeyVaultPropertiesCurrentVersionedKeyIdentifier *string
	AccountPropertiesEncryptionKeyVaultPropertiesLastKeyRotationTimestamp      time.Time
	AccountPropertiesAccessTier                                                string

	AccountPropertiesAzureFilesIdentityBasedAuthenticationDirectoryServiceOptions string

	AccountPropertiesAzureFilesIdentityBasedAuthenticationActiveDirectoryPropertiesDomainName        *string
	AccountPropertiesAzureFilesIdentityBasedAuthenticationActiveDirectoryPropertiesNetBiosDomainName *string
	AccountPropertiesAzureFilesIdentityBasedAuthenticationActiveDirectoryPropertiesForestName        *string
	AccountPropertiesAzureFilesIdentityBasedAuthenticationActiveDirectoryPropertiesDomainGUID        *string
	AccountPropertiesAzureFilesIdentityBasedAuthenticationActiveDirectoryPropertiesDomainSid         *string
	AccountPropertiesAzureFilesIdentityBasedAuthenticationActiveDirectoryPropertiesAzureStorageSid   *string
	AccountPropertiesEnableHTTPSTrafficOnly                                                          *bool

	AccountPropertiesNetworkRuleSetBypass              string
	AccountPropertiesNetworkRuleSetVirtualNetworkRules []AccountVirtualNetworkRule `gorm:"constraint:OnDelete:CASCADE;"`
	AccountPropertiesNetworkRuleSetIPRules             []AccountIPRule             `gorm:"constraint:OnDelete:CASCADE;"`
	AccountPropertiesNetworkRuleSetDefaultAction       string
	AccountPropertiesIsHnsEnabled                      *bool

	AccountPropertiesGeoReplicationStatsStatus       string
	AccountPropertiesGeoReplicationStatsLastSyncTime time.Time
	AccountPropertiesGeoReplicationStatsCanFailover  *bool
	AccountPropertiesFailoverInProgress              *bool
	AccountPropertiesLargeFileSharesState            string
	AccountPropertiesPrivateEndpointConnections      []AccountPrivateEndpointConnection `gorm:"constraint:OnDelete:CASCADE;"`

	AccountPropertiesRoutingPreferenceRoutingChoice             string
	AccountPropertiesRoutingPreferencePublishMicrosoftEndpoints *bool
	AccountPropertiesRoutingPreferencePublishInternetEndpoints  *bool

	AccountPropertiesBlobRestoreStatusStatus        string
	AccountPropertiesBlobRestoreStatusFailureReason *string
	AccountPropertiesBlobRestoreStatusRestoreID     *string

	AccountPropertiesBlobRestoreStatusParametersTimeToRestore time.Time
	AccountPropertiesBlobRestoreStatusParametersBlobRanges    []AccountBlobRestoreRange `gorm:"constraint:OnDelete:CASCADE;"`
	AccountPropertiesAllowBlobPublicAccess                    *bool
	AccountPropertiesMinimumTLSVersion                        string
	Location                                                  *string
	ResourceID                                                *string
	Name                                                      *string
	Type                                                      *string
}

func (Account) TableName() string {
	return "azure_storage_accounts"
}

type AccountVirtualNetworkRule struct {
	ID                       uint `gorm:"primarykey"`
	AccountID                uint
	AccountNetworkRuleSetID  uint
	VirtualNetworkResourceID *string
	Action                   string
	State                    string
}

func (AccountVirtualNetworkRule) TableName() string {
	return "azure_storage_accountvirtualnetworkrules"
}

type AccountIPRule struct {
	ID                      uint `gorm:"primarykey"`
	AccountID               uint
	AccountNetworkRuleSetID uint
	IPAddressOrRange        *string
	Action                  string
}

func (AccountIPRule) TableName() string {
	return "azure_storage_accountiprules"
}

type AccountPrivateEndpointConnection struct {
	ID                  uint `gorm:"primarykey"`
	AccountID           uint
	AccountPropertiesID uint

	PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateStatus         string
	PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateDescription    *string
	PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateActionRequired *string
	PrivateEndpointConnectionPropertiesProvisioningState                               string
	ResourceID                                                                         *string
	Name                                                                               *string
	Type                                                                               *string
}

func (AccountPrivateEndpointConnection) TableName() string {
	return "azure_storage_accountprivateendpointconnections"
}

type AccountBlobRestoreRange struct {
	ID                             uint `gorm:"primarykey"`
	AccountID                      uint
	AccountBlobRestoreParametersID uint
	StartRange                     *string
	EndRange                       *string
}

func (AccountBlobRestoreRange) TableName() string {
	return "azure_storage_accountblobrestoreranges"
}

func transformAccount(value storage.Account) Account {
	return Account{

		SkuName: string(value.Sku.Name),
		SkuTier: string(value.Sku.Tier),
		Kind:    string(value.Kind),

		//IdentityPrincipalID: value.Identity.PrincipalID,
		//IdentityTenantID: value.Identity.TenantID,
		//IdentityType: value.Identity.Type,

		AccountPropertiesProvisioningState: string(value.AccountProperties.ProvisioningState),

		AccountPropertiesPrimaryEndpointsBlob:  value.AccountProperties.PrimaryEndpoints.Blob,
		AccountPropertiesPrimaryEndpointsQueue: value.AccountProperties.PrimaryEndpoints.Queue,
		AccountPropertiesPrimaryEndpointsTable: value.AccountProperties.PrimaryEndpoints.Table,
		AccountPropertiesPrimaryEndpointsFile:  value.AccountProperties.PrimaryEndpoints.File,
		AccountPropertiesPrimaryEndpointsWeb:   value.AccountProperties.PrimaryEndpoints.Web,
		AccountPropertiesPrimaryEndpointsDfs:   value.AccountProperties.PrimaryEndpoints.Dfs,

		//AccountPropertiesPrimaryEndpointsMicrosoftEndpointsBlob: value.AccountProperties.PrimaryEndpoints.MicrosoftEndpoints.Blob,
		//AccountPropertiesPrimaryEndpointsMicrosoftEndpointsQueue: value.AccountProperties.PrimaryEndpoints.MicrosoftEndpoints.Queue,
		//AccountPropertiesPrimaryEndpointsMicrosoftEndpointsTable: value.AccountProperties.PrimaryEndpoints.MicrosoftEndpoints.Table,
		//AccountPropertiesPrimaryEndpointsMicrosoftEndpointsFile: value.AccountProperties.PrimaryEndpoints.MicrosoftEndpoints.File,
		//AccountPropertiesPrimaryEndpointsMicrosoftEndpointsWeb: value.AccountProperties.PrimaryEndpoints.MicrosoftEndpoints.Web,
		//AccountPropertiesPrimaryEndpointsMicrosoftEndpointsDfs: value.AccountProperties.PrimaryEndpoints.MicrosoftEndpoints.Dfs,

		//AccountPropertiesPrimaryEndpointsInternetEndpointsBlob: value.AccountProperties.PrimaryEndpoints.InternetEndpoints.Blob,
		//AccountPropertiesPrimaryEndpointsInternetEndpointsFile: value.AccountProperties.PrimaryEndpoints.InternetEndpoints.File,
		//AccountPropertiesPrimaryEndpointsInternetEndpointsWeb: value.AccountProperties.PrimaryEndpoints.InternetEndpoints.Web,
		//AccountPropertiesPrimaryEndpointsInternetEndpointsDfs: value.AccountProperties.PrimaryEndpoints.InternetEndpoints.Dfs,
		AccountPropertiesPrimaryLocation: value.AccountProperties.PrimaryLocation,
		AccountPropertiesStatusOfPrimary: string(value.AccountProperties.StatusOfPrimary),
		//AccountPropertiesLastGeoFailoverTime: value.AccountProperties.LastGeoFailoverTime.ToTime(),
		AccountPropertiesSecondaryLocation: value.AccountProperties.SecondaryLocation,
		AccountPropertiesStatusOfSecondary: string(value.AccountProperties.StatusOfSecondary),
		AccountPropertiesCreationTime:      value.AccountProperties.CreationTime.ToTime(),

		//AccountPropertiesCustomDomainName: value.AccountProperties.CustomDomain.Name,
		//AccountPropertiesCustomDomainUseSubDomainName: value.AccountProperties.CustomDomain.UseSubDomainName,

		AccountPropertiesSecondaryEndpointsBlob:  value.AccountProperties.SecondaryEndpoints.Blob,
		AccountPropertiesSecondaryEndpointsQueue: value.AccountProperties.SecondaryEndpoints.Queue,
		AccountPropertiesSecondaryEndpointsTable: value.AccountProperties.SecondaryEndpoints.Table,
		AccountPropertiesSecondaryEndpointsFile:  value.AccountProperties.SecondaryEndpoints.File,
		AccountPropertiesSecondaryEndpointsWeb:   value.AccountProperties.SecondaryEndpoints.Web,
		AccountPropertiesSecondaryEndpointsDfs:   value.AccountProperties.SecondaryEndpoints.Dfs,

		//AccountPropertiesSecondaryEndpointsMicrosoftEndpointsBlob: value.AccountProperties.SecondaryEndpoints.MicrosoftEndpoints.Blob,
		//AccountPropertiesSecondaryEndpointsMicrosoftEndpointsQueue: value.AccountProperties.SecondaryEndpoints.MicrosoftEndpoints.Queue,
		//AccountPropertiesSecondaryEndpointsMicrosoftEndpointsTable: value.AccountProperties.SecondaryEndpoints.MicrosoftEndpoints.Table,
		//AccountPropertiesSecondaryEndpointsMicrosoftEndpointsFile: value.AccountProperties.SecondaryEndpoints.MicrosoftEndpoints.File,
		//AccountPropertiesSecondaryEndpointsMicrosoftEndpointsWeb: value.AccountProperties.SecondaryEndpoints.MicrosoftEndpoints.Web,
		//AccountPropertiesSecondaryEndpointsMicrosoftEndpointsDfs: value.AccountProperties.SecondaryEndpoints.MicrosoftEndpoints.Dfs,

		//AccountPropertiesSecondaryEndpointsInternetEndpointsBlob: value.AccountProperties.SecondaryEndpoints.InternetEndpoints.Blob,
		//AccountPropertiesSecondaryEndpointsInternetEndpointsFile: value.AccountProperties.SecondaryEndpoints.InternetEndpoints.File,
		//AccountPropertiesSecondaryEndpointsInternetEndpointsWeb: value.AccountProperties.SecondaryEndpoints.InternetEndpoints.Web,
		//AccountPropertiesSecondaryEndpointsInternetEndpointsDfs: value.AccountProperties.SecondaryEndpoints.InternetEndpoints.Dfs,

		AccountPropertiesEncryptionServicesBlobEnabled:         value.AccountProperties.Encryption.Services.Blob.Enabled,
		AccountPropertiesEncryptionServicesBlobLastEnabledTime: value.AccountProperties.Encryption.Services.Blob.LastEnabledTime.ToTime(),
		AccountPropertiesEncryptionServicesBlobKeyType:         string(value.AccountProperties.Encryption.Services.Blob.KeyType),

		AccountPropertiesEncryptionServicesFileEnabled:         value.AccountProperties.Encryption.Services.File.Enabled,
		AccountPropertiesEncryptionServicesFileLastEnabledTime: value.AccountProperties.Encryption.Services.File.LastEnabledTime.ToTime(),
		AccountPropertiesEncryptionServicesFileKeyType:         string(value.AccountProperties.Encryption.Services.File.KeyType),

		//AccountPropertiesEncryptionServicesTableEnabled: value.AccountProperties.Encryption.Services.Table.Enabled,
		//AccountPropertiesEncryptionServicesTableLastEnabledTime: value.AccountProperties.Encryption.Services.Table.LastEnabledTime.ToTime(),
		//AccountPropertiesEncryptionServicesTableKeyType: string(value.AccountProperties.Encryption.Services.Table.KeyType),

		//AccountPropertiesEncryptionServicesQueueEnabled: value.AccountProperties.Encryption.Services.Queue.Enabled,
		//AccountPropertiesEncryptionServicesQueueLastEnabledTime: value.AccountProperties.Encryption.Services.Queue.LastEnabledTime.ToTime(),
		//AccountPropertiesEncryptionServicesQueueKeyType: string(value.AccountProperties.Encryption.Services.Queue.KeyType),
		AccountPropertiesEncryptionKeySource:                       string(value.AccountProperties.Encryption.KeySource),
		AccountPropertiesEncryptionRequireInfrastructureEncryption: value.AccountProperties.Encryption.RequireInfrastructureEncryption,

		//AccountPropertiesEncryptionKeyVaultPropertiesKeyName: value.AccountProperties.Encryption.KeyVaultProperties.KeyName,
		//AccountPropertiesEncryptionKeyVaultPropertiesKeyVersion: value.AccountProperties.Encryption.KeyVaultProperties.KeyVersion,
		//AccountPropertiesEncryptionKeyVaultPropertiesKeyVaultURI: value.AccountProperties.Encryption.KeyVaultProperties.KeyVaultURI,
		//AccountPropertiesEncryptionKeyVaultPropertiesCurrentVersionedKeyIdentifier: value.AccountProperties.Encryption.KeyVaultProperties.CurrentVersionedKeyIdentifier,
		//AccountPropertiesEncryptionKeyVaultPropertiesLastKeyRotationTimestamp: value.AccountProperties.Encryption.KeyVaultProperties.LastKeyRotationTimestamp.ToTime(),
		AccountPropertiesAccessTier: string(value.AccountProperties.AccessTier),

		//AccountPropertiesAzureFilesIdentityBasedAuthenticationDirectoryServiceOptions: string(value.AccountProperties.AzureFilesIdentityBasedAuthentication.DirectoryServiceOptions),

		//AccountPropertiesAzureFilesIdentityBasedAuthenticationActiveDirectoryPropertiesDomainName: value.AccountProperties.AzureFilesIdentityBasedAuthentication.ActiveDirectoryProperties.DomainName,
		//AccountPropertiesAzureFilesIdentityBasedAuthenticationActiveDirectoryPropertiesNetBiosDomainName: value.AccountProperties.AzureFilesIdentityBasedAuthentication.ActiveDirectoryProperties.NetBiosDomainName,
		//AccountPropertiesAzureFilesIdentityBasedAuthenticationActiveDirectoryPropertiesForestName: value.AccountProperties.AzureFilesIdentityBasedAuthentication.ActiveDirectoryProperties.ForestName,
		//AccountPropertiesAzureFilesIdentityBasedAuthenticationActiveDirectoryPropertiesDomainGUID: value.AccountProperties.AzureFilesIdentityBasedAuthentication.ActiveDirectoryProperties.DomainGUID,
		//AccountPropertiesAzureFilesIdentityBasedAuthenticationActiveDirectoryPropertiesDomainSid: value.AccountProperties.AzureFilesIdentityBasedAuthentication.ActiveDirectoryProperties.DomainSid,
		//AccountPropertiesAzureFilesIdentityBasedAuthenticationActiveDirectoryPropertiesAzureStorageSid: value.AccountProperties.AzureFilesIdentityBasedAuthentication.ActiveDirectoryProperties.AzureStorageSid,
		AccountPropertiesEnableHTTPSTrafficOnly: value.AccountProperties.EnableHTTPSTrafficOnly,

		AccountPropertiesNetworkRuleSetBypass:              string(value.AccountProperties.NetworkRuleSet.Bypass),
		AccountPropertiesNetworkRuleSetVirtualNetworkRules: transformAccountVirtualNetworkRules(*value.AccountProperties.NetworkRuleSet.VirtualNetworkRules),
		AccountPropertiesNetworkRuleSetIPRules:             transformAccountIPRules(*(value.AccountProperties.NetworkRuleSet.IPRules)),
		AccountPropertiesNetworkRuleSetDefaultAction:       string(value.AccountProperties.NetworkRuleSet.DefaultAction),
		AccountPropertiesIsHnsEnabled:                      value.AccountProperties.IsHnsEnabled,

		//AccountPropertiesGeoReplicationStatsStatus: string(value.AccountProperties.GeoReplicationStats.Status),
		//AccountPropertiesGeoReplicationStatsLastSyncTime: value.AccountProperties.GeoReplicationStats.LastSyncTime.ToTime(),
		//AccountPropertiesGeoReplicationStatsCanFailover: value.AccountProperties.GeoReplicationStats.CanFailover,
		AccountPropertiesFailoverInProgress:         value.AccountProperties.FailoverInProgress,
		AccountPropertiesLargeFileSharesState:       string(value.AccountProperties.LargeFileSharesState),
		AccountPropertiesPrivateEndpointConnections: transformAccountPrivateEndpointConnections(*value.AccountProperties.PrivateEndpointConnections),

		//AccountPropertiesRoutingPreferenceRoutingChoice: string(value.AccountProperties.RoutingPreference.RoutingChoice),
		//AccountPropertiesRoutingPreferencePublishMicrosoftEndpoints: value.AccountProperties.RoutingPreference.PublishMicrosoftEndpoints,
		//AccountPropertiesRoutingPreferencePublishInternetEndpoints: value.AccountProperties.RoutingPreference.PublishInternetEndpoints,

		//AccountPropertiesBlobRestoreStatusStatus: string(value.AccountProperties.BlobRestoreStatus.Status),
		//AccountPropertiesBlobRestoreStatusFailureReason: value.AccountProperties.BlobRestoreStatus.FailureReason,
		//AccountPropertiesBlobRestoreStatusRestoreID: value.AccountProperties.BlobRestoreStatus.RestoreID,

		//AccountPropertiesBlobRestoreStatusParametersTimeToRestore: value.AccountProperties.BlobRestoreStatus.Parameters.TimeToRestore.ToTime(),
		//AccountPropertiesBlobRestoreStatusParametersBlobRanges: transformAccountBlobRestoreRanges(*value.AccountProperties.BlobRestoreStatus.Parameters.BlobRanges),
		AccountPropertiesAllowBlobPublicAccess: value.AccountProperties.AllowBlobPublicAccess,
		AccountPropertiesMinimumTLSVersion:     string(value.AccountProperties.MinimumTLSVersion),
		Location:                               value.Location,
		//ResourceID: value.Id,
		Name: value.Name,
		Type: value.Type,
	}
}

func transformAccounts(values []storage.Account) []Account {
	var tValues []Account
	for _, v := range values {
		tValues = append(tValues, transformAccount(v))
	}
	return tValues
}

func transformAccountVirtualNetworkRule(value storage.VirtualNetworkRule) AccountVirtualNetworkRule {
	return AccountVirtualNetworkRule{
		VirtualNetworkResourceID: value.VirtualNetworkResourceID,
		Action:                   string(value.Action),
		State:                    string(value.State),
	}
}

func transformAccountVirtualNetworkRules(values []storage.VirtualNetworkRule) []AccountVirtualNetworkRule {
	var tValues []AccountVirtualNetworkRule
	for _, v := range values {
		tValues = append(tValues, transformAccountVirtualNetworkRule(v))
	}
	return tValues
}

func transformAccountIPRule(value storage.IPRule) AccountIPRule {
	return AccountIPRule{
		IPAddressOrRange: value.IPAddressOrRange,
		Action:           string(value.Action),
	}
}

func transformAccountIPRules(values []storage.IPRule) []AccountIPRule {
	var tValues []AccountIPRule
	for _, v := range values {
		tValues = append(tValues, transformAccountIPRule(v))
	}
	return tValues
}

func transformAccountPrivateEndpointConnection(value storage.PrivateEndpointConnection) AccountPrivateEndpointConnection {
	return AccountPrivateEndpointConnection{
		PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateStatus:         string(value.PrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState.Status),
		PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateDescription:    value.PrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState.Description,
		PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateActionRequired: value.PrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState.ActionRequired,
		PrivateEndpointConnectionPropertiesProvisioningState:                               string(value.PrivateEndpointConnectionProperties.ProvisioningState),
		//ResourceID: value.Id,
		Name: value.Name,
		Type: value.Type,
	}
}

func transformAccountPrivateEndpointConnections(values []storage.PrivateEndpointConnection) []AccountPrivateEndpointConnection {
	var tValues []AccountPrivateEndpointConnection
	for _, v := range values {
		tValues = append(tValues, transformAccountPrivateEndpointConnection(v))
	}
	return tValues
}

func transformAccountBlobRestoreRange(value storage.BlobRestoreRange) AccountBlobRestoreRange {
	return AccountBlobRestoreRange{
		StartRange: value.StartRange,
		EndRange:   value.EndRange,
	}
}

func transformAccountBlobRestoreRanges(values []storage.BlobRestoreRange) []AccountBlobRestoreRange {
	var tValues []AccountBlobRestoreRange
	for _, v := range values {
		tValues = append(tValues, transformAccountBlobRestoreRange(v))
	}
	return tValues
}

type AccountConfig struct {
	Filter string
}

func MigrateAccount(db *database.Database) error {
	err := db.AutoMigrate(
		&Account{},
		&AccountVirtualNetworkRule{},
		&AccountIPRule{},
		&AccountPrivateEndpointConnection{},
		&AccountBlobRestoreRange{},
	)
	if err != nil {
		return err
	}

	return nil
}

func Accounts(subscriptionID string, auth autorest.Authorizer, db *database.Database, log *zap.Logger, gConfig interface{}) error {
	var config AccountConfig
	ctx := context.Background()
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	svc := storage.NewAccountsClient("d61dc1dc-e66f-4429-ae51-4e469be0b0ca")
	svc.Authorizer = auth
	output, err := svc.List(ctx)

	for output.NotDone() {
		vals := output.Values()
		db.ChunkedCreate(transformAccounts(vals))
		log.Info("populating Accounts", zap.Int("count", len(vals)))
		output.NextWithContext(ctx)
	}

	//db.Where("project_id = ?", c.projectID).Delete(&Account{})
	//var tValues []*Account
	//for _, items := range output.Items {
	//	tValues = append(tValues, transformAccounts(items)...)
	//}

	return nil
}
