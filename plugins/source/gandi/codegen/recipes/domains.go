package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/go-gandi/go-gandi/domain"
)

func DomainResources() []Resource {
	return []Resource{
		{
			DataStruct:          &domain.Details{},
			PKColumns:           []string{"id"},
			Relations:           []string{"DomainLiveDNS()", "DomainWebRedirections()", "DomainGlueRecords()", "DomainDNSSecKeys()"},
			TableName:           "gandi_domains",
			PreResourceResolver: "getDomain",
			Template:            "resource_manual",
			Package:             "domains",
			TableFuncName:       "Domains",
			Filename:            "domains.go",
			ResolverFuncName:    "fetchDomains",
		},
		{
			DataStruct: &domain.LiveDNS{},
			PKColumns:  []string{"fqdn"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "fqdn",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("fqdn")`,
				},
			},
			TableName:        "gandi_domain_live_dns",
			Template:         "resource_manual",
			Package:          "domains",
			Filename:         "domain_live_dns.go",
			TableFuncName:    "DomainLiveDNS",
			ResolverFuncName: "fetchDomainLiveDNS",
		},
		{
			DataStruct: &domain.DNSSECKey{},
			PKColumns:  []string{"fqdn", "id"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "fqdn",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("fqdn")`,
				},
			},
			TableName:        "gandi_domain_dnssec_keys",
			Template:         "resource_manual",
			Package:          "domains",
			Filename:         "domain_dnssec_keys.go",
			TableFuncName:    "DomainDNSSecKeys",
			ResolverFuncName: "fetchDomainDNSSecKeys",
		},
		{
			DataStruct: &domain.GlueRecord{},
			PKColumns:  []string{"domain_fqdn", "name"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "domain_fqdn",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("fqdn")`,
				},
			},
			TableName:        "gandi_domain_glue_records",
			Template:         "resource_manual",
			Package:          "domains",
			Filename:         "domain_glue_records.go",
			TableFuncName:    "DomainGlueRecords",
			ResolverFuncName: "fetchDomainGlueRecords",
		},
		{
			DataStruct: &domain.WebRedirection{},
			PKColumns:  []string{"domain_fqdn", "host", "type"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "domain_fqdn",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("fqdn")`,
				},
			},
			TableName:        "gandi_domain_web_redirections",
			Template:         "resource_manual",
			Package:          "domains",
			Filename:         "domain_web_redirections.go",
			TableFuncName:    "DomainWebRedirections",
			ResolverFuncName: "fetchDomainWebRedirections",
		},
	}
}
