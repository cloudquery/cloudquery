package route53

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Route53Domains() *schema.Table {
	return &schema.Table{
		Name:          "aws_route53_domains",
		Description:   "The domain names registered with Amazon Route 53.",
		Resolver:      fetchRoute53Domains,
		Multiplex:     client.AccountMultiplex,
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "domain_name"}},
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "admin_contact_address_line1",
				Description: "First line of the contact's address.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AdminContact.AddressLine1"),
			},
			{
				Name:        "admin_contact_address_line2",
				Description: "Second line of contact's address, if any.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AdminContact.AddressLine2"),
			},
			{
				Name:        "admin_contact_city",
				Description: "The city of the contact's address.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AdminContact.City"),
			},
			{
				Name:        "admin_contact_type",
				Description: "Indicates whether the contact is a person, company, association, or public organization.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AdminContact.ContactType"),
			},
			{
				Name:        "admin_contact_country_code",
				Description: "Code for the country of the contact's address.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AdminContact.CountryCode"),
			},
			{
				Name:        "admin_contact_email",
				Description: "Email address of the contact.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AdminContact.Email"),
			},
			{
				Name:        "admin_contact_fax",
				Description: "Fax number of the contact.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AdminContact.Fax"),
			},
			{
				Name:        "admin_contact_first_name",
				Description: "First name of contact.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AdminContact.FirstName"),
			},
			{
				Name:        "admin_contact_last_name",
				Description: "Last name of contact.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AdminContact.LastName"),
			},
			{
				Name:        "admin_contact_organization_name",
				Description: "Name of the organization for contact types other than PERSON.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AdminContact.OrganizationName"),
			},
			{
				Name:        "admin_contact_phone_number",
				Description: "The phone number of the contact.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AdminContact.PhoneNumber"),
			},
			{
				Name:        "admin_contact_state",
				Description: "The state or province of the contact's city.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AdminContact.State"),
			},
			{
				Name:        "admin_contact_zip_code",
				Description: "The zip or postal code of the contact's address.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AdminContact.ZipCode"),
			},
			{
				Name:        "admin_contact_extra_params",
				Description: "A mapping of name to value parameter pairs required by certain top-level domains.",
				Type:        schema.TypeJSON,
				Resolver:    resolveRoute53DomainContactExtraParams(func(d *route53domains.GetDomainDetailOutput) *types.ContactDetail { return d.AdminContact }),
			},
			{
				Name:        "domain_name",
				Description: "The name of a domain.",
				Type:        schema.TypeString,
			},
			{
				Name:        "registrant_contact_address_line1",
				Description: "First line of the contact's address.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RegistrantContact.AddressLine1"),
			},
			{
				Name:        "registrant_contact_address_line2",
				Description: "Second line of contact's address, if any.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RegistrantContact.AddressLine2"),
			},
			{
				Name:        "registrant_contact_city",
				Description: "The city of the contact's address.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RegistrantContact.City"),
			},
			{
				Name:        "registrant_contact_type",
				Description: "Indicates whether the contact is a person, company, association, or public organization.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RegistrantContact.ContactType"),
			},
			{
				Name:        "registrant_contact_country_code",
				Description: "Code for the country of the contact's address.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RegistrantContact.CountryCode"),
			},
			{
				Name:        "registrant_contact_email",
				Description: "Email address of the contact.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RegistrantContact.Email"),
			},
			{
				Name:        "registrant_contact_fax",
				Description: "Fax number of the contact.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RegistrantContact.Fax"),
			},
			{
				Name:        "registrant_contact_first_name",
				Description: "First name of contact.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RegistrantContact.FirstName"),
			},
			{
				Name:        "registrant_contact_last_name",
				Description: "Last name of contact.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RegistrantContact.LastName"),
			},
			{
				Name:        "registrant_contact_organization_name",
				Description: "Name of the organization for contact types other than PERSON.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RegistrantContact.OrganizationName"),
			},
			{
				Name:        "registrant_contact_phone_number",
				Description: "The phone number of the contact.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RegistrantContact.PhoneNumber"),
			},
			{
				Name:        "registrant_contact_state",
				Description: "The state or province of the contact's city.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RegistrantContact.State"),
			},
			{
				Name:        "registrant_contact_zip_code",
				Description: "The zip or postal code of the contact's address.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RegistrantContact.ZipCode"),
			},
			{
				Name:        "registrant_contact_extra_params",
				Description: "A mapping of name to value parameter pairs required by certain top-level domains.",
				Type:        schema.TypeJSON,
				Resolver:    resolveRoute53DomainContactExtraParams(func(d *route53domains.GetDomainDetailOutput) *types.ContactDetail { return d.RegistrantContact }),
			},
			{
				Name:        "tech_contact_address_line1",
				Description: "First line of the contact's address.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TechContact.AddressLine1"),
			},
			{
				Name:        "tech_contact_address_line2",
				Description: "Second line of contact's address, if any.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TechContact.AddressLine2"),
			},
			{
				Name:        "tech_contact_city",
				Description: "The city of the contact's address.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TechContact.City"),
			},
			{
				Name:        "tech_contact_type",
				Description: "Indicates whether the contact is a person, company, association, or public organization.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TechContact.ContactType"),
			},
			{
				Name:        "tech_contact_country_code",
				Description: "Code for the country of the contact's address.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TechContact.CountryCode"),
			},
			{
				Name:        "tech_contact_email",
				Description: "Email address of the contact.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TechContact.Email"),
			},
			{
				Name:        "tech_contact_fax",
				Description: "Fax number of the contact.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TechContact.Fax"),
			},
			{
				Name:        "tech_contact_first_name",
				Description: "First name of contact.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TechContact.FirstName"),
			},
			{
				Name:        "tech_contact_last_name",
				Description: "Last name of contact.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TechContact.LastName"),
			},
			{
				Name:        "tech_contact_organization_name",
				Description: "Name of the organization for contact types other than PERSON.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TechContact.OrganizationName"),
			},
			{
				Name:        "tech_contact_phone_number",
				Description: "The phone number of the contact.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TechContact.PhoneNumber"),
			},
			{
				Name:        "tech_contact_state",
				Description: "The state or province of the contact's city.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TechContact.State"),
			},
			{
				Name:        "tech_contact_zip_code",
				Description: "The zip or postal code of the contact's address.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TechContact.ZipCode"),
			},
			{
				Name:        "tech_contact_extra_params",
				Description: "A mapping of name to value parameter pairs required by certain top-level domains.",
				Type:        schema.TypeJSON,
				Resolver:    resolveRoute53DomainContactExtraParams(func(d *route53domains.GetDomainDetailOutput) *types.ContactDetail { return d.TechContact }),
			},
			{
				Name:        "abuse_contact_email",
				Description: "Email address to contact to report incorrect contact information for a domain, to report that the domain is being used to send spam, to report that someone is cybersquatting on a domain name, or report some other type of abuse.",
				Type:        schema.TypeString,
			},
			{
				Name:        "abuse_contact_phone",
				Description: "Phone number for reporting abuse.",
				Type:        schema.TypeString,
			},
			{
				Name:        "admin_privacy",
				Description: "Specifies whether contact information is concealed from WHOIS queries",
				Type:        schema.TypeBool,
			},
			{
				Name:        "auto_renew",
				Description: "Specifies whether the domain registration is set to renew automatically.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "creation_date",
				Description: "The date when the domain was created as found in the response to a WHOIS query.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "dns_sec",
				Description: "Reserved for future use.",
				Type:        schema.TypeString,
			},
			{
				Name:        "expiration_date",
				Description: "The date when the registration for the domain is set to expire",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "registrant_privacy",
				Description: "Specifies whether contact information is concealed from WHOIS queries",
				Type:        schema.TypeBool,
			},
			{
				Name:        "registrar_name",
				Description: "Name of the registrar of the domain as identified in the registry",
				Type:        schema.TypeString,
			},
			{
				Name:        "registrar_url",
				Description: "Web address of the registrar.",
				Type:        schema.TypeString,
			},
			{
				Name:        "registry_domain_id",
				Description: "Reserved for future use.",
				Type:        schema.TypeString,
			},
			{
				Name:        "reseller",
				Description: "Reseller of the domain",
				Type:        schema.TypeString,
			},
			{
				Name:        "status_list",
				Description: "An array of domain name status codes, also known as Extensible Provisioning Protocol (EPP) status codes",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "tech_privacy",
				Description: "Specifies whether contact information is concealed from WHOIS queries",
				Type:        schema.TypeBool,
			},
			{
				Name:        "updated_date",
				Description: "The last updated date of the domain as found in the response to a WHOIS query.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "who_is_server",
				Description: "The fully qualified name of the WHOIS server that can answer the WHOIS query for the domain.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "A list of tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveRoute53DomainTags,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_route53_domain_nameservers",
				Description: "Nameserver includes the following elements.",
				Resolver:    fetchRoute53DomainNameservers,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"domain_cq_id", "name"}},
				Columns: []schema.Column{
					{
						Name:        "domain_cq_id",
						Description: "Unique CloudQuery ID of aws_route53_domains table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "The fully qualified host name of the name server",
						Type:        schema.TypeString,
					},
					{
						Name:        "glue_ips",
						Description: "Glue IP address of a name server entry",
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
func fetchRoute53Domains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Route53Domains
	var input route53domains.ListDomainsInput
	optsFunc := func(options *route53domains.Options) {
		// Set region to default global region
		options.Region = "us-east-1"
	}
	for {
		output, err := svc.ListDomains(ctx, &input, optsFunc)
		if err != nil {
			return diag.WrapError(err)
		}

		for _, v := range output.Domains {
			d, err := svc.GetDomainDetail(ctx, &route53domains.GetDomainDetailInput{DomainName: v.DomainName}, optsFunc)
			if err != nil {
				return diag.WrapError(err)
			}
			res <- d
		}

		if aws.ToString(output.NextPageMarker) == "" {
			break
		}
		input.Marker = output.NextPageMarker
	}
	return nil
}

func fetchRoute53DomainNameservers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	d, ok := parent.Item.(*route53domains.GetDomainDetailOutput)
	if !ok {
		return fmt.Errorf("not a *route53domains.GetDomainDetailOutput instance: %T", parent.Item)
	}
	res <- d.Nameservers
	return nil
}

