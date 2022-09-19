package recipes

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DNSRecordResources() []Resource {
	return []Resource{
		{
			DefaultColumns: []codegen.ColumnDefinition{AccountIDColumn}, // ZoneIDColumn is already in the response
			Multiplex:      "client.ZoneMultiplex",

			CFStruct:   &cloudflare.DNSRecord{},
			PrimaryKey: "id",
			Template:   "resource_manual",
			SkipFields: []string{"Meta", "Data"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:        "meta",
					Description: "Extra Cloudflare-specific information about the record.",
					Type:        schema.TypeJSON,
				},
				{
					Name:        "data",
					Description: "Metadata about the record.",
					Type:        schema.TypeJSON,
				},
			},
			TableName:        "cloudflare_dns_records",
			TableFuncName:    "DNSRecords",
			Filename:         "dns_records.go",
			Package:          "dns_records",
			ResolverFuncName: "fetchDNSRecords",
		},
	}
}
