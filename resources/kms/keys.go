package kms

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/mitchellh/mapstructure"
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

func (c *Client) transformKeyListEntrys(values *[]types.KeyListEntry) ([]*Key, error) {
	var tValues []*Key
	ctx := context.Background()
	for _, value := range *values {
		output, err := c.svc.DescribeKey(ctx, &kms.DescribeKeyInput{
			KeyId: value.KeyId,
		})
		if err != nil {
			return nil, err
		}
		var outputKeyRotation *kms.GetKeyRotationStatusOutput
		if string(output.KeyMetadata.Origin) != "EXTERNAL" {
			outputKeyRotation, err = c.svc.GetKeyRotationStatus(ctx, &kms.GetKeyRotationStatusInput{
				KeyId: value.KeyId,
			})
			if err != nil {
				return nil, err
			}
		}

		res := Key{
			Region:                c.region,
			AccountID:             c.accountID,
			Arn:                   value.KeyArn,
			KeyId:                 value.KeyId,
			CloudHsmClusterId:     output.KeyMetadata.CloudHsmClusterId,
			CreationDate:          output.KeyMetadata.CreationDate,
			CustomKeyStoreId:      output.KeyMetadata.CustomKeyStoreId,
			CustomerMasterKeySpec: aws.String(string(output.KeyMetadata.CustomerMasterKeySpec)),
			DeletionDate:          output.KeyMetadata.DeletionDate,
			Description:           output.KeyMetadata.Description,
			Enabled:               &output.KeyMetadata.Enabled,
			ExpirationModel:       aws.String(string(output.KeyMetadata.ExpirationModel)),
			Manager:               aws.String(string(output.KeyMetadata.KeyManager)),
			KeyState:              aws.String(string(output.KeyMetadata.KeyState)),
			KeyUsage:              aws.String(string(output.KeyMetadata.KeyUsage)),
			Origin:                aws.String(string(output.KeyMetadata.Origin)),
			ValidTo:               output.KeyMetadata.ValidTo,
		}

		if outputKeyRotation != nil {
			res.RotationEnabled = &outputKeyRotation.KeyRotationEnabled
		}

		for _, algorithm := range output.KeyMetadata.EncryptionAlgorithms {
			res.EncryptionAlgorithms = append(res.EncryptionAlgorithms, &KeyEncryptionAlgorithm{
				Region:    c.region,
				AccountID: c.accountID,
				name:      string(algorithm),
			})
		}

		for _, algorithm := range output.KeyMetadata.SigningAlgorithms {
			res.SigningAlgorithms = append(res.SigningAlgorithms, &KeySigningAlgorithm{
				Region:    c.region,
				AccountID: c.accountID,
				name:      string(algorithm),
			})
		}
		tValues = append(tValues, &res)
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
	ctx := context.Background()
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(KeyTables...)

	for {
		output, err := c.svc.ListKeys(ctx, &config)
		if err != nil {
			return err
		}
		tValues, err := c.transformKeyListEntrys(&output.Keys)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(tValues)
		c.log.Info("Fetched resources", "resource", "kms.keys", "count", len(output.Keys))
		if aws.ToString(output.NextMarker) == "" {
			break
		}
		config.Marker = output.NextMarker
	}
	return nil
}
