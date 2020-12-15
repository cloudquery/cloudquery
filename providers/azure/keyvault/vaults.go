package keyvault

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2019-09-01/keyvault"
	"github.com/Azure/go-autorest/autorest"
	"github.com/cloudquery/cloudquery/providers/azure/utils"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"regexp"
)

type Vault struct {
	ID             uint `gorm:"primarykey"`
	SubscriptionID string
	ResourceID     *string
	Name           *string
	Type           *string
	Location       *string
	Tags           []*VaultTag `gorm:"constraint:OnDelete:CASCADE;"`
	Keys           []*Key      `gorm:"constraint:OnDelete:CASCADE;"`

	TenantID *string

	SkuFamily                    *string
	SkuName                      string
	AccessPolicies               []*VaultAccessPolicy `gorm:"constraint:OnDelete:CASCADE;"`
	VaultURI                     *string
	EnabledForDeployment         *bool
	EnabledForDiskEncryption     *bool
	EnabledForTemplateDeployment *bool
	EnableSoftDelete             *bool
	SoftDeleteRetentionInDays    *int32
	EnableRbacAuthorization      *bool
	CreateMode                   string
	EnablePurgeProtection        *bool

	NetworkAclsBypass              string
	NetworkAclsDefaultAction       string
	NetworkAclsIPRules             []*VaultIPRule                    `gorm:"constraint:OnDelete:CASCADE;"`
	NetworkAclsVirtualNetworkRules []*VaultVirtualNetworkRule        `gorm:"constraint:OnDelete:CASCADE;"`
	PrivateEndpointConnections     []*VaultPrivateEndpointConnection `gorm:"constraint:OnDelete:CASCADE;"`
}

func (Vault) TableName() string {
	return "azure_keyvault_vaults"
}

type VaultTag struct {
	ID      uint
	VaultID uint
	Key     string
	Value   *string
}

func (VaultTag) TableName() string {
	return "azure_keyvault_tags"
}

type VaultIPRule struct {
	ID      uint `gorm:"primarykey"`
	VaultID uint
	Value   *string
}

func (VaultIPRule) TableName() string {
	return "azure_keyvault_vault_ip_rules"
}

type VaultVirtualNetworkRule struct {
	ID         uint `gorm:"primarykey"`
	VaultID    uint
	ResourceID *string
}

func (VaultVirtualNetworkRule) TableName() string {
	return "azure_keyvault_vault_virtual_network_rules"
}

type VaultPrivateEndpointConnection struct {
	ID      uint `gorm:"primarykey"`
	VaultID uint

	ResourceID                                      *string
	PrivateLinkServiceConnectionStateStatus         string
	PrivateLinkServiceConnectionStateDescription    *string
	PrivateLinkServiceConnectionStateActionRequired *string
	ProvisioningState                               string
}

func (VaultPrivateEndpointConnection) TableName() string {
	return "azure_keyvault_vault_private_endpoint_connections"
}

type VaultAccessPolicy struct {
	ID            uint `gorm:"primarykey"`
	VaultID       uint
	TenantID      *string
	ObjectID      *string
	ApplicationID *string

	KeyPermissions         []*VaultKeyPermission         `gorm:"constraint:OnDelete:CASCADE;"`
	SecretPermissions      []*VaultSecretPermission      `gorm:"constraint:OnDelete:CASCADE;"`
	CertificatePermissions []*VaultCertificatePermission `gorm:"constraint:OnDelete:CASCADE;"`
	StoragePermissions     []*VaultStoragePermission     `gorm:"constraint:OnDelete:CASCADE;"`
}

func (VaultAccessPolicy) TableName() string {
	return "azure_keyvault_vault_access_policies"
}

type VaultKeyPermission struct {
	ID                  uint `gorm:"primarykey"`
	VaultAccessPolicyID uint
	Value               string
}

func (VaultKeyPermission) TableName() string {
	return "azure_keyvault_vault_access_policy_key_permissions"
}

type VaultSecretPermission struct {
	ID                  uint `gorm:"primarykey"`
	VaultAccessPolicyID uint
	Value               string
}

func (VaultSecretPermission) TableName() string {
	return "azure_keyvault_vault_access_policy_secret_permissions"
}

