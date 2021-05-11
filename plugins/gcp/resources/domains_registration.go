package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	domains "google.golang.org/api/domains/v1beta1"
)

func DomainsRegistration() *schema.Table {
	return &schema.Table{
		Name:         "gcp_domains_registrations",
		Resolver:     fetchDomainsRegistrations,
		Multiplex:    client.ProjectMultiplex,
		DeleteFilter: client.DeleteProjectFilter,
		IgnoreError:  client.IgnoreErrorHandler,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "admin_contact_email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.AdminContact.Email"),
			},
			{
				Name:     "admin_contact_fax_number",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.AdminContact.FaxNumber"),
			},
			{
				Name:     "admin_contact_phone_number",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.AdminContact.PhoneNumber"),
			},
			{
				Name:     "admin_contact_postal_address_address_lines",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ContactSettings.AdminContact.PostalAddress.AddressLines"),
			},
			{
				Name:     "admin_contact_postal_address_administrative_area",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.AdminContact.PostalAddress.AdministrativeArea"),
			},
			{
				Name:     "admin_contact_postal_address_language_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.AdminContact.PostalAddress.LanguageCode"),
			},
			{
				Name:     "admin_contact_postal_address_locality",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.AdminContact.PostalAddress.Locality"),
			},
			{
				Name:     "admin_contact_postal_address_organization",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.AdminContact.PostalAddress.Organization"),
			},
			{
				Name:     "admin_contact_postal_address_postal_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.AdminContact.PostalAddress.PostalCode"),
			},
			{
				Name:     "admin_contact_postal_address_recipients",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ContactSettings.AdminContact.PostalAddress.Recipients"),
			},
			{
				Name:     "admin_contact_postal_address_region_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.AdminContact.PostalAddress.RegionCode"),
			},
			{
				Name:     "admin_contact_postal_address_revision",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ContactSettings.AdminContact.PostalAddress.Revision"),
			},
			{
				Name:     "admin_contact_postal_address_sorting_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.AdminContact.PostalAddress.SortingCode"),
			},
			{
				Name:     "admin_contact_postal_address_sublocality",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.AdminContact.PostalAddress.Sublocality"),
			},
			{
				Name:     "privacy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.Privacy"),
			},
			{
				Name:     "registrant_contact_email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.RegistrantContact.Email"),
			},
			{
				Name:     "registrant_contact_fax_number",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.RegistrantContact.FaxNumber"),
			},
			{
				Name:     "registrant_contact_phone_number",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.RegistrantContact.PhoneNumber"),
			},
			{
				Name:     "registrant_contact_postal_address_address_lines",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ContactSettings.RegistrantContact.PostalAddress.AddressLines"),
			},
			{
				Name:     "registrant_contact_postal_address_administrative_area",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.RegistrantContact.PostalAddress.AdministrativeArea"),
			},
			{
				Name:     "registrant_contact_postal_address_language_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.RegistrantContact.PostalAddress.LanguageCode"),
			},
			{
				Name:     "registrant_contact_postal_address_locality",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.RegistrantContact.PostalAddress.Locality"),
			},
			{
				Name:     "registrant_contact_postal_address_organization",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.RegistrantContact.PostalAddress.Organization"),
			},
			{
				Name:     "registrant_contact_postal_address_postal_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.RegistrantContact.PostalAddress.PostalCode"),
			},
			{
				Name:     "registrant_contact_postal_address_recipients",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ContactSettings.RegistrantContact.PostalAddress.Recipients"),
			},
			{
				Name:     "registrant_contact_postal_address_region_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.RegistrantContact.PostalAddress.RegionCode"),
			},
			{
				Name:     "registrant_contact_postal_address_revision",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ContactSettings.RegistrantContact.PostalAddress.Revision"),
			},
			{
				Name:     "registrant_contact_postal_address_sorting_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.RegistrantContact.PostalAddress.SortingCode"),
			},
			{
				Name:     "registrant_contact_postal_address_sublocality",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.RegistrantContact.PostalAddress.Sublocality"),
			},
			{
				Name:     "technical_contact_email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.TechnicalContact.Email"),
			},
			{
				Name:     "technical_contact_fax_number",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.TechnicalContact.FaxNumber"),
			},
			{
				Name:     "technical_contact_phone_number",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.TechnicalContact.PhoneNumber"),
			},
			{
				Name:     "technical_contact_postal_address_address_lines",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ContactSettings.TechnicalContact.PostalAddress.AddressLines"),
			},
			{
				Name:     "technical_contact_postal_address_administrative_area",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.TechnicalContact.PostalAddress.AdministrativeArea"),
			},
			{
				Name:     "technical_contact_postal_address_language_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.TechnicalContact.PostalAddress.LanguageCode"),
			},
			{
				Name:     "technical_contact_postal_address_locality",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.TechnicalContact.PostalAddress.Locality"),
			},
			{
				Name:     "technical_contact_postal_address_organization",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.TechnicalContact.PostalAddress.Organization"),
			},
			{
				Name:     "technical_contact_postal_address_postal_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.TechnicalContact.PostalAddress.PostalCode"),
			},
			{
				Name:     "technical_contact_postal_address_recipients",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ContactSettings.TechnicalContact.PostalAddress.Recipients"),
			},
			{
				Name:     "technical_contact_postal_address_region_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.TechnicalContact.PostalAddress.RegionCode"),
			},
			{
				Name:     "technical_contact_postal_address_revision",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ContactSettings.TechnicalContact.PostalAddress.Revision"),
			},
			{
				Name:     "technical_contact_postal_address_sorting_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.TechnicalContact.PostalAddress.SortingCode"),
			},
			{
				Name:     "technical_contact_postal_address_sublocality",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactSettings.TechnicalContact.PostalAddress.Sublocality"),
			},
			{
				Name: "create_time",
				Type: schema.TypeString,
			},
			{
				Name:     "custom_dns_ds_records",
				Type:     schema.TypeJSON,
				Resolver: resolveDomainsRegistrationCustomDNSDsRecords,
			},
			{
				Name:     "custom_dns_name_servers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("DnsSettings.CustomDns.NameServers"),
			},
			{
				Name:     "google_domains_dns_ds_records",
				Type:     schema.TypeJSON,
				Resolver: resolveDomainsRegistrationGoogleDomainsDNSDsRecords,
			},
			{
				Name:     "google_domains_dns_ds_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DnsSettings.GoogleDomainsDns.DsState"),
			},
			{
				Name:     "google_domains_dns_name_servers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("DnsSettings.GoogleDomainsDns.NameServers"),
			},
			{
				Name: "domain_name",
				Type: schema.TypeString,
			},
			{
				Name: "expire_time",
				Type: schema.TypeString,
			},
			{
				Name: "issues",
				Type: schema.TypeStringArray,
			},
			{
				Name: "labels",
				Type: schema.TypeJSON,
			},
			{
				Name:     "management_settings_renewal_method",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ManagementSettings.RenewalMethod"),
			},
			{
				Name:     "management_settings_transfer_lock_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ManagementSettings.TransferLockState"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "state",
				Type: schema.TypeString,
			},
			{
				Name: "supported_privacy",
				Type: schema.TypeStringArray,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "gcp_domains_registration_glue_records",
				Resolver: fetchDomainsRegistrationGlueRecords,
				Columns: []schema.Column{
					{
						Name:     "registration_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "host_name",
						Type: schema.TypeString,
					},
					{
						Name: "ipv4_addresses",
						Type: schema.TypeStringArray,
					},
					{
						Name: "ipv6_addresses",
						Type: schema.TypeStringArray,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchDomainsRegistrations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Domain.Projects.Locations.Registrations.List("projects/" + c.ProjectId + "/location/-").Context(ctx)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}
		res <- output.Registrations
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func resolveDomainsRegistrationCustomDNSDsRecords(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	reg := resource.Item.(*domains.Registration)
	if reg.DnsSettings == nil || reg.DnsSettings.CustomDns == nil {
		return nil
	}
	data, err := json.Marshal(reg.DnsSettings.CustomDns.DsRecords)
	if err != nil {
		return fmt.Errorf("failed to marshal custom_dns_ds_records. %w", err)
	}
	return resource.Set("custom_dns_ds_records", data)
}

func resolveDomainsRegistrationGoogleDomainsDNSDsRecords(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	reg := resource.Item.(*domains.Registration)
	data, err := json.Marshal(reg.DnsSettings.GoogleDomainsDns.DsRecords)
	if err != nil {
		return fmt.Errorf("failed to marshal google_domains_dns_ds_records. %w", err)
	}
	return resource.Set("google_domains_dns_ds_records", data)
}

func fetchDomainsRegistrationGlueRecords(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	reg := parent.Item.(*domains.Registration)
	if reg.DnsSettings != nil {
		res <- reg.DnsSettings.GlueRecords
	}
	return nil
}
