package kms

import (
	"context"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/cloudkms/v1"
	"sync/atomic"
)

type CryptoKey struct {
	ID               uint `gorm:"primarykey"`
	ProjectID        string
	CreateTime       string
	Name             string
	NextRotationTime string

	PrimaryAlgorithm string

	CaviumCerts          []*CryptoKeyCaviumCert          `gorm:"constraint:OnDelete:CASCADE;"`
	GoogleCardCerts      []*CryptoKeyGoogleCardCert      `gorm:"constraint:OnDelete:CASCADE;"`
	GooglePartitionCerts []*CryptoKeyGooglePartitionCert `gorm:"constraint:OnDelete:CASCADE;"`

	PrimaryAttestationContent string
	PrimaryAttestationFormat  string

	PrimaryCreateTime       string
	PrimaryDestroyEventTime string
	PrimaryDestroyTime      string

	PrimaryExternalProtectionLevelOptionsExternalKeyUri string

	PrimaryGenerateTime        string
	PrimaryImportFailureReason string
	PrimaryImportJob           string
	PrimaryImportTime          string
	PrimaryName                string
	PrimaryProtectionLevel     string
	PrimaryState               string

	Purpose        string
	RotationPeriod string

	VersionTemplateAlgorithm       string
	VersionTemplateProtectionLevel string
}

func (CryptoKey) TableName() string {
	return "gcp_kms_crypto_keys"
}

type CryptoKeyCaviumCert struct {
	ID                  uint `gorm:"primarykey"`
	CryptoKeyID uint
	Value                        string
}

func (CryptoKeyCaviumCert) TableName() string {
	return "gcp_kms_key_cavium_certs"
}

type CryptoKeyGoogleCardCert struct {
	ID                  uint `gorm:"primarykey"`
	CryptoKeyID uint
	Value                        string
}

func (CryptoKeyGoogleCardCert) TableName() string {
	return "gcp_kms_key_card_certs"
}

type CryptoKeyGooglePartitionCert struct {
	ID                  uint `gorm:"primarykey"`
	CryptoKeyID uint
	Value                        string
}

func (CryptoKeyGooglePartitionCert) TableName() string {
	return "gcp_kms_key_google_partition_certs"
}

func (c *Client) transformCryptoKeys(values []*cloudkms.CryptoKey) []*CryptoKey {
	var tValues []*CryptoKey
	for _, value := range values {
		tValue := CryptoKey{
			ProjectID:        c.projectID,
			CreateTime:       value.CreateTime,
			Name:             value.Name,
			NextRotationTime: value.NextRotationTime,
			Purpose:          value.Purpose,
			RotationPeriod:   value.RotationPeriod,
		}
		if value.Primary != nil {

			tValue.PrimaryAlgorithm = value.Primary.Algorithm
			tValue.PrimaryCreateTime = value.Primary.CreateTime
			tValue.PrimaryDestroyEventTime = value.Primary.DestroyEventTime
			tValue.PrimaryDestroyTime = value.Primary.DestroyTime
			tValue.PrimaryGenerateTime = value.Primary.GenerateTime
			tValue.PrimaryImportFailureReason = value.Primary.ImportFailureReason
			tValue.PrimaryImportJob = value.Primary.ImportJob
			tValue.PrimaryImportTime = value.Primary.ImportTime
			tValue.PrimaryName = value.Primary.Name
			tValue.PrimaryProtectionLevel = value.Primary.ProtectionLevel
			tValue.PrimaryState = value.Primary.State

		}
		if value.VersionTemplate != nil {

			tValue.VersionTemplateAlgorithm = value.VersionTemplate.Algorithm
			tValue.VersionTemplateProtectionLevel = value.VersionTemplate.ProtectionLevel

		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}
func (c *Client) transformCryptoKeyCertificateChainsCaviumCerts(values []string) []*CryptoKeyCaviumCert {
	var tValues []*CryptoKeyCaviumCert
	for _, v := range values {
		tValues = append(tValues, &CryptoKeyCaviumCert{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformCryptoKeyCertificateChainsGoogleCardCerts(values []string) []*CryptoKeyGoogleCardCert {
	var tValues []*CryptoKeyGoogleCardCert
	for _, v := range values {
		tValues = append(tValues, &CryptoKeyGoogleCardCert{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformCryptoKeyCertificateChainsGooglePartitionCerts(values []string) []*CryptoKeyGooglePartitionCert {
	var tValues []*CryptoKeyGooglePartitionCert
	for _, v := range values {
		tValues = append(tValues, &CryptoKeyGooglePartitionCert{
			Value: v,
		})
	}
	return tValues
}

var CryptoKeyTables = []interface{}{
	&CryptoKey{},
	&CryptoKeyCaviumCert{},
	&CryptoKeyGoogleCardCert{},
	&CryptoKeyGooglePartitionCert{},
}

func (c *Client) cryptoKeys(_ interface{}) error {
	c.db.Where("project_id", c.projectID).Delete(CryptoKeyTables...)
	locations, err := c.getAllKmsLocations()
	if err != nil {
		return err
	}
	var counter uint32
	g, _ := errgroup.WithContext(context.Background())
	for _, l := range locations {
		l := l
		g.Go(func() error {
			keyrings, err := c.getLocationKeyRings(l)
			if err != nil {
				return err
			}
			for _, key := range keyrings {
				count, err := c.getCryptoKeys(key)
				if err != nil {
					return err
				}
				atomic.AddUint32(&counter, count)
			}
			return nil
		})
	}

	err = g.Wait()
	if err != nil {
		return err
	}
	c.log.Info("Fetched resources", "resource", "kms.keys", "count", counter)
	return nil
}

func (c *Client) getCryptoKeys(key *cloudkms.KeyRing) (uint32, error) {
	nextPageToken := ""
	var count uint32 = 0
	for {
		call := c.svc.Projects.Locations.KeyRings.CryptoKeys.List(key.Name)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return count, err
		}

		c.db.ChunkedCreate(c.transformCryptoKeys(output.CryptoKeys))
		count += uint32(len(output.CryptoKeys))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return count, nil
}

func (c *Client) getLocationKeyRings(location *cloudkms.Location) ([]*cloudkms.KeyRing, error) {
	var keyrings []*cloudkms.KeyRing
	call := c.svc.Projects.Locations.KeyRings.List(location.Name)
	nextPageToken := ""
	for {
		call.PageToken(nextPageToken)
		resp, err := call.Do()
		if err != nil {
			return nil, err
		}
		keyrings = append(keyrings, resp.KeyRings...)

		if resp.NextPageToken == "" {
			break
		}
		nextPageToken = resp.NextPageToken
	}
	return keyrings, nil
}

func (c *Client) getAllKmsLocations() ([]*cloudkms.Location, error) {

	var locations []*cloudkms.Location
	call := c.svc.Projects.Locations.List("projects/" + c.projectID)
	nextPageToken := ""
	for {
		call.PageToken(nextPageToken)
		resp, err := call.Do()
		if err != nil {
			return nil, err
		}
		locations = append(locations, resp.Locations...)

		if resp.NextPageToken == "" {
			break
		}
		nextPageToken = resp.NextPageToken
	}
	return locations, nil
}
