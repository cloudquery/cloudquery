package kms

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type Key struct {
	_         interface{} `neo:"raw:MERGE (a:AWSAccount {account_id: $account_id}) MERGE (a) - [:Resource] -> (n)"`
	ID        uint        `gorm:"primarykey"`
	AccountID string
	Region    string

	Arn                   *string `neo:"unique"`
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

func (Key) TableName() string {
	return "aws_kms_keys"
}

type KeyEncryptionAlgorithm struct {
	ID        uint   `gorm:"primarykey"`
	KeyID     uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	name string
}

func (KeyEncryptionAlgorithm) TableName() string {
	return "aws_kms_key_encryption_algorithms"
}

type KeySigningAlgorithm struct {
	ID        uint   `gorm:"primarykey"`
	KeyID     uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	name string
}

func (KeySigningAlgorithm) TableName() string {
	return "aws_kms_key_signing_algorithms"
}

func (c *Client) transformKeyListEntry(value *kms.KeyListEntry) (*Key, error) {
	output, err := c.svc.DescribeKey(&kms.DescribeKeyInput{
		KeyId: value.KeyId,
	})
	if err != nil {
		return nil, err
	}
	outputKeyRotation, err := c.svc.GetKeyRotationStatus(&kms.GetKeyRotationStatusInput{
		KeyId: value.KeyId,
	})
	if err != nil {
		return nil, err
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
			Region:    c.region,
			AccountID: c.accountID,
			name:      aws.StringValue(algorithm),
		})
	}

	for _, algorithm := range output.KeyMetadata.SigningAlgorithms {
		res.SigningAlgorithms = append(res.SigningAlgorithms, &KeySigningAlgorithm{
			Region:    c.region,
			AccountID: c.accountID,
			name:      aws.StringValue(algorithm),
		})
	}

	return &res, nil
}

func (c *Client) transformKeyListEntrys(values []*kms.KeyListEntry) ([]*Key, error) {
	var tValues []*Key
	for _, v := range values {
		tValue, err := c.transformKeyListEntry(v)
		if err != nil {
			return nil, err
		}
		tValues = append(tValues, tValue)
	}
	return tValues, nil
}

var KeyTables = []interface{}{
	&Key{},
	&KeySigningAlgorithm{},
	&KeyEncryptionAlgorithm{},
}

func (c *Client) keys(gConfig interface{}) error {
	var config kms.ListKeysInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(KeyTables...)

	for {
		output, err := c.svc.ListKeys(&config)
		if err != nil {
			return err
		}
		tValues, err := c.transformKeyListEntrys(output.Keys)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(tValues)
		c.log.Info("Fetched resources", zap.String("resource", "kms.keys"), zap.Int("count", len(output.Keys)))
		if aws.StringValue(output.NextMarker) == "" {
			break
		}
		config.Marker = output.NextMarker
	}
	return nil
}
