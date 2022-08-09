package services

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cq-provider-cloudflare/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource certificate_packs --config certificate_packs.hcl --output .
func CertificatePacks() *schema.Table {
	return &schema.Table{
		Name:        "cloudflare_certificate_packs",
		Description: "CertificatePack is the overarching structure of a certificate pack response.",
		Resolver:    fetchCertificatePacks,
		Multiplex:   client.ZoneMultiplex,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAccountId,
			},
			{
				Name:        "zone_id",
				Description: "The Zone ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveZoneId,
			},
			{
				Name:        "id",
				Description: "The unique identifier for a certificate_pack",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "type",
				Description: "Type of certificate pack",
				Type:        schema.TypeString,
			},
			{
				Name:        "hosts",
				Description: "comma separated list of valid host names for the certificate packs. Must contain the zone apex, may not contain more than 50 hosts, and may not be empty.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "primary_certificate",
				Description: "Identifier of the primary certificate in a pack",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "cloudflare_certificate_pack_certificates",
				Description: "CertificatePackCertificate is the base structure of a TLS certificate that is contained within a certificate pack.",
				Resolver:    fetchCertificatePackCertificates,
				Columns: []schema.Column{
					{
						Name:        "certificate_pack_cq_id",
						Description: "Unique CloudQuery ID of cloudflare_certificate_packs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ID"),
					},
					{
						Name: "hosts",
						Type: schema.TypeStringArray,
					},
					{
						Name: "issuer",
						Type: schema.TypeString,
					},
					{
						Name: "signature",
						Type: schema.TypeString,
					},
					{
						Name: "status",
						Type: schema.TypeString,
					},
					{
						Name: "bundle_method",
						Type: schema.TypeString,
					},
					{
						Name:     "geo_restrictions_label",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("GeoRestrictions.Label"),
					},
					{
						Name:     "zone_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ZoneID"),
					},
					{
						Name: "uploaded_on",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "modified_on",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "expires_on",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "priority",
						Type: schema.TypeBigInt,
					},
				},
			},
			{
				Name:        "cloudflare_certificate_pack_validation_records",
				Description: "SSLValidationRecord displays Domain Control Validation tokens.",
				Resolver:    fetchCertificatePackValidationRecords,
				Columns: []schema.Column{
					{
						Name:        "certificate_pack_cq_id",
						Description: "Unique CloudQuery ID of cloudflare_certificate_packs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name: "cname_target",
						Type: schema.TypeString,
					},
					{
						Name: "cname_name",
						Type: schema.TypeString,
					},
					{
						Name: "txt_name",
						Type: schema.TypeString,
					},
					{
						Name: "txt_value",
						Type: schema.TypeString,
					},
					{
						Name:     "http_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("HTTPUrl"),
					},
					{
						Name:     "http_body",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("HTTPBody"),
					},
					{
						Name: "emails",
						Type: schema.TypeStringArray,
					},
				},
			},
			{
				Name:        "cloudflare_certificate_pack_validation_errors",
				Description: "SSLValidationError represents errors that occurred during SSL validation.",
				Resolver:    fetchCertificatePackValidationErrors,
				Columns: []schema.Column{
					{
						Name:        "certificate_pack_cq_id",
						Description: "Unique CloudQuery ID of cloudflare_certificate_packs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name: "message",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchCertificatePacks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	zoneId := svc.ZoneId

	packs, err := svc.ClientApi.ListCertificatePacks(ctx, zoneId)
	if err != nil {
		return diag.WrapError(err)
	}
	res <- packs
	return nil
}
func fetchCertificatePackCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(cloudflare.CertificatePack)
	res <- r.Certificates
	return nil
}
func fetchCertificatePackValidationRecords(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(cloudflare.CertificatePack)
	res <- r.ValidationRecords
	return nil
}
func fetchCertificatePackValidationErrors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(cloudflare.CertificatePack)
	res <- r.ValidationErrors
	return nil
}
