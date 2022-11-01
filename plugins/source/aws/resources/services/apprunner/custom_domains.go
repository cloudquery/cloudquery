// Code generated by codegen; DO NOT EDIT.

package apprunner

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func CustomDomains() *schema.Table {
	return &schema.Table{
		Name:        "aws_apprunner_custom_domains",
		Description: "https://docs.aws.amazon.com/apprunner/latest/api/API_CustomDomain.html",
		Resolver:    fetchApprunnerCustomDomains,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "domain_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DomainName"),
			},
			{
				Name:     "enable_www_subdomain",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableWWWSubdomain"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "certificate_validation_records",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CertificateValidationRecords"),
			},
		},
	}
}
