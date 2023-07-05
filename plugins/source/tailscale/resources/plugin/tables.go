package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/tailscale/resources/services/acl"
	"github.com/cloudquery/cloudquery/plugins/source/tailscale/resources/services/device"
	"github.com/cloudquery/cloudquery/plugins/source/tailscale/resources/services/dns"
	"github.com/cloudquery/cloudquery/plugins/source/tailscale/resources/services/key"
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/cloudquery/plugin-sdk/v4/docs"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

var customExceptions = map[string]string{
	"acls":        "Access Control Lists (ACLs)",
	"dns":         "Domain Name System (DNS)",
	"nameservers": "Name Servers",
	"searchpaths": "Search Paths",
}

func titleTransformer(table *schema.Table) {
	if table.Title != "" {
		return
	}
	exceptions := make(map[string]string)
	for k, v := range docs.DefaultTitleExceptions {
		exceptions[k] = v
	}
	for k, v := range customExceptions {
		exceptions[k] = v
	}
	csr := caser.New(caser.WithCustomExceptions(exceptions))
	table.Title =  csr.ToTitle(table.Name)
	for _, rel := range table.Relations {
		titleTransformer(rel)
	}
}

func addCqIDs(table *schema.Table) {
	havePks := len(table.PrimaryKeys()) > 0
	cqIdColumn := schema.CqIDColumn
	if !havePks {
		cqIdColumn.PrimaryKey = true
	}
	table.Columns = append(
		schema.ColumnList{
			cqIdColumn,
			schema.CqParentIDColumn,
		},
		table.Columns...,
	)
	for _, rel := range table.Relations {
		addCqIDs(rel)
	}
}

func Tables() schema.Tables {
	tables := schema.Tables{
		acl.Acls(),
		device.Devices(),
		dns.Nameservers(),
		dns.Preferences(),
		dns.Searchpaths(),
		key.Keys(),
	}

	if err := transformers.TransformTables(tables); err != nil {
		panic(err)
	}
	for _, table := range tables {
		addCqIDs(table)
		titleTransformer(table)
	}
	return tables
}
