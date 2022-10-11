package recipes

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/codegen"
)

func CertificatePackResources() []Resource {
	return []Resource{
		{
			DefaultColumns:   []codegen.ColumnDefinition{AccountIDColumn, ZoneIDColumn},
			Multiplex:        "client.ZoneMultiplex",
			CFStruct:         &cloudflare.CertificatePack{},
			PrimaryKey:       "id",
			Template:         "resource_manual",
			TableName:        "cloudflare_certificate_packs",
			TableFuncName:    "CertificatePacks",
			Filename:         "certificate_packs.go",
			Package:          "certificate_packs",
			ResolverFuncName: "fetchCertificatePacks",
		},
	}
}
