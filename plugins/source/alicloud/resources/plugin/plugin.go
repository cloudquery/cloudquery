package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/resources/services/bss"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/resources/services/ecs"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/resources/services/oss"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"alicloud",
		Version,
		// Note:  this list should only include top-level tables
		[]*schema.Table{
			bss.BillOverview(),
			bss.Bill(),
			bss.BillDetails(),
			ecs.Instances(),
			oss.Buckets(),
		},
		client.New,
	)
}
