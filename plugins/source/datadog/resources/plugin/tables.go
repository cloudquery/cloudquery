package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/datadog/resources/services/users"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Tables() []*schema.Table {
	return []*schema.Table{
		users.Users(),
	}
}
