package recipes

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/codegen"
)

func CertificatePackResources() []*Resource {
	return []*Resource{
		{
			ExtraColumns: []codegen.ColumnDefinition{AccountIDColumn, ZoneIDColumn},
			Multiplex:    "client.ZoneMultiplex",
			DataStruct:   &cloudflare.CertificatePack{},
			PKColumns:    []string{"id"},
			Service:      "certificate_packs",
		},
	}
}
