// Code generated by codegen; DO NOT EDIT.

package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/fastly/resources/services/services"
	"github.com/cloudquery/cloudquery/plugins/source/fastly/resources/services/stats"
	"github.com/cloudquery/plugin-sdk/schema"
)

func tables() []*schema.Table {
	return []*schema.Table{
		services.Services(),
		stats.StatsRegions(),
	}
}