type VaultCertificatePermission struct {
	ID                  uint `gorm:"primarykey"`
	VaultAccessPolicyID uint
	Value               string
}

func (VaultCertificatePermission) TableName() string {
	return "azure_keyvault_vault_access_policy_certificate_permissions"
}

type VaultStoragePermission struct {
	ID                  uint `gorm:"primarykey"`
	VaultAccessPolicyID uint
	Value               string
}

func (VaultStoragePermission) TableName() string {
	return "azure_keyvault_vault_access_policy_storage_permissions"
}

func transformVaults(subscriptionID string, auth autorest.Authorizer, values []keyvault.Vault) ([]*Vault, error) {
	var tValues []*Vault
	ctx := context.Background()
	resourceGroupRe := regexp.MustCompile("resourceGroups/([a-zA-Z0-9-_]+)/")
	for _, value := range values {
		tValue := Vault{
			SubscriptionID: subscriptionID,
			ResourceID:     value.ID,
			Name:           value.Name,
			Type:           value.Type,
			Location:       value.Location,
			Tags:           transformVaultTags(value.Tags),
		}

		match := resourceGroupRe.FindStringSubmatch(*value.ID)
		if len(match) < 2 {
			return nil, fmt.Errorf("coultn't extract resource group from %s", *value.ID)
		}
		svc := keyvault.NewKeysClient(subscriptionID)
		svc.Authorizer = auth
		keysResult, err := svc.List(ctx, match[1], *value.Name)
		if err != nil {
			return nil, err
		}
		for keysResult.NotDone() {
			keys := keysResult.Values()
			tValue.Keys = append(tValue.Keys, transformKeys(subscriptionID, keys)...)
			err := keysResult.NextWithContext(ctx)
			if err != nil {
				return nil, err
			}
		}

		if value.Properties != nil {
			tValue.TenantID = utils.AzureUUIDToString(value.Properties.TenantID)
			tValue.AccessPolicies = transformVaultAccessPolicyEntries(value.Properties.AccessPolicies)
			tValue.VaultURI = value.Properties.VaultURI
			tValue.EnabledForDeployment = value.Properties.EnabledForDeployment
			tValue.EnabledForDiskEncryption = value.Properties.EnabledForDiskEncryption
			tValue.EnabledForTemplateDeployment = value.Properties.EnabledForTemplateDeployment
			tValue.EnableSoftDelete = value.Properties.EnableSoftDelete
			tValue.SoftDeleteRetentionInDays = value.Properties.SoftDeleteRetentionInDays
			tValue.EnableRbacAuthorization = value.Properties.EnableRbacAuthorization
			tValue.CreateMode = string(value.Properties.CreateMode)
			tValue.EnablePurgeProtection = value.Properties.EnablePurgeProtection

			if value.Properties.NetworkAcls != nil {
				tValue.NetworkAclsIPRules = transformVaultIPRules(value.Properties.NetworkAcls.IPRules)
				tValue.NetworkAclsVirtualNetworkRules = transformVaultVirtualNetworkRules(value.Properties.NetworkAcls.VirtualNetworkRules)
			}

			tValue.PrivateEndpointConnections = transformVaultPrivateEndpointConnectionItems(value.Properties.PrivateEndpointConnections)
		}
		tValues = append(tValues, &tValue)
	}
	return tValues, nil
}

