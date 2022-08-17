package kms

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/cloudkms/v1"
)

type KeyRing struct {
	*cloudkms.KeyRing
	Location string
}

func KmsKeyrings() *schema.Table {
	return &schema.Table{
		Name:        "gcp_kms_keyrings",
		Description: "A KeyRing is a toplevel logical grouping of CryptoKeys.",
		Resolver:    fetchKmsKeyrings,
		Multiplex:   client.ProjectMultiplex,

		PostResourceResolver: client.AddGcpMetadata,
		Options:              schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "location", "name"}},
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "location",
				Description: "Location of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "create_time",
				Description: "The time at which this KeyRing was created",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.DateResolver("CreateTime"),
			},
			{
				Name:        "name",
				Description: "The resource name for the KeyRing in the format `projects/*/locations/*/keyRings/*`",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:                 "gcp_kms_keyring_crypto_keys",
				Description:          "A CryptoKey represents a logical key that can be used for cryptographic operations.",
				Resolver:             fetchKmsKeyringCryptoKeys,
				PostResourceResolver: client.AddGcpMetadata,
				Columns: []schema.Column{
					{
						Name:        "keyring_cq_id",
						Description: "Unique ID of gcp_kms_keyrings table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "project_id",
						Description: "GCP Project Id of the resource",
						Type:        schema.TypeString,
						Resolver:    client.ResolveProject,
					},
					{
						Name:        "location",
						Description: "Location of the resource",
						Type:        schema.TypeString,
						Resolver:    resolveKmsKeyringCryptKeyLocation,
					},
					{
						Name:        "policy",
						Description: "Access control policy for a resource",
						Type:        schema.TypeJSON,
						Resolver:    resolveKmsKeyringCryptoKeyPolicy,
					},
					{
						Name:        "create_time",
						Description: "The time at which this CryptoKey was created",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.DateResolver("CreateTime"),
					},
					{
						Name:        "labels",
						Description: "Labels with user-defined metadata For more information, see Labeling Keys (https://cloudgooglecom/kms/docs/labeling-keys)",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "name",
						Description: "The resource name for this CryptoKey in the format `projects/*/locations/*/keyRings/*/cryptoKeys/*`",
						Type:        schema.TypeString,
					},
					{
						Name:        "next_rotation_time",
						Description: "At next_rotation_time, the Key Management Service will automatically: 1 Create a new version of this CryptoKey 2 Mark the new version as primary Key rotations performed manually via CreateCryptoKeyVersion and UpdateCryptoKeyPrimaryVersion do not affect next_rotation_time Keys with purpose ENCRYPT_DECRYPT support automatic rotation For other keys, this field must be omitted",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.DateResolver("NextRotationTime"),
					},
					{
						Name:        "primary_algorithm",
						Description: "The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports  Possible values:   \"CRYPTO_KEY_VERSION_ALGORITHM_UNSPECIFIED\" - Not specified   \"GOOGLE_SYMMETRIC_ENCRYPTION\" - Creates symmetric encryption keys   \"RSA_SIGN_PSS_2048_SHA256\" - RSASSA-PSS 2048 bit key with a SHA256 digest   \"RSA_SIGN_PSS_3072_SHA256\" - RSASSA-PSS 3072 bit key with a SHA256 digest   \"RSA_SIGN_PSS_4096_SHA256\" - RSASSA-PSS 4096 bit key with a SHA256 digest   \"RSA_SIGN_PSS_4096_SHA512\" - RSASSA-PSS 4096 bit key with a SHA512 digest   \"RSA_SIGN_PKCS1_2048_SHA256\" - RSASSA-PKCS1-v1_5 with a 2048 bit key and a SHA256 digest   \"RSA_SIGN_PKCS1_3072_SHA256\" - RSASSA-PKCS1-v1_5 with a 3072 bit key and a SHA256 digest   \"RSA_SIGN_PKCS1_4096_SHA256\" - RSASSA-PKCS1-v1_5 with a 4096 bit key and a SHA256 digest   \"RSA_SIGN_PKCS1_4096_SHA512\" - RSASSA-PKCS1-v1_5 with a 4096 bit key and a SHA512 digest   \"RSA_DECRYPT_OAEP_2048_SHA256\" - RSAES-OAEP 2048 bit key with a SHA256 digest   \"RSA_DECRYPT_OAEP_3072_SHA256\" - RSAES-OAEP 3072 bit key with a SHA256 digest   \"RSA_DECRYPT_OAEP_4096_SHA256\" - RSAES-OAEP 4096 bit key with a SHA256 digest   \"RSA_DECRYPT_OAEP_4096_SHA512\" - RSAES-OAEP 4096 bit key with a SHA512 digest   \"EC_SIGN_P256_SHA256\" - ECDSA on the NIST P-256 curve with a SHA256 digest   \"EC_SIGN_P384_SHA384\" - ECDSA on the NIST P-384 curve with a SHA384 digest   \"EXTERNAL_SYMMETRIC_ENCRYPTION\" - Algorithm representing symmetric encryption by an external key manager",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.Algorithm"),
					},
					{
						Name:          "primary_attestation_cert_chains_cavium_certs",
						Description:   "Cavium certificate chain corresponding to the attestation",
						Type:          schema.TypeStringArray,
						IgnoreInTests: true,
						Resolver:      schema.PathResolver("Primary.Attestation.CertChains.CaviumCerts"),
					},
					{
						Name:          "primary_attestation_cert_chains_google_card_certs",
						Description:   "Google card certificate chain corresponding to the attestation",
						Type:          schema.TypeStringArray,
						IgnoreInTests: true,
						Resolver:      schema.PathResolver("Primary.Attestation.CertChains.GoogleCardCerts"),
					},
					{
						Name:          "primary_attestation_cert_chains_google_partition_certs",
						Description:   "Google partition certificate chain corresponding to the attestation",
						Type:          schema.TypeStringArray,
						IgnoreInTests: true,
						Resolver:      schema.PathResolver("Primary.Attestation.CertChains.GooglePartitionCerts"),
					},
					{
						Name:        "primary_attestation_content",
						Description: "The attestation data provided by the HSM when the key operation was performed",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.Attestation.Content"),
					},
					{
						Name:        "primary_attestation_format",
						Description: "The format of the attestation data  Possible values:   \"ATTESTATION_FORMAT_UNSPECIFIED\" - Not specified   \"CAVIUM_V1_COMPRESSED\" - Cavium HSM attestation compressed with gzip Note that this format is defined by Cavium and subject to change at any time   \"CAVIUM_V2_COMPRESSED\" - Cavium HSM attestation V2 compressed with gzip This is a new format introduced in Cavium's version 32-08",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.Attestation.Format"),
					},
					{
						Name:        "primary_create_time",
						Description: "The time at which this CryptoKeyVersion was created",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.DateResolver("Primary.CreateTime"),
					},
					{
						Name:          "primary_destroy_event_time",
						Description:   "The time this CryptoKeyVersion's key material was destroyed Only present if state is DESTROYED",
						Type:          schema.TypeTimestamp,
						IgnoreInTests: true,
						Resolver:      schema.DateResolver("Primary.DestroyEventTime"),
					},
					{
						Name:          "primary_destroy_time",
						Description:   "The time this CryptoKeyVersion's key material is scheduled for destruction Only present if state is DESTROY_SCHEDULED",
						Type:          schema.TypeTimestamp,
						IgnoreInTests: true,
						Resolver:      schema.DateResolver("Primary.DestroyTime"),
					},
					{
						Name:        "primary_external_protection_level_options_external_key_uri",
						Description: "The URI for an external resource that this CryptoKeyVersion represents",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.ExternalProtectionLevelOptions.ExternalKeyUri"),
					},
					{
						Name:        "primary_generate_time",
						Description: "The time this CryptoKeyVersion's key material was generated",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.DateResolver("Primary.GenerateTime"),
					},
					{
						Name:        "primary_import_failure_reason",
						Description: "The root cause of an import failure Only present if state is IMPORT_FAILED",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.ImportFailureReason"),
					},
					{
						Name:        "primary_import_job",
						Description: "The name of the ImportJob used to import this CryptoKeyVersion Only present if the underlying key material was imported",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.ImportJob"),
					},
					{
						Name:          "primary_import_time",
						Description:   "The time at which this CryptoKeyVersion's key material was imported",
						Type:          schema.TypeTimestamp,
						IgnoreInTests: true,
						Resolver:      schema.DateResolver("Primary.ImportTime"),
					},
					{
						Name:        "primary_name",
						Description: "The resource name for this CryptoKeyVersion in the format `projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*`",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.Name"),
					},
					{
						Name:        "primary_protection_level",
						Description: "The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.ProtectionLevel"),
					},
					{
						Name:        "primary_state",
						Description: "The current state of the CryptoKeyVersion",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.State"),
					},
					{
						Name:        "purpose",
						Description: "Immutable The immutable purpose of this CryptoKey",
						Type:        schema.TypeString,
					},
					{
						Name:        "rotation_period",
						Description: "next_rotation_time will be advanced by this period when the service automatically rotates a key Must be at least 24 hours and at most 876,000 hours If rotation_period is set, next_rotation_time must also be set Keys with purpose ENCRYPT_DECRYPT support automatic rotation For other keys, this field must be omitted",
						Type:        schema.TypeString,
					},
					{
						Name:        "version_template_algorithm",
						Description: "Algorithm to use when creating a CryptoKeyVersion based on this template",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VersionTemplate.Algorithm"),
					},
					{
						Name:        "version_template_protection_level",
						Description: "ProtectionLevel to use when creating a CryptoKeyVersion based on this template Immutable Defaults to SOFTWARE",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VersionTemplate.ProtectionLevel"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchKmsKeyrings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	locations, err := getAllKmsLocations(ctx, c)
	if err != nil {
		return errors.WithStack(fmt.Errorf("failed to get kms locations. %w", err))
	}
	nextPageToken := ""
	for _, l := range locations {
		call := c.Services.Kms.Projects.Locations.KeyRings.List(l.Name)
		for {
			call.PageToken(nextPageToken)
			resp, err := call.Do()
			if err != nil {
				return errors.WithStack(err)
			}

			rings := make([]*KeyRing, len(resp.KeyRings))
			for i, k := range resp.KeyRings {
				rings[i] = &KeyRing{
					KeyRing:  k,
					Location: l.LocationId,
				}
			}
			res <- rings

			if resp.NextPageToken == "" {
				break
			}
			nextPageToken = resp.NextPageToken
		}
	}
	return nil
}
func fetchKmsKeyringCryptoKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	keyRing := parent.Item.(*KeyRing)
	nextPageToken := ""
	call := c.Services.Kms.Projects.Locations.KeyRings.CryptoKeys.List(keyRing.Name)
	for {
		call.PageToken(nextPageToken)
		resp, err := call.Do()
		if err != nil {
			return errors.WithStack(err)
		}

		res <- resp.CryptoKeys

		if resp.NextPageToken == "" {
			break
		}
		nextPageToken = resp.NextPageToken
	}
	return nil
}
func resolveKmsKeyringCryptoKeyPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	client_ := meta.(*client.Client)
	p := resource.Item.(*cloudkms.CryptoKey)
	resp, err := client_.Services.Kms.Projects.Locations.KeyRings.CryptoKeys.GetIamPolicy(p.Name).Do()
	if err != nil {
		return errors.WithStack(err)
	}

	var policy map[string]interface{}
	data, err := json.Marshal(resp)
	if err != nil {
		return errors.WithStack(err)
	}
	if err := json.Unmarshal(data, &policy); err != nil {
		return errors.WithStack(err)
	}

	return errors.WithStack(resource.Set(c.Name, policy))
}

func resolveKmsKeyringCryptKeyLocation(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	keyRing := resource.Parent.Item.(*KeyRing)
	// CryptoKey location is the same as it's keyring location
	return errors.WithStack(resource.Set(c.Name, keyRing.Location))
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func getAllKmsLocations(ctx context.Context, c *client.Client) ([]*cloudkms.Location, error) {
	var locations []*cloudkms.Location
	call := c.Services.Kms.Projects.Locations.List("projects/" + c.ProjectId)
	nextPageToken := ""
	for {
		resp, err := call.PageToken(nextPageToken).Do()
		if err != nil {
			return nil, errors.WithStack(err)
		}

		locations = append(locations, resp.Locations...)

		if resp.NextPageToken == "" {
			break
		}
		nextPageToken = resp.NextPageToken
	}
	return locations, nil
}
