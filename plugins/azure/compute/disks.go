package compute

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-06-01/compute"
	"github.com/Azure/go-autorest/autorest"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cloudquery/providers/azure/utils"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"log"
	"time"
)

type Disk struct {
	ID                uint        `gorm:"primarykey"`
	_                 interface{} `neo:"raw:MERGE (a:AzureSubscription {subscription_id: $subscription_id}) MERGE (a) - [:Resource] -> (n)"`
	SubscriptionID    string
	ManagedBy         *string
	ManagedByExtended []*DiskManagedByExtended `gorm:"constraint:OnDelete:CASCADE;"`

	SkuName string
	SkuTier *string
	Zones   []*DiskZones `gorm:"constraint:OnDelete:CASCADE;"`

	TimeCreated      *time.Time
	OsType           string
	HyperVGeneration string

	CreationDataCreateOption     string
	CreationDataStorageAccountID *string

	CreationDataImageReferenceID  *string
	CreationDataImageReferenceLun *int32

	CreationDataGalleryImageReferenceID  *string
	CreationDataGalleryImageReferenceLun *int32
	CreationDataSourceURI                *string
	CreationDataSourceResourceID         *string
	CreationDataSourceUniqueID           *string
	CreationDataUploadSizeBytes          *int64
	DiskSizeGB                           *int32
	DiskSizeBytes                        *int64
	UniqueID                             *string

	EncryptionSettingsCollectionEnabled         *bool
	EncryptionSettingsCollectionSettings        []*DiskEncryptionSetting `gorm:"constraint:OnDelete:CASCADE;"`
	EncryptionSettingsCollectionSettingsVersion *string
	ProvisioningState                           *string
	DiskIOPSReadWrite                           *int64
	DiskMBpsReadWrite                           *int64
	DiskIOPSReadOnly                            *int64
	DiskMBpsReadOnly                            *int64
	DiskState                                   string

	EncryptionDiskEncryptionSetID *string
	EncryptionType                string
	MaxShares                     *int32
	ShareInfo                     []*DiskShareInfo `gorm:"constraint:OnDelete:CASCADE;"`
	NetworkAccessPolicy           string
	DiskAccessID                  *string
	ResourceID                    *string
	Name                          *string
	Type                          *string
	Location                      *string
	Tags                          []*DiskTag `gorm:"constraint:OnDelete:CASCADE;"`
}

func (Disk) TableName() string {
	return "azure_compute_disks"
}

type DiskManagedByExtended struct {
	ID             uint   `gorm:"primarykey"`
	DiskID         uint   `neo:"ignore"`
	SubscriptionID string `gorm:"-"`
	Value          string
}
type DiskZones struct {
	ID             uint   `gorm:"primarykey"`
	DiskID         uint   `neo:"ignore"`
	SubscriptionID string `gorm:"-"`
	Value          string
}
type DiskEncryptionSetting struct {
	ID             uint   `gorm:"primarykey"`
	DiskID         uint   `neo:"ignore"`
	SubscriptionID string `gorm:"-"`

	DiskEncryptionKeySourceVaultID *string
	DiskEncryptionKeySecretURL     *string

	KeyEncryptionKeySourceVaultID *string
	KeyEncryptionKeyURL           *string
}

func (DiskEncryptionSetting) TableName() string {
	return "azure_compute_disk_encryption_settings"
}

type DiskShareInfo struct {
	ID             uint   `gorm:"primarykey"`
	DiskID         uint   `neo:"ignore"`
	SubscriptionID string `gorm:"-"`
	VMURI          *string
}

func (DiskShareInfo) TableName() string {
	return "azure_compute_disk_share_info"
}

type DiskTag struct {
	ID             uint   `gorm:"primarykey"`
	DiskID         uint   `neo:"ignore"`
	SubscriptionID string `gorm:"-"`

	Key   string
	Value *string
}

func (DiskTag) TableName() string {
	return "azure_compute_disk_tags"
}

