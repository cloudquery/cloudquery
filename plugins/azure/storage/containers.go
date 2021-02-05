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

type ListContainerItem struct {
	ID        uint `gorm:"primarykey"`
	AccountID string

	ContainerPropertiesVersion                     *string
	ContainerPropertiesDeleted                     *bool
	ContainerPropertiesDeletedTime                 time.Time
	ContainerPropertiesRemainingRetentionDays      *int32
	ContainerPropertiesDefaultEncryptionScope      *string
	ContainerPropertiesDenyEncryptionScopeOverride *bool
	ContainerPropertiesPublicAccess                string
	ContainerPropertiesLastModifiedTime            time.Time
	ContainerPropertiesLeaseStatus                 string
	ContainerPropertiesLeaseState                  string
	ContainerPropertiesLeaseDuration               string

	ContainerPropertiesImmutabilityPolicyImmutabilityPolicyPropertyImmutabilityPeriodSinceCreationInDays *int32
	ContainerPropertiesImmutabilityPolicyImmutabilityPolicyPropertyState                                 string
	ContainerPropertiesImmutabilityPolicyImmutabilityPolicyPropertyAllowProtectedAppendWrites            *bool
	ContainerPropertiesImmutabilityPolicyEtag                                                            *string
	ContainerPropertiesImmutabilityPolicyUpdateHistory                                                   []ListContainerItemUpdateHistoryProperty `gorm:"constraint:OnDelete:CASCADE;"`

	ContainerPropertiesLegalHoldHasLegalHold *bool
	ContainerPropertiesLegalHoldTags         []ListContainerItemTagProperty `gorm:"constraint:OnDelete:CASCADE;"`
	ContainerPropertiesHasLegalHold          *bool
	ContainerPropertiesHasImmutabilityPolicy *bool
	Etag                                     *string
	ResourceID                               *string
	Name                                     *string
	Type                                     *string
}

func (ListContainerItem) TableName() string {
	return "azure_storage_listcontaineritems"
}

type ListContainerItemUpdateHistoryProperty struct {
	ID                                              uint `gorm:"primarykey"`
	ListContainerItemID                             uint
	ListContainerItemImmutabilityPolicyPropertiesID uint
	Update                                          string
	ImmutabilityPeriodSinceCreationInDays           *int32
	Timestamp                                       time.Time
	ObjectIdentifier                                *string
	TenantID                                        *string
	Upn                                             *string
}

func (ListContainerItemUpdateHistoryProperty) TableName() string {
	return "azure_storage_listcontaineritemupdatehistorypropertys"
}

type ListContainerItemTagProperty struct {
	ID                                     uint `gorm:"primarykey"`
	ListContainerItemID                    uint
	ListContainerItemLegalHoldPropertiesID uint
	Tag                                    *string
	Timestamp                              time.Time
	ObjectIdentifier                       *string
	TenantID                               *string
	Upn                                    *string
}

func (ListContainerItemTagProperty) TableName() string {
	return "azure_storage_listcontaineritemtagpropertys"
}

func transformListContainerItem(value storage.ListContainerItem) ListContainerItem {
	res := ListContainerItem{
		//SubscriptionID: SubscriptionID,

		ContainerPropertiesVersion:                     value.ContainerProperties.Version,
		ContainerPropertiesDeleted:                     value.ContainerProperties.Deleted,
		ContainerPropertiesRemainingRetentionDays:      value.ContainerProperties.RemainingRetentionDays,
		ContainerPropertiesDefaultEncryptionScope:      value.ContainerProperties.DefaultEncryptionScope,
		ContainerPropertiesDenyEncryptionScopeOverride: value.ContainerProperties.DenyEncryptionScopeOverride,
		ContainerPropertiesPublicAccess:                string(value.ContainerProperties.PublicAccess),
		ContainerPropertiesLastModifiedTime:            value.ContainerProperties.LastModifiedTime.ToTime(),
		ContainerPropertiesLeaseStatus:                 string(value.ContainerProperties.LeaseStatus),
		ContainerPropertiesLeaseState:                  string(value.ContainerProperties.LeaseState),
		ContainerPropertiesLeaseDuration:               string(value.ContainerProperties.LeaseDuration),

		//ContainerPropertiesLegalHoldHasLegalHold: value.ContainerProperties.LegalHold.HasLegalHold,
		ContainerPropertiesHasLegalHold:          value.ContainerProperties.HasLegalHold,
		ContainerPropertiesHasImmutabilityPolicy: value.ContainerProperties.HasImmutabilityPolicy,
		Etag:                                     value.Etag,
		ResourceID:                               value.ID,
		Name:                                     value.Name,
		Type:                                     value.Type,
	}

	if value.ContainerProperties.ImmutabilityPolicy != nil && value.ContainerProperties.ImmutabilityPolicy.ImmutabilityPolicyProperty != nil {
		res.ContainerPropertiesImmutabilityPolicyImmutabilityPolicyPropertyImmutabilityPeriodSinceCreationInDays = value.ContainerProperties.ImmutabilityPolicy.ImmutabilityPolicyProperty.ImmutabilityPeriodSinceCreationInDays
		res.ContainerPropertiesImmutabilityPolicyImmutabilityPolicyPropertyState = string(value.ContainerProperties.ImmutabilityPolicy.ImmutabilityPolicyProperty.State)
		res.ContainerPropertiesImmutabilityPolicyImmutabilityPolicyPropertyAllowProtectedAppendWrites = value.ContainerProperties.ImmutabilityPolicy.ImmutabilityPolicyProperty.AllowProtectedAppendWrites
		res.ContainerPropertiesImmutabilityPolicyEtag = value.ContainerProperties.ImmutabilityPolicy.Etag
	}

	if value.ContainerProperties.DeletedTime != nil {
		res.ContainerPropertiesDeletedTime = value.ContainerProperties.DeletedTime.ToTime()
	}

	if value.ContainerProperties.ImmutabilityPolicy != nil && value.ContainerProperties.ImmutabilityPolicy.UpdateHistory != nil {
		res.ContainerPropertiesImmutabilityPolicyUpdateHistory = transformListContainerItemUpdateHistoryPropertys(*value.ContainerProperties.ImmutabilityPolicy.UpdateHistory)
	}

	if value.ContainerProperties.LegalHold != nil && value.ContainerProperties.LegalHold.Tags != nil {
		res.ContainerPropertiesLegalHoldTags = transformListContainerItemTagPropertys(*value.ContainerProperties.LegalHold.Tags)
	}

	return res
}

