package kms

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"log"
	"time"
)

type Key struct {
	ID        uint `gorm:"primarykey"`
	AccountID string
	Region    string

	Arn                   *string
	KeyId                 *string
	RotationEnabled       *bool
	CloudHsmClusterId     *string
	CreationDate          *time.Time
	CustomKeyStoreId      *string
	CustomerMasterKeySpec *string
	DeletionDate          *time.Time
	Description           *string
	Enabled               *bool
	EncryptionAlgorithms  []*KeyEncryptionAlgorithm `gorm:"constraint:OnDelete:CASCADE;"`
	ExpirationModel       *string
	Manager               *string
	KeyState              *string
	KeyUsage              *string
	Origin                *string
	SigningAlgorithms     []*KeySigningAlgorithm `gorm:"constraint:OnDelete:CASCADE;"`
	ValidTo               *time.Time
}

func (Key)TableName() string {
	return "aws_kms_keys"
}

type KeyEncryptionAlgorithm struct {
	ID    uint `gorm:"primarykey"`
	KeyID uint
	name  string
}

func (KeyEncryptionAlgorithm)TableName() string {
	return "aws_kms_key_encryption_algorithms"
}

type KeySigningAlgorithm struct {
	ID    uint `gorm:"primarykey"`
	KeyID uint
	name  string
}

func (KeySigningAlgorithm)TableName() string {
	return "aws_kms_key_signing_algorithms"
}

func (c *Client) transformKeyListEntry(value *kms.KeyListEntry) *Key {
	output, err := c.svc.DescribeKey(&kms.DescribeKeyInput{
		KeyId: value.KeyId,
	})
	if err != nil {
		log.Fatal(err)
	}
	outputKeyRotation, err := c.svc.GetKeyRotationStatus(&kms.GetKeyRotationStatusInput{
		KeyId: value.KeyId,
	})
	if err != nil {
		log.Fatal(err)
	}
	res := Key{
		Region:                c.region,
		AccountID:             c.accountID,
		Arn:                   value.KeyArn,
		KeyId:                 value.KeyId,
		RotationEnabled:       outputKeyRotation.KeyRotationEnabled,
		CloudHsmClusterId:     output.KeyMetadata.CloudHsmClusterId,
		CreationDate:          output.KeyMetadata.CreationDate,
		CustomKeyStoreId:      output.KeyMetadata.CustomKeyStoreId,
		CustomerMasterKeySpec: output.KeyMetadata.CustomerMasterKeySpec,
		DeletionDate:          output.KeyMetadata.DeletionDate,
		Description:           output.KeyMetadata.Description,
		Enabled:               output.KeyMetadata.Enabled,
		ExpirationModel:       output.KeyMetadata.ExpirationModel,
		Manager:               output.KeyMetadata.KeyManager,
		KeyState:              output.KeyMetadata.KeyState,
		KeyUsage:              output.KeyMetadata.KeyUsage,
		Origin:                output.KeyMetadata.Origin,
		ValidTo:               output.KeyMetadata.ValidTo,
	}

	for _, algorithm := range output.KeyMetadata.EncryptionAlgorithms {
		res.EncryptionAlgorithms = append(res.EncryptionAlgorithms, &KeyEncryptionAlgorithm{
			name: aws.StringValue(algorithm),
		})
	}

	for _, algorithm := range output.KeyMetadata.SigningAlgorithms {
		res.SigningAlgorithms = append(res.SigningAlgorithms, &KeySigningAlgorithm{
			name: aws.StringValue(algorithm),
		})
	}

	return &res
}

func (c *Client) transformKeyListEntrys(values []*kms.KeyListEntry) []*Key {
	var tValues []*Key
	for _, v := range values {
		tValues = append(tValues, c.transformKeyListEntry(v))
	}
	return tValues
}

func (c *Client) keys(gConfig interface{}) error {
	var config kms.ListKeysInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["kmsKeys"] {
		err := c.db.AutoMigrate(
			&Key{},
			&KeySigningAlgorithm{},
			&KeyEncryptionAlgorithm{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["kmsKeys"] = true
	}
	for {
		output, err := c.svc.ListKeys(&config)
		if err != nil {
			return err
		}
		c.db.Where("region = ?", c.region).Where("account_id = ?", c.accountID).Delete(&Key{})
		common.ChunkedCreate(c.db, c.transformKeyListEntrys(output.Keys))
		c.log.Info("Fetched resources", zap.String("resource", "kms.keys"), zap.Int("count", len(output.Keys)))
		if aws.StringValue(output.NextMarker) == "" {
			break
		}
		config.Marker = output.NextMarker
	}
	return nil
}
