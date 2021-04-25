package resources

import (
	"context"
	"fmt"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/cloudkms/v1"
)

func KmsKeyring() *schema.Table {
	return &schema.Table{
		Name:                 "gcp_kms_keyrings",
		Resolver:             fetchKmsKeyrings,
		Multiplex:            client.ProjectMultiplex,
		DeleteFilter:         client.DeleteProjectFilter,
		IgnoreError:          client.IgnoreErrorHandler,
		PostResourceResolver: client.AddGcpMetadata,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name: "location",
				Type: schema.TypeString,
			},
			{
				Name: "create_time",
				Type: schema.TypeString,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:                 "gcp_kms_keyring_crypto_keys",
				Resolver:             fetchKmsKeyringCryptoKeys,
				IgnoreError:          client.IgnoreErrorHandler,
				PostResourceResolver: client.AddGcpMetadata,
				Columns: []schema.Column{
					{
						Name:     "keyring_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "project_id",
						Type:     schema.TypeString,
						Resolver: client.ResolveProject,
					},
					{
						Name: "location",
						Type: schema.TypeString,
					},
					{
						Name: "create_time",
						Type: schema.TypeString,
					},
					{
						Name: "labels",
						Type: schema.TypeJSON,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "next_rotation_time",
						Type: schema.TypeString,
					},
					{
						Name:     "primary_algorithm",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Primary.Algorithm"),
					},
					{
						Name:     "primary_attestation_cert_chains_cavium_certs",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("Primary.Attestation.CertChains.CaviumCerts"),
					},
					{
						Name:     "primary_attestation_cert_chains_google_card_certs",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("Primary.Attestation.CertChains.GoogleCardCerts"),
					},
					{
						Name:     "primary_attestation_cert_chains_google_partition_certs",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("Primary.Attestation.CertChains.GooglePartitionCerts"),
					},
					{
						Name:     "primary_attestation_content",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Primary.Attestation.Content"),
					},
					{
						Name:     "primary_attestation_format",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Primary.Attestation.Format"),
					},
					{
						Name:     "primary_create_time",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Primary.CreateTime"),
					},
					{
						Name:     "primary_destroy_event_time",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Primary.DestroyEventTime"),
					},
					{
						Name:     "primary_destroy_time",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Primary.DestroyTime"),
					},
					{
						Name:     "primary_external_protection_level_options_external_key_uri",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Primary.ExternalProtectionLevelOptions.ExternalKeyUri"),
					},
					{
						Name:     "primary_generate_time",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Primary.GenerateTime"),
					},
					{
						Name:     "primary_import_failure_reason",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Primary.ImportFailureReason"),
					},
					{
						Name:     "primary_import_job",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Primary.ImportJob"),
					},
					{
						Name:     "primary_import_time",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Primary.ImportTime"),
					},
					{
						Name:     "primary_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Primary.Name"),
					},
					{
						Name:     "primary_protection_level",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Primary.ProtectionLevel"),
					},
					{
						Name:     "primary_state",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Primary.State"),
					},
					{
						Name: "purpose",
						Type: schema.TypeString,
					},
					{
						Name: "rotation_period",
						Type: schema.TypeString,
					},
					{
						Name:     "version_template_algorithm",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("VersionTemplate.Algorithm"),
					},
					{
						Name:     "version_template_protection_level",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("VersionTemplate.ProtectionLevel"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchKmsKeyrings(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
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
