package keyvault

import (
	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2019-09-01/keyvault"
)

type Key struct {
	ID             uint `gorm:"primarykey"`
	VaultID        uint `neo:"ignore"`
	SubscriptionID string

	AttributesEnabled       *bool
	AttributesNotBefore     *int64
	AttributesExpires       *int64
	AttributesCreated       *int64
	AttributesUpdated       *int64
	AttributesRecoveryLevel string
	Kty                     string
	KeyOps                  []*KeyOp `gorm:"constraint:OnDelete:CASCADE;"`
	KeySize                 *int32
	CurveName               string
	KeyURI                  *string
	KeyURIWithVersion       *string
	ResourceID              *string
	Name                    *string
	Type                    *string
	Location                *string
	Tags                    []*KeyTag `gorm:"constraint:OnDelete:CASCADE;"`
}

func (Key) TableName() string {
	return "azure_keyvault_keys"
}

type KeyOp struct {
	ID             uint   `gorm:"primarykey"`
	KeyID          uint   `neo:"ignore"`
	SubscriptionID string `gorm:"-"`
	Value          string
}

func (KeyOp) TableName() string {
	return "azure_keyvault_key_ops"
}

type KeyTag struct {
	ID             uint   `gorm:"primarykey"`
	KeyID          uint   `neo:"ignore"`
	SubscriptionID string `gorm:"-"`

	Key   string
	Value *string
}

func (KeyTag) TableName() string {
	return "azure_keyvault_keys_tags"
}

func transformKeys(subscriptionID string, values []keyvault.Key) []*Key {
	var tValues []*Key
	for _, value := range values {
		tValue := Key{
			SubscriptionID: subscriptionID,
			ResourceID:     value.ID,
			Name:           value.Name,
			Type:           value.Type,
			Location:       value.Location,
			Tags:           transformKeyTags(subscriptionID, value.Tags),
		}

		if value.Attributes != nil {
			tValue.AttributesCreated = value.Attributes.Created
			tValue.AttributesEnabled = value.Attributes.Enabled
			tValue.AttributesExpires = value.Attributes.Expires
			tValue.AttributesNotBefore = value.Attributes.NotBefore
			tValue.AttributesRecoveryLevel = string(value.Attributes.RecoveryLevel)
			tValue.AttributesUpdated = value.Attributes.Updated
		}

		if value.KeyProperties != nil {
			tValue.Kty = string(value.KeyProperties.Kty)
			tValue.KeyOps = transformKeyPropertiesKeyOps(subscriptionID, value.KeyProperties.KeyOps)
			tValue.KeySize = value.KeyProperties.KeySize
			tValue.CurveName = string(value.KeyProperties.CurveName)
			tValue.KeyURI = value.KeyProperties.KeyURI
			tValue.KeyURIWithVersion = value.KeyProperties.KeyURIWithVersion
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func transformKeyPropertiesKeyOps(subscriptionID string, values *[]keyvault.JSONWebKeyOperation) []*KeyOp {
	var tValues []*KeyOp
	if values == nil {
		return nil
	}
	for _, v := range *values {
		tValues = append(tValues, &KeyOp{
			SubscriptionID: subscriptionID,
			Value:          string(v),
		})
	}
	return tValues
}

func transformKeyTags(subscriptionID string, values map[string]*string) []*KeyTag {
	var tValues []*KeyTag
	for k, v := range values {
		tValue := KeyTag{
			SubscriptionID: subscriptionID,
			Key:            k,
			Value:          v,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

type KeyConfig struct {
	Filter string
}
