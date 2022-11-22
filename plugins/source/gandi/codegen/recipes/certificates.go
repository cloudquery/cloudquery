package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/go-gandi/go-gandi/certificate"
)

func CertificateResources() []Resource {
	return []Resource{
		{
			DataStruct:   &certificate.CertificateType{},
			PKColumns:    []string{"id"},
			ExtraColumns: []codegen.ColumnDefinition{SharingIDColumn},
			TableName:    "gandi_certificates",

			Template:         "resource_manual",
			Package:          "certificates",
			TableFuncName:    "Certificates",
			Filename:         "certificates.go",
			ResolverFuncName: "fetchCertificates",
		},
		{
			DataStruct:   &certificate.Package{},
			PKColumns:    []string{"name"},
			ExtraColumns: []codegen.ColumnDefinition{SharingIDColumn},
			TableName:    "gandi_certificate_packages",

			Template:         "resource_manual",
			Package:          "certificates",
			TableFuncName:    "Packages",
			Filename:         "packages.go",
			ResolverFuncName: "fetchPackages",
		},
	}
}
