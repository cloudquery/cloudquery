package domains

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	domains "google.golang.org/api/domains/v1beta1"
)

func DomainsRegistration() *schema.Table {
	return &schema.Table{
		Name:        "gcp_domains_registrations",
		Description: "The `Registration` resource facilitates managing and configuring domain name registrations To create a new `Registration` resource, find a suitable domain name by calling the `SearchDomains` method with a query to see available domain name options After choosing a name, call `RetrieveRegisterParameters` to ensure availability and obtain information like pricing, which is needed to build a call to `RegisterDomain`",
		Resolver:    fetchDomainsRegistrations,
		Multiplex:   client.ProjectMultiplex,

		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "name"}},
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "admin_contact_email",
				Description: "Required Email address of the contact",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.AdminContact.Email"),
			},
			{
				Name:        "admin_contact_fax_number",
				Description: "Fax number of the contact in international format For example, \"+1-800-555-0123\"",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.AdminContact.FaxNumber"),
			},
			{
				Name:        "admin_contact_phone_number",
				Description: "Required Phone number of the contact in international format For example, \"+1-800-555-0123\"",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.AdminContact.PhoneNumber"),
			},
			{
				Name:        "admin_contact_postal_address_address_lines",
				Description: "Unstructured address lines describing the lower levels of an address",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("ContactSettings.AdminContact.PostalAddress.AddressLines"),
			},
			{
				Name:        "admin_contact_postal_address_administrative_area",
				Description: "Optional Highest administrative subdivision which is used for postal addresses of a country or region",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.AdminContact.PostalAddress.AdministrativeArea"),
			},
			{
				Name:        "admin_contact_postal_address_language_code",
				Description: "Optional BCP-47 language code of the contents of this address (if known)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.AdminContact.PostalAddress.LanguageCode"),
			},
			{
				Name:        "admin_contact_postal_address_locality",
				Description: "Optional Generally refers to the city/town portion of the address",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.AdminContact.PostalAddress.Locality"),
			},
			{
				Name:        "admin_contact_postal_address_organization",
				Description: "The name of the organization at the address",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.AdminContact.PostalAddress.Organization"),
			},
			{
				Name:        "admin_contact_postal_address_postal_code",
				Description: "Postal code of the address Not all countries use or require postal codes to be present",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.AdminContact.PostalAddress.PostalCode"),
			},
			{
				Name:        "admin_contact_postal_address_recipients",
				Description: "The recipient at the address This field may, under certain circumstances, contain multiline information For example, it might contain \"care of\" information",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("ContactSettings.AdminContact.PostalAddress.Recipients"),
			},
			{
				Name:        "admin_contact_postal_address_region_code",
				Description: "Required CLDR region code of the country/region of the address This is never inferred and it is up to the user to ensure the value is correct See http://cldrunicodeorg/ and http://wwwunicodeorg/cldr/charts/30/supplemental/territory_informationhtml for details Example: \"CH\" for Switzerland",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.AdminContact.PostalAddress.RegionCode"),
			},
			{
				Name:        "admin_contact_postal_address_revision",
				Description: "The schema revision of the `PostalAddress` This must be set to 0, which is the latest revision All new revisions **must** be backward compatible with old revisions",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ContactSettings.AdminContact.PostalAddress.Revision"),
			},
			{
				Name:        "admin_contact_postal_address_sorting_code",
				Description: "Optional Additional, country-specific, sorting code",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.AdminContact.PostalAddress.SortingCode"),
			},
			{
				Name:        "admin_contact_postal_address_sublocality",
				Description: "Sublocality of the address For example, this can be neighborhoods, boroughs, districts",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.AdminContact.PostalAddress.Sublocality"),
			},
			{
				Name:        "privacy",
				Description: "Required Privacy setting for the contacts associated with the `Registration`  Possible values:   \"CONTACT_PRIVACY_UNSPECIFIED\" - The contact privacy settings are undefined   \"PUBLIC_CONTACT_DATA\" - All the data from `ContactSettings` is publicly available When setting this option, you must also provide a `PUBLIC_CONTACT_DATA_ACKNOWLEDGEMENT` in the `contact_notices` field of the request   \"PRIVATE_CONTACT_DATA\" - None of the data from `ContactSettings` is publicly available Instead, proxy contact data is published for your domain Email sent to the proxy email address is forwarded to the registrant's email address Cloud Domains provides this privacy proxy service at no additional cost   \"REDACTED_CONTACT_DATA\" - Some data from `ContactSettings` is publicly available The actual information redacted depends on the domain For details, see [the registration privacy article](https://supportgooglecom/domains/answer/3251242)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.Privacy"),
			},
			{
				Name:        "registrant_contact_email",
				Description: "Required Email address of the contact",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.RegistrantContact.Email"),
			},
			{
				Name:        "registrant_contact_fax_number",
				Description: "Fax number of the contact in international format For example, \"+1-800-555-0123\"",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.RegistrantContact.FaxNumber"),
			},
			{
				Name:        "registrant_contact_phone_number",
				Description: "Required Phone number of the contact in international format For example, \"+1-800-555-0123\"",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.RegistrantContact.PhoneNumber"),
			},
			{
				Name:        "registrant_contact_postal_address_address_lines",
				Description: "Unstructured address lines describing the lower levels of an address",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("ContactSettings.RegistrantContact.PostalAddress.AddressLines"),
			},
			{
				Name:        "registrant_contact_postal_address_administrative_area",
				Description: "Optional Highest administrative subdivision which is used for postal addresses of a country or region",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.RegistrantContact.PostalAddress.AdministrativeArea"),
			},
			{
				Name:        "registrant_contact_postal_address_language_code",
				Description: "Optional BCP-47 language code of the contents of this address (if known)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.RegistrantContact.PostalAddress.LanguageCode"),
			},
			{
				Name:        "registrant_contact_postal_address_locality",
				Description: "Optional Generally refers to the city/town portion of the address",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.RegistrantContact.PostalAddress.Locality"),
			},
			{
				Name:        "registrant_contact_postal_address_organization",
				Description: "The name of the organization at the address",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.RegistrantContact.PostalAddress.Organization"),
			},
			{
				Name:        "registrant_contact_postal_address_postal_code",
				Description: "Postal code of the address Not all countries use or require postal codes to be present",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.RegistrantContact.PostalAddress.PostalCode"),
			},
			{
				Name:        "registrant_contact_postal_address_recipients",
				Description: "The recipient at the address This field may, under certain circumstances, contain multiline information For example, it might contain \"care of\" information",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("ContactSettings.RegistrantContact.PostalAddress.Recipients"),
			},
			{
				Name:        "registrant_contact_postal_address_region_code",
				Description: "Required CLDR region code of the country/region of the address This is never inferred and it is up to the user to ensure the value is correct See http://cldrunicodeorg/ and http://wwwunicodeorg/cldr/charts/30/supplemental/territory_informationhtml for details Example: \"CH\" for Switzerland",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.RegistrantContact.PostalAddress.RegionCode"),
			},
			{
				Name:        "registrant_contact_postal_address_revision",
				Description: "The schema revision of the `PostalAddress` This must be set to 0, which is the latest revision All new revisions **must** be backward compatible with old revisions",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ContactSettings.RegistrantContact.PostalAddress.Revision"),
			},
			{
				Name:        "registrant_contact_postal_address_sorting_code",
				Description: "Optional Additional, country-specific, sorting code",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.RegistrantContact.PostalAddress.SortingCode"),
			},
			{
				Name:        "registrant_contact_postal_address_sublocality",
				Description: "Sublocality of the address For example, this can be neighborhoods, boroughs, districts",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.RegistrantContact.PostalAddress.Sublocality"),
			},
			{
				Name:        "technical_contact_email",
				Description: "Required Email address of the contact",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.TechnicalContact.Email"),
			},
			{
				Name:        "technical_contact_fax_number",
				Description: "Fax number of the contact in international format For example, \"+1-800-555-0123\"",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.TechnicalContact.FaxNumber"),
			},
			{
				Name:        "technical_contact_phone_number",
				Description: "Required Phone number of the contact in international format For example, \"+1-800-555-0123\"",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.TechnicalContact.PhoneNumber"),
			},
			{
				Name:        "technical_contact_postal_address_address_lines",
				Description: "Unstructured address lines describing the lower levels of an address",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("ContactSettings.TechnicalContact.PostalAddress.AddressLines"),
			},
			{
				Name:        "technical_contact_postal_address_administrative_area",
				Description: "Optional Highest administrative subdivision which is used for postal addresses of a country or region",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.TechnicalContact.PostalAddress.AdministrativeArea"),
			},
			{
				Name:        "technical_contact_postal_address_language_code",
				Description: "Optional BCP-47 language code of the contents of this address (if known)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.TechnicalContact.PostalAddress.LanguageCode"),
			},
			{
				Name:        "technical_contact_postal_address_locality",
				Description: "Generally refers to the city/town portion of the address Examples: US city, IT comune, UK post town In regions of the world where localities are not well defined or do not fit into this structure well, leave locality empty and use address_lines",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.TechnicalContact.PostalAddress.Locality"),
			},
			{
				Name:        "technical_contact_postal_address_organization",
				Description: "The name of the organization at the address",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.TechnicalContact.PostalAddress.Organization"),
			},
			{
				Name:        "technical_contact_postal_address_postal_code",
				Description: "Postal code of the address Not all countries use or require postal codes to be present",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.TechnicalContact.PostalAddress.PostalCode"),
			},
			{
				Name:        "technical_contact_postal_address_recipients",
				Description: "The recipient at the address This field may, under certain circumstances, contain multiline information For example, it might contain \"care of\" information",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("ContactSettings.TechnicalContact.PostalAddress.Recipients"),
			},
			{
				Name:        "technical_contact_postal_address_region_code",
				Description: "Required CLDR region code of the country/region of the address This is never inferred and it is up to the user to ensure the value is correct See http://cldrunicodeorg/ and http://wwwunicodeorg/cldr/charts/30/supplemental/territory_informationhtml for details Example: \"CH\" for Switzerland",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.TechnicalContact.PostalAddress.RegionCode"),
			},
			{
				Name:        "technical_contact_postal_address_revision",
				Description: "The schema revision of the `PostalAddress` This must be set to 0, which is the latest revision All new revisions **must** be backward compatible with old revisions",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ContactSettings.TechnicalContact.PostalAddress.Revision"),
			},
			{
				Name:        "technical_contact_postal_address_sorting_code",
				Description: "Optional Additional, country-specific, sorting code",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.TechnicalContact.PostalAddress.SortingCode"),
			},
			{
				Name:        "technical_contact_postal_address_sublocality",
				Description: "Sublocality of the address For example, this can be neighborhoods, boroughs, districts",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactSettings.TechnicalContact.PostalAddress.Sublocality"),
			},
			{
				Name:        "create_time",
				Description: "The creation timestamp of the `Registration` resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "custom_dns_ds_records",
				Description: "The list of DS records for this domain, which are used to enable DNSSEC The domain's DNS provider can provide the values to set here If this field is empty, DNSSEC is disabled",
				Type:        schema.TypeJSON,
				Resolver:    resolveDomainsRegistrationCustomDNSDsRecords,
			},
			{
				Name:        "custom_dns_name_servers",
				Description: "Required A list of name servers that store the DNS zone for this domain Each name server is a domain name, with Unicode domain names expressed in Punycode format",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("DnsSettings.CustomDns.NameServers"),
			},
			{
				Name:        "google_domains_dns_ds_records",
				Description: "The list of DS records published for this domain The list is automatically populated when `ds_state` is `DS_RECORDS_PUBLISHED`, otherwise it remains empty",
				Type:        schema.TypeJSON,
				Resolver:    resolveDomainsRegistrationGoogleDomainsDNSDsRecords,
			},
			{
				Name:        "google_domains_dns_ds_state",
				Description: "Required The state of DS records for this domain Used to enable or disable automatic DNSSEC  Possible values:   \"DS_STATE_UNSPECIFIED\" - DS state is unspecified   \"DS_RECORDS_UNPUBLISHED\" - DNSSEC is disabled for this domain No DS records for this domain are published in the parent DNS zone   \"DS_RECORDS_PUBLISHED\" - DNSSEC is enabled for this domain Appropriate DS records for this domain are published in the parent DNS zone This option is valid only if the DNS zone referenced in the `Registration`'s `dns_provider` field is already DNSSEC-signed",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DnsSettings.GoogleDomainsDns.DsState"),
			},
			{
				Name:        "google_domains_dns_name_servers",
				Description: "A list of name servers that store the DNS zone for this domain Each name server is a domain name, with Unicode domain names expressed in Punycode format This field is automatically populated with the name servers assigned to the Google Domains DNS zone",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("DnsSettings.GoogleDomainsDns.NameServers"),
			},
			{
				Name:        "domain_name",
				Description: "Required Immutable The domain name Unicode domain names must be expressed in Punycode format",
				Type:        schema.TypeString,
			},
			{
				Name:        "expire_time",
				Description: "The expiration timestamp of the `Registration`",
				Type:        schema.TypeString,
			},
			{
				Name:        "issues",
				Description: "The set of issues with the `Registration` that require attention  Possible values:   \"ISSUE_UNSPECIFIED\" - The issue is undefined   \"CONTACT_SUPPORT\" - Contact the Cloud Support team to resolve a problem with this domain   \"UNVERIFIED_EMAIL\" - [ICANN](https://icannorg/) requires verification of the email address in the `Registration`'s `contact_settingsregistrant_contact` field To verify the email address, follow the instructions in the email the `registrant_contact` receives following registration If you do not complete email verification within 15 days of registration, the domain is suspended To resend the verification email, call ConfigureContactSettings and provide the current `registrant_contactemail`",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "labels",
				Description: "Set of labels associated with the `Registration`",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "management_settings_renewal_method",
				Description: "The renewal method for this `Registration`  Possible values:   \"RENEWAL_METHOD_UNSPECIFIED\" - The renewal method is undefined   \"AUTOMATIC_RENEWAL\" - The domain is automatically renewed each year  To disable automatic renewals, export the domain by calling `ExportRegistration`    \"MANUAL_RENEWAL\" - The domain must be explicitly renewed each year before its `expire_time` This option is only available when the `Registration` is in state `EXPORTED` To manage the domain's current billing and renewal settings, go to [Google Domains](https://domainsgoogle/)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagementSettings.RenewalMethod"),
			},
			{
				Name:        "management_settings_transfer_lock_state",
				Description: "Controls whether the domain can be transferred to another registrar  Possible values:   \"TRANSFER_LOCK_STATE_UNSPECIFIED\" - The state is unspecified   \"UNLOCKED\" - The domain is unlocked and can be transferred to another registrar   \"LOCKED\" - The domain is locked and cannot be transferred to another registrar",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagementSettings.TransferLockState"),
			},
			{
				Name:        "name",
				Description: "Name of the `Registration` resource, in the format `projects/*/locations/*/registrations/`",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "The state of the `Registration`  Possible values:   \"STATE_UNSPECIFIED\" - The state is undefined   \"REGISTRATION_PENDING\" - The domain is being registered   \"REGISTRATION_FAILED\" - The domain registration failed You can delete resources in this state to allow registration to be retried   \"ACTIVE\" - The domain is registered and operational The domain renews automatically as long as it remains in this state   \"SUSPENDED\" - The domain is suspended and inoperative For more details, see the `issues` field   \"EXPORTED\" - The domain has been exported from Cloud Domains to [Google Domains](https://domainsgoogle/) You can no longer update it with this API, and information shown about it may be stale Without further action, domains in this state expire at their `expire_time` You can delete the resource after the `expire_time` has passed",
				Type:        schema.TypeString,
			},
			{
				Name:        "supported_privacy",
				Description: "Set of options for the `contact_settingsprivacy` field that this `Registration` supports  Possible values:   \"CONTACT_PRIVACY_UNSPECIFIED\" - The contact privacy settings are undefined   \"PUBLIC_CONTACT_DATA\" - All the data from `ContactSettings` is publicly available When setting this option, you must also provide a `PUBLIC_CONTACT_DATA_ACKNOWLEDGEMENT` in the `contact_notices` field of the request   \"PRIVATE_CONTACT_DATA\" - None of the data from `ContactSettings` is publicly available Instead, proxy contact data is published for your domain Email sent to the proxy email address is forwarded to the registrant's email address Cloud Domains provides this privacy proxy service at no additional cost   \"REDACTED_CONTACT_DATA\" - Some data from `ContactSettings` is publicly available The actual information redacted depends on the domain For details, see [the registration privacy article](https://supportgooglecom/domains/answer/3251242)",
				Type:        schema.TypeStringArray,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "gcp_domains_registration_glue_records",
				Description: "Defines a host on your domain that is a DNS name server for your domain and/or other domains Glue records are a way of making the IP address of a name server known, even when it serves DNS queries for its parent domain For example, when `nsexamplecom` is a name server for `examplecom`, the host `nsexamplecom` must have a glue record to break the circular DNS reference",
				Resolver:    fetchDomainsRegistrationGlueRecords,
				Columns: []schema.Column{
					{
						Name:        "registration_cq_id",
						Description: "Unique ID of gcp_domains_registrations table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "registration_name",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("name"),
					},
					{
						Name:        "host_name",
						Description: "Required Domain name of the host in Punycode format",
						Type:        schema.TypeString,
					},
					{
						Name:        "ipv4_addresses",
						Description: "List of IPv4 addresses corresponding to this host in the standard decimal format (eg `198511001`) At least one of `ipv4_address` and `ipv6_address` must be set",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "ipv6_addresses",
						Description: "List of IPv6 addresses corresponding to this host in the standard hexadecimal format (eg `2001:db8::`) At least one of `ipv4_address` and `ipv6_address` must be set",
						Type:        schema.TypeStringArray,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchDomainsRegistrations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Domain.Projects.Locations.Registrations.List("projects/" + c.ProjectId + "/locations/-").PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		res <- output.Registrations
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func resolveDomainsRegistrationCustomDNSDsRecords(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	reg := resource.Item.(*domains.Registration)
	if reg.DnsSettings == nil || reg.DnsSettings.CustomDns == nil {
		return nil
	}
	data, err := json.Marshal(reg.DnsSettings.CustomDns.DsRecords)
	if err != nil {
		return errors.WithStack(fmt.Errorf("failed to marshal custom_dns_ds_records. %w", err))
	}
	return errors.WithStack(resource.Set(c.Name, data))
}
func resolveDomainsRegistrationGoogleDomainsDNSDsRecords(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	reg := resource.Item.(*domains.Registration)
	if reg.DnsSettings == nil || reg.DnsSettings.GoogleDomainsDns == nil {
		return nil
	}
	data, err := json.Marshal(reg.DnsSettings.GoogleDomainsDns.DsRecords)
	if err != nil {
		return errors.WithStack(fmt.Errorf("failed to marshal google_domains_dns_ds_records. %w", err))
	}
	return errors.WithStack(resource.Set("google_domains_dns_ds_records", data))
}
func fetchDomainsRegistrationGlueRecords(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	reg := parent.Item.(*domains.Registration)
	if reg.DnsSettings != nil {
		res <- reg.DnsSettings.GlueRecords
	}
	return nil
}
