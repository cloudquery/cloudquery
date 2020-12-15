package compute

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-06-01/compute"
	"github.com/Azure/go-autorest/autorest"
	"github.com/cloudquery/cloudquery/providers/azure/utils"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type Disk struct {
	ID                uint `gorm:"primarykey"`
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

	EncryptionSettingsCollectionEnabled                   *bool
	EncryptionSettingsCollectionEncryptionSettings        []*DiskEncryptionSettingsElement `gorm:"constraint:OnDelete:CASCADE;"`
	EncryptionSettingsCollectionEncryptionSettingsVersion *string
	ProvisioningState                                     *string
	DiskIOPSReadWrite                                     *int64
	DiskMBpsReadWrite                                     *int64
	DiskIOPSReadOnly                                      *int64
	DiskMBpsReadOnly                                      *int64
	DiskState                                             string

	EncryptionDiskEncryptionSetID *string
	EncryptionType                string
	MaxShares                     *int32
	ShareInfo                     []*DiskShareInfoElement `gorm:"constraint:OnDelete:CASCADE;"`
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
	ID     uint `gorm:"primarykey"`
	DiskID uint
	Value  string
}
type DiskZones struct {
	ID     uint `gorm:"primarykey"`
	DiskID uint
	Value  string
}
type DiskEncryptionSettingsElement struct {
	ID     uint `gorm:"primarykey"`
	DiskID uint

	DiskEncryptionKeySourceVaultID *string
	DiskEncryptionKeySecretURL     *string

	KeyEncryptionKeySourceVaultID *string
	KeyEncryptionKeyURL           *string
}

func (DiskEncryptionSettingsElement) TableName() string {
	return "azure_compute_disk_encryption_settings_elements"
}

type DiskShareInfoElement struct {
	ID     uint `gorm:"primarykey"`
	DiskID uint
	VMURI  *string
}

func (DiskShareInfoElement) TableName() string {
	return "azure_compute_disk_share_info_elements"
}

type DiskTag struct {
	ID     uint
	DiskID uint
	Key    string
	Value  *string
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
			Tags:           transformDiskTags(value.Tags),
		}

		if value.ManagedByExtended != nil {
			tValue.ManagedByExtended = transformDiskManagedByExtended(*value.ManagedByExtended)
		}

		if value.Zones != nil {
			tValue.Zones = transformDiskZones(*value.Zones)
		}

		if value.Sku != nil {
			tValue.SkuName = string(value.Sku.Name)
			tValue.SkuTier = value.Sku.Tier
		}

		if value.DiskProperties != nil {

			tValue.TimeCreated = utils.AzureDateToTime(value.DiskProperties.TimeCreated)
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
				tValue.ShareInfo = transformDiskShareInfoElements(value.DiskProperties.ShareInfo)
			}
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}
func transformDiskManagedByExtended(values []string) []*DiskManagedByExtended {
	var tValues []*DiskManagedByExtended
	for _, v := range values {
		tValues = append(tValues, &DiskManagedByExtended{
			Value: v,
		})
	}
	return tValues
}

func transformDiskZones(values []string) []*DiskZones {
	var tValues []*DiskZones
	for _, v := range values {
		tValues = append(tValues, &DiskZones{
			Value: v,
		})
	}
	return tValues
}

func transformDiskEncryptionSettingsElements(values *[]compute.EncryptionSettingsElement) []*DiskEncryptionSettingsElement {
	var tValues []*DiskEncryptionSettingsElement
	for _, value := range *values {
		tValue := DiskEncryptionSettingsElement{}
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

func transformDiskShareInfoElements(values *[]compute.ShareInfoElement) []*DiskShareInfoElement {
	var tValues []*DiskShareInfoElement
	for _, value := range *values {
		tValue := DiskShareInfoElement{
			VMURI: value.VMURI,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func transformDiskTags(values map[string]*string) []*DiskTag {
	var tValues []*DiskTag
	for k, v := range values {
		tValue := DiskTag{
			Key:   k,
			Value: v,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

type DiskConfig struct {
	Filter string
}

func MigrateDisk(db *gorm.DB) error {
	err := db.AutoMigrate(
		&Disk{},
		&DiskEncryptionSettingsElement{},
		&DiskShareInfoElement{},
		&DiskTag{},
	)
	if err != nil {
		return err
	}

	return nil
}

func Disks(subscriptionID string, auth autorest.Authorizer, db *gorm.DB, log *zap.Logger, gConfig interface{}) error {
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

	db.Where("subscription_id = ?", subscriptionID).Delete(&Disk{})
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
		common.ChunkedCreate(db, tValues)
		log.Info("Fetched resources", zap.Int("count", len(tValues)))
	}

	return nil
}
