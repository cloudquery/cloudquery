package models

type RegionalConfig struct {
	EbsEncryptionEnabledByDefault bool
	EbsDefaultKmsKeyId            *string
}
