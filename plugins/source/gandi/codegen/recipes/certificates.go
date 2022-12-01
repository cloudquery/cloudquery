package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/go-gandi/go-gandi/certificate"
)

func CertificateResources() []*Resource {
	return []*Resource{
		{
			DataStruct:         &certificate.CertificateType{},
			PKColumns:          []string{"id"},
			ExtraColumns:       []codegen.ColumnDefinition{SharingIDColumn},
			SkipSubserviceName: true,
		},
		{
			DataStruct:   &certificate.Package{},
			PKColumns:    []string{"name"},
			ExtraColumns: []codegen.ColumnDefinition{SharingIDColumn},
		},
	}
}
