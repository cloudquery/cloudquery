package recipes

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DNSRecordResources() []*Resource {
	return []*Resource{
		{
			Multiplex: "client.ZoneMultiplex",

			DataStruct: &cloudflare.DNSRecord{},
			PKColumns:  []string{"id"},
			SkipFields: []string{"Meta", "Data"},
			ExtraColumns: []codegen.ColumnDefinition{
				AccountIDColumn, // ZoneIDColumn is already in the response
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
			Service: "dns_records",
		},
	}
}
