package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/resources/services/bss"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/resources/services/oss"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

func Plugin() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"alicloud",
		Version,
		[]*schema.Table{
			bss.BillOverview(),
			bss.Bill(),

			oss.Buckets(),
			oss.BucketStats(),
		},
		client.New,
	)
}