func transformVaultTags(values map[string]*string) []*VaultTag {
	var tValues []*VaultTag
	for k, v := range values {
		tValue := VaultTag{
			Key:   k,
			Value: v,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func transformVaultPermissionsKeys(values *[]keyvault.KeyPermissions) []*VaultKeyPermission {
	var tValues []*VaultKeyPermission
	if values == nil {
		return nil
	}
	for _, v := range *values {
		tValues = append(tValues, &VaultKeyPermission{
			Value: string(v),
		})
	}
	return tValues
}

func transformVaultPermissionsSecrets(values *[]keyvault.SecretPermissions) []*VaultSecretPermission {
	var tValues []*VaultSecretPermission
	if values == nil {
		return nil
	}
	for _, v := range *values {
		tValues = append(tValues, &VaultSecretPermission{
			Value: string(v),
		})
	}
	return tValues
}

func transformVaultPermissionsCertificates(values *[]keyvault.CertificatePermissions) []*VaultCertificatePermission {
	var tValues []*VaultCertificatePermission
	if values == nil {
		return nil
	}
	for _, v := range *values {
		tValues = append(tValues, &VaultCertificatePermission{
			Value: string(v),
		})
	}
	return tValues
}

func transformVaultPermissionsStorage(values *[]keyvault.StoragePermissions) []*VaultStoragePermission {
	var tValues []*VaultStoragePermission
	if values == nil {
		return nil
	}
	for _, v := range *values {
		tValues = append(tValues, &VaultStoragePermission{
			Value: string(v),
		})
	}
	return tValues
}

func transformVaultAccessPolicyEntries(values *[]keyvault.AccessPolicyEntry) []*VaultAccessPolicy {
	var tValues []*VaultAccessPolicy
	if values == nil {
		return nil
	}
	for _, value := range *values {
		tValue := VaultAccessPolicy{
			TenantID:      utils.AzureUUIDToString(value.TenantID),
			ObjectID:      value.ObjectID,
			ApplicationID: utils.AzureUUIDToString(value.ApplicationID),
		}
		if value.Permissions != nil {
			tValue.KeyPermissions = transformVaultPermissionsKeys(value.Permissions.Keys)
			tValue.SecretPermissions = transformVaultPermissionsSecrets(value.Permissions.Secrets)
			tValue.CertificatePermissions = transformVaultPermissionsCertificates(value.Permissions.Certificates)
			tValue.StoragePermissions = transformVaultPermissionsStorage(value.Permissions.Storage)
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func transformVaultIPRules(values *[]keyvault.IPRule) []*VaultIPRule {
	var tValues []*VaultIPRule
	if values == nil {
		return nil
	}
	for _, value := range *values {
		tValue := VaultIPRule{
			Value: value.Value,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func transformVaultVirtualNetworkRules(values *[]keyvault.VirtualNetworkRule) []*VaultVirtualNetworkRule {
	var tValues []*VaultVirtualNetworkRule
	if values == nil {
		return nil
	}
	for _, value := range *values {
		tValue := VaultVirtualNetworkRule{
			ResourceID: value.ID,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func transformVaultPrivateEndpointConnectionItems(values *[]keyvault.PrivateEndpointConnectionItem) []*VaultPrivateEndpointConnection {
	var tValues []*VaultPrivateEndpointConnection
	if values == nil {
		return nil
	}
	for _, value := range *values {
		tValue := VaultPrivateEndpointConnection{}
		if value.PrivateEndpointConnectionProperties != nil {

			tValue.ProvisioningState = string(value.PrivateEndpointConnectionProperties.ProvisioningState)
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

type VaultConfig struct {
	Filter string
}

func MigrateVault(db *gorm.DB) error {
	err := db.AutoMigrate(
		&Vault{},
		&VaultTag{},
		&VaultIPRule{},
		&VaultVirtualNetworkRule{},
		&VaultPrivateEndpointConnection{},
		&VaultAccessPolicy{},
		&VaultKeyPermission{},
		&VaultSecretPermission{},
		&VaultStoragePermission{},
		&VaultCertificatePermission{},
		&Key{},
		&KeyTag{},
	)
	if err != nil {
		return err
	}

	return nil
}

func Vaults(subscriptionID string, auth autorest.Authorizer, db *gorm.DB, log *zap.Logger, gConfig interface{}) error {
	var config VaultConfig
	ctx := context.Background()
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	svc := keyvault.NewVaultsClient(subscriptionID)
	svc.Authorizer = auth
	maxResults := int32(1000)
	output, err := svc.ListBySubscription(ctx, &maxResults)
	if err != nil {
		return err
	}

	db.Where("subscription_id = ?", subscriptionID).Delete(&Vault{})
	if !output.NotDone() {
		log.Info("Fetched resources", zap.Int("count", 0))
	}
	for output.NotDone() {
		values := output.Values()
		err := output.NextWithContext(ctx)
		if err != nil {
			return err
		}
		tValues, err := transformVaults(subscriptionID, auth, values)
		if err != nil {
			return err
		}
		common.ChunkedCreate(db, tValues)
		log.Info("Fetched resources", zap.Int("count", len(tValues)))
	}

	return nil
}