func transformListContainerItems(values []storage.ListContainerItem) []ListContainerItem {
	var tValues []ListContainerItem
	for _, v := range values {
		tValues = append(tValues, transformListContainerItem(v))
	}
	return tValues
}

func transformListContainerItemUpdateHistoryProperty(value storage.UpdateHistoryProperty) ListContainerItemUpdateHistoryProperty {
	return ListContainerItemUpdateHistoryProperty{
		Update:                                string(value.Update),
		ImmutabilityPeriodSinceCreationInDays: value.ImmutabilityPeriodSinceCreationInDays,
		Timestamp:                             value.Timestamp.ToTime(),
		ObjectIdentifier:                      value.ObjectIdentifier,
		TenantID:                              value.TenantID,
		Upn:                                   value.Upn,
	}
}

func transformListContainerItemUpdateHistoryPropertys(values []storage.UpdateHistoryProperty) []ListContainerItemUpdateHistoryProperty {
	var tValues []ListContainerItemUpdateHistoryProperty
	for _, v := range values {
		tValues = append(tValues, transformListContainerItemUpdateHistoryProperty(v))
	}
	return tValues
}

func transformListContainerItemTagProperty(value storage.TagProperty) ListContainerItemTagProperty {
	return ListContainerItemTagProperty{
		Tag:              value.Tag,
		Timestamp:        value.Timestamp.ToTime(),
		ObjectIdentifier: value.ObjectIdentifier,
		TenantID:         value.TenantID,
		Upn:              value.Upn,
	}
}

func transformListContainerItemTagPropertys(values []storage.TagProperty) []ListContainerItemTagProperty {
	var tValues []ListContainerItemTagProperty
	for _, v := range values {
		tValues = append(tValues, transformListContainerItemTagProperty(v))
	}
	return tValues
}

type ListContainerItemConfig struct {
	Filter string
}

func MigrateListContainerItem(db *database.Database) error {
	err := db.AutoMigrate(
		&ListContainerItem{},
		&ListContainerItemUpdateHistoryProperty{},
		&ListContainerItemTagProperty{},
	)
	if err != nil {
		return err
	}

	return nil
}

func ListContainerItems(subscriptionID string, auth autorest.Authorizer, db *database.Database, log *zap.Logger, gConfig interface{}) error {
	var config ListContainerItemConfig
	ctx := context.Background()
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	svc := storage.NewBlobContainersClient(subscriptionID)
	svc.Authorizer = auth
	output, err := svc.List(ctx, "testResourceGroup", "gigwitstorage", "", "", "")
	if err != nil {
		return err
	}
	for output.NotDone() {
		vals := output.Values()
		db.ChunkedCreate(transformListContainerItems(vals))
		log.Info("populating ListContainerItems", zap.Int("count", len(vals)))
		output.NextWithContext(ctx)
	}

	//db.Where("project_id = ?", c.projectID).Delete(&ListContainerItem{})
	//var tValues []*ListContainerItem
	//for _, items := range output.Items {
	//	tValues = append(tValues, transformListContainerItems(items)...)
	//}

	return nil
}
