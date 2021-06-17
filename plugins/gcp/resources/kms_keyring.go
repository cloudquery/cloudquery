package resources

import (
	"context"
	"fmt"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	cloudkms "google.golang.org/api/cloudkms/v1"
)

func KmsKeyring() *schema.Table {
	return &schema.Table{
		Name:                 "gcp_kms_keyrings",
		Description:          "A KeyRing is a toplevel logical grouping of CryptoKeys.",
		Resolver:             fetchKmsKeyrings,
		Multiplex:            client.ProjectMultiplex,
		IgnoreError:          client.IgnoreErrorHandler,
		DeleteFilter:         client.DeleteProjectFilter,
		PostResourceResolver: client.AddGcpMetadata,
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
				Type:        schema.TypeString,
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
						Name:        "keyring_id",
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
					},
					{
						Name:        "create_time",
						Description: "The time at which this CryptoKey was created",
						Type:        schema.TypeString,
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
						Type:        schema.TypeString,
					},
					{
						Name:        "primary_algorithm",
						Description: "The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.Algorithm"),
					},
					{
						Name:        "primary_attestation_cert_chains_cavium_certs",
						Description: "Cavium certificate chain corresponding to the attestation",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("Primary.Attestation.CertChains.CaviumCerts"),
					},
					{
						Name:        "primary_attestation_cert_chains_google_card_certs",
						Description: "Google card certificate chain corresponding to the attestation",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("Primary.Attestation.CertChains.GoogleCardCerts"),
					},
					{
						Name:        "primary_attestation_cert_chains_google_partition_certs",
						Description: "Google partition certificate chain corresponding to the attestation",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("Primary.Attestation.CertChains.GooglePartitionCerts"),
					},
					{
						Name:        "primary_attestation_content",
						Description: "The attestation data provided by the HSM when the key operation was performed",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.Attestation.Content"),
					},
					{
						Name:        "primary_attestation_format",
						Description: "The format of the attestation data",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.Attestation.Format"),
					},
					{
						Name:        "primary_create_time",
						Description: "The time at which this CryptoKeyVersion was created",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.CreateTime"),
					},
					{
						Name:        "primary_destroy_event_time",
						Description: "The time this CryptoKeyVersion's key material was destroyed Only present if state is DESTROYED",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.DestroyEventTime"),
					},
					{
						Name:        "primary_destroy_time",
						Description: "The time this CryptoKeyVersion's key material is scheduled for destruction Only present if state is DESTROY_SCHEDULED",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.DestroyTime"),
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
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.GenerateTime"),
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
						Name:        "primary_import_time",
						Description: "The time at which this CryptoKeyVersion's key material was imported",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.ImportTime"),
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
						Description: "Required Algorithm to use when creating a CryptoKeyVersion based on this template For backwards compatibility, GOOGLE_SYMMETRIC_ENCRYPTION is implied if both this field is omitted and CryptoKeypurpose is ENCRYPT_DECRYPT",
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
func fetchKmsKeyrings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	locations, err := getAllKmsLocations(ctx, c.ProjectId, c.Services.Kms)
	if err != nil {
		return fmt.Errorf("failed to get kms locations. %w", err)
	}
	nextPageToken := ""
	for _, l := range locations {
		call := c.Services.Kms.Projects.Locations.KeyRings.List(l.Name).Context(ctx)
		for {
			call.PageToken(nextPageToken)
			resp, err := call.Do()
			if err != nil {
				return err
			}
			res <- resp.KeyRings

			if resp.NextPageToken == "" {
				break
			}
			nextPageToken = resp.NextPageToken
		}
	}
	return nil
}
func fetchKmsKeyringCryptoKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	keyRing := parent.Item.(*cloudkms.KeyRing)
	nextPageToken := ""
	call := c.Services.Kms.Projects.Locations.KeyRings.CryptoKeys.List(keyRing.Name).Context(ctx)
	for {
		call.PageToken(nextPageToken)
		resp, err := call.Do()
		if err != nil {
			return err
		}
		res <- resp.CryptoKeys

		if resp.NextPageToken == "" {
			break
		}
		nextPageToken = resp.NextPageToken
	}
	return nil
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func getAllKmsLocations(ctx context.Context, projectId string, kms *cloudkms.Service) ([]*cloudkms.Location, error) {
	var locations []*cloudkms.Location
	call := kms.Projects.Locations.List("projects/" + projectId).Context(ctx)
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