func resolveRoute53DomainTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, col schema.Column) error {
	c := meta.(*client.Client)
	svc := c.Services().Route53Domains
	d, ok := resource.Item.(*route53domains.GetDomainDetailOutput)
	if !ok {
		return fmt.Errorf("not a *route53domains.GetDomainDetailOutput instance: %T", resource.Item)
	}
	out, err := svc.ListTagsForDomain(ctx, &route53domains.ListTagsForDomainInput{DomainName: d.DomainName}, func(options *route53domains.Options) {
		// Set region to default global region
		options.Region = "us-east-1"
	})
	if err != nil {
		return diag.WrapError(err)
	}
	tags := make(map[string]string, len(out.TagList))
	for _, v := range out.TagList {
		key := aws.ToString(v.Key)
		if key == "" {
			continue
		}
		tags[key] = aws.ToString(v.Value)
	}
	return resource.Set(col.Name, tags)
}

func resolveRoute53DomainContactExtraParams(extractValue func(*route53domains.GetDomainDetailOutput) *types.ContactDetail) func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, col schema.Column) error {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, col schema.Column) error {
		d, ok := resource.Item.(*route53domains.GetDomainDetailOutput)
		if !ok {
			return fmt.Errorf("not a *route53domains.GetDomainDetailOutput instance: %T", resource.Item)
		}
		detail := extractValue(d)
		if detail == nil {
			return nil
		}
		m := make(map[string]string, len(detail.ExtraParams))
		for _, p := range detail.ExtraParams {
			m[string(p.Name)] = aws.ToString(p.Value)
		}
		b, err := json.Marshal(m)
		if err != nil {
			return diag.WrapError(err)
		}
		return resource.Set(col.Name, b)
	}
}