func transformDisks(subscriptionID string, values []compute.Disk) []*Disk {
	var tValues []*Disk
	for _, value := range values {
		tValue := Disk{
			SubscriptionID: subscriptionID,
			ManagedBy:      value.ManagedBy,
			ResourceID:     value.ID,
			Name:           value.Name,
			Type:           value.Type,
			Location:       value.Location,
			Tags:           transformDiskTags(subscriptionID, value.Tags),
		}

		if value.ManagedByExtended != nil {
			tValue.ManagedByExtended = transformDiskManagedByExtended(subscriptionID, *value.ManagedByExtended)
		}

		if value.EncryptionSettingsCollection != nil {
			tValue.EncryptionSettingsCollectionSettings = transformDiskEncryptionSettingsElements(subscriptionID, value.EncryptionSettingsCollection.EncryptionSettings)
			tValue.EncryptionSettingsCollectionEnabled = value.EncryptionSettingsCollection.Enabled
			tValue.EncryptionSettingsCollectionSettingsVersion = value.EncryptionSettingsCollection.EncryptionSettingsVersion
		}

		if value.Encryption != nil {
			tValue.EncryptionDiskEncryptionSetID = value.Encryption.DiskEncryptionSetID
			tValue.EncryptionType = string(value.Encryption.Type)
		}

		if value.Zones != nil {
			tValue.Zones = transformDiskZones(subscriptionID, *value.Zones)
		}

		if value.Sku != nil {
			tValue.SkuName = string(value.Sku.Name)
			tValue.SkuTier = value.Sku.Tier
		}

		if value.DiskProperties != nil {
			location, err := time.LoadLocation("UTC")
			if err != nil {
				log.Fatal(err)
			}
			timeCreated := utils.AzureDateToTime(value.DiskProperties.TimeCreated).In(location)
			tValue.TimeCreated = &timeCreated
			tValue.OsType = string(value.DiskProperties.OsType)
			tValue.HyperVGeneration = string(value.DiskProperties.HyperVGeneration)
			tValue.DiskSizeGB = value.DiskProperties.DiskSizeGB
			tValue.DiskSizeBytes = value.DiskProperties.DiskSizeBytes
			tValue.UniqueID = value.DiskProperties.UniqueID
			tValue.ProvisioningState = value.DiskProperties.ProvisioningState
			tValue.DiskIOPSReadWrite = value.DiskProperties.DiskIOPSReadWrite
			tValue.DiskMBpsReadWrite = value.DiskProperties.DiskMBpsReadWrite
			tValue.DiskIOPSReadOnly = value.DiskProperties.DiskIOPSReadOnly
			tValue.DiskMBpsReadOnly = value.DiskProperties.DiskMBpsReadOnly
			tValue.DiskState = string(value.DiskProperties.DiskState)
			tValue.MaxShares = value.DiskProperties.MaxShares
			tValue.NetworkAccessPolicy = string(value.DiskProperties.NetworkAccessPolicy)
			tValue.DiskAccessID = value.DiskProperties.DiskAccessID

			if value.DiskProperties.ShareInfo != nil {
				tValue.ShareInfo = transformDiskShareInfoElements(subscriptionID, value.DiskProperties.ShareInfo)
			}
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}
func transformDiskManagedByExtended(subscriptionID string, values []string) []*DiskManagedByExtended {
	var tValues []*DiskManagedByExtended
	for _, v := range values {
		tValues = append(tValues, &DiskManagedByExtended{
			SubscriptionID: subscriptionID,
			Value:          v,
		})
	}
	return tValues
}

func transformDiskZones(subscriptionID string, values []string) []*DiskZones {
	var tValues []*DiskZones
	for _, v := range values {
		tValues = append(tValues, &DiskZones{
			SubscriptionID: subscriptionID,
			Value:          v,
		})
	}
	return tValues
}

func transformDiskEncryptionSettingsElements(subscriptionID string, values *[]compute.EncryptionSettingsElement) []*DiskEncryptionSetting {
	var tValues []*DiskEncryptionSetting
	for _, value := range *values {
		tValue := DiskEncryptionSetting{
			SubscriptionID: subscriptionID,
		}
		if value.DiskEncryptionKey != nil {
			if value.DiskEncryptionKey.SourceVault != nil {
				return nil
			}
			tValue.DiskEncryptionKeySecretURL = value.DiskEncryptionKey.SecretURL
		}
		if value.KeyEncryptionKey != nil {
			if value.KeyEncryptionKey.SourceVault != nil {
				tValue.KeyEncryptionKeySourceVaultID = value.KeyEncryptionKey.SourceVault.ID
			}
			tValue.KeyEncryptionKeyURL = value.KeyEncryptionKey.KeyURL
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func transformDiskShareInfoElements(subscriptionID string, values *[]compute.ShareInfoElement) []*DiskShareInfo {
	var tValues []*DiskShareInfo
	for _, value := range *values {
		tValue := DiskShareInfo{
			SubscriptionID: subscriptionID,
			VMURI:          value.VMURI,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func transformDiskTags(subscriptionID string, values map[string]*string) []*DiskTag {
	var tValues []*DiskTag
	for k, v := range values {
		tValue := DiskTag{
			SubscriptionID: subscriptionID,
			Key:            k,
			Value:          v,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

type DiskConfig struct {
	Filter string
}

var DiskTables = []interface{}{
	&Disk{},
	&DiskEncryptionSetting{},
	&DiskShareInfo{},
	&DiskTag{},
}

func Disks(subscriptionID string, auth autorest.Authorizer, db *database.Database, log *zap.Logger, gConfig interface{}) error {
	var config DiskConfig
	ctx := context.Background()
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	svc := compute.NewDisksClient(subscriptionID)
	svc.Authorizer = auth
	output, err := svc.List(ctx)
	if err != nil {
		return err
	}

	db.Where("subscription_id", subscriptionID).Delete(DiskTables...)
	if !output.NotDone() {
		log.Info("Fetched resources", zap.Int("count", 0))
	}
	for output.NotDone() {
		values := output.Values()
		err := output.NextWithContext(ctx)
		if err != nil {
			return err
		}
		tValues := transformDisks(subscriptionID, values)
		db.ChunkedCreate(tValues)
		log.Info("Fetched resources", zap.Int("count", len(tValues)))
	}

	return nil
}
